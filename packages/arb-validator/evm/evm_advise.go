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

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Result interface {
	IsResult()

	GetEthMsg() EthBridgeMessage
}

type Return struct {
	Msg     EthBridgeMessage
	ArbCall ArbMessage

	ReturnVal []byte
	Logs      []Log
}

func (e Return) GetEthMsg() EthBridgeMessage {
	return e.Msg
}

func (e Return) IsResult() {}

func (e Return) String() string {
	var sb strings.Builder
	sb.WriteString("EVMReturn(func: ")
	sb.WriteString(e.ArbCall.GetFuncName())
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
	Msg       EthBridgeMessage
	ArbCall   ArbMessage
	ReturnVal []byte
}

func (e Revert) GetEthMsg() EthBridgeMessage {
	return e.Msg
}

func (e Revert) IsResult() {}

func (e Revert) String() string {
	var sb strings.Builder
	sb.WriteString("EVMRevert(func: ")
	sb.WriteString(e.ArbCall.GetFuncName())
	sb.WriteString(", returnVal: ")
	sb.WriteString(hexutil.Encode(e.ReturnVal))
	sb.WriteString(")")
	return sb.String()
}

type Stop struct {
	Msg     EthBridgeMessage
	ArbCall ArbMessage
	Logs    []Log
}

func (e Stop) GetEthMsg() EthBridgeMessage {
	return e.Msg
}

func (e Stop) IsResult() {}

func (e Stop) String() string {
	var sb strings.Builder
	sb.WriteString("EVMStop(func: ")
	sb.WriteString(e.ArbCall.GetFuncName())
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

type BadSequenceNum struct {
	Msg     EthBridgeMessage
	ArbCall ArbMessage
}

func (e BadSequenceNum) GetEthMsg() EthBridgeMessage {
	return e.Msg
}

func (e BadSequenceNum) IsResult() {}

func (e BadSequenceNum) String() string {
	var sb strings.Builder
	sb.WriteString("BadSequenceNum(func: ")
	sb.WriteString(e.ArbCall.GetFuncName())
	sb.WriteString("])")
	return sb.String()
}

type Invalid struct {
	Msg     EthBridgeMessage
	ArbCall ArbMessage
}

func (e Invalid) GetEthMsg() EthBridgeMessage {
	return e.Msg
}

func (e Invalid) IsResult() {}

func (e Invalid) String() string {
	var sb strings.Builder
	sb.WriteString("Invalid(func: ")
	sb.WriteString(e.ArbCall.GetFuncName())
	sb.WriteString("])")
	return sb.String()
}

type FuncCall struct {
	funcID [4]byte
	logs   value.Value
}

const (
	RevertCode      = 0
	InvalidCode     = 1
	ReturnCode      = 2
	StopCode        = 3
	BadSequenceCode = 4
)

type Log struct {
	ContractID value.IntValue
	Data       []byte
	Topics     []common.Hash
}

func (l Log) String() string {
	var sb strings.Builder
	sb.WriteString("Log(contract: ")
	sb.WriteString(l.ContractID.String())
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
		return Log{}, errors.New("log must be a tuple")
	}
	if tupVal.Len() < 3 {
		return Log{}, fmt.Errorf("log tuple must be at least size 3, but is %v", tupVal)
	}
	contractIDVal, _ := tupVal.GetByInt64(0)
	contractIDInt, ok := contractIDVal.(value.IntValue)
	if !ok {
		return Log{}, errors.New("log contract id must be an int")
	}
	logDataByteVal, _ := tupVal.GetByInt64(1)
	logData, err := ByteStackToHex(logDataByteVal)
	if err != nil {
		return Log{}, err
	}
	topics := make([]common.Hash, 0, tupVal.Len()-2)
	for _, topicVal := range tupVal.Contents()[2:] {
		topicValInt, ok := topicVal.(value.IntValue)
		if !ok {
			return Log{}, errors.New("log topic must be an int")
		}
		topics = append(topics, topicValInt.ToBytes())
	}

	return Log{contractIDInt, logData, topics}, nil
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

type ArbMessage interface {
	GetFuncName() string
}

type EthTransferMessage struct {
	Dest   common.Address
	Amount *big.Int
}

func (m EthTransferMessage) GetFuncName() string {
	return "EthTransfer"
}

func NewEthTransferMessageFromValue(val value.Value) (EthTransferMessage, error) {
	tup, ok := val.(value.TupleValue)
	if !ok {
		return EthTransferMessage{}, errors.New("msg must be tuple value")
	}
	if tup.Len() != 2 {
		return EthTransferMessage{}, fmt.Errorf("expected tuple of length 2, but recieved %v", tup)
	}
	destVal, _ := tup.GetByInt64(0)
	amountVal, _ := tup.GetByInt64(1)

	destInt, ok := destVal.(value.IntValue)
	if !ok {
		return EthTransferMessage{}, errors.New("dest must be an int")
	}

	destBytes := destInt.ToBytes()
	var destAddress common.Address
	copy(destAddress[:], destBytes[:20])

	amountInt, ok := amountVal.(value.IntValue)
	if !ok {
		return EthTransferMessage{}, errors.New("amount must be an int")
	}

	return EthTransferMessage{
		Dest:   destAddress,
		Amount: amountInt.BigInt(),
	}, nil
}

type TokenTransferMessage struct {
	TokenAddress common.Address
	Dest         common.Address
	Amount       *big.Int
}

func (m TokenTransferMessage) GetFuncName() string {
	return "TokenTransfer"
}

func NewTokenTransferMessageFromValue(val value.Value) (TokenTransferMessage, error) {
	tup, ok := val.(value.TupleValue)
	if !ok {
		return TokenTransferMessage{}, errors.New("msg must be tuple value")
	}
	if tup.Len() != 3 {
		return TokenTransferMessage{}, fmt.Errorf("expected tuple of length 3, but recieved %v", tup)
	}
	tokenAddressVal, _ := tup.GetByInt64(0)
	destVal, _ := tup.GetByInt64(1)
	amountVal, _ := tup.GetByInt64(2)

	tokenAddressInt, ok := tokenAddressVal.(value.IntValue)
	if !ok {
		return TokenTransferMessage{}, errors.New("token address must be an int")
	}

	tokenAddressBytes := tokenAddressInt.ToBytes()
	var tokenAddress common.Address
	copy(tokenAddress[:], tokenAddressBytes[:20])

	destInt, ok := destVal.(value.IntValue)
	if !ok {
		return TokenTransferMessage{}, errors.New("dest must be an int")
	}

	destBytes := destInt.ToBytes()
	var destAddress common.Address
	copy(destAddress[:], destBytes[:20])

	amountInt, ok := amountVal.(value.IntValue)
	if !ok {
		return TokenTransferMessage{}, errors.New("amount must be an int")
	}

	return TokenTransferMessage{
		TokenAddress: tokenAddress,
		Dest:         destAddress,
		Amount:       amountInt.BigInt(),
	}, nil
}

type TxMessage struct {
	To          common.Address
	SequenceNum *big.Int
	Amount      *big.Int
	Data        []byte
}

func (m TxMessage) GetFuncName() string {
	return hexutil.Encode(m.Data[:4])
}

func NewTxMessageFromValue(val value.Value) (TxMessage, error) {
	tup, ok := val.(value.TupleValue)
	if !ok {
		return TxMessage{}, errors.New("msg must be tuple value")
	}
	if tup.Len() != 4 {
		return TxMessage{}, fmt.Errorf("expected tuple of length 3, but recieved %v", tup)
	}
	toVal, _ := tup.GetByInt64(0)
	sequenceNumVal, _ := tup.GetByInt64(1)
	valueVal, _ := tup.GetByInt64(2)
	dataVal, _ := tup.GetByInt64(3)

	toValInt, ok := toVal.(value.IntValue)
	if !ok {
		return TxMessage{}, errors.New("to must be an int")
	}

	toBytes := toValInt.ToBytes()
	var toAddress common.Address
	copy(toAddress[:], toBytes[:20])

	sequenceNumInt, ok := sequenceNumVal.(value.IntValue)
	if !ok {
		return TxMessage{}, errors.New("sequenceNum must be an int")
	}

	valueInt, ok := valueVal.(value.IntValue)
	if !ok {
		return TxMessage{}, errors.New("value must be an int")
	}

	data, err := ByteStackToHex(dataVal)
	if err != nil {
		return TxMessage{}, nil
	}

	return TxMessage{
		To:          toAddress,
		SequenceNum: sequenceNumInt.BigInt(),
		Amount:      valueInt.BigInt(),
		Data:        data,
	}, nil
}

type EthBridgeMessage struct {
	Type        uint8
	BlockNumber *big.Int
	TxHash      common.Hash
	Sender      common.Address
}

func NewEthBridgeMessageFromValue(val value.Value) (EthBridgeMessage, value.Value, error) {
	tup, ok := val.(value.TupleValue)
	if !ok {
		return EthBridgeMessage{}, nil, errors.New("msg must be tuple value")
	}
	if tup.Len() != 3 {
		return EthBridgeMessage{}, nil, fmt.Errorf("expected tuple of length 3, but recieved %v", tup)
	}
	blockNumberVal, _ := tup.GetByInt64(0)
	txHashVal, _ := tup.GetByInt64(1)
	restVal, _ := tup.GetByInt64(2)

	blockNumberInt, ok := blockNumberVal.(value.IntValue)
	if !ok {
		return EthBridgeMessage{}, nil, errors.New("block number must be an int")
	}

	txHashInt, ok := txHashVal.(value.IntValue)
	if !ok {
		return EthBridgeMessage{}, nil, errors.New("tx hash must be an int")
	}

	txHashBytes := txHashInt.ToBytes()
	var txHash common.Hash
	copy(txHash[:], txHashBytes[:])

	restValTup, ok := restVal.(value.TupleValue)
	if !ok {
		return EthBridgeMessage{}, nil, errors.New("message must be a tup")
	}

	typeVal, _ := restValTup.GetByInt64(0)
	senderVal, _ := restValTup.GetByInt64(1)
	messageVal, _ := restValTup.GetByInt64(2)

	typeInt, ok := typeVal.(value.IntValue)
	if !ok {
		return EthBridgeMessage{}, nil, errors.New("type must be an int")
	}
	typecode := uint8(typeInt.BigInt().Uint64())

	senderInt, ok := senderVal.(value.IntValue)
	if !ok {
		return EthBridgeMessage{}, nil, errors.New("sender must be an int")
	}

	senderBytes := senderInt.ToBytes()
	var senderAddress common.Address
	copy(senderAddress[:], senderBytes[:20])

	return EthBridgeMessage{
		Type:        typecode,
		BlockNumber: blockNumberInt.BigInt(),
		TxHash:      txHash,
		Sender:      senderAddress,
	}, messageVal, nil
}

func ParseArbMessage(typecode uint8, messageVal value.Value) (ArbMessage, error) {
	switch typecode {
	case 0:
		return NewTxMessageFromValue(messageVal)
	case 1:
		return NewEthTransferMessageFromValue(messageVal)
	case 2:
		return NewTokenTransferMessageFromValue(messageVal)
	case 3:
		return NewTokenTransferMessageFromValue(messageVal)
	default:
		return nil, errors.New("Invalid message type")
	}
}

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

	ethMsg, messageVal, err := NewEthBridgeMessageFromValue(origMsgVal)
	if err != nil {
		return nil, err
	}
	arbMessage, err := ParseArbMessage(ethMsg.Type, messageVal)
	if err != nil {
		return nil, err
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
		returnBytes, err := ByteStackToHex(returnVal)
		if err != nil {
			return nil, err
		}
		return Return{ethMsg, arbMessage, returnBytes, logs}, nil
	case RevertCode:
		// EVM Revert
		returnVal, _ := tup.GetByInt64(2)
		returnBytes, err := ByteStackToHex(returnVal)
		if err != nil {
			return nil, err
		}
		return Revert{ethMsg, arbMessage, returnBytes}, nil
	case StopCode:
		// EVM Stop
		logVal, _ := tup.GetByInt64(1)
		logs, err := LogStackToLogs(logVal)
		if err != nil {
			return nil, err
		}
		return Stop{ethMsg, arbMessage, logs}, nil
	case BadSequenceCode:
		return BadSequenceNum{ethMsg, arbMessage}, nil
	case InvalidCode:
		return Invalid{ethMsg, arbMessage}, nil
	default:
		// Unknown type
		return nil, errors.New("unknown return code")
	}
}
