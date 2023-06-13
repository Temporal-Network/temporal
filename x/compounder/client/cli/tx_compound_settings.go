package cli

import (
	"encoding/json"
	"github.com/Temporal-Network/temporal/x/compounder/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func CmdCreateCompoundSettings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-compound-settings [delegator] [validator-settings] [amount-to-remain] [frequency]",
		Short: "Create a new CompoundSettings",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get value arguments
			argValidatorSettingsRaw := new([]types.ValidatorSetting)
			err = json.Unmarshal([]byte(args[1]), argValidatorSettingsRaw)
			if err != nil {
				return err
			}

			argValidatorSettings := make([]*types.ValidatorSetting, 0, 100)
			for _, vs := range *argValidatorSettingsRaw {
				argValidatorSettings = append(argValidatorSettings, &vs)
			}

			argAmountToRemain, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}
			argFrequency := new(types.Frequency)
			err = json.Unmarshal([]byte(args[3]), argFrequency)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCompoundSettings(
				clientCtx.GetFromAddress().String(),
				argValidatorSettings,
				argAmountToRemain,
				argFrequency,
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

func CmdUpdateCompoundSettings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-compound-settings [delegator] [validator-settings] [amount-to-remain] [frequency]",
		Short: "Update a CompoundSettings",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get value arguments
			argValidatorSettingsRaw := new([]types.ValidatorSetting)
			err = json.Unmarshal([]byte(args[1]), argValidatorSettingsRaw)
			if err != nil {
				return err
			}

			argValidatorSettings := make([]*types.ValidatorSetting, 0, 100)
			for _, vs := range *argValidatorSettingsRaw {
				argValidatorSettings = append(argValidatorSettings, &vs)
			}

			argAmountToRemain, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}
			argFrequency := new(types.Frequency)
			err = json.Unmarshal([]byte(args[3]), argFrequency)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateCompoundSettings(
				clientCtx.GetFromAddress().String(),
				argValidatorSettings,
				argAmountToRemain,
				argFrequency,
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

func CmdDeleteCompoundSettings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-compound-settings [index-123]",
		Short: "Delete a CompoundSettings",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteCompoundSettings(
				clientCtx.GetFromAddress().String(),
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
