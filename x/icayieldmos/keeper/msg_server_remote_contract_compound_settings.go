package keeper

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
)

func (k msgServer) isRemoteContractCompoundSettingsValid(
	ctx sdk.Context,
	contractRemoteZone uint64,
	remoteDelegatorAddress string,
	compoundSettings string,
	frequency uint64,
	remoteContractAddress string,
	creator string) error {

	contractRemoteZoneObj, found := k.GetContractRemoteZone(ctx, contractRemoteZone)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, fmt.Sprintf("unable to find ContractRemoteZone with id: %d", contractRemoteZone))
	}

	if len(remoteDelegatorAddress) == 0 || remoteDelegatorAddress == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid remoteDelegatorAddress: %s", remoteDelegatorAddress))
	}

	// TODO: This validation will only work on 44/118 deriv path addresses, when adding a non-44/118 network add a different check...
	_, base64, err := bech32.DecodeAndConvert(creator)
	if err != nil {
		return err
	}

	remoteAddressConverted, err := bech32.ConvertAndEncode(contractRemoteZoneObj.Bech32Prefix, base64)
	if err != nil {
		return err
	}

	if remoteDelegatorAddress != remoteAddressConverted {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("remoteDelegatorAddress address: %s doesn't remoteAddressConverted: %s", remoteDelegatorAddress, remoteAddressConverted))
	}

	if len(compoundSettings) == 0 || compoundSettings == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, fmt.Sprintf("RemoteCompoundSettings can't be empty"))
	}

	if frequency < 100 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, fmt.Sprintf("FrequencyInSeconds can't be less than 100"))
	}

	if len(remoteContractAddress) == 0 || remoteContractAddress == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, fmt.Sprintf("RemoteContractAddress can't be empty"))
	}

	return nil
}

func (k msgServer) CreateRemoteContractCompoundSettings(goCtx context.Context, msg *types.MsgCreateRemoteContractCompoundSettings) (*types.MsgCreateRemoteContractCompoundSettingsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.isRemoteContractCompoundSettingsValid(
		ctx,
		msg.ContractRemoteZone,
		msg.RemoteDelegatorAddress,
		msg.CompoundSettings,
		msg.FrequencyInSeconds,
		msg.RemoteContractAddress,
		msg.Creator)
	if err != nil {
		return nil, err
	}

	var remoteContractCompoundSettings = types.RemoteContractCompoundSettings{
		Creator:                msg.Creator,
		ContractRemoteZone:     msg.ContractRemoteZone,
		RemoteDelegatorAddress: msg.RemoteDelegatorAddress,
		CompoundSettings:       msg.CompoundSettings,
		FrequencyInSeconds:     msg.FrequencyInSeconds,
		RemoteContractAddress:  msg.RemoteContractAddress,
	}

	id := k.AppendRemoteContractCompoundSettings(
		ctx,
		remoteContractCompoundSettings,
	)

	return &types.MsgCreateRemoteContractCompoundSettingsResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateRemoteContractCompoundSettings(goCtx context.Context, msg *types.MsgUpdateRemoteContractCompoundSettings) (*types.MsgUpdateRemoteContractCompoundSettingsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.isRemoteContractCompoundSettingsValid(
		ctx,
		msg.ContractRemoteZone,
		msg.RemoteDelegatorAddress,
		msg.CompoundSettings,
		msg.FrequencyInSeconds,
		msg.RemoteContractAddress,
		msg.Creator)
	if err != nil {
		return nil, err
	}

	var remoteContractCompoundSettings = types.RemoteContractCompoundSettings{
		Creator:                msg.Creator,
		Id:                     msg.Id,
		ContractRemoteZone:     msg.ContractRemoteZone,
		RemoteDelegatorAddress: msg.RemoteDelegatorAddress,
		CompoundSettings:       msg.CompoundSettings,
		FrequencyInSeconds:     msg.FrequencyInSeconds,
		RemoteContractAddress:  msg.RemoteContractAddress,
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

func (k msgServer) DeleteRemoteContractCompoundSettings(goCtx context.Context, msg *types.MsgDeleteRemoteContractCompoundSettings) (*types.MsgDeleteRemoteContractCompoundSettingsResponse, error) {
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
