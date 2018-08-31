/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package role

import (
	"database/sql"

	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ory/herodot"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"github.com/rubenv/sql-migrate"
)

var migrations = &migrate.MemoryMigrationSource{
	Migrations: []*migrate.Migration{
		{
			Id: "1",
			Up: []string{`CREATE TABLE IF NOT EXISTS keto_role (
	id      	varchar(255) NOT NULL PRIMARY KEY
)`, `CREATE TABLE IF NOT EXISTS keto_role_member (
	member		varchar(255) NOT NULL,
	role_id		varchar(255) NOT NULL,
	FOREIGN KEY (role_id) REFERENCES keto_role(id) ON DELETE CASCADE,
	PRIMARY KEY (member, role_id)
)`},
			Down: []string{
				"DROP TABLE keto_role",
				"DROP TABLE keto_role_member",
			},
		},
	},
}

type SQLManager struct {
	DB *sqlx.DB

	TableRole      string
	TableMember    string
	TableMigration string
}

func NewSQLManager(db *sqlx.DB) *SQLManager {
	return &SQLManager{
		DB:             db,
		TableRole:      "keto_role",
		TableMember:    "keto_role_member",
		TableMigration: "keto_role_migration",
	}
}

func (m *SQLManager) CreateSchemas() (int, error) {
	migrate.SetTable(m.TableMigration)
	n, err := migrate.Exec(m.DB.DB, m.DB.DriverName(), migrations, migrate.Up)
	if err != nil {
		return 0, errors.Wrapf(err, "Could not migrate sql schema, applied %d migrations", n)
	}
	return n, nil
}

func (m *SQLManager) createRole(role string) func(tx *sqlx.Tx) error {
	return func(tx *sqlx.Tx) error {
		_, err := tx.Exec(m.DB.Rebind(fmt.Sprintf("INSERT INTO %s (id) VALUES (?)", m.TableRole)), role)

		return errors.WithStack(err)
	}
}

func (m *SQLManager) CreateRole(r *Role) error {
	if r.ID == "" {
		r.ID = uuid.New()
	}

	return m.applyInTransaction(m.createRole(r.ID), m.addRoleMembers(r.ID, r.Members))
}

func (m *SQLManager) GetRole(id string) (*Role, error) {
	var found string
	if err := m.DB.Get(&found, m.DB.Rebind(fmt.Sprintf("SELECT id from %s WHERE id = ?", m.TableRole)), id); err != nil {
		return nil, errors.WithStack(err)
	}

	var q []string
	if err := m.DB.Select(&q, m.DB.Rebind(fmt.Sprintf("SELECT member from %s WHERE role_id = ?", m.TableMember)), found); err == sql.ErrNoRows {
		return nil, errors.WithStack(&herodot.ErrorNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Role{
		ID:      found,
		Members: q,
	}, nil
}

func (m *SQLManager) deleteRole(id string) func(tx *sqlx.Tx) error {
	return func(tx *sqlx.Tx) error {
		_, err := tx.Exec(m.DB.Rebind(fmt.Sprintf("DELETE FROM %s WHERE id=?", m.TableRole)), id)

		return errors.WithStack(err)
	}
}

func (m *SQLManager) DeleteRole(id string) error {
	return m.applyInTransaction(m.deleteRole(id))
}

func (m *SQLManager) addRoleMembers(role string, subjects []string) func(tx *sqlx.Tx) error {
	return func(tx *sqlx.Tx) error {
		query := fmt.Sprintf("INSERT INTO %s (role_id, member) VALUES (?, ?)", m.TableMember)

		for _, subject := range subjects {
			if _, err := tx.Exec(m.DB.Rebind(query), role, subject); err != nil {
				return errors.WithStack(err)
			}
		}

		return nil
	}
}

func (m *SQLManager) AddRoleMembers(role string, subjects []string) error {
	return m.applyInTransaction(m.addRoleMembers(role, subjects))
}

func (m *SQLManager) removeGroupMembers(role string, subjects []string) func(tx *sqlx.Tx) error {
	return func(tx *sqlx.Tx) error {
		query := fmt.Sprintf("DELETE FROM %s WHERE member=? AND role_id=?", m.TableMember)

		for _, subject := range subjects {
			if _, err := tx.Exec(m.DB.Rebind(query), subject, role); err != nil {
				return errors.WithStack(err)
			}
		}

		return nil
	}
}

func (m *SQLManager) RemoveRoleMembers(role string, subjects []string) error {
	return m.applyInTransaction(m.removeGroupMembers(role, subjects))
}

func (m *SQLManager) FindRolesByMember(member string, limit, offset int) ([]Role, error) {
	var ids []string
	if err := m.DB.Select(&ids, m.DB.Rebind(fmt.Sprintf("SELECT role_id from %s WHERE member = ? GROUP BY role_id ORDER BY role_id LIMIT ? OFFSET ?", m.TableMember)), member, limit, offset); err == sql.ErrNoRows {
		return nil, errors.WithStack(&herodot.ErrorNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	var roles = make([]Role, len(ids))
	for k, id := range ids {
		role, err := m.GetRole(id)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		roles[k] = *role
	}

	return roles, nil
}

func (m *SQLManager) ListRoles(limit, offset int) ([]Role, error) {
	var ids []string
	if err := m.DB.Select(&ids, m.DB.Rebind(fmt.Sprintf("SELECT id from %s LIMIT ? OFFSET ?", m.TableRole)), limit, offset); err == sql.ErrNoRows {
		return nil, errors.WithStack(&herodot.ErrorNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	var roles = make([]Role, len(ids))
	for k, id := range ids {
		role, err := m.GetRole(id)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		roles[k] = *role
	}

	return roles, nil
}

func (m *SQLManager) UpdateRole(role Role) error {
	return m.applyInTransaction(m.deleteRole(role.ID), m.createRole(role.ID), m.addRoleMembers(role.ID, role.Members))
}

func (m *SQLManager) applyInTransaction(executors ...func(tx *sqlx.Tx) error) error {
	tx, err := m.DB.Beginx()
	if err != nil {
		return errors.Wrap(err, "Could not begin transaction")
	}

	for _, exec := range executors {
		if err := exec(tx); err != nil {
			if err := tx.Rollback(); err != nil {
				return errors.WithStack(err)
			}

			return err
		}
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.WithStack(err)
		}
		return errors.Wrap(err, "Could not commit transaction")
	}

	return nil
}
