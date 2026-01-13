// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ory/x/configx"
	"github.com/ory/x/logrusx"
)

func TestNewOPLConfigWatcher(t *testing.T) {
	var hits atomic.Int64
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		if _, err := io.WriteString(w, testOPL); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
	t.Cleanup(ts.Close)
	ctx := context.Background()
	cfg, err := NewDefault(ctx, nil, logrusx.New("", ""), configx.SkipValidation())
	require.NoError(t, err)
	cw, err := newOPLConfigWatcher(ctx, cfg, ts.URL)
	require.NoError(t, err)
	require.EqualValues(t, 1, hits.Load(), "HTTP request made")
	_, err = cw.GetNamespaceByName(ctx, "User")
	require.NoError(t, err)
	_, err = cw.GetNamespaceByName(ctx, "Document")
	require.NoError(t, err)

	cache.Wait()

	cw, err = newOPLConfigWatcher(ctx, cfg, ts.URL)
	require.NoError(t, err)
	require.EqualValues(t, 1, hits.Load(), "content was cached")
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
