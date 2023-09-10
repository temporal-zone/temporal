package keeper

import (
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/record/types"
	"time"
)

// SetDelegationHistory set a specific delegationHistory in the store from its index
func (k Keeper) SetDelegationHistory(ctx sdk.Context, delegationHistory types.DelegationHistory) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationHistoryKeyPrefix))
	b := k.cdc.MustMarshal(&delegationHistory)
	store.Set(types.DelegationHistoryKey(delegationHistory.Address), b)
}

// GetDelegationHistory returns a delegationHistory from its index
func (k Keeper) GetDelegationHistory(ctx sdk.Context, address string) (val types.DelegationHistory, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationHistoryKeyPrefix))

	b := store.Get(types.DelegationHistoryKey(address))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDelegationHistory removes a delegationHistory from the store
func (k Keeper) RemoveDelegationHistory(ctx sdk.Context, address string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationHistoryKeyPrefix))
	store.Delete(types.DelegationHistoryKey(address))
}

// GetAllDelegationHistory returns all delegationHistory
func (k Keeper) GetAllDelegationHistory(ctx sdk.Context) (list []types.DelegationHistory) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationHistoryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DelegationHistory
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// CheckDelegationHistoryRecords checks if the accounts DelegationHistory record needs to be updated
func (k Keeper) CheckDelegationHistoryRecords(ctx sdk.Context, delAddr sdk.AccAddress) {
	//current staked amount
	delegatedAmount := k.CalcTotalDelegatedAmount(ctx, delAddr)

	//if no DelegationHistory exists, create a new one
	delegationHistory, found := k.GetDelegationHistory(ctx, delAddr.String())
	if !found {
		delegationHistory = k.NewDelegationHistory(ctx, delAddr, delegatedAmount)
		k.SetDelegationHistory(ctx, delegationHistory)

		return
	}

	//calculates the difference between the sum of DelegationTimestamps balances in the DelegationHistory and the current amount staked
	difference := k.CalcDelegationHistoryDifference(delegatedAmount, delegationHistory)

	if difference.IsZero() {
		return
	}

	//on a positive difference only add to a DelegationTimestamp or a new one to a DelegationHistory
	if difference.IsPositive() {
		delegationHistory = k.AddDelegationTimestamp(ctx, delegationHistory, difference)
		k.SetDelegationHistory(ctx, delegationHistory)

		return
	}

	//on a negative difference it might adjust and/or remove a DelegationTimestamp in a DelegationHistory
	if difference.IsNegative() {
		delegationHistory = k.RemoveDelegationTimestamps(delegationHistory, difference)
		k.SetDelegationHistory(ctx, delegationHistory)

		return
	}
}

// CalcTotalDelegatedAmount gets all delegations shares, sums them and converts to bond denom
func (k Keeper) CalcTotalDelegatedAmount(ctx sdk.Context, delAddr sdk.AccAddress) math.Int {
	delegatedAmount := sdk.NewDec(0)
	delegations := k.stakingKeeper.GetAllDelegatorDelegations(ctx, delAddr)

	for _, delegation := range delegations {
		val, found := k.stakingKeeper.GetValidator(ctx, delegation.GetValidatorAddr())
		if !found {
			panic("unable to find validator: " + delegation.GetValidatorAddr().String())
		}

		delegatedAmount = delegatedAmount.Add(val.TokensFromShares(delegation.Shares))
	}

	return delegatedAmount.TruncateInt()
}

// CalcDelegationHistoryDifference calculates the difference between the sum of bonded amounts in a DelegationHistory and the delegation amount
func (k Keeper) CalcDelegationHistoryDifference(delegationAmount math.Int, delegationHistory types.DelegationHistory) math.Int {
	delegatedAmountHistory := sdk.NewInt(0)
	for _, delegationHistory := range delegationHistory.GetHistory() {
		delegatedAmountHistory = delegatedAmountHistory.Add(delegationHistory.GetBalance().Amount)
	}

	return delegationAmount.Sub(delegatedAmountHistory)
}

// AddDelegationTimestamp adds or adjusts a DelegationTimestamp on an existing DelegationHistory record
func (k Keeper) AddDelegationTimestamp(ctx sdk.Context, delegationHistory types.DelegationHistory, amount math.Int) types.DelegationHistory {
	addedToExisting := false
	newDelTimestamp := k.NewDelegationTimestamp(ctx, amount)
	for _, delTimestamp := range delegationHistory.GetHistory() {
		if delTimestamp.GetTimestamp().Equal(newDelTimestamp.GetTimestamp()) {
			delTimestamp.Balance = delTimestamp.Balance.Add(newDelTimestamp.GetBalance())
			addedToExisting = true
		}
	}

	if !addedToExisting {
		delegationHistory.History = append(delegationHistory.GetHistory(), &newDelTimestamp)
	}

	return delegationHistory
}

// RemoveDelegationTimestamps removes or reduces DelegationsTimestamp(s) on a DelegationHistory
func (k Keeper) RemoveDelegationTimestamps(delegationHistory types.DelegationHistory, difference math.Int) types.DelegationHistory {
	absoluteDifference := difference.Abs()

	for i := len(delegationHistory.GetHistory()) - 1; i >= 0; i-- {
		delegationTimestamp := delegationHistory.GetHistory()[i]

		if delegationTimestamp.GetBalance().Amount.Equal(absoluteDifference) {
			delegationHistory.History = delegationHistory.History[:len(delegationHistory.History)-1]
			break
		} else if delegationTimestamp.GetBalance().Amount.GT(absoluteDifference) {
			delegationHistory.History[i].Balance.Amount = delegationTimestamp.GetBalance().Amount.Sub(absoluteDifference)
			break
		} else {
			absoluteDifference = absoluteDifference.Sub(delegationTimestamp.GetBalance().Amount)
			delegationHistory.History = delegationHistory.History[:len(delegationHistory.History)-1]
		}
	}

	return delegationHistory
}

// PruneDelegationHistory prune the DelegationTimestamp history in a DelegationHistory
func (k Keeper) PruneDelegationHistory(delegationHistory types.DelegationHistory) types.DelegationHistory {
	delegationHistoryNew := types.DelegationHistory{}
	delegationHistoryNew.Address = delegationHistory.Address

	//remove any DelegationTimestamp that have a 0 amount amd compress DelegationHistory's down to a daily frequency
	for _, delegationTimestamp := range delegationHistory.GetHistory() {
		if !delegationTimestamp.GetBalance().Amount.Equal(math.NewInt(0)) {
			delegationHistoryNew.History = append(delegationHistoryNew.GetHistory(), delegationTimestamp)
		}
	}

	return delegationHistoryNew
}

// NewDelegationTimestamp creates a new DelegationTimestamp
func (k Keeper) NewDelegationTimestamp(ctx sdk.Context, amount math.Int) types.DelegationTimestamp {
	bt := time.Unix(ctx.BlockTime().Unix(), 0).UTC()
	dt := time.Date(bt.Year(), bt.Month(), bt.Day(), 0, 0, 0, 0, bt.Location())
	return types.DelegationTimestamp{
		Timestamp: dt,
		Balance: sdk.NewCoin(
			k.stakingKeeper.BondDenom(ctx),
			amount,
		),
	}
}

// NewDelegationHistory creates a new DelegationHistory
func (k Keeper) NewDelegationHistory(ctx sdk.Context, delAddr sdk.AccAddress, delegatedAmount math.Int) types.DelegationHistory {
	delegationTimestamp := k.NewDelegationTimestamp(ctx, delegatedAmount)

	delegationHistory := types.DelegationHistory{
		Address: delAddr.String(),
		History: []*types.DelegationTimestamp{&delegationTimestamp},
	}

	return delegationHistory
}
