package compounder

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/compounder/keeper"
	"github.com/temporal-zone/temporal/x/compounder/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the compoundSettings
	for _, elem := range genState.CompoundSettingsList {
		k.SetCompoundSettings(ctx, elem)
	}
	// Set all the previousCompounding
	for _, elem := range genState.PreviousCompoundingList {
		k.SetPreviousCompounding(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.CompoundSettingsList = k.GetAllCompoundSettings(ctx)
	genesis.PreviousCompoundingList = k.GetAllPreviousCompounding(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
