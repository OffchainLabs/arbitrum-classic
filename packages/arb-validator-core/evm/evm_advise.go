/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

type ResultType int

const (
	RevertCode      ResultType = 0
	InvalidCode                = 1
	ReturnCode                 = 2
	StopCode                   = 3
	BadSequenceCode            = 4
)

type Result interface {
	IsResult()
	GetReturnData() []byte
	GetLogs() []Log

	GetDeliveredMessage() value.Value
	Type() ResultType
}

func ResultAsValue(result Result) value.TupleValue {
	tup, _ := value.NewTupleFromSlice([]value.Value{
		result.GetDeliveredMessage(),
		LogsToLogStack(result.GetLogs()),
		message.BytesToByteStack(result.GetReturnData()),
		value.NewInt64Value(int64(result.Type())),
	})
	return tup
}

type Return struct {
	Delivered value.Value
	ReturnVal []byte
	Logs      []Log
}

func (e Return) Type() ResultType {
	return ReturnCode
}

func (e Return) GetDeliveredMessage() value.Value {
	return e.Delivered
}

func (e Return) IsResult() {}

func (e Return) GetReturnData() []byte {
	return e.ReturnVal
}

func (e Return) GetLogs() []Log {
	return e.Logs
}

func (e Return) String() string {
	var sb strings.Builder
	sb.WriteString("EVMReturn(returnVal: ")
	sb.WriteString(hexutil.Encode(e.ReturnVal))
	sb.WriteString(", logs: [")
	for i, log := range e.Logs {
		sb.WriteString(log.String())
		if i != len(e.Logs)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("])")
	return sb.String()
}

type Revert struct {
	Delivered value.Value
	ReturnVal []byte
}

func (e Revert) Type() ResultType {
	return RevertCode
}

func (e Revert) GetDeliveredMessage() value.Value {
	return e.Delivered
}

func (e Revert) IsResult() {}

func (e Revert) GetReturnData() []byte {
	return e.ReturnVal
}

func (e Revert) GetLogs() []Log {
	return nil
}

func (e Revert) String() string {
	var sb strings.Builder
	sb.WriteString("EVMRevert(returnVal: ")
	sb.WriteString(hexutil.Encode(e.ReturnVal))
	sb.WriteString(")")
	return sb.String()
}

type Stop struct {
	Delivered value.Value
	Logs      []Log
}

func NewRandomStop(msg message.ExecutionMessage, logCount int32) Stop {
	logs := make([]Log, 0, logCount)
	for i := int32(0); i < logCount; i++ {
		logs = append(logs, NewRandomLog(3))
	}
	single := message.NewRandomSingleDelivered(msg)
	return Stop{
		Delivered: single.AsInboxValue(),
		Logs:      logs,
	}
}

func (e Stop) Type() ResultType {
	return StopCode
}

func (e Stop) GetDeliveredMessage() value.Value {
	return e.Delivered
}

func (e Stop) IsResult() {}

func (e Stop) GetReturnData() []byte {
	return nil
}

func (e Stop) GetLogs() []Log {
	return e.Logs
}

func (e Stop) String() string {
	var sb strings.Builder
	sb.WriteString("EVMStop(logs: [")
	for i, log := range e.Logs {
		sb.WriteString(log.String())
		if i != len(e.Logs)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("])")
	return sb.String()
}

type BadSequenceNum struct {
	Delivered value.Value
}

func (e BadSequenceNum) Type() ResultType {
	return BadSequenceCode
}

func (e BadSequenceNum) GetDeliveredMessage() value.Value {
	return e.Delivered
}

func (e BadSequenceNum) IsResult() {}

func (e BadSequenceNum) GetReturnData() []byte {
	return nil
}

func (e BadSequenceNum) GetLogs() []Log {
	return nil
}

func (e BadSequenceNum) String() string {
	return "BadSequenceNum()"
}

type Invalid struct {
	Delivered value.Value
}

func (e Invalid) Type() ResultType {
	return InvalidCode
}

func (e Invalid) GetDeliveredMessage() value.Value {
	return e.Delivered
}

func (e Invalid) IsResult() {}

func (e Invalid) GetReturnData() []byte {
	return nil
}

func (e Invalid) GetLogs() []Log {
	return nil
}

func (e Invalid) String() string {
	return "Invalid()"
}

type FuncCall struct {
	funcID [4]byte
	logs   value.Value
}

func ProcessLog(val value.Value) (Result, error) {
	tup, ok := val.(value.TupleValue)
	if !ok {
		return nil, errors.New("advise expected tuple value")
	}
	if tup.Len() != 4 {
		return nil, fmt.Errorf("advise expected tuple of length 4, but recieved %v", tup)
	}
	origMsgVal, _ := tup.GetByInt64(0)

	returnCodeVal, _ := tup.GetByInt64(3)
	returnCode, ok := returnCodeVal.(value.IntValue)
	if !ok {
		return nil, errors.New("return code must be an int")
	}

	switch ResultType(returnCode.BigInt().Uint64()) {
	case ReturnCode:
		// EVM Return
		logVal, _ := tup.GetByInt64(1)
		logs, err := LogStackToLogs(logVal)
		if err != nil {
			return nil, err
		}
		returnVal, err := tup.GetByInt64(2)
		returnBytes, err := message.ByteStackToHex(returnVal)
		if err != nil {
			return nil, err
		}
		return Return{origMsgVal, returnBytes, logs}, nil
	case RevertCode:
		// EVM Revert
		returnVal, _ := tup.GetByInt64(2)
		returnBytes, err := message.ByteStackToHex(returnVal)
		if err != nil {
			return nil, err
		}
		return Revert{origMsgVal, returnBytes}, nil
	case StopCode:
		// EVM Stop
		logVal, _ := tup.GetByInt64(1)
		logs, err := LogStackToLogs(logVal)
		if err != nil {
			return nil, err
		}
		return Stop{origMsgVal, logs}, nil
	case BadSequenceCode:
		return BadSequenceNum{origMsgVal}, nil
	case InvalidCode:
		return Invalid{origMsgVal}, nil
	default:
		// Unknown type
		return nil, fmt.Errorf("unknown return code %v for message %v", returnCode.BigInt(), val)
	}
}
