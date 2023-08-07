package cli_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"strconv"
	"testing"
	"time"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	stakingCli "github.com/cosmos/cosmos-sdk/x/staking/client/cli"
	"github.com/temporal-zone/temporal/testutil/network"
	"github.com/temporal-zone/temporal/testutil/nullify"
	recordCli "github.com/temporal-zone/temporal/x/record/client/cli"
	recordTypes "github.com/temporal-zone/temporal/x/record/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithDelegationHistoryObjects(t *testing.T, n int) (*network.Network, []recordTypes.DelegationHistory) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := recordTypes.GenesisState{}

	delegationTimestamps := make([]*recordTypes.DelegationTimestamp, 1)
	delegationTimestamp := recordTypes.DelegationTimestamp{
		Timestamp: time.Now(),
		Balance:   sdk.Coin{},
	}

	delegationTimestamps[0] = &delegationTimestamp

	for i := 0; i < n; i++ {
		delegationHistory := recordTypes.DelegationHistory{
			Address: strconv.Itoa(i),
			History: delegationTimestamps,
		}
		//nullify.Fill(&delegationHistory)
		state.DelegationHistoryList = append(state.DelegationHistoryList, delegationHistory)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[recordTypes.ModuleName] = buf
	return network.New(t, cfg), state.DelegationHistoryList
}

func TestShowDelegationHistory(t *testing.T) {
	net, objs := networkWithDelegationHistoryObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		desc      string
		idAddress string

		args []string
		err  error
		obj  recordTypes.DelegationHistory
	}{
		{
			desc:      "found",
			idAddress: objs[0].Address,

			args: common,
			obj:  objs[0],
		},
		{
			desc:      "not found",
			idAddress: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idAddress,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, recordCli.CmdShowDelegationHistory(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp recordTypes.QueryGetDelegationHistoryResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.DelegationHistory)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.DelegationHistory),
				)
			}
		})
	}
}

func TestListDelegationHistory(t *testing.T) {
	numberOfDelegationHistoryObjs := 5
	net, objs := networkWithDelegationHistoryObjects(t, numberOfDelegationHistoryObjs)

	ctx := net.Validators[0].ClientCtx
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
	args := request(nil, 0, uint64(10000), true)
	out, err := clitestutil.ExecTestCLICmd(ctx, recordCli.CmdListDelegationHistory(), args)
	require.NoError(t, err)
	var resp recordTypes.QueryAllDelegationHistoryResponse
	require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))

	//Get number of validators in test network
	out, err = clitestutil.ExecTestCLICmd(ctx, stakingCli.GetCmdQueryValidators(), args)
	require.NoError(t, err)
	var respValidators stakingTypes.QueryValidatorsResponse
	require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &respValidators))

	//Filter out the validator DelegationHistory objects
	valObjs := make([]recordTypes.DelegationHistory, 0)
	for _, validator := range respValidators.GetValidators() {
		valAddr, err := sdk.ValAddressFromBech32(validator.OperatorAddress)
		require.NoError(t, err)
		valAccAddress := sdk.AccAddress(valAddr)
		for _, delHistory := range resp.GetDelegationHistory() {
			if valAccAddress.String() == delHistory.Address {
				valObjs = append(valObjs, delHistory)
			}
		}
	}

	// Make sure the number of DelegationHistory objects match so far
	require.Equal(t, len(resp.GetDelegationHistory()), len(valObjs)+numberOfDelegationHistoryObjs)

	objs = append(objs, valObjs...)

	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, recordCli.CmdListDelegationHistory(), args)
			require.NoError(t, err)
			var resp recordTypes.QueryAllDelegationHistoryResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.DelegationHistory), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.DelegationHistory),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, recordCli.CmdListDelegationHistory(), args)
			require.NoError(t, err)
			var resp recordTypes.QueryAllDelegationHistoryResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.DelegationHistory), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.DelegationHistory),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, recordCli.CmdListDelegationHistory(), args)
		require.NoError(t, err)
		var resp recordTypes.QueryAllDelegationHistoryResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.DelegationHistory),
		)
	})
}
