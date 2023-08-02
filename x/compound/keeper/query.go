package keeper

import (
	"temporal/x/compound/types"
)

var _ types.QueryServer = Keeper{}
