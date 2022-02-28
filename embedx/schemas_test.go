package embedx

import (
	"context"
	"testing"

	"github.com/ory/jsonschema/v3"
	"github.com/stretchr/testify/require"
)

func TestConfigSchema(t *testing.T) {
	c := jsonschema.NewCompiler()
	require.NoError(t, AddConfigSchema(c))

	_, err := c.Compile(context.Background(), ConfigSchemaID)
	require.NoError(t, err)
}
