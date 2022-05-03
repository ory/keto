package migrations

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/ory/x/popx"
	"github.com/ory/x/sqlcon"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/persistence/sql"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

// Partially transferred from tree at https://github.com/ory/keto/tree/88cedc35b5bcb79ee54e361e00b9d7f60f27b431
// Right before https://github.com/ory/keto/pull/638

type (
	dependencies interface {
		persistence.Provider
		x.LoggerProvider
		config.Provider

		PopConnection(ctx context.Context) (*pop.Connection, error)
		PopConnectionWithOpts(ctx context.Context, dets ...func(*pop.ConnectionDetails)) (*pop.Connection, error)
	}
	toSingleTableMigrator struct {
		d       dependencies
		PerPage int
	}

	relationTuple struct {
		// An ID field is required to make pop happy. The actual ID is a composite primary key.
		ID         string               `db:"shard_id" json:"-"`
		Object     string               `db:"object" json:"object"`
		Relation   string               `db:"relation" json:"relation"`
		Subject    string               `db:"subject" json:"subject"`
		CommitTime time.Time            `db:"commit_time" json:"commit_time"`
		Namespace  *namespace.Namespace `db:"-" json:"-"`
	}
	relationTuples []*relationTuple
	contextKey     string

	ErrInvalidTuples []*relationTuple
)

var (
	//go:embed namespace_migrations/*.sql
	namespaceMigrations embed.FS
)

const namespaceCtxKey contextKey = "namespace"

func tableFromNamespace(n *namespace.Namespace) string {
	return fmt.Sprintf("keto_%0.10d_relation_tuples", n.ID)
}

func namespaceIDFromTable(t string) (nID int32, err error) {
	_, err = fmt.Sscanf(t, "keto_%d_relation_tuples", &nID)
	return nID, errors.WithStack(err)
}

func migrationTableFromNamespace(n *namespace.Namespace) string {
	return fmt.Sprintf("keto_namespace_%0.10d_migrations", n.ID)
}

func namespaceTableFromContext(ctx context.Context) string {
	n, ok := ctx.Value(namespaceCtxKey).(*namespace.Namespace)
	if n == nil || !ok {
		panic("namespace context key not set")
	}
	return tableFromNamespace(n)
}

func (e ErrInvalidTuples) Error() string {
	msg := "found non-deserializable relationtuples: "
	raw, err := json.Marshal(e)
	if err != nil {
		msg += "internal error: " + err.Error()
	} else {
		msg += string(raw)
	}
	return msg
}

func (e ErrInvalidTuples) Is(other error) bool {
	_, ok := other.(ErrInvalidTuples)
	return ok
}

func (r *relationTuple) toInternal() (*relationtuple.InternalRelationTuple, error) {
	if r == nil {
		return nil, nil
	}

	sub, err := relationtuple.SubjectFromString(r.Subject)
	if err != nil {
		return nil, err
	}

	return &relationtuple.InternalRelationTuple{
		Relation:  r.Relation,
		Object:    r.Object,
		Namespace: r.Namespace.Name,
		Subject:   sub,
	}, nil
}

func (relationTuples) TableName(ctx context.Context) string {
	return namespaceTableFromContext(ctx)
}

func (relationTuple) TableName(ctx context.Context) string {
	return namespaceTableFromContext(ctx)
}

func NewToSingleTableMigrator(d dependencies) *toSingleTableMigrator {
	return &toSingleTableMigrator{
		d:       d,
		PerPage: 100,
	}
}

func (m *toSingleTableMigrator) NamespaceMigrationBox(ctx context.Context, n *namespace.Namespace) (*popx.MigrationBox, error) {
	c, err := m.d.PopConnectionWithOpts(ctx, func(d *pop.ConnectionDetails) {
		d.Options = map[string]string{
			"migration_table_name": migrationTableFromNamespace(n),
		}
	})
	if err != nil {
		return nil, err
	}

	return popx.NewMigrationBox(
		namespaceMigrations,
		popx.NewMigrator(c, m.d.Logger(), nil, 0),
		popx.WithTemplateValues(map[string]interface{}{
			"tableName": tableFromNamespace(n),
		}),
	)
}

func (m *toSingleTableMigrator) GetOldRelationTuples(ctx context.Context, n *namespace.Namespace, page, perPage int) (relationTuples, bool, error) {
	q := m.d.Persister().Connection(ctx).
		WithContext(context.WithValue(ctx, namespaceCtxKey, n)).
		Order("object, relation, subject, commit_time").
		Paginate(page, perPage)

	var res relationTuples
	if err := q.All(&res); err != nil {
		return nil, false, sqlcon.HandleError(err)
	}
	for _, r := range res {
		r.Namespace = n
	}
	return res, q.Paginator.Page < q.Paginator.TotalPages, nil
}

func (m *toSingleTableMigrator) InsertOldRelationTuples(ctx context.Context, n *namespace.Namespace, rs ...*relationtuple.InternalRelationTuple) error {
	for _, r := range rs {
		if r.Subject == nil {
			return errors.New("subject is not allowed to be nil")
		}

		m.d.Logger().WithFields(r.ToLoggerFields()).Trace("creating in legacy database")

		if err := m.d.Persister().Connection(context.WithValue(ctx, namespaceCtxKey, n)).Create(&relationTuple{
			ID:         "testing only",
			Object:     r.Object,
			Relation:   r.Relation,
			Subject:    r.Subject.String(),
			CommitTime: time.Now(),
		}); err != nil {
			return err
		}
	}
	return nil
}

func (m *toSingleTableMigrator) MigrateNamespace(ctx context.Context, n *namespace.Namespace) error {
	p, ok := m.d.Persister().(*sql.Persister)
	if !ok {
		panic("got unexpected persister")
	}

	var irrecoverableRTs ErrInvalidTuples

	if err := p.Transaction(ctx, func(ctx context.Context, _ *pop.Connection) error {
		for page := 1; ; page++ {
			rs, hasNext, err := m.GetOldRelationTuples(ctx, n, page, m.PerPage)
			if err != nil {
				return err
			}

			for _, r := range rs {
				ri, err := r.toInternal()
				if err != nil {
					m.d.Logger().WithField("relation_tuple", r).WithField("hint", "").WithError(err).Warn("Skipping relation tuple, it seems to be in a broken state. Please recreate it manually.")
					irrecoverableRTs = append(irrecoverableRTs, r)
					continue
				}
				rt := &sql.RelationTuple{
					ID:         uuid.Must(uuid.NewV4()),
					CommitTime: r.CommitTime,
				}
				if err := rt.FromInternal(ctx, p, ri); err != nil {
					return err
				}

				m.d.Logger().WithFields(ri.ToLoggerFields()).Debug("creating in new table...")
				if err := sqlcon.HandleError(
					p.CreateWithNetwork(ctx, rt),
				); err != nil {
					return err
				}
			}

			if !hasNext {
				break
			}
		}

		return nil
	}); err != nil {
		return err
	}

	if len(irrecoverableRTs) != 0 {
		return irrecoverableRTs
	}
	return nil
}

func (m *toSingleTableMigrator) LegacyNamespaces(ctx context.Context) ([]*namespace.Namespace, error) {
	c, err := m.d.PopConnection(ctx)
	if err != nil {
		return nil, err
	}

	var query *pop.Query
	switch d := c.Dialect.Name(); d {
	case "sqlite3":
		query = c.RawQuery("SELECT name FROM sqlite_master WHERE type='table' AND name LIKE 'keto_%_relation_tuples'")
	case "postgres":
		query = c.RawQuery("SELECT tablename FROM pg_catalog.pg_tables WHERE tablename LIKE 'keto_%_relation_tuples'")
	case "cockroach":
		query = c.RawQuery("SELECT table_name FROM information_schema.tables WHERE table_name LIKE 'keto_%_relation_tuples'")
	case "mysql":
		query = c.RawQuery("SELECT table_name FROM information_schema.tables WHERE table_name LIKE 'keto_%_relation_tuples' AND table_schema = DATABASE()")
	default:
		panic("got unknown database dialect " + d)
	}

	var tableNames []string
	if err := sqlcon.HandleError(query.All(&tableNames)); err != nil {
		return nil, err
	}
	m.d.Logger().Debugf("Found tables %v", tableNames)

	nm, err := m.d.Config(ctx).NamespaceManager()
	if err != nil {
		return nil, err
	}

	namespaces := make([]*namespace.Namespace, len(tableNames))
	for i := range tableNames {
		cID, err := namespaceIDFromTable(tableNames[i])
		if err != nil {
			return nil, err
		}
		namespaces[i], err = nm.GetNamespaceByConfigID(ctx, cID)
		if err != nil {
			return nil, err
		}
	}

	return namespaces, nil
}

func (m *toSingleTableMigrator) MigrateDown(ctx context.Context, n *namespace.Namespace) error {
	mb, err := m.NamespaceMigrationBox(ctx, n)
	if err != nil {
		return err
	}
	if err := mb.Down(ctx, 0); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
