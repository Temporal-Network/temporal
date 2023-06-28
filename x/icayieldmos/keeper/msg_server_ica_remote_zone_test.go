package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

    "github.com/temporal-zone/temporal/x/icayieldmos/types"
)

func TestICARemoteZoneMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateICARemoteZone(ctx, &types.MsgCreateICARemoteZone{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestICARemoteZoneMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateICARemoteZone
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateICARemoteZone{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateICARemoteZone{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateICARemoteZone{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateICARemoteZone(ctx, &types.MsgCreateICARemoteZone{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateICARemoteZone(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestICARemoteZoneMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteICARemoteZone
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteICARemoteZone{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteICARemoteZone{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteICARemoteZone{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateICARemoteZone(ctx, &types.MsgCreateICARemoteZone{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteICARemoteZone(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
