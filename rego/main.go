package main

import (
	"fmt"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"context"
	"path/filepath"
	"os"
	"github.com/pkg/errors"
	"io/ioutil"
	"github.com/open-policy-agent/opa/storage/inmem"
	"github.com/open-policy-agent/opa/topdown"
	"github.com/ory/ladon"
	"encoding/json"
	"strings"
)

type DataModel struct {
	Store DataModelStore `json:"store"`
}
type DataModelLadon struct {
	Exact DataModelLadonAll `json:"exact"`
	Regex DataModelLadonAll `json:"regex"`
}
type DataModelLadonAll struct {
	Policies ladon.Policies      `json:"policies"`
	Roles    map[string][]string `json:"roles"`
}
type DataModelStore struct {
	Ladon DataModelLadon `json:"ladon"`
}

func main() {
	files, err := loadFiles(".")
	if err != nil {
		panic(err)
	}

	compiler, err := newCompiler(files)
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(&DataModel{
		Store: DataModelStore{
			Ladon: DataModelLadon{
				Exact: DataModelLadonAll{
					Policies: ladon.Policies{
						&ladon.DefaultPolicy{
							Actions:    []string{"actions:1"},
							Subjects:   []string{"subjects:1"},
							Resources:  []string{"resources:1"},
							Conditions: ladon.Conditions{},
							Effect:     ladon.AllowAccess,
						},
					},
					Roles: map[string][]string{},
				},
				Regex: DataModelLadonAll{},
			},
		},
	})
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{}
	if err := json.Unmarshal(b, &DataModel{
		Store: DataModelStore{
			Ladon: DataModelLadon{
				Exact: DataModelLadonAll{
					Policies: ladon.Policies{
						&ladon.DefaultPolicy{
							Actions:    []string{"actions:1"},
							Subjects:   []string{"subjects:1"},
							Resources:  []string{"resources:1"},
							Conditions: ladon.Conditions{},
							Effect:     ladon.AllowAccess,
						},
					},
					Roles: map[string][]string{},
				},
				Regex: DataModelLadonAll{},
			},
		},
	}); err != nil {
		panic(err)
	}

	tracer := topdown.NewBufferTracer()
	store := inmem.NewFromObject(data)

	r := rego.New(
		rego.Query("data.ladon.exact.allow"),
		rego.Compiler(compiler),
		rego.Store(store),
		rego.Tracer(tracer),
		rego.Input(
			map[string]interface{}{
				"action":   "actions:1",
				"subject":  "subjects:1",
				"resource": "resources:1",
			},
		),
	)

	// Run evaluation.
	rs, err := r.Eval(context.TODO())
	if err != nil {
		panic(err)
	}

	//for k, e := range *tracer {
	//	fmt.Printf("Got tracer event (%d): %s\n", k, e)
	//}

	if len(rs) != 1 || len(rs[0].Expressions) != 1 {
		panic(fmt.Sprintf("Expected exactly one result, got %d - %+v", len(rs), rs))
	}

	result, ok := rs[0].Expressions[0].Value.(bool)
	if !ok {
		panic(fmt.Sprintf("Expected bool but got %+v", rs[0].Expressions[0].Value))
	}

	fmt.Printf("Got result: %v\n", result)
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
