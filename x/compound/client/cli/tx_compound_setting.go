package cli

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/temporal-zone/temporal/x/compound/types"
)

func CmdCreateCompoundSetting() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-compound-setting [validator-setting] [amount-to-remain] [frequency]",
		Short: "Create a new CompoundSetting for yourself",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get value arguments
			argValidatorSettingBytes := new([]types.ValidatorSetting)
			err = json.Unmarshal([]byte(args[0]), argValidatorSettingBytes)
			if err != nil {
				return err
			}

			argValidatorSettings := make([]*types.ValidatorSetting, 0, 100)
			for _, vs := range *argValidatorSettingBytes {
				argValidatorSettings = append(argValidatorSettings, &vs)
			}

			argAmountToRemain, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}
			argFrequency, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCompoundSetting(
				clientCtx.GetFromAddress().String(),
				argValidatorSettings,
				argAmountToRemain,
				argFrequency,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateCompoundSetting() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-compound-setting [validator-setting] [amount-to-remain] [frequency]",
		Short: "Update your own CompoundSetting",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get value arguments
			argValidatorSettingBytes := new([]types.ValidatorSetting)
			err = json.Unmarshal([]byte(args[0]), argValidatorSettingBytes)
			if err != nil {
				return err
			}

			argValidatorSettings := make([]*types.ValidatorSetting, 0, 100)
			for _, vs := range *argValidatorSettingBytes {
				argValidatorSettings = append(argValidatorSettings, &vs)
			}

			argAmountToRemain, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}
			argFrequency, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateCompoundSetting(
				clientCtx.GetFromAddress().String(),
				argValidatorSettings,
				argAmountToRemain,
				argFrequency,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteCompoundSetting() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-compound-setting",
		Short: "Delete your own CompoundSetting",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteCompoundSetting(
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
