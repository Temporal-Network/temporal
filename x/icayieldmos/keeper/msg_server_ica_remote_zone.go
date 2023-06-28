package keeper

import (
    "fmt"
	"context"

    "github.com/temporal-zone/temporal/x/icayieldmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


func (k msgServer) CreateICARemoteZone(goCtx context.Context,  msg *types.MsgCreateICARemoteZone) (*types.MsgCreateICARemoteZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    var iCARemoteZone = types.ICARemoteZone{
        Creator: msg.Creator,
        ConnectionId: msg.ConnectionId,
        RemoteChainId: msg.RemoteChainId,
        Active: msg.Active,
        Bech32Prefix: msg.Bech32Prefix,
    }

    id := k.AppendICARemoteZone(
        ctx,
        iCARemoteZone,
    )

	return &types.MsgCreateICARemoteZoneResponse{
	    Id: id,
	}, nil
}

func (k msgServer) UpdateICARemoteZone(goCtx context.Context,  msg *types.MsgUpdateICARemoteZone) (*types.MsgUpdateICARemoteZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    var iCARemoteZone = types.ICARemoteZone{
		Creator: msg.Creator,
		Id:      msg.Id,
    	ConnectionId: msg.ConnectionId,
    	RemoteChainId: msg.RemoteChainId,
    	Active: msg.Active,
    	Bech32Prefix: msg.Bech32Prefix,
	}

    // Checks that the element exists
    val, found := k.GetICARemoteZone(ctx, msg.Id)
    if !found {
        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
    }

    // Checks if the msg creator is the same as the current owner
    if msg.Creator != val.Creator {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }

	k.SetICARemoteZone(ctx, iCARemoteZone)

	return &types.MsgUpdateICARemoteZoneResponse{}, nil
}

func (k msgServer) DeleteICARemoteZone(goCtx context.Context,  msg *types.MsgDeleteICARemoteZone) (*types.MsgDeleteICARemoteZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // Checks that the element exists
    val, found := k.GetICARemoteZone(ctx, msg.Id)
    if !found {
        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
    }

    // Checks if the msg creator is the same as the current owner
    if msg.Creator != val.Creator {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }

	k.RemoveICARemoteZone(ctx, msg.Id)

	return &types.MsgDeleteICARemoteZoneResponse{}, nil
}
