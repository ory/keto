package storage

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/open-policy-agent/opa/storage"
	"github.com/ory/sqlcon"
	"github.com/pkg/errors"
	"github.com/rubenv/sql-migrate"
)

type sqlItem struct {
	Key        string `db:"key"`
	Collection string `db:"collection"`
	Data       string `db:"data"`
}

var Migrations = map[string]*migrate.MemoryMigrationSource{
	"mysql": {
		Migrations: []*migrate.Migration{
			{
				Id: "1",
				Up: []string{
					`CREATE TABLE IF NOT EXISTS rego_data (
    id 					INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
	collection 			VARCHAR(64) NOT NULL,
	` + "`key`" + ` 	VARCHAR(64) NOT NULL,
	` + "`data`" + ` 	JSON,
	UNIQUE KEY rego_data_uidx_ck (` + "`collection`" + `, ` + "`key`" + `)
)`,
				},
				Down: []string{
					"DROP TABLE rego_data",
				},
			},
		},
	},
	"postgres": {
		Migrations: []*migrate.Migration{
			{
				Id: "1",
				Up: []string{
					`CREATE TABLE IF NOT EXISTS rego_data (
    id 			SERIAL PRIMARY KEY,
	collection 	VARCHAR(64) NOT NULL,
	key 		VARCHAR(64) NOT NULL,
	data		JSON
)`,
					`CREATE UNIQUE INDEX rego_data_uidx_ck ON rego_data (collection, key)`,
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
	database := db.DriverName()
	switch database {
	case "pgx", "pq":
		database = "postgres"
	}

	migrate.SetTable("keto_storage_migration")
	n, err := migrate.Exec(db.DB, db.DriverName(), Migrations[database], migrate.Up)
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
	switch database := m.db.DriverName(); database {
	case "mysql":
		query = "INSERT INTO rego_data (`key`, `collection`, `data`) VALUES (:key, :collection, :data) ON DUPLICATE KEY UPDATE `data`=:data"
	case "pgx", "pq", "postgres":
		query = `INSERT INTO rego_data (key, collection, data) VALUES (:key, :collection, :data) ON CONFLICT(collection, key) DO UPDATE SET data = :data`
	default:
		return errors.Errorf("unknown database driver: %s", database)
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
	if err := m.db.SelectContext(
		ctx,
		&items,
		m.db.Rebind("SELECT data FROM rego_data WHERE collection=? ORDER BY id ASC LIMIT ? OFFSET ?"), collection, limit, offset,
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
	var query string
	switch database := m.db.DriverName(); database {
	case "mysql":
		query = "SELECT data FROM rego_data WHERE `collection`=? AND `key`=?"
	case "pgx", "pq", "postgres":
		query = `SELECT data FROM rego_data WHERE collection=? AND key=?`
	default:
		return errors.Errorf("unknown database driver: %s", database)
	}

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
	var query string
	switch database := m.db.DriverName(); database {
	case "mysql":
		query = "DELETE FROM rego_data WHERE `key`=:key AND collection=:collection"
	case "pgx", "pq", "postgres":
		query = "DELETE FROM rego_data WHERE key=:key AND collection=:collection"
	default:
		return errors.Errorf("unknown database driver: %s", database)
	}

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
			m.db.Rebind("SELECT data FROM rego_data WHERE collection=?"), s,
		); err != nil {
			return nil, errors.WithStack(err)
		}
		return items, nil
	})
}
