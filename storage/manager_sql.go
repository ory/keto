package storage

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/open-policy-agent/opa/storage"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"

	"github.com/ory/x/dbal"
	"github.com/ory/x/sqlcon"
)

type sqlItem struct {
	Key        string `db:"pkey"`
	Collection string `db:"collection"`
	Data       string `db:"document"`
}

var Migrations = map[string]*migrate.MemoryMigrationSource{
	dbal.DriverMySQL: {
		Migrations: []*migrate.Migration{
			{
				Id: "1",
				Up: []string{
					`CREATE TABLE IF NOT EXISTS rego_data (
    id 					INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
	collection 			VARCHAR(64) NOT NULL,
	pkey 				VARCHAR(64) NOT NULL,
	document		 	JSON,
	UNIQUE KEY rego_data_uidx_ck (collection, pkey)
)`,
				},
				Down: []string{
					"DROP TABLE rego_data",
				},
			},
		},
	},
	dbal.DriverPostgreSQL: {
		Migrations: []*migrate.Migration{
			{
				Id: "1",
				Up: []string{
					`CREATE TABLE IF NOT EXISTS rego_data (
    id 			SERIAL PRIMARY KEY,
	collection 	VARCHAR(64) NOT NULL,
	pkey 		VARCHAR(64) NOT NULL,
	document	JSON
)`,
					`CREATE UNIQUE INDEX rego_data_uidx_ck ON rego_data (collection, pkey)`,
				},
				Down: []string{
					"DROP TABLE rego_data",
				},
			},
		},
	},
}

type SQLManager struct {
	db *sqlx.DB
}

func NewSQLManager(db *sqlx.DB) *SQLManager {
	return &SQLManager{
		db: db,
	}
}

func (m *SQLManager) CreateSchemas(db *sqlx.DB) (int, error) {
	migrate.SetTable("keto_storage_migration")
	n, err := migrate.Exec(db.DB, dbal.Canonicalize(m.db.DriverName()), Migrations[dbal.MustCanonicalize(db.DriverName())], migrate.Up)
	if err != nil {
		return 0, errors.Wrapf(err, "could not migrate sql schema completely, applied only %d migrations", n)
	}
	return n, nil
}

func (m *SQLManager) Upsert(ctx context.Context, collection, key string, value interface{}) error {
	b := bytes.NewBuffer(nil)
	if err := json.NewEncoder(b).Encode(value); err != nil {
		return errors.WithStack(err)
	}

	var query string
	switch database := dbal.Canonicalize(m.db.DriverName()); database {
	case dbal.DriverMySQL:
		query = "INSERT INTO rego_data (pkey, collection, document) VALUES (:pkey, :collection, :document) ON DUPLICATE KEY UPDATE document=:document"
	case dbal.DriverPostgreSQL:
		query = `INSERT INTO rego_data (pkey, collection, document) VALUES (:pkey, :collection, :document) ON CONFLICT(collection, pkey) DO UPDATE SET document = :document`
	default:
		return errors.Errorf("unknown database driver: %s", m.db.DriverName())
	}

	if _, err := m.db.NamedExecContext(ctx, query, &sqlItem{
		Key:        key,
		Collection: collection,
		Data:       b.String(),
	}); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (m *SQLManager) List(ctx context.Context, collection string, value interface{}, limit, offset int) error {
	var items []string
	query := "SELECT document FROM rego_data WHERE collection=? ORDER BY id ASC LIMIT ? OFFSET ?"
	if err := m.db.SelectContext(
		ctx,
		&items,
		m.db.Rebind(query), collection, limit, offset,
	); err != nil {
		return sqlcon.HandleError(err)
	}

	ji := make([]json.RawMessage, len(items))
	for k, v := range items {
		ji[k] = json.RawMessage(v)
	}

	return roundTrip(&ji, value)
}

func (m *SQLManager) Get(ctx context.Context, collection, key string, value interface{}) error {
	query := "SELECT document FROM rego_data WHERE collection=? AND pkey=?"
	var item string
	if err := m.db.GetContext(
		ctx,
		&item,
		m.db.Rebind(query), collection, key,
	); err != nil {
		return sqlcon.HandleError(err)
	}

	ji := json.RawMessage(item)
	return roundTrip(&ji, value)
}

func (m *SQLManager) Delete(ctx context.Context, collection, key string) error {
	query := "DELETE FROM rego_data WHERE pkey=:pkey AND collection=:collection"
	if _, err := m.db.NamedExecContext(ctx, query, &sqlItem{
		Key:        key,
		Collection: collection,
	}); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (m *SQLManager) Storage(ctx context.Context, schema string, collections []string) (storage.Store, error) {
	return toRegoStore(ctx, schema, collections, func(i context.Context, s string) ([]json.RawMessage, error) {
		var items []json.RawMessage
		if err := m.db.SelectContext(
			ctx,
			&items,
			m.db.Rebind("SELECT document FROM rego_data WHERE collection=? ORDER BY id ASC"), s,
		); err != nil {
			return nil, errors.WithStack(err)
		}
		return items, nil
	})
}
