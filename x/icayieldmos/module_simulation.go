package icayieldmos

import (
	"math/rand"

	"github.com/Temporal-Network/temporal/testutil/sample"
	icayieldmossimulation "github.com/Temporal-Network/temporal/x/icayieldmos/simulation"
	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = icayieldmossimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateConnection = "op_weight_msg_create_connection"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateConnection int = 100

	opWeightMsgCreateContractRemoteZone = "op_weight_msg_contract_remote_zone"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateContractRemoteZone int = 100

	opWeightMsgUpdateContractRemoteZone = "op_weight_msg_contract_remote_zone"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateContractRemoteZone int = 100

	opWeightMsgDeleteContractRemoteZone = "op_weight_msg_contract_remote_zone"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteContractRemoteZone int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	icayieldmosGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
			ContractRemoteZoneList: []types.ContractRemoteZone{
		{
			Id: 0,
			Creator: sample.AccAddress(),

		},
		{
			Id: 1,
			Creator: sample.AccAddress(),

		},
	},
	ContractRemoteZoneCount: 2,
	// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&icayieldmosGenesis)
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

	var weightMsgCreateConnection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateConnection, &weightMsgCreateConnection, nil,
		func(_ *rand.Rand) {
			weightMsgCreateConnection = defaultWeightMsgCreateConnection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateConnection,
		icayieldmossimulation.SimulateMsgCreateConnection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateContractRemoteZone int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateContractRemoteZone, &weightMsgCreateContractRemoteZone, nil,
		func(_ *rand.Rand) {
			weightMsgCreateContractRemoteZone = defaultWeightMsgCreateContractRemoteZone
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateContractRemoteZone,
		icayieldmossimulation.SimulateMsgCreateContractRemoteZone(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateContractRemoteZone int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateContractRemoteZone, &weightMsgUpdateContractRemoteZone, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateContractRemoteZone = defaultWeightMsgUpdateContractRemoteZone
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateContractRemoteZone,
		icayieldmossimulation.SimulateMsgUpdateContractRemoteZone(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteContractRemoteZone int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteContractRemoteZone, &weightMsgDeleteContractRemoteZone, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteContractRemoteZone = defaultWeightMsgDeleteContractRemoteZone
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteContractRemoteZone,
		icayieldmossimulation.SimulateMsgDeleteContractRemoteZone(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
