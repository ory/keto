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
				  github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
				  github.com/ory/cli \
				  github.com/anchore/grype \
				  golang.org/x/tools/cmd/stringer

define make-go-dependency
  # go install is responsible for not re-building when the code hasn't changed
  .bin/$2: .bin/go.mod .bin/go.sum Makefile
		cd .bin; GOBIN=$(PWD)/.bin/ go install $1
endef
$(foreach dep, $(GO_DEPENDENCIES), $(eval $(call make-go-dependency,$(dep),$(notdir $(dep)))))

# versioned package paths
$(eval $(call make-go-dependency,github.com/mikefarah/yq/v4,yq))

.bin/ory: Makefile
		bash <(curl https://raw.githubusercontent.com/ory/meta/master/install.sh) -d -b .bin ory v0.1.0
		touch -a -m .bin/ory

node_modules: package.json package-lock.json Makefile
		npm ci

.PHONY: .bin/clidoc
.bin/clidoc:
		go build -o .bin/clidoc ./cmd/clidoc/.

docs/cli: .bin/clidoc
		clidoc .

.PHONY: format
format: .bin/goimports node_modules
		goimports -w -local github.com/ory/keto *.go internal cmd contrib ketoctx embedx
		npm run format

.PHONY: install
install:
		go install -tags sqlite .

.PHONY: docker
docker:
		docker build -t oryd/keto:latest -f .docker/Dockerfile-build .

# Generates the SDKs
.PHONY: sdk
sdk: .bin/swagger .bin/ory node_modules
		rm -rf internal/httpclient internal/httpclient-next
		swagger generate spec -m -o spec/swagger.json \
			-c github.com/ory/keto \
			-c github.com/ory/x/healthx \
			-x internal/httpclient \
			-x internal/e2e
		ory dev swagger sanitize ./spec/swagger.json
		swagger validate ./spec/swagger.json
		CIRCLE_PROJECT_USERNAME=ory CIRCLE_PROJECT_REPONAME=keto \
				ory dev openapi migrate \
					--health-path-tags metadata \
					-p https://raw.githubusercontent.com/ory/x/master/healthx/openapi/patch.yaml \
					-p file://.schema/openapi/patches/meta.yaml \
					spec/swagger.json spec/api.json

		mkdir -p internal/httpclient internal/httpclient-next
		swagger generate client -f ./spec/swagger.json -t internal/httpclient -A Ory_Keto

		npm run openapi-generator-cli -- generate -i "spec/api.json" \
				-g go \
				-o "internal/httpclient-next" \
				--git-user-id ory \
				--git-repo-id keto-client-go \
				--git-host github.com \
				-t .schema/openapi/templates/go \
				-c .schema/openapi/gen.go.yml

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

.PHONY: cve-scan
cve-scan: docker .bin/grype
		grype oryd/keto:latest

.PHONY: post-release
post-release: .bin/yq
		cat docker-compose.yml | yq '.services.keto.image = "oryd/keto:'$$DOCKER_TAG'"' | sponge docker-compose.yml
		cat docker-compose-mysql.yml | yq '.services.keto-migrate.image = "oryd/keto:'$$DOCKER_TAG'"' | sponge docker-compose-mysql.yml
		cat docker-compose-postgres.yml | yq '.services.keto-migrate.image = "oryd/keto:'$$DOCKER_TAG'"' | sponge docker-compose-postgres.yml

.PHONY: generate
generate: .bin/swagger .bin/stringer
		PATH=$(PWD)/.bin:${PATH} go generate ./...