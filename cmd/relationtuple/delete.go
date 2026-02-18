// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/cmd/helpers"
	"github.com/ory/keto/ketoapi"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"
)

const (
	MaxFileSizeBytes = 10 * 1024 * 1024 // 10 MB
	MaxDirDepth      = 3
)

func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <subject_namespace>:<subject_id> <relation> <object_namespace>:<object_id>",
		Short: "Delete relationship tuples from inline arguments or JSON files and folders",
		Long: "Delete relationship tuples from inline arguments or JSON files and folders.\n\n" +
			"Inline example:\n" +
			"	keto relation-tuple delete User:alice owner Doc:readme\n\n" +

			"From file or folder:\n" +
			"	keto relation-tuple delete -f relationships1.json -f relationships2.json\n" +
			"	keto relation-tuple delete -f relationships-dir1 -f relationships-dir2\n\n" +

			"If a directory is provided, all JSON files inside it are processed.\n" +
			"Use '-' as filename to read from STD_IN:\n" +
			"	keto relation-tuple delete -f -",
		Args: cobra.ArbitraryArgs,
		RunE: transactTuples(rts.RelationTupleDelta_ACTION_DELETE),
	}

	registerFileFlag(cmd.Flags())
	registerPackageFlags(cmd.Flags())

	return cmd
}

func transactTuples(action rts.RelationTupleDelta_Action) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		files, err := cmd.Flags().GetStringSlice(FlagFile)
		if err != nil {
			return err
		}
		if len(files) > 0 {
			return transactTuplesFromFiles(cmd, action, files)
		}

		if len(args) != 3 {
			return fmt.Errorf("expected inline arguments or JSON files and folders")
		}

		return transactTupleFromInlineArgs(cmd, action, args)
	}
}

func transactTupleFromInlineArgs(cmd *cobra.Command, action rts.RelationTupleDelta_Action, args []string) error {
	namespace, object, err := helpers.ParseNamespaceObject(cmd, args[2:])
	if err != nil {
		return err
	}

	sub, err := helpers.ParseSubject(args[0])
	if err != nil {
		return fmt.Errorf("could not parse subject %q: %w", args[0], err)
	}
	tuple := &rts.RelationTuple{
		Namespace: namespace,
		Object:    object,
		Relation:  args[1],
		Subject:   sub,
	}

	conn, err := client.GetWriteConn(cmd)
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	cl := rts.NewWriteServiceClient(conn)
	_, err = cl.TransactRelationTuples(cmd.Context(), &rts.TransactRelationTuplesRequest{
		RelationTupleDeltas: []*rts.RelationTupleDelta{{
			Action:        action,
			RelationTuple: tuple,
		}},
	})
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not make request: %s\n", err)
		return cmdx.FailSilently(cmd)
	}

	c, err := NewProtoCollection([]*rts.RelationTuple{tuple})
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "error printing the response: %s\n", err)
	}
	cmdx.PrintTable(cmd, c)
	return nil
}

func transactTuplesFromFiles(cmd *cobra.Command, action rts.RelationTupleDelta_Action, files []string) error {
	conn, err := client.GetWriteConn(cmd)
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	var tuples []*ketoapi.RelationTuple
	var deltas []*rts.RelationTupleDelta
	for _, fn := range files {
		parsed, err := readTuplesFromPath(cmd, fn, 0, decodeTuples)
		if err != nil {
			// silencing the usage message as from here on, there can't be any syntax error in the command usage,
			// but rather an error with the content of the files
			_, _ = fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
			return cmdx.FailSilently(cmd)
		}
		for _, t := range parsed {
			tuples = append(tuples, t)
			deltas = append(deltas, &rts.RelationTupleDelta{
				Action:        action,
				RelationTuple: t.ToProto(),
			})
		}
	}

	cl := rts.NewWriteServiceClient(conn)
	_, err = cl.TransactRelationTuples(cmd.Context(), &rts.TransactRelationTuplesRequest{
		RelationTupleDeltas: deltas,
	})
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not make request: %s\n", err)
		return cmdx.FailSilently(cmd)
	}

	cmdx.PrintTable(cmd, NewAPICollection(tuples))
	return nil
}

// TupleDecoder reads relation tuples from a reader. Implementations handle
// format-specific decoding (e.g. JSON, human-readable).
type TupleDecoder func(r io.Reader) ([]*ketoapi.RelationTuple, error)

// readTuplesFromPath reads relation tuples from a file path, directory, or
// stdin (when path is "-"). The depth parameter tracks directory nesting to
// enforce MaxDirDepth. The decode function handles format-specific parsing.
func readTuplesFromPath(cmd *cobra.Command, path string, depth int, decode TupleDecoder) ([]*ketoapi.RelationTuple, error) {
	if path == "-" {
		tuples, err := decode(cmd.InOrStdin())
		if err != nil {
			return nil, fmt.Errorf("stdin: %w", err)
		}
		return tuples, nil
	}

	cleanPath := filepath.Clean(path)
	info, err := os.Stat(cleanPath)
	if err != nil {
		return nil, fmt.Errorf("error getting stats for %s: %w", path, err)
	}

	if info.IsDir() {
		if depth >= MaxDirDepth {
			return nil, fmt.Errorf("maximum directory depth of %d exceeded at %s", MaxDirDepth, path)
		}
		tuples, err := readTuplesFromDir(cleanPath, decode)
		if err != nil {
			return nil, err
		}
		return tuples, nil
	}

	return readTuplesFromFile(cleanPath, decode)
}

func readTuplesFromDir(dir string, decode TupleDecoder) ([]*ketoapi.RelationTuple, error) {
	var tuples []*ketoapi.RelationTuple

	fsys := os.DirFS(dir)
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return fmt.Errorf("walking %s: %w", path, walkErr)
		}

		depth := 0
		if path != "." {
			depth = strings.Count(path, "/") + 1
		}

		if depth > MaxDirDepth {
			return fmt.Errorf("maximum directory depth of %d exceeded at %s", MaxDirDepth, path)
		}

		if d.IsDir() {
			return nil
		}

		fullPath := filepath.Join(dir, filepath.FromSlash(path))

		t, err := readTuplesFromFile(fullPath, decode)
		if err != nil {
			return err
		}
		tuples = append(tuples, t...)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tuples, nil
}

// readTuplesFromFile opens a single file and decodes its content using the
// provided decoder.
func readTuplesFromFile(path string, decode TupleDecoder) ([]*ketoapi.RelationTuple, error) {
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("opening file %s: %w", path, err)
	}
	defer func() { _ = f.Close() }()

	tuples, err := decode(f)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return tuples, nil
}

// decodeTuples decodes JSON from the reader as either a single RelationTuple
// or an array of RelationTuples. Input is limited to MaxFileSizeBytes.
func decodeTuples(r io.Reader) ([]*ketoapi.RelationTuple, error) {
	data, err := io.ReadAll(io.LimitReader(r, MaxFileSizeBytes+1))
	if err != nil {
		return nil, fmt.Errorf("reading input: %w", err)
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("empty input")
	}
	if len(data) > MaxFileSizeBytes {
		return nil, fmt.Errorf("input exceeds maximum size of %d bytes", MaxFileSizeBytes)
	}

	// Trim UTF-8 BOM first, then leading/trailing whitespace
	trimmed := bytes.TrimPrefix(data, []byte{0xEF, 0xBB, 0xBF})
	trimmed = bytes.TrimSpace(trimmed)

	if len(trimmed) == 0 {
		return nil, fmt.Errorf("empty input after trimming whitespace")
	}

	decoder := json.NewDecoder(bytes.NewReader(trimmed))
	decoder.DisallowUnknownFields()

	if trimmed[0] == '[' {
		var tuples []*ketoapi.RelationTuple
		if err := decoder.Decode(&tuples); err != nil {
			return nil, fmt.Errorf("could not decode: %w", err)
		}
		return tuples, nil
	}

	var tuple ketoapi.RelationTuple
	if err := decoder.Decode(&tuple); err != nil {
		return nil, fmt.Errorf("could not decode: %w", err)
	}
	return []*ketoapi.RelationTuple{&tuple}, nil
}
