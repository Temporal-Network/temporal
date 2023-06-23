package keeper

import (
	"github.com/temporal-zone/temporal/x/yieldmos/types"
)

var _ types.QueryServer = Keeper{}
