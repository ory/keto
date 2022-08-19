package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/ory/x/flagx"
	"golang.org/x/oauth2"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials/oauth"
)

type contextKeys string

const (
	FlagReadRemote        = "read-remote"
	FlagWriteRemote       = "write-remote"
	FlagInsecureTransport = "insecure"

	EnvReadRemote  = "KETO_READ_REMOTE"
	EnvWriteRemote = "KETO_WRITE_REMOTE"
	EnvAuthToken   = "ORY_PAT"

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

func getToken(cmd *cobra.Command) (token string) {
	return os.Getenv(EnvAuthToken)
}

func GetReadConn(cmd *cobra.Command) (*grpc.ClientConn, error) {
	return Conn(cmd.Context(),
		getRemote(cmd, FlagReadRemote, EnvReadRemote),
		getToken(cmd),
		flagx.MustGetBool(cmd, FlagInsecureTransport),
	)
}

func GetWriteConn(cmd *cobra.Command) (*grpc.ClientConn, error) {
	return Conn(cmd.Context(),
		getRemote(cmd, FlagWriteRemote, EnvWriteRemote),
		getToken(cmd),
		flagx.MustGetBool(cmd, FlagInsecureTransport),
	)
}

func Conn(ctx context.Context, remote, token string, insecureTransport bool) (*grpc.ClientConn, error) {
	timeout := 3 * time.Second
	if d, ok := ctx.Value(ContextKeyTimeout).(time.Duration); ok {
		timeout = d
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	dialOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithDisableHealthCheck(),
	}
	dialOpts = append(dialOpts, transportCredentials(remote, insecureTransport))
	if token != "" {
		dialOpts = append(dialOpts,
			grpc.WithPerRPCCredentials(
				oauth.NewOauthAccess(&oauth2.Token{AccessToken: token})))
	}

	return grpc.DialContext(ctx, remote, dialOpts...)
}

func transportCredentials(remote string, insecureTransport bool) grpc.DialOption {
	if insecureTransport {
		return grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	host, _, err := net.SplitHostPort(remote)
	if err == nil && (host == "127.0.0.1" || host == "localhost") {
		return grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			// nolint only set for local domain.
			InsecureSkipVerify: true,
		}))
	}

	// Defaults to the default host root CA bundle
	return grpc.WithTransportCredentials(credentials.NewTLS(nil))
}

func RegisterRemoteURLFlags(flags *pflag.FlagSet) {
	flags.String(FlagReadRemote, "127.0.0.1:4466", "Remote address of the read API endpoint.")
	flags.String(FlagWriteRemote, "127.0.0.1:4467", "Remote address of the write API endpoint.")
	flags.Bool(FlagInsecureTransport, false, "Do not use transport encryption.")
}
