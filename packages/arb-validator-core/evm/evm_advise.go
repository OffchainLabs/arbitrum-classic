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
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

type Result interface {
	IsResult()
	GetLogs() []Log

	GetEthMsg() EthBridgeMessage
}

const (
	RevertCode      = 0
	InvalidCode     = 1
	ReturnCode      = 2
	StopCode        = 3
	BadSequenceCode = 4
)

type Return struct {
	Msg       EthBridgeMessage
	ReturnVal []byte
	Logs      []Log
}

func (e Return) GetEthMsg() EthBridgeMessage {
	return e.Msg
}

func (e Return) IsResult() {}

func (e Return) GetLogs() []Log {
	return e.Logs
}

func (e Return) String() string {
	var sb strings.Builder
	sb.WriteString("EVMReturn(func: ")
	sb.WriteString(e.Msg.ArbCall.GetFuncName())
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
	sb.WriteString(e.Msg.ArbCall.String())
	return sb.String()
}

type Revert struct {
	Msg       EthBridgeMessage
	ReturnVal []byte
}

func (e Revert) GetEthMsg() EthBridgeMessage {
	return e.Msg
}

func (e Revert) IsResult() {}

func (e Revert) GetLogs() []Log {
	return nil
}

func (e Revert) String() string {
	var sb strings.Builder
	sb.WriteString("EVMRevert(func: ")
	sb.WriteString(e.Msg.ArbCall.GetFuncName())
	sb.WriteString(", returnVal: ")
	sb.WriteString(hexutil.Encode(e.ReturnVal))
	sb.WriteString(") from transaction ")
	sb.WriteString(e.Msg.ArbCall.String())
	return sb.String()
}

type Stop struct {
	Msg  EthBridgeMessage
	Logs []Log
}

func (e Stop) GetEthMsg() EthBridgeMessage {
	return e.Msg
}

func (e Stop) IsResult() {}

func (e Stop) GetLogs() []Log {
	return e.Logs
}

func (e Stop) String() string {
	var sb strings.Builder
	sb.WriteString("EVMStop(func: ")
	sb.WriteString(e.Msg.ArbCall.GetFuncName())
	sb.WriteString(", logs: [")
	for i, log := range e.Logs {
		sb.WriteString(log.String())
		if i != len(e.Logs)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]) from transaction ")
	sb.WriteString(e.Msg.ArbCall.String())
	return sb.String()
}

type BadSequenceNum struct {
	Msg EthBridgeMessage
}

func (e BadSequenceNum) GetEthMsg() EthBridgeMessage {
	return e.Msg
}

func (e BadSequenceNum) IsResult() {}

func (e BadSequenceNum) GetLogs() []Log {
	return nil
}

func (e BadSequenceNum) String() string {
	var sb strings.Builder
	sb.WriteString("BadSequenceNum(func: ")
	sb.WriteString(e.Msg.ArbCall.GetFuncName())
	sb.WriteString("]) from transaction ")
	sb.WriteString(e.Msg.ArbCall.String())
	return sb.String()
}

type Invalid struct {
	Msg EthBridgeMessage
}

func (e Invalid) GetEthMsg() EthBridgeMessage {
	return e.Msg
}

func (e Invalid) IsResult() {}

func (e Invalid) GetLogs() []Log {
	return nil
}

func (e Invalid) String() string {
	var sb strings.Builder
	sb.WriteString("Invalid(func: ")
	sb.WriteString(e.Msg.ArbCall.GetFuncName())
	sb.WriteString("]) from transaction ")
	sb.WriteString(e.Msg.ArbCall.String())
	return sb.String()
}

type FuncCall struct {
	funcID [4]byte
	logs   value.Value
}

type EthBridgeMessage struct {
	BlockNumber *big.Int
	Timestamp   *big.Int
	TxHash      common.Hash
	ArbCall     message.ExecutionMessage
}

func NewEthBridgeMessageFromValue(val value.Value, chain common.Address) (EthBridgeMessage, error) {
	tup, ok := val.(value.TupleValue)
	invalid := EthBridgeMessage{}
	if !ok {
		return invalid, errors.New("msg must be tuple value")
	}
	if tup.Len() != 4 {
		return invalid, fmt.Errorf("expected tuple of length 4, but recieved %v", tup)
	}
	blockNumberVal, _ := tup.GetByInt64(0)
	timestampVal, _ := tup.GetByInt64(1)
	txHashVal, _ := tup.GetByInt64(2)
	restVal, _ := tup.GetByInt64(3)

	blockNumberInt, ok := blockNumberVal.(value.IntValue)
	if !ok {
		return invalid, errors.New("block number must be an int")
	}

	timestampInt, ok := timestampVal.(value.IntValue)
	if !ok {
		return invalid, errors.New("timestamp must be an int")
	}

	txHashInt, ok := txHashVal.(value.IntValue)
	if !ok {
		return invalid, errors.New("tx hash must be an int")
	}

	txHashBytes := txHashInt.ToBytes()
	var txHash common.Hash
	copy(txHash[:], txHashBytes[:])

	restValTup, ok := restVal.(value.TupleValue)
	if !ok {
		return invalid, errors.New("message must be a tup")
	}

	typeVal, _ := restValTup.GetByInt64(0)
	typeInt, ok := typeVal.(value.IntValue)
	if !ok {
		return invalid, errors.New("type must be an int")
	}
	typecode := uint8(typeInt.BigInt().Uint64())

	arbMessage, err := message.UnmarshalExecuted(message.Type(typecode), restValTup, chain)
	if err != nil {
		return invalid, err
	}

	return EthBridgeMessage{
		BlockNumber: blockNumberInt.BigInt(),
		Timestamp:   timestampInt.BigInt(),
		TxHash:      txHash,
		ArbCall:     arbMessage,
	}, nil
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
	ethMsg, err := NewEthBridgeMessageFromValue(origMsgVal, chain)
	if err != nil {
		return nil, err
	}

	returnCodeVal, _ := tup.GetByInt64(3)
	returnCode, ok := returnCodeVal.(value.IntValue)
	if !ok {
		return nil, errors.New("return code must be an int")
	}

	switch returnCode.BigInt().Uint64() {
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
		return Return{ethMsg, returnBytes, logs}, nil
	case RevertCode:
		// EVM Revert
		returnVal, _ := tup.GetByInt64(2)
		returnBytes, err := message.ByteStackToHex(returnVal)
		if err != nil {
			return nil, err
		}
		return Revert{ethMsg, returnBytes}, nil
	case StopCode:
		// EVM Stop
		logVal, _ := tup.GetByInt64(1)
		logs, err := LogStackToLogs(logVal)
		if err != nil {
			return nil, err
		}
		return Stop{ethMsg, logs}, nil
	case BadSequenceCode:
		return BadSequenceNum{ethMsg}, nil
	case InvalidCode:
		return Invalid{ethMsg}, nil
	default:
		// Unknown type
		return nil, fmt.Errorf("unknown return code %v for message %v", returnCode.BigInt(), val)
	}
}

func NewVMResultValue(delivered message.Delivered, returnCode int, data []byte, logs []Log) (value.TupleValue, error) {
	msg, ok := delivered.Message.(message.SingleMessage)
	if !ok {
		return value.TupleValue{}, errors.New("can only make stop from single message")
	}
	val := message.DeliveredValue(delivered, msg)

	tup, _ := value.NewTupleFromSlice([]value.Value{
		val,
		LogsToLogStack(logs),
		message.BytesToByteStack(data),
		value.NewInt64Value(int64(returnCode)),
	})
	return tup, nil
}
