package types

// DONTCOVER

import (
	sdkerr "cosmossdk.io/errors"
)

// x/record module sentinel errors
var (
	ErrSample = sdkerr.Register(ModuleName, 1100, "sample error")
)
