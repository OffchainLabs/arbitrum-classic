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
	"bytes"
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/pkg/errors"
	"math/big"
)

type ResultType int

const (
	ReturnCode               ResultType = 0
	RevertCode               ResultType = 1
	CongestionCode           ResultType = 2
	InsufficientGasFundsCode ResultType = 3
	InsufficientTxFundsCode  ResultType = 4
	BadSequenceCode          ResultType = 5
	InvalidMessageFormatCode ResultType = 6
	ContractAlreadyExists    ResultType = 7
	UnknownErrorCode         ResultType = 255
)

type Result interface {
}

type TxResult struct {
	IncomingRequest IncomingRequest
	ResultCode      ResultType
	ReturnData      []byte
	EVMLogs         []Log
	GasUsed         *big.Int
	GasPrice        *big.Int
	CumulativeGas   *big.Int
	TxIndex         *big.Int
	StartLogIndex   *big.Int
}

func CompareResults(res1 *TxResult, res2 *TxResult) []string {
	var differences []string
	differences = append(differences, CompareIncomingRequests(res1.IncomingRequest, res2.IncomingRequest)...)
	if res1.ResultCode != res2.ResultCode {
		differences = append(differences, fmt.Sprintf("different result code %v and %v", res1.ResultCode, res2.ResultCode))
	}
	if !bytes.Equal(res1.ReturnData, res2.ReturnData) {
		differences = append(differences, fmt.Sprintf("different return data 0x%X and 0x%X", res1.ReturnData, res2.ReturnData))
	}
	if len(res1.EVMLogs) != len(res2.EVMLogs) {

	} else {
		for i, log1 := range res1.EVMLogs {
			log2 := res2.EVMLogs[i]
			differences = append(differences, CompareLogs(log1, log2)...)
		}
	}
	if res1.GasUsed.Cmp(res2.GasUsed) != 0 {
		differences = append(differences, fmt.Sprintf("different gas used %v and %v", res1.GasUsed, res2.GasUsed))
	}
	if res1.GasPrice.Cmp(res2.GasPrice) != 0 {
		differences = append(differences, fmt.Sprintf("different gas price %v and %v", res1.GasPrice, res2.GasPrice))
	}
	if res1.CumulativeGas.Cmp(res2.CumulativeGas) != 0 {
		differences = append(differences, fmt.Sprintf("different cumulative gas %v and %v", res1.CumulativeGas, res2.CumulativeGas))
	}
	if res1.TxIndex.Cmp(res2.TxIndex) != 0 {
		differences = append(differences, fmt.Sprintf("different tx index %v and %v", res1.TxIndex, res2.TxIndex))
	}
	if res1.StartLogIndex.Cmp(res2.StartLogIndex) != 0 {
		differences = append(differences, fmt.Sprintf("different start log index %v and %v", res1.StartLogIndex, res2.StartLogIndex))
	}
	return differences
}

func (r *TxResult) String() string {
	return fmt.Sprintf(
		"TxResult(%v, %v, %v, %v, %v, %v)",
		r.IncomingRequest,
		r.ResultCode,
		hexutil.Encode(r.ReturnData),
		r.EVMLogs,
		r.GasUsed,
		r.GasPrice,
	)
}

func (r *TxResult) EthLogs(blockHash common.Hash) []*types.Log {
	evmLogs := make([]*types.Log, 0, len(r.EVMLogs))
	logIndex := r.StartLogIndex.Uint64()
	for _, l := range r.EVMLogs {
		ethLog := &types.Log{
			Address:     l.Address.ToEthAddress(),
			Topics:      common.NewEthHashesFromHashes(l.Topics),
			Data:        l.Data,
			BlockNumber: r.IncomingRequest.ChainTime.BlockNum.AsInt().Uint64(),
			TxHash:      r.IncomingRequest.MessageID.ToEthHash(),
			TxIndex:     uint(r.TxIndex.Uint64()),
			BlockHash:   blockHash.ToEthHash(),
			Index:       uint(logIndex),
		}
		logIndex++
		evmLogs = append(evmLogs, ethLog)
	}
	return evmLogs
}

func (r *TxResult) ToEthReceipt(blockHash common.Hash) *types.Receipt {
	contractAddress := ethcommon.Address{}
	if r.IncomingRequest.Kind == message.L2Type && r.ResultCode == ReturnCode {
		msg, err := message.L2Message{Data: r.IncomingRequest.Data}.AbstractMessage()
		if err == nil {
			if msg, ok := msg.(message.AbstractTransaction); ok {
				emptyAddress := common.Address{}
				if msg.Destination() == emptyAddress {
					copy(contractAddress[:], r.ReturnData[12:])
				}
			}
		}
	}

	status := uint64(0)
	if r.ResultCode == ReturnCode {
		status = 1
	}

	evmLogs := r.EthLogs(blockHash)
	return &types.Receipt{
		PostState:         []byte{0},
		Status:            status,
		CumulativeGasUsed: r.CumulativeGas.Uint64(),
		Bloom:             types.BytesToBloom(types.LogsBloom(evmLogs)),
		Logs:              evmLogs,
		TxHash:            r.IncomingRequest.MessageID.ToEthHash(),
		ContractAddress:   contractAddress,
		GasUsed:           r.GasUsed.Uint64(),
		BlockHash:         blockHash.ToEthHash(),
		BlockNumber:       r.IncomingRequest.ChainTime.BlockNum.AsInt(),
		TransactionIndex:  uint(r.TxIndex.Uint64()),
	}
}

func parseTxResult(l1MsgVal value.Value, resultInfo value.Value, gasInfo value.Value, chainInfo value.Value) (*TxResult, error) {
	resultTup, ok := resultInfo.(*value.TupleValue)
	if !ok || resultTup.Len() != 3 {
		return nil, errors.Errorf("advise expected result info tuple of length 3, but recieved %v", resultTup)
	}

	// Tuple size already verified above, so error can be ignored
	resultCode, _ := resultTup.GetByInt64(0)
	returnData, _ := resultTup.GetByInt64(1)
	evmLogs, _ := resultTup.GetByInt64(2)

	gasInfoTup, ok := gasInfo.(*value.TupleValue)
	if !ok || gasInfoTup.Len() != 2 {
		return nil, errors.Errorf("advise expected gas info tuple of length 2, but recieved %v", gasInfoTup)
	}

	// Tuple size already verified above, so error can be ignored
	gasUsed, _ := gasInfoTup.GetByInt64(0)
	gasPrice, _ := gasInfoTup.GetByInt64(1)

	chainInfoTup, ok := chainInfo.(*value.TupleValue)
	if !ok || chainInfoTup.Len() != 3 {
		return nil, errors.Errorf("advise expected tx block data tuple of length 3, but recieved %v", resultTup)
	}

	// Tuple size already verified above, so error can be ignored
	cumulativeGas, _ := chainInfoTup.GetByInt64(0)
	txIndex, _ := chainInfoTup.GetByInt64(1)
	startLogIndex, _ := chainInfoTup.GetByInt64(2)

	l1Msg, err := NewIncomingRequestFromValue(l1MsgVal)
	if err != nil {
		return nil, err
	}
	returnBytes, err := inbox.ByteArrayToBytes(returnData)
	if err != nil {
		return nil, errors.Wrap(err, "umarshalling return data")
	}
	logs, err := LogStackToLogs(evmLogs)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshaling logs")
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

	return &TxResult{
		IncomingRequest: l1Msg,
		ResultCode:      ResultType(resultCodeInt.BigInt().Uint64()),
		ReturnData:      returnBytes,
		EVMLogs:         logs,
		GasUsed:         gasUsedInt.BigInt(),
		GasPrice:        gasPriceInt.BigInt(),
		CumulativeGas:   cumulativeGasInt.BigInt(),
		TxIndex:         txIndexInt.BigInt(),
		StartLogIndex:   startLogIndexInt.BigInt(),
	}, nil
}

func NewResultFromValue(val value.Value) (Result, error) {
	tup, ok := val.(*value.TupleValue)
	if !ok || tup.Len() == 0 {
		return nil, errors.New("expected result to be nonempty tuple")
	}

	// Tuple size already verified above, so error can be ignored
	kind, _ := tup.GetByInt64(0)
	kindInt, ok := kind.(value.IntValue)
	if !ok {
		return nil, errors.New(" result kind must be an int")
	}

	if kindInt.BigInt().Uint64() == 0 {
		if tup.Len() != 5 {
			return nil, errors.Errorf("tx result expected tuple of length 5, but recieved len %v: %v", tup.Len(), tup)
		}

		// Tuple size already verified above, so error can be ignored
		l1MsgVal, _ := tup.GetByInt64(1)
		resultInfo, _ := tup.GetByInt64(2)
		gasInfo, _ := tup.GetByInt64(3)
		chainInfo, _ := tup.GetByInt64(4)
		return parseTxResult(l1MsgVal, resultInfo, gasInfo, chainInfo)
	} else if kindInt.BigInt().Uint64() == 1 {
		if tup.Len() != 8 {
			return nil, errors.Errorf("block result expected tuple of length 8, but received len %v: %v", tup.Len(), tup)
		}

		// Tuple size already verified above, so error can be ignored
		blockNum, _ := tup.GetByInt64(1)
		timestamp, _ := tup.GetByInt64(2)
		gasLimit, _ := tup.GetByInt64(3)
		blockStatsRaw, _ := tup.GetByInt64(4)
		chainStatsRaw, _ := tup.GetByInt64(5)
		gasStats, _ := tup.GetByInt64(6)
		previousHeight, _ := tup.GetByInt64(7)

		return parseBlockResult(blockNum, timestamp, gasLimit, blockStatsRaw, chainStatsRaw, gasStats, previousHeight)
	} else {
		return nil, errors.New("unknown result kind")
	}
}

func NewTxResultFromValue(val value.Value) (*TxResult, error) {
	res, err := NewResultFromValue(val)
	if err != nil {
		return nil, err
	}
	txRes, ok := res.(*TxResult)
	if !ok {
		return nil, errors.Errorf("got %T but expected TxResult", res)
	}
	return txRes, nil
}

func NewRandomResult(logCount int32) *TxResult {
	logs := make([]Log, 0, logCount)
	for i := int32(0); i < logCount; i++ {
		logs = append(logs, NewRandomLog(3))
	}
	return &TxResult{
		IncomingRequest: NewRandomIncomingRequest(),
		ResultCode:      ReturnCode,
		ReturnData:      common.RandBytes(200),
		EVMLogs:         logs,
		GasUsed:         common.RandBigInt(),
		GasPrice:        common.RandBigInt(),
		CumulativeGas:   common.RandBigInt(),
		TxIndex:         common.RandBigInt(),
		StartLogIndex:   common.RandBigInt(),
	}
}
