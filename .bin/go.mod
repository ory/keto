module github.com/ory/keto/.bin

go 1.16

replace github.com/goreleaser/nfpm => github.com/goreleaser/nfpm v1.10.2

replace github.com/ory/kratos-client-go => github.com/ory/kratos-client-go v0.5.4-alpha.1.0.20210210170256-960b093d8bf9

replace github.com/ory/kratos/corp => github.com/ory/kratos/corp v0.0.0-20210118092700-c2358be1e867

replace github.com/oleiade/reflections => github.com/oleiade/reflections v1.0.1

require (
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/Microsoft/go-winio v0.4.15 // indirect
	github.com/bufbuild/buf v0.31.1
	github.com/fatih/color v1.10.0 // indirect
	github.com/go-swagger/go-swagger v0.26.1
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/jmoiron/sqlx v1.2.1-0.20190826204134-d7d95172beb5 // indirect
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/ory/cli v0.0.49
	github.com/pseudomuto/protoc-gen-doc v1.4.1
	github.com/rogpeppe/go-internal v1.6.1 // indirect
	github.com/smartystreets/assertions v1.0.0 // indirect
	golang.org/x/tools v0.1.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.26.0
	honnef.co/go/tools v0.0.1-2020.1.6 // indirect
)
