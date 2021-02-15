package main

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
)

func TestExample(t *testing.T) {
	// capture errors from main()
	defer func() {
		require.Nil(t, recover())
	}()

	f, err := os.Create(filepath.Join(t.TempDir(), "mock_output"))
	require.NoError(t, err)

	os.Stdout, f = f, os.Stdout
	main()
	os.Stdout, f = f, os.Stdout

	out, err := ioutil.ReadFile(f.Name())
	require.NoError(t, err)
	_, _ = os.Stdout.Write(out)

	assert.Equal(t, string(out), "Successfully created tuple.\n")

	conn, err := grpc.Dial("127.0.0.1:4466", grpc.WithInsecure())
	require.NoError(t, err)

	client := acl.NewReadServiceClient(conn)
	resp, err := client.ListRelationTuples(context.Background(), &acl.ListRelationTuplesRequest{Query: &acl.ListRelationTuplesRequest_Query{Namespace: "messages"}})
	require.NoError(t, err)
	require.Len(t, resp.RelationTuples, 1)

	assert.Equal(t, "messages", resp.RelationTuples[0].Namespace)
	assert.Equal(t, "02y_15_4w350m3", resp.RelationTuples[0].Object)
	assert.Equal(t, "decypher", resp.RelationTuples[0].Relation)
	assert.Equal(t, "john", resp.RelationTuples[0].Subject.GetId())
}
