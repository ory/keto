module github.com/ory/keto/.bin

go 1.16

replace github.com/goreleaser/nfpm => github.com/goreleaser/nfpm v1.10.2

replace github.com/ory/kratos/corp => github.com/ory/kratos/corp v0.0.0-20210118092700-c2358be1e867

replace github.com/oleiade/reflections => github.com/oleiade/reflections v1.0.1

require (
	github.com/bufbuild/buf v0.31.1
	github.com/go-swagger/go-swagger v0.26.1
	github.com/goreleaser/godownloader v0.1.1-0.20200426152203-fd8ad8f7dd78
	github.com/mattn/goveralls v0.0.7
	github.com/ory/go-acc v0.2.6
	github.com/pseudomuto/protoc-gen-doc v1.4.1
	golang.org/x/tools v0.1.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.26.0
)
