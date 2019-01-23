init:
		go get -u \
			github.com/ory/x/tools/listx \
			github.com/sqs/goreturns \
			github.com/go-swagger/go-swagger/cmd/swagger

format:
		goreturns -w -local github.com/ory $$(listx .)
		# goimports -w -v -local github.com/ory $$(listx .)

swagger:
		swagger generate spec -m -o ./docs/api.swagger.json

build-sdk:
		rm -rf ./sdk/go/keto/swagger
		rm -rf ./sdk/js/swagger
		rm -rf ./sdk/php/swagger

		java -jar scripts/swagger-codegen-cli-2.2.3.jar generate -i ./docs/api.swagger.json -l go -o ./sdk/go/keto/swagger
		java -jar scripts/swagger-codegen-cli-2.2.3.jar generate -i ./docs/api.swagger.json -l javascript -o ./sdk/js/swagger
		java -jar scripts/swagger-codegen-cli-2.2.3.jar generate -i ./docs/api.swagger.json -l php -o ./sdk/php/ \
			--invoker-package keto\\SDK --git-repo-id swagger --git-user-id ory --additional-properties "packagePath=swagger,description=Client for keto"

		git checkout HEAD -- sdk/go/keto/swagger/api_client.go

		# goreturns -w -i -local github.com/ory $$(listx ./sdk/go)

		rm -f ./sdk/js/swagger/package.json
		rm -rf ./sdk/js/swagger/test
		rm -f ./sdk/php/swagger/composer.json ./sdk/php/swagger/phpunit.xml.dist
		rm -rf ./sdk/php/swagger/test

install-stable:
		KETO_LATEST=$$(git describe --abbrev=0 --tags)
		git checkout $$KETO_LATEST
		GO111MODULE=on go install \
				-ldflags "-X github.com/ory/keto/cmd.Version=$$KETO_LATEST -X github.com/ory/keto/cmd.BuildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` -X github.com/ory/keto/cmd.GitHash=`git rev-parse HEAD`" \
				.
		git checkout master

install:
		GO111MODULE=on go install .
