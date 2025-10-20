package ketodriver

import (
	"context"
	_ "embed"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/persistence/sql"
	"github.com/ory/keto/ketoctx"
	"github.com/spf13/pflag"
)

var SQLMigrations = sql.Migrations

// NewRegistry returns a new default keto registry. It is a simple wrapper around internal driver.NewDefaultRegistry.
func NewRegistry(ctx context.Context, flags *pflag.FlagSet, withoutNetwork bool, opts []ketoctx.Option) (driver.Registry, error) {
	return driver.NewDefaultRegistry(ctx, flags, withoutNetwork, opts)
}
