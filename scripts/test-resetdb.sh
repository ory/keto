#!/usr/bin/env bash

set -euxo pipefail

docker rm -f test_keto_postgres || true
docker rm -f test_keto_mysql || true
docker rm -f test_keto_cockroach || true

postgres_port="$(docker port "$(docker run --name test_keto_postgres -e "POSTGRES_PASSWORD=secret" -e "POSTGRES_DB=postgres" -p 0.0.0.0:0:5432 -d postgres:11.8)" 5432 | sed 's/.*:\([0-9]*\)/\1/')"
mysql_port="$(docker port "$(docker run --name test_keto_mysql -e "MYSQL_ROOT_PASSWORD=secret" -p 0.0.0.0:0:3306 -d mysql:8.0)" 3306 | sed 's/.*:\([0-9]*\)/\1/')"
cockroach_port="$(docker port "$(docker run --name test_keto_cockroach -p 0.0.0.0:0:26257 -d cockroachdb/cockroach:v20.2.4 start-single-node --insecure)" 26257 | sed 's/.*:\([0-9]*\)/\1/')"

TEST_DATABASE_POSTGRESQL=$(printf "postgres://postgres:secret@localhost:%s/postgres?sslmode=disable" "$postgres_port")
TEST_DATABASE_MYSQL=$(printf "mysql://root:secret@(localhost:%s)/mysql?parseTime=true&multiStatements=true" "$mysql_port")
TEST_DATABASE_COCKROACHDB=$(printf "cockroach://root@localhost:%s/defaultdb?sslmode=disable" "$cockroach_port")

export TEST_DATABASE_POSTGRESQL
export TEST_DATABASE_MYSQL
export TEST_DATABASE_COCKROACHDB

# undo set from above
set +e
set +u
set +x
set +o pipefail
