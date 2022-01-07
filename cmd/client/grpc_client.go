package client

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ory/x/flagx"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
)

type contextKeys string

const (
	FlagReadRemote  = "read-remote"
	FlagWriteRemote = "write-remote"

	EnvReadRemote  = "KETO_READ_REMOTE"
	EnvWriteRemote = "KETO_WRITE_REMOTE"

	ContextKeyTimeout contextKeys = "timeout"
)

func getRemote(cmd *cobra.Command, flagRemote, envRemote string) (remote string) {
	defer (func() {
		if strings.HasPrefix(remote, "http://") || strings.HasPrefix(remote, "https://") {
			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "remote \"%s\" seems to be an http URL instead of a remote address\n", remote)
		}
	})()

	if cmd.Flags().Changed(flagRemote) {
		return flagx.MustGetString(cmd, flagRemote)
	} else if remote, isSet := os.LookupEnv(envRemote); isSet {
		return remote
	}

	// no value is set, use fallback from the flag and warn about that
	remote = flagx.MustGetString(cmd, flagRemote)
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
	timeout := 3 * time.Second
	if d, ok := ctx.Value(ContextKeyTimeout).(time.Duration); ok {
		timeout = d
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return grpc.DialContext(ctx, remote, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDisableHealthCheck())
}

func RegisterRemoteURLFlags(flags *pflag.FlagSet) {
	flags.String(FlagReadRemote, "127.0.0.1:4466", "Remote address of the read API endpoint.")
	flags.String(FlagWriteRemote, "127.0.0.1:4467", "Remote address of the write API endpoint.")
}
