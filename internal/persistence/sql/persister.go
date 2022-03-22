package sql

import (
	"context"
	"embed"
	"fmt"
	"reflect"
	"strconv"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/ory/x/fsx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/networkx"
	"github.com/ory/x/otelx"
	"github.com/ory/x/popx"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoctx"
)

type (
	Persister struct {
		conn *pop.Connection
		d    dependencies
		nid  uuid.UUID
	}
	internalPagination struct {
		Page, PerPage int
	}
	dependencies interface {
		config.Provider
		x.LoggerProvider
		ketoctx.ContextualizerProvider

		PopConnection(ctx context.Context) (*pop.Connection, error)
	}
)

const (
	defaultPageSize int = 100
)

var (
	//go:embed migrations/sql/*.sql
	migrations embed.FS

	_ persistence.Persister = &Persister{}
)

func NewPersister(ctx context.Context, reg dependencies, nid uuid.UUID) (*Persister, error) {
	conn, err := reg.PopConnection(ctx)
	if err != nil {
		return nil, err
	}

	p := &Persister{
		d:    reg,
		nid:  nid,
		conn: conn,
	}

	return p, nil
}

func NewMigrationBox(c *pop.Connection, logger *logrusx.Logger, tracer *otelx.Tracer) (*popx.MigrationBox, error) {
	return popx.NewMigrationBox(fsx.Merge(migrations, networkx.Migrations), popx.NewMigrator(c, logger, tracer, 0))
}

func (p *Persister) Connection(ctx context.Context) *pop.Connection {
	return popx.GetConnection(ctx, p.conn.WithContext(ctx))
}

func (p *Persister) CreateWithNetwork(ctx context.Context, v interface{}) error {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Ptr && rv.Elem().Kind() != reflect.Struct {
		panic("expected to get *struct in create")
	}
	nID := rv.Elem().FieldByName("NetworkID")
	if !nID.IsValid() || !nID.CanSet() {
		panic("expected struct to have a 'NetworkID uuid.UUID' field")
	}
	nID.Set(reflect.ValueOf(p.NetworkID(ctx)))

	return p.Connection(ctx).Create(v)
}

func (p *Persister) QueryWithNetwork(ctx context.Context) *pop.Query {
	return p.Connection(ctx).Where("nid = ?", p.NetworkID(ctx))
}

func (p *Persister) Transaction(ctx context.Context, f func(ctx context.Context, c *pop.Connection) error) error {
	return popx.Transaction(ctx, p.conn.WithContext(ctx), f)
}

func (p *Persister) NetworkID(ctx context.Context) uuid.UUID {
	return p.d.Contextualizer().Network(ctx, p.nid)
}

func internalPaginationFromOptions(opts ...x.PaginationOptionSetter) (*internalPagination, error) {
	xp := x.GetPaginationOptions(opts...)
	ip := &internalPagination{
		PerPage: xp.Size,
	}
	if ip.PerPage == 0 {
		ip.PerPage = defaultPageSize
	}
	return ip, ip.parsePageToken(xp.Token)
}

func (p *internalPagination) parsePageToken(t string) error {
	if t == "" {
		p.Page = 1
		return nil
	}

	i, err := strconv.ParseUint(t, 10, 32)
	if err != nil {
		return errors.WithStack(persistence.ErrMalformedPageToken)
	}

	p.Page = int(i)
	return nil
}

func (p *internalPagination) encodeNextPageToken() string {
	return fmt.Sprintf("%d", p.Page+1)
}
