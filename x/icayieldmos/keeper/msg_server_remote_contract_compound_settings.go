package keeper

import (
    "fmt"
	"context"

    "github.com/Temporal-Network/temporal/x/icayieldmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


func (k msgServer) CreateRemoteContractCompoundSettings(goCtx context.Context,  msg *types.MsgCreateRemoteContractCompoundSettings) (*types.MsgCreateRemoteContractCompoundSettingsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    var remoteContractCompoundSettings = types.RemoteContractCompoundSettings{
        Creator: msg.Creator,
        ContractRemoteZone: msg.ContractRemoteZone,
        RemoteDelegatorAddress: msg.RemoteDelegatorAddress,
        CompoundSettings: msg.CompoundSettings,
        FrequencyInSeconds: msg.FrequencyInSeconds,
        RemoteContractAddress: msg.RemoteContractAddress,
    }

    id := k.AppendRemoteContractCompoundSettings(
        ctx,
        remoteContractCompoundSettings,
    )

	return &types.MsgCreateRemoteContractCompoundSettingsResponse{
	    Id: id,
	}, nil
}

func (k msgServer) UpdateRemoteContractCompoundSettings(goCtx context.Context,  msg *types.MsgUpdateRemoteContractCompoundSettings) (*types.MsgUpdateRemoteContractCompoundSettingsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    var remoteContractCompoundSettings = types.RemoteContractCompoundSettings{
		Creator: msg.Creator,
		Id:      msg.Id,
    	ContractRemoteZone: msg.ContractRemoteZone,
    	RemoteDelegatorAddress: msg.RemoteDelegatorAddress,
    	CompoundSettings: msg.CompoundSettings,
    	FrequencyInSeconds: msg.FrequencyInSeconds,
    	RemoteContractAddress: msg.RemoteContractAddress,
	}

    // Checks that the element exists
    val, found := k.GetRemoteContractCompoundSettings(ctx, msg.Id)
    if !found {
        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
    }

    // Checks if the msg creator is the same as the current owner
    if msg.Creator != val.Creator {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }

	k.SetRemoteContractCompoundSettings(ctx, remoteContractCompoundSettings)

	return &types.MsgUpdateRemoteContractCompoundSettingsResponse{}, nil
}

func (k msgServer) DeleteRemoteContractCompoundSettings(goCtx context.Context,  msg *types.MsgDeleteRemoteContractCompoundSettings) (*types.MsgDeleteRemoteContractCompoundSettingsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // Checks that the element exists
    val, found := k.GetRemoteContractCompoundSettings(ctx, msg.Id)
    if !found {
        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
    }

    // Checks if the msg creator is the same as the current owner
    if msg.Creator != val.Creator {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }

	k.RemoveRemoteContractCompoundSettings(ctx, msg.Id)

	return &types.MsgDeleteRemoteContractCompoundSettingsResponse{}, nil
}
