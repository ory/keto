SHELL=/bin/bash -o pipefail

export PATH := .bin:${PATH}
export PWD := $(shell pwd)

GO_DEPENDENCIES = github.com/go-swagger/go-swagger/cmd/swagger \
				  golang.org/x/tools/cmd/goimports \
				  github.com/mattn/goveralls \
				  github.com/ory/cli \
				  github.com/ory/go-acc \
				  github.com/bufbuild/buf/cmd/buf \
				  google.golang.org/protobuf/cmd/protoc-gen-go \
				  google.golang.org/grpc/cmd/protoc-gen-go-grpc \
				  github.com/goreleaser/godownloader \
				  github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

define make-go-dependency
  # go install is responsible for not re-building when the code hasn't changed
  .bin/$(notdir $1): go.mod go.sum Makefile
		GOBIN=$(PWD)/.bin/ go install $1
endef
$(foreach dep, $(GO_DEPENDENCIES), $(eval $(call make-go-dependency, $(dep))))
$(call make-lint-dependency)

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
sdk: .bin/swagger .bin/cli
		swagger generate spec -m -o ./spec/api.json -x internal/httpclient -x proto/ory/keto -x docker
		cli dev swagger sanitize ./spec/api.json
		swagger flatten --with-flatten=remove-unused -o ./spec/api.json ./spec/api.json
		swagger validate ./spec/api.json
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
