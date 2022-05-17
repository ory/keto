module github.com/ory/keto

replace google.golang.org/protobuf v1.25.1-0.20201020201750-d3470999428b => google.golang.org/protobuf v1.25.0

replace github.com/soheilhy/cmux => github.com/soheilhy/cmux v0.1.5-0.20210114230657-cdd3331e3e7c

replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2

replace github.com/oleiade/reflections => github.com/oleiade/reflections v1.0.1

replace github.com/gobuffalo/packr => github.com/gobuffalo/packr v1.30.1

replace github.com/ory/keto/proto => ./proto

replace github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt/v4 v4.0.0

replace github.com/containerd/containerd => github.com/containerd/containerd v1.6.1

replace github.com/opencontainers/runc => github.com/opencontainers/runc v1.0.3

require (
	github.com/cenkalti/backoff/v3 v3.0.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/errors v0.20.1
	github.com/go-openapi/runtime v0.20.0
	github.com/go-openapi/strfmt v0.20.3
	github.com/go-openapi/swag v0.19.15
	github.com/go-openapi/validate v0.20.3
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gobuffalo/pop/v6 v6.0.1
	github.com/gofrs/uuid v4.1.0+incompatible
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/luna-duclos/instrumentedsql v1.1.3
	github.com/mikefarah/yq/v4 v4.19.1
	github.com/ory/analytics-go/v4 v4.0.3
	github.com/ory/graceful v0.1.2
	github.com/ory/herodot v0.9.13
	github.com/ory/jsonschema/v3 v3.0.7
	github.com/ory/keto/proto v0.0.0-00010101000000-000000000000
	github.com/ory/x v0.0.389
	github.com/pelletier/go-toml v1.9.4
	github.com/phayes/freeport v0.0.0-20180830031419-95f893ade6f2
	github.com/pkg/errors v0.9.1
	github.com/rs/cors v1.8.0
	github.com/segmentio/objconv v1.0.1
	github.com/sirupsen/logrus v1.8.1
	github.com/soheilhy/cmux v0.1.5
	github.com/spf13/cobra v1.4.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.1
	github.com/tidwall/gjson v1.14.0
	github.com/tidwall/sjson v1.2.4
	github.com/urfave/negroni v1.0.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.30.0
	go.opentelemetry.io/otel v1.7.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
)

require (
	cloud.google.com/go v0.99.0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/a8m/envsubst v1.3.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.1.2 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/cockroachdb/cockroach-go/v2 v2.2.7 // indirect
	github.com/containerd/containerd v1.5.8 // indirect
	github.com/containerd/continuity v0.2.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/docker/cli v20.10.11+incompatible // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/docker v20.10.9+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/elliotchance/orderedmap v1.4.0 // indirect
	github.com/evanphx/json-patch v4.11.0+incompatible // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/analysis v0.20.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.5 // indirect
	github.com/go-openapi/loads v0.20.2 // indirect
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/gobuffalo/envy v1.10.1 // indirect
	github.com/gobuffalo/fizz v1.14.0 // indirect
	github.com/gobuffalo/flect v0.2.4 // indirect
	github.com/gobuffalo/github_flavored_markdown v1.1.1 // indirect
	github.com/gobuffalo/helpers v0.6.4 // indirect
	github.com/gobuffalo/nulls v0.4.1 // indirect
	github.com/gobuffalo/plush/v4 v4.1.9 // indirect
	github.com/gobuffalo/tags/v3 v3.1.2 // indirect
	github.com/gobuffalo/validate/v3 v3.3.1 // indirect
	github.com/goccy/go-yaml v1.9.5 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.0.0 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/inhies/go-bytesize v0.0.0-20210819104631-275770b98743 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.10.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.2.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.9.0 // indirect
	github.com/jackc/pgx/v4 v4.14.0 // indirect
	github.com/jandelgado/gcov2lcov v1.0.5 // indirect
	github.com/jinzhu/copier v0.3.5 // indirect
	github.com/jmoiron/sqlx v1.3.4 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/knadh/koanf v1.4.0 // indirect
	github.com/lib/pq v1.10.4 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/microcosm-cc/bluemonday v1.0.16 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/nyaruka/phonenumbers v1.0.73 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/opencontainers/runc v1.1.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/openzipkin/zipkin-go v0.4.0 // indirect
	github.com/ory/dockertest/v3 v3.8.1 // indirect
	github.com/ory/go-acc v0.2.6 // indirect
	github.com/ory/viper v1.7.5 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pkg/profile v1.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/seatgeek/logrus-gelf-formatter v0.0.0-20210414080842-5b05eb8ff761 // indirect
	github.com/segmentio/backo-go v0.0.0-20200129164019-23eae7c10bd3 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/sourcegraph/annotate v0.0.0-20160123013949-f4cad6c6324d // indirect
	github.com/sourcegraph/syntaxhighlight v0.0.0-20170531221838-bd320f5d308e // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/timtadh/data-structures v0.5.3 // indirect
	github.com/timtadh/lexmachine v0.2.2 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/xtgo/uuid v0.0.0-20140804021211-a0b114877d4c // indirect
	go.mongodb.org/mongo-driver v1.7.3 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.25.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.29.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.4.0 // indirect
	go.opentelemetry.io/contrib/propagators/jaeger v1.4.0 // indirect
	go.opentelemetry.io/contrib/samplers/jaegerremote v0.0.0-20220314184135-32895002a444 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.5.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.6.3 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.6.3 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.6.3 // indirect
	go.opentelemetry.io/otel/exporters/zipkin v1.7.0 // indirect
	go.opentelemetry.io/otel/internal/metric v0.27.0 // indirect
	go.opentelemetry.io/otel/metric v0.27.0 // indirect
	go.opentelemetry.io/otel/sdk v1.7.0 // indirect
	go.opentelemetry.io/otel/trace v1.7.0 // indirect
	go.opentelemetry.io/proto/otlp v0.15.0 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/op/go-logging.v1 v1.0.0-20160211212156-b2cb9fa56473 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

go 1.17
