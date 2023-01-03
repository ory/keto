// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package namespace

import (
	"embed"
	"io"
	"text/template"

	"github.com/pkg/errors"
)

//go:embed config_template/*
var configTemplate embed.FS

// GenerateOPLConfig derives an Ory Permission Language config from the
// namespaces and writes it to out. The OPL config is functionally equivalent to
// the list of namespaces.
func GenerateOPLConfig(namespaces []string, out io.Writer) error {
	t, err := template.New("config_template").ParseFS(configTemplate, "config_template/*")
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(t.ExecuteTemplate(out,
		"namespaces.ts.tmpl",
		struct{ Namespaces []string }{Namespaces: namespaces}))
}
