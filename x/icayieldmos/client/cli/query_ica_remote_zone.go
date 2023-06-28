package cli

import (
    "context"
    "strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
    "github.com/temporal-zone/temporal/x/icayieldmos/types"
)

func CmdListICARemoteZone() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-ica-remote-zone",
		Short: "list all ICARemoteZone",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllICARemoteZoneRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.ICARemoteZoneAll(context.Background(), params)
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

func CmdShowICARemoteZone() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-ica-remote-zone [id]",
		Short: "shows a ICARemoteZone",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            queryClient := types.NewQueryClient(clientCtx)

            id, err := strconv.ParseUint(args[0], 10, 64)
            if err != nil {
                return err
            }

            params := &types.QueryGetICARemoteZoneRequest{
                Id: id,
            }

            res, err := queryClient.ICARemoteZone(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
