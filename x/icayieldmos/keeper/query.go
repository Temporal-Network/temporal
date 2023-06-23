package keeper

import (
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
)

var _ types.QueryServer = Keeper{}
