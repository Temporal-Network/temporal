package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateContractRemoteZone = "create_contract_remote_zone"
	TypeMsgUpdateContractRemoteZone = "update_contract_remote_zone"
	TypeMsgDeleteContractRemoteZone = "delete_contract_remote_zone"
)

var _ sdk.Msg = &MsgCreateContractRemoteZone{}

func NewMsgCreateContractRemoteZone(creator string, connectionId string, remoteChainId string, active bool) *MsgCreateContractRemoteZone {
	return &MsgCreateContractRemoteZone{
		Creator:       creator,
		ConnectionId:  connectionId,
		RemoteChainId: remoteChainId,
		Active:        active,
	}
}

func (msg *MsgCreateContractRemoteZone) Route() string {
	return RouterKey
}

func (msg *MsgCreateContractRemoteZone) Type() string {
	return TypeMsgCreateContractRemoteZone
}

func (msg *MsgCreateContractRemoteZone) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateContractRemoteZone) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateContractRemoteZone) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateContractRemoteZone{}

func NewMsgUpdateContractRemoteZone(creator string, id uint64, connectionId string, remoteChainId string, active bool) *MsgUpdateContractRemoteZone {
	return &MsgUpdateContractRemoteZone{
		Id:            id,
		Creator:       creator,
		ConnectionId:  connectionId,
		RemoteChainId: remoteChainId,
		Active:        active,
	}
}

func (msg *MsgUpdateContractRemoteZone) Route() string {
	return RouterKey
}

func (msg *MsgUpdateContractRemoteZone) Type() string {
	return TypeMsgUpdateContractRemoteZone
}

func (msg *MsgUpdateContractRemoteZone) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateContractRemoteZone) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateContractRemoteZone) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteContractRemoteZone{}

func NewMsgDeleteContractRemoteZone(creator string, id uint64) *MsgDeleteContractRemoteZone {
	return &MsgDeleteContractRemoteZone{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteContractRemoteZone) Route() string {
	return RouterKey
}

func (msg *MsgDeleteContractRemoteZone) Type() string {
	return TypeMsgDeleteContractRemoteZone
}

func (msg *MsgDeleteContractRemoteZone) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteContractRemoteZone) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteContractRemoteZone) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
