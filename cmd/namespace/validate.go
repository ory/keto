package namespace

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/ory/keto/internal/driver/config"

	"github.com/ory/x/jsonschemax"

	"github.com/ory/jsonschema/v3"
	"github.com/ory/x/cmdx"
	"github.com/segmentio/objconv/yaml"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/namespace"
)

func NewValidateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate <namespace.yml> [<namespace2.yml> ...]",
		Args:  cobra.MinimumNArgs(1),
		Short: "Validate namespace files",
		Long:  "Validate one or more namespace yaml files and get human readable errors. Useful for debugging.",
		RunE: func(cmd *cobra.Command, args []string) error {
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

func validateNamespaceFile(cmd *cobra.Command, fn string) (*namespace.Namespace, error) {
	if configSchema == nil {
		c := jsonschema.NewCompiler()
		if err := c.AddResource(schemaPath, bytes.NewBuffer(config.Schema)); err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not add the config schema file to the compiler. This is an internal error that should be reported. Thanks ;)\n%+v\n", err)
			return nil, cmdx.FailSilently(cmd)
		}

		var err error
		configSchema, err = c.Compile(schemaPath + "#/definitions/namespace")
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not compile the config schema file. This is an internal error that should be reported. Thanks ;)\n%+v\n", err)
			return nil, cmdx.FailSilently(cmd)
		}
	}

	fc, err := ioutil.ReadFile(fn)
	if err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not read file \"%s\": %+v\n", fn, err)
		return nil, cmdx.FailSilently(cmd)
	}

	var val map[string]interface{}
	if err := yaml.Unmarshal(fc, &val); err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Encountered yaml unmarshal error for \"%s\": %+v\n", fn, err)
		return nil, cmdx.FailSilently(cmd)
	}

	if err := configSchema.ValidateInterface(val); err != nil {
		jsonschemax.FormatValidationErrorForCLI(cmd.ErrOrStderr(), config.Schema, err)
		return nil, cmdx.FailSilently(cmd)
	}

	var n namespace.Namespace
	if err := yaml.Unmarshal(fc, &n); err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Encountered yaml unmarshal error for \"%s\": %+v\n", fn, err)
		return nil, cmdx.FailSilently(cmd)
	}

	return &n, nil
}
