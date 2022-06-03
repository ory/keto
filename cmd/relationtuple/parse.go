package relationtuple

import (
	"fmt"
	"github.com/ory/keto/ketoapi"
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
		Short: "Parse human readable relation tuples",
		Long: "Parse human readable relation tuples as used in the documentation.\n" +
			"Supports various output formats. Especially useful for piping into other commands by using `--format json`.\n" +
			"Ignores comments (starting with `//`) and blank lines.",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var rts []*ketoapi.RelationTuple
			for _, fn := range args {
				rtss, err := parseFile(cmd, fn)
				if err != nil {
					return err
				}
				rts = append(rts, rtss...)
			}

			if len(rts) == 1 {
				cmdx.PrintRow(cmd, rts[0])
				return nil
			}
			cmdx.PrintTable(cmd, NewAPICollection(rts))
			return nil
		},
	}

	cmdx.RegisterFormatFlags(cmd.Flags())

	return cmd
}

func parseFile(cmd *cobra.Command, fn string) ([]*ketoapi.RelationTuple, error) {
	var f io.Reader
	if fn == "-" {
		// set human readable filename here for debug and error messages
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

		rt, err := (&ketoapi.RelationTuple{}).FromString(row)
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not decode %s:%d\n  %s\n\n%v\n", fn, i+1, row, err)
			return nil, cmdx.FailSilently(cmd)
		}
		rts = append(rts, rt)
	}

	return rts, nil
}
