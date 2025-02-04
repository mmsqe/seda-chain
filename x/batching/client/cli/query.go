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
		GetCmdQueryBatchByHeight(),
		GetCmdQueryBatches(),
		GetCmdQueryTreeEntries(),
		GetCmdQueryBatchSignatures(),
		GetCmdQueryDataResult(),
		GetCmdQueryBatchAssignment(),
	)
	return cmd
}

// GetCmdQueryBatch returns the command for querying a specific batch.
func GetCmdQueryBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batch <optional_batch_number>",
		Short: "Query a batch",
		Long:  "Query the latest batch whose signatures have been collected or, by providing its batch number, a specific batch.",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			var batchNum uint64
			if len(args) != 0 {
				batchNum, err = strconv.ParseUint(args[0], 10, 64)
				if err != nil {
					return err
				}
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

// GetCmdQueryBatchByHeight returns the command for querying a
// batch using a block height.
func GetCmdQueryBatchByHeight() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batch-by-height <block_height>",
		Short: "Get a batch given its creation block height",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			blockHeight, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}
			req := &types.QueryBatchForHeightRequest{
				BlockHeight: blockHeight,
			}
			res, err := queryClient.BatchForHeight(cmd.Context(), req)
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

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := queryClient.Batches(cmd.Context(), &types.QueryBatchesRequest{Pagination: pageReq})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "batches")
	return cmd
}

// GetCmdQueryTreeEntries returns the command for querying tree entries
// corresponding to the given batch number.
func GetCmdQueryTreeEntries() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tree-entries <batch_number>",
		Short: "Get tree entries given its batch number",
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
			req := &types.QueryTreeEntriesRequest{
				BatchNumber: batchNum,
			}
			res, err := queryClient.TreeEntries(cmd.Context(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryBatchSignatures returns the command for querying batch
// signatures of a given batch.
func GetCmdQueryBatchSignatures() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "batch-signatures <batch_number>",
		Aliases: []string{"batch-sigs"},
		Short:   "Get batch signatures given its batch number",
		Args:    cobra.ExactArgs(1),
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
			req := &types.QueryBatchSignaturesRequest{
				BatchNumber: batchNum,
			}
			res, err := queryClient.BatchSignatures(cmd.Context(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryBatch returns the command for querying a data result
// associated with a given data request.
func GetCmdQueryDataResult() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "data-result <data_request_id>",
		Short: "Get the data result associated with a given data request",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryDataResultRequest{
				DataRequestId: args[0],
			}
			res, err := queryClient.DataResult(cmd.Context(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryBatchAssignment returns the command for querying the
// batch number assigned to the given data request.
func GetCmdQueryBatchAssignment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batch-assignment <data_request_id>",
		Short: "Get the batch number assigned to a given data request",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryBatchAssignmentRequest{
				DataRequestId: args[0],
			}
			res, err := queryClient.BatchAssignment(cmd.Context(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
