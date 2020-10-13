package client

import (
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
	return grpc.Dial(remote, grpc.WithInsecure(), grpc.WithBlock())
}

func RegisterRemoteURLFlag(flags *pflag.FlagSet) {
	flags.StringP(FlagRemoteURL, "r", "", "TODO")
}
