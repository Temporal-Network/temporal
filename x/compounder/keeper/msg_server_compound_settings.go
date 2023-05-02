package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"temporal/x/compounder/types"
)

func (k msgServer) CreateCompoundSettings(goCtx context.Context, msg *types.MsgCreateCompoundSettings) (*types.MsgCreateCompoundSettingsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetCompoundSettings(
		ctx,
		msg.Delegator,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "delegator already has a CompoundSettings defined, use MsgUpdateCompoundSettings instead")
	}

	var compoundSettings = types.CompoundSettings{
		Delegator:         msg.Delegator,
		ValidatorSettings: msg.ValidatorSettings,
		AmountToRemain:    msg.AmountToRemain,
		Frequency:         msg.Frequency,
	}

	k.SetCompoundSettings(
		ctx,
		compoundSettings,
	)
	return &types.MsgCreateCompoundSettingsResponse{}, nil
}

func (k msgServer) UpdateCompoundSettings(goCtx context.Context, msg *types.MsgUpdateCompoundSettings) (*types.MsgUpdateCompoundSettingsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetCompoundSettings(
		ctx,
		msg.Delegator,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "can't find CompoundSettings for Delegator")
	}

	// Checks if the the msg delegator is the same as the current owner
	if msg.Delegator != valFound.Delegator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var compoundSettings = types.CompoundSettings{
		Delegator:         msg.Delegator,
		ValidatorSettings: msg.ValidatorSettings,
		AmountToRemain:    msg.AmountToRemain,
		Frequency:         msg.Frequency,
	}

	k.SetCompoundSettings(ctx, compoundSettings)

	return &types.MsgUpdateCompoundSettingsResponse{}, nil
}

func (k msgServer) DeleteCompoundSettings(goCtx context.Context, msg *types.MsgDeleteCompoundSettings) (*types.MsgDeleteCompoundSettingsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetCompoundSettings(
		ctx,
		msg.Delegator,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg delegator is the same as the current owner
	if msg.Delegator != valFound.Delegator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveCompoundSettings(
		ctx,
		msg.Delegator,
	)

	return &types.MsgDeleteCompoundSettingsResponse{}, nil
}
