// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package configx

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/pflag"
)

// RegisterFlags registers the config file flag.
func RegisterFlags(flags *pflag.FlagSet) {
	flags.StringSliceP("config", "c", []string{}, "Path to one or more .json, .yaml, .yml, .toml config files. Values are loaded in the order provided, meaning that the last config file overwrites values from the previous config file.")
}

// host = unix:/path/to/socket => port is discarded, otherwise format as host:port
func GetAddress(host string, port int) string {
	if strings.HasPrefix(host, "unix:") {
		return host
	}
	return fmt.Sprintf("%s:%d", host, port)
}

func (s *Serve) GetAddress() string {
	return GetAddress(s.Host, s.Port)
}

// AddSchemaResources adds the config schema partials to the compiler.
// The interface is specified instead of `jsonschema.Compiler` to allow the use of any jsonschema library fork or version.
func AddSchemaResources(c interface {
	AddResource(url string, r io.Reader) error
}) error {
	if err := c.AddResource(TLSConfigSchemaID, bytes.NewReader(TLSConfigSchema)); err != nil {
		return err
	}
	if err := c.AddResource(ServeConfigSchemaID, bytes.NewReader(ServeConfigSchema)); err != nil {
		return err
	}
	return c.AddResource(CORSConfigSchemaID, bytes.NewReader(CORSConfigSchema))
}

func cleanPrefix(prefix string) string {
	if len(prefix) > 0 {
		prefix = strings.TrimRight(prefix, ".") + "."
	}
	return prefix
}
