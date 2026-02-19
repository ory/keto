// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ory/jsonschema/v3"
	"github.com/ory/x/cmdx"
	"github.com/ory/x/randx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/ketoapi"
	"github.com/ory/keto/ketoctx"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
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

func ParseSubject(s string) (*rts.Subject, error) {
	if strings.Contains(s, ":") {
		su, err := (&ketoapi.SubjectSet{}).FromString(s)
		if err != nil {
			return nil, err
		}

		return rts.NewSubjectSet(su.Namespace, su.Object, su.Relation), nil
	}
	return rts.NewSubjectID(s), nil
}

// ParseNamespaceObject parses namespace and object from args that may be in the
// combined format ["namespace:object"] or the legacy format ["namespace", "object"].
// It writes a deprecation warning to cmd.ErrOrStderr if the legacy format is used.
func ParseNamespaceObject(cmd *cobra.Command, args []string) (namespace, object string, err error) {
	switch len(args) {
	case 2:
		_, _ = fmt.Fprintln(cmd.ErrOrStderr(), "Warning: passing namespace and object as separate arguments is deprecated. Use <object_namespace>:<object_id> instead.")
		return args[0], args[1], nil
	case 1:
		namespace, object, ok := strings.Cut(args[0], ":")
		// empty ObjectID is allowed
		if !ok || namespace == "" {
			return "", "", fmt.Errorf("expected <object_namespace>:<object_id> format, got %q", args[0])
		}
		return namespace, object, nil
	default:
		return "", "", fmt.Errorf("unexpected number of arguments for <object_namespace>:<object_id>: got %d arguments - %s", len(args), strings.Join(args, ","))
	}
}

func RandomTupleWithSubjectSet(ns1, ns2 string) *ketoapi.RelationTuple {
	return &ketoapi.RelationTuple{
		Namespace: ns1,
		Object:    randx.MustString(8, randx.AlphaNum),
		Relation:  randx.MustString(8, randx.AlphaNum),
		SubjectSet: &ketoapi.SubjectSet{
			Namespace: ns2,
			Object:    randx.MustString(8, randx.AlphaNum),
		},
	}
}

func RandomTupleWithSubjectID(ns1 string) *ketoapi.RelationTuple {
	return &ketoapi.RelationTuple{
		Namespace: ns1,
		Object:    randx.MustString(8, randx.AlphaNum),
		Relation:  randx.MustString(8, randx.AlphaNum),
		SubjectID: new(randx.MustString(8, randx.AlphaNum)),
	}
}
