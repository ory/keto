package relationtuple

import (
	"encoding/json"
	"fmt"
	acl "github.com/ory/keto/api/keto/acl/v1alpha1"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/spf13/cobra"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/cmd/client"
)

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create <relation-tuple.json> [<relation-tuple-dir>]",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := client.GetGRPCConn(cmd)
			if err != nil {
				return err
			}

			var tuples []*relationtuple.InternalRelationTuple
			var deltas []*acl.RelationTupleWriteDelta
			for _, fn := range args {
				tuple, err := readTuplesFromArg(cmd, fn)
				if err != nil {
					return err
				}
				for _, t := range tuple {
					tuples = append(tuples, t)
					deltas = append(deltas, &acl.RelationTupleWriteDelta{
						Action:        acl.RelationTupleWriteDelta_INSERT,
						RelationTuple: t.ToGRPC(),
					})
				}
			}

			cl := acl.NewWriteServiceClient(conn)

			_, err = cl.WriteRelationTuples(cmd.Context(), &acl.WriteRelationTuplesRequest{
				RelationTupleDeltas: deltas,
			})
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Error doing the request: %s\n", err)
				return cmdx.FailSilently(cmd)
			}

			cmdx.PrintTable(cmd, relationtuple.NewRelationCollection(tuples))
			return nil
		},
	}
	cmd.Flags().AddFlagSet(packageFlags)

	return cmd
}

func readTuplesFromArg(cmd *cobra.Command, arg string) ([]*relationtuple.InternalRelationTuple, error) {
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
			fi, err := ioutil.ReadDir(arg)
			if err != nil {
				return nil, err
			}

			var tuples []*relationtuple.InternalRelationTuple
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

	var r relationtuple.InternalRelationTuple
	err := json.NewDecoder(f).Decode(&r)
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not decode: %s\n", err)
		return nil, cmdx.FailSilently(cmd)
	}

	return []*relationtuple.InternalRelationTuple{&r}, nil
}
