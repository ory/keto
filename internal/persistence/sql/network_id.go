package sql

import (
	"context"
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
	"github.com/ory/x/sqlcon"
	"github.com/pkg/errors"
)

type networkID struct {
	ID        uuid.UUID `db:"network_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (*networkID) TableName() string {
	return "keto_networks"
}

func (p *Persister) NetworkID(ctx context.Context) (uuid.UUID, error) {
	if p.networkIDCached == uuid.Nil {
		var nID networkID

		if err := p.connection(ctx).First(&nID); errors.Is(err, sql.ErrNoRows) {
			var err error
			nID.ID, err = uuid.NewV4()
			if err != nil {
				return uuid.Nil, errors.WithStack(err)
			}

			if err := p.connection(ctx).Create(&nID); err != nil {
				return uuid.Nil, sqlcon.HandleError(err)
			}
		} else if err != nil {
			return uuid.Nil, sqlcon.HandleError(err)
		}

		p.networkIDCached = nID.ID
	}

	return p.networkIDCached, nil
}
