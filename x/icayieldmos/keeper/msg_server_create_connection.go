package keeper

import (
	"context"

	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateConnection(goCtx context.Context, msg *types.MsgCreateConnection) (*types.MsgCreateConnectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateConnectionResponse{}, nil
}
