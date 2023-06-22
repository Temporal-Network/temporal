package keeper

import (
	"fmt"
	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
)

// TODO: Change the amount of compoundings per block to be a module level param
const numberOfCompounds = 100

// TODO Create unit test
func (k Keeper) ShouldCompoundingHappen(ctx sdk.Context, remoteContractCompoundingSetting types.RemoteContractCompoundSettings) (bool, error) {
	previousRemoteCompounding, found := k.GetPreviousRemoteCompounding(ctx, remoteContractCompoundingSetting.Id)
	if !found {
		return true, nil
	}

	secondsSincePreviousCompounding := ctx.BlockTime().Sub(previousRemoteCompounding.Timestamp).Seconds()
	// Just a sanity check to make sure the time since last compound is not negative.
	if secondsSincePreviousCompounding < 0 {
		return false, sdkerrors.Wrap(sdkerrors.ErrInvalidType, fmt.Sprintf("time since last compounding is negative: %f", secondsSincePreviousCompounding))
	}

	if uint64(secondsSincePreviousCompounding) > remoteContractCompoundingSetting.FrequencyInSeconds {
		return true, nil
	}

	return false, nil
}

// TODO Create unit test
func (k Keeper) RunRemoteCompounding(ctx sdk.Context) {
	remoteContractCompoundingSettings := k.GetAllRemoteContractCompoundSettings(ctx)

	// TODO: Profile/Benchmark what 100, 1K, 10K, 100K of compounds in one block does to block time/other resource usage?
	numberOfCompoundsTemp := numberOfCompounds

	for _, remoteContractCompoundingSetting := range remoteContractCompoundingSettings {
		shouldCompoundingHappen, err := k.ShouldCompoundingHappen(ctx, remoteContractCompoundingSetting)
		if err != nil {
			k.Logger(ctx).Info(err.Error())
			continue
		}

		if !shouldCompoundingHappen {
			continue
		}

		numberOfCompoundsTemp--

		err = k.SendCompoundMsg(ctx, remoteContractCompoundingSetting)
		if err != nil {
			k.Logger(ctx).Info(err.Error())
			continue
		}

		if numberOfCompoundsTemp <= 0 {
			return
		}
	}
}

// TODO Create unit test
func (k Keeper) SendCompoundMsg(ctx sdk.Context, remoteContractCompoundingSetting types.RemoteContractCompoundSettings) error {
	contractRemoteZone, found := k.GetContractRemoteZone(ctx, remoteContractCompoundingSetting.ContractRemoteZone)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, fmt.Sprintf("ContractRemoteZone not found: %d", remoteContractCompoundingSetting.ContractRemoteZone))
	}

	contractMsg := types.RawContractMessage(remoteContractCompoundingSetting.CompoundSettings)

	owner := k.GetModuleAddress()

	remoteAddress, found := k.GetRemoteModuleAddress(ctx, contractRemoteZone.ConnectionId, owner.String())
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "remote ICA address not found for connectionId: %s and owner: %s", contractRemoteZone.ConnectionId, owner)
	}

	coins := sdk.NewCoins()

	msgExecuteContract := types.MsgExecuteContract{
		Sender:   remoteAddress,
		Contract: remoteContractCompoundingSetting.RemoteContractAddress,
		Msg:      contractMsg,
		Funds:    coins,
	}

	msgs := []proto.Message{&msgExecuteContract}

	fmt.Println(msgExecuteContract.String())

	err := k.SubmitTxs(ctx, msgs, contractRemoteZone.ConnectionId)
	if err != nil {
		return err
	}

	previousRemoteCompounding := types.PreviousRemoteCompounding{
		RemoteContractCompoundSetting: remoteContractCompoundingSetting.Id,
		Timestamp:                     ctx.BlockTime(),
	}

	k.SetPreviousRemoteCompounding(ctx, previousRemoteCompounding)

	return nil
}
