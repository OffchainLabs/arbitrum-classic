package inbox

import (
	"encoding/json"
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type JSONValue interface {
}

type JSONTuple struct {
	Tuple []JSONValue `json:"Tuple"`
}

type JSONInt struct {
	Int string `json:"Int"`
}

type TestVector struct {
	Version int         `json:"format_version"`
	Inbox   []JSONValue `json:"inbox"`
	Logs    []JSONValue `json:"logs"`
	Sends   []JSONValue `json:"sends"`
}

func TestVectorJSON(inbox []InboxMessage, logs []value.Value, sends []value.Value) ([]byte, error) {
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
	jsonSends := make([]JSONValue, 0, len(sends))
	for _, avmSend := range sends {
		val, err := valueToJSON(avmSend)
		if err != nil {
			return nil, err
		}
		jsonSends = append(jsonSends, val)
	}
	vector := TestVector{
		Version: 1,
		Inbox:   jsonInbox,
		Logs:    jsonLogs,
		Sends:   jsonSends,
	}
	return json.Marshal(vector)
}

func valueToJSON(val value.Value) (JSONValue, error) {
	switch val := val.(type) {
	case value.IntValue:
		return JSONInt{Int: val.BigInt().Text(16)}, nil
	case value.TupleValue:
		vals := make([]JSONValue, 0)
		for _, subVal := range val.Contents() {
			jsonSubVal, err := valueToJSON(subVal)
			if err != nil {
				return nil, err
			}
			vals = append(vals, jsonSubVal)
		}
		return JSONTuple{Tuple: vals}, nil
	default:
		return nil, errors.New("unsupported type")
	}
}
