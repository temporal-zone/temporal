package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/temporal-zone/temporal/x/record/types"
)

var _ = strconv.Itoa(0)

func CmdDeleteUserInstruction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-user-instruction [remote-address] [chain-id] [contract-address]",
		Short: "Broadcast message DeleteUserInstruction",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argChainId := args[1]
			argContractAddress := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteUserInstruction(
				clientCtx.GetFromAddress().String(),
				argAddress,
				argChainId,
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
