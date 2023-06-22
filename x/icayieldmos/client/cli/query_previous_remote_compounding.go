package cli

import (
	"context"
	"strconv"

	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListPreviousRemoteCompounding() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-previous-remote-compounding",
		Short: "list all previousRemoteCompounding",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllPreviousRemoteCompoundingRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.PreviousRemoteCompoundingAll(context.Background(), params)
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

func CmdShowPreviousRemoteCompounding() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-previous-remote-compounding [remote-contract-compound-setting]",
		Short: "shows a previousRemoteCompounding",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argRemoteContractCompoundSetting, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetPreviousRemoteCompoundingRequest{
				RemoteContractCompoundSetting: argRemoteContractCompoundSetting,
			}

			res, err := queryClient.PreviousRemoteCompounding(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
