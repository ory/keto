SHELL=/bin/bash -o pipefail

export PWD := $(shell pwd)
export PATH := ${PWD}/.bin:${PATH}

GO_DEPENDENCIES = golang.org/x/tools/cmd/goimports \
				  github.com/mattn/goveralls \
				  github.com/ory/go-acc \
				  github.com/bufbuild/buf/cmd/buf \
				  google.golang.org/protobuf/cmd/protoc-gen-go \
				  google.golang.org/grpc/cmd/protoc-gen-go-grpc \
				  github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
				  github.com/josephburnett/jd \
				  github.com/mikefarah/yq/v4 \
				  golang.org/x/tools/cmd/stringer \
				  github.com/mdempsky/go114-fuzz-build

SCRIPT_DEPENDENCIES = swagger \
					protoc \
					grype \
					trivy \
					ory \
					protoc-gen-js

define make-go-dependency
  # go install is responsible for not re-building when the code hasn't changed
  .bin/$2: .bin/go.sum
		cd .bin; GOBIN=$(PWD)/.bin go install $1
endef
$(foreach dep, $(GO_DEPENDENCIES), $(eval $(call make-go-dependency,$(dep),$(notdir $(dep)))))

define make-script-dependency
  # each script is responsible to figure out whether it should re-install
  .PHONY: .bin/$1
  .bin/$1:
		./scripts/install-$1.sh
endef
$(foreach dep, $(SCRIPT_DEPENDENCIES), $(eval $(call make-script-dependency,$(dep))))

.bin/yq: .bin/go.mod .bin/go.sum
	cd .bin; GOBIN=$(PWD)/.bin go install github.com/mikefarah/yq/v4

.PHONY: .bin/clidoc
.bin/clidoc:
	go build -o .bin/clidoc ./cmd/clidoc/.

.PHONY: format
format: .bin/ory .bin/goimports node_modules
	.bin/ory dev headers copyright --type=open-source --exclude=.bin --exclude=internal/httpclient --exclude=proto
	.bin/goimports -w -local github.com/ory/keto *.go internal cmd contrib ketoctx ketoapi embedx
	npm exec -- prettier --write .

.PHONY: install
install:
	go install -tags sqlite .

.PHONY: docker
docker:
	docker build -t oryd/keto:latest -f .docker/Dockerfile-build .

# Generates the SDKs
.PHONY: sdk
sdk: .bin/swagger .bin/ory node_modules
	rm -rf internal/httpclient
	swagger generate spec -m -o spec/swagger.json \
		-c github.com/ory/keto \
		-c github.com/ory/x/healthx \
		-x internal/httpclient \
		-x internal/e2e
	.bin/ory dev swagger sanitize ./spec/swagger.json
	swagger validate ./spec/swagger.json
	CIRCLE_PROJECT_USERNAME=ory CIRCLE_PROJECT_REPONAME=keto \
		.bin/ory dev openapi migrate \
			--health-path-tags metadata \
			-p https://raw.githubusercontent.com/ory/x/master/healthx/openapi/patch.yaml \
			-p file://.schema/openapi/patches/meta.yaml \
			spec/swagger.json spec/api.json

	mkdir -p internal/httpclient

	npm run openapi-generator-cli -- generate -i "spec/api.json" \
		-g go \
		-o "internal/httpclient" \
		--git-user-id ory \
		--git-repo-id keto-client-go \
		--git-host github.com \
		-t .schema/openapi/templates/go \
		-c .schema/openapi/gen.go.yml

	rm internal/httpclient/go.{mod,sum}

	make format

.PHONY: build
build:
	go build -tags sqlite

#
# Generate APIs and client stubs from the definitions
#
.PHONY: buf-gen
buf-gen: .bin/buf .bin/protoc .bin/protoc-gen-go .bin/protoc-gen-go-grpc .bin/protoc-gen-js .bin/protoc-gen-doc node_modules
	buf generate proto
	make format
	@echo "All code was generated successfully!"

#
# Lint API definitions
#
.PHONY: buf-lint
buf-lint: .bin/buf
	cd proto; buf lint
	@echo "All lint checks passed successfully!"

#
# Generate after linting succeeded
#
.PHONY: buf
buf: buf-lint buf-gen

.PHONY: test-e2e
test-e2e:
	go test -tags sqlite -failfast -v ./internal/e2e

.PHONY: test-docs-samples
test-docs-samples: .bin/jd
	cd ./contrib/docs-code-samples \
	&& \
	npm i \
	&& \
	npm test

.PHONY: fuzz-test
fuzz-test:
	go test -tags=sqlite -fuzz=FuzzParser -fuzztime=10s ./internal/schema

.PHONY: libfuzzer-fuzz-test
libfuzzer-fuzz-test: .bin/go114-fuzz-build
	mkdir -p .fuzzer
	.bin/go114-fuzz-build -o ./.fuzzer/parser.a ./internal/schema
	clang -fsanitize=fuzzer ./.fuzzer/parser.a -o ./.fuzzer/parser
	./.fuzzer/parser -timeout=1 -max_total_time=10 -use_value_profile

.PHONY: cve-scan
cve-scan: docker .bin/grype
	grype oryd/keto:latest

.PHONY: post-release
post-release: .bin/yq
	cat docker-compose.yml | yq '.services.keto.image = "oryd/keto:'$$DOCKER_TAG'"' | sponge docker-compose.yml
	cat docker-compose-mysql.yml | yq '.services.keto-migrate.image = "oryd/keto:'$$DOCKER_TAG'"' | sponge docker-compose-mysql.yml
	cat docker-compose-postgres.yml | yq '.services.keto-migrate.image = "oryd/keto:'$$DOCKER_TAG'"' | sponge docker-compose-postgres.yml

.PHONY: generate
generate: .bin/stringer
	go generate ./...
	make format

licenses: .bin/licenses node_modules  # checks open-source licenses
	.bin/licenses

.bin/licenses: Makefile
	curl https://raw.githubusercontent.com/ory/ci/master/licenses/install | sh

.bin/ory: Makefile
	curl https://raw.githubusercontent.com/ory/meta/master/install.sh | bash -s -- -b .bin ory v0.1.48
	touch .bin/ory

node_modules: package-lock.json
	npm ci
	touch node_modules
