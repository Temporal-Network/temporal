package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RemoteContractCompoundSettingsAll(goCtx context.Context, req *types.QueryAllRemoteContractCompoundSettingsRequest) (*types.QueryAllRemoteContractCompoundSettingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var remoteContractCompoundSettingss []types.RemoteContractCompoundSettings
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	remoteContractCompoundSettingsStore := prefix.NewStore(store, types.KeyPrefix(types.RemoteContractCompoundSettingsKey))

	pageRes, err := query.Paginate(remoteContractCompoundSettingsStore, req.Pagination, func(key []byte, value []byte) error {
		var remoteContractCompoundSettings types.RemoteContractCompoundSettings
		if err := k.cdc.Unmarshal(value, &remoteContractCompoundSettings); err != nil {
			return err
		}

		remoteContractCompoundSettingss = append(remoteContractCompoundSettingss, remoteContractCompoundSettings)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRemoteContractCompoundSettingsResponse{RemoteContractCompoundSettings: remoteContractCompoundSettingss, Pagination: pageRes}, nil
}

func (k Keeper) RemoteContractCompoundSettings(goCtx context.Context, req *types.QueryGetRemoteContractCompoundSettingsRequest) (*types.QueryGetRemoteContractCompoundSettingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	remoteContractCompoundSettings, found := k.GetRemoteContractCompoundSettings(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetRemoteContractCompoundSettingsResponse{RemoteContractCompoundSettings: remoteContractCompoundSettings}, nil
}
