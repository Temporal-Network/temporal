package keeper

import (
	"temporal/x/compounder/types"
)

var _ types.QueryServer = Keeper{}
