module github.com/ory/keto

replace google.golang.org/protobuf v1.25.1-0.20201020201750-d3470999428b => google.golang.org/protobuf v1.25.0

replace github.com/soheilhy/cmux => github.com/soheilhy/cmux v0.1.5-0.20210114230657-cdd3331e3e7c

replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2

replace github.com/oleiade/reflections => github.com/oleiade/reflections v1.0.1

replace github.com/gobuffalo/packr => github.com/gobuffalo/packr v1.30.1

replace github.com/ory/keto/proto => ./proto

replace github.com/gobuffalo/pop/v5 => github.com/gobuffalo/pop/v5 v5.3.4-0.20210716080118-0ecad2589ac4

replace github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt/v4 v4.0.0

require (
	github.com/cenkalti/backoff/v3 v3.0.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/errors v0.20.0
	github.com/go-openapi/runtime v0.19.26
	github.com/go-openapi/strfmt v0.20.0
	github.com/go-openapi/swag v0.19.14
	github.com/go-openapi/validate v0.20.2
	github.com/gobuffalo/pop/v5 v5.3.4
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/julienschmidt/httprouter v1.3.0
	github.com/luna-duclos/instrumentedsql v1.1.3
	github.com/luna-duclos/instrumentedsql/opentracing v0.0.0-20201103091713-40d03108b6f4
	github.com/ory/analytics-go/v4 v4.0.0
	github.com/ory/graceful v0.1.1
	github.com/ory/herodot v0.9.6
	github.com/ory/jsonschema/v3 v3.0.3
	github.com/ory/keto/proto v0.0.0-00010101000000-000000000000
	github.com/ory/x v0.0.280
	github.com/pelletier/go-toml v1.8.1
	github.com/phayes/freeport v0.0.0-20180830031419-95f893ade6f2
	github.com/pkg/errors v0.9.1
	github.com/rs/cors v1.6.0
	github.com/segmentio/objconv v1.0.1
	github.com/sirupsen/logrus v1.8.1
	github.com/soheilhy/cmux v0.1.4
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/tidwall/gjson v1.7.1
	github.com/tidwall/sjson v1.1.5
	github.com/urfave/negroni v1.0.0
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
)

go 1.16
