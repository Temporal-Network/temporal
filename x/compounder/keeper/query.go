package keeper

import (
	"github.com/temporal-zone/temporal/x/compounder/types"
)

var _ types.QueryServer = Keeper{}
