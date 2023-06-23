package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
)

func (k msgServer) CreateConnection(goCtx context.Context, msg *types.MsgCreateConnection) (*types.MsgCreateConnectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.RegisterInterchainAccount(ctx, msg.ConnectionId)
	if err != nil {
		return &types.MsgCreateConnectionResponse{}, err
	}

	return &types.MsgCreateConnectionResponse{}, nil
}
