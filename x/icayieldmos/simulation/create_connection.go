package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/temporal-zone/temporal/x/icayieldmos/keeper"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
)

func SimulateMsgCreateConnection(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateConnection{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateConnection simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateConnection simulation not implemented"), nil, nil
	}
}
