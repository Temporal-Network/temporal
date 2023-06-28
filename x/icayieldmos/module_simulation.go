package icayieldmos

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/temporal-zone/temporal/testutil/sample"
	icayieldmossimulation "github.com/temporal-zone/temporal/x/icayieldmos/simulation"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
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

	opWeightMsgSendTestContractMessages = "op_weight_msg_send_test_contract_messages"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendTestContractMessages int = 100

	opWeightMsgCreateRemoteContractCompoundSettings = "op_weight_msg_remote_contract_compound_settings"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateRemoteContractCompoundSettings int = 100

	opWeightMsgUpdateRemoteContractCompoundSettings = "op_weight_msg_remote_contract_compound_settings"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateRemoteContractCompoundSettings int = 100

	opWeightMsgDeleteRemoteContractCompoundSettings = "op_weight_msg_remote_contract_compound_settings"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteRemoteContractCompoundSettings int = 100

	opWeightMsgCreateICARemoteZone = "op_weight_msg_ica_remote_zone"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateICARemoteZone int = 100

	opWeightMsgUpdateICARemoteZone = "op_weight_msg_ica_remote_zone"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateICARemoteZone int = 100

	opWeightMsgDeleteICARemoteZone = "op_weight_msg_ica_remote_zone"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteICARemoteZone int = 100

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
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		ContractRemoteZoneCount: 2,
		RemoteContractCompoundSettingsList: []types.RemoteContractCompoundSettings{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		RemoteContractCompoundSettingsCount: 2,
			ICARemoteZoneList: []types.ICARemoteZone{
		{
			Id: 0,
			Creator: sample.AccAddress(),

		},
		{
			Id: 1,
			Creator: sample.AccAddress(),

		},
	},
	ICARemoteZoneCount: 2,
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

	var weightMsgSendTestContractMessages int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSendTestContractMessages, &weightMsgSendTestContractMessages, nil,
		func(_ *rand.Rand) {
			weightMsgSendTestContractMessages = defaultWeightMsgSendTestContractMessages
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendTestContractMessages,
		icayieldmossimulation.SimulateMsgSendTestContractMessages(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateRemoteContractCompoundSettings int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateRemoteContractCompoundSettings, &weightMsgCreateRemoteContractCompoundSettings, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRemoteContractCompoundSettings = defaultWeightMsgCreateRemoteContractCompoundSettings
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateRemoteContractCompoundSettings,
		icayieldmossimulation.SimulateMsgCreateRemoteContractCompoundSettings(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateRemoteContractCompoundSettings int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateRemoteContractCompoundSettings, &weightMsgUpdateRemoteContractCompoundSettings, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRemoteContractCompoundSettings = defaultWeightMsgUpdateRemoteContractCompoundSettings
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateRemoteContractCompoundSettings,
		icayieldmossimulation.SimulateMsgUpdateRemoteContractCompoundSettings(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteRemoteContractCompoundSettings int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteRemoteContractCompoundSettings, &weightMsgDeleteRemoteContractCompoundSettings, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteRemoteContractCompoundSettings = defaultWeightMsgDeleteRemoteContractCompoundSettings
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteRemoteContractCompoundSettings,
		icayieldmossimulation.SimulateMsgDeleteRemoteContractCompoundSettings(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateICARemoteZone int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateICARemoteZone, &weightMsgCreateICARemoteZone, nil,
		func(_ *rand.Rand) {
			weightMsgCreateICARemoteZone = defaultWeightMsgCreateICARemoteZone
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateICARemoteZone,
		icayieldmossimulation.SimulateMsgCreateICARemoteZone(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateICARemoteZone int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateICARemoteZone, &weightMsgUpdateICARemoteZone, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateICARemoteZone = defaultWeightMsgUpdateICARemoteZone
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateICARemoteZone,
		icayieldmossimulation.SimulateMsgUpdateICARemoteZone(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteICARemoteZone int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteICARemoteZone, &weightMsgDeleteICARemoteZone, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteICARemoteZone = defaultWeightMsgDeleteICARemoteZone
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteICARemoteZone,
		icayieldmossimulation.SimulateMsgDeleteICARemoteZone(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
