// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package namespace

import (
	"bytes"
	"fmt"
	"os"

	"github.com/ory/keto/embedx"

	"github.com/ory/jsonschema/v3"
	"github.com/ory/x/cmdx"
	"github.com/ory/x/configx"
	"github.com/ory/x/jsonschemax"
	"github.com/ory/x/logrusx"
	"github.com/segmentio/objconv/yaml"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
)

func NewValidateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Deprecated: "The legacy namespaces are deprecated. Please use the new Ory Permission Language instead.",
		Aliases:    []string{"validate"},
		Use:        "validate-legacy <namespace.yml> [<namespace2.yml> ...] | validate -c <config.yaml>",
		Short:      "Validate legacy namespace definitions",
		Long: `validate-legacy
Validates namespace definitions. Parses namespace yaml files or configuration
files passed via the configuration flag. Returns human readable errors. Useful for
debugging.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfFlag := cmd.Flag(configx.FlagConfig)
			if cfFlag.Changed {
				cfiles, err := cmd.Flags().GetStringSlice(configx.FlagConfig)
				if err != nil {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Failed to read config command line flag\n%+v\n", err)
					return cmdx.FailSilently(cmd)
				}
				for _, fn := range cfiles {
					err := validateConfigFile(cmd, fn)
					if err != nil {
						return err
					}
				}
			}

			// user passed a list of namespace files
			for _, fn := range args {
				_, err := validateNamespaceFile(cmd, fn)
				if err != nil {
					return err
				}
			}

			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Congrats, all files are valid!")
			return nil
		},
	}

	return cmd
}

var configSchema *jsonschema.Schema

const schemaPath = "github.com/ory/keto/.schema/config.schema.json"

func getSchema(cmd *cobra.Command) (*jsonschema.Schema, error) {
	if configSchema == nil {
		c := jsonschema.NewCompiler()
		if err := c.AddResource(schemaPath, bytes.NewBuffer(embedx.ConfigSchema)); err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not add the config schema file to the compiler. This is an internal error that should be reported. Thanks ;)\n%+v\n", err)
			return nil, cmdx.FailSilently(cmd)
		}

		var err error
		configSchema, err = c.Compile(cmd.Context(), schemaPath+"#/definitions/namespace")
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not compile the config schema file. This is an internal error that should be reported. Thanks ;)\n%+v\n", err)
			return nil, cmdx.FailSilently(cmd)
		}
	}
	return configSchema, nil
}

func validateNamespaceFile(cmd *cobra.Command, fn string) (*namespace.Namespace, error) {
	fc, err := os.ReadFile(fn)
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not read file \"%s\": %+v\n", fn, err)
		return nil, cmdx.FailSilently(cmd)
	}

	parse, err := config.GetParser(fn)
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Unable to infer file type from \"%s\": %+v\n", fn, err)
		return nil, cmdx.FailSilently(cmd)
	}

	return validateNamespaceBytes(cmd, fn, fc, parse)
}

func validateNamespaceBytes(cmd *cobra.Command, name string, b []byte, parser config.Parser) (*namespace.Namespace, error) {
	schema, err := getSchema(cmd)
	if err != nil {
		return nil, err
	}

	var val map[string]interface{}
	if err := parser(b, &val); err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Encountered unmarshal error for \"%s\": %+v\n", name, err)
		return nil, cmdx.FailSilently(cmd)
	}

	if err := schema.ValidateInterface(val); err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "File %s was not a valid namespace file. Reasons:\n", name)
		jsonschemax.FormatValidationErrorForCLI(cmd.ErrOrStderr(), embedx.ConfigSchema, err)
		return nil, cmdx.FailSilently(cmd)
	}

	var n namespace.Namespace
	if err := parser(b, &n); err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Encountered unmarshal error for \"%s\": %+v\n", name, err)
		return nil, cmdx.FailSilently(cmd)
	}

	return &n, nil
}

func validateConfigFile(cmd *cobra.Command, fn string) error {
	fc, err := os.ReadFile(fn)
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not read file \"%s\": %+v\n", fn, err)
		return cmdx.FailSilently(cmd)
	}

	parse, err := config.GetParser(fn)
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Unable to infer file type from \"%s\": %+v\n", fn, err)
		return cmdx.FailSilently(cmd)
	}

	var val map[string]interface{}
	if err := parse(fc, &val); err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Encountered parse error for \"%s\": %+v\n", fn, err)
		return cmdx.FailSilently(cmd)
	}

	if ns, ok := val["namespaces"]; ok {
		// ns can either be a ws url or a list of namespace objects)
		switch t := ns.(type) {
		case string:
			logger := logrusx.New("cmd", "0")
			cw, err := config.NewNamespaceWatcher(cmd.Context(), logger, t)
			if err != nil {
				_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Encountered error reading config: %+v\n", err)
				return cmdx.FailSilently(cmd)
			}

			files := cw.NamespaceFiles()
			for _, file := range files {
				_, err := validateNamespaceBytes(cmd, file.Name, file.Contents, file.Parser)
				if err != nil {
					return err
				}
			}
			return nil
		case []interface{}:
			for i, obj := range t {
				fc, err := yaml.Marshal(obj)
				if err != nil {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Internal error. Failed to marshal yaml %+v\n", err)
					return cmdx.FailSilently(cmd)
				}
				_, err = validateNamespaceBytes(cmd, fmt.Sprintf("index: %d", i), fc, yaml.Unmarshal)
				if err != nil {
					return err
				}
			}
			return nil
		case []map[string]interface{}:
			for i, obj := range t {
				fc, err := yaml.Marshal(obj)
				if err != nil {
					_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Internal error. Failed to marshal yaml %+v\n", err)
					return cmdx.FailSilently(cmd)
				}
				_, err = validateNamespaceBytes(cmd, fmt.Sprintf("index: %d", i), fc, yaml.Unmarshal)
				if err != nil {
					return err
				}
			}
			return nil
		default:
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "unknown type %T for key 'namespaces' in config file: %v\n", t, t)
			return cmdx.FailSilently(cmd)
		}
	} else {
		return nil
	}
}
