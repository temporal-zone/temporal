package cli_test

import (
	"fmt"
	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/temporal-zone/temporal/x/compound/types"
	"strconv"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/temporal-zone/temporal/testutil/network"
	"github.com/temporal-zone/temporal/x/compound/client/cli"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCreateCompoundSetting(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 2
	state := types.GenesisState{Params: types.NewParams(uint64(100), uint64(5))}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf

	net := network.New(t, cfg)
	val := net.Validators[0]
	val1 := net.Validators[1]
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

	tests := []struct {
		desc               string
		valSetting         string
		amountToRemain     string
		frequency          string
		compoundValidators []string

		args []string
		err  error
		code uint32
	}{
		{
			desc:               "valid 1",
			valSetting:         fmt.Sprintf("[{\"validatorAddress\":\"%s\",\"percentToCompound\":50}]", val.ValAddress.String()),
			amountToRemain:     "10utprl",
			frequency:          "111",
			compoundValidators: []string{val.ValAddress.String()},

			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
			},
		},
		{
			desc: "valid 2",
			valSetting: fmt.Sprintf(
				"[{\"validatorAddress\":\"%s\",\"percentToCompound\":50},{\"validatorAddress\":\"%s\",\"percentToCompound\":50}]",
				val.ValAddress.String(),
				val1.ValAddress.String()),
			amountToRemain:     "10utprl",
			frequency:          "111",
			compoundValidators: []string{val.ValAddress.String(), val1.ValAddress.String()},

			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			require.NoError(t, net.WaitForNextBlock())

			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdDeleteCompoundSetting(), tc.args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}

			require.NoError(t, net.WaitForNextBlock())

			var args []string
			fields := []string{tc.valSetting, tc.amountToRemain, tc.frequency}
			args = append(args, fields...)
			args = append(args, tc.args...)
			out, err = clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateCompoundSetting(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			require.NoError(t, net.WaitForNextBlock())

			var resp sdk.TxResponse
			require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NoError(t, clitestutil.CheckTxCode(net, ctx, resp.TxHash, tc.code))

			_ = request(nil, 0, uint64(10000), true)
			args = append([]string{val.Address.String()}, fmt.Sprintf("--%s=json", tmcli.OutputFlag))
			out, err = clitestutil.ExecTestCLICmd(ctx, cli.CmdShowCompoundSetting(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			var compoundSetting types.QueryGetCompoundSettingResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &compoundSetting))
			require.NotNil(t, compoundSetting.GetCompoundSetting())
			require.Equal(t, len(compoundSetting.GetCompoundSetting().ValidatorSetting), len(tc.compoundValidators))

			for i, _ := range tc.compoundValidators {
				require.Equal(t, compoundSetting.GetCompoundSetting().ValidatorSetting[i].ValidatorAddress, tc.compoundValidators[i])
			}
		})
	}
}

func TestUpdateCompoundSetting(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	valSetting := fmt.Sprintf("[{\"validatorAddress\":\"%s\",\"percentToCompound\":50}]", val.ValAddress.String())

	fields := []string{valSetting, "10token", "111"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}
	var args []string
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateCompoundSetting(), args)
	require.NoError(t, err)

	tests := []struct {
		desc string

		args []string
		code uint32
		err  error
	}{
		{
			desc: "valid",

			args: common,
		},
		{
			desc: "key not found",

			args: common,
			code: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			require.NoError(t, net.WaitForNextBlock())

			var args []string
			args = append(args, fields...)
			args = append(args, tc.args...)

			if tc.desc == "key not found" {
				out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdDeleteCompoundSetting(), tc.args)

				require.NoError(t, err)

				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NoError(t, clitestutil.CheckTxCode(net, ctx, resp.TxHash, 0))
			}

			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdUpdateCompoundSetting(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			var resp sdk.TxResponse
			require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NoError(t, clitestutil.CheckTxCode(net, ctx, resp.TxHash, tc.code))
		})
	}
}

func TestDeleteCompoundSetting(t *testing.T) {
	net := network.New(t)

	val := net.Validators[0]
	ctx := val.ClientCtx

	valSetting := fmt.Sprintf("[{\"validatorAddress\":\"%s\",\"percentToCompound\":50}]", val.ValAddress.String())

	fields := []string{valSetting, "10" + net.Config.BondDenom, "111"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}
	var args []string
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateCompoundSetting(), args)
	require.NoError(t, err)

	tests := []struct {
		desc string

		args []string
		code uint32
		err  error
	}{
		{
			desc: "valid",

			args: common,
		},
		{
			desc: "key not found",

			args: common,
			code: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			require.NoError(t, net.WaitForNextBlock())

			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdDeleteCompoundSetting(), tc.args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			var resp sdk.TxResponse
			require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NoError(t, clitestutil.CheckTxCode(net, ctx, resp.TxHash, tc.code))
		})
	}
}
