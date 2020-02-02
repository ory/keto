module github.com/ory/keto

require (
	github.com/akutz/goof v0.1.2 // indirect
	github.com/akutz/gotil v0.1.0
	github.com/go-errors/errors v1.0.1
	github.com/go-openapi/errors v0.19.2
	github.com/go-openapi/runtime v0.19.5
	github.com/go-openapi/strfmt v0.19.3
	github.com/go-openapi/swag v0.19.5
	github.com/go-openapi/validate v0.19.3
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-swagger/go-swagger v0.21.1-0.20200107003254-1c98855b472d
	github.com/gobuffalo/packr v1.24.1
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gorilla/mux v1.7.1 // indirect
	github.com/gorilla/sessions v1.1.3
	github.com/jmoiron/sqlx v1.2.0
	github.com/julienschmidt/httprouter v1.2.0
	github.com/kardianos/osext v0.0.0-20170510131534-ae77be60afb1 // indirect
	github.com/lib/pq v1.0.0
	github.com/open-policy-agent/opa v0.10.1
	github.com/ory/go-acc v0.0.0-20181118080137-ddc355013f90
	github.com/ory/graceful v0.1.1
	github.com/ory/herodot v0.6.2
	github.com/ory/sdk/swagutil v0.0.0-20200202121523-307941feee4b
	github.com/ory/viper v1.5.6
	github.com/ory/x v0.0.93
	github.com/pborman/uuid v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/pkg/profile v1.3.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a // indirect
	github.com/rs/cors v1.6.0
	github.com/rubenv/sql-migrate v0.0.0-20190327083759-54bad0a9b051
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/sqs/goreturns v0.0.0-20181028201513-538ac6014518
	github.com/stretchr/testify v1.4.0
	github.com/urfave/negroni v1.0.0
	github.com/yashtewari/glob-intersection v0.0.0-20180916065949-5c77d914dd0b // indirect
	golang.org/x/tools v0.0.0-20191224055732-dd894d0a8a40
)

// Fix for https://github.com/golang/lint/issues/436
replace github.com/golang/lint => github.com/golang/lint v0.0.0-20181217174547-8f45f776aaf1

go 1.13
