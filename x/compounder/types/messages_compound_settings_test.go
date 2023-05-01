package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"temporal/testutil/sample"
)

func TestMsgCreateCompoundSettings_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateCompoundSettings
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateCompoundSettings{
				Delegator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateCompoundSettings{
				Delegator: sample.AccAddress(),
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

func TestMsgUpdateCompoundSettings_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateCompoundSettings
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateCompoundSettings{
				Delegator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateCompoundSettings{
				Delegator: sample.AccAddress(),
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

func TestMsgDeleteCompoundSettings_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteCompoundSettings
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteCompoundSettings{
				Delegator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteCompoundSettings{
				Delegator: sample.AccAddress(),
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
