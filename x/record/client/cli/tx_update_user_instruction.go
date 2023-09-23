package cli

import (
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/temporal-zone/temporal/x/record/types"
)

var _ = strconv.Itoa(0)

func CmdUpdateUserInstruction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-user-instruction [remote-address] [chain-id] [frequency] [expires] [instruction] [strategy-id] [contract-address]",
		Short: "Broadcast message UpdateUserInstruction",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRemoteAddress := args[0]

			argChainId := args[1]

			argFrequency, err := cast.ToInt64E(args[2])
			if err != nil {
				return err
			}
			argExpires, err := cast.ToInt64E(args[3])
			if err != nil {
				return err
			}
			argExpiresTime := time.Unix(argExpires, 0)

			argInstruction := args[4]

			argStrategyId, err := cast.ToInt64E(args[5])
			if err != nil {
				return err
			}

			argContractAddress := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateUserInstruction(
				clientCtx.GetFromAddress().String(),
				argRemoteAddress,
				argChainId,
				argFrequency,
				argExpiresTime,
				argInstruction,
				argStrategyId,
				argContractAddress,
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
