package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateICARemoteZone = "create_ica_remote_zone"
	TypeMsgUpdateICARemoteZone = "update_ica_remote_zone"
	TypeMsgDeleteICARemoteZone = "delete_ica_remote_zone"
)

var _ sdk.Msg = &MsgCreateICARemoteZone{}

func NewMsgCreateICARemoteZone(creator string, connectionId string, remoteChainId string, active bool, bech32Prefix string) *MsgCreateICARemoteZone {
  return &MsgCreateICARemoteZone{
		Creator: creator,
    ConnectionId: connectionId,
    RemoteChainId: remoteChainId,
    Active: active,
    Bech32Prefix: bech32Prefix,
	}
}

func (msg *MsgCreateICARemoteZone) Route() string {
  return RouterKey
}

func (msg *MsgCreateICARemoteZone) Type() string {
  return TypeMsgCreateICARemoteZone
}

func (msg *MsgCreateICARemoteZone) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreateICARemoteZone) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateICARemoteZone) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

var _ sdk.Msg = &MsgUpdateICARemoteZone{}

func NewMsgUpdateICARemoteZone(creator string, id uint64, connectionId string, remoteChainId string, active bool, bech32Prefix string) *MsgUpdateICARemoteZone {
  return &MsgUpdateICARemoteZone{
        Id: id,
		Creator: creator,
    ConnectionId: connectionId,
    RemoteChainId: remoteChainId,
    Active: active,
    Bech32Prefix: bech32Prefix,
	}
}

func (msg *MsgUpdateICARemoteZone) Route() string {
  return RouterKey
}

func (msg *MsgUpdateICARemoteZone) Type() string {
  return TypeMsgUpdateICARemoteZone
}

func (msg *MsgUpdateICARemoteZone) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateICARemoteZone) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateICARemoteZone) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
   return nil
}

var _ sdk.Msg = &MsgDeleteICARemoteZone{}

func NewMsgDeleteICARemoteZone(creator string, id uint64) *MsgDeleteICARemoteZone {
  return &MsgDeleteICARemoteZone{
        Id: id,
		Creator: creator,
	}
} 
func (msg *MsgDeleteICARemoteZone) Route() string {
  return RouterKey
}

func (msg *MsgDeleteICARemoteZone) Type() string {
  return TypeMsgDeleteICARemoteZone
}

func (msg *MsgDeleteICARemoteZone) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteICARemoteZone) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteICARemoteZone) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
  return nil
}
