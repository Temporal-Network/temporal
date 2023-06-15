package keeper

import (
    "fmt"
	"context"

    "github.com/Temporal-Network/temporal/x/icayieldmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


func (k msgServer) CreateContractRemoteZone(goCtx context.Context,  msg *types.MsgCreateContractRemoteZone) (*types.MsgCreateContractRemoteZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    var contractRemoteZone = types.ContractRemoteZone{
        Creator: msg.Creator,
        ConnectionId: msg.ConnectionId,
        RemoteChainId: msg.RemoteChainId,
        Active: msg.Active,
    }

    id := k.AppendContractRemoteZone(
        ctx,
        contractRemoteZone,
    )

	return &types.MsgCreateContractRemoteZoneResponse{
	    Id: id,
	}, nil
}

func (k msgServer) UpdateContractRemoteZone(goCtx context.Context,  msg *types.MsgUpdateContractRemoteZone) (*types.MsgUpdateContractRemoteZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    var contractRemoteZone = types.ContractRemoteZone{
		Creator: msg.Creator,
		Id:      msg.Id,
    	ConnectionId: msg.ConnectionId,
    	RemoteChainId: msg.RemoteChainId,
    	Active: msg.Active,
	}

    // Checks that the element exists
    val, found := k.GetContractRemoteZone(ctx, msg.Id)
    if !found {
        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
    }

    // Checks if the msg creator is the same as the current owner
    if msg.Creator != val.Creator {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }

	k.SetContractRemoteZone(ctx, contractRemoteZone)

	return &types.MsgUpdateContractRemoteZoneResponse{}, nil
}

func (k msgServer) DeleteContractRemoteZone(goCtx context.Context,  msg *types.MsgDeleteContractRemoteZone) (*types.MsgDeleteContractRemoteZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // Checks that the element exists
    val, found := k.GetContractRemoteZone(ctx, msg.Id)
    if !found {
        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
    }

    // Checks if the msg creator is the same as the current owner
    if msg.Creator != val.Creator {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }

	k.RemoveContractRemoteZone(ctx, msg.Id)

	return &types.MsgDeleteContractRemoteZoneResponse{}, nil
}
