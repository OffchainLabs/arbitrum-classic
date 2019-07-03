/*
 * Copyright 2019, Offchain Labs, Inc.
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

	"github.com/miguelmota/go-solidity-sha3"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arb-util/protocol"
	"github.com/offchainlabs/arb-util/value"
)

type Result interface {
	IsResult()

	GetEthMsg() EthMsg
}

type Return struct {
	Msg       EthMsg
	ReturnVal []byte
	Logs      []Log
}

func (e Return) GetEthMsg() EthMsg {
	return e.Msg
}

func (e Return) IsResult() {}

func (e Return) String() string {
	var sb strings.Builder
	sb.WriteString("EVMReturn(func: ")
	sb.WriteString(hexutil.Encode(e.Msg.Data.CallData.Data[:4]))
	sb.WriteString(", returnVal: ")
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
	Msg       EthMsg
	ReturnVal []byte
}

func (e Revert) GetEthMsg() EthMsg {
	return e.Msg
}

func (e Revert) IsResult() {}

func (e Revert) String() string {
	var sb strings.Builder
	sb.WriteString("EVMRevert(func: ")
	sb.WriteString(hexutil.Encode(e.Msg.Data.CallData.Data[:4]))
	sb.WriteString(", returnVal: ")
	sb.WriteString(hexutil.Encode(e.ReturnVal))
	sb.WriteString(")")
	return sb.String()
}

type Stop struct {
	Msg  EthMsg
	Logs []Log
}

func (e Stop) GetEthMsg() EthMsg {
	return e.Msg
}

func (e Stop) IsResult() {}

func (e Stop) String() string {
	var sb strings.Builder
	sb.WriteString("EVMStop(func: ")
	sb.WriteString(hexutil.Encode(e.Msg.Data.CallData.Data[:4]))
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

type FuncCall struct {
	funcID [4]byte
	logs   value.Value
}

const (
	EVM_REVERT_CODE       = 0
	EVM_INVALID_CODE      = 1
	EVM_RETURN_CODE       = 2
	EVM_STOP_CODE         = 3
	EVM_BAD_SEQUENCE_CODE = 4
)

type Log struct {
	ContractId value.IntValue
	Data       []byte
	Topics     [][32]byte
}

func (l Log) String() string {
	var sb strings.Builder
	sb.WriteString("Log(contract: ")
	sb.WriteString(l.ContractId.String())
	sb.WriteString(", topics: [")
	for i, topic := range l.Topics {
		sb.WriteString(hexutil.Encode(topic[:]))
		if i != len(l.Topics)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("], data:")
	sb.WriteString(hexutil.Encode(l.Data))
	sb.WriteString(")")
	return ""
}

func LogValToLog(val value.Value) (Log, error) {
	tupVal, ok := val.(value.TupleValue)
	if !ok {
		return Log{}, errors.New("Log must be a tuple")
	}
	if tupVal.Len() < 3 {
		return Log{}, fmt.Errorf("Log tuple must be at least size 3, but is %v", tupVal)
	}
	contractIdVal, _ := tupVal.GetByInt64(0)
	contractIdInt, ok := contractIdVal.(value.IntValue)
	if !ok {
		return Log{}, errors.New("Log contract id must be an int")
	}
	logDataByteVal, _ := tupVal.GetByInt64(1)
	logData, err := SizedByteArrayToHex(logDataByteVal)
	if err != nil {
		return Log{}, err
	}
	topics := make([][32]byte, 0, tupVal.Len()-2)
	for _, topicVal := range tupVal.Contents()[2:] {
		topicValInt, ok := topicVal.(value.IntValue)
		if !ok {
			return Log{}, errors.New("Log topic must be an int")
		}
		topics = append(topics, topicValInt.ToBytes())
	}

	return Log{contractIdInt, logData, topics}, nil
}

func LogStackToLogs(val value.Value) ([]Log, error) {
	logValues, err := StackValueToList(val)
	if err != nil {
		return nil, err
	}
	logs := make([]Log, 0, len(logValues))
	for _, logVal := range logValues {
		log, err := LogValToLog(logVal)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}

type EthCallData struct {
	Data        []byte
	ContractId  *big.Int
	SequenceNum *big.Int
}

func NewEthCallDataFromValue(val value.Value) (EthCallData, error) {
	tup, ok := val.(value.TupleValue)
	if !ok {
		return EthCallData{}, errors.New("msg must be tuple value")
	}
	if tup.Len() != 3 {
		return EthCallData{}, fmt.Errorf("expected tuple of length 3, but recieved %v", tup)
	}
	dataVal, _ := tup.GetByInt64(0)
	contractIDVal, _ := tup.GetByInt64(1)
	sequenceNumVal, _ := tup.GetByInt64(2)

	contractIDInt, ok := contractIDVal.(value.IntValue)
	if !ok {
		return EthCallData{}, errors.New("contractID must be an int")
	}

	sequenceNumInt, ok := sequenceNumVal.(value.IntValue)
	if !ok {
		return EthCallData{}, errors.New("sequenceNum must be an int")
	}

	data, err := SizedByteArrayToHex(dataVal)
	if err != nil {
		return EthCallData{}, nil
	}

	return EthCallData{
		data,
		contractIDInt.BigInt(),
		sequenceNumInt.BigInt(),
	}, nil
}

type EthMsgData struct {
	CallData  EthCallData
	Timestamp *big.Int
	Number    *big.Int
	TxHash    [32]byte

	dataHash [32]byte
}

func NewEthMsgDataFromValue(val value.Value) (EthMsgData, error) {
	tup, ok := val.(value.TupleValue)
	if !ok {
		return EthMsgData{}, errors.New("msg must be tuple value")
	}
	if tup.Len() != 4 {
		return EthMsgData{}, fmt.Errorf("expected tuple of length 4, but recieved %v", tup)
	}
	dataVal, _ := tup.GetByInt64(0)
	timestampVal, _ := tup.GetByInt64(1)
	numberVal, _ := tup.GetByInt64(2)
	txHashVal, _ := tup.GetByInt64(3)

	callData, err := NewEthCallDataFromValue(dataVal)
	if err != nil {
		return EthMsgData{}, err
	}

	timestampInt, ok := timestampVal.(value.IntValue)
	if !ok {
		return EthMsgData{}, errors.New("timestamp must be an int")
	}

	numberInt, ok := numberVal.(value.IntValue)
	if !ok {
		return EthMsgData{}, errors.New("number must be an int")
	}

	txHashInt, ok := txHashVal.(value.IntValue)
	if !ok {
		return EthMsgData{}, errors.New("txhash must be an int")
	}

	return EthMsgData{
		callData,
		timestampInt.BigInt(),
		numberInt.BigInt(),
		txHashInt.ToBytes(),
		dataVal.Hash(),
	}, nil
}

type EthMsg struct {
	Data      EthMsgData
	TokenType [21]byte
	Currency  *big.Int
	Caller    [32]byte
}

func (msg EthMsg) MsgHash(vmId [32]byte) [32]byte {
	ret := [32]byte{}
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(vmId),
		solsha3.Bytes32(msg.Data.dataHash),
		solsha3.Uint256(msg.Currency),
		msg.TokenType[:],
	))
	return ret
}

func NewEthMsgFromValue(val value.Value) (EthMsg, error) {
	msg, err := protocol.NewMessageFromValue(val)
	if err != nil {
		return EthMsg{}, err
	}
	ethMsgData, err := NewEthMsgDataFromValue(msg.Data)
	if err != nil {
		return EthMsg{}, err
	}
	return EthMsg{
		ethMsgData,
		msg.TokenType,
		msg.Currency,
		msg.Destination,
	}, nil
}

// [logs, contract_num, func_code, return_val, return_code]
func ProcessLog(val value.Value) (Result, error) {
	tup, ok := val.(value.TupleValue)
	if !ok {
		return nil, errors.New("advise expected tuple value")
	}
	if tup.Len() != 4 {
		return nil, fmt.Errorf("advise expected tuple of length 4, but recieved %v", tup)
	}
	returnCodeVal, _ := tup.GetByInt64(3)
	returnCode, ok := returnCodeVal.(value.IntValue)
	if !ok {
		return nil, errors.New("return code must be an int")
	}

	origMsgVal, _ := tup.GetByInt64(0)

	ethMsg, err := NewEthMsgFromValue(origMsgVal)
	if err != nil {
		return nil, err
	}

	switch returnCode.BigInt().Uint64() {
	case EVM_RETURN_CODE:
		// EVM Return
		logVal, _ := tup.GetByInt64(1)
		logs, err := LogStackToLogs(logVal)
		if err != nil {
			return nil, err
		}
		returnVal, err := tup.GetByInt64(2)
		returnBytes, err := SizedByteArrayToHex(returnVal)
		if err != nil {
			return nil, err
		}
		return Return{ethMsg, returnBytes, logs}, nil
	case EVM_REVERT_CODE:
		// EVM Revert
		returnVal, _ := tup.GetByInt64(2)
		returnBytes, err := SizedByteArrayToHex(returnVal)
		if err != nil {
			return nil, err
		}
		return Revert{ethMsg, returnBytes}, nil
	case EVM_STOP_CODE:
		// EVM Stop
		logVal, _ := tup.GetByInt64(1)
		logs, err := LogStackToLogs(logVal)
		if err != nil {
			return nil, err
		}
		return Stop{ethMsg, logs}, nil
	case EVM_BAD_SEQUENCE_CODE:
		return nil, errors.New("bad sequence number")
	case EVM_INVALID_CODE:
		return nil, errors.New("invalid tx")
	default:
		// Unknown type
		return nil, errors.New("unknown return code")
	}
}
