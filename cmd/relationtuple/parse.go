// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/ketoapi"
)

func NewParseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parse",
		Short: "Parse human readable relationships",
		Long: "Parse human readable relationships as used in the documentation.\n" +
			"Supports various output formats. Especially useful for piping into other commands by using `--format json`.\n" +
			"Ignores comments (lines starting with `//`) and blank lines.\n\n" +
			"From file or folder:\n" +
			"\tketo relation-tuple parse -f tuples1.txt -f tuples2.txt\n" +
			"\tketo relation-tuple parse -f tuples-dir\n\n" +
			"Use '-' as filename to read from STD_IN:\n" +
			"\tketo relation-tuple parse -f -",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			files, err := cmd.Flags().GetStringSlice(FlagFile)
			if err != nil {
				return err
			}
			if len(files) == 0 {
				return fmt.Errorf("at least one file must be specified with -f")
			}

			var rts []*ketoapi.RelationTuple
			for _, fn := range files {
				rtss, err := readTuplesFromPath(cmd, fn, 0, parseHumanReadable)
				if err != nil {
					_, _ = fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
					return cmdx.FailSilently(cmd)
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

	registerFileFlag(cmd.Flags())
	cmdx.RegisterFormatFlags(cmd.Flags())

	return cmd
}

// parseHumanReadable reads human-readable relation tuples from the reader.
// Input is limited to MaxFileSizeBytes.
func parseHumanReadable(r io.Reader) ([]*ketoapi.RelationTuple, error) {
	fc, err := io.ReadAll(io.LimitReader(r, int64(MaxFileSizeBytes)+1))
	if err != nil {
		return nil, fmt.Errorf("reading input: %w", err)
	}
	if len(fc) > MaxFileSizeBytes {
		return nil, fmt.Errorf("input exceeds maximum size of %d bytes", MaxFileSizeBytes)
	}

	parts := strings.Split(string(fc), "\n")
	rts := make([]*ketoapi.RelationTuple, 0, len(parts))
	for i, row := range parts {
		row = strings.TrimSpace(row)
		// ignore comments and empty lines
		if row == "" || strings.HasPrefix(row, "//") {
			continue
		}

		rt, err := (&ketoapi.RelationTuple{}).FromString(row)
		if err != nil {
			return nil, fmt.Errorf("line %d: could not decode %q: %w", i+1, row, err)
		}
		rts = append(rts, rt)
	}

	return rts, nil
}
