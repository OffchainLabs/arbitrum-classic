package inbox

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/pkg/errors"
	"math/big"
)

type JSONValue struct {
	Tuple *[]JSONValue `json:"Tuple,omitempty"`
	Int   *string      `json:"Int,omitempty"`
}

type TestVector struct {
	Version int         `json:"format_version"`
	Inbox   []JSONValue `json:"inbox"`
	Logs    []JSONValue `json:"logs"`
	Sends   []string    `json:"sends"`
}

func TestVectorJSON(inbox []InboxMessage, logs []value.Value, sends [][]byte) ([]byte, error) {
	jsonInbox := make([]JSONValue, 0, len(inbox))
	for _, msg := range inbox {
		val, err := valueToJSON(msg.AsValue())
		if err != nil {
			return nil, err
		}
		jsonInbox = append(jsonInbox, val)
	}
	jsonLogs := make([]JSONValue, 0, len(logs))
	for _, avmLog := range logs {
		val, err := valueToJSON(avmLog)
		if err != nil {
			return nil, err
		}
		jsonLogs = append(jsonLogs, val)
	}
	hexSends := make([]string, 0, len(sends))
	for _, avmSend := range sends {
		hexSends = append(hexSends, hexutil.Encode(avmSend))
	}
	vector := TestVector{
		Version: 1,
		Inbox:   jsonInbox,
		Logs:    jsonLogs,
		Sends:   hexSends,
	}
	return json.Marshal(vector)
}

func LoadTestVector(data []byte) ([]InboxMessage, []value.Value, [][]byte, error) {
	testVector := new(TestVector)
	if err := json.Unmarshal(data, testVector); err != nil {
		return nil, nil, nil, err
	}
	inboxMessages := make([]InboxMessage, 0, len(testVector.Inbox))
	for _, msg := range testVector.Inbox {
		val, err := jsonToValue(msg)
		if err != nil {
			return nil, nil, nil, err
		}
		im, err := NewInboxMessageFromValue(val)
		if err != nil {
			return nil, nil, nil, err
		}
		inboxMessages = append(inboxMessages, im)
	}
	avmLogs := make([]value.Value, 0, len(testVector.Logs))
	for _, avmLog := range testVector.Logs {
		val, err := jsonToValue(avmLog)
		if err != nil {
			return nil, nil, nil, err
		}
		avmLogs = append(avmLogs, val)
	}
	avmSends := make([][]byte, 0, len(testVector.Sends))
	for _, avmSend := range testVector.Sends {
		val, err := hexutil.Decode(avmSend)
		if err != nil {
			return nil, nil, nil, err
		}
		avmSends = append(avmSends, val)
	}
	return inboxMessages, avmLogs, avmSends, nil
}

func valueToJSON(val value.Value) (JSONValue, error) {
	switch val := val.(type) {
	case value.IntValue:
		intString := val.BigInt().Text(16)
		return JSONValue{Int: &intString}, nil
	case *value.TupleValue:
		vals := make([]JSONValue, 0)
		for _, subVal := range val.Contents() {
			jsonSubVal, err := valueToJSON(subVal)
			if err != nil {
				return JSONValue{}, err
			}
			vals = append(vals, jsonSubVal)
		}
		return JSONValue{Tuple: &vals}, nil
	default:
		return JSONValue{}, errors.New("unsupported type")
	}
}

func jsonToValue(val JSONValue) (value.Value, error) {
	if val.Int != nil {
		intVal, ok := new(big.Int).SetString(*val.Int, 16)
		if !ok {
			return nil, errors.New("invalid int value")
		}
		return value.NewIntValue(intVal), nil
	} else if val.Tuple != nil {
		vals := make([]value.Value, 0)
		for _, jsonSubVal := range *val.Tuple {
			subVal, err := jsonToValue(jsonSubVal)
			if err != nil {
				return nil, err
			}
			vals = append(vals, subVal)
		}
		return value.NewTupleFromSlice(vals)
	} else {
		return nil, errors.New("unsupported json value")
	}
}
