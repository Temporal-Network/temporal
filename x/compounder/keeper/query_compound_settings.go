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

func (k Keeper) CompoundSettingsAll(goCtx context.Context, req *types.QueryAllCompoundSettingsRequest) (*types.QueryAllCompoundSettingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var compoundSettingss []types.CompoundSettings
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	compoundSettingsStore := prefix.NewStore(store, types.KeyPrefix(types.CompoundSettingsKeyPrefix))

	pageRes, err := query.Paginate(compoundSettingsStore, req.Pagination, func(key []byte, value []byte) error {
		var compoundSettings types.CompoundSettings
		if err := k.cdc.Unmarshal(value, &compoundSettings); err != nil {
			return err
		}

		compoundSettingss = append(compoundSettingss, compoundSettings)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCompoundSettingsResponse{CompoundSettings: compoundSettingss, Pagination: pageRes}, nil
}

func (k Keeper) CompoundSettings(goCtx context.Context, req *types.QueryGetCompoundSettingsRequest) (*types.QueryGetCompoundSettingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetCompoundSettings(
		ctx,
		req.Index123,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCompoundSettingsResponse{CompoundSettings: val}, nil
}
