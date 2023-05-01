package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateCompoundSettings = "create_compound_settings"
	TypeMsgUpdateCompoundSettings = "update_compound_settings"
	TypeMsgDeleteCompoundSettings = "delete_compound_settings"
)

var _ sdk.Msg = &MsgCreateCompoundSettings{}

func NewMsgCreateCompoundSettings(
	delegator string,
	index123 string,
	validatorSettings *ValidatorSetting,
	amountToRemain sdk.Coin,
	frequency *Frequency,

) *MsgCreateCompoundSettings {
	return &MsgCreateCompoundSettings{
		Delegator:         delegator,
		Index123:          index123,
		ValidatorSettings: validatorSettings,
		AmountToRemain:    amountToRemain,
		Frequency:         frequency,
	}
}

func (msg *MsgCreateCompoundSettings) Route() string {
	return RouterKey
}

func (msg *MsgCreateCompoundSettings) Type() string {
	return TypeMsgCreateCompoundSettings
}

func (msg *MsgCreateCompoundSettings) GetSigners() []sdk.AccAddress {
	delegator, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{delegator}
}

func (msg *MsgCreateCompoundSettings) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCompoundSettings) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateCompoundSettings{}

func NewMsgUpdateCompoundSettings(
	delegator string,
	index123 string,
	validatorSettings *ValidatorSetting,
	amountToRemain sdk.Coin,
	frequency *Frequency,

) *MsgUpdateCompoundSettings {
	return &MsgUpdateCompoundSettings{
		Delegator:         delegator,
		Index123:          index123,
		ValidatorSettings: validatorSettings,
		AmountToRemain:    amountToRemain,
		Frequency:         frequency,
	}
}

func (msg *MsgUpdateCompoundSettings) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCompoundSettings) Type() string {
	return TypeMsgUpdateCompoundSettings
}

func (msg *MsgUpdateCompoundSettings) GetSigners() []sdk.AccAddress {
	delegator, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{delegator}
}

func (msg *MsgUpdateCompoundSettings) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCompoundSettings) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteCompoundSettings{}

func NewMsgDeleteCompoundSettings(
	delegator string,
	index123 string,

) *MsgDeleteCompoundSettings {
	return &MsgDeleteCompoundSettings{
		Delegator: delegator,
		Index123:  index123,
	}
}
func (msg *MsgDeleteCompoundSettings) Route() string {
	return RouterKey
}

func (msg *MsgDeleteCompoundSettings) Type() string {
	return TypeMsgDeleteCompoundSettings
}

func (msg *MsgDeleteCompoundSettings) GetSigners() []sdk.AccAddress {
	delegator, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{delegator}
}

func (msg *MsgDeleteCompoundSettings) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteCompoundSettings) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address (%s)", err)
	}
	return nil
}
