package keeper

import (
	"context"

	"github.com/Temporal-Network/temporal/x/compounder/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	if msg.Frequency != nil {
		msg.Frequency.OnceEvery = CheckFrequency(msg.Frequency.OnceEvery)
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

	if msg.Frequency != nil {
		msg.Frequency.OnceEvery = CheckFrequency(msg.Frequency.OnceEvery)
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

// TODO: CheckFrequency needs test coverage
// CheckFrequency checks to make sure frequency should be no less than X seconds or blocks.
func CheckFrequency(onceEvery sdk.Int) sdk.Int {
	// TODO: Change minimumCompoundingFrequency to be a module param
	minimumCompoundingFrequency := sdk.NewInt(100)
	if onceEvery.LT(minimumCompoundingFrequency) {
		return minimumCompoundingFrequency
	}

	return onceEvery
}
