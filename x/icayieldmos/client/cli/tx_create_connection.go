package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
)

var _ = strconv.Itoa(0)

func CmdCreateConnection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-connection [connection-id]",
		Short: "Broadcast message CreateConnection",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argConnectionId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateConnection(
				clientCtx.GetFromAddress().String(),
				argConnectionId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
