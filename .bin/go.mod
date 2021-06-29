module github.com/ory/keto/.bin

go 1.16

replace github.com/goreleaser/nfpm => github.com/goreleaser/nfpm v1.10.2

replace github.com/ory/kratos-client-go => github.com/ory/kratos-client-go v0.6.3-alpha.1

replace github.com/ory/kratos/corp => github.com/ory/kratos/corp v0.0.0-20210118092700-c2358be1e867

replace github.com/oleiade/reflections => github.com/oleiade/reflections v1.0.1

replace github.com/ory/cli => github.com/ory/cli v0.0.57-0.20210629114108-ae1184abec67

replace github.com/gobuffalo/pop/v5 => github.com/gobuffalo/pop/v5 v5.3.4-0.20210608105745-bb07a373cc0e

replace github.com/mattn/go-sqlite3 => github.com/mattn/go-sqlite3 v1.14.7-0.20210414154423-1157a4212dcb

replace github.com/ory/kratos => github.com/ory/kratos v0.6.3-alpha.1.0.20210608145203-b5c1658e01ca

require (
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.7 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/bufbuild/buf v0.31.1
	github.com/go-swagger/go-swagger v0.26.1
	github.com/goreleaser/godownloader v0.1.1-0.20200426152203-fd8ad8f7dd78
	github.com/mattn/goveralls v0.0.7
	github.com/ory/cli v0.0.54
	github.com/ory/go-acc v0.2.6
	github.com/pseudomuto/protoc-gen-doc v1.4.1
	golang.org/x/tools v0.1.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.26.0
)
