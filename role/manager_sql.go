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
	"github.com/ory/hydra/pkg"
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
				"DROP TABLE hydra_warden_group",
				"DROP TABLE hydra_warden_group_member",
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

func (m *SQLManager) CreateRole(g *Role) error {
	if g.ID == "" {
		g.ID = uuid.New()
	}
	if _, err := m.DB.Exec(m.DB.Rebind(fmt.Sprintf("INSERT INTO %s (id) VALUES (?)", m.TableRole)), g.ID); err != nil {
		return errors.WithStack(err)
	}

	return m.AddRoleMembers(g.ID, g.Members)
}

func (m *SQLManager) GetRole(id string) (*Role, error) {
	var found string
	if err := m.DB.Get(&found, m.DB.Rebind(fmt.Sprintf("SELECT id from %s WHERE id = ?", m.TableRole)), id); err != nil {
		return nil, errors.WithStack(err)
	}

	var q []string
	if err := m.DB.Select(&q, m.DB.Rebind(fmt.Sprintf("SELECT member from %s WHERE role_id = ?", m.TableMember)), found); err == sql.ErrNoRows {
		return nil, errors.WithStack(pkg.ErrNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Role{
		ID:      found,
		Members: q,
	}, nil
}

func (m *SQLManager) DeleteRole(id string) error {
	if _, err := m.DB.Exec(m.DB.Rebind(fmt.Sprintf("DELETE FROM %s WHERE id=?", m.TableRole)), id); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (m *SQLManager) AddRoleMembers(group string, subjects []string) error {
	tx, err := m.DB.Beginx()
	if err != nil {
		return errors.Wrap(err, "Could not begin transaction")
	}

	query := fmt.Sprintf("INSERT INTO %s (role_id, member) VALUES (?, ?)", m.TableMember)
	for _, subject := range subjects {
		if _, err := tx.Exec(m.DB.Rebind(query), group, subject); err != nil {
			if err := tx.Rollback(); err != nil {
				return errors.WithStack(err)
			}
			return errors.WithStack(err)
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

func (m *SQLManager) RemoveRoleMembers(group string, subjects []string) error {
	tx, err := m.DB.Beginx()
	if err != nil {
		return errors.Wrap(err, "Could not begin transaction")
	}
	for _, subject := range subjects {
		if _, err := m.DB.Exec(m.DB.Rebind(fmt.Sprintf("DELETE FROM %s WHERE member=? AND role_id=?", m.TableMember)), subject, group); err != nil {
			if err := tx.Rollback(); err != nil {
				return errors.WithStack(err)
			}
			return errors.WithStack(err)
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

func (m *SQLManager) FindRolesByMember(member string, limit, offset int) ([]Role, error) {
	var ids []string
	if err := m.DB.Select(&ids, m.DB.Rebind(fmt.Sprintf("SELECT role_id from %s WHERE member = ? GROUP BY role_id ORDER BY role_id LIMIT ? OFFSET ?", m.TableMember)), member, limit, offset); err == sql.ErrNoRows {
		return nil, errors.WithStack(pkg.ErrNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	var groups = make([]Role, len(ids))
	for k, id := range ids {
		group, err := m.GetRole(id)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		groups[k] = *group
	}

	return groups, nil
}

func (m *SQLManager) ListRoles(limit, offset int) ([]Role, error) {
	var ids []string
	if err := m.DB.Select(&ids, m.DB.Rebind(fmt.Sprintf("SELECT role_id from %s GROUP BY role_id ORDER BY role_id LIMIT ? OFFSET ?", m.TableMember)), limit, offset); err == sql.ErrNoRows {
		return nil, errors.WithStack(pkg.ErrNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	var groups = make([]Role, len(ids))
	for k, id := range ids {
		group, err := m.GetRole(id)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		groups[k] = *group
	}

	return groups, nil
}
