// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

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
	FlagAuthority                    = "authority"

	EnvReadRemote  = "KETO_READ_REMOTE"
	EnvWriteRemote = "KETO_WRITE_REMOTE"
	EnvAuthToken   = "KETO_BEARER_TOKEN" // nosec G101 -- just the key, not the value
	EnvAuthority   = "KETO_AUTHORITY"

	ContextKeyTimeout contextKeys = "timeout"
)

type connectionDetails struct {
	token, authority     string
	skipHostVerification bool
	noTransportSecurity  bool
}

func (d *connectionDetails) dialOptions() (opts []grpc.DialOption) {
	if d.token != "" {
		opts = append(opts,
			grpc.WithPerRPCCredentials(
				oauth.NewOauthAccess(&oauth2.Token{AccessToken: d.token})))
	}
	if d.authority != "" {
		opts = append(opts, grpc.WithAuthority(d.authority))
	}

	// TLS settings
	switch {
	case d.noTransportSecurity:
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	case d.skipHostVerification:
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			// nolint explicity set through scary flag
			InsecureSkipVerify: true,
		})))
	default:
		// Defaults to the default host root CA bundle
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	}
	return opts
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

func getAuthority(cmd *cobra.Command) string {
	if cmd.Flags().Changed(FlagAuthority) {
		return flagx.MustGetString(cmd, FlagAuthority)
	}
	return os.Getenv(EnvAuthority)
}

func getConnectionDetails(cmd *cobra.Command) connectionDetails {
	return connectionDetails{
		token:                os.Getenv(EnvAuthToken),
		authority:            getAuthority(cmd),
		skipHostVerification: flagx.MustGetBool(cmd, FlagInsecureSkipHostVerification),
		noTransportSecurity:  flagx.MustGetBool(cmd, FlagInsecureNoTransportSecurity),
	}
}

func GetReadConn(cmd *cobra.Command) (*grpc.ClientConn, error) {
	return Conn(cmd.Context(),
		getRemote(cmd, FlagReadRemote, EnvReadRemote),
		getConnectionDetails(cmd),
	)
}

func GetWriteConn(cmd *cobra.Command) (*grpc.ClientConn, error) {
	return Conn(cmd.Context(),
		getRemote(cmd, FlagWriteRemote, EnvWriteRemote),
		getConnectionDetails(cmd),
	)
}

func Conn(ctx context.Context, remote string, details connectionDetails) (*grpc.ClientConn, error) {
	timeout := 3 * time.Second
	if d, ok := ctx.Value(ContextKeyTimeout).(time.Duration); ok {
		timeout = d
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return grpc.DialContext(
		ctx,
		remote,
		append([]grpc.DialOption{
			grpc.WithBlock(),
			grpc.WithDisableHealthCheck(),
		}, details.dialOptions()...)...,
	)
}

func RegisterRemoteURLFlags(flags *pflag.FlagSet) {
	flags.String(FlagReadRemote, "127.0.0.1:4466", "Remote address of the read API endpoint.")
	flags.String(FlagWriteRemote, "127.0.0.1:4467", "Remote address of the write API endpoint.")
	flags.String(FlagAuthority, "", "Set the authority header for the remote gRPC server.")
	flags.Bool(FlagInsecureNoTransportSecurity, false, "Disables transport security. Do not use this in production.")
	flags.Bool(FlagInsecureSkipHostVerification, false, "Disables hostname verification. Do not use this in production.")
}
