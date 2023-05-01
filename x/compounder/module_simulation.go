package compounder

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"temporal/testutil/sample"
	compoundersimulation "temporal/x/compounder/simulation"
	"temporal/x/compounder/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = compoundersimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateCompoundSettings = "op_weight_msg_compound_settings"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCompoundSettings int = 100

	opWeightMsgUpdateCompoundSettings = "op_weight_msg_compound_settings"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateCompoundSettings int = 100

	opWeightMsgDeleteCompoundSettings = "op_weight_msg_compound_settings"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteCompoundSettings int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	compounderGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		CompoundSettingsList: []types.CompoundSettings{
			{
				Delegator: sample.AccAddress(),
				Index123:  "0",
			},
			{
				Delegator: sample.AccAddress(),
				Index123:  "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&compounderGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateCompoundSettings int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateCompoundSettings, &weightMsgCreateCompoundSettings, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCompoundSettings = defaultWeightMsgCreateCompoundSettings
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCompoundSettings,
		compoundersimulation.SimulateMsgCreateCompoundSettings(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateCompoundSettings int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateCompoundSettings, &weightMsgUpdateCompoundSettings, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCompoundSettings = defaultWeightMsgUpdateCompoundSettings
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCompoundSettings,
		compoundersimulation.SimulateMsgUpdateCompoundSettings(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteCompoundSettings int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteCompoundSettings, &weightMsgDeleteCompoundSettings, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteCompoundSettings = defaultWeightMsgDeleteCompoundSettings
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteCompoundSettings,
		compoundersimulation.SimulateMsgDeleteCompoundSettings(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
