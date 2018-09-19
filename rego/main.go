package main

import (
	"fmt"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"context"
	"runtime/debug"
	"path/filepath"
	"os"
	"github.com/pkg/errors"
	"io/ioutil"
	"github.com/ory/ladon"
	"github.com/open-policy-agent/opa/storage/inmem"
	"github.com/open-policy-agent/opa/topdown"
)

func main() {
	files, err := loadFiles(".")
	if err != nil {
		panic(err)
	}

	compiler, err := newCompiler(files)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"policies": map[string]interface{}{"exact": ladon.Policies{}},
	}

	topdown.RegisterFunctionalBuiltin2()
	tracer := topdown.NewBufferTracer()
	store := inmem.NewFromObject(data)

	r := rego.New(
		rego.Query("policies.exact"),
		rego.Compiler(compiler),
		rego.Input(interface{}("data.ladon.allowed_exact")),
		rego.Store(store),
		rego.Tracer(tracer),
	)

	// Run evaluation.
	rs, err := r.Eval(context.TODO())
	if err != nil {
		panic(err)
	}

	if len(rs) > 0 {
		panic(fmt.Sprintf("It's suspicious that a result was found, got %d results: %+v", len(rs), rs))
	}
}

func loadFiles(directory string) (map[string][]byte, error) {
	m := map[string][]byte{}
	if err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.WithStack(err)
		}

		if info.IsDir() {
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

func newCompiler(files map[string][]byte) (*ast.Compiler, error) {
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
