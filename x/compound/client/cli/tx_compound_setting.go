package cli

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"temporal/x/compound/types"
)

func CmdCreateCompoundSetting() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-compound-setting [index-123] [validator-setting] [amount-to-remain] [frequency]",
		Short: "Create a new CompoundSetting",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex123 := args[0]

			// Get value arguments
			argValidatorSetting := new(types.ValidatorSetting)
			err = json.Unmarshal([]byte(args[1]), argValidatorSetting)
			if err != nil {
				return err
			}
			argAmountToRemain, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}
			argFrequency, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCompoundSetting(
				clientCtx.GetFromAddress().String(),
				indexIndex123,
				argValidatorSetting,
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
		Use:   "update-compound-setting [index-123] [validator-setting] [amount-to-remain] [frequency]",
		Short: "Update a CompoundSetting",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex123 := args[0]

			// Get value arguments
			argValidatorSetting := new(types.ValidatorSetting)
			err = json.Unmarshal([]byte(args[1]), argValidatorSetting)
			if err != nil {
				return err
			}
			argAmountToRemain, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}
			argFrequency, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateCompoundSetting(
				clientCtx.GetFromAddress().String(),
				indexIndex123,
				argValidatorSetting,
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
		Use:   "delete-compound-setting [index-123]",
		Short: "Delete a CompoundSetting",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexIndex123 := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteCompoundSetting(
				clientCtx.GetFromAddress().String(),
				indexIndex123,
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
