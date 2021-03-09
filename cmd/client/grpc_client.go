package client

import (
	"context"
	"fmt"
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

func getRemote(cmd *cobra.Command, flagRemote, envRemote string) string {
	if cmd.Flags().Changed(flagRemote) {
		return flagx.MustGetString(cmd, flagRemote)
	} else if remote, isSet := os.LookupEnv(envRemote); isSet {
		return remote
	}

	// no value is set, use fallback from the flag and warn about that
	remote := flagx.MustGetString(cmd, flagRemote)
	_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "neither flag --%s nor env var %s are set, falling back to %s\n", flagRemote, envRemote, remote)
	return remote
}

func GetReadConn(cmd *cobra.Command) (*grpc.ClientConn, error) {
	return Conn(cmd.Context(), getRemote(cmd, FlagReadRemote, EnvReadRemote))
}

func GetWriteConn(cmd *cobra.Command) (*grpc.ClientConn, error) {
	return Conn(cmd.Context(), getRemote(cmd, FlagWriteRemote, EnvWriteRemote))
}

func Conn(ctx context.Context, remote string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return grpc.DialContext(ctx, remote, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDisableHealthCheck())
}

func RegisterRemoteURLFlags(flags *pflag.FlagSet) {
	flags.String(FlagReadRemote, "127.0.0.1:4466", "Remote URL of the read API endpoint.")
	flags.String(FlagWriteRemote, "127.0.0.1:4467", "Remote URL of the write API endpoint.")
}
