package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type DistrKeeper interface {
	// Methods imported from distr should be defined here
	CalculateDelegationRewards(ctx sdk.Context, val stakingTypes.ValidatorI, del stakingTypes.DelegationI, endingPeriod uint64) (rewards sdk.DecCoins)
	IncrementValidatorPeriod(ctx sdk.Context, val stakingTypes.ValidatorI) uint64
}

type StakingKeeper interface {
	// Methods imported from staking should be defined here
	Delegate(ctx sdk.Context, delAddr sdk.AccAddress, bondAmt math.Int, tokenSrc stakingTypes.BondStatus, validator stakingTypes.Validator, subtractAccount bool) (sdk.Dec, error)
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (stakingTypes.Validator, bool)
	IterateDelegations(
		ctx sdk.Context, delegator sdk.AccAddress,
		fn func(index int64, delegation stakingTypes.DelegationI) (stop bool),
	)
	Validator(sdk.Context, sdk.ValAddress) stakingTypes.ValidatorI
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	// Methods imported from account should be defined here
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	// Methods imported from bank should be defined here
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}
