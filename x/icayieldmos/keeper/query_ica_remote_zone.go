package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ICARemoteZoneAll(goCtx context.Context, req *types.QueryAllICARemoteZoneRequest) (*types.QueryAllICARemoteZoneResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var iCARemoteZones []types.ICARemoteZone
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	iCARemoteZoneStore := prefix.NewStore(store, types.KeyPrefix(types.ICARemoteZoneKey))

	pageRes, err := query.Paginate(iCARemoteZoneStore, req.Pagination, func(key []byte, value []byte) error {
		var iCARemoteZone types.ICARemoteZone
		if err := k.cdc.Unmarshal(value, &iCARemoteZone); err != nil {
			return err
		}

		iCARemoteZones = append(iCARemoteZones, iCARemoteZone)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllICARemoteZoneResponse{ICARemoteZone: iCARemoteZones, Pagination: pageRes}, nil
}

func (k Keeper) ICARemoteZone(goCtx context.Context, req *types.QueryGetICARemoteZoneRequest) (*types.QueryGetICARemoteZoneResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	iCARemoteZone, found := k.GetICARemoteZone(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetICARemoteZoneResponse{ICARemoteZone: iCARemoteZone}, nil
}
