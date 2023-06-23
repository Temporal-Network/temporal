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

func (k Keeper) ContractRemoteZoneAll(goCtx context.Context, req *types.QueryAllContractRemoteZoneRequest) (*types.QueryAllContractRemoteZoneResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var contractRemoteZones []types.ContractRemoteZone
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	contractRemoteZoneStore := prefix.NewStore(store, types.KeyPrefix(types.ContractRemoteZoneKey))

	pageRes, err := query.Paginate(contractRemoteZoneStore, req.Pagination, func(key []byte, value []byte) error {
		var contractRemoteZone types.ContractRemoteZone
		if err := k.cdc.Unmarshal(value, &contractRemoteZone); err != nil {
			return err
		}

		contractRemoteZones = append(contractRemoteZones, contractRemoteZone)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllContractRemoteZoneResponse{ContractRemoteZone: contractRemoteZones, Pagination: pageRes}, nil
}

func (k Keeper) ContractRemoteZone(goCtx context.Context, req *types.QueryGetContractRemoteZoneRequest) (*types.QueryGetContractRemoteZoneResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	contractRemoteZone, found := k.GetContractRemoteZone(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetContractRemoteZoneResponse{ContractRemoteZone: contractRemoteZone}, nil
}
