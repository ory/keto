package dbx

import (
	"testing"

	"github.com/gobuffalo/pop/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_withDbName(t *testing.T) {
	type args struct {
		dsn string
		db  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "postgres",
		args: args{
			dsn: "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
			db:  "mydb",
		},
		want: "postgres://postgres:postgres@localhost:5432/mydb?sslmode=disable",
	}, {
		name: "cockroach",
		args: args{
			dsn: "cockroach://root@localhost:49364/defaultdb?sslmode=disable",
			db:  "foo",
		},
		want: "cockroach://root@localhost:49364/foo?sslmode=disable",
	}, {
		name: "mysql",
		args: args{
			dsn: "mysql://root:secret@(localhost:49394)/mysql?parseTime=true&multiStatements=true",
			db:  "testdb",
		},
		want: "mysql://root:secret@tcp(localhost:49394)/testdb?multiStatements=true&parseTime=true",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := withDbName(tt.args.dsn, tt.args.db); got != tt.want {
				t.Errorf("\nwant %q\ngot  %q", tt.want, got)
			}
		})
	}
}

func Test_GetDSNs_can_connect_to_each_db(t *testing.T) {
	for _, db := range GetDSNs(t, false) {
		db := db
		t.Run("dsn="+db.Name, func(t *testing.T) {
			t.Parallel()
			conn, err := pop.NewConnection(&pop.ConnectionDetails{URL: db.Conn})
			require.NoError(t, err)
			assert.NoError(t, conn.Open())
			assert.NoError(t, Ping(conn))
			assert.NoError(t, conn.Close())
		})
	}
}
