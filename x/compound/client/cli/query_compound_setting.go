package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"temporal/x/compound/types"
)

func CmdListCompoundSetting() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-compound-setting",
		Short: "list all CompoundSetting",
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

			params := &types.QueryAllCompoundSettingRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.CompoundSettingAll(cmd.Context(), params)
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

func CmdShowCompoundSetting() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-compound-setting [index-123]",
		Short: "shows a CompoundSetting",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argDelegator := args[0]

			params := &types.QueryGetCompoundSettingRequest{
				Delegator: argDelegator,
			}

			res, err := queryClient.CompoundSetting(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
