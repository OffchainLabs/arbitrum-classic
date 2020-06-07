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

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
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

	GetEthMsg() message.SingleDelivered
	Type() ResultType
}

func ResultAsValue(result Result) value.TupleValue {
	tup, _ := value.NewTupleFromSlice([]value.Value{
		result.GetEthMsg().AsInboxValue(),
		LogsToLogStack(result.GetLogs()),
		message.BytesToByteStack(result.GetReturnData()),
		value.NewInt64Value(int64(result.Type())),
	})
	return tup
}

type Return struct {
	Msg       message.SingleDelivered
	ReturnVal []byte
	Logs      []Log
}

func (e Return) Type() ResultType {
	return ReturnCode
}

func (e Return) GetEthMsg() message.SingleDelivered {
	return e.Msg
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
	sb.WriteString("EVMReturn(func: ")
	sb.WriteString(e.Msg.ExectutedMessage().GetFuncName())
	sb.WriteString(", returnVal: ")
	sb.WriteString(hexutil.Encode(e.ReturnVal))
	sb.WriteString(", logs: [")
	for i, log := range e.Logs {
		sb.WriteString(log.String())
		if i != len(e.Logs)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]) from transaction ")
	sb.WriteString(e.Msg.Message.String())
	return sb.String()
}

type Revert struct {
	Msg       message.SingleDelivered
	ReturnVal []byte
}

func (e Revert) Type() ResultType {
	return RevertCode
}

func (e Revert) GetEthMsg() message.SingleDelivered {
	return e.Msg
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
	sb.WriteString("EVMRevert(func: ")
	sb.WriteString(e.Msg.ExectutedMessage().GetFuncName())
	sb.WriteString(", returnVal: ")
	sb.WriteString(hexutil.Encode(e.ReturnVal))
	sb.WriteString(") from transaction ")
	sb.WriteString(e.Msg.Message.String())
	return sb.String()
}

type Stop struct {
	Msg  message.SingleDelivered
	Logs []Log
}

func NewRandomStop(msg message.ExecutionMessage, logCount int32) Stop {
	logs := make([]Log, 0, logCount)
	for i := int32(0); i < logCount; i++ {
		logs = append(logs, NewRandomLog(3))
	}

	return Stop{
		Msg:  message.NewRandomSingleDelivered(msg),
		Logs: logs,
	}
}

func (e Stop) Type() ResultType {
	return StopCode
}

func (e Stop) GetEthMsg() message.SingleDelivered {
	return e.Msg
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
	sb.WriteString("EVMStop(func: ")
	sb.WriteString(e.Msg.ExectutedMessage().GetFuncName())
	sb.WriteString(", logs: [")
	for i, log := range e.Logs {
		sb.WriteString(log.String())
		if i != len(e.Logs)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]) from transaction ")
	sb.WriteString(e.Msg.Message.String())
	return sb.String()
}

type BadSequenceNum struct {
	Msg message.SingleDelivered
}

func (e BadSequenceNum) Type() ResultType {
	return BadSequenceCode
}

func (e BadSequenceNum) GetEthMsg() message.SingleDelivered {
	return e.Msg
}

func (e BadSequenceNum) IsResult() {}

func (e BadSequenceNum) GetReturnData() []byte {
	return nil
}

func (e BadSequenceNum) GetLogs() []Log {
	return nil
}

func (e BadSequenceNum) String() string {
	var sb strings.Builder
	sb.WriteString("BadSequenceNum(func: ")
	sb.WriteString(e.Msg.ExectutedMessage().GetFuncName())
	sb.WriteString("]) from transaction ")
	sb.WriteString(e.Msg.Message.String())
	return sb.String()
}

type Invalid struct {
	Msg message.SingleDelivered
}

func (e Invalid) Type() ResultType {
	return InvalidCode
}

func (e Invalid) GetEthMsg() message.SingleDelivered {
	return e.Msg
}

func (e Invalid) IsResult() {}

func (e Invalid) GetReturnData() []byte {
	return nil
}

func (e Invalid) GetLogs() []Log {
	return nil
}

func (e Invalid) String() string {
	var sb strings.Builder
	sb.WriteString("Invalid(func: ")
	sb.WriteString(e.Msg.ExectutedMessage().GetFuncName())
	sb.WriteString("]) from transaction ")
	sb.WriteString(e.Msg.Message.String())
	return sb.String()
}

type FuncCall struct {
	funcID [4]byte
	logs   value.Value
}

func ProcessLog(val value.Value, chain common.Address) (Result, error) {
	tup, ok := val.(value.TupleValue)
	if !ok {
		return nil, errors.New("advise expected tuple value")
	}
	if tup.Len() != 4 {
		return nil, fmt.Errorf("advise expected tuple of length 4, but recieved %v", tup)
	}
	origMsgVal, _ := tup.GetByInt64(0)
	ethMsg, err := message.UnmarshalDelivered(origMsgVal, chain)
	if err != nil {
		return nil, err
	}

	singleDelivered, err := message.NewSingleDelivered(ethMsg)

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
		return Return{singleDelivered, returnBytes, logs}, nil
	case RevertCode:
		// EVM Revert
		returnVal, _ := tup.GetByInt64(2)
		returnBytes, err := message.ByteStackToHex(returnVal)
		if err != nil {
			return nil, err
		}
		return Revert{singleDelivered, returnBytes}, nil
	case StopCode:
		// EVM Stop
		logVal, _ := tup.GetByInt64(1)
		logs, err := LogStackToLogs(logVal)
		if err != nil {
			return nil, err
		}
		return Stop{singleDelivered, logs}, nil
	case BadSequenceCode:
		return BadSequenceNum{singleDelivered}, nil
	case InvalidCode:
		return Invalid{singleDelivered}, nil
	default:
		// Unknown type
		return nil, fmt.Errorf("unknown return code %v for message %v", returnCode.BigInt(), val)
	}
}
