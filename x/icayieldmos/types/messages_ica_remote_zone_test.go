package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/temporal-zone/temporal/testutil/sample"
)

func TestMsgCreateICARemoteZone_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateICARemoteZone
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateICARemoteZone{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateICARemoteZone{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateICARemoteZone_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateICARemoteZone
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateICARemoteZone{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateICARemoteZone{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteICARemoteZone_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteICARemoteZone
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteICARemoteZone{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteICARemoteZone{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
