SHELL=/bin/bash -o pipefail

export PATH := .bin:${PATH}
export PWD := $(shell pwd)

GO_DEPENDENCIES = github.com/go-swagger/go-swagger/cmd/swagger \
				  golang.org/x/tools/cmd/goimports \
				  github.com/mattn/goveralls \
				  github.com/ory/go-acc \
				  github.com/bufbuild/buf/cmd/buf \
				  google.golang.org/protobuf/cmd/protoc-gen-go \
				  google.golang.org/grpc/cmd/protoc-gen-go-grpc \
				  github.com/goreleaser/godownloader \
				  github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

define make-go-dependency
  # go install is responsible for not re-building when the code hasn't changed
  .bin/$(notdir $1): .bin/go.mod .bin/go.sum Makefile
		cd .bin; GOBIN=$(PWD)/.bin/ go install $1
endef
$(foreach dep, $(GO_DEPENDENCIES), $(eval $(call make-go-dependency, $(dep))))
$(call make-lint-dependency)

.bin/ory: Makefile
		bash <(curl https://raw.githubusercontent.com/ory/cli/master/install.sh) -b .bin v0.0.57
		touch -a -m .bin/ory

node_modules: package.json package-lock.json Makefile
		npm ci

.bin/clidoc:
		go build -o .bin/clidoc ./cmd/clidoc/.

docs/cli: .bin/clidoc
		clidoc .

.PHONY: format
format: .bin/goimports node_modules
		goimports -w -local github.com/ory/keto *.go internal cmd contrib
		npm run format

.PHONY: install
install:
		go install -tags sqlite .

.PHONY: docker
docker:
		docker build -t oryd/keto:latest -f .docker/Dockerfile-build .

# Generates the SDKs
.PHONY: sdk
sdk: .bin/swagger .bin/ory
		swagger generate spec -m -o ./spec/swagger.json -x internal/httpclient -x proto/ory/keto -x docker
		ory dev swagger sanitize ./spec/swagger.json
		swagger flatten --with-flatten=remove-unused -o ./spec/swagger.json ./spec/swagger.json
		swagger validate ./spec/swagger.json

		CIRCLE_PROJECT_USERNAME=ory CIRCLE_PROJECT_REPONAME=kratos \
			ory dev openapi migrate \
				--health-path-tags metadata \
				-p https://raw.githubusercontent.com/ory/x/master/healthx/openapi/patch.yaml \
				-p file://spec/patches/subjects.yml \
				spec/swagger.json spec/api.json

		rm -rf internal/httpclient
		mkdir -p internal/httpclient
		swagger generate client -f ./spec/api.json -t internal/httpclient -A Ory_Keto
		make format

.PHONY: build
build:
		go build -tags sqlite

#
# Generate APIs and client stubs from the definitions
#
.PHONY: buf-gen
buf-gen: .bin/buf .bin/protoc-gen-go .bin/protoc-gen-go-grpc .bin/protoc-gen-doc node_modules
		buf generate \
		&& \
		echo "All code was generated successfully!"

#
# Lint API definitions
#
.PHONY: buf-lint
buf-lint: .bin/buf
		buf check lint \
		&& \
		echo "All lint checks passed successfully!"

#
# Generate after linting succeeded
#
.PHONY: buf
buf: buf-lint buf-gen

.PHONY: test-e2e
test-e2e:
		go test -tags sqlite -failfast -v ./internal/e2e

.PHONY: test-docs-samples
test-docs-samples:
		cd ./contrib/docs-code-samples \
		&& \
		npm i \
		&& \
		npm test

.PHONY: migrations-render
migrations-render: .bin/ory
		ory dev pop migration render internal/persistence/sql/migrations/templates internal/persistence/sql/migrations/sql

.PHONY: migrations-render-replace
migrations-render-replace: .bin/ory
		ory dev pop migration render -r internal/persistence/sql/migrations/templates internal/persistence/sql/migrations/sql
