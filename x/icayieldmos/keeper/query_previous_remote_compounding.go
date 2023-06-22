package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PreviousRemoteCompoundingAll(goCtx context.Context, req *types.QueryAllPreviousRemoteCompoundingRequest) (*types.QueryAllPreviousRemoteCompoundingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var previousRemoteCompoundings []types.PreviousRemoteCompounding
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	previousRemoteCompoundingStore := prefix.NewStore(store, types.KeyPrefix(types.PreviousRemoteCompoundingKeyPrefix))

	pageRes, err := query.Paginate(previousRemoteCompoundingStore, req.Pagination, func(key []byte, value []byte) error {
		var previousRemoteCompounding types.PreviousRemoteCompounding
		if err := k.cdc.Unmarshal(value, &previousRemoteCompounding); err != nil {
			return err
		}

		previousRemoteCompoundings = append(previousRemoteCompoundings, previousRemoteCompounding)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPreviousRemoteCompoundingResponse{PreviousRemoteCompounding: previousRemoteCompoundings, Pagination: pageRes}, nil
}

func (k Keeper) PreviousRemoteCompounding(goCtx context.Context, req *types.QueryGetPreviousRemoteCompoundingRequest) (*types.QueryGetPreviousRemoteCompoundingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetPreviousRemoteCompounding(
	    ctx,
	    req.RemoteContractCompoundSetting,
        )
	if !found {
	    return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPreviousRemoteCompoundingResponse{PreviousRemoteCompounding: val}, nil
}