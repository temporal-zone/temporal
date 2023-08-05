package keeper

import (
	"github.com/temporal-zone/temporal/x/record/types"
)

var _ types.QueryServer = Keeper{}
