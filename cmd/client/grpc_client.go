package client

import (
	"context"
	"crypto/tls"
	"fmt"
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
	FlagReadRemote  = "read-remote"
	FlagWriteRemote = "write-remote"
	FlagOplRemote   = "syntax-remote"

	FlagInsecureNoTransportSecurity  = "insecure-disable-transport-security"
	FlagInsecureSkipHostVerification = "insecure-skip-hostname-verification"

	EnvReadRemote  = "KETO_READ_REMOTE"
	EnvWriteRemote = "KETO_WRITE_REMOTE"
	EnvAuthToken   = "KETO_BEARER_TOKEN" // nosec G101 -- just the key, not the value

	ContextKeyTimeout contextKeys = "timeout"
)

type securityFlags struct {
	skipHostVerification bool
	noTransportSecurity  bool
}

func (sf *securityFlags) transportCredentials() grpc.DialOption {
	switch {
	case sf.noTransportSecurity:
		return grpc.WithTransportCredentials(insecure.NewCredentials())

	case sf.skipHostVerification:
		return grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			// nolint explicity set through scary flag
			InsecureSkipVerify: true,
		}))

	default:
		// Defaults to the default host root CA bundle
		return grpc.WithTransportCredentials(credentials.NewTLS(nil))
	}
}

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

func getSecurityFlags(cmd *cobra.Command) securityFlags {
	return securityFlags{
		skipHostVerification: flagx.MustGetBool(cmd, FlagInsecureSkipHostVerification),
		noTransportSecurity:  flagx.MustGetBool(cmd, FlagInsecureNoTransportSecurity),
	}
}

func GetReadConn(cmd *cobra.Command) (*grpc.ClientConn, error) {
	return Conn(cmd.Context(),
		getRemote(cmd, FlagReadRemote, EnvReadRemote),
		getToken(cmd),
		getSecurityFlags(cmd),
	)
}

func GetWriteConn(cmd *cobra.Command) (*grpc.ClientConn, error) {
	return Conn(cmd.Context(),
		getRemote(cmd, FlagWriteRemote, EnvWriteRemote),
		getToken(cmd),
		getSecurityFlags(cmd),
	)
}

func Conn(ctx context.Context, remote, token string, security securityFlags) (*grpc.ClientConn, error) {
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
	dialOpts = append(dialOpts, security.transportCredentials())
	if token != "" {
		dialOpts = append(dialOpts,
			grpc.WithPerRPCCredentials(
				oauth.NewOauthAccess(&oauth2.Token{AccessToken: token})))
	}

	return grpc.DialContext(ctx, remote, dialOpts...)
}

func RegisterRemoteURLFlags(flags *pflag.FlagSet) {
	flags.String(FlagReadRemote, "127.0.0.1:4466", "Remote address of the read API endpoint.")
	flags.String(FlagWriteRemote, "127.0.0.1:4467", "Remote address of the write API endpoint.")
	flags.Bool(FlagInsecureNoTransportSecurity, false, "Disables transport security. Do not use this in production.")
	flags.Bool(FlagInsecureSkipHostVerification, false, "Disables hostname verification. Do not use this in production.")
}
