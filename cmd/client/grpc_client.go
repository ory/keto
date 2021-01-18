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
	FlagBasicRemote      = "basic-remote"
	FlagPrivilegedRemote = "privileged-remote"

	EnvBasicRemote      = "KETO_BASIC_REMOTE"
	EnvPrivilegedRemote = "KETO_PRIVILEGED_REMOTE"
)

func getBasicRemote(cmd *cobra.Command) string {
	remote := flagx.MustGetString(cmd, FlagBasicRemote)
	if remote == "" {
		remote = os.Getenv(EnvBasicRemote)
	}
	if remote == "" {
		remote = getPrivilegedRemote(cmd)
	}
	return remote
}

func getPrivilegedRemote(cmd *cobra.Command) string {
	remote := flagx.MustGetString(cmd, FlagPrivilegedRemote)
	if remote == "" {
		remote = os.Getenv(EnvPrivilegedRemote)
	}
	return remote
}

func GetBasicConn(cmd *cobra.Command) (*grpc.ClientConn, error) {
	remote := getBasicRemote(cmd)
	ctx, cancel := context.WithTimeout(cmd.Context(), 3*time.Second)
	defer cancel()
	return grpc.DialContext(ctx, remote, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDisableHealthCheck())
}

func GetPrivilegedConn(cmd *cobra.Command) (*grpc.ClientConn, error) {
	remote := getPrivilegedRemote(cmd)
	ctx, cancel := context.WithTimeout(cmd.Context(), 3*time.Second)
	defer cancel()
	return grpc.DialContext(ctx, remote, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDisableHealthCheck())
}

func RegisterRemoteURLFlags(flags *pflag.FlagSet) {
	flags.String(FlagBasicRemote, "127.0.0.1:4466", "Remote URL of the privileged API endpoint.")
	flags.String(FlagPrivilegedRemote, "127.0.0.1:4467", "Remote URL of the basic API endpoint.")
}
