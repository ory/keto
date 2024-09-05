module github.com/ory/keto/proto

go 1.21

toolchain go1.22.5

require (
	github.com/envoyproxy/protoc-gen-validate v1.0.4
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.0
	github.com/stretchr/testify v1.8.4
	google.golang.org/genproto v0.0.0-20230822172742-b8732ec3820d
	google.golang.org/grpc v1.66.0
	google.golang.org/protobuf v1.34.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240604185151-ef581f913117 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract v0.9.0-alpha.0.pre.0
