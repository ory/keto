module github.com/ory/keto

replace google.golang.org/protobuf v1.25.1-0.20201020201750-d3470999428b => google.golang.org/protobuf v1.25.0

replace github.com/soheilhy/cmux => github.com/soheilhy/cmux v0.1.5-0.20210114230657-cdd3331e3e7c

replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2

replace github.com/oleiade/reflections => github.com/oleiade/reflections v1.0.1

replace github.com/gobuffalo/packr => github.com/gobuffalo/packr v1.30.1

replace github.com/ory/keto/proto => ./proto

replace github.com/gobuffalo/pop/v5 => github.com/gobuffalo/pop/v5 v5.3.4-0.20210716080118-0ecad2589ac4

replace github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt/v4 v4.0.0

replace github.com/ory/x => github.com/ory/x v0.0.289-0.20211001134443-5386a420dd32

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
	github.com/ory/x v0.0.288
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

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/DataDog/datadog-go v4.0.0+incompatible // indirect
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/cenkalti/backoff/v4 v4.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/cockroachdb/cockroach-go/v2 v2.1.1 // indirect
	github.com/containerd/containerd v1.4.3 // indirect
	github.com/containerd/continuity v0.0.0-20200107194136-26c1120b8d41 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v17.12.0-ce-rc1.0.20201201034508-7d75c1d40d88+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/elastic/go-licenser v0.3.1 // indirect
	github.com/elastic/go-sysinfo v1.1.1 // indirect
	github.com/elastic/go-windows v1.0.0 // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-openapi/analysis v0.20.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.5 // indirect
	github.com/go-openapi/loads v0.20.2 // indirect
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gobuffalo/envy v1.9.0 // indirect
	github.com/gobuffalo/fizz v1.13.1-0.20201104174146-3416f0e6618f // indirect
	github.com/gobuffalo/flect v0.2.1 // indirect
	github.com/gobuffalo/github_flavored_markdown v1.1.0 // indirect
	github.com/gobuffalo/helpers v0.6.1 // indirect
	github.com/gobuffalo/here v0.6.0 // indirect
	github.com/gobuffalo/nulls v0.3.0 // indirect
	github.com/gobuffalo/packd v1.0.0 // indirect
	github.com/gobuffalo/plush/v4 v4.0.0 // indirect
	github.com/gobuffalo/tags/v3 v3.1.0 // indirect
	github.com/gobuffalo/validate/v3 v3.2.0 // indirect
	github.com/gofrs/uuid/v3 v3.1.2 // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.6.8 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/inhies/go-bytesize v0.0.0-20201103132853-d0aed0d254f8 // indirect
	github.com/instana/go-sensor v1.29.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.8.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.0.6 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.7.0 // indirect
	github.com/jackc/pgx/v4 v4.11.0 // indirect
	github.com/jandelgado/gcov2lcov v1.0.4 // indirect
	github.com/jcchavezs/porto v0.1.0 // indirect
	github.com/jmoiron/sqlx v1.3.3 // indirect
	github.com/joeshaw/multierror v0.0.0-20140124173710-69b34d4ec901 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/knadh/koanf v1.2.2 // indirect
	github.com/lib/pq v1.10.1 // indirect
	github.com/looplab/fsm v0.1.0 // indirect
	github.com/magiconair/properties v1.8.4 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/markbates/pkger v0.17.1 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/microcosm-cc/bluemonday v1.0.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/term v0.0.0-20201216013528-df9cb8a40635 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opencontainers/runc v1.0.0-rc9 // indirect
	github.com/opentracing-contrib/go-observer v0.0.0-20170622124052-a52f23424492 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5 // indirect
	github.com/openzipkin/zipkin-go v0.2.2 // indirect
	github.com/ory/dockertest/v3 v3.6.5 // indirect
	github.com/ory/go-acc v0.2.6 // indirect
	github.com/ory/viper v1.7.5 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/pkg/profile v1.2.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/procfs v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.5.2 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/santhosh-tekuri/jsonschema v1.2.4 // indirect
	github.com/seatgeek/logrus-gelf-formatter v0.0.0-20210414080842-5b05eb8ff761 // indirect
	github.com/segmentio/backo-go v0.0.0-20200129164019-23eae7c10bd3 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sourcegraph/annotate v0.0.0-20160123013949-f4cad6c6324d // indirect
	github.com/sourcegraph/syntaxhighlight v0.0.0-20170531221838-bd320f5d308e // indirect
	github.com/spf13/afero v1.5.1 // indirect
	github.com/spf13/cast v1.3.2-0.20200723214538-8d17101741c8 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/sqs/goreturns v0.0.0-20181028201513-538ac6014518 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/tidwall/match v1.0.3 // indirect
	github.com/tidwall/pretty v1.1.0 // indirect
	github.com/tinylib/msgp v1.1.2 // indirect
	github.com/uber/jaeger-client-go v2.22.1+incompatible // indirect
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	github.com/xtgo/uuid v0.0.0-20140804021211-a0b114877d4c // indirect
	go.elastic.co/apm v1.13.0 // indirect
	go.elastic.co/apm/module/apmhttp v1.13.0 // indirect
	go.elastic.co/apm/module/apmot v1.13.0 // indirect
	go.elastic.co/fastjson v1.1.0 // indirect
	go.mongodb.org/mongo-driver v1.4.6 // indirect
	go.opentelemetry.io/contrib v0.20.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.20.0 // indirect
	go.opentelemetry.io/otel v0.20.0 // indirect
	go.opentelemetry.io/otel/metric v0.20.0 // indirect
	go.opentelemetry.io/otel/trace v0.20.0 // indirect
	go.uber.org/atomic v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/net v0.0.0-20210316092652-d523dce5a7f4 // indirect
	golang.org/x/sys v0.0.0-20210320140829-1e4c9ba3b0c4 // indirect
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/tools v0.1.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20210503173045-b96a97608f20 // indirect
	gopkg.in/DataDog/dd-trace-go.v1 v1.27.0 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gotest.tools v2.2.0+incompatible // indirect
	howett.net/plist v0.0.0-20181124034731-591f970eefbb // indirect
)

go 1.17
