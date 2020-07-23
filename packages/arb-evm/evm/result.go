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
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/l2message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	errors2 "github.com/pkg/errors"
	"math/big"
)

type ResultType int

const (
	ReturnCode               ResultType = 0
	RevertCode                          = 1
	CongestionCode                      = 2
	InsufficientGasFundsCode            = 3
	InsufficientTxFundsCode             = 4
	BadSequenceCode                     = 5
	InvalidMessageFormatCode            = 6
	UnknownErrorCode                    = 255
)

type Result struct {
	L1Message     message.InboxMessage
	ResultCode    ResultType
	ReturnData    []byte
	EVMLogs       []Log
	GasUsed       *big.Int
	GasPrice      *big.Int
	CumulativeGas *big.Int
	TxIndex       *big.Int
	StartLogIndex *big.Int
}

func (r *Result) String() string {
	return fmt.Sprintf(
		"Result(%v, %v, %v, %v, %v, %v)",
		r.L1Message,
		r.ResultCode,
		hexutil.Encode(r.ReturnData),
		r.EVMLogs,
		r.GasUsed,
		r.GasPrice,
	)
}

func (r *Result) AsValue() value.Value {
	tup, _ := value.NewTupleFromSlice([]value.Value{
		r.L1Message.AsValue(),
		value.NewInt64Value(int64(r.ResultCode)),
		message.BytesToByteStack(r.ReturnData),
		LogsToLogStack(r.EVMLogs),
		value.NewIntValue(r.GasUsed),
		value.NewIntValue(r.GasPrice),
	})
	return tup
}

func (r *Result) ToEthReceipt(blockHash common.Hash) (*types.Receipt, error) {
	contractAddress := ethcommon.Address{}
	if r.L1Message.Kind == message.L2Type {
		msg, err := l2message.NewL2MessageFromData(r.L1Message.Data)
		if err == nil {
			if msg, ok := msg.(l2message.Transaction); ok {
				emptyAddress := common.Address{}
				if msg.DestAddress == emptyAddress {
					copy(contractAddress[:], r.ReturnData[12:])
				}
			}
		}
	}

	status := uint64(0)
	if r.ResultCode == ReturnCode {
		status = 1
	}

	var evmLogs []*types.Log
	logIndex := r.StartLogIndex.Uint64()
	for _, l := range r.EVMLogs {
		ethLog := &types.Log{
			Address:     l.Address.ToEthAddress(),
			Topics:      common.NewEthHashesFromHashes(l.Topics),
			Data:        l.Data,
			BlockNumber: r.L1Message.ChainTime.BlockNum.AsInt().Uint64(),
			TxHash:      r.L1Message.MessageID().ToEthHash(),
			TxIndex:     uint(r.TxIndex.Uint64()),
			BlockHash:   blockHash.ToEthHash(),
			Index:       uint(logIndex),
		}
		logIndex++
		evmLogs = append(evmLogs, ethLog)
	}

	return &types.Receipt{
		PostState:         []byte{0},
		Status:            status,
		CumulativeGasUsed: r.CumulativeGas.Uint64(),
		Bloom:             types.BytesToBloom(types.LogsBloom(evmLogs).Bytes()),
		Logs:              evmLogs,
		TxHash:            r.L1Message.MessageID().ToEthHash(),
		ContractAddress:   contractAddress,
		GasUsed:           r.GasUsed.Uint64(),
		BlockHash:         blockHash.ToEthHash(),
		BlockNumber:       r.L1Message.ChainTime.BlockNum.AsInt(),
		TransactionIndex:  uint(r.TxIndex.Uint64()),
	}, nil
}

func NewResultFromValue(val value.Value) (*Result, error) {
	tup, ok := val.(value.TupleValue)
	if !ok || tup.Len() != 4 {
		return nil, fmt.Errorf("advise expected tuple of length 5, but recieved %v", tup)
	}
	l1MsgVal, _ := tup.GetByInt64(0)
	resultInfo, _ := tup.GetByInt64(1)
	gasInfo, _ := tup.GetByInt64(2)
	chainInfo, _ := tup.GetByInt64(3)

	resultTup, ok := resultInfo.(value.TupleValue)
	if !ok || resultTup.Len() != 3 {
		return nil, fmt.Errorf("advise expected result info tuple of length 3, but recieved %v", resultTup)
	}
	resultCode, _ := resultTup.GetByInt64(0)
	returnData, _ := resultTup.GetByInt64(1)
	evmLogs, _ := resultTup.GetByInt64(2)

	gasInfoTup, ok := gasInfo.(value.TupleValue)
	if !ok || gasInfoTup.Len() != 2 {
		return nil, fmt.Errorf("advise expected gas info tuple of length 2, but recieved %v", gasInfoTup)
	}
	gasUsed, _ := gasInfoTup.GetByInt64(0)
	gasPrice, _ := gasInfoTup.GetByInt64(1)

	chainInfoTup, ok := chainInfo.(value.TupleValue)
	if !ok || chainInfoTup.Len() != 3 {
		return nil, fmt.Errorf("advise expected tx block data tuple of length 3, but recieved %v", resultTup)
	}
	cumulativeGas, _ := chainInfoTup.GetByInt64(0)
	txIndex, _ := chainInfoTup.GetByInt64(1)
	startLogIndex, _ := chainInfoTup.GetByInt64(2)

	l1Msg, err := message.NewInboxMessageFromValue(l1MsgVal)
	if err != nil {
		return nil, err
	}
	returnBytes, err := message.ByteStackToHex(returnData)
	if err != nil {
		return nil, errors2.Wrap(err, "umarshalling return data")
	}
	logs, err := LogStackToLogs(evmLogs)
	if err != nil {
		return nil, errors2.Wrap(err, "unmarshaling logs")
	}
	resultCodeInt, ok := resultCode.(value.IntValue)
	if !ok {
		return nil, errors.New("resultCode must be an int")
	}
	gasUsedInt, ok := gasUsed.(value.IntValue)
	if !ok {
		return nil, errors.New("gasUsed must be an int")
	}
	gasPriceInt, ok := gasPrice.(value.IntValue)
	if !ok {
		return nil, errors.New("gasPrice must be an int")
	}
	cumulativeGasInt, ok := cumulativeGas.(value.IntValue)
	if !ok {
		return nil, errors.New("cumulativeGas must be an int")
	}
	txIndexInt, ok := txIndex.(value.IntValue)
	if !ok {
		return nil, errors.New("txIndex must be an int")
	}
	startLogIndexInt, ok := startLogIndex.(value.IntValue)
	if !ok {
		return nil, errors.New("startLogIndex must be an int")
	}

	return &Result{
		L1Message:     l1Msg,
		ResultCode:    ResultType(resultCodeInt.BigInt().Uint64()),
		ReturnData:    returnBytes,
		EVMLogs:       logs,
		GasUsed:       gasUsedInt.BigInt(),
		GasPrice:      gasPriceInt.BigInt(),
		CumulativeGas: cumulativeGasInt.BigInt(),
		TxIndex:       txIndexInt.BigInt(),
		StartLogIndex: startLogIndexInt.BigInt(),
	}, nil
}

func NewRandomResult(msg message.Message, logCount int32) *Result {
	logs := make([]Log, 0, logCount)
	for i := int32(0); i < logCount; i++ {
		logs = append(logs, NewRandomLog(3))
	}
	return &Result{
		L1Message:     message.NewRandomInboxMessage(msg),
		ResultCode:    ReturnCode,
		ReturnData:    common.RandBytes(200),
		EVMLogs:       logs,
		GasUsed:       common.RandBigInt(),
		GasPrice:      common.RandBigInt(),
		CumulativeGas: common.RandBigInt(),
		TxIndex:       common.RandBigInt(),
		StartLogIndex: common.RandBigInt(),
	}
}
