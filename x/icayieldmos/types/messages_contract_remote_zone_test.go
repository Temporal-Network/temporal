package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/Temporal-Network/temporal/testutil/sample"
)

func TestMsgCreateContractRemoteZone_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateContractRemoteZone
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateContractRemoteZone{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateContractRemoteZone{
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

func TestMsgUpdateContractRemoteZone_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateContractRemoteZone
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateContractRemoteZone{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateContractRemoteZone{
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

func TestMsgDeleteContractRemoteZone_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteContractRemoteZone
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteContractRemoteZone{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteContractRemoteZone{
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
