package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/temporal-zone/temporal/testutil/keeper"
	"github.com/temporal-zone/temporal/x/compounder/keeper"
	"github.com/temporal-zone/temporal/x/compounder/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCompoundSettingsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.CompounderKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	delegator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateCompoundSettings{Delegator: delegator}
		_, err := srv.CreateCompoundSettings(wctx, expected)

		if i == 0 {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
		rst, found := k.GetCompoundSettings(ctx,
			expected.Delegator,
		)
		require.True(t, found)
		require.Equal(t, expected.Delegator, rst.Delegator)
	}
}

func TestCompoundSettingsMsgServerUpdate(t *testing.T) {
	delegator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateCompoundSettings
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateCompoundSettings{Delegator: delegator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateCompoundSettings{Delegator: "B"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CompounderKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateCompoundSettings{Delegator: delegator}
			_, err := srv.CreateCompoundSettings(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateCompoundSettings(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetCompoundSettings(ctx,
					expected.Delegator,
				)
				require.True(t, found)
				require.Equal(t, expected.Delegator, rst.Delegator)
			}
		})
	}
}

func TestCompoundSettingsMsgServerDelete(t *testing.T) {
	delegator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteCompoundSettings
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteCompoundSettings{Delegator: delegator},
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteCompoundSettings{Delegator: "B"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CompounderKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateCompoundSettings(wctx, &types.MsgCreateCompoundSettings{Delegator: delegator})
			require.NoError(t, err)
			_, err = srv.DeleteCompoundSettings(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetCompoundSettings(ctx,
					tc.request.Delegator,
				)
				require.False(t, found)
			}
		})
	}
}
