package sql

import (
	"context"

	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/pop/v5"

	"github.com/ory/keto/persistence"
)

type (
	Persister struct {
		conn *pop.Connection
		mb   pop.MigrationBox
	}
)

var (
	migrations = packr.New("migrations", "migrations")

	_ persistence.Persister = &Persister{}
)

func NewPersister(c *pop.Connection) (*Persister, error) {
	mb, err := pop.NewMigrationBox(migrations, c)
	if err != nil {
		return nil, err
	}
	return &Persister{
		mb:   mb,
		conn: c,
	}, nil
}

func (p *Persister) MigrateUp(_ context.Context) error {
	return p.mb.Up()
}
