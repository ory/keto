SHELL=/bin/bash -o pipefail

EXECUTABLES = docker-compose docker node npm go
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))

export GO111MODULE := on
export PATH := .bin:${PATH}

.PHONY: deps
deps:
ifneq ("$(shell base64 Makefile) $(shell base64 go.mod) $(shell base64 go.sum)","$(shell cat .bin/.lock)")
		go build -o .bin/go-acc github.com/ory/go-acc
		go build -o .bin/goreturns github.com/sqs/goreturns
		go build -o .bin/mockgen github.com/golang/mock/mockgen
		go build -o .bin/swagger github.com/go-swagger/go-swagger/cmd/swagger
		go build -o golang.org/x/tools/cmd/goimports
		go build -o .bin/ory github.com/ory/cli
		go build -o .bin/go-bindata github.com/go-bindata/go-bindata/go-bindata
		go build -o .bin/buf github.com/bufbuild/buf/cmd/buf
		go build -o .bin/protoc-gen-go google.golang.org/protobuf/cmd/protoc-gen-go
		go build -o .bin/protoc-gen-go-grpc google.golang.org/grpc/cmd/protoc-gen-go-grpc
		npm i
		echo "v0" > .bin/.lock
		echo "$$(base64 Makefile) $$(base64 go.mod) $$(base64 go.sum)" > .bin/.lock
endif

.PHONY: format
format:
		goimports -w -local github.com/ory/keto *.go internal cmd contrib

.PHONY: install-stable
install-stable: deps
		KETO_LATEST=$$(git describe --abbrev=0 --tags)
		git checkout $$KETO_LATEST
		GO111MODULE=on go install \
				-ldflags "-X github.com/ory/keto/cmd.Version=$$KETO_LATEST -X github.com/ory/keto/cmd.Date=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` -X github.com/ory/keto/cmd.Commit=`git rev-parse HEAD`" \
				.
		rm pkged.go
		git checkout master

.PHONY: install
install: deps
		GO111MODULE=on go install .

# Generates the SDKs
.PHONY: sdk
sdk: deps
		swagger generate spec -m -o ./.schema/api.swagger.json -x internal/httpclient
		ory dev swagger sanitize ./.schema/api.swagger.json
		swagger flatten --with-flatten=remove-unused -o ./.schema/api.swagger.json ./.schema/api.swagger.json
		swagger validate ./.schema/api.swagger.json
		rm -rf internal/httpclient
		mkdir -p internal/httpclient
		swagger generate client -f ./.schema/api.swagger.json -t internal/httpclient -A Ory_Keto
		make format

.PHONY: build
build: deps
		go build -tags sqlite

#
# Generate APIs and client stubs from the definitions
#
.PHONY: buf-gen
buf-gen: deps
		buf generate \
		&& \
		echo "TODO: generate gapic client at ./client" \
		&& \
		echo "All code was generated successfully!"

#
# Lint API definitions
#
.PHONY: buf-lint
buf-lint: deps
		buf check lint \
		&& \
		echo "All lint checks passed successfully!"

#
# Generate after linting succeeded
#
.PHONY: buf
buf: buf-lint buf-gen

.PHONY: reset-testdb
reset-testdb:
		source scripts/test-resetdb.sh

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
