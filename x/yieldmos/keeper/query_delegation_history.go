package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"temporal/x/yieldmos/types"
)

func (k Keeper) DelegationHistoryAll(goCtx context.Context, req *types.QueryAllDelegationHistoryRequest) (*types.QueryAllDelegationHistoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var delegationHistorys []types.DelegationHistory
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	delegationHistoryStore := prefix.NewStore(store, types.KeyPrefix(types.DelegationHistoryKeyPrefix))

	pageRes, err := query.Paginate(delegationHistoryStore, req.Pagination, func(key []byte, value []byte) error {
		var delegationHistory types.DelegationHistory
		if err := k.cdc.Unmarshal(value, &delegationHistory); err != nil {
			return err
		}

		delegationHistorys = append(delegationHistorys, delegationHistory)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDelegationHistoryResponse{DelegationHistory: delegationHistorys, Pagination: pageRes}, nil
}

func (k Keeper) DelegationHistory(goCtx context.Context, req *types.QueryGetDelegationHistoryRequest) (*types.QueryGetDelegationHistoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetDelegationHistory(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDelegationHistoryResponse{DelegationHistory: val}, nil
}
