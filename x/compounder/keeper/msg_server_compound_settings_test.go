package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "temporal/testutil/keeper"
	"temporal/x/compounder/keeper"
	"temporal/x/compounder/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCompoundSettingsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.CompounderKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	delegator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateCompoundSettings{Delegator: delegator,
			Index123: strconv.Itoa(i),
		}
		_, err := srv.CreateCompoundSettings(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetCompoundSettings(ctx,
			expected.Index123,
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
			desc: "Completed",
			request: &types.MsgUpdateCompoundSettings{Delegator: delegator,
				Index123: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateCompoundSettings{Delegator: "B",
				Index123: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateCompoundSettings{Delegator: delegator,
				Index123: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CompounderKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateCompoundSettings{Delegator: delegator,
				Index123: strconv.Itoa(0),
			}
			_, err := srv.CreateCompoundSettings(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateCompoundSettings(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetCompoundSettings(ctx,
					expected.Index123,
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
			desc: "Completed",
			request: &types.MsgDeleteCompoundSettings{Delegator: delegator,
				Index123: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteCompoundSettings{Delegator: "B",
				Index123: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteCompoundSettings{Delegator: delegator,
				Index123: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CompounderKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateCompoundSettings(wctx, &types.MsgCreateCompoundSettings{Delegator: delegator,
				Index123: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteCompoundSettings(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetCompoundSettings(ctx,
					tc.request.Index123,
				)
				require.False(t, found)
			}
		})
	}
}
