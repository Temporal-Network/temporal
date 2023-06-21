package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

    "github.com/Temporal-Network/temporal/x/icayieldmos/types"
)

func TestRemoteContractCompoundSettingsMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateRemoteContractCompoundSettings(ctx, &types.MsgCreateRemoteContractCompoundSettings{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestRemoteContractCompoundSettingsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateRemoteContractCompoundSettings
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateRemoteContractCompoundSettings{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateRemoteContractCompoundSettings{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateRemoteContractCompoundSettings{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateRemoteContractCompoundSettings(ctx, &types.MsgCreateRemoteContractCompoundSettings{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateRemoteContractCompoundSettings(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestRemoteContractCompoundSettingsMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteRemoteContractCompoundSettings
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteRemoteContractCompoundSettings{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteRemoteContractCompoundSettings{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteRemoteContractCompoundSettings{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateRemoteContractCompoundSettings(ctx, &types.MsgCreateRemoteContractCompoundSettings{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteRemoteContractCompoundSettings(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
