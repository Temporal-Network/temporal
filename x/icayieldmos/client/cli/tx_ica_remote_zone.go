package cli

import (
    "strconv"
	
	 "github.com/spf13/cast"
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
)

func CmdCreateICARemoteZone() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-ica-remote-zone [connection-id] [remote-chain-id] [active] [bech-32-prefix]",
		Short: "Create a new ICARemoteZone",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
	  	 argConnectionId := args[0]
		 argRemoteChainId := args[1]
		 argActive, err := cast.ToBoolE(args[2])
            		if err != nil {
                		return err
            		}
		 argBech32Prefix := args[3]
		
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateICARemoteZone(clientCtx.GetFromAddress().String(), argConnectionId, argRemoteChainId, argActive, argBech32Prefix)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdUpdateICARemoteZone() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-ica-remote-zone [id] [connection-id] [remote-chain-id] [active] [bech-32-prefix]",
		Short: "Update a ICARemoteZone",
		Args:  cobra.ExactArgs(5),
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
        
	  		argBech32Prefix := args[4]
        
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateICARemoteZone(clientCtx.GetFromAddress().String(), id, argConnectionId, argRemoteChainId, argActive, argBech32Prefix)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdDeleteICARemoteZone() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-ica-remote-zone [id]",
		Short: "Delete a ICARemoteZone by id",
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

			msg := types.NewMsgDeleteICARemoteZone(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}
