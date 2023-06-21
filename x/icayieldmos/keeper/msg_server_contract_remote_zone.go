package keeper

import (
	"context"
	"fmt"
	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) isContractRemoteZoneValid(ctx sdk.Context, connectionId string, remoteChainId string, bech32Prefix string) error {
	owner := k.accountKeeper.GetModuleAddress(types.ModuleName)

	if !k.DoesInterchainAccountExist(ctx, connectionId, owner.String()) {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, fmt.Sprintf("ica account not found for connectionId: %s and owner: %s", connectionId, owner.String()))
	}

	if !k.IsChannelOpen(ctx, connectionId, owner.String()) {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, fmt.Sprintf("channel not open for connectionId: %s and owner: %s", connectionId, owner.String()))
	}

	_, found := k.GetRemoteModuleAddress(ctx, connectionId, owner.String())
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, fmt.Sprintf("remote address not found for connectionId: %s and owner: %s", connectionId, owner.String()))
	}

	if len(bech32Prefix) == 0 || bech32Prefix == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("bech32Prefix can't be empty"))
	}

	remoteZones := k.GetAllContractRemoteZone(ctx)

	for _, remoteZone := range remoteZones {
		if connectionId == remoteZone.ConnectionId && remoteChainId == remoteZone.RemoteChainId {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("duplicate remote zone not allowed for connectionID: %s and remoteChainId: %s", connectionId, remoteChainId))
		}
	}

	return nil
}

// TODO Restrict who can create ContracRemoteZones
func (k msgServer) CreateContractRemoteZone(goCtx context.Context, msg *types.MsgCreateContractRemoteZone) (*types.MsgCreateContractRemoteZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.isContractRemoteZoneValid(ctx, msg.ConnectionId, msg.RemoteChainId, msg.Bech32Prefix)
	if err != nil {
		return nil, err
	}

	var contractRemoteZone = types.ContractRemoteZone{
		Creator:       msg.Creator,
		ConnectionId:  msg.ConnectionId,
		RemoteChainId: msg.RemoteChainId,
		Active:        msg.Active,
		Bech32Prefix:  msg.Bech32Prefix,
	}

	id := k.AppendContractRemoteZone(ctx, contractRemoteZone)

	return &types.MsgCreateContractRemoteZoneResponse{Id: id}, nil
}

// TODO Restrict who can update ContracRemoteZones
func (k msgServer) UpdateContractRemoteZone(goCtx context.Context, msg *types.MsgUpdateContractRemoteZone) (*types.MsgUpdateContractRemoteZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.isContractRemoteZoneValid(ctx, msg.ConnectionId, msg.RemoteChainId, msg.Bech32Prefix)
	if err != nil {
		return nil, err
	}

	var contractRemoteZone = types.ContractRemoteZone{
		Creator:       msg.Creator,
		Id:            msg.Id,
		ConnectionId:  msg.ConnectionId,
		RemoteChainId: msg.RemoteChainId,
		Active:        msg.Active,
		Bech32Prefix:  msg.Bech32Prefix,
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

// TODO Restrict who can delete ContracRemoteZones
func (k msgServer) DeleteContractRemoteZone(goCtx context.Context, msg *types.MsgDeleteContractRemoteZone) (*types.MsgDeleteContractRemoteZoneResponse, error) {
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
