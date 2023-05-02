package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"temporal/x/compounder/types"
)

func CmdListCompoundSettings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-compound-settings",
		Short: "list all CompoundSettings",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllCompoundSettingsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.CompoundSettingsAll(context.Background(), params)
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

func CmdShowCompoundSettings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-compound-settings [delegator]",
		Short: "shows a CompoundSettings",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argDelegator := args[0]

			params := &types.QueryGetCompoundSettingsRequest{
				Delegator: argDelegator,
			}

			res, err := queryClient.CompoundSettings(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
