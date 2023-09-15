// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ory/x/configx"
	"github.com/ory/x/logrusx"
	"github.com/stretchr/testify/require"
)

func TestNewOPLConfigWatcher(t *testing.T) {
	hits := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits++
		io.WriteString(w, testOPL)
	}))
	t.Cleanup(ts.Close)
	ctx := context.Background()
	cfg, err := NewDefault(ctx, nil, logrusx.New("", ""), configx.SkipValidation())
	require.NoError(t, err)
	cw, err := newOPLConfigWatcher(ctx, cfg, ts.URL)
	require.NoError(t, err)
	require.Equal(t, 1, hits, "HTTP request made")
	_, err = cw.GetNamespaceByName(ctx, "User")
	require.NoError(t, err)
	_, err = cw.GetNamespaceByName(ctx, "Document")
	require.NoError(t, err)

	cache.Wait()

	cw, err = newOPLConfigWatcher(ctx, cfg, ts.URL)
	require.NoError(t, err)
	require.Equal(t, 1, hits, "content was cached")
	_, err = cw.GetNamespaceByName(ctx, "User")
	require.NoError(t, err)
	_, err = cw.GetNamespaceByName(ctx, "Document")
	require.NoError(t, err)
}

var testOPL = `
import { Namespace } from "@ory/keto-namespace-types"

class User implements Namespace {}
class Document implements Namespace {}
`
