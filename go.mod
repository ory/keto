module github.com/ory/keto

replace google.golang.org/protobuf v1.25.1-0.20201020201750-d3470999428b => google.golang.org/protobuf v1.25.0

replace github.com/soheilhy/cmux => github.com/soheilhy/cmux v0.1.5-0.20210114230657-cdd3331e3e7c

replace github.com/ory/dockertest/v3 => github.com/ory/dockertest/v3 v3.6.3

require (
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/bufbuild/buf v0.31.1
	github.com/cenkalti/backoff/v3 v3.0.0
	github.com/containerd/continuity v0.0.0-20200228182428-0f16d7a0959c // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/go-bindata/go-bindata v3.1.1+incompatible
	github.com/go-openapi/errors v0.20.0
	github.com/go-openapi/runtime v0.19.26
	github.com/go-openapi/strfmt v0.20.0
	github.com/go-openapi/swag v0.19.13
	github.com/go-openapi/validate v0.20.1
	github.com/go-swagger/go-swagger v0.26.1
	github.com/gobuffalo/pop/v5 v5.3.3
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/sessions v1.1.3
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/julienschmidt/httprouter v1.2.0
	github.com/mattn/goveralls v0.0.8
	github.com/ory/cli v0.0.11
	github.com/ory/go-acc v0.2.6
	github.com/ory/graceful v0.1.1
	github.com/ory/herodot v0.9.2
	github.com/ory/jsonschema/v3 v3.0.1
	github.com/ory/x v0.0.199
	github.com/pelletier/go-toml v1.8.1
	github.com/phayes/freeport v0.0.0-20180830031419-95f893ade6f2
	github.com/pkg/errors v0.9.1
	github.com/rs/cors v1.6.0
	github.com/segmentio/objconv v1.0.1
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/soheilhy/cmux v0.1.4
	github.com/spf13/cobra v1.0.1-0.20201006035406-b97b5ead31f7
	github.com/spf13/pflag v1.0.5
	github.com/sqs/goreturns v0.0.0-20181028201513-538ac6014518
	github.com/stretchr/testify v1.6.1
	github.com/tidwall/gjson v1.6.0
	github.com/tidwall/sjson v1.1.1
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	github.com/urfave/negroni v1.0.0
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
	golang.org/x/tools v0.0.0-20201125231158-b5590deeca9b
	google.golang.org/genproto v0.0.0-20201117123952-62d171c70ae1
	google.golang.org/grpc v1.36.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.25.1-0.20201020201750-d3470999428b
	gotest.tools v2.2.0+incompatible // indirect
)

go 1.16
