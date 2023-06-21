package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/Temporal-Network/temporal/testutil/sample"
)

func TestMsgCreateRemoteContractCompoundSettings_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateRemoteContractCompoundSettings
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateRemoteContractCompoundSettings{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateRemoteContractCompoundSettings{
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

func TestMsgUpdateRemoteContractCompoundSettings_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateRemoteContractCompoundSettings
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateRemoteContractCompoundSettings{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateRemoteContractCompoundSettings{
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

func TestMsgDeleteRemoteContractCompoundSettings_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteRemoteContractCompoundSettings
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteRemoteContractCompoundSettings{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteRemoteContractCompoundSettings{
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
