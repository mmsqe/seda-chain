package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/sedaprotocol/seda-chain/x/wasm-storage/types"
)

// GetQueryCmd returns the CLI query commands for this module
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdQueryOracleProgram(),
		GetCmdQueryExecutorWasm(),
		GetCmdQueryOraclePrograms(),
		GetCmdQueryExecutorWasms(),
		GetCmdQueryCoreContractRegistry(),
	)
	return cmd
}

// GetCmdQueryOracleProgram returns the command for querying oracle program.
func GetCmdQueryOracleProgram() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle-program <hash>",
		Short: "Get oracle program given its hash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryOracleProgramRequest{
				Hash: args[0],
			}
			res, err := queryClient.OracleProgram(cmd.Context(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryExecutorWasm returns the command for querying an executor Wasm.
func GetCmdQueryExecutorWasm() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "executor-wasm <hash>",
		Short: "Get an executor wasm given its hash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryExecutorWasmRequest{
				Hash: args[0],
			}
			res, err := queryClient.ExecutorWasm(cmd.Context(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryOraclePrograms returns the command for querying
// oracle programs in the store.
func GetCmdQueryOraclePrograms() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-oracle-programs",
		Short: "List hashes and expiration heights of all oracle programs",
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
			res, err := queryClient.OraclePrograms(
				cmd.Context(),
				&types.QueryOracleProgramsRequest{
					Pagination: pageReq,
				},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "oracle programs")
	return cmd
}

// GetCmdQueryExecutorWasms returns the command for querying all
// executor wasms.
func GetCmdQueryExecutorWasms() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-executor-wasms",
		Short: "List hashes of all executor wasms",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ExecutorWasms(cmd.Context(), &types.QueryExecutorWasmsRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryCoreContractRegistry returns the command for querying
// Core Contract registry.
func GetCmdQueryCoreContractRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "core-contract-registry",
		Short: "Get the address of Core Contract",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.CoreContractRegistry(
				cmd.Context(),
				&types.QueryCoreContractRegistryRequest{},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
