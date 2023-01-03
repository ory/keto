// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"errors"
	"fmt"

	"github.com/ory/jsonschema/v3"
	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/ketoctx"
)

func NewRegistry(cmd *cobra.Command, opts []ketoctx.Option) (driver.Registry, error) {
	reg, err := driver.NewDefaultRegistry(cmd.Context(), cmd.Flags(), false, opts)
	if errors.Is(err, persistence.ErrNetworkMigrationsMissing) {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Migrations were not applied yet, please apply them first.")
		return nil, cmdx.FailSilently(cmd)
	} else if validationErr := new(jsonschema.ValidationError); errors.As(err, &validationErr) {
		// the configx provider already printed the validation error
		return nil, cmdx.FailSilently(cmd)
	} else if err != nil {
		return nil, err
	}

	return reg, nil
}
