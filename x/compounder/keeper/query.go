package keeper

import (
	"github.com/Temporal-Network/temporal/x/compounder/types"
)

var _ types.QueryServer = Keeper{}
