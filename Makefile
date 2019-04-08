.PHONY: format
format:
		goreturns -w -local github.com/ory $$(listx .)
		# goimports -w -v -local github.com/ory $$(listx .)

.PHONY: swagger
swagger:
		swagger generate spec -m -o ./docs/api.swagger.json

.PHONY: sdk
sdk:
		GO111MODULE=on go mod tidy
		GO111MODULE=on go mod vendor
		GO111MODULE=off swagger generate spec -m -o ./docs/api.swagger.json
		GO111MODULE=off swagger validate ./docs/api.swagger.json

		rm -rf ./sdk/go/keto/*
		rm -rf ./sdk/js/swagger
		rm -rf ./sdk/php/swagger

		GO111MODULE=off swagger generate client -f ./docs/api.swagger.json -t sdk/go/keto -A Ory_Keto

		java -jar scripts/swagger-codegen-cli-2.2.3.jar generate -i ./docs/api.swagger.json -l javascript -o ./sdk/js/swagger
		java -jar scripts/swagger-codegen-cli-2.2.3.jar generate -i ./docs/api.swagger.json -l php -o ./sdk/php/ \
			--invoker-package keto\\SDK --git-repo-id swagger --git-user-id ory --additional-properties "packagePath=swagger,description=Client for keto"

		make format

		rm -f ./sdk/js/swagger/package.json
		rm -rf ./sdk/js/swagger/test
		rm -f ./sdk/php/swagger/composer.json ./sdk/php/swagger/phpunit.xml.dist
		rm -rf ./sdk/php/swagger/test
		rm -rf ./vendor

.PHONY: install-stable
install-stable:
		KETO_LATEST=$$(git describe --abbrev=0 --tags)
		git checkout $$KETO_LATEST
		$(go env GOPATH)/bin/packr
		GO111MODULE=on go install \
				-ldflags "-X github.com/ory/keto/cmd.Version=$$KETO_LATEST -X github.com/ory/keto/cmd.Date=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` -X github.com/ory/keto/cmd.Commit=`git rev-parse HEAD`" \
				.
		$(go env GOPATH)/bin/packr clean
		git checkout master

.PHONY: install
install:
		$(go env GOPATH)/bin/packr
		GO111MODULE=on go install .
		$(go env GOPATH)/bin/packr clean
