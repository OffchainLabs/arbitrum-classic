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
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type ResultType int

const (
	ReturnCode                ResultType = 0
	RevertCode                ResultType = 1
	CongestionCode            ResultType = 2
	InsufficientGasFundsCode  ResultType = 3
	InsufficientTxFundsCode   ResultType = 4
	BadSequenceCode           ResultType = 5
	InvalidMessageFormatCode  ResultType = 6
	ContractAlreadyExists     ResultType = 7
	ExceededTxGasLimit        ResultType = 8
	InsufficientGasForBaseFee ResultType = 9
	MinArbGasForContractTx    ResultType = 10
	GasPriceTooLow            ResultType = 11
	NoGasForAutoRedeem        ResultType = 12
	ForbiddenSender           ResultType = 13
)

func (r ResultType) String() string {
	switch r {
	case ReturnCode:
		return "Return"
	case RevertCode:
		return "Revert"
	case CongestionCode:
		return "Congestion"
	case InsufficientGasFundsCode:
		return "InsufficientGasFunds"
	case InsufficientTxFundsCode:
		return "InsufficientTxFunds"
	case BadSequenceCode:
		return "BadSequence"
	case InvalidMessageFormatCode:
		return "InvalidMessageFormat"
	case ContractAlreadyExists:
		return "ContractAlreadyExists"
	case ExceededTxGasLimit:
		return "ExceededTxGasLimit"
	case InsufficientGasForBaseFee:
		return "InsufficientGasForBaseFee"
	case MinArbGasForContractTx:
		return "MinArbGasForContractTx"
	case GasPriceTooLow:
		return "GasPriceTooLow"
	case NoGasForAutoRedeem:
		return "NoGasForAutoRedeem"
	default:
		return fmt.Sprintf("%v", int(r))
	}
}

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
	FeeStats        *FeeStats
}

type revertError struct {
	error
	reason interface{}
}

// ErrorCode returns the JSON error code for a revertal.
// See: https://github.com/ethereum/wiki/wiki/JSON-RPC-Error-Codes-Improvement-Proposal
func (e revertError) ErrorCode() int {
	return 3
}

// ErrorData returns the hex encoded revert reason.
func (e revertError) ErrorData() interface{} {
	return e.reason
}

type ganacheErrorData struct {
	Error  string `json:"error"`
	Return string `json:"return"`
	Reason string `json:"reason"`
}

func HandleCallError(res *TxResult, ganacheMode bool) error {
	if res == nil {
		logger.Warn().Msg("missing tx error result")
		return vm.ErrExecutionReverted
	}
	if len(res.ReturnData) > 0 {
		err := vm.ErrExecutionReverted
		reason := ""
		revertReason, unpackError := abi.UnpackRevert(res.ReturnData)
		if unpackError == nil {
			err = errors.Errorf("execution reverted: %v", revertReason)
			reason = revertReason
		}

		var errorReason interface{}
		if ganacheMode {
			errMap := make(map[string]ganacheErrorData)
			errMap[res.IncomingRequest.MessageID.String()] = ganacheErrorData{
				Error:  err.Error(),
				Return: hexutil.Encode(res.ReturnData),
				Reason: reason,
			}
			errorReason = errMap
		} else {
			errorReason = hexutil.Encode(res.ReturnData)
		}

		return revertError{
			error:  err,
			reason: errorReason,
		}
	} else if res.ResultCode == CongestionCode {
		return errors.New("tx dropped due to L2 congestion")
	} else if res.ResultCode == InsufficientTxFundsCode {
		return vm.ErrInsufficientBalance
	} else if res.ResultCode == InsufficientGasFundsCode {
		return errors.New("not enough funds for gas")
	} else if res.ResultCode == BadSequenceCode {
		return errors.New("invalid transaction nonce")
	} else if res.ResultCode == InvalidMessageFormatCode {
		return errors.New("invalid message format")
	} else if res.ResultCode == RevertCode {
		return vm.ErrExecutionReverted
	} else if res.ResultCode == GasPriceTooLow {
		return errors.New("gas price too low")
	} else if res.ResultCode == ForbiddenSender {
		return errors.New("forbidden sender address")
	} else {
		return errors.Errorf("execution reverted: error code %v", res.ResultCode)
	}
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
		"TxResult(request=%v, resultCode=%v, returnData=%v, evmLogs=%v, gasUsed=%v, gasPrice=%v)",
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
			BlockNumber: r.IncomingRequest.L2BlockNumber.Uint64(),
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

func (r *TxResult) CalcGasUsed() *big.Int {
	if r.FeeStats.Price.L2Computation.Cmp(big.NewInt(0)) == 0 {
		return r.GasUsed
	} else {
		return r.FeeStats.GasUsed()
	}
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
		GasUsed:           r.CalcGasUsed().Uint64(),
		BlockHash:         blockHash.ToEthHash(),
		BlockNumber:       r.IncomingRequest.L2BlockNumber,
		TransactionIndex:  uint(r.TxIndex.Uint64()),
	}
}

type FeeSet struct {
	L1Transaction *big.Int
	L1Calldata    *big.Int
	L2Storage     *big.Int
	L2Computation *big.Int
}

func (fs *FeeSet) Total() *big.Int {
	total := new(big.Int).Add(fs.L1Transaction, fs.L1Calldata)
	total = total.Add(total, fs.L2Storage)
	total = total.Add(total, fs.L2Computation)
	return total
}

func NewFeeSetFromValue(val value.Value) (*FeeSet, error) {
	tup, ok := val.(*value.TupleValue)
	if !ok || tup.Len() != 4 {
		return nil, errors.Errorf("expected fee set tuple of length 4, but recieved %v", val)
	}
	l1Transaction, _ := tup.GetByInt64(0)
	l1Calldata, _ := tup.GetByInt64(1)
	l2Storage, _ := tup.GetByInt64(2)
	l2Computation, _ := tup.GetByInt64(3)

	l1TransactionInt, ok := l1Transaction.(value.IntValue)
	if !ok {
		return nil, errors.New("l1Transaction must be an int")
	}
	l1CalldataInt, ok := l1Calldata.(value.IntValue)
	if !ok {
		return nil, errors.New("l1Calldata must be an int")
	}
	l2StorageInt, ok := l2Storage.(value.IntValue)
	if !ok {
		return nil, errors.New("l2Storage must be an int")
	}
	l2ComputationInt, ok := l2Computation.(value.IntValue)
	if !ok {
		return nil, errors.New("l2Computation must be an int")
	}

	return &FeeSet{
		L1Transaction: l1TransactionInt.BigInt(),
		L1Calldata:    l1CalldataInt.BigInt(),
		L2Storage:     l2StorageInt.BigInt(),
		L2Computation: l2ComputationInt.BigInt(),
	}, nil
}

type FeeStats struct {
	Price                  *FeeSet
	UnitsUsed              *FeeSet
	Paid                   *FeeSet
	Aggregator             *common.Address
	NoFeeGasEstimationMode bool
}

func (fs *FeeStats) String() string {
	return fmt.Sprintf("FeeStats{Prices=%v, Units=%v, Paid=%v, Aggregator=%v}", fs.Price, fs.UnitsUsed, fs.Paid, fs.Aggregator)
}

func (fs *FeeStats) PayTarget() *FeeSet {
	return &FeeSet{
		L1Transaction: new(big.Int).Mul(fs.Price.L1Transaction, fs.UnitsUsed.L1Transaction),
		L1Calldata:    new(big.Int).Mul(fs.Price.L1Calldata, fs.UnitsUsed.L1Calldata),
		L2Storage:     new(big.Int).Mul(fs.Price.L2Storage, fs.UnitsUsed.L2Storage),
		L2Computation: new(big.Int).Mul(fs.Price.L2Computation, fs.UnitsUsed.L2Computation),
	}
}

func (fs *FeeStats) GasUsed() *big.Int {
	return new(big.Int).Div(fs.Paid.Total(), fs.Price.L2Computation)
}

func (fs *FeeStats) TargetGasUsed() *big.Int {
	return new(big.Int).Div(fs.PayTarget().Total(), fs.Price.L2Computation)
}

func NewFeeStatsFromValue(val value.Value) (*FeeStats, error) {
	tup, ok := val.(*value.TupleValue)
	if !ok || tup.Len() < 4 || tup.Len() > 5 {
		return nil, errors.Errorf("expected gas fee tuple of length 4 or 5, but recieved %v", val)
	}
	pricesVal, _ := tup.GetByInt64(0)
	unitsVal, _ := tup.GetByInt64(1)
	paidVal, _ := tup.GetByInt64(2)
	aggregator, _ := tup.GetByInt64(3)
	noFeeGasEstimationMode := false
	if tup.Len() == 5 {
		noFeeGasEstimationModeVal, _ := tup.GetByInt64(4)
		var err error
		noFeeGasEstimationMode, err = NewBoolFromValue(noFeeGasEstimationModeVal)
		if err != nil {
			return nil, err
		}
	}

	prices, err := NewFeeSetFromValue(pricesVal)
	if err != nil {
		return nil, err
	}
	units, err := NewFeeSetFromValue(unitsVal)
	if err != nil {
		return nil, err
	}
	paid, err := NewFeeSetFromValue(paidVal)
	if err != nil {
		return nil, err
	}
	aggregatorInt, ok := aggregator.(value.IntValue)
	if !ok {
		return nil, errors.New("aggregator must be an int")
	}
	rawAggregatorAddress := inbox.NewAddressFromInt(aggregatorInt)
	blankAddress := common.Address{}
	var aggAddress *common.Address
	if rawAggregatorAddress != blankAddress {
		aggAddress = &rawAggregatorAddress
	}
	return &FeeStats{
		Price:                  prices,
		UnitsUsed:              units,
		Paid:                   paid,
		Aggregator:             aggAddress,
		NoFeeGasEstimationMode: noFeeGasEstimationMode,
	}, nil
}

func parseTxResult(l1MsgVal value.Value, resultInfo value.Value, gasInfo value.Value, chainInfo value.Value, feeStatsVal value.Value) (*TxResult, error) {
	resultTup, ok := resultInfo.(*value.TupleValue)
	if !ok || resultTup.Len() != 3 {
		return nil, errors.Errorf("expected result info tuple of length 3, but recieved %v", resultTup)
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

	feeStats, err := NewFeeStatsFromValue(feeStatsVal)
	if err != nil {
		return nil, err
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
		FeeStats:        feeStats,
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
		if tup.Len() != 6 {
			return nil, errors.Errorf("tx result expected tuple of length 6, but recieved len %v: %v", tup.Len(), tup)
		}

		// Tuple size already verified above, so error can be ignored
		l1MsgVal, _ := tup.GetByInt64(1)
		resultInfo, _ := tup.GetByInt64(2)
		gasInfo, _ := tup.GetByInt64(3)
		chainInfo, _ := tup.GetByInt64(4)
		feeStats, _ := tup.GetByInt64(5)
		return parseTxResult(l1MsgVal, resultInfo, gasInfo, chainInfo, feeStats)
	} else if kindInt.BigInt().Uint64() == 1 {
		if tup.Len() != 8 {
			return nil, errors.Errorf("block result expected tuple of length 8, but received len %v: %v", tup.Len(), tup)
		}

		// Tuple size already verified above, so error can be ignored
		blockNum, _ := tup.GetByInt64(1)
		timestamp, _ := tup.GetByInt64(2)
		blockStatsRaw, _ := tup.GetByInt64(3)
		chainStatsRaw, _ := tup.GetByInt64(4)
		gasStats, _ := tup.GetByInt64(5)
		previousHeight, _ := tup.GetByInt64(6)
		l1BlockNum, _ := tup.GetByInt64(7)

		return parseBlockResult(blockNum, timestamp, blockStatsRaw, chainStatsRaw, gasStats, previousHeight, l1BlockNum)
	} else if kindInt.BigInt().Uint64() == 2 {
		return NewSendResultFromValue(tup)
	} else if kindInt.BigInt().Uint64() == 3 {
		return NewMerkleRootLogResultFromValue(tup)
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
