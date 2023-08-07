package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type StakingKeeper interface {
	// Methods imported from staking should be defined here
	GetDelegation(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (stakingTypes.Delegation, bool)
	BondDenom(ctx sdk.Context) string
	GetValidator(ctx sdk.Context, valAddr sdk.ValAddress) (stakingTypes.Validator, bool)
	GetAllDelegatorDelegations(ctx sdk.Context, delAddr sdk.AccAddress) []stakingTypes.Delegation
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}
