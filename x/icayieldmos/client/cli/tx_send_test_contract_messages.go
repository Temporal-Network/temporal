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

func CmdSendTestContractMessages() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-test-contract-messages [connection-id] [remoteDelegatorAddress] [remoteValidatorAddress] [remoteGranteeAddress]",
		Short: "Broadcast message SendTestContractMessages",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argConnectionId := args[0]
			argRemoteDelegatorAddress := args[1]
			argRemoteValidatorAddress := args[2]
			argRemoteGranteeAddress := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSendTestContractMessages(
				clientCtx.GetFromAddress().String(),
				argConnectionId,
				argRemoteDelegatorAddress,
				argRemoteValidatorAddress,
				argRemoteGranteeAddress,
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
