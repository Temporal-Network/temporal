package cli

import (
	"context"
	"strconv"

	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListContractRemoteZone() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-contract-remote-zone",
		Short: "list all ContractRemoteZone",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllContractRemoteZoneRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ContractRemoteZoneAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowContractRemoteZone() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-contract-remote-zone [id]",
		Short: "shows a ContractRemoteZone",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetContractRemoteZoneRequest{
				Id: id,
			}

			res, err := queryClient.ContractRemoteZone(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
