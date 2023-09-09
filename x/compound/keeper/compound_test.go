package keeper_test

import (
	"cosmossdk.io/math"
	sdkmath "cosmossdk.io/math"
	"fmt"
	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingCli "github.com/cosmos/cosmos-sdk/x/staking/client/cli"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	"github.com/temporal-zone/temporal/app/apptesting"
	"github.com/temporal-zone/temporal/x/compound/client/cli"
	"github.com/temporal-zone/temporal/x/compound/keeper"
	compTypes "github.com/temporal-zone/temporal/x/compound/types"
	recordCli "github.com/temporal-zone/temporal/x/record/client/cli"
	recordTypes "github.com/temporal-zone/temporal/x/record/types"
	"strconv"
	"testing"
	"time"

	"github.com/temporal-zone/temporal/testutil/network"
	"github.com/temporal-zone/temporal/x/compound/types"

	compoundCli "github.com/temporal-zone/temporal/x/compound/client/cli"
)

func TestShouldCompoundHappen(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	cases := []struct {
		name        string
		cs          compTypes.CompoundSetting
		blockHeight int64
		expected    bool
		expectedErr error
		prevComp    compTypes.PreviousCompound
	}{
		{
			name: "every 10 blocks true",
			cs: compTypes.CompoundSetting{
				Delegator: "delegator2",
				Frequency: 10,
			},
			blockHeight: s.Ctx.BlockHeight(),
			expected:    true,
			expectedErr: nil,
			prevComp: compTypes.PreviousCompound{
				Delegator:   "delegator1",
				BlockHeight: 0,
			},
		},
		{
			name: "every 10 blocks false",
			cs: compTypes.CompoundSetting{
				Delegator: "delegator2",
				Frequency: 10,
			},
			blockHeight: s.Ctx.BlockHeight(),
			expected:    false,
			expectedErr: nil,
			prevComp: compTypes.PreviousCompound{
				Delegator:   "delegator2",
				BlockHeight: 1,
			},
		},
		{
			name: "no previous comp",
			cs: compTypes.CompoundSetting{
				Delegator: "delegator3",
				Frequency: 10,
			},
			blockHeight: s.Ctx.BlockHeight(),
			expected:    true,
			expectedErr: nil,
			prevComp:    compTypes.PreviousCompound{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.name != "no previous comp" {
				s.App.CompoundKeeper.SetPreviousCompound(s.Ctx, tc.prevComp)
			}

			actual := s.App.CompoundKeeper.ShouldCompoundHappen(s.Ctx, tc.cs)
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestHandleLeftOverAmount(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	cases := []struct {
		name                 string
		compoundActions      []keeper.StakingCompoundAction
		totalCompoundPercent math.Int
		amountToCompound     sdk.Coin
		expectedResult       []keeper.StakingCompoundAction
	}{
		{
			name: "total compound percent is 100% 1",
			compoundActions: []keeper.StakingCompoundAction{
				{ValidatorAddress: "cosmosval1", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(50))},
				{ValidatorAddress: "cosmosval2", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(50))},
			},
			totalCompoundPercent: sdk.NewInt(100),
			amountToCompound:     sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100)),
			expectedResult: []keeper.StakingCompoundAction{
				{ValidatorAddress: "cosmosval1", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(50))},
				{ValidatorAddress: "cosmosval2", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(50))},
			},
		},
		{
			name: "total compound percent is 100% 2",
			compoundActions: []keeper.StakingCompoundAction{
				{ValidatorAddress: "cosmosval1", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(24))},
				{ValidatorAddress: "cosmosval2", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(80))},
			},
			totalCompoundPercent: sdk.NewInt(100),
			amountToCompound:     sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(105)),
			expectedResult: []keeper.StakingCompoundAction{
				{ValidatorAddress: "cosmosval1", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(25))},
				{ValidatorAddress: "cosmosval2", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(80))},
			},
		},

		{
			name: "total compound percent is 100% 2",
			compoundActions: []keeper.StakingCompoundAction{
				{ValidatorAddress: "cosmosval1", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))},
				{ValidatorAddress: "cosmosval2", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(20))},
			},
			totalCompoundPercent: sdk.NewInt(30),
			amountToCompound:     sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100)),
			expectedResult: []keeper.StakingCompoundAction{
				{ValidatorAddress: "cosmosval1", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))},
				{ValidatorAddress: "cosmosval2", Balance: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(20))},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := s.App.CompoundKeeper.HandleLeftOverAmount(tc.compoundActions, tc.totalCompoundPercent, tc.amountToCompound)
			require.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestStakingCompoundAmount(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()
	bondDenom := s.App.StakingKeeper.BondDenom(s.Ctx)

	delegations := []distrTypes.DelegationDelegatorReward{
		{
			ValidatorAddress: "cosmosvalopr11",
			Reward: sdk.DecCoins{
				sdk.NewDecCoin(bondDenom, sdk.NewInt(100)),
			},
		},
		{
			ValidatorAddress: "cosmosvaloper12",
			Reward: sdk.DecCoins{
				sdk.NewDecCoin(bondDenom, sdk.NewInt(200)),
			},
		},
	}

	outstandingRewards := s.App.CompoundKeeper.StakingCompoundAmount(delegations)

	require.Equal(t, sdk.NewInt(300), outstandingRewards.Amount)
	require.Equal(t, bondDenom, outstandingRewards.Denom)
}

func TestExtraCompoundAmount(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()
	bondDenom := s.App.StakingKeeper.BondDenom(s.Ctx)

	cs := compTypes.CompoundSetting{
		AmountToRemain: sdk.NewCoin(bondDenom, sdk.NewInt(100000)),
	}
	walletBalance := sdk.NewCoin(bondDenom, sdk.NewInt(200000))

	extraCompoundAmount := s.App.CompoundKeeper.ExtraCompoundAmount(cs, walletBalance)
	expectedExtraCompoundAmount := sdk.NewCoin(bondDenom, sdk.NewInt(100000))
	require.Equal(t, expectedExtraCompoundAmount, extraCompoundAmount)

	walletBalance = sdk.NewCoin(bondDenom, sdk.NewInt(90000))
	extraCompoundAmount = s.App.CompoundKeeper.ExtraCompoundAmount(cs, walletBalance)
	expectedExtraCompoundAmount = sdk.NewCoin(bondDenom, sdk.NewInt(0))
	require.Equal(t, expectedExtraCompoundAmount, extraCompoundAmount)

	cs.AmountToRemain = sdk.Coin{}
	extraCompoundAmount = s.App.CompoundKeeper.ExtraCompoundAmount(cs, walletBalance)
	expectedExtraCompoundAmount = sdk.NewCoin(bondDenom, sdk.NewInt(0))
	require.Equal(t, expectedExtraCompoundAmount, extraCompoundAmount)
}

// TestBuildCompoundActions tests the BuildCompoundActions function
func TestBuildCompoundActions(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	cs := compTypes.CompoundSetting{
		Delegator: "cosmos1delegator",
		ValidatorSetting: []*compTypes.ValidatorSetting{
			{
				ValidatorAddress:  "cosmosval1",
				PercentToCompound: 10,
			},
			{
				ValidatorAddress:  "cosmosval2",
				PercentToCompound: 20,
			},
		},
		AmountToRemain: sdk.NewCoin("uatom", sdk.NewInt(100000)),
	}
	amountToCompound := sdk.NewCoin("uatom", sdk.NewInt(1000000))

	totalCompoundPercent, compoundActions := s.App.CompoundKeeper.BuildCompoundActions(cs, amountToCompound)

	require.Equal(t, sdk.NewInt(30), totalCompoundPercent)
	require.Equal(t, 2, len(compoundActions))
	require.Equal(t, "cosmosval1", compoundActions[0].ValidatorAddress)
	require.Equal(t, "cosmosval2", compoundActions[1].ValidatorAddress)
	require.Equal(t, sdk.NewCoin("uatom", sdk.NewInt(100000)), compoundActions[0].Balance)
	require.Equal(t, sdk.NewCoin("uatom", sdk.NewInt(200000)), compoundActions[1].Balance)

	cs = compTypes.CompoundSetting{
		Delegator: "cosmos1delegator",
		ValidatorSetting: []*compTypes.ValidatorSetting{
			{
				ValidatorAddress:  "cosmosval1",
				PercentToCompound: 80,
			},
			{
				ValidatorAddress:  "cosmosval2",
				PercentToCompound: 20,
			},
		},
		AmountToRemain: sdk.NewCoin("uatom", sdk.NewInt(100000)),
	}

	amountToCompound = sdk.NewCoin("uatom", sdk.NewInt(10_000_000))

	totalCompoundPercent, compoundActions = s.App.CompoundKeeper.BuildCompoundActions(cs, amountToCompound)

	require.Equal(t, sdk.NewInt(100), totalCompoundPercent)
	require.Equal(t, 2, len(compoundActions))
	require.Equal(t, "cosmosval1", compoundActions[0].ValidatorAddress)
	require.Equal(t, "cosmosval2", compoundActions[1].ValidatorAddress)
	require.Equal(t, sdk.NewCoin("uatom", sdk.NewInt(8_000_000)), compoundActions[0].Balance)
	require.Equal(t, sdk.NewCoin("uatom", sdk.NewInt(2_000_000)), compoundActions[1].Balance)

	cs = compTypes.CompoundSetting{
		Delegator: "cosmos1delegator",
		ValidatorSetting: []*compTypes.ValidatorSetting{
			{
				ValidatorAddress:  "cosmosval1",
				PercentToCompound: 80,
			},
			{
				ValidatorAddress:  "cosmosval2",
				PercentToCompound: 20,
			},
		},
		AmountToRemain: sdk.NewCoin("uatom", sdk.NewInt(100000)),
	}

	amountToCompound = sdk.NewCoin("uatom", sdk.NewInt(10_000_081))

	totalCompoundPercent, compoundActions = s.App.CompoundKeeper.BuildCompoundActions(cs, amountToCompound)

	require.Equal(t, sdk.NewInt(100), totalCompoundPercent)
	require.Equal(t, 2, len(compoundActions))
	require.Equal(t, "cosmosval1", compoundActions[0].ValidatorAddress)
	require.Equal(t, "cosmosval2", compoundActions[1].ValidatorAddress)
	require.Equal(t, sdk.NewCoin("uatom", sdk.NewInt(8_000_064)), compoundActions[0].Balance)
	require.Equal(t, sdk.NewCoin("uatom", sdk.NewInt(2_000_016)), compoundActions[1].Balance)
}

func TestTotalCompoundAmount(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()
	bondDenom := s.App.StakingKeeper.BondDenom(s.Ctx)

	delegations := []distrTypes.DelegationDelegatorReward{
		{
			ValidatorAddress: "cosmosvalopr11",
			Reward: sdk.DecCoins{
				sdk.NewDecCoin(bondDenom, sdk.NewInt(100)),
			},
		},
		{
			ValidatorAddress: "cosmosvaloper12",
			Reward: sdk.DecCoins{
				sdk.NewDecCoin(bondDenom, sdk.NewInt(200)),
			},
		},
	}

	walletBalance := sdk.NewCoin(bondDenom, sdk.NewInt(1000))
	cs := compTypes.CompoundSetting{
		Delegator:        "cosmos11",
		ValidatorSetting: []*compTypes.ValidatorSetting{},
		AmountToRemain:   sdk.NewCoin(bondDenom, sdk.NewInt(500)),
	}

	total := s.App.CompoundKeeper.TotalCompoundAmount(delegations, walletBalance, cs)
	if total.Amount.Int64() != 800 {
		t.Errorf("Total compound amount is incorrect, got: %d, want: %d.", total.Amount.Int64(), 800)
	}

	walletBalance = sdk.NewCoin(bondDenom, sdk.NewInt(1500))
	cs = compTypes.CompoundSetting{
		Delegator:        "cosmos12",
		ValidatorSetting: []*compTypes.ValidatorSetting{},
		AmountToRemain:   sdk.NewCoin(bondDenom, sdk.NewInt(1500)),
	}

	total = s.App.CompoundKeeper.TotalCompoundAmount(delegations, walletBalance, cs)
	if total.Amount.Int64() != 300 {
		t.Errorf("Total compound amount is incorrect, got: %d, want: %d.", total.Amount.Int64(), 300)
	}

	delegations = []distrTypes.DelegationDelegatorReward{}
	walletBalance = sdk.NewCoin(bondDenom, sdk.NewInt(1500))
	cs = compTypes.CompoundSetting{
		Delegator:        "cosmos13",
		ValidatorSetting: []*compTypes.ValidatorSetting{},
		AmountToRemain:   sdk.NewCoin(bondDenom, sdk.NewInt(1500)),
	}

	total = s.App.CompoundKeeper.TotalCompoundAmount(delegations, walletBalance, cs)
	if total.Amount.Int64() != 0 {
		t.Errorf("Total compound amount is incorrect, got: %d, want: %d.", total.Amount.Int64(), 0)
	}

	walletBalance = sdk.NewCoin(bondDenom, sdk.NewInt(0))
	cs = compTypes.CompoundSetting{
		Delegator:        "cosmos13",
		ValidatorSetting: []*compTypes.ValidatorSetting{},
		AmountToRemain:   sdk.NewCoin(bondDenom, sdk.NewInt(1500)),
	}

	total = s.App.CompoundKeeper.TotalCompoundAmount(delegations, walletBalance, cs)
	if total.Amount.Int64() != 0 {
		t.Errorf("Total compound amount is incorrect, got: %d, want: %d.", total.Amount.Int64(), 0)
	}

	walletBalance = sdk.NewCoin(bondDenom, sdk.NewInt(1500))
	cs = compTypes.CompoundSetting{
		Delegator:        "cosmos13",
		ValidatorSetting: []*compTypes.ValidatorSetting{},
		AmountToRemain:   sdk.NewCoin(bondDenom, sdk.NewInt(1000)),
	}

	total = s.App.CompoundKeeper.TotalCompoundAmount(delegations, walletBalance, cs)
	if total.Amount.Int64() != 500 {
		t.Errorf("Total compound amount is incorrect, got: %d, want: %d.", total.Amount.Int64(), 500)
	}
}

func TestRecordCompound(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	// Test case 1: New delegator record
	address1 := "address1"
	blockHeight1 := s.Ctx.BlockHeight()
	s.App.CompoundKeeper.RecordCompound(s.Ctx, address1)

	value, found := s.App.CompoundKeeper.GetPreviousCompound(s.Ctx, address1)
	if !found {
		t.Fatalf("Expected to find the address in the store but not found")
	}
	if value.Delegator != address1 {
		t.Fatalf("Expected delegator address to be %s but got %s", address1, value.Delegator)
	}
	if value.BlockHeight != blockHeight1 {
		t.Fatalf("Expected block time to be %d but got %d", blockHeight1, value.BlockHeight)
	}

	// Test case 2: Update existing record
	s.App.CompoundKeeper.RecordCompound(s.Ctx, address1)

	value, found = s.App.CompoundKeeper.GetPreviousCompound(s.Ctx, address1)
	if !found {
		t.Fatalf("Expected to find the address in the store but not found")
	}
	if value.Delegator != address1 {
		t.Fatalf("Expected delegator address to be %s but got %s", address1, value.Delegator)
	}
	if value.BlockHeight != blockHeight1 {
		t.Fatalf("Expected block time to be %d but got %d", blockHeight1, value.BlockHeight)
	}
}

func TestCalculateCompoundingAmount(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	var tests = []struct {
		name              string
		rewardAmount      math.Int
		percentToCompound uint64
		expectedAmount    math.Int
	}{
		{
			name:              "Test case 1",
			rewardAmount:      sdk.NewInt(100),
			percentToCompound: 10,
			expectedAmount:    sdk.NewInt(10),
		},
		{
			name:              "Test case 2",
			rewardAmount:      sdk.NewInt(100),
			percentToCompound: 30,
			expectedAmount:    sdk.NewInt(30),
		},
		{
			name:              "Test case 3",
			rewardAmount:      sdk.NewInt(200),
			percentToCompound: 20,
			expectedAmount:    sdk.NewInt(40),
		},
		{
			name:              "Test case 4",
			rewardAmount:      sdk.NewInt(200),
			percentToCompound: 0,
			expectedAmount:    sdk.NewInt(0),
		},
		{
			name:              "Zero case",
			rewardAmount:      sdk.NewInt(0),
			percentToCompound: 50,
			expectedAmount:    math.NewInt(0),
		},
	}

	for _, test := range tests {
		result := s.App.CompoundKeeper.CalculateCompoundAmount(sdk.NewCoin("test", test.rewardAmount), test.percentToCompound)
		if !result.Sub(test.expectedAmount).IsZero() {
			t.Errorf("Test case %s failed: expected %s but got %s", test.name, test.expectedAmount, result)
		}
	}
}

func networkWithCustomParams(t *testing.T) *network.Network {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{Params: types.NewParams(100, 5)}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg)
}

// TestRunCompounding writes a CompoundSetting to state and waits the minimumCompoundFrequency and checks if the compound happened
func TestRunCompounding(t *testing.T) {
	net := networkWithCustomParams(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}

	//Get existing DelegationHistory objects from state
	queryArgs := request(nil, 0, uint64(10000), true)
	out, err := clitestutil.ExecTestCLICmd(ctx, recordCli.CmdListDelegationHistory(), queryArgs)
	require.NoError(t, err)
	var currentDelHistoryList recordTypes.QueryAllDelegationHistoryResponse
	require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &currentDelHistoryList))

	//Get validators in test network
	out, err = clitestutil.ExecTestCLICmd(ctx, stakingCli.GetCmdQueryValidators(), queryArgs)
	require.NoError(t, err)
	var respValidators stakingTypes.QueryValidatorsResponse
	require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &respValidators))

	// When a validator is created they self delegate some amount so there should be one of each of the below for each validator
	require.Equal(t, len(currentDelHistoryList.GetDelegationHistory()), len(respValidators.GetValidators()))

	for i := range respValidators.GetValidators() {
		require.Equal(t, len(currentDelHistoryList.GetDelegationHistory()[i].GetHistory()), 1)
	}

	//Get existing PreviousCompound objects from state
	out, err = clitestutil.ExecTestCLICmd(ctx, compoundCli.CmdListPreviousCompound(), queryArgs)
	require.NoError(t, err)
	var currentPrevCompoundList compTypes.QueryAllPreviousCompoundResponse
	require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &currentPrevCompoundList))

	//Since no CompoundSettings are present so far, there should not be any PreviousCompounds
	require.Equal(t, len(currentPrevCompoundList.GetPreviousCompound()), 0)

	valSetting := fmt.Sprintf("[{\"validatorAddress\":\"%s\",\"percentToCompound\":50}]", val.ValAddress.String())

	broadcastArgs := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}

	tests := []struct {
		desc        string
		idDelegator string
		fields      []string

		args []string
		err  error
		code uint32
	}{
		{
			desc:        "InvalidValidatorSettings",
			idDelegator: strconv.Itoa(0),
			fields:      []string{"null", "10token", "600"},
			code:        18,
			args:        broadcastArgs,
		},
		{
			desc:        "ValidValidatorSettings",
			idDelegator: strconv.Itoa(1),
			fields:      []string{valSetting, "", "5"},
			code:        0,
			args:        broadcastArgs,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			require.NoError(t, net.WaitForNextBlock())

			var args []string
			args = append(args, tc.fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateCompoundSetting(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			var resp sdk.TxResponse
			require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NoError(t, clitestutil.CheckTxCode(net, ctx, resp.TxHash, tc.code))

			err = net.WaitForNextBlock()
			require.NoError(t, err)

			if tc.desc == "ValidValidatorSettings" {
				//Get existing PreviousCompound objects from state, 1 should exist
				args = request(nil, 0, uint64(10000), true)
				out, err = clitestutil.ExecTestCLICmd(ctx, compoundCli.CmdListPreviousCompound(), args)
				require.NoError(t, err)
				var currentPrevCompoundListStart compTypes.QueryAllPreviousCompoundResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &currentPrevCompoundListStart))
				require.Equal(t, len(currentPrevCompoundListStart.GetPreviousCompound()), 1)

				height, err := net.LatestHeight()
				require.NoError(t, err)

				// ~2 seconds blocks
				_, err = net.WaitForHeightWithTimeout(5+height, time.Second*15)
				require.NoError(t, err)

				//Get existing DelegationHistory objects from state
				args = request(nil, 0, uint64(10000), true)
				out, err = clitestutil.ExecTestCLICmd(ctx, recordCli.CmdListDelegationHistory(), args)
				require.NoError(t, err)
				var currentDelegetaionHistoryList recordTypes.QueryAllDelegationHistoryResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &currentDelegetaionHistoryList))

				//There should only be 1 DelegationHistory object with multiple DelegationTimestamps
				delhistory := currentDelegetaionHistoryList.GetDelegationHistory()
				require.Equal(t, len(delhistory), 1)

				prevComps := delhistory[0].GetHistory()
				require.GreaterOrEqual(t, len(prevComps), 1)

				//The DelegationHistory should match the CompoundSetting that was just created
				require.Equal(t, delhistory[0].Address, val.Address.String())
			}
		})
	}
}
