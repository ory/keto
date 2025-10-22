SHELL=/bin/bash -o pipefail

export PWD				:= $(shell pwd)
export PATH				:= ${PWD}/.bin:${PATH}
export IMAGE_TAG		:= $(if $(IMAGE_TAG),$(IMAGE_TAG),latest)

SCRIPT_DEPENDENCIES = grype \
					trivy \
					ory \
					licenses

define make-script-dependency
  # each script is responsible to figure out whether it should re-install
  .PHONY: .bin/$1
  .bin/$1:
		./scripts/install-$1.sh
endef
$(foreach dep, $(SCRIPT_DEPENDENCIES), $(eval $(call make-script-dependency,$(dep))))

.PHONY: .bin/clidoc
.bin/clidoc:
	go build -o .bin/clidoc ./cmd/clidoc/.

authors:  # updates the AUTHORS file
	curl https://raw.githubusercontent.com/ory/ci/master/authors/authors.sh | env PRODUCT="Ory Keto" bash

.PHONY: format
format: .bin/ory node_modules
	.bin/ory dev headers copyright --type=open-source --exclude=.bin --exclude=internal/httpclient --exclude=proto --exclude=oryx
	go tool goimports -w -local github.com/ory/keto *.go internal cmd contrib ketoctx ketoapi embedx
	npm exec -- prettier --write .
	go tool buf format -w

.PHONY: install
install:
	go install -tags sqlite .

.PHONY: docker
docker:
	DOCKER_BUILDKIT=1 DOCKER_CONTENT_TRUST=1 docker build --progress=plain -t oryd/keto:${IMAGE_TAG} -f .docker/Dockerfile-build .

# Generates the SDKs
.PHONY: sdk
sdk: buf .bin/ory node_modules
	rm -rf internal/httpclient
	.bin/ory dev swagger sanitize ./spec/api.swagger.json
	sed -i -f ./.schema/openapi/patches/replacements.sed ./spec/api.swagger.json
	go tool swagger validate ./spec/api.swagger.json
	CIRCLE_PROJECT_USERNAME=ory CIRCLE_PROJECT_REPONAME=keto \
		.bin/ory dev openapi migrate \
			--health-path-tags metadata \
			-p https://raw.githubusercontent.com/ory/x/master/healthx/openapi/patch.yaml \
			-p file://.schema/openapi/patches/meta.yaml \
			-p file://.schema/openapi/patches/checkServices.yaml \
			spec/api.swagger.json spec/api.json

	mkdir -p internal/httpclient

	npm run openapi-generator-cli -- generate -i "spec/api.json" \
		-g go \
		-o "internal/httpclient" \
		--git-user-id ory \
		--git-repo-id keto-client-go \
		--git-host github.com \
		--api-name-suffix "Api" \
		--global-property apiTests=false \
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
buf-gen: node_modules
	go tool buf generate proto
	make format
	@echo "All code was generated successfully!"

#
# Lint API definitions
#
.PHONY: buf-lint
buf-lint:
	go tool buf lint ./proto
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
test-docs-samples:
	go tool -n jd # Apparently on the first run the path is the temporary build output and will be deleted again. Later invocations use the correct go build cache path.
	PATH=$$PATH:$$(dirname "$$(go tool -n jd)") && \
		(cd ./contrib/docs-code-samples && \
		npm i && \
		npm test)

.PHONY: fuzz-test
fuzz-test:
	go test -tags=sqlite -fuzz=FuzzParser -fuzztime=10s ./internal/schema

.PHONY: libfuzzer-fuzz-test
libfuzzer-fuzz-test: .bin/go114-fuzz-build
	mkdir -p .fuzzer/fuzz_parser_corpus
	.bin/go114-fuzz-build -o ./.fuzzer/parser.a -func LibfuzzerFuzzParser ./internal/schema
	clang -fsanitize=fuzzer ./.fuzzer/parser.a -o ./.fuzzer/parser
	./.fuzzer/parser -use_value_profile=1 -timeout=1 ./.fuzzer/fuzz_parser_corpus ./.fuzzer/fuzz_parser_seeds

.PHONY: libfuzzer-fuzz-test-minimize
libfuzzer-fuzz-test-minimize: .bin/go114-fuzz-build
	mkdir -p .fuzzer/fuzz_parser_corpus
	mv .fuzzer/fuzz_parser_corpus .fuzzer/fuzz_parser_old_corpus
	mkdir -p .fuzzer/fuzz_parser_corpus
	.bin/go114-fuzz-build -o ./.fuzzer/parser.a -func LibfuzzerFuzzParser ./internal/schema
	clang -fsanitize=fuzzer ./.fuzzer/parser.a -o ./.fuzzer/parser
	./.fuzzer/parser -runs=0 -merge=1 ./.fuzzer/fuzz_parser_corpus ./.fuzzer/fuzz_parser_seeds ./.fuzzer/fuzz_parser_old_corpus

.PHONY: cve-scan
cve-scan: docker .bin/grype
	grype oryd/keto:latest

.PHONY: pre-release
pre-release:
	go tool yq '.services.keto.image = "oryd/keto:'$$DOCKER_TAG'"' -i docker-compose.yml
	go tool yq '.services.keto-migrate.image = "oryd/keto:'$$DOCKER_TAG'"' -i docker-compose-mysql.yml
	go tool yq '.services.keto-migrate.image = "oryd/keto:'$$DOCKER_TAG'"' -i docker-compose-postgres.yml

.PHONY: post-release
post-release:
	echo "nothing to do"

.PHONY: generate
generate: .bin/stringer
	go generate ./...
	make format

licenses: .bin/licenses node_modules  # checks open-source licenses
	.bin/licenses

node_modules: package-lock.json
	npm ci
	touch node_modules
