/*
* Copyright 2020-2021, Offchain Labs, Inc.
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
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type callFrameAddresses struct {
	executor common.Address
	code     common.Address
}

func (cf callFrameAddresses) String() string {
	return fmt.Sprintf("{executor: %v, code: %v}", cf.executor, cf.code)
}

func (cf callFrameAddresses) MarshalZerologObject(e *zerolog.Event) {
	e.Str("executor", cf.executor.Hex()).Str("code", cf.code.Hex())
}

func (cf callFrameAddresses) same() bool {
	return cf.executor == cf.code
}

func getAddresses(val value.Value) (callFrameAddresses, error) {
	tup, ok := val.(*value.TupleValue)
	if !ok || tup.Len() != 2 {
		return callFrameAddresses{}, errors.New("addresses wasn't a 2-tuple")
	}
	executorAddressVal, _ := tup.GetByInt64(0)
	codeAddressVal, _ := tup.GetByInt64(1)

	executorAddressInt, ok := executorAddressVal.(value.IntValue)
	if !ok {
		return callFrameAddresses{}, errors.New("expected first address to be int")
	}

	codeAddressInt, ok := codeAddressVal.(value.IntValue)
	if !ok {
		return callFrameAddresses{}, errors.New("expected second address to be int")
	}
	cf := callFrameAddresses{
		executor: inbox.NewAddressFromInt(executorAddressInt),
		code:     inbox.NewAddressFromInt(codeAddressInt),
	}
	return cf, nil
}

type DebugPrintLog struct {
	txId    common.Hash
	current callFrameAddresses
	parent  callFrameAddresses
	kind    string
	pc      *uint64
}

func (d *DebugPrintLog) String() string {
	ret := ""
	ret += fmt.Sprintf("DebugPrintLog{txId: %v, kind: %v", d.txId, d.kind)
	if d.current.same() {
		ret += fmt.Sprintf(", current: %v", d.current.code)
	} else {
		ret += fmt.Sprintf(", current: %v", d.current)
	}
	if d.parent.same() {
		emptyAddress := common.Address{}
		if d.parent.code != emptyAddress {
			ret += fmt.Sprintf(", parent: %v", d.parent.code)
		}
	} else {
		ret += fmt.Sprintf(", parent: %v", d.parent)
	}
	if d.pc != nil {
		ret += fmt.Sprintf(", pc: %v", *d.pc)
	}
	ret += "}"
	return ret
}

func (d *DebugPrintLog) MarshalZerologObject(e *zerolog.Event) {
	e.Hex("tx_id", d.txId[:]).Str("kind", d.kind)
	if d.current.same() {
		e.Str("current", d.current.code.Hex())
	} else {
		e.Object("current", d.current)
	}

	if d.parent.same() {
		emptyAddress := common.Address{}
		if d.parent.code != emptyAddress {
			e.Str("parent", d.parent.code.Hex())
		}
	} else {
		e.Object("parent", d.parent)
	}
	if d.pc != nil {
		e.Uint64("pc", *d.pc)
	}
}

func generateLog(txID, currentFrame, parentFrame value.Value, kind string, pc *uint64) (*DebugPrintLog, error) {
	txIDInt, ok := txID.(value.IntValue)
	if !ok {
		return nil, errors.New("expected txid to be int")
	}
	current, err := getAddresses(currentFrame)
	if err != nil {
		return nil, err
	}
	parent, err := getAddresses(parentFrame)
	if err != nil {
		return nil, err
	}
	return &DebugPrintLog{
		txId:    txIDInt.ToBytes(),
		current: current,
		parent:  parent,
		kind:    kind,
		pc:      pc,
	}, nil
}

func decodeEVMCallError(errorCode uint64) string {
	switch errorCode {
	case 0:
		return "application code error"
	case 1:
		return "Failed to transfer eth balance to contract: insufficient balance or unknown error"
	case 2:
		return "Can't pay for gas to contract"
	case 3:
		return "Failed to transfer eth balance to EOA"
	case 4:
		return "Can't pay for gas in constructor"
	case 5:
		return "Somehow the constructor didn't have storage"
	case 7:
		return "Should never reach end of call entry function"
	case 8:
		return "Should never reach end of call return function"
	case 9:
		return "Should never reach end of call return function"
	case 10:
		return "Called evmCallStack_getTopFrameMemoryOrDie while not in global stack frame"
	case 11:
		return "EVM code tried to jump to a forbidden EVM jump destination"
	case 12:
		return "Shouldn't reach at end of evmOp_getjumpaddr function"
	case 14:
		return "Called evmCallStack_queueMessage while not in global stack frame"
	case 15:
		return "Can't pay for gas in constructor"
	case 17:
		return "Called arbAddressTable_txcall while not in EVM tx"
	case 18:
		return "Called arbBLS_txcall or arbFunctionTable_txcall while not in EVM tx"
	case 19:
		return "Called arbosTest_txcall while not in EVM tx or error in snapshotAuxStack or restoreAuxStackAndCall"
	case 20:
		return "Chain hasn't been initialized"
	case 21:
		return "Called arbsys_txcall while not in EVM tx"
	case 22:
		return "Called arbowner_txcall while not in EVM tx"
	case 23:
		return "generateCodeForEvmSegment pushN without data"
	default:
		return "unknown"
	}
}

type EVMCallError struct {
	description string
	errorCode   uint64
}

func (e *EVMCallError) String() string {
	return fmt.Sprintf("EVMCallError{description: %v, errorCode: %v}", e.description, e.errorCode)
}

func (e *EVMCallError) MarshalZerologObject(event *zerolog.Event) {
	event.
		Str("description", e.description).
		Str("kind", "evm_call_error").
		Uint64("error_code", e.errorCode)
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

type ErrorHandlerError struct {
}

func (e *ErrorHandlerError) String() string {
	return fmt.Sprintf("ErrorHandlerError{error_in_error_handler}")
}

func (e *ErrorHandlerError) MarshalZerologObject(event *zerolog.Event) {
	event.Str("kind", "error_in_error_handler")
}

type RawDebugPrint struct {
	Val value.Value
}

func (r *RawDebugPrint) String() string {
	return fmt.Sprintf("RawDebugPrint{%v}", r.Val)
}

func (r *RawDebugPrint) MarshalZerologObject(event *zerolog.Event) {
	event.
		Str("kind", "raw").
		Str("value", r.Val.String())
}

type EVMLogLine interface {
	zerolog.LogObjectMarshaler
}

func NewLogLineFromValue(d value.Value) (EVMLogLine, error) {
	tup, ok := d.(*value.TupleValue)
	if !ok || tup.Len() == 0 {
		return nil, errors.New("expected debugprint to be tuple")
	}
	// Tuple already checked to be at least size 1
	debugPrintType, _ := tup.GetByInt64(0)
	debugPrintTypeInt, ok := debugPrintType.(value.IntValue)
	if !ok {
		return nil, errors.New("expected debugprint typecode to be int")
	}
	typ := debugPrintTypeInt.BigInt().Uint64()
	if typ == 664 {
		if tup.Len() != 2 {
			return nil, errors.New("expected type 664 to be 2-tuple")
		}
		subCodeVal, _ := tup.GetByInt64(1)
		subCodeInt, ok := subCodeVal.(value.IntValue)
		if !ok {
			return nil, errors.New("expected type 664 to have subcode")
		}
		errorCode := subCodeInt.BigInt().Uint64()
		errorStr := decodeEVMCallError(errorCode)
		return &EVMCallError{
			description: errorStr,
			errorCode:   errorCode,
		}, nil
	} else if typ == 666 && tup.Len() == 1 {
		return &ErrorHandlerError{}, nil
	} else if typ == 665 || typ == 666 {
		var kind string
		if typ == 665 {
			kind = "out_of_gas"
		} else {
			kind = "hit_error_handler"
		}

		if tup.Len() != 4 {
			return nil, errors.New("expected 665 or 666 to be 4-tuple")
		}
		txID, _ := tup.GetByInt64(1)
		currentFrame, _ := tup.GetByInt64(2)
		parentFrame, _ := tup.GetByInt64(3)

		return generateLog(txID, currentFrame, parentFrame, kind, nil)
	} else if typ == 10000 {
		if tup.Len() != 6 {
			return nil, errors.New("expected type 10000 to be 6-tuple")
		}
		evmPC, _ := tup.GetByInt64(2)
		txID, _ := tup.GetByInt64(3)
		currentFrame, _ := tup.GetByInt64(4)
		parentFrame, _ := tup.GetByInt64(5)

		evmPCInt, ok := evmPC.(value.IntValue)
		if !ok {
			return nil, errors.New("expected pc to be in")
		}
		pc := evmPCInt.BigInt().Uint64()

		return generateLog(txID, currentFrame, parentFrame, "evm_revert", &pc)
	} else if typ == 20000 {
		vals, err := NewTraceFromDebugPrint(d)
		if err != nil {
			return nil, err
		}
		return &EVMTrace{Items: vals}, nil
	} else {
		return &RawDebugPrint{Val: d}, nil
	}
}
