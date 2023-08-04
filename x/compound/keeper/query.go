package keeper

import (
	"github.com/temporal-zone/temporal/x/compound/types"
)

var _ types.QueryServer = Keeper{}
