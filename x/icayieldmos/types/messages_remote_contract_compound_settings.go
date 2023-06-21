package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateRemoteContractCompoundSettings = "create_remote_contract_compound_settings"
	TypeMsgUpdateRemoteContractCompoundSettings = "update_remote_contract_compound_settings"
	TypeMsgDeleteRemoteContractCompoundSettings = "delete_remote_contract_compound_settings"
)

var _ sdk.Msg = &MsgCreateRemoteContractCompoundSettings{}

func NewMsgCreateRemoteContractCompoundSettings(creator string, contractRemoteZone uint64, remoteDelegatorAddress string, compoundSettings string, frequencyInSeconds uint64, remoteContractAddress string) *MsgCreateRemoteContractCompoundSettings {
	return &MsgCreateRemoteContractCompoundSettings{
		Creator:                creator,
		ContractRemoteZone:     contractRemoteZone,
		RemoteDelegatorAddress: remoteDelegatorAddress,
		CompoundSettings:       compoundSettings,
		FrequencyInSeconds:     frequencyInSeconds,
		RemoteContractAddress:  remoteContractAddress,
	}
}

func (msg *MsgCreateRemoteContractCompoundSettings) Route() string {
	return RouterKey
}

func (msg *MsgCreateRemoteContractCompoundSettings) Type() string {
	return TypeMsgCreateRemoteContractCompoundSettings
}

func (msg *MsgCreateRemoteContractCompoundSettings) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRemoteContractCompoundSettings) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRemoteContractCompoundSettings) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateRemoteContractCompoundSettings{}

func NewMsgUpdateRemoteContractCompoundSettings(creator string, id uint64, contractRemoteZone uint64, remoteDelegatorAddress string, compoundSettings string, frequencyInSeconds uint64, remoteContractAddress string) *MsgUpdateRemoteContractCompoundSettings {
	return &MsgUpdateRemoteContractCompoundSettings{
		Id:                     id,
		Creator:                creator,
		ContractRemoteZone:     contractRemoteZone,
		RemoteDelegatorAddress: remoteDelegatorAddress,
		CompoundSettings:       compoundSettings,
		FrequencyInSeconds:     frequencyInSeconds,
		RemoteContractAddress:  remoteContractAddress,
	}
}

func (msg *MsgUpdateRemoteContractCompoundSettings) Route() string {
	return RouterKey
}

func (msg *MsgUpdateRemoteContractCompoundSettings) Type() string {
	return TypeMsgUpdateRemoteContractCompoundSettings
}

func (msg *MsgUpdateRemoteContractCompoundSettings) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateRemoteContractCompoundSettings) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateRemoteContractCompoundSettings) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteRemoteContractCompoundSettings{}

func NewMsgDeleteRemoteContractCompoundSettings(creator string, id uint64) *MsgDeleteRemoteContractCompoundSettings {
	return &MsgDeleteRemoteContractCompoundSettings{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteRemoteContractCompoundSettings) Route() string {
	return RouterKey
}

func (msg *MsgDeleteRemoteContractCompoundSettings) Type() string {
	return TypeMsgDeleteRemoteContractCompoundSettings
}

func (msg *MsgDeleteRemoteContractCompoundSettings) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteRemoteContractCompoundSettings) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRemoteContractCompoundSettings) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
