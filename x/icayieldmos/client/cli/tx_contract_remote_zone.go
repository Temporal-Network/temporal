package cli

import (
    "strconv"
	
	 "github.com/spf13/cast"
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
)

func CmdCreateContractRemoteZone() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-contract-remote-zone [connection-id] [remote-chain-id] [active]",
		Short: "Create a new ContractRemoteZone",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
	  	 argConnectionId := args[0]
		 argRemoteChainId := args[1]
		 argActive, err := cast.ToBoolE(args[2])
            		if err != nil {
                		return err
            		}
		
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateContractRemoteZone(clientCtx.GetFromAddress().String(), argConnectionId, argRemoteChainId, argActive)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdUpdateContractRemoteZone() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-contract-remote-zone [id] [connection-id] [remote-chain-id] [active]",
		Short: "Update a ContractRemoteZone",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            id, err := strconv.ParseUint(args[0], 10, 64)
            if err != nil {
                return err
            }

	    
	  		argConnectionId := args[1]
        
	  		argRemoteChainId := args[2]
        
	  		argActive, err := cast.ToBoolE(args[3])
            		if err != nil {
                		return err
            		}
        
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateContractRemoteZone(clientCtx.GetFromAddress().String(), id, argConnectionId, argRemoteChainId, argActive)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdDeleteContractRemoteZone() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-contract-remote-zone [id]",
		Short: "Delete a ContractRemoteZone by id",
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

			msg := types.NewMsgDeleteContractRemoteZone(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}
