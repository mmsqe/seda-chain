package keeper_test

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sedaprotocol/seda-chain/x/wasm-storage/keeper"
	"github.com/sedaprotocol/seda-chain/x/wasm-storage/types"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name            string
		tallyInputAsHex string
		outliers        []int
		reveals         []types.RevealBody
		consensus       bool
		wantErr         error
	}{
		{
			name:            "None filter",
			tallyInputAsHex: "00",
			outliers:        []int{0, 0, 0, 0, 0},
			reveals: []types.RevealBody{
				{},
				{},
				{},
				{},
				{},
			},
			consensus: true,
			wantErr:   nil,
		},
		{
			name:            "Mode filter - Happy Path",
			tallyInputAsHex: "01000000000000000b726573756C742E74657874", // json_path = result.text
			outliers:        []int{0, 0, 1, 0, 1, 0, 0},
			reveals: []types.RevealBody{
				{Reveal: `{"high_level_prop1":"ignore this", "result": {"text": "A", "number": 0}}`},
				{Reveal: `{"makes_this_json":"ignore this", "result": {"text": "A", "number": 10}}`},
				{Reveal: `{"unstructured":"ignore this", "result": {"text": "B", "number": 101}}`},
				{Reveal: `{"but":"ignore this", "result": {"text": "A", "number": 10}}`},
				{Reveal: `{"it_does_not":"ignore this", "result": {"text": "C", "number": 10}}`},
				{Reveal: `{"matter":"ignore this", "result": {"text": "A", "number": 10}}`},
				{Reveal: `{"matter":"ignore this", "result": {"text": "A", "number": 10}}`},
			},
			consensus: true,
			wantErr:   nil,
		},
		{
			name:            "Mode filter - One outlier but consensus",
			tallyInputAsHex: "01000000000000000b726573756C742E74657874", // json_path = result.text
			outliers:        []int{0, 0, 1},
			reveals: []types.RevealBody{
				{Reveal: `{"result": {"text": "A", "number": 0}}`},
				{Reveal: `{"result": {"text": "A", "number": 10}}`},
				{Reveal: `{"result": {"text": "B", "number": 101}}`},
			},
			consensus: true,
			wantErr:   nil,
		},
		{
			name:            "Mode filter - Multiple modes",
			tallyInputAsHex: "01000000000000000b726573756C742E74657874", // json_path = result.text
			outliers:        []int{0, 0, 0, 0, 0, 0, 1},
			reveals: []types.RevealBody{
				{Reveal: `{"result": {"text": "A"}`},
				{Reveal: `{"result": {"text": "A"}`},
				{Reveal: `{"result": {"text": "A"}`},
				{Reveal: `{"result": {"text": "B"}`},
				{Reveal: `{"result": {"text": "B"}`},
				{Reveal: `{"result": {"text": "B"}`},
				{Reveal: `{"result": {"text": "C"}`},
			},
			consensus: false,
			wantErr:   nil,
		},
		{
			name:            "Mode filter - One corrupt reveal but consensus",
			tallyInputAsHex: "01000000000000000b726573756C742E74657874", // json_path = result.text
			outliers:        []int{0, 1, 0},
			reveals: []types.RevealBody{
				{Reveal: `{"result": {"text": "A", "number": 0}}`},
				{Reveal: `{"resultt": {"text": "A", "number": 10}}`},
				{Reveal: `{"result": {"text": "A", "number": 101}}`},
			},
			consensus: true,
			wantErr:   nil,
		},
		{
			name:            "Mode filter - Consensus due to non exit code",
			tallyInputAsHex: "01000000000000000b726573756C742E74657874", // json_path = result.text
			outliers:        []int{0, 0, 0, 0, 0, 0},
			reveals: []types.RevealBody{
				{
					ExitCode: 1,
					Reveal:   `{"high_level_prop1":"ignore this", "result": {"text": "A", "number": 0}}`,
				},
				{
					ExitCode: 1,
					Reveal:   `{"makes_this_json":"ignore this", "result": {"text": "A", "number": 10}}`,
				},
				{
					ExitCode: 1,
					Reveal:   `{"unstructured":"ignore this", "result": {"text": "B", "number": 101}}`,
				},
				{Reveal: `{"but":"ignore this", "result": {"text": "B", "number": 10}}`},
				{Reveal: `{"it_does_not":"ignore this", "result": {"text": "C", "number": 10}}`},
				{Reveal: `{"matter":"ignore this", "result": {"text": "C", "number": 10}}`},
			},
			consensus: false,
			wantErr:   types.ErrCorruptReveals,
		},
		{
			name:            "Mode filter - Valid reveal marked outlier due to non exit code [still consensus]",
			tallyInputAsHex: "01000000000000000b726573756C742E74657874", // json_path = result.text
			outliers:        []int{1, 0, 0, 1, 0, 0, 0},
			reveals: []types.RevealBody{
				{
					ExitCode: 1,
					Reveal:   `{"xx":"ignore this", "result": {"text": "A", "number": 0}}`,
				},
				{Reveal: `{"xx":"ignore this", "result": {"text": "A", "number": 10}}`},
				{Reveal: `{"xx":"ignore this", "result": {"text": "A", "number": 101}}`},
				{Reveal: `{"xx":"ignore this", "result": {"text": "B", "number": 10}}`},
				{Reveal: `{"xx":"ignore this", "result": {"text": "A", "number": 10}}`},
				{Reveal: `{"xx":"ignore this", "result": {"text": "A", "number": 10}}`},
				{Reveal: `{"xx":"ignore this", "result": {"text": "A", "number": 10}}`},
			},
			consensus: true,
			wantErr:   nil,
		},
		{
			name:            "Mode filter - Consensus not reached due to exit code",
			tallyInputAsHex: "01000000000000000b726573756C742E74657874", // json_path = result.text
			outliers:        []int{1, 0, 0, 1, 1, 0},
			reveals: []types.RevealBody{
				{
					ExitCode: 1,
					Reveal:   `{"result": {"text": "A", "number": 0}}`,
				},
				{Reveal: `{"result": {"text": "A", "number": 10}}`},
				{Reveal: `{"result": {"text": "A", "number": 101}}`},
				{Reveal: `{"result": {"text": "B", "number": 10}}`},
				{Reveal: `{"result": {"text": "C", "number": 10}}`},
				{Reveal: `{"result": {"text": "A", "number": 10}}`},
			},
			consensus: false,
			wantErr:   nil,
		},
		{
			name:            "Mode filter - Consensus not reached due to corrupt reveal",
			tallyInputAsHex: "01000000000000000b726573756C742E74657874", // json_path = result.text
			outliers:        []int{1, 0, 0, 1, 1, 0},
			reveals: []types.RevealBody{
				{Reveal: `{"resalt": {"text": "A", "number": 0}}`},
				{Reveal: `{"result": {"text": "A", "number": 10}}`},
				{Reveal: `{"result": {"text": "A", "number": 101}}`},
				{Reveal: `{"result": {"text": "B", "number": 10}}`},
				{Reveal: `{"result": {"text": "C", "number": 10}}`},
				{Reveal: `{"result": {"text": "A", "number": 10}}`},
			},
			consensus: false,
			wantErr:   nil,
		},
		{
			name:            "Standard deviation filter uint64",
			tallyInputAsHex: "02000000000016E36003000000000000000B726573756C742E74657874", // max_sigma = 1.5, number_type = uint64, json_path = result.text
			outliers:        []int{1, 0, 0, 0, 0, 1},
			reveals: []types.RevealBody{
				{Reveal: `{"result": {"text": "AAAAAAAAAAQ=", "number": 0}}`},   // 4
				{Reveal: `{"result": {"text": "AAAAAAAAAAU=", "number": 10}}`},  // 5
				{Reveal: `{"result": {"text": "AAAAAAAAAAY=", "number": 101}}`}, // 6
				{Reveal: `{"result": {"text": "AAAAAAAAAAc=", "number": 0}}`},   // 7
				{Reveal: `{"result": {"text": "AAAAAAAAAAg=", "number": 0}}`},   // 8
				{Reveal: `{"result": {"text": "AAAAAAAAAAk=", "number": 0}}`},   // 9

			},
			consensus: true,
			wantErr:   nil,
		},
		{
			name:            "Standard deviation filter int64",
			tallyInputAsHex: "02000000000016E36001000000000000000b726573756C742E74657874", // max_sigma = 1.5, number_type = int64, json_path = result.text
			outliers:        []int{1, 0, 0, 0, 0, 1},
			reveals: []types.RevealBody{
				{Reveal: `{"result": {"text": "AAAAAAAAAAQ=", "number": 0}}`},   // 4
				{Reveal: `{"result": {"text": "AAAAAAAAAAU=", "number": 10}}`},  // 5
				{Reveal: `{"result": {"text": "AAAAAAAAAAY=", "number": 101}}`}, // 6
				{Reveal: `{"result": {"text": "AAAAAAAAAAc=", "number": 0}}`},   // 7
				{Reveal: `{"result": {"text": "AAAAAAAAAAg=", "number": 0}}`},   // 8
				{Reveal: `{"result": {"text": "AAAAAAAAAAk=", "number": 0}}`},   // 9
			},
			consensus: true,
			wantErr:   nil,
		},
		{
			name:            "Standard deviation filter - Empty reveal",
			tallyInputAsHex: "02000000000016E36001000000000000000b726573756C742E74657874", // max_sigma = 1.5, number_type = uint64, json_path = result.text
			outliers:        []int{},
			reveals:         []types.RevealBody{},
			consensus:       false,
			wantErr:         types.ErrEmptyReveals,
		},
		{
			name:            "Standard deviation filter - Single reveal",
			tallyInputAsHex: "02000000000016E36001000000000000000b726573756C742E74657874", // max_sigma = 1.5, number_type = uint64, json_path = result.text
			outliers:        []int{0},
			reveals: []types.RevealBody{
				{Reveal: `{"result": {"text": "AAAAAAAAAAQ=", "number": 0}}`}, // 4
			},
			consensus: true,
			wantErr:   nil,
		},
		{
			name:            "Standard deviation filter - One corrupt reveal",
			tallyInputAsHex: "02000000000016E36001000000000000000b726573756C742E74657874", // max_sigma = 1.5, number_type = uint64, json_path = result.text
			outliers:        []int{1, 0, 0, 0, 1, 1},
			reveals: []types.RevealBody{
				{Reveal: `{"result": {"text": "AAAAAAAAAAQ=", "number": 0}}`},   // 4
				{Reveal: `{"result": {"text": "AAAAAAAAAAU=", "number": 10}}`},  // 5
				{Reveal: `{"result": {"text": "AAAAAAAAAAY=", "number": 101}}`}, // 6
				{Reveal: `{"result": {"text": "AAAAAAAAAAc=", "number": 0}}`},   // 7
				{Reveal: `{"result": {"number": 0}}`},                           // corrupt
				{Reveal: `{"result": {"text": "AAAAAAAAAAk=", "number": 0}}`},   // 9
			},
			consensus: false,
			wantErr:   nil,
		},
		{
			name:            "Standard deviation filter - Max sigma 1.55",
			tallyInputAsHex: "02000000000017A6B003000000000000000b726573756C742E74657874", // max_sigma = 1.55, number_type = uint64, json_path = result.text
			outliers:        []int{1, 0, 0, 0, 0, 1},
			reveals: []types.RevealBody{
				{Reveal: `{"result": {"text": "AAAAAAAAAAQ=", "number": 0}}`},   // 4
				{Reveal: `{"result": {"text": "AAAAAAAAAAU=", "number": 10}}`},  // 5
				{Reveal: `{"result": {"text": "AAAAAAAAAAY=", "number": 101}}`}, // 6
				{Reveal: `{"result": {"text": "AAAAAAAAAAc=", "number": 0}}`},   // 7
				{Reveal: `{"result": {"text": "AAAAAAAAAAg=", "number": 0}}`},   // 8
				{Reveal: `{"result": {"text": "AAAAAAAAAAk=", "number": 0}}`},   // 9
			},
			consensus: true,
			wantErr:   nil,
		},
		{
			name:            "Standard deviation filter - Max sigma 1.45",
			tallyInputAsHex: "02000000000016201003000000000000000b726573756C742E74657874", // max_sigma = 1.45, number_type = uint64, json_path = result.text
			outliers:        []int{1, 1, 0, 0, 1, 1},
			reveals: []types.RevealBody{
				{Reveal: `{"result": {"text": "AAAAAAAAAAQ=", "number": 0}}`},   // 4
				{Reveal: `{"result": {"text": "AAAAAAAAAAU=", "number": 10}}`},  // 5
				{Reveal: `{"result": {"text": "AAAAAAAAAAY=", "number": 101}}`}, // 6
				{Reveal: `{"result": {"text": "AAAAAAAAAAc=", "number": 0}}`},   // 7
				{Reveal: `{"result": {"text": "AAAAAAAAAAg=", "number": 0}}`},   // 8
				{Reveal: `{"result": {"text": "AAAAAAAAAAk=", "number": 0}}`},   // 9
			},
			consensus: false,
			wantErr:   nil,
		},
		{
			name:            "Standard deviation filter int64 with negative reveals",
			tallyInputAsHex: "02000000000016E36001000000000000000b726573756C742E74657874", // max_sigma = 1.5, number_type = int64, json_path = result.text
			outliers:        []int{1, 0, 0, 0, 0, 1},
			reveals: []types.RevealBody{
				{Reveal: `{"result": {"text": "//////////w=", "number": 0}}`},   // -4
				{Reveal: `{"result": {"text": "//////////s=", "number": 10}}`},  // -5
				{Reveal: `{"result": {"text": "//////////o=", "number": 101}}`}, // -6
				{Reveal: `{"result": {"text": "//////////k=", "number": 0}}`},   // -7
				{Reveal: `{"result": {"text": "//////////g=", "number": 0}}`},   // -8
				{Reveal: `{"result": {"text": "//////////c=", "number": 0}}`},   // -9
			},
			consensus: true,
			wantErr:   nil,
		},
		{
			name:            "Standard deviation filter int64 median -0.5",
			tallyInputAsHex: "02000000000007A12001000000000000000b726573756C742E74657874", // max_sigma = 0.5, number_type = int64, json_path = result.text
			outliers:        []int{1, 0, 0, 1},
			reveals: []types.RevealBody{
				{Reveal: `{"result": {"text": "AAAAAAAAAAE=", "number": 0}}`},  // 1
				{Reveal: `{"result": {"text": "AAAAAAAAAAA=", "number": 0}}`},  // 0
				{Reveal: `{"result": {"text": "//////////8=", "number": 10}}`}, // -1
				{Reveal: `{"result": {"text": "//////////4=", "number": 10}}`}, // -2
			},
			consensus: false,
			wantErr:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter, err := hex.DecodeString(tt.tallyInputAsHex)
			require.NoError(t, err)

			// For illustration
			for i := 0; i < len(tt.reveals); i++ {
				tt.reveals[i].Reveal = base64.StdEncoding.EncodeToString([]byte(tt.reveals[i].Reveal))
			}

			outliers, cons, err := keeper.ApplyFilter(filter, tt.reveals)
			require.ErrorIs(t, err, tt.wantErr)
			require.Equal(t, tt.outliers, outliers)
			require.Equal(t, tt.consensus, cons)
		})
	}
}
