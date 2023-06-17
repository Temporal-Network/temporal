package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendTestContractMessages = "send_test_contract_messages"

var _ sdk.Msg = &MsgSendTestContractMessages{}

func NewMsgSendTestContractMessages(creator string, connectionId string) *MsgSendTestContractMessages {
	return &MsgSendTestContractMessages{
		Creator:      creator,
		ConnectionId: connectionId,
	}
}

func (msg *MsgSendTestContractMessages) Route() string {
	return RouterKey
}

func (msg *MsgSendTestContractMessages) Type() string {
	return TypeMsgSendTestContractMessages
}

func (msg *MsgSendTestContractMessages) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendTestContractMessages) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendTestContractMessages) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
