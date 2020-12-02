package sql

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ory/keto/internal/namespace"

	"github.com/markbates/pkger"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/pkgerx"

	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/x"
)

type (
	Persister struct {
		conn       *pop.Connection
		mb         *pkgerx.MigrationBox
		namespaces namespace.Manager
	}
	internalPagination struct {
		Offset uint
		Limit  uint
	}
	contextKeys string
)

const (
	pageTokenEnd                      = "no other page"
	transactionContextKey contextKeys = "ongoing transaction"
	defaultPageSize       uint        = 100
)

var (
	migrations = pkger.Dir("/internal/persistence/sql/migrations")

	_ persistence.Persister = &Persister{}
)

func NewPersister(c *pop.Connection, l *logrusx.Logger, namespaces namespace.Manager) (*Persister, error) {
	mb, err := pkgerx.NewMigrationBox(migrations, c, l)
	if err != nil {
		return nil, err
	}
	return &Persister{
		mb:         mb,
		conn:       c,
		namespaces: namespaces,
	}, nil
}

func (p *Persister) MigrateUp(_ context.Context) error {
	return p.mb.Up()
}

func (p *Persister) connection(ctx context.Context) *pop.Connection {
	tx := ctx.Value(transactionContextKey)
	if tx == nil {
		return p.conn
	}
	return tx.(*pop.Connection)
}

func (p *Persister) transaction(ctx context.Context, f func(context.Context, *pop.Connection) error) error {
	tx := ctx.Value(transactionContextKey)
	if tx != nil {
		return f(ctx, tx.(*pop.Connection))
	}

	return p.conn.Transaction(func(tx *pop.Connection) error {
		return f(context.WithValue(ctx, transactionContextKey, tx), tx)
	})
}

func internalPaginationFromOptions(opts ...x.PaginationOptionSetter) (*internalPagination, error) {
	xp := x.GetPaginationOptions(opts...)
	ip := &internalPagination{
		Limit: xp.Size,
	}
	if ip.Limit == 0 {
		ip.Limit = defaultPageSize
	}
	return ip, ip.parsePageToken(xp.Token)
}

func (p *internalPagination) parsePageToken(t string) error {
	if t == pageTokenEnd {
		p.Limit = 0
		p.Offset = 0
		return nil
	}

	if t == "" {
		p.Offset = 0
		return nil
	}

	i, err := strconv.ParseUint(t, 10, 32)
	if err != nil {
		return errors.WithStack(persistence.ErrMalformedPageToken)
	}

	p.Offset = uint(i)
	return nil
}

func (p *internalPagination) encodeNextPageToken() string {
	return fmt.Sprintf("%d", p.Offset+p.Limit)
}
