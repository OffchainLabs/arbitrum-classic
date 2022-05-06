/*
* Copyright 2021, Offchain Labs, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package evm

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type CallType int

const (
	Call         CallType = 0
	CallCode     CallType = 1
	DelegateCall CallType = 2
	StaticCall   CallType = 3
)

func (c CallType) String() string {
	switch c {
	case Call:
		return "Call"
	case CallCode:
		return "CallCode"
	case DelegateCall:
		return "DelegateCall"
	case StaticCall:
		return "StaticCall"
	default:
		return "Unknown"
	}
}

func (c CallType) RPCString() string {
	switch c {
	case Call:
		return "call"
	case CallCode:
		return "callcode"
	case DelegateCall:
		return "delegatecall"
	case StaticCall:
		return "staticcall"
	default:
		return "Unknown"
	}
}

type TraceItem interface {
	zerolog.LogObjectMarshaler
	fmt.Stringer
}

type CallTrace struct {
	Type     CallType
	Data     []byte
	Value    *big.Int
	From     common.Address
	To       *common.Address
	Gas      *big.Int
	GasPrice *big.Int
	PC       *uint64
}

func (t *CallTrace) String() string {
	return fmt.Sprintf(
		"%v(from=%v,to=%v,value=%v,gas=%v,gasPrice=%v,data=%v)",
		t.Type,
		t.From,
		t.To,
		t.Value,
		t.Gas,
		t.GasPrice,
		hexutil.Encode(t.Data),
	)
}

func (t *CallTrace) MarshalZerologObject(event *zerolog.Event) {
	event.
		Int("type", int(t.Type)).
		Str("from", t.From.Hex()).
		Str("value", t.Value.String()).
		Str("gas", t.Gas.String()).
		Str("gasPrice", t.GasPrice.String()).
		Hex("data", t.Data)
	if t.To != nil {
		event.Str("to", t.To.Hex())
	}
}

type ReturnTrace struct {
	Result     ResultType
	ReturnData []byte
	GasUsed    *big.Int
	PC         *uint64
}

func (t *ReturnTrace) String() string {
	return fmt.Sprintf("%v(gasUsed=%v, data=%v)", t.Result, t.GasUsed, hexutil.Encode(t.ReturnData))
}

func (t *ReturnTrace) MarshalZerologObject(event *zerolog.Event) {
	event.
		Int("result", int(t.Result)).
		Hex("data", t.ReturnData).
		Str("gasUsed", t.GasUsed.String())
}

type CreateTrace struct {
	Code            []byte
	ContractAddress common.Address
	PC              *uint64
}

func (t *CreateTrace) String() string {
	return fmt.Sprintf(
		"Create(contract=%v, code=%v)",
		t.ContractAddress,
		hexutil.Encode(t.Code),
	)
}

func (t *CreateTrace) MarshalZerologObject(event *zerolog.Event) {
	event.
		Hex("code", t.Code).
		Str("contract", t.ContractAddress.Hex())
}

type Create2Trace struct {
	Code            []byte
	Creator         common.Address
	Salt            *big.Int
	ContractAddress common.Address
	PC              *uint64
}

func (t *Create2Trace) String() string {
	return fmt.Sprintf(
		"Create2(contract=%v, creator=%v, salt=%v, code=%v)",
		t.ContractAddress,
		t.Creator,
		t.Salt,
		hexutil.Encode(t.Code),
	)
}

func (t *Create2Trace) MarshalZerologObject(event *zerolog.Event) {
	event.
		Hex("code", t.Code).
		Str("creator", t.Creator.Hex()).
		Str("salt", t.Salt.String()).
		Str("contract", t.ContractAddress.Hex())
}

func NewTraceFromDebugPrint(val value.Value) ([]TraceItem, error) {
	tup, ok := val.(*value.TupleValue)
	if !ok || tup.Len() != 2 {
		return nil, errors.New("expected trace to be tuple of length 2")
	}
	code, _ := tup.GetByInt64(0)
	traceData, _ := tup.GetByInt64(1)

	codeInt, ok := code.(value.IntValue)
	if !ok {
		return nil, errors.New("code must be an int")
	}
	if codeInt.BigInt().Cmp(big.NewInt(20000)) != 0 {
		return nil, errors.New("code must be 20000")
	}
	rawItems, err := inbox.StackValueToList(traceData)
	if err != nil {
		return nil, err
	}
	traceItems := make([]TraceItem, 0, len(rawItems))
	for i := range rawItems {
		item, err := newTraceItem(rawItems[len(rawItems)-1-i])
		if err != nil {
			return nil, err
		}
		traceItems = append(traceItems, item)
	}
	return traceItems, nil
}

func newTraceItem(val value.Value) (TraceItem, error) {
	tup, ok := val.(*value.TupleValue)
	if !ok || tup.Len() != 2 {
		return nil, errors.New("expected trace data to be tuple of length 2")
	}
	codeVal, _ := tup.GetByInt64(0)
	traceData, _ := tup.GetByInt64(1)
	codeInt, ok := codeVal.(value.IntValue)
	if !ok {
		return nil, errors.New("trace code must be an int")
	}
	code := codeInt.BigInt()
	if code.Cmp(big.NewInt(0)) == 0 {
		return newCallTraceItem(traceData)
	} else if code.Cmp(big.NewInt(1)) == 0 {
		return newReturnTraceItem(traceData)
	} else if code.Cmp(big.NewInt(2)) == 0 {
		return newCreateTraceItem(traceData)
	} else if code.Cmp(big.NewInt(3)) == 0 {
		return newCreate2TraceItem(traceData)
	} else {
		return nil, errors.New("unknown trace item type")
	}
}

func convertByteArrayToBytes(val value.Value) ([]byte, error) {
	tup, ok := val.(*value.TupleValue)
	if !ok || tup.Len() != 3 {
		return nil, errors.New("expected ByteArray to be tuple of length 3")
	}
	size, _ := tup.GetByInt64(0)
	offset, _ := tup.GetByInt64(1)
	buffer, _ := tup.GetByInt64(2)

	sizeInt, ok := size.(value.IntValue)
	if !ok {
		return nil, errors.New("bytearray size must be an int")
	}
	offsetInt, ok := offset.(value.IntValue)
	if !ok {
		return nil, errors.New("offset must be an int")
	}
	bufferBuf, ok := buffer.(*value.Buffer)
	if !ok {
		return nil, errors.New("buf must be a buffer")
	}
	return inbox.BufOffsetAndLengthToBytes(sizeInt.BigInt(), offsetInt.BigInt(), bufferBuf), nil
}

func parsePc(pcVal value.Value) (*uint64, error) {
	pcInt, ok := pcVal.(value.IntValue)
	if !ok {
		return nil, errors.New("to must be an int")
	}
	t := pcInt.BigInt()

	var pc *uint64
	if t.Cmp(math.MaxBig256) != 0 {
		v := t.Uint64()
		pc = &v
	}
	return pc, nil
}

func newCallTraceItem(val value.Value) (*CallTrace, error) {
	tup, ok := val.(*value.TupleValue)
	if !ok || tup.Len() != 8 {
		return nil, errors.New("expected call trace to be tuple of length 8")
	}
	callTypeVal, _ := tup.GetByInt64(0)
	callDataVal, _ := tup.GetByInt64(1)
	callValueVal, _ := tup.GetByInt64(2)
	fromVal, _ := tup.GetByInt64(3)
	toVal, _ := tup.GetByInt64(4)
	gasVal, _ := tup.GetByInt64(5)
	gasPriceVal, _ := tup.GetByInt64(6)
	pcVal, _ := tup.GetByInt64(7)

	callTypeInt, ok := callTypeVal.(value.IntValue)
	if !ok {
		return nil, errors.New("call type must be an int")
	}
	callType := CallType(callTypeInt.BigInt().Int64())
	callData, err := convertByteArrayToBytes(callDataVal)
	if err != nil {
		return nil, err
	}
	callValue, ok := callValueVal.(value.IntValue)
	if !ok {
		return nil, errors.New("call value must be an int")
	}
	fromInt, ok := fromVal.(value.IntValue)
	if !ok {
		return nil, errors.New("from must be an int")
	}
	from := inbox.NewAddressFromInt(fromInt)
	toInt, ok := toVal.(value.IntValue)
	if !ok {
		return nil, errors.New("to must be an int")
	}
	toRaw := inbox.NewAddressFromInt(toInt)
	var to *common.Address
	emptyAddress := common.Address{}
	if toRaw != emptyAddress {
		to = &toRaw
	}
	gas, ok := gasVal.(value.IntValue)
	if !ok {
		return nil, errors.New("gas must be an int")
	}
	gasPrice, ok := gasPriceVal.(value.IntValue)
	if !ok {
		return nil, errors.New("gas price must be an int")
	}
	pc, err := parsePc(pcVal)
	if err != nil {
		return nil, err
	}
	return &CallTrace{
		Type:     callType,
		Data:     callData,
		Value:    callValue.BigInt(),
		From:     from,
		To:       to,
		Gas:      gas.BigInt(),
		GasPrice: gasPrice.BigInt(),
		PC:       pc,
	}, nil
}

func newReturnTraceItem(val value.Value) (*ReturnTrace, error) {
	tup, ok := val.(*value.TupleValue)
	if !ok || tup.Len() != 4 {
		return nil, errors.New("expected return trace to be tuple of length 4")
	}
	resultVal, _ := tup.GetByInt64(0)
	returnDataVal, _ := tup.GetByInt64(1)
	gasUsedVal, _ := tup.GetByInt64(2)
	pcVal, _ := tup.GetByInt64(3)

	resultInt, ok := resultVal.(value.IntValue)
	if !ok {
		return nil, errors.New("result must be an int")
	}
	returnData, err := convertByteArrayToBytes(returnDataVal)
	if err != nil {
		return nil, err
	}
	gasUsed, ok := gasUsedVal.(value.IntValue)
	if !ok {
		return nil, errors.New("gas used must be an int")
	}
	pc, err := parsePc(pcVal)
	if err != nil {
		return nil, err
	}

	return &ReturnTrace{
		Result:     ResultType(resultInt.BigInt().Int64()),
		ReturnData: returnData,
		GasUsed:    gasUsed.BigInt(),
		PC:         pc,
	}, nil
}

func newCreateTraceItem(val value.Value) (*CreateTrace, error) {
	tup, ok := val.(*value.TupleValue)
	if !ok || tup.Len() != 3 {
		return nil, errors.New("expected create trace to be tuple of length 3")
	}
	codeVal, _ := tup.GetByInt64(0)
	contractAddressVal, _ := tup.GetByInt64(1)
	pcVal, _ := tup.GetByInt64(2)

	code, err := convertByteArrayToBytes(codeVal)
	if err != nil {
		return nil, err
	}
	contractAddress, ok := contractAddressVal.(value.IntValue)
	if !ok {
		return nil, errors.New("contract address must be an int")
	}
	pc, err := parsePc(pcVal)
	if err != nil {
		return nil, err
	}

	return &CreateTrace{
		Code:            code,
		ContractAddress: inbox.NewAddressFromInt(contractAddress),
		PC:              pc,
	}, nil
}

func newCreate2TraceItem(val value.Value) (*Create2Trace, error) {
	tup, ok := val.(*value.TupleValue)
	if !ok || tup.Len() != 5 {
		return nil, errors.New("expected create2 trace to be tuple of length 5")
	}
	codeVal, _ := tup.GetByInt64(0)
	creatorAddressVal, _ := tup.GetByInt64(1)
	saltVal, _ := tup.GetByInt64(2)
	contractAddressVal, _ := tup.GetByInt64(3)
	pcVal, _ := tup.GetByInt64(4)

	code, err := convertByteArrayToBytes(codeVal)
	if err != nil {
		return nil, err
	}
	creatorAddress, ok := creatorAddressVal.(value.IntValue)
	if !ok {
		return nil, errors.New("creator address must be an int")
	}
	salt, ok := saltVal.(value.IntValue)
	if !ok {
		return nil, errors.New("salt must be an int")
	}
	contractAddress, ok := contractAddressVal.(value.IntValue)
	if !ok {
		return nil, errors.New("contract address must be an int")
	}
	pc, err := parsePc(pcVal)
	if err != nil {
		return nil, err
	}

	return &Create2Trace{
		Code:            code,
		Creator:         inbox.NewAddressFromInt(creatorAddress),
		Salt:            salt.BigInt(),
		ContractAddress: inbox.NewAddressFromInt(contractAddress),
		PC:              pc,
	}, nil
}

type Frame interface {
	GetCallFrame() *CallFrame
}

type CallFrame struct {
	Call   *CallTrace
	Return *ReturnTrace
	Nested []Frame
}

func (f *CallFrame) GetCallFrame() *CallFrame {
	return f
}

type CreateFrame struct {
	Create *CreateTrace
	*CallFrame
}

type Create2Frame struct {
	Create *Create2Trace
	*CallFrame
}

type EVMTrace struct {
	Items []TraceItem
}

func (e *EVMTrace) String() string {
	builder := &strings.Builder{}
	builder.WriteString("Tx trace:")
	for _, item := range e.Items {
		builder.WriteString("\n")
		builder.WriteString(item.String())
	}
	return builder.String()
}

func (e *EVMTrace) MarshalZerologObject(event *zerolog.Event) {
	array := zerolog.Arr()
	for _, item := range e.Items {
		array = array.Object(item)
	}
	event.Array("items", array)
}

func (e *EVMTrace) FrameTree() (Frame, error) {
	if len(e.Items) == 0 {
		return nil, nil
	}
	items := e.Items
	var frames []Frame
	for i := 0; i < len(items); i++ {
		getCallAfterCreate := func() (*CallTrace, error) {
			if i+1 >= len(items) {
				return nil, errors.New("expected item after create")
			}
			i++
			createCall, ok := items[i].(*CallTrace)
			if !ok {
				return nil, errors.New("expected call after create")
			}
			return createCall, nil
		}
		switch item := items[i].(type) {
		case *CreateTrace:
			createCall, err := getCallAfterCreate()
			if err != nil {
				return nil, err
			}
			frames = append(frames, &CreateFrame{
				Create: item,
				CallFrame: &CallFrame{
					Call: createCall,
				},
			})
		case *Create2Trace:
			createCall, err := getCallAfterCreate()
			if err != nil {
				return nil, err
			}
			frames = append(frames, &Create2Frame{
				Create: item,
				CallFrame: &CallFrame{
					Call: createCall,
				},
			})
		case *CallTrace:
			if item.Type == DelegateCall {
				var from *common.Address
				for i := 0; i < len(frames); i++ {
					prevFrame := frames[len(frames)-1-i]
					prevCall := prevFrame.GetCallFrame().Call
					if prevCall.Type != DelegateCall && prevCall.Type != CallCode {
						from = prevCall.To
						break
					}
				}
				if from == nil {
					return nil, errors.New("expected to find non-delegatecall or callcode in call tree")
				}
				item.From = *from
			}
			frames = append(frames, &CallFrame{
				Call: item,
			})
		case *ReturnTrace:
			if len(frames) == 0 {
				return nil, errors.New("returned while not in call")
			}
			frames[len(frames)-1].GetCallFrame().Return = item
			if len(frames) == 1 {
				if i != len(items)-1 {
					return nil, errors.New("finished")
				}
				return frames[0], nil
			} else {
				parent := frames[len(frames)-2].GetCallFrame()
				parent.Nested = append(parent.Nested, frames[len(frames)-1])
				frames = frames[:len(frames)-1]
			}
		}
	}
	return nil, errors.New("expected to end on return")
}
