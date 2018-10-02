package engine

import (
	"github.com/open-policy-agent/opa/ast"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func walk(directory string) (map[string][]byte, error) {
	m := map[string][]byte{}
	if err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.WithStack(err)
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".rego" {
			return nil
		}

		if strings.Contains(path, "_test.rego") {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return errors.WithStack(err)
		}

		d, err := ioutil.ReadAll(f)
		if err != nil {
			return errors.WithStack(err)
		}

		m[path] = d

		return nil
	}); err != nil {
		return nil, err
	}

	return m, nil
}

func NewCompiler(directory string) (*ast.Compiler, error) {
	files, err := walk(directory)
	if err != nil {
		return nil, err
	}

	modules := map[string]*ast.Module{}
	for file, content := range files {
		parsed, err := ast.ParseModule(file, string(content))
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
