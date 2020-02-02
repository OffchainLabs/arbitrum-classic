package value

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type TestCase struct {
	Value string `json:"value"`
	Hash  string `json:"hash"`
	Name  string `json:"name"`
}

func TestTupleHash(t *testing.T) {
	jsonFile, err := os.Open("test_cases.json")
	if err != nil {
		t.Error(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var testCases []TestCase
	err = json.Unmarshal(byteValue, &testCases)
	if err != nil {
		t.Error(err)
	}
	for _, testCase := range testCases {
		test := testCase
		t.Run(testCase.Name, func(t *testing.T) {
			valBytes, err := hexutil.Decode("0x" + test.Value)
			if err != nil {
				t.Error(err)
			}
			val, err := UnmarshalValue(bytes.NewReader(valBytes))
			if err != nil {
				t.Error(err)
			}
			valHash := val.Hash()

			hashBytes, err := hexutil.Decode("0x" + test.Hash)
			if err != nil {
				t.Error(err)
			}

			if !bytes.Equal(valHash[:], hashBytes) {
				t.Errorf(
					"calculated wrong hash value: %v, expected hash value is: %v",
					valHash,
					hexutil.Encode(hashBytes),
				)
			}
		})
	}
}
