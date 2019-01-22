package engine

import (
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/open-policy-agent/opa/ast"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func walk(directory packr.Box, logger logrus.FieldLogger) (map[string]string, error) {
	m := map[string]string{}
	if err := directory.Walk(func(path string, file packr.File) error {
		if filepath.Ext(path) != ".rego" || filepath.Ext(path) != ".rego.go" {
			return nil
		}

		if strings.Contains(path, "_test.rego") || strings.Contains(path, "_test.rego.go") {
			return nil
		}

		m[path] = directory.String(path)
		logger.WithField("file", path).Debugf("Successfully loaded rego file")

		return nil
	}); err != nil {
		return nil, err
	}

	return m, nil
}

func NewCompiler(directory packr.Box, logger logrus.FieldLogger) (*ast.Compiler, error) {
	files, err := walk(directory, logger)
	if err != nil {
		return nil, err
	}

	modules := map[string]*ast.Module{}
	for file, content := range files {
		parsed, err := ast.ParseModule(file, content)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		modules[file] = parsed
	}

	compiler := ast.NewCompiler()
	compiler.Compile(modules)

	if compiler.Failed() {
		return nil, errors.Errorf("unable to compile module with payload: %s", compiler.Errors)
	}

	return compiler, nil
}
