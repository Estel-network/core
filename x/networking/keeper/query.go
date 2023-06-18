package keeper

import (
	"core/x/networking/types"
)

var _ types.QueryServer = Keeper{}
