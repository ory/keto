package relationtuple

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/relationtuple"
)

func newParseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parse",
		Short: "Parse human readable relation tuples.",
		Long:  "Parse human readable relation tuples as used in the documentation. Supports various output formats. Especially useful for piping into other commands.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var rts []*relationtuple.InternalRelationTuple
			for _, fn := range args {
				rtss, err := parseFile(cmd, fn)
				if err != nil {
					return err
				}
				rts = append(rts, rtss...)
			}

			cmdx.PrintTable(cmd, relationtuple.NewRelationCollection(rts))
			return nil
		},
	}

	cmdx.RegisterFormatFlags(cmd.Flags())

	return cmd
}

func parseFile(cmd *cobra.Command, fn string) ([]*relationtuple.InternalRelationTuple, error) {
	var f io.Reader
	if fn == "-" {
		// set human readable filename here for debug output
		fn = "stdin"
		f = cmd.InOrStdin()
	} else {
		ff, err := os.Open(fn)
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not open file %s: %v\n", fn, err)
			return nil, cmdx.FailSilently(cmd)
		}
		defer ff.Close()
		f = ff
	}

	fc, err := io.ReadAll(f)
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could read file %s: %v\n", fn, err)
		return nil, cmdx.FailSilently(cmd)
	}

	parts := strings.Split(string(fc), "\n")
	rts := make([]*relationtuple.InternalRelationTuple, 0, len(parts))
	for i, row := range parts {
		row = strings.TrimSpace(row)
		// ignore comments and empty lines
		if row == "" || strings.HasPrefix(row, "//") {
			continue
		}

		rt, err := (&relationtuple.InternalRelationTuple{}).FromString(row)
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not decode %s:%d\n  %s\n\n%v\n", fn, i+1, row, err)
			return nil, cmdx.FailSilently(cmd)
		}
		rts = append(rts, rt)
	}

	return rts, nil
}
