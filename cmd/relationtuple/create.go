// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"github.com/spf13/cobra"

	rts "github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2"
)

const FlagFile = "file"

func NewCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <subject_namespace>:<subject_id> <relation> <object_namespace>:<object_id>",
		Short: "Create relationship tuples from inline arguments or JSON files and folders",
		Long: `Create relationship tuples from inline arguments or JSON files and folders.
Inline example:
	keto relation-tuple create User:alice owner Doc:readme

From file or folder:
	keto relation-tuple create -f relationships1.json -f relationships2.json
	keto relation-tuple create -f relationships-dir1 -f relationships-dir2

If a directory is provided, all JSON files inside it are processed.
Use '-' as filename to read from STD_IN:
	keto relation-tuple create -f -`,
		Args: cobra.ArbitraryArgs,
		RunE: transactTuples(rts.RelationTupleDelta_ACTION_INSERT),
	}

	registerFileFlag(cmd.Flags())
	registerPackageFlags(cmd.Flags())

	return cmd
}
