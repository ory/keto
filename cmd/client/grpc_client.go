package client

import (
	"context"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
)

const (
	FlagRemoteURL = "remote"
)

func GetGRPCConn(cmd *cobra.Command) (*grpc.ClientConn, error) {
	remote, err := cmd.Flags().GetString(FlagRemoteURL)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return grpc.DialContext(ctx, remote, grpc.WithInsecure(), grpc.WithBlock())
}

func RegisterRemoteURLFlag(flags *pflag.FlagSet) {
	flags.StringP(FlagRemoteURL, "r", os.Getenv("KETO_GRPC_URL"), "TODO")
}
