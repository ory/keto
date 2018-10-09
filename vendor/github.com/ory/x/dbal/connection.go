package dbal

import (
	"net/url"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ory/x/sqlcon"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func Connect(db string, logger logrus.FieldLogger, memf func() error, sqlf func(db *sqlx.DB) error) error {
	if db == "memory" {
		return memf()
	} else if db == "" {
		return errors.New("No database URL provided")
	}

	u, err := url.Parse(db)
	if err != nil {
		return errors.WithStack(err)
	}

	switch u.Scheme {
	case "postgres":
		fallthrough
	case "mysql":
		c, err := sqlcon.NewSQLConnection(db, logger)
		if err != nil {
			return errors.WithStack(err)
		}

		cdb, err := c.GetDatabaseRetry(time.Second*15, time.Minute*5)
		if err != nil {
			return errors.WithStack(err)
		}

		return sqlf(cdb)
	}

	return errors.Errorf("The provided database URL %s can not be handled", db)
}
