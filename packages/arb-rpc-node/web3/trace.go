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

package web3

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	arbcommon "github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type TraceAction struct {
	CallType string          `json:"callType,omitempty"`
	From     common.Address  `json:"from"`
	Gas      hexutil.Uint64  `json:"gas"`
	Input    hexutil.Bytes   `json:"input,omitempty"`
	Init     hexutil.Bytes   `json:"init,omitempty"`
	To       *common.Address `json:"to,omitempty"`
	Value    *hexutil.Big    `json:"value"`
}

type TraceCallResult struct {
	Address *common.Address `json:"address,omitempty"`
	Code    *hexutil.Bytes  `json:"code,omitempty"`
	GasUsed hexutil.Uint64  `json:"gasUsed"`
	Output  *hexutil.Bytes  `json:"output,omitempty"`
}

type TraceFrame struct {
	Action              TraceAction      `json:"action"`
	BlockHash           *hexutil.Bytes   `json:"blockHash,omitempty"`
	BlockNumber         *uint64          `json:"blockNumber,omitempty"`
	Result              *TraceCallResult `json:"result,omitempty"`
	Error               *string          `json:"error,omitempty"`
	Subtraces           int              `json:"subtraces"`
	TraceAddress        []int            `json:"traceAddress"`
	TransactionHash     *hexutil.Bytes   `json:"transactionHash,omitempty"`
	TransactionPosition *uint64          `json:"transactionPosition,omitempty"`
	Type                string           `json:"type"`
}

type TraceResult struct {
	Output             hexutil.Bytes     `json:"output"`
	StateDiff          *int              `json:"stateDiff"`
	Trace              []TraceFrame      `json:"trace"`
	VmTrace            *int              `json:"vmTrace"`
	DestroyedContracts *[]common.Address `json:"destroyedContracts"`
}

type Trace struct {
	s          *Server
	coreConfig *configuration.Core
}

func NewTracer(s *Server, coreConfig *configuration.Core) *Trace {
	return &Trace{s: s, coreConfig: coreConfig}
}

func splitEmissionsByLog(emissions []core.MachineEmission) map[uint64][]value.Value {
	splitEmissions := make(map[uint64][]value.Value)
	for _, emission := range emissions {
		splitEmissions[emission.LogCount.Uint64()] = append(splitEmissions[emission.LogCount.Uint64()], emission.Value)
	}
	return splitEmissions
}

func extractValuesFromEmissions(emissions []core.MachineEmission) []value.Value {
	values := make([]value.Value, 0, len(emissions))
	for _, emission := range emissions {
		values = append(values, emission.Value)
	}
	return values
}

func extractTrace(debugPrints []value.Value) (*evm.EVMTrace, error) {
	var trace *evm.EVMTrace
	for _, debugPrint := range debugPrints {
		parsedLog, err := evm.NewLogLineFromValue(debugPrint)
		if err != nil {
			return nil, err
		}
		foundTrace, ok := parsedLog.(*evm.EVMTrace)
		if ok {
			if trace != nil {
				return nil, errors.New("found multiple traces")
			}
			trace = foundTrace
		}
	}
	if trace == nil {
		return nil, errors.New("found no trace")
	}
	return trace, nil
}

func getDestroyedContracts(snap *snapshot.Snapshot, frames []TraceFrame) ([]common.Address, error) {
	maybeDestroyedContracts := make(map[common.Address]struct{})
	for _, frame := range frames {
		if frame.Type == "call" && frame.Action.CallType == "call" {
			maybeDestroyedContracts[*frame.Action.To] = struct{}{}
		}
		if frame.Type == "create" && frame.Result != nil {
			maybeDestroyedContracts[*frame.Result.Address] = struct{}{}
		}
	}

	deletedContracts := make([]common.Address, 0, len(maybeDestroyedContracts))
	for con := range maybeDestroyedContracts {
		txCount, err := snap.GetTransactionCount(arbcommon.NewAddressFromEth(con))
		if err != nil {
			return nil, err
		}
		if txCount.Sign() == 0 {
			// If nonce is 0, contract must have been destroyed and if it was destroyed, the nonce must be 0
			deletedContracts = append(deletedContracts, con)
		}
	}
	return deletedContracts, nil
}

func renderTraceFrames(txRes *evm.TxResult, trace *evm.EVMTrace) ([]TraceFrame, error) {
	receipt := txRes.ToEthReceipt(arbcommon.Hash{})
	frame, err := trace.FrameTree()
	if err != nil {
		return nil, err
	}

	type trackedFrame struct {
		f            evm.Frame
		traceAddress []int
	}

	frames := []trackedFrame{{f: frame, traceAddress: make([]int, 0)}}
	resFrames := make([]TraceFrame, 0)
	for len(frames) > 0 {
		frame := frames[0]
		frames = frames[1:]

		callFrame := frame.f.GetCallFrame()
		action := TraceAction{
			From:  callFrame.Call.From.ToEthAddress(),
			Gas:   hexutil.Uint64(callFrame.Call.Gas.Uint64()),
			Value: (*hexutil.Big)(callFrame.Call.Value),
		}

		var result *TraceCallResult
		var callErr *string
		if callFrame.Return.Result == evm.ReturnCode {
			result = &TraceCallResult{
				GasUsed: hexutil.Uint64(callFrame.Return.GasUsed.Uint64()),
			}
		} else {
			tmp := callFrame.Return.Result.String()
			callErr = &tmp
		}

		var frameType string
		switch frame := frame.f.(type) {
		case *evm.CallFrame:
			// Top level call could actually be contract creation
			if len(resFrames) == 0 && callFrame.Call.To == nil {
				frameType = "create"
				action.Init = callFrame.Call.Data
				if result != nil {
					result.Address = &receipt.ContractAddress
					result.Code = (*hexutil.Bytes)(&callFrame.Return.ReturnData)
				}
			} else {
				frameType = "call"
				action.Input = callFrame.Call.Data
				toTmp := callFrame.Call.To.ToEthAddress()
				action.To = &toTmp
				action.CallType = callFrame.Call.Type.RPCString()
				if result != nil {
					result.Output = (*hexutil.Bytes)(&callFrame.Return.ReturnData)
				}
			}
		case *evm.CreateFrame:
			frameType = "create"
			action.Init = frame.Create.Code
			if result != nil {
				tmp := frame.Create.ContractAddress.ToEthAddress()
				result.Address = &tmp
				result.Code = (*hexutil.Bytes)(&callFrame.Return.ReturnData)
			}
		case *evm.Create2Frame:
			frameType = "create"
			action.Init = frame.Create.Code
			if result != nil {
				tmp := frame.Create.ContractAddress.ToEthAddress()
				result.Address = &tmp
				result.Code = (*hexutil.Bytes)(&callFrame.Return.ReturnData)
			}
		}

		resFrames = append(resFrames, TraceFrame{
			Action:       action,
			Result:       result,
			Error:        callErr,
			Subtraces:    len(callFrame.Nested),
			TraceAddress: frame.traceAddress,
			Type:         frameType,
		})
		for i, nested := range callFrame.Nested {
			nestedTrace := make([]int, 0)
			nestedTrace = append(nestedTrace, frame.traceAddress...)
			nestedTrace = append(nestedTrace, i)
			frames = append(frames, trackedFrame{f: nested, traceAddress: nestedTrace})
		}
	}
	return resFrames, nil
}

func authenticateTraceType(traceTypes []string) (bool, error) {
	types := make(map[string]struct{})
	for _, typ := range traceTypes {
		if typ != "trace" && typ != "deletedContracts" {
			return false, errors.Errorf("unsupported trace type: %v", typ)
		}
		types[typ] = struct{}{}
	}
	if _, found := types["trace"]; !found {
		return false, errors.New("must specify trace type as 'trace'")
	}
	_, traceDestroys := types["deletedContracts"]
	return traceDestroys, nil
}

type CallTraceRequest struct {
	callArgs   CallTxArgs
	traceTypes []string
}

func (at *CallTraceRequest) UnmarshalJSON(b []byte) error {
	fields := []interface{}{&at.callArgs, &at.traceTypes}
	if err := json.Unmarshal(b, &fields); err != nil {
		return err
	}
	if len(fields) != 2 {
		return errors.New("expected two arguments per call")
	}
	return nil
}

func (t *Trace) transaction(txHash hexutil.Bytes) (*rawTxTrace, *machine.BlockInfo, error) {
	res, blockInfo, _, logNumber, err := t.s.getTransactionInfoByHash(txHash)
	if err != nil || res == nil {
		return nil, nil, err
	}
	blockNumber := res.IncomingRequest.L2BlockNumber.Uint64()
	cursor, err := t.s.srv.GetExecutionCursorAtEndOfBlock(blockNumber-1, true)
	if err != nil {
		return nil, nil, err
	}
	maxGas := int64(t.coreConfig.CheckpointMaxExecutionGas)
	if maxGas == 0 {
		maxGas = 100000000000
	}
	debugPrints, err := t.s.srv.AdvanceExecutionCursorWithTracing(
		cursor,
		big.NewInt(maxGas),
		true,
		true,
		logNumber,
		new(big.Int).Add(logNumber, big.NewInt(1)),
	)
	if err != nil {
		return nil, nil, err
	}
	vmTrace, err := extractTrace(extractValuesFromEmissions(debugPrints))
	if err != nil {
		return nil, nil, err
	}
	frames, err := renderTraceFrames(res, vmTrace)
	return &rawTxTrace{
		frames: frames,
		res:    res,
	}, blockInfo, err
}

type rawTxTrace struct {
	frames []TraceFrame
	res    *evm.TxResult
}

func (t *Trace) block(blockNum rpc.BlockNumberOrHash) ([]*rawTxTrace, *machine.BlockInfo, error) {
	blockInfo, err := t.s.blockInfoForNumberOrHash(blockNum)
	if err != nil || blockInfo == nil {
		return nil, nil, err
	}
	blockLog, txResults, err := t.s.srv.GetMachineBlockResults(blockInfo)
	if err != nil {
		return nil, nil, err
	}

	cursor, err := t.s.srv.GetExecutionCursorAtEndOfBlock(blockInfo.Header.Number.Uint64()-1, true)
	if err != nil {
		return nil, nil, err
	}
	maxGas := int64(t.coreConfig.CheckpointMaxExecutionGas)
	if maxGas == 0 {
		maxGas = 100000000000
	}

	firstIndex := blockLog.FirstAVMLog()
	lastIndex := new(big.Int).Add(blockLog.FirstAVMLog(), blockLog.BlockStats.TxCount)
	debugPrints, err := t.s.srv.AdvanceExecutionCursorWithTracing(
		cursor,
		big.NewInt(maxGas),
		true,
		true,
		firstIndex,
		lastIndex,
	)

	res := make([]*rawTxTrace, 0, len(debugPrints))
	for logIndex, debugPrints := range splitEmissionsByLog(debugPrints) {
		if logIndex < firstIndex.Uint64() {
			return nil, nil, errors.New("expected log index to be greater the first in the block")
		}
		if logIndex >= lastIndex.Uint64() {
			return nil, nil, errors.Errorf("expected log index to be less then the last in the block")
		}
		logOffset := logIndex - firstIndex.Uint64()
		txRes := txResults[logOffset]
		trace, err := extractTrace(debugPrints)
		failMsg := logger.
			Warn().
			Uint64("block", blockInfo.Header.Number.Uint64()).
			Str("txhash", txRes.IncomingRequest.MessageID.String()).
			Err(err)
		if err != nil {
			failMsg.Msg("error getting trace for transaction")
			continue
		}
		frames, err := renderTraceFrames(txRes, trace)
		if err != nil {
			failMsg.Msg("error rending trace for transaction")
			continue
		}
		res = append(res, &rawTxTrace{
			frames: frames,
			res:    txRes,
		})
	}
	return res, blockInfo, nil
}

func (t *Trace) handleCallRequest(callArgs CallTxArgs, traceDestroys bool, snap *snapshot.Snapshot) (*TraceResult, error) {
	from, msg := buildCallMsg(callArgs)
	// We're mutating so we need unique ownership
	snap = snap.Clone()
	callRes, debugPrints, err := snap.AddContractMessage(msg, from, t.s.maxAVMGas)
	if err != nil {
		return nil, err
	}
	if callRes.ResultCode != evm.ReturnCode && callRes.ResultCode != evm.RevertCode {
		return nil, evm.HandleCallError(callRes, t.s.ganacheMode)
	}
	vmTrace, err := extractTrace(debugPrints)
	if err != nil {
		return nil, err
	}
	frames, err := renderTraceFrames(callRes, vmTrace)
	if err != nil {
		return nil, err
	}
	var destroyed *[]common.Address
	if traceDestroys {
		destroyedTmp, err := getDestroyedContracts(snap, frames)
		if err != nil {
			return nil, err
		}
		destroyed = &destroyedTmp
	}

	return &TraceResult{
		Output:             callRes.ReturnData,
		Trace:              frames,
		DestroyedContracts: destroyed,
	}, nil
}

func (t *Trace) Call(callArgs CallTxArgs, traceTypes []string, blockNum rpc.BlockNumberOrHash) (*TraceResult, error) {
	traceDestroys, err := authenticateTraceType(traceTypes)
	if err != nil {
		return nil, err
	}

	snap, err := t.s.getSnapshotForNumberOrHash(blockNum)
	if err != nil {
		return nil, err
	}
	return t.handleCallRequest(callArgs, traceDestroys, snap)
}

func (t *Trace) CallMany(calls []*CallTraceRequest, blockNum rpc.BlockNumberOrHash) ([]*TraceResult, error) {
	traceDestroys := make([]bool, 0, len(calls))
	for _, call := range calls {
		traceDestroy, err := authenticateTraceType(call.traceTypes)
		if err != nil {
			return nil, err
		}
		traceDestroys = append(traceDestroys, traceDestroy)
	}
	snap, err := t.s.getSnapshotForNumberOrHash(blockNum)
	if err != nil {
		return nil, err
	}

	traces := make([]*TraceResult, 0, len(calls))
	for i, call := range calls {
		frame, err := t.handleCallRequest(call.callArgs, traceDestroys[i], snap)
		if err != nil {
			return nil, err
		}
		traces = append(traces, frame)
	}
	return traces, nil
}

func (t *Trace) ReplayBlockTransactions(blockNum rpc.BlockNumberOrHash, traceTypes []string) ([]*TraceResult, error) {
	traceDestroys, err := authenticateTraceType(traceTypes)
	if err != nil {
		return nil, err
	}
	// TODO: Handle destroyed contract tracing
	_ = traceDestroys
	txTraces, blockInfo, err := t.block(blockNum)
	if err != nil {
		return nil, err
	}
	results := make([]*TraceResult, 0)
	for _, txTrace := range txTraces {
		chainContext := newChainContext(txTrace.res, blockInfo)
		for i := range txTrace.frames {
			txTrace.frames[i].TransactionHash = chainContext.transactionHash
		}
		results = append(results, &TraceResult{
			Output: txTrace.res.ReturnData,
			Trace:  txTrace.frames,
		})
	}
	return results, nil
}

func (t *Trace) ReplayTransaction(txHash hexutil.Bytes, traceTypes []string) (*TraceResult, error) {
	traceDestroys, err := authenticateTraceType(traceTypes)
	if err != nil {
		return nil, err
	}
	// TODO: Handle destroyed contract tracing
	_ = traceDestroys
	txTrace, _, err := t.transaction(txHash)
	if err != nil || txTrace.res == nil {
		return nil, err
	}
	return &TraceResult{
		Output: txTrace.res.ReturnData,
		Trace:  txTrace.frames,
	}, nil
}

type chainContext struct {
	blockHash           *hexutil.Bytes
	blockNumber         *uint64
	transactionHash     *hexutil.Bytes
	transactionPosition *uint64
}

func newChainContext(res *evm.TxResult, blockInfo *machine.BlockInfo) *chainContext {
	blockHash := hexutil.Bytes(blockInfo.Header.Hash().Bytes())
	blockNumber := res.IncomingRequest.L2BlockNumber.Uint64()
	txIndex := res.TxIndex.Uint64()
	txHash := hexutil.Bytes(res.IncomingRequest.MessageID.Bytes())
	return &chainContext{
		blockHash:           &blockHash,
		blockNumber:         &blockNumber,
		transactionHash:     &txHash,
		transactionPosition: &txIndex,
	}
}

func addChainContext(frame *TraceFrame, context *chainContext) {
	frame.TransactionHash = context.transactionHash
	frame.TransactionPosition = context.transactionPosition
	frame.BlockNumber = context.blockNumber
	frame.BlockHash = context.blockHash
}

func (t *Trace) Transaction(txHash hexutil.Bytes) ([]TraceFrame, error) {
	txTrace, blockInfo, err := t.transaction(txHash)
	if err != nil || txTrace == nil {
		return nil, err
	}
	chainContext := newChainContext(txTrace.res, blockInfo)
	for i := range txTrace.frames {
		addChainContext(&txTrace.frames[i], chainContext)
	}
	return txTrace.frames, nil
}

func (t *Trace) Get(txHash hexutil.Bytes, path []hexutil.Uint64) (*TraceFrame, error) {
	frames, err := t.Transaction(txHash)
	if err != nil {
		return nil, err
	}
	for _, frame := range frames {
		if len(path) != len(frame.TraceAddress) {
			continue
		}
		for i, addr := range frame.TraceAddress {
			if uint64(path[i]) != uint64(addr) {
				continue
			}
		}
		return &frame, nil
	}
	return nil, nil
}

func (t *Trace) Block(blockNum rpc.BlockNumberOrHash) ([]TraceFrame, error) {
	txTraces, blockInfo, err := t.block(blockNum)
	if err != nil {
		return nil, err
	}
	traces := make([]TraceFrame, 0)
	for _, txTrace := range txTraces {
		chainContext := newChainContext(txTrace.res, blockInfo)
		for i := range txTrace.frames {
			addChainContext(&txTrace.frames[i], chainContext)
		}
		traces = append(traces, txTrace.frames...)
	}
	return traces, nil
}
