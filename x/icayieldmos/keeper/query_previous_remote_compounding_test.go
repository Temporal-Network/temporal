package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/temporal-zone/temporal/testutil/keeper"
	"github.com/temporal-zone/temporal/testutil/nullify"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPreviousRemoteCompoundingQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPreviousRemoteCompounding(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetPreviousRemoteCompoundingRequest
		response *types.QueryGetPreviousRemoteCompoundingResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetPreviousRemoteCompoundingRequest{
				RemoteContractCompoundSetting: msgs[0].RemoteContractCompoundSetting,
			},
			response: &types.QueryGetPreviousRemoteCompoundingResponse{PreviousRemoteCompounding: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetPreviousRemoteCompoundingRequest{
				RemoteContractCompoundSetting: msgs[1].RemoteContractCompoundSetting,
			},
			response: &types.QueryGetPreviousRemoteCompoundingResponse{PreviousRemoteCompounding: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetPreviousRemoteCompoundingRequest{
				RemoteContractCompoundSetting: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.PreviousRemoteCompounding(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestPreviousRemoteCompoundingQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPreviousRemoteCompounding(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPreviousRemoteCompoundingRequest {
		return &types.QueryAllPreviousRemoteCompoundingRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PreviousRemoteCompoundingAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PreviousRemoteCompounding), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PreviousRemoteCompounding),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PreviousRemoteCompoundingAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PreviousRemoteCompounding), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PreviousRemoteCompounding),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.PreviousRemoteCompoundingAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.PreviousRemoteCompounding),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.PreviousRemoteCompoundingAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
