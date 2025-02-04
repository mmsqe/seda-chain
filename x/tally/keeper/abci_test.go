package keeper_test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"

	"github.com/sedaprotocol/seda-wasm-vm/tallyvm"

	"github.com/sedaprotocol/seda-chain/x/tally/keeper/testdata"
	"github.com/sedaprotocol/seda-chain/x/tally/types"
	wasmstoragetypes "github.com/sedaprotocol/seda-chain/x/wasm-storage/types"
)

// TestTallyVM tests tally VM using a sample tally wasm that performs
// preliminary checks on the given reveal data.
func TestTallyVM(t *testing.T) {
	cases := []struct {
		name        string
		requestJSON []byte
		args        []string
		expErr      string
	}{
		{
			name: "Three reveals",
			requestJSON: []byte(`{
				"commits":{},
				"dr_binary_id":"9471d36add157cd7eaa32a42b5ddd091d5d5d396bf9ad67938a4fc40209df6cf",
				"dr_inputs":"",
				"gas_limit":5000000000,
				"gas_price":"10",
				"height":1661661742461173125,
				"id":"fba5314c57e52da7d1a2245d18c670fde1cb8c237062d2a1be83f449ace0932e",
				"memo":"",
				"payback_address":"",
				"replication_factor":3,
				"reveals":{
				   "1b85dfb9420e6757630a0db2280fa1787ec8c1e419a6aca76dbbfe8ef6e17521":{
					  "exit_code":0,
					  "gas_used":10,
					  "reveal":"Ng==",
					  "salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"
				   },
				   "1dae290cd880b79d21079d89aee3460cf8a7d445fb35cade70cf8aa96924441c":{
					  "exit_code":0,
					  "gas_used":10,
					  "reveal":"LQ==",
					  "salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"
				   },
				   "421e735518ef77fc1209a9d3585cdf096669b52ea68549e2ce048d4919b4c8c0":{
					  "exit_code":0,
					  "gas_used":10,
					  "reveal":"DQ==",
					  "salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"
				   }
				},
				"seda_payload":"",
				"tally_binary_id":"8ade60039246740faa80bf424fc29e79fe13b32087043e213e7bc36620111f6b",
				"tally_inputs":"AAEBAQE=",
				"version":"1.0.0"
			 }`),
			args:   []string{"6d792d74616c6c792d696e70757473", "[{\"reveal\":[54],\"salt\":[211,159,121,219,109,120,111,102,218,223,158,61,107,199,122,219,183,57,237,221,157,209,215,117,111,70,182,113,238,185,115,142,158,221,189,159,219,151,54,239,126,58,225,183,188,109,174,95],\"exit_code\":0,\"gas_used\":\"10\"},{\"reveal\":[45],\"salt\":[211,159,121,219,109,120,111,102,218,223,158,61,107,199,122,219,183,57,237,221,157,209,215,117,111,70,182,113,238,185,115,142,158,221,189,159,219,151,54,239,126,58,225,183,188,109,174,95],\"exit_code\":0,\"gas_used\":\"10\"},{\"reveal\":[13],\"salt\":[211,159,121,219,109,120,111,102,218,223,158,61,107,199,122,219,183,57,237,221,157,209,215,117,111,70,182,113,238,185,115,142,158,221,189,159,219,151,54,239,126,58,225,183,188,109,174,95],\"exit_code\":0,\"gas_used\":\"10\"}]", "[0,0,0]"},
			expErr: "",
		},
		{
			name: "One less outlier provided",
			requestJSON: []byte(`{
				"commits":{},
				"dr_binary_id":"9471d36add157cd7eaa32a42b5ddd091d5d5d396bf9ad67938a4fc40209df6cf",
				"dr_inputs":"",
				"gas_limit":5000000000,
				"gas_price":"10",
				"height":1661661742461173125,
				"id":"fba5314c57e52da7d1a2245d18c670fde1cb8c237062d2a1be83f449ace0932e",
				"memo":"",
				"payback_address":"",
				"replication_factor":3,
				"reveals":{
				   "1b85dfb9420e6757630a0db2280fa1787ec8c1e419a6aca76dbbfe8ef6e17521":{
					  "exit_code":0,
					  "gas_used":10,
					  "reveal":"Ng==",
					  "salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"
				   },
				   "1dae290cd880b79d21079d89aee3460cf8a7d445fb35cade70cf8aa96924441c":{
					  "exit_code":0,
					  "gas_used":10,
					  "reveal":"LQ==",
					  "salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"
				   },
				   "421e735518ef77fc1209a9d3585cdf096669b52ea68549e2ce048d4919b4c8c0":{
					  "exit_code":0,
					  "gas_used":10,
					  "reveal":"DQ==",
					  "salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"
				   }
				},
				"seda_payload":"",
				"tally_binary_id":"8ade60039246740faa80bf424fc29e79fe13b32087043e213e7bc36620111f6b",
				"tally_inputs":"AAEBAQE=",
				"version":"1.0.0"
			 }`),
			args:   []string{"6d792d74616c6c792d696e70757473", "[{\"reveal\":[54],\"salt\":[211,159,121,219,109,120,111,102,218,223,158,61,107,199,122,219,183,57,237,221,157,209,215,117,111,70,182,113,238,185,115,142,158,221,189,159,219,151,54,239,126,58,225,183,188,109,174,95],\"exit_code\":0,\"gas_used\":\"10\"},{\"reveal\":[45],\"salt\":[211,159,121,219,109,120,111,102,218,223,158,61,107,199,122,219,183,57,237,221,157,209,215,117,111,70,182,113,238,185,115,142,158,221,189,159,219,151,54,239,126,58,225,183,188,109,174,95],\"exit_code\":0,\"gas_used\":\"10\"},{\"reveal\":[13],\"salt\":[211,159,121,219,109,120,111,102,218,223,158,61,107,199,122,219,183,57,237,221,157,209,215,117,111,70,182,113,238,185,115,142,158,221,189,159,219,151,54,239,126,58,225,183,188,109,174,95],\"exit_code\":0,\"gas_used\":\"10\"}]", "[0,0]"},
			expErr: "abort: Number of reveals (3) does not equal number of consensus reports (2)",
		},
		{
			name: "One reveal",
			requestJSON: []byte(`{
				"commits":{},
				"dr_binary_id":"9471d36add157cd7eaa32a42b5ddd091d5d5d396bf9ad67938a4fc40209df6cf",
				"dr_inputs":"",
				"gas_limit":5000000000,
				"gas_price":"10",
				"height":9859593541233596221,
				"id":"d4e40f45fbf529134926acf529baeb6d4f37b5c380d7ab6b934833e7c00d725f",
				"memo":"",
				"payback_address":"",
				"replication_factor":1,
				"reveals":{
				   "c9a4c8f1e70a0059a88b4768a920e41c95c587b8387ea3286d8fa4ee3b68b038":{
					  "exit_code":0,
					  "gas_used":10,
					  "reveal":"Yw==",
					  "salt":"f837455a930a66464f1c50586dc745a6b14ea807727c6069acac24c9558b6dbf"
				   }
				},
				"seda_payload":"",
				"tally_binary_id":"8ade60039246740faa80bf424fc29e79fe13b32087043e213e7bc36620111f6b",
				"tally_inputs":"AAEBAQE=",
				"version":"1.0.0"
			 }`),
			args:   []string{"6d792d74616c6c792d696e70757473", "[{\"reveal\":[99],\"salt\":[127,205,251,227,158,90,247,125,26,235,174,58,225,253,92,231,78,124,233,215,59,227,150,186,111,94,30,107,205,59,239,110,220,235,78,189,105,198,156,219,135,61,231,159,27,233,214,223],\"exit_code\":0,\"gas_used\":\"10\"}]", "[0]"},
			expErr: "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var req types.Request
			err := json.Unmarshal(tc.requestJSON, &req)
			require.NoError(t, err)

			result := tallyvm.ExecuteTallyVm(testdata.SampleTallyWasm(), tc.args, map[string]string{
				"VM_MODE":               "tally",
				"CONSENSUS":             fmt.Sprintf("%v", true),
				"DR_ID":                 req.ID,
				"DR_INPUT":              req.DrInputs,
				"BINARY_ID":             req.DrBinaryID,
				"DR_REPLICATION_FACTOR": fmt.Sprintf("%v", req.ReplicationFactor),
				"DR_GAS_PRICE":          req.GasPrice,
				"DR_GAS_LIMIT":          fmt.Sprintf("%v", req.GasLimit),
				"DR_MEMO":               req.Memo,
				"DR_PAYBACK_ADDRESS":    req.PaybackAddress,
			})

			if tc.expErr != "" {
				require.Contains(t, result.Stderr[0], tc.expErr)
			} else {
				require.Equal(t, 0, len(result.Stderr))

				bz, err := hex.DecodeString(tc.args[0])
				require.NoError(t, err)
				require.Contains(t, string(result.Result), string(bz))
			}
		})
	}
}

// TestTallyVM_EnvVars tests passing environment variables to tally VM.
func TestTallyVM_EnvVars(t *testing.T) {
	cases := []struct {
		name        string
		requestJSON []byte
		args        []string
		expResult   string
		expErr      string
	}{
		{
			name: "Test passing all environment variables",
			requestJSON: []byte(`{
				"commits":{},
				"dr_binary_id":"9471d36add157cd7eaa32a42b5ddd091d5d5d396bf9ad67938a4fc40209df6cf",
				"dr_inputs":"",
				"gas_limit":5000000000,
				"gas_price":"10",
				"height":1661661742461173200,
				"id":"fba5314c57e52da7d1a2245d18c670fde1cb8c237062d2a1be83f449ace0932e",
				"memo":"mock_data_request_num_one",
				"payback_address":"YrzimoSJXwpA7ju71AkhkirkDCU=",
				"consensus_filter":"AQAAAAAAAAALcmVzdWx0LnRleHQ=",
				"replication_factor":3,
				"reveals":{},
				"seda_payload":"",
				"tally_binary_id":"5f3b31bff28c64a143119ee6389d62e38767672daace9c36db54fa2d18e9f391",
				"tally_inputs":"AAEBAQE=",
				"version":"1.0.0"
			}`),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var req types.Request
			err := json.Unmarshal(tc.requestJSON, &req)
			require.NoError(t, err)

			envs := map[string]string{
				"VM_MODE":               "tally",
				"CONSENSUS":             fmt.Sprintf("%v", true),
				"DR_ID":                 req.ID,
				"DR_INPUT":              req.DrInputs,
				"BINARY_ID":             req.DrBinaryID,
				"DR_REPLICATION_FACTOR": fmt.Sprintf("%v", req.ReplicationFactor),
				"DR_GAS_PRICE":          req.GasPrice,
				"DR_GAS_LIMIT":          fmt.Sprintf("%v", req.GasLimit),
				"DR_MEMO":               req.Memo,
				"DR_PAYBACK_ADDRESS":    req.PaybackAddress,
			}

			result := tallyvm.ExecuteTallyVm(testdata.SampleTallyWasm2(), tc.args, envs)

			require.Equal(t, 0, len(result.Stderr))
			for key := range envs {
				require.Contains(t, string(result.Result), fmt.Sprintf("%s=%s", key, envs[key]))
			}
		})
	}
}

// TestExecuteTally tests ExecuteTally using mock fetch data up to the
// point where VM execution results are posted to the contract.
func TestExecuteTally(t *testing.T) {
	f := initFixture(t)
	ctx := f.Context()

	// NOTE: consensus_filter = "AQAAAAAAAAALcmVzdWx0LnRleHQ="
	// represents mode filter with data path "result.text".
	tests := []struct {
		name       string
		resp       []byte
		wantErrStr []string
	}{
		{
			name: "None filter",
			resp: []byte(`[
				{
				   "consensus_filter":"AAEBAQE=",
				   "dr_binary_id":"9471d36add157cd7eaa32a42b5ddd091d5d5d396bf9ad67938a4fc40209df6cf",
				   "dr_inputs":"",
				   "gas_limit":5000000000,
				   "gas_price":"10",
				   "height":1661661742461173125,
				   "id":"fba5314c57e52da7d1a2245d18c670fde1cb8c237062d2a1be83f449ace0932e",
				   "memo":"",
				   "payback_address":"",
				   "replication_factor":3,
				   "commits":{
					  "1b85dfb9420e6757630a0db2280fa1787ec8c1e419a6aca76dbbfe8ef6e17521":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f",
					  "1dae290cd880b79d21079d89aee3460cf8a7d445fb35cade70cf8aa96924441c":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f",
					  "421e735518ef77fc1209a9d3585cdf096669b52ea68549e2ce048d4919b4c8c0":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"
				   },
				   "reveals":{
					  "1b85dfb9420e6757630a0db2280fa1787ec8c1e419a6aca76dbbfe8ef6e17521":{
						 "exit_code":0,
						 "gas_used":10,
						 "reveal":"Ng==",
						 "salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"
					  },
					  "1dae290cd880b79d21079d89aee3460cf8a7d445fb35cade70cf8aa96924441c":{
						 "exit_code":0,
						 "gas_used":10,
						 "reveal":"LQ==",
						 "salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"
					  },
					  "421e735518ef77fc1209a9d3585cdf096669b52ea68549e2ce048d4919b4c8c0":{
						 "exit_code":0,
						 "gas_used":10,
						 "reveal":"DQ==",
						 "salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"
					  }
				   },
				   "seda_payload":"",
				   "tally_binary_id":"8ade60039246740faa80bf424fc29e79fe13b32087043e213e7bc36620111f6b",
				   "tally_inputs":"bXktdGFsbHktaW5wdXRz",
				   "version":"1.0.0"
				},
				{
				   "consensus_filter":"AAEBAQE=",
				   "dr_binary_id":"9471d36add157cd7eaa32a42b5ddd091d5d5d396bf9ad67938a4fc40209df6cf",
				   "dr_inputs":"",
				   "gas_limit":5000000000,
				   "gas_price":"10",
				   "height":9859593541233596221,
				   "id":"d4e40f45fbf529134926acf529baeb6d4f37b5c380d7ab6b934833e7c00d725f",
				   "memo":"",
				   "payback_address":"",
				   "replication_factor":1,
				   "commits":{
				      "c9a4c8f1e70a0059a88b4768a920e41c95c587b8387ea3286d8fa4ee3b68b038":"f837455a930a66464f1c50586dc745a6b14ea807727c6069acac24c9558b6dbf"
				   },
				   "reveals":{
					  "c9a4c8f1e70a0059a88b4768a920e41c95c587b8387ea3286d8fa4ee3b68b038":{
						 "exit_code":0,
						 "gas_used":10,
						 "reveal":"Yw==",
						 "salt":"f837455a930a66464f1c50586dc745a6b14ea807727c6069acac24c9558b6dbf"
					  }
				   },
				   "seda_payload":"",
				   "tally_binary_id":"8ade60039246740faa80bf424fc29e79fe13b32087043e213e7bc36620111f6b",
				   "tally_inputs":"bXktdGFsbHktaW5wdXRz",
				   "version":"1.0.0"
				}
			 ]`),
			wantErrStr: []string{"seda_common::msgs::data_requests::types::DataRequest; key:", "not found"},
		},
		{
			name: "Mode filter",
			resp: []byte(`[
				{
				   "dr_binary_id":"9471d36add157cd7eaa32a42b5ddd091d5d5d396bf9ad67938a4fc40209df6cf",
				   "dr_inputs":"",
				   "gas_limit":5000000000,
				   "gas_price":"10",
				   "height":1661661742461173200,
				   "id":"fba5314c57e52da7d1a2245d18c670fde1cb8c237062d2a1be83f449ace0932e",
				   "memo":"mock_data_request_num_one",
				   "payback_address":"YrzimoSJXwpA7ju71AkhkirkDCU=",
				   "consensus_filter":"AQAAAAAAAAALcmVzdWx0LnRleHQ=",
				   "replication_factor":3,
				   "commits":{
				   	  "1b85dfb9420e6757630a0db2280fa1787ec8c1e419a6aca76dbbfe8ef6e17521":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f",
				   	  "1dae290cd880b79d21079d89aee3460cf8a7d445fb35cade70cf8aa96924441c":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f",
				   	  "421e735518ef77fc1209a9d3585cdf096669b52ea68549e2ce048d4919b4c8c0":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"
				   },
				   "reveals":{
					  "1b85dfb9420e6757630a0db2280fa1787ec8c1e419a6aca76dbbfe8ef6e17521":{
						 "exit_code":0,
						 "gas_used":10,
						 "reveal":"eyJyZXN1bHQiOiB7InRleHQiOiAiQSIsICJudW1iZXIiOiAxMH19",
						 "salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"
					  },
					  "1dae290cd880b79d21079d89aee3460cf8a7d445fb35cade70cf8aa96924441c":{
						 "exit_code":0,
						 "gas_used":10,
						 "reveal":"eyJyZXN1bHQiOiB7InRleHQiOiAiQSIsICJudW1iZXIiOiAyMH19",
						 "salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"
					  },
					  "421e735518ef77fc1209a9d3585cdf096669b52ea68549e2ce048d4919b4c8c0":{
						 "exit_code":0,
						 "gas_used":10,
						 "reveal":"eyJyZXN1bHQiOiB7InRleHQiOiAiQiIsICJudW1iZXIiOiAxMH19",
						 "salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"
					  }
				   },
				   "seda_payload":"",
				   "tally_binary_id":"5f3b31bff28c64a143119ee6389d62e38767672daace9c36db54fa2d18e9f391",
				   "tally_inputs":"AAEBAQE=",
				   "version":"1.0.0"
				},
				{
				   "dr_binary_id":"9471d36add157cd7eaa32a42b5ddd091d5d5d396bf9ad67938a4fc40209df6cf",
				   "dr_inputs":"",
				   "gas_limit":5000000000,
				   "gas_price":"10",
				   "height":9859593541233596000,
				   "id":"d4e40f45fbf529134926acf529baeb6d4f37b5c380d7ab6b934833e7c00d725f",
				   "memo":"mock_data_request_num_two",
				   "payback_address":"YrzimoSJXwpA7ju71AkhkirkDCU=",
				   "consensus_filter":"AQAAAAAAAAALcmVzdWx0LnRleHQ=",
				   "replication_factor":1,
				   "commits":{
				   	  "c9a4c8f1e70a0059a88b4768a920e41c95c587b8387ea3286d8fa4ee3b68b038":"f837455a930a66464f1c50586dc745a6b14ea807727c6069acac24c9558b6dbf"
				   },
				   "reveals":{
					  "c9a4c8f1e70a0059a88b4768a920e41c95c587b8387ea3286d8fa4ee3b68b038":{
						 "exit_code":0,
						 "gas_used":10,
						 "reveal":"eyJyZXN1bHQiOiB7InRleHQiOiAiQiIsICJudW1iZXIiOiAxMH19",
						 "salt":"f837455a930a66464f1c50586dc745a6b14ea807727c6069acac24c9558b6dbf"
					  }
				   },
				   "seda_payload":"",
				   "tally_binary_id":"5f3b31bff28c64a143119ee6389d62e38767672daace9c36db54fa2d18e9f391",
				   "tally_inputs":"AAEBAQE=",
				   "version":"1.0.0"
				}
			 ]`),
			wantErrStr: []string{"seda_common::msgs::data_requests::types::DataRequest; key:", "not found"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f.mockViewKeeper.EXPECT().QuerySmart(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.resp, nil)

			// Store the tally wasms.
			tallyWasm := wasmstoragetypes.NewOracleProgram(testdata.SampleTallyWasm(), ctx.BlockTime(), ctx.BlockHeight(), 100)
			err := f.wasmStorageKeeper.OracleProgram.Set(ctx, tallyWasm.Hash, tallyWasm)
			require.NoError(t, err)

			tallyWasm = wasmstoragetypes.NewOracleProgram(testdata.SampleTallyWasm2(), ctx.BlockTime(), ctx.BlockHeight(), 100)
			err = f.wasmStorageKeeper.OracleProgram.Set(ctx, tallyWasm.Hash, tallyWasm)
			require.NoError(t, err)

			// Contract should return not found in response to post data result
			// since the fetch data was mocked.
			err = f.tallyKeeper.ProcessTallies(ctx, f.coreContractAddr)
			if len(tt.wantErrStr) != 0 {
				for _, errStr := range tt.wantErrStr {
					require.Contains(t, err.Error(), errStr)
				}
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestEndBlock(t *testing.T) {
	f := initFixture(t)

	tests := []struct {
		name      string
		resp      []byte
		expErrLog []string
	}{
		{
			name:      "Invalid gas_limit and gas_used",
			resp:      []byte(`[{"commits":{"1b85dfb9420e6757630a0db2280fa1787ec8c1e419a6aca76dbbfe8ef6e17521":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f","1dae290cd880b79d21079d89aee3460cf8a7d445fb35cade70cf8aa96924441c":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f","421e735518ef77fc1209a9d3585cdf096669b52ea68549e2ce048d4919b4c8c0":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"},"dr_binary_id":"9471d36add157cd7eaa32a42b5ddd091d5d5d396bf9ad67938a4fc40209df6cf","dr_inputs":"","gas_limit":"10","gas_price":"10","height":1661661742461173200,"id":"fba5314c57e52da7d1a2245d18c670fde1cb8c237062d2a1be83f449ace0932e","memo":"","payback_address":"","consensus_filter":"A","replication_factor":3,"reveals":{"1b85dfb9420e6757630a0db2280fa1787ec8c1e419a6aca76dbbfe8ef6e17521":{"exit_code":0,"gas_used":"10","reveal":"eyJyZXN1bHQiOiB7InRleHQiOiAiQSIsICJudW1iZXIiOiAxMH19","salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"},"1dae290cd880b79d21079d89aee3460cf8a7d445fb35cade70cf8aa96924441c":{"exit_code":0,"gas_used":"10","reveal":"eyJyZXN1bHQiOiB7InRleHQiOiAiQSIsICJudW1iZXIiOiAyMH19","salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"},"421e735518ef77fc1209a9d3585cdf096669b52ea68549e2ce048d4919b4c8c0":{"exit_code":0,"gas_used":"10","reveal":"eyJyZXN1bHQiOiB7InRleHQiOiAiQiIsICJudW1iZXIiOiAxMH19","salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"}},"seda_payload":"","tally_binary_id":"8ade60039246740faa80bf424fc29e79fe13b32087043e213e7bc36620111f6b","tally_inputs":"AAEBAQE=","version":"1.0.0"},{"commits":{"c9a4c8f1e70a0059a88b4768a920e41c95c587b8387ea3286d8fa4ee3b68b038":"f837455a930a66464f1c50586dc745a6b14ea807727c6069acac24c9558b6dbf"},"dr_binary_id":"9471d36add157cd7eaa32a42b5ddd091d5d5d396bf9ad67938a4fc40209df6cf","dr_inputs":"","gas_limit":5000000000,"gas_price":"10","height":9859593541233596000,"id":"d4e40f45fbf529134926acf529baeb6d4f37b5c380d7ab6b934833e7c00d725f","memo":"","payback_address":"","consensus_filter":"AQAAAAAAAAALcmVzdWx0LnRleHQ=","replication_factor":1,"reveals":{"c9a4c8f1e70a0059a88b4768a920e41c95c587b8387ea3286d8fa4ee3b68b038":{"exit_code":0,"gas_used":"10","reveal":"eyJyZXN1bHQiOiB7InRleHQiOiAiQiIsICJudW1iZXIiOiAxMH19","salt":"f837455a930a66464f1c50586dc745a6b14ea807727c6069acac24c9558b6dbf"}},"seda_payload":"","tally_binary_id":"8ade60039246740faa80bf424fc29e79fe13b32087043e213e7bc36620111f6b","tally_inputs":"AAEBAQE=","version":"1.0.0"}]`),
			expErrLog: []string{"ERR", "json: cannot unmarshal string into Go struct field Request.gas_limit of type uint64"},
		},
		{
			name:      "Valid fetch format",
			resp:      []byte(`[{"commits":{"1b85dfb9420e6757630a0db2280fa1787ec8c1e419a6aca76dbbfe8ef6e17521":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f","1dae290cd880b79d21079d89aee3460cf8a7d445fb35cade70cf8aa96924441c":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f","421e735518ef77fc1209a9d3585cdf096669b52ea68549e2ce048d4919b4c8c0":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"},"dr_binary_id":"9471d36add157cd7eaa32a42b5ddd091d5d5d396bf9ad67938a4fc40209df6cf","dr_inputs":"","gas_limit":10,"gas_price":"10","height":1661661742461173200,"id":"fba5314c57e52da7d1a2245d18c670fde1cb8c237062d2a1be83f449ace0932e","memo":"","payback_address":"","consensus_filter":"A","replication_factor":3,"reveals":{"1b85dfb9420e6757630a0db2280fa1787ec8c1e419a6aca76dbbfe8ef6e17521":{"exit_code":0,"gas_used":10,"reveal":"eyJyZXN1bHQiOiB7InRleHQiOiAiQSIsICJudW1iZXIiOiAxMH19","salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"},"1dae290cd880b79d21079d89aee3460cf8a7d445fb35cade70cf8aa96924441c":{"exit_code":0,"gas_used":10,"reveal":"eyJyZXN1bHQiOiB7InRleHQiOiAiQSIsICJudW1iZXIiOiAyMH19","salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"},"421e735518ef77fc1209a9d3585cdf096669b52ea68549e2ce048d4919b4c8c0":{"exit_code":0,"gas_used":10,"reveal":"eyJyZXN1bHQiOiB7InRleHQiOiAiQiIsICJudW1iZXIiOiAxMH19","salt":"05952214b2ba3549a8d627c57d2d0dd1b0a2ce65c46e3b2f25c273464be8ba5f"}},"seda_payload":"","tally_binary_id":"8ade60039246740faa80bf424fc29e79fe13b32087043e213e7bc36620111f6b","tally_inputs":"AAEBAQE=","version":"1.0.0"},{"commits":{"c9a4c8f1e70a0059a88b4768a920e41c95c587b8387ea3286d8fa4ee3b68b038":"f837455a930a66464f1c50586dc745a6b14ea807727c6069acac24c9558b6dbf"},"dr_binary_id":"9471d36add157cd7eaa32a42b5ddd091d5d5d396bf9ad67938a4fc40209df6cf","dr_inputs":"","gas_limit":5000000000,"gas_price":"10","height":9859593541233596000,"id":"d4e40f45fbf529134926acf529baeb6d4f37b5c380d7ab6b934833e7c00d725f","memo":"","payback_address":"","consensus_filter":"AQAAAAAAAAALcmVzdWx0LnRleHQ=","replication_factor":1,"reveals":{"c9a4c8f1e70a0059a88b4768a920e41c95c587b8387ea3286d8fa4ee3b68b038":{"exit_code":0,"gas_used":10,"reveal":"eyJyZXN1bHQiOiB7InRleHQiOiAiQiIsICJudW1iZXIiOiAxMH19","salt":"f837455a930a66464f1c50586dc745a6b14ea807727c6069acac24c9558b6dbf"}},"seda_payload":"","tally_binary_id":"8ade60039246740faa80bf424fc29e79fe13b32087043e213e7bc36620111f6b","tally_inputs":"AAEBAQE=","version":"1.0.0"}]`),
			expErrLog: []string{"type: seda_common::msgs::data_requests::types::DataRequest; key: [00, 16, 64, 61, 74, 61, 5F, 72, 65, 71, 75, 65, 73, 74, 5F, 70, 6F, 6F, 6C, 5F, 72, 65, 71, 73, FB, A5, 31, 4C, 57, E5, 2D, A7, D1, A2, 24, 5D, 18, C6, 70, FD, E1, CB, 8C, 23, 70, 62, D2, A1, BE, 83, F4, 49, AC, E0, 93, 2E] not found: execute wasm contract failed"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f.mockViewKeeper.EXPECT().QuerySmart(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.resp, nil)

			err := f.tallyKeeper.EndBlock(f.Context())
			// Error must be handled by defer function and logged as an error.
			require.NoError(t, err)
			if len(tt.expErrLog) != 0 {
				for _, exp := range tt.expErrLog {
					fmt.Println(f.logBuf.String())
					require.Contains(t, f.logBuf.String(), exp)
				}
			}
		})
	}
}
