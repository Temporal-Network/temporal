package keeper

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	types2 "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/gogo/protobuf/proto"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
)

// TODO: Change the amount of compoundings per block to be a module level param
const numberOfICACompounds = 100

// TODO Create unit test
func (k Keeper) ShouldICACompoundingHappen() (bool, error) {
	return false, nil
}

// TODO Create unit test
func (k Keeper) RunICACompounding(ctx sdk.Context, msg types.MsgSendTestContractMessages) {
	remoteDelegatorAddress := msg.RemoteDelegatorAddress
	remoteValidatorAddress := msg.RemoteValidatorAddress
	remoteGranteeAddress := msg.RemoteGranteeAddress

	amount := sdk.NewCoin("ujunox", sdk.NewInt(100000))

	delegateMsg := types2.MsgDelegate{
		DelegatorAddress: remoteDelegatorAddress,
		ValidatorAddress: remoteValidatorAddress,
		Amount:           amount,
	}

	msgExecute, err := k.NewAuthzMsgExecute(remoteGranteeAddress, []sdk.Msg{&delegateMsg})
	if err != nil {
		k.Logger(ctx).Info(err.Error())
	}

	msgs := []proto.Message{&msgExecute}

	err = k.SubmitTxs(ctx, msgs, "connection-0")
	if err != nil {
		k.Logger(ctx).Info(err.Error())
	}

}

// TODO Create unit test
func (k Keeper) SendICACompoundMsg(ctx sdk.Context, remoteContractCompoundingSetting types.RemoteContractCompoundSettings) error {
	return nil
}

func (k Keeper) NewAuthzMsgExecute(remoteGranteeAddress string, msgs []sdk.Msg) (authz.MsgExec, error) {
	msgsAny := make([]*cdctypes.Any, len(msgs))
	for i, msg := range msgs {
		any, err := cdctypes.NewAnyWithValue(msg)
		if err != nil {
			return authz.MsgExec{}, err
		}

		msgsAny[i] = any
	}

	return authz.MsgExec{Grantee: remoteGranteeAddress, Msgs: msgsAny}, nil
}
