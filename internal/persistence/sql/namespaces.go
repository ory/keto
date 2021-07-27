package sql

import (
	"context"
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/namespace"
)

type namespaceRecord struct {
	NetworkID uuid.UUID `db:"nid"`
	ID        uuid.UUID `db:"id"`
	SerialID  int       `db:"serial_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (*namespaceRecord) TableName(context.Context) string {
	return "keto_namespace_ids"
}

func (p *Persister) cacheNamespaceMapping(serialID int, id uuid.UUID) {
	p.namespaceIDCache.Set(serialID, id, 1)
	p.namespaceIDCache.Set(id[:], serialID, 1)
}

// GetNamespaceID returns the internal UUID for use in the relation tuple table.
// This additional indirection is to allow distributed configs later on, so that we don't have
// to sync on serial configIDs by directly using UUIDs.
func (p *Persister) GetNamespaceID(ctx context.Context, name string) (uuid.UUID, error) {
	nm, err := p.d.Config().NamespaceManager()
	if err != nil {
		return uuid.Nil, err
	}
	n, err := nm.GetNamespaceByName(ctx, name)
	if err != nil {
		return uuid.Nil, err
	}

	if id, ok := p.namespaceIDCache.Get(n.ID); ok {
		return id.(uuid.UUID), nil
	}

	dbEntry := namespaceRecord{}
	if err := p.QueryWithNetwork(ctx).Where("serial_id = ?", n.ID).First(&dbEntry); errors.Is(err, sql.ErrNoRows) {
		dbEntry.ID = uuid.Must(uuid.NewV4())
		dbEntry.SerialID = n.ID

		if err := p.CreateWithNetwork(ctx, &dbEntry); err != nil {
			return uuid.Nil, errors.WithStack(err)
		}
		// fall through to adding to the cache
	} else if err != nil {
		return uuid.Nil, errors.WithStack(err)
	}

	p.cacheNamespaceMapping(n.ID, dbEntry.ID)

	return dbEntry.ID, nil
}

// GetNamespaceConfigID returns the serial ID associated with the namespace
// as stored in the relation tuple table.
// Reverse lookup of GetNamespaceID.
func (p *Persister) GetNamespaceConfigID(ctx context.Context, id uuid.UUID) (*namespace.Namespace, error) {
	configID, ok := p.namespaceIDCache.Get(id[:])
	if !ok {
		dbEntry := namespaceRecord{}
		if err := p.QueryWithNetwork(ctx).Where("id = ?", id).First(&dbEntry); errors.Is(err, sql.ErrNoRows) {
			return nil, errors.WithStack(err)
		}

		p.cacheNamespaceMapping(dbEntry.SerialID, id)

		configID = dbEntry.SerialID
	}

	nm, err := p.d.Config().NamespaceManager()
	if err != nil {
		return nil, err
	}
	return nm.GetNamespaceByConfigID(ctx, configID.(int))
}
