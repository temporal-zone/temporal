package cli

import (
	"github.com/spf13/cast"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/temporal-zone/temporal/x/record/types"
)

var _ = strconv.Itoa(0)

func CmdCreateUserInstruction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-user-instruction [remote-address] [chain-id] [frequency] [expires] [instructions] [strategy-id] [contract-address]",
		Short: "Broadcast message CreateUserInstruction",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

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

			msg := types.NewMsgCreateUserInstruction(
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
