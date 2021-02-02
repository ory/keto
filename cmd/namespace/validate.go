package namespace

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/ory/x/jsonschemax"

	"github.com/markbates/pkger"
	"github.com/ory/jsonschema/v3"
	"github.com/ory/x/cmdx"
	"github.com/segmentio/objconv/yaml"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/namespace"
)

const configSchemaPath = "github.comory/keto:/.schema/config.schema.json"

func NewValidateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate <namespace.yml> [<namespace2.yml> ...]",
		Args:  cobra.MinimumNArgs(1),
		Short: "Validate a namespace file.",
		Long:  "Validate a namespace file and get human readable errors.",
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

var (
	configSchema    *jsonschema.Schema
	configSchemaRaw []byte
)

func validateNamespaceFile(cmd *cobra.Command, fn string) (*namespace.Namespace, error) {
	if configSchema == nil || len(configSchemaRaw) == 0 {
		sf, err := pkger.Open(configSchemaPath)
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not open the config schema file. This is an internal error that should be reported. Thanks ;)\n%+v\n", err)
			return nil, cmdx.FailSilently(cmd)
		}

		configSchemaRaw, err = ioutil.ReadAll(sf)
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not read the config schema file. This is an internal error that should be reported. Thanks ;)\n%+v\n", err)
			return nil, cmdx.FailSilently(cmd)
		}

		c := jsonschema.NewCompiler()
		if err := c.AddResource(configSchemaPath, bytes.NewBuffer(configSchemaRaw)); err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not add the config schema file to the compiler. This is an internal error that should be reported. Thanks ;)\n%+v\n", err)
			return nil, cmdx.FailSilently(cmd)
		}

		configSchema, err = c.Compile(configSchemaPath + "#/definitions/namespace")
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
		jsonschemax.FormatValidationErrorForCLI(cmd.ErrOrStderr(), configSchemaRaw, err)
		return nil, cmdx.FailSilently(cmd)
	}

	var n namespace.Namespace
	if err := yaml.Unmarshal(fc, &n); err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Encountered yaml unmarshal error for \"%s\": %+v\n", fn, err)
		return nil, cmdx.FailSilently(cmd)
	}

	return &n, nil
}
