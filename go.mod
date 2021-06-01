module github.com/ory/keto

replace google.golang.org/protobuf v1.25.1-0.20201020201750-d3470999428b => google.golang.org/protobuf v1.25.0

replace github.com/soheilhy/cmux => github.com/soheilhy/cmux v0.1.5-0.20210114230657-cdd3331e3e7c

replace github.com/ory/dockertest/v3 => github.com/ory/dockertest/v3 v3.6.5

replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2

replace github.com/ory/kratos-client-go => github.com/ory/kratos-client-go v0.5.4-alpha.1.0.20210210170256-960b093d8bf9

replace github.com/ory/kratos/corp => github.com/ory/kratos/corp v0.0.0-20210118092700-c2358be1e867

replace github.com/seatgeek/logrus-gelf-formatter => github.com/zepatrik/logrus-gelf-formatter v0.0.0-20210305135027-b8b3731dba10

replace gopkg.in/DataDog/dd-trace-go.v1 => gopkg.in/DataDog/dd-trace-go.v1 v1.27.1

replace github.com/oleiade/reflections => github.com/oleiade/reflections v1.0.1

replace github.com/goreleaser/nfpm => github.com/goreleaser/nfpm v1.10.2

replace github.com/ory/keto/proto/ory/keto/acl/v1alpha1 => ./proto/ory/keto/acl/v1alpha1

require (
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.7 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/DataDog/datadog-go v4.7.0+incompatible // indirect
	github.com/bufbuild/buf v0.31.1
	github.com/cenkalti/backoff/v3 v3.0.0
	github.com/containerd/containerd v1.5.2 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/errors v0.20.0
	github.com/go-openapi/runtime v0.19.26
	github.com/go-openapi/strfmt v0.20.0
	github.com/go-openapi/swag v0.19.14
	github.com/go-openapi/validate v0.20.2
	github.com/go-swagger/go-swagger v0.26.1
	github.com/gobuffalo/pop/v5 v5.3.4
	github.com/golang/mock v1.5.0
	github.com/goreleaser/godownloader v0.1.1-0.20200426152203-fd8ad8f7dd78
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/julienschmidt/httprouter v1.3.0
	github.com/luna-duclos/instrumentedsql v1.1.3
	github.com/luna-duclos/instrumentedsql/opentracing v0.0.0-20201103091713-40d03108b6f4
	github.com/mattn/goveralls v0.0.8
	github.com/ory/analytics-go/v4 v4.0.0
	github.com/ory/cli v0.0.49
	github.com/ory/go-acc v0.2.6
	github.com/ory/graceful v0.1.1
	github.com/ory/herodot v0.9.3
	github.com/ory/jsonschema/v3 v3.0.3
	github.com/ory/keto/proto/ory/keto/acl/v1alpha1 v0.0.0-00010101000000-000000000000
	github.com/ory/x v0.0.242
	github.com/pelletier/go-toml v1.8.1
	github.com/phayes/freeport v0.0.0-20180830031419-95f893ade6f2
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pseudomuto/protoc-gen-doc v1.4.1
	github.com/rs/cors v1.6.0
	github.com/segmentio/objconv v1.0.1
	github.com/sirupsen/logrus v1.8.1
	github.com/soheilhy/cmux v0.1.4
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/sqs/goreturns v0.0.0-20181028201513-538ac6014518
	github.com/stretchr/testify v1.7.0
	github.com/tidwall/gjson v1.7.1
	github.com/tidwall/sjson v1.1.5
	github.com/urfave/negroni v1.0.0
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	golang.org/x/tools v0.1.0
	google.golang.org/grpc v1.37.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.26.0
)

go 1.16
