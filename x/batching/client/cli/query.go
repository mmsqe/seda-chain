package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/sedaprotocol/seda-chain/x/batching/types"
)

// GetQueryCmd returns the CLI query commands for batching module.
func GetQueryCmd(_ string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdQueryBatch(),
		GetCmdQueryBatches(),
	)
	return cmd
}

// GetCmdQueryBatch returns the command for querying a batch.
func GetCmdQueryBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batch <batch_number>",
		Short: "Get a batch given its batch number",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			batchNum, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			req := &types.QueryBatchRequest{
				BatchNumber: batchNum,
			}
			res, err := queryClient.Batch(cmd.Context(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryBatches returns the command for querying all batches.
func GetCmdQueryBatches() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batches",
		Short: "List all batches in the store",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Batches(cmd.Context(), &types.QueryBatchesRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}