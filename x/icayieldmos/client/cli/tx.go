package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateConnection())
	cmd.AddCommand(CmdCreateContractRemoteZone())
	cmd.AddCommand(CmdUpdateContractRemoteZone())
	cmd.AddCommand(CmdDeleteContractRemoteZone())
	cmd.AddCommand(CmdSendTestContractMessages())
	cmd.AddCommand(CmdCreateRemoteContractCompoundSettings())
	cmd.AddCommand(CmdUpdateRemoteContractCompoundSettings())
	cmd.AddCommand(CmdDeleteRemoteContractCompoundSettings())
	cmd.AddCommand(CmdCreateICARemoteZone())
	cmd.AddCommand(CmdUpdateICARemoteZone())
	cmd.AddCommand(CmdDeleteICARemoteZone())
// this line is used by starport scaffolding # 1

	return cmd
}
