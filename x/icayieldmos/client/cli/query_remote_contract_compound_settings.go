package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
)

func CmdListRemoteContractCompoundSettings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-remote-contract-compound-settings",
		Short: "list all RemoteContractCompoundSettings",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllRemoteContractCompoundSettingsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.RemoteContractCompoundSettingsAll(context.Background(), params)
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

func CmdShowRemoteContractCompoundSettings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-remote-contract-compound-settings [id]",
		Short: "shows a RemoteContractCompoundSettings",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetRemoteContractCompoundSettingsRequest{
				Id: id,
			}

			res, err := queryClient.RemoteContractCompoundSettings(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
