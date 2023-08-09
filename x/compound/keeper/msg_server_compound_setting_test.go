package keeper_test

import (
	"github.com/temporal-zone/temporal/app/apptesting"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/temporal-zone/temporal/testutil/keeper"
	"github.com/temporal-zone/temporal/x/compound/keeper"
	"github.com/temporal-zone/temporal/x/compound/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCompoundSettingMsgServerCreate(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()
	srv := keeper.NewMsgServerImpl(s.App.CompoundKeeper)

	vals := s.App.StakingKeeper.GetAllValidators(s.Ctx)
	require.True(t, len(vals) > 0)

	val := vals[0]
	valAddress, err := sdk.ValAddressFromBech32(val.OperatorAddress)
	require.NoError(t, err)

	valSetting := types.ValidatorSetting{
		ValidatorAddress:  valAddress.String(),
		PercentToCompound: 100,
	}

	valSettings := []*types.ValidatorSetting{&valSetting}

	for i := 0; i < 5; i++ {
		delegator := strconv.Itoa(i)
		expected := &types.MsgCreateCompoundSetting{Delegator: delegator, ValidatorSetting: valSettings}
		_, err := srv.CreateCompoundSetting(s.Ctx, expected)
		require.NoError(t, err)
		rst, found := s.App.CompoundKeeper.GetCompoundSetting(s.Ctx, expected.Delegator)
		require.True(t, found)
		require.Equal(t, expected.Delegator, rst.Delegator)
	}
}

func TestCompoundSettingMsgServerUpdate(t *testing.T) {
	delegatorA := "A"
	delegatorB := "B"
	delegatorC := "C"
	delegatorD := "D"

	s := apptesting.SetupSuitelessTestHelper()
	srv := keeper.NewMsgServerImpl(s.App.CompoundKeeper)

	vals := s.App.StakingKeeper.GetAllValidators(s.Ctx)
	require.True(t, len(vals) > 0)

	val := vals[0]
	valAddress, err := sdk.ValAddressFromBech32(val.OperatorAddress)
	require.NoError(t, err)

	valSettingUpdate := types.ValidatorSetting{
		ValidatorAddress:  valAddress.String(),
		PercentToCompound: 50,
	}

	valSettingsUpdate := []*types.ValidatorSetting{&valSettingUpdate}

	valSettingInvalidOver := types.ValidatorSetting{
		ValidatorAddress:  valAddress.String(),
		PercentToCompound: 150,
	}

	valSettingsInvalidOver := []*types.ValidatorSetting{&valSettingInvalidOver}

	valSettingInvalidUnder := types.ValidatorSetting{
		ValidatorAddress:  valAddress.String(),
		PercentToCompound: 0,
	}

	valSettingsInvalidUnder := []*types.ValidatorSetting{&valSettingInvalidUnder}

	valSettingInvalidTotalOver1 := types.ValidatorSetting{
		ValidatorAddress:  valAddress.String(),
		PercentToCompound: 75,
	}

	valSettingInvalidTotalOver2 := types.ValidatorSetting{
		ValidatorAddress:  valAddress.String(),
		PercentToCompound: 75,
	}

	valSettingsInvalidTotalOver := []*types.ValidatorSetting{&valSettingInvalidTotalOver1, &valSettingInvalidTotalOver2}

	tests := []struct {
		desc      string
		delegator string
		request   *types.MsgUpdateCompoundSetting
		err       error
	}{
		{
			desc:      "Completed",
			delegator: delegatorA,
			request:   &types.MsgUpdateCompoundSetting{Delegator: delegatorA, ValidatorSetting: valSettingsUpdate},
		},
		{
			desc:      "KeyNotFound",
			delegator: delegatorB,
			request:   &types.MsgUpdateCompoundSetting{Delegator: delegatorB, ValidatorSetting: valSettingsUpdate},
			err:       sdkerrors.ErrKeyNotFound,
		},
		{
			desc:      "InvalidValidatorSettingsOver",
			delegator: delegatorB,
			request:   &types.MsgUpdateCompoundSetting{Delegator: delegatorB, ValidatorSetting: valSettingsInvalidOver},
			err:       sdkerrors.ErrInvalidRequest,
		},
		{
			desc:      "InvalidValidatorSettingsUnder",
			delegator: delegatorC,
			request:   &types.MsgUpdateCompoundSetting{Delegator: delegatorC, ValidatorSetting: valSettingsInvalidUnder},
			err:       sdkerrors.ErrInvalidRequest,
		},
		{
			desc:      "InvalidValidatorSettingsTotalOver",
			delegator: delegatorD,
			request:   &types.MsgUpdateCompoundSetting{Delegator: delegatorD, ValidatorSetting: valSettingsInvalidTotalOver},
			err:       sdkerrors.ErrInvalidRequest,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			valSetting := types.ValidatorSetting{
				ValidatorAddress:  valAddress.String(),
				PercentToCompound: 75,
			}

			valSettings := []*types.ValidatorSetting{&valSetting}

			expected := &types.MsgCreateCompoundSetting{Delegator: tc.delegator, ValidatorSetting: valSettings}
			_, err := srv.CreateCompoundSetting(s.Ctx, expected)
			require.NoError(t, err)

			if tc.desc == "KeyNotFound" {
				del := &types.MsgDeleteCompoundSetting{Delegator: tc.delegator}
				_, err = srv.DeleteCompoundSetting(s.Ctx, del)
				require.NoError(t, err)
			}

			_, err = srv.UpdateCompoundSetting(s.Ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := s.App.CompoundKeeper.GetCompoundSetting(s.Ctx, expected.Delegator)
				require.True(t, found)
				require.Equal(t, expected.Delegator, rst.Delegator)
			}
		})
	}
}

func TestCompoundSettingMsgServerDelete(t *testing.T) {
	delegatorA := "A"
	delegatorB := "B"

	s := apptesting.SetupSuitelessTestHelper()
	srv := keeper.NewMsgServerImpl(s.App.CompoundKeeper)

	vals := s.App.StakingKeeper.GetAllValidators(s.Ctx)
	require.True(t, len(vals) > 0)

	val := vals[0]
	valAddress, err := sdk.ValAddressFromBech32(val.OperatorAddress)
	require.NoError(t, err)

	valSetting := types.ValidatorSetting{
		ValidatorAddress:  valAddress.String(),
		PercentToCompound: 50,
	}

	valSettings := []*types.ValidatorSetting{&valSetting}

	tests := []struct {
		desc      string
		delegator string
		request   *types.MsgDeleteCompoundSetting
		err       error
	}{
		{
			desc:      "Completed",
			delegator: delegatorA,
			request:   &types.MsgDeleteCompoundSetting{Delegator: delegatorA},
		},
		{
			desc:      "KeyNotFound",
			delegator: delegatorB,
			request:   &types.MsgDeleteCompoundSetting{Delegator: delegatorB},
			err:       sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := srv.CreateCompoundSetting(s.Ctx, &types.MsgCreateCompoundSetting{Delegator: tc.delegator, ValidatorSetting: valSettings})
			require.NoError(t, err)

			if tc.desc == "KeyNotFound" {
				del := &types.MsgDeleteCompoundSetting{Delegator: tc.delegator}
				_, err := srv.DeleteCompoundSetting(s.Ctx, del)
				require.NoError(t, err)
			}

			_, err = srv.DeleteCompoundSetting(s.Ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := s.App.CompoundKeeper.GetCompoundSetting(s.Ctx, tc.request.Delegator)
				require.False(t, found)
			}
		})
	}
}

func TestCheckFrequency(t *testing.T) {
	tests := []struct {
		desc    string
		request uint64
		err     error
	}{
		{
			desc:    "Under",
			request: 10,
		},
		{
			desc:    "Over",
			request: 700,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CompoundKeeper(t)

			onceEvery := k.CheckFrequency(ctx, tc.request)

			if tc.request < k.MinimumCompoundFrequency(ctx) {
				require.GreaterOrEqual(t, onceEvery, k.MinimumCompoundFrequency(ctx))
			}

			if tc.request > k.MinimumCompoundFrequency(ctx) {
				require.Equal(t, onceEvery, tc.request)
			}
		})
	}
}
