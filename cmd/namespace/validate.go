package namespace

import (
	"fmt"
	"io/ioutil"

	"github.com/markbates/pkger"
	"github.com/ory/jsonschema/v3"
	"github.com/ory/x/cmdx"
	"github.com/ory/x/viperx"
	"github.com/segmentio/objconv/yaml"
	"github.com/spf13/cobra"

	"github.com/ory/keto/internal/namespace"
)

const namespaceSchemaPath = "/.schema/namespace.schema.json"

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

var namespaceSchema *jsonschema.Schema

func validateNamespaceFile(cmd *cobra.Command, fn string) (*namespace.Namespace, error) {
	if namespaceSchema == nil {
		sf, err := pkger.Open(namespaceSchemaPath)
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not open the namespace schema file. This is an internal error that should be reported. Thanks ;)\n%+v\n", err)
			return nil, cmdx.FailSilently(cmd)
		}

		c := jsonschema.NewCompiler()
		if err := c.AddResource(namespaceSchemaPath, sf); err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not add the namespace schema file to the compiler. This is an internal error that should be reported. Thanks ;)\n%+v\n", err)
			return nil, cmdx.FailSilently(cmd)
		}

		namespaceSchema, err = c.Compile(namespaceSchemaPath)
		if err != nil {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Could not compile the namespace schema file. This is an internal error that should be reported. Thanks ;)\n%+v\n", err)
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

	if err := namespaceSchema.ValidateInterface(val); err != nil {
		viperx.PrintHumanReadableValidationErrors(cmd.ErrOrStderr(), err)
		return nil, cmdx.FailSilently(cmd)
	}

	var n namespace.Namespace
	if err := yaml.Unmarshal(fc, &n); err != nil {
		_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Encountered yaml unmarshal error for \"%s\": %+v\n", fn, err)
		return nil, cmdx.FailSilently(cmd)
	}

	return &n, nil
}
