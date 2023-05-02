package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"temporal/x/compounder/types"
)

func (k Keeper) PreviousCompoundingAll(goCtx context.Context, req *types.QueryAllPreviousCompoundingRequest) (*types.QueryAllPreviousCompoundingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var previousCompoundings []types.PreviousCompounding
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	previousCompoundingStore := prefix.NewStore(store, types.KeyPrefix(types.PreviousCompoundingKeyPrefix))

	pageRes, err := query.Paginate(previousCompoundingStore, req.Pagination, func(key []byte, value []byte) error {
		var previousCompounding types.PreviousCompounding
		if err := k.cdc.Unmarshal(value, &previousCompounding); err != nil {
			return err
		}

		previousCompoundings = append(previousCompoundings, previousCompounding)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPreviousCompoundingResponse{PreviousCompounding: previousCompoundings, Pagination: pageRes}, nil
}

func (k Keeper) PreviousCompounding(goCtx context.Context, req *types.QueryGetPreviousCompoundingRequest) (*types.QueryGetPreviousCompoundingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetPreviousCompounding(
		ctx,
		req.Delegator,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPreviousCompoundingResponse{PreviousCompounding: val}, nil
}
