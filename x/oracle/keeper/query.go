package keeper

import (
	"core/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
