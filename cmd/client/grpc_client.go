package client

import (
	"context"
	"os"
	"time"

	"github.com/ory/x/flagx"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
)

const (
	FlagReadRemote  = "read-remote"
	FlagWriteRemote = "write-remote"

	EnvReadRemote  = "KETO_READ_REMOTE"
	EnvWriteRemote = "KETO_WRITE_REMOTE"
)

func getReadRemote(cmd *cobra.Command) string {
	remote := flagx.MustGetString(cmd, FlagReadRemote)
	if remote == "" {
		remote = os.Getenv(EnvReadRemote)
	}
	return remote
}

func getWriteRemote(cmd *cobra.Command) string {
	remote := flagx.MustGetString(cmd, FlagWriteRemote)
	if remote == "" {
		remote = os.Getenv(EnvWriteRemote)
	}
	return remote
}

func GetReadConn(cmd *cobra.Command) (*grpc.ClientConn, error) {
	remote := getReadRemote(cmd)
	ctx, cancel := context.WithTimeout(cmd.Context(), 3*time.Second)
	defer cancel()
	return grpc.DialContext(ctx, remote, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDisableHealthCheck())
}

func GetWriteConn(cmd *cobra.Command) (*grpc.ClientConn, error) {
	remote := getWriteRemote(cmd)
	ctx, cancel := context.WithTimeout(cmd.Context(), 3*time.Second)
	defer cancel()
	return grpc.DialContext(ctx, remote, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDisableHealthCheck())
}

func RegisterRemoteURLFlags(flags *pflag.FlagSet) {
	flags.String(FlagReadRemote, "127.0.0.1:4466", "Remote URL of the read API endpoint.")
	flags.String(FlagWriteRemote, "127.0.0.1:4467", "Remote URL of the write API endpoint.")
}
