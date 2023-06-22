package icayieldmos

import (
	"github.com/Temporal-Network/temporal/x/icayieldmos/keeper"
	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the contractRemoteZone
	for _, elem := range genState.ContractRemoteZoneList {
		k.SetContractRemoteZone(ctx, elem)
	}

	// Set contractRemoteZone count
	k.SetContractRemoteZoneCount(ctx, genState.ContractRemoteZoneCount)
	// Set all the remoteContractCompoundSettings
for _, elem := range genState.RemoteContractCompoundSettingsList {
	k.SetRemoteContractCompoundSettings(ctx, elem)
}

// Set remoteContractCompoundSettings count
k.SetRemoteContractCompoundSettingsCount(ctx, genState.RemoteContractCompoundSettingsCount)
// Set all the previousRemoteCompounding
for _, elem := range genState.PreviousRemoteCompoundingList {
	k.SetPreviousRemoteCompounding(ctx, elem)
}
// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	genesis.ContractRemoteZoneList = k.GetAllContractRemoteZone(ctx)
	genesis.ContractRemoteZoneCount = k.GetContractRemoteZoneCount(ctx)
	genesis.RemoteContractCompoundSettingsList = k.GetAllRemoteContractCompoundSettings(ctx)
genesis.RemoteContractCompoundSettingsCount = k.GetRemoteContractCompoundSettingsCount(ctx)
genesis.PreviousRemoteCompoundingList = k.GetAllPreviousRemoteCompounding(ctx)
// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
