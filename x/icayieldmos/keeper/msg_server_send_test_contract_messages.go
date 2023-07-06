package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
)

func (k msgServer) SendTestContractMessages(goCtx context.Context, msg *types.MsgSendTestContractMessages) (*types.MsgSendTestContractMessagesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.QueryDelegations(ctx)

	//k.RunICACompounding(ctx, *msg)

	return &types.MsgSendTestContractMessagesResponse{}, nil
}

/*
func (k msgServer) SendTestContractMessages(goCtx context.Context, msg *types.MsgSendTestContractMessages) (*types.MsgSendTestContractMessagesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	json := `{"compound":{"comp_prefs":{"relative":[{"destination":{"juno_staking":{"validator_address":"junovaloper1cvk3scc8akj4psfzre7xxyt9453ug6x7ul5gw3"}},"amount":"1000000000000000000"}]},"delegator_address":"juno17dtl0mjt3t77kpuhg2edqzjpszulwhgz2qxy4v"}}`
	contractMsg := types.RawContractMessage(json)

	owner := k.GetModuleAddress()

	remoteAddress, found := k.GetRemoteModuleAddress(ctx, msg.ConnectionId, owner.String())
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "remote ICA address not found for connectionId: %s and owner: %s", msg.ConnectionId, owner)
	}

	coins := sdk.NewCoins()

	msgExecuteContract := types.MsgExecuteContract{
		Sender:   remoteAddress,
		Contract: "juno1unyuj8qnmygvzuex3dwmg9yzt9alhvyeat0uu0jedg2wj33efl5qewxkxd",
		Msg:      contractMsg,
		Funds:    coins,
	}

	msgs := []proto.Message{&msgExecuteContract}

	fmt.Println(msgExecuteContract.String())

	err := k.SubmitTxs(ctx, msgs, msg.ConnectionId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &types.MsgSendTestContractMessagesResponse{}, nil
}
*/
