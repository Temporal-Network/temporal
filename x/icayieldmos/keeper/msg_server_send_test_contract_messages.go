package keeper

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/proto"

	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendTestContractMessages(goCtx context.Context, msg *types.MsgSendTestContractMessages) (*types.MsgSendTestContractMessagesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	json := `{"msg":{"compound":{"comp_prefs":{"relative":[{"destination":{"juno_staking":{"validator_address":"junovaloper1m55p4c956dawa95uhzz027p4pwm3fedfkdtnyj"}},"amount":"1000000000000000000"}]},"delegator_address":"juno1f49xq0rmah39sk58aaxq6gnqcvupee7jgl90tn"}},"funds":[]}"`
	contractMsg := []byte(json)

	msgExecuteContract := types.MsgExecuteContract{
		Sender:   k.accountKeeper.GetModuleAddress(types.ModuleName).String(),
		Contract: "juno1hy58yc42xhwq28qhct4xv0yr9gmndcwzdkfh035pkqng0rxw80qqlpxelj",
		Msg:      contractMsg,
		Funds:    nil,
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
