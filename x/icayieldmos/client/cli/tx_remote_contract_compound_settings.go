package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
)

func CmdCreateRemoteContractCompoundSettings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-remote-contract-compound-settings [contract-remote-zone] [remote-delegator-address] [compound-settings] [frequency-in-seconds] [remote-contract-address]",
		Short: "Create a new RemoteContractCompoundSettings",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argContractRemoteZone, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			argRemoteDelegatorAddress := args[1]
			argCompoundSettings := args[2]
			argFrequencyInSeconds, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return err
			}
			argRemoteContractAddress := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateRemoteContractCompoundSettings(clientCtx.GetFromAddress().String(), argContractRemoteZone, argRemoteDelegatorAddress, argCompoundSettings, argFrequencyInSeconds, argRemoteContractAddress)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateRemoteContractCompoundSettings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-remote-contract-compound-settings [id] [contract-remote-zone] [remote-delegator-address] [compound-settings] [frequency-in-seconds] [remote-contract-address]",
		Short: "Update a RemoteContractCompoundSettings",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argContractRemoteZone, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			argRemoteDelegatorAddress := args[2]

			argCompoundSettings := args[3]

			argFrequencyInSeconds, err := strconv.ParseUint(args[4], 10, 64)
			if err != nil {
				return err
			}

			argRemoteContractAddress := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateRemoteContractCompoundSettings(clientCtx.GetFromAddress().String(), id, argContractRemoteZone, argRemoteDelegatorAddress, argCompoundSettings, argFrequencyInSeconds, argRemoteContractAddress)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteRemoteContractCompoundSettings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-remote-contract-compound-settings [id]",
		Short: "Delete a RemoteContractCompoundSettings by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteRemoteContractCompoundSettings(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
