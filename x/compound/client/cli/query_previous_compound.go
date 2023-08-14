package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/temporal-zone/temporal/x/compound/types"
)

func CmdListPreviousCompound() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-previous-compound",
		Short: "list all PreviousCompound",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllPreviousCompoundRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.PreviousCompoundAll(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowPreviousCompound() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-previous-compound [delegator]",
		Short: "shows a PreviousCompound",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argDelegator := args[0]

			params := &types.QueryGetPreviousCompoundRequest{
				Delegator: argDelegator,
			}

			res, err := queryClient.PreviousCompound(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
