package keeper

import (
	"core/x/reserve/types"
)

var _ types.QueryServer = Keeper{}
