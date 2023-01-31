package validate_test

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/x/validate"
)

func toURL(t *testing.T, s string) *url.URL {
	u, err := url.Parse(s)
	require.NoError(t, err)
	return u
}

func TestValidateNoExtraParams(t *testing.T) {
	for _, tt := range []struct {
		name      string
		req       *http.Request
		assertErr assert.ErrorAssertionFunc
	}{
		{
			name:      "empty",
			req:       &http.Request{URL: toURL(t, "https://example.com")},
			assertErr: assert.NoError,
		},
		{
			name:      "all params",
			req:       &http.Request{URL: toURL(t, "https://example.com?foo=1&bar=baz")},
			assertErr: assert.NoError,
		},
		{
			name:      "extra params",
			req:       &http.Request{URL: toURL(t, "https://example.com?foo=1&bar=2&baz=3")},
			assertErr: assert.Error,
		},
	} {
		t.Run("case="+tt.name, func(t *testing.T) {
			err := validate.NoExtraQueryParams(tt.req, "foo", "bar")
			tt.assertErr(t, err)
		})
	}
}

func TestValidateHasEmptyBody(t *testing.T) {
	for _, tt := range []struct {
		name      string
		req       *http.Request
		assertErr assert.ErrorAssertionFunc
	}{
		{
			name:      "empty body",
			req:       &http.Request{Body: io.NopCloser(strings.NewReader(""))},
			assertErr: assert.NoError,
		},
		{
			name:      "non-empty body",
			req:       &http.Request{Body: io.NopCloser(strings.NewReader("content"))},
			assertErr: assert.Error,
		},
	} {
		t.Run("case="+tt.name, func(t *testing.T) {
			err := validate.HasEmptyBody(tt.req)
			tt.assertErr(t, err)
		})
	}

}
