package keeper

import (
	"cosmossdk.io/math"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	compTypes "github.com/temporal-zone/temporal/x/compound/types"
	"strconv"
)

type StakingCompoundAction struct {
	Delegator        string
	ValidatorAddress string
	Balance          sdk.Coin
}

var VALIDATORS = make(map[string]sdk.ValAddress)
var DELEGATORS = make(map[string]sdk.AccAddress)

// RunCompounding gets all CompoundSettings and attempts to Compound them
func (k Keeper) RunCompounding(ctx sdk.Context) {
	numberOfCompoundsTemp := k.NumberOfCompoundsPerBlock(ctx)
	compSettings := k.GetAllCompoundSetting(ctx)

	for _, compSetting := range compSettings {
		if !k.ShouldCompoundHappen(ctx, compSetting) {
			continue
		}

		compoundHappened, err := k.Compound(ctx, compSetting)
		if err != nil {
			// A compounding error should not halt the network, so it gets logged to error.
			k.Logger(ctx).Error("compound error: " + err.Error())
			continue
		}

		if compoundHappened {
			numberOfCompoundsTemp--
		}

		if numberOfCompoundsTemp <= 0 {
			break
		}
	}

	k.Logger(ctx).Info("Compounds in this block: " + strconv.Itoa(int(k.NumberOfCompoundsPerBlock(ctx)-numberOfCompoundsTemp)))
}

func (k Keeper) Compound(ctx sdk.Context, cs compTypes.CompoundSetting) (bool, error) {
	address, err := GetAccAddress(cs.Delegator)
	if err != nil {
		return false, err
	}

	// Get all active delegations
	delegations, err := k.DelegationTotalRewards(ctx, address)
	if err != nil {
		return false, err
	}

	// Calculate total rewards that can be claimed and delegated
	walletBalance := k.bankKeeper.GetBalance(ctx, address, sdk.DefaultBondDenom)
	amountToCompound := k.TotalCompoundAmount(delegations, walletBalance, cs)
	if amountToCompound.Amount.IsNegative() {
		//TODO: Better error output that logs the delegations and compound settings
		return false, errors.New("amountToCompound is below 0 for: " + cs.Delegator)
	}

	if amountToCompound.Amount.IsZero() {
		return false, nil
	}

	// Calculate each CompoundSettings validators compound amount
	totalCompoundPercent, compoundActions := k.BuildCompoundActions(cs, amountToCompound)
	if len(compoundActions) == 0 {
		return false, nil
	}

	if totalCompoundPercent.GT(sdk.NewInt(100)) {
		return false, errors.New("totalCompoundPercent can't be over 100")
	}

	// Handle any leftover amount if 100% of rewards are to be compounded by adding any leftover amount to their first validator
	compoundActions = k.HandleLeftOverAmount(compoundActions, totalCompoundPercent, amountToCompound)

	// If withdraw address is different change it temporarily
	withdrawAddr := k.distrKeeper.GetDelegatorWithdrawAddr(ctx, address)
	if !withdrawAddr.Equals(address) {
		k.distrKeeper.SetDelegatorWithdrawAddr(ctx, address, address)
	}

	// Claim all staking rewards, there is an edge case where if multiple validators worth of rewards are being
	// compounded to a single validator and the compounding amount is greater than the sum of the staking reward being
	// claimed on the delegate and the wallet balance, a panic will occur as the network will try to delegate more than
	// the account has available to be delegated.
	err = k.ClaimDelegationRewards(ctx, address, delegations)

	// Change withdraw address back to what it was
	if !withdrawAddr.Equals(address) {
		k.distrKeeper.SetDelegatorWithdrawAddr(ctx, address, withdrawAddr)
	}

	// Execute all CompoundActions
	for _, compoundAction := range compoundActions {
		err := Delegate(ctx, k, compoundAction, address)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func GetAccAddress(delegatorAddress string) (sdk.AccAddress, error) {
	accAddr, found := DELEGATORS[delegatorAddress]
	if !found {
		var err error
		accAddr, err = sdk.AccAddressFromBech32(delegatorAddress)
		if err != nil {
			return nil, err
		}

		DELEGATORS[delegatorAddress] = accAddr
	}

	return accAddr, nil
}

func GetValAddress(validatorAddress string) (sdk.ValAddress, error) {
	valAddr, found := VALIDATORS[validatorAddress]
	if !found {
		var err error
		valAddr, err = sdk.ValAddressFromBech32(validatorAddress)
		if err != nil {
			return nil, err
		}

		VALIDATORS[validatorAddress] = valAddr
	}

	return valAddr, nil
}

func (k Keeper) ClaimDelegationRewards(ctx sdk.Context, address sdk.AccAddress, delegations []distrTypes.DelegationDelegatorReward) error {
	for _, delegation := range delegations {
		valAddr, err := GetValAddress(delegation.ValidatorAddress)
		if err != nil {
			return err
		}

		_, err = k.distrKeeper.WithdrawDelegationRewards(ctx, address, valAddr)
		if err != nil {
			return err
		}
	}

	return nil
}

// ShouldCompoundHappen compares the last time a compounding happened
func (k Keeper) ShouldCompoundHappen(ctx sdk.Context, cs compTypes.CompoundSetting) bool {
	previousCompound, found := k.GetPreviousCompound(ctx, cs.Delegator)
	if !found {
		return true
	}

	nextCompoundTime := previousCompound.BlockHeight + cs.Frequency

	return ctx.BlockHeight() >= nextCompoundTime
}

// Delegate is a helper method that delegates
func Delegate(ctx sdk.Context, k Keeper, compoundAction StakingCompoundAction, address sdk.AccAddress) error {
	valAddr, err := GetValAddress(compoundAction.ValidatorAddress)
	if err != nil {
		return err
	}

	validator, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return errors.New("validator not found")
	}

	_, err = k.stakingKeeper.Delegate(ctx, address, compoundAction.Balance.Amount, stakingTypes.Unbonded, validator, true)
	if err != nil {
		return err
	}

	k.RecordCompound(ctx, address.String())

	return nil
}

// HandleLeftOverAmount calculates any leftover amount if totalCompoundPercent is 100%
func (k Keeper) HandleLeftOverAmount(compoundActions []StakingCompoundAction, totalCompoundPercent math.Int, amountToCompound sdk.Coin) []StakingCompoundAction {
	if totalCompoundPercent.Equal(sdk.NewInt(100)) {
		amountToBeCompounded := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(0))
		for _, compoundAmount := range compoundActions {
			amountToBeCompounded = amountToBeCompounded.Add(compoundAmount.Balance)
		}

		leftOver := amountToCompound.Amount.Sub(amountToBeCompounded.Amount)
		if leftOver.GT(sdk.NewInt(0)) {
			compoundActions[0].Balance.Amount = compoundActions[0].Balance.Amount.Add(leftOver)
		}
	}

	return compoundActions
}

// BuildCompoundActions creates the delegation actions that need to happen
func (k Keeper) BuildCompoundActions(cs compTypes.CompoundSetting, amountToCompound sdk.Coin) (math.Int, []StakingCompoundAction) {
	totalCompoundPercent := sdk.NewInt(0)
	compoundActions := make([]StakingCompoundAction, 0, 1)
	for _, valSetting := range cs.ValidatorSetting {
		validatorCompoundAmount := k.CalculateCompoundAmount(amountToCompound, valSetting.PercentToCompound)

		stakingCompoundAction := StakingCompoundAction{
			ValidatorAddress: valSetting.ValidatorAddress,
			Balance:          sdk.NewCoin(amountToCompound.Denom, validatorCompoundAmount),
		}

		compoundActions = append(compoundActions, stakingCompoundAction)

		totalCompoundPercent = totalCompoundPercent.Add(math.NewInt(int64(valSetting.PercentToCompound)))
	}

	return totalCompoundPercent, compoundActions
}

// TotalCompoundAmount sums all delegations and extra balance amount
func (k Keeper) TotalCompoundAmount(delegations []distrTypes.DelegationDelegatorReward, walletBalance sdk.Coin, cs compTypes.CompoundSetting) sdk.Coin {
	// Sum the total staking claims
	outstandingRewards := k.StakingCompoundAmount(delegations)

	// Extra balance above CompoundSettings.AmountToRemain
	extraCompoundAmount := k.ExtraCompoundAmount(cs, walletBalance)

	return outstandingRewards.Add(extraCompoundAmount)
}

func (k Keeper) StakingCompoundAmount(delegations []distrTypes.DelegationDelegatorReward) sdk.Coin {
	outstandingRewards := sdk.Coin{Denom: sdk.DefaultBondDenom, Amount: sdk.NewInt(0)}
	for _, delegation := range delegations {
		for _, reward := range delegation.Reward {
			if reward.Denom == sdk.DefaultBondDenom {
				outstandingRewards = outstandingRewards.AddAmount(reward.Amount.TruncateInt())
			}
		}
	}

	return outstandingRewards
}

// ExtraCompoundAmount calcs the diff between CompoundSettings.AmountToRemain and the wallet balance
func (k Keeper) ExtraCompoundAmount(cs compTypes.CompoundSetting, walletBalance sdk.Coin) sdk.Coin {
	extraCompoundAmount := sdk.Coin{Denom: sdk.DefaultBondDenom, Amount: sdk.NewInt(0)}

	if !cs.AmountToRemain.IsValid() {
		return extraCompoundAmount
	}

	if walletBalance.Denom == cs.AmountToRemain.Denom && walletBalance.Amount.GT(cs.AmountToRemain.Amount) {
		extraCompoundAmount = walletBalance.Sub(cs.AmountToRemain)
	}

	return extraCompoundAmount
}

// RecordCompound records the compounding timestamp
func (k Keeper) RecordCompound(ctx sdk.Context, address string) {
	value, _ := k.GetPreviousCompound(ctx, address)

	value.Delegator = address
	value.BlockHeight = ctx.BlockHeight()

	k.SetPreviousCompound(ctx, value)
}

// CalculateCompoundAmount calcs the compounding amount
func (k Keeper) CalculateCompoundAmount(rewardAmount sdk.Coin, percentToCompound uint64) math.Int {
	amountToCompound := rewardAmount.Amount.Mul(math.NewInt(int64(percentToCompound))).Quo(sdk.NewInt(100))

	return amountToCompound
}

// DelegationTotalRewards the total rewards accrued by each validator
func (k Keeper) DelegationTotalRewards(ctx sdk.Context, delegator sdk.AccAddress) ([]distrTypes.DelegationDelegatorReward, error) {
	total := sdk.DecCoins{}
	var delRewards []distrTypes.DelegationDelegatorReward

	k.stakingKeeper.IterateDelegations(
		ctx, delegator,
		func(_ int64, del stakingTypes.DelegationI) (stop bool) {
			valAddr := del.GetValidatorAddr()
			val := k.stakingKeeper.Validator(ctx, valAddr)
			endingPeriod := k.distrKeeper.IncrementValidatorPeriod(ctx, val)
			delReward := k.distrKeeper.CalculateDelegationRewards(ctx, val, del, endingPeriod)

			delRewards = append(delRewards, distrTypes.NewDelegationDelegatorReward(valAddr, delReward))
			total = total.Add(delReward...)
			return false
		},
	)

	return delRewards, nil
}
