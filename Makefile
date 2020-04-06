.PHONY: format
format:
		goreturns -w -local github.com/ory $$(listx .)

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

# Generates the SDKs
.PHONY: sdk
sdk:
		$$(go env GOPATH)/bin/swagger generate spec -m -o ./.schemas/api.swagger.json -x internal/httpclient
		$$(go env GOPATH)/bin/swagutil sanitize ./.schemas/api.swagger.json
		$$(go env GOPATH)/bin/swagger flatten --with-flatten=remove-unused -o ./.schemas/api.swagger.json ./.schemas/api.swagger.json
		$$(go env GOPATH)/bin/swagger validate ./.schemas/api.swagger.json
		rm -rf internal/httpclient
		mkdir -p internal/httpclient
		$$(go env GOPATH)/bin/swagger generate client -f ./.schemas/api.swagger.json -t internal/httpclient -A Ory_Keto
		make format
