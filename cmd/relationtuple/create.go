// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ory/keto/ketoapi"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/spf13/cobra"

	"github.com/ory/x/cmdx"
)

func NewCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <relationships.json> [<relationships-dir>]",
		Short: "Create relationships from JSON files",
		Long: "Create relationships from JSON files.\n" +
			"A directory will be traversed and all relationships will be created.\n" +
			"Pass the special filename `-` to read from STD_IN.",
		Args: cobra.MinimumNArgs(1),
		RunE: transactRelationTuples(rts.RelationTupleDelta_ACTION_INSERT),
	}
	registerPackageFlags(cmd.Flags())

	return cmd
}

func readTuplesFromArg(cmd *cobra.Command, arg string) ([]*ketoapi.RelationTuple, error) {
	var f io.Reader
	if arg == "-" {
		f = cmd.InOrStdin()
	} else {
		stats, err := os.Stat(arg)
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Error getting stats for %s: %s\n", arg, err)
			return nil, cmdx.FailSilently(cmd)
		}

		if stats.IsDir() {
			fi, err := os.ReadDir(arg)
			if err != nil {
				return nil, err
			}

			var tuples []*ketoapi.RelationTuple
			for _, child := range fi {
				t, err := readTuplesFromArg(cmd, filepath.Join(arg, child.Name()))
				if err != nil {
					return nil, err
				}
				tuples = append(tuples, t...)
			}
			return tuples, nil
		}

		f, err = os.Open(arg)
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Error processing arg %s: %s\n", arg, err)
			return nil, cmdx.FailSilently(cmd)
		}
	}

	fc, err := io.ReadAll(f)
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could read file %s: %s\n", arg, err)
		return nil, cmdx.FailSilently(cmd)
	}

	decoder := json.NewDecoder(bytes.NewReader(fc))
	decoder.DisallowUnknownFields()
	// it is ok to not validate beforehand because json.Unmarshal will report errors
	if fc[0] == '[' {
		var ts []*ketoapi.RelationTuple
		if err := decoder.Decode(&ts); err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not decode: %s\n", err)
			return nil, cmdx.FailSilently(cmd)
		}
		return ts, nil
	}

	var r ketoapi.RelationTuple
	if err := decoder.Decode(&r); err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not decode: %s\n", err)
		return nil, cmdx.FailSilently(cmd)
	}

	return []*ketoapi.RelationTuple{&r}, nil
}
