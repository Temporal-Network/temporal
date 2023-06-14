package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateConnection = "create_connection"

var _ sdk.Msg = &MsgCreateConnection{}

func NewMsgCreateConnection(creator string, connectionId string) *MsgCreateConnection {
	return &MsgCreateConnection{
		Creator:      creator,
		ConnectionId: connectionId,
	}
}

func (msg *MsgCreateConnection) Route() string {
	return RouterKey
}

func (msg *MsgCreateConnection) Type() string {
	return TypeMsgCreateConnection
}

func (msg *MsgCreateConnection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateConnection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateConnection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
