package namespace

import (
	"embed"
	"github.com/ory/keto/internal/driver/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const FlagOut = "out"

func registerOutputFlag(cmd *cobra.Command) {
	cmd.Flags().StringP(FlagOut, "o", ".", "output directory, will be created if it does not exist")
}

//go:embed config_template/*
var configTemplate embed.FS
var version string

func init() {
	version = config.Version
	if version == "master" || version == "" {
		version = "latest"
	}
}

func generateConfigFiles(nspaces []string, out string) error {
	t, err := template.New("config_template").ParseFS(configTemplate, "config_template/*")
	if err != nil {
		return errors.WithStack(err)
	}
	return fs.WalkDir(configTemplate, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return errors.WithStack(err)
		}
		if d.IsDir() {
			return nil
		}
		orig, _ := configTemplate.Open(path)
		defer orig.Close()
		other, err := os.Create(filepath.Join(out, strings.TrimSuffix(d.Name(), ".tmpl")))
		if err != nil {
			return err
		}
		defer other.Close()
		if !strings.HasSuffix(path, ".tmpl") {
			_, err := io.Copy(other, orig)
			if err != nil {
				return errors.WithStack(err)
			}
			return nil
		}

		return errors.WithStack(t.ExecuteTemplate(other, d.Name(), struct {
			Namespaces []string
			Version    string
		}{
			Namespaces: nspaces,
			Version:    version,
		}))
	})
}
