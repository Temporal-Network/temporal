package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateConnection{}, "icayieldmos/CreateConnection", nil)
	cdc.RegisterConcrete(&MsgCreateContractRemoteZone{}, "icayieldmos/CreateContractRemoteZone", nil)
	cdc.RegisterConcrete(&MsgUpdateContractRemoteZone{}, "icayieldmos/UpdateContractRemoteZone", nil)
	cdc.RegisterConcrete(&MsgDeleteContractRemoteZone{}, "icayieldmos/DeleteContractRemoteZone", nil)
	cdc.RegisterConcrete(&MsgSendTestContractMessages{}, "icayieldmos/SendTestContractMessages", nil)
	cdc.RegisterConcrete(&MsgCreateRemoteContractCompoundSettings{}, "icayieldmos/CreateRemoteContractCompoundSettings", nil)
cdc.RegisterConcrete(&MsgUpdateRemoteContractCompoundSettings{}, "icayieldmos/UpdateRemoteContractCompoundSettings", nil)
cdc.RegisterConcrete(&MsgDeleteRemoteContractCompoundSettings{}, "icayieldmos/DeleteRemoteContractCompoundSettings", nil)
// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateConnection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateContractRemoteZone{},
		&MsgUpdateContractRemoteZone{},
		&MsgDeleteContractRemoteZone{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendTestContractMessages{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCreateRemoteContractCompoundSettings{},
	&MsgUpdateRemoteContractCompoundSettings{},
	&MsgDeleteRemoteContractCompoundSettings{},
)
// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
