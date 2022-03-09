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
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	arbcommon "github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type TraceAction struct {
	CallType string          `json:"callType,omitempty"`
	From     common.Address  `json:"from"`
	Gas      hexutil.Uint64  `json:"gas"`
	Input    *hexutil.Bytes  `json:"input,omitempty"`
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

func getDestroyedContracts(ctx context.Context, snap *snapshot.Snapshot, frames []TraceFrame) ([]common.Address, error) {
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
		txCount, err := snap.GetTransactionCount(ctx, arbcommon.NewAddressFromEth(con))
		if err != nil {
			return nil, err
		}
		if txCount.Sign() == 0 {
			// If nonce is 0, contract must have been destroyed and if it was destroyed, the nonce must be 0
			deletedContracts = append(deletedContracts, con)
		}
	}
	sort.Slice(deletedContracts, func(i, j int) bool {
		return bytes.Compare(deletedContracts[i].Bytes(), deletedContracts[j].Bytes()) < 0
	})
	return deletedContracts, nil
}

func renderTraceFrames(txRes *evm.TxResult, trace *evm.EVMTrace) ([]TraceFrame, error) {
	frame, err := trace.FrameTree()
	if err != nil || frame == nil {
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
			if len(resFrames) == 0 && txRes.IsContractCreation() {
				frameType = "create"
				// Call frame has no input for contract construction
				if txRes.IncomingRequest.Kind == message.L2Type || txRes.IncomingRequest.Kind == message.EthDepositTxType {
					abstractMessage, err := message.L2Message{Data: txRes.IncomingRequest.Data}.AbstractMessage()
					if err == nil {
						if msg, ok := abstractMessage.(message.EthConvertable); ok {
							action.Init = msg.AsEthTx().Data()
						}
					}
				}

				if result != nil {
					topLevelContractAddress, gotAddress := txRes.GetCreatedContractAddress()
					if gotAddress {
						result.Address = &topLevelContractAddress
					}
					// Return data contains the created contract address so we can't get the created code from that
					// We'll get it by querying the code instead
				}
			} else {
				frameType = "call"
				tmpInput := hexutil.Bytes(callFrame.Call.Data)
				action.Input = &tmpInput
				var toTmp common.Address
				if callFrame.Call.To != nil {
					toTmp = callFrame.Call.To.ToEthAddress()
				}
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
	sort.SliceStable(resFrames, func(i, j int) bool {
		addr1 := resFrames[i].TraceAddress
		addr2 := resFrames[j].TraceAddress
		maxLength := len(addr1)
		if len(addr2) > maxLength {
			maxLength = len(addr2)
		}
		for i := 0; i < maxLength; i++ {
			if i >= len(addr1) {
				return true
			}
			if i >= len(addr2) {
				return false
			}
			if addr1[i] < addr2[i] {
				return true
			}
		}
		return false
	})
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

func (t *Trace) getSnapAfterTx(ctx context.Context, cursor core.ExecutionCursor) (*snapshot.Snapshot, error) {
	mach, err := t.s.srv.GetLookup().TakeMachine(cursor)
	if err != nil {
		return nil, err
	}
	snapTime := inbox.ChainTime{
		BlockNum:  arbcommon.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}
	endOfBlock := message.NewInboxMessage(
		message.EndBlockMessage{},
		arbcommon.Address{},
		big.NewInt(0),
		big.NewInt(0),
		inbox.ChainTime{
			BlockNum:  arbcommon.NewTimeBlocksInt(0),
			Timestamp: big.NewInt(0),
		},
	)

	_, _, _, err = mach.ExecuteAssertionAdvanced(
		ctx,
		1000000000000,
		true,
		[]inbox.InboxMessage{endOfBlock},
		nil,
		true,
		false,
	)
	if err != nil {
		return nil, err
	}
	return snapshot.NewSnapshot(ctx, mach, snapTime, big.NewInt(math.MaxInt64))
}

func needsTopLevelCreate(frames []TraceFrame) bool {
	return len(frames) > 0 &&
		frames[0].Type == "create" &&
		frames[0].Result != nil &&
		frames[0].Result.Address != nil
}

func fillInTopLevelCreate(ctx context.Context, frames []TraceFrame, snap *snapshot.Snapshot) {
	createdCode, err := snap.GetCode(ctx, arbcommon.NewAddressFromEth(*frames[0].Result.Address))
	if err != nil {
		logger.Warn().Msg("failed to retrieve code for contract")
	} else {
		frames[0].Result.Code = (*hexutil.Bytes)(&createdCode)
	}
}

func (t *Trace) traceTransaction(ctx context.Context, cursor core.ExecutionCursor, res *evm.TxResult, logNumber *big.Int, traceDestroyed bool) (*rawTxTrace, error) {
	maxGas := int64(t.coreConfig.CheckpointMaxExecutionGas)
	if maxGas == 0 {
		maxGas = 100000000000
	}
	debugPrints, err := t.s.srv.GetLookup().AdvanceExecutionCursorWithTracing(
		cursor,
		big.NewInt(maxGas),
		true,
		true,
		logNumber,
		new(big.Int).Add(logNumber, big.NewInt(1)),
	)
	if err != nil {
		return nil, err
	}
	vmTrace, err := extractTrace(extractValuesFromEmissions(debugPrints))
	if err != nil {
		return nil, err
	}
	frames, err := renderTraceFrames(res, vmTrace)
	if err != nil {
		return nil, err
	}

	var snap *snapshot.Snapshot
	neadsCode := needsTopLevelCreate(frames)
	if neadsCode || traceDestroyed {
		snap, err = t.getSnapAfterTx(ctx, cursor.Clone())
		if err != nil {
			return nil, err
		}
	}
	if neadsCode {
		fillInTopLevelCreate(ctx, frames, snap)
	}

	var destroyed *[]common.Address
	if traceDestroyed {
		destroyedTmp, err := getDestroyedContracts(ctx, snap, frames)
		if err != nil {
			return nil, err
		}
		destroyed = &destroyedTmp
	}

	return &rawTxTrace{
		frames:    frames,
		res:       res,
		destroyed: destroyed,
	}, nil
}

func (t *Trace) transaction(ctx context.Context, txHash hexutil.Bytes, traceDestroyed bool) (*rawTxTrace, *machine.BlockInfo, error) {
	res, blockInfo, _, logNumber, err := t.s.getTransactionInfoByHash(txHash)
	if err != nil || res == nil {
		return nil, nil, err
	}
	blockNumber := res.IncomingRequest.L2BlockNumber.Uint64()
	cursor, err := t.s.srv.GetLookup().GetExecutionCursorAtEndOfBlock(blockNumber-1, true)
	if err != nil {
		return nil, nil, err
	}
	txTrace, err := t.traceTransaction(ctx, cursor, res, logNumber, traceDestroyed)
	return txTrace, blockInfo, err
}

type rawTxTrace struct {
	frames    []TraceFrame
	res       *evm.TxResult
	destroyed *[]common.Address
}

func (t *Trace) block(ctx context.Context, blockNum rpc.BlockNumberOrHash, traceDestroyed bool) ([]*rawTxTrace, *machine.BlockInfo, core.ExecutionCursor, error) {
	blockInfo, err := t.s.blockInfoForNumberOrHash(blockNum)
	if err != nil || blockInfo == nil {
		return nil, nil, nil, err
	}
	blockLog, txResults, err := t.s.srv.GetMachineBlockResults(blockInfo)
	if err != nil {
		return nil, nil, nil, err
	}

	cursor, err := t.s.srv.GetLookup().GetExecutionCursorAtEndOfBlock(blockInfo.Header.Number.Uint64()-1, true)
	if err != nil {
		return nil, nil, nil, err
	}
	maxGas := int64(t.coreConfig.CheckpointMaxExecutionGas)
	if maxGas == 0 {
		maxGas = 100000000000
	}

	logIndex := blockLog.FirstAVMLog()
	res := make([]*rawTxTrace, 0, len(txResults))
	for i := uint64(0); i < blockLog.BlockStats.TxCount.Uint64(); i++ {
		txRes := txResults[i]
		failMsg := logger.
			Warn().
			Uint64("block", blockInfo.Header.Number.Uint64()).
			Str("txhash", txRes.IncomingRequest.MessageID.String()).
			Err(err)
		txTrace, err := t.traceTransaction(ctx, cursor, txRes, logIndex, traceDestroyed)
		if err != nil {
			failMsg.Msg("error getting trace for transaction")
			continue
		}
		res = append(res, txTrace)
		logIndex.Add(logIndex, big.NewInt(1))
	}
	return res, blockInfo, cursor, nil
}

func (t *Trace) handleCallRequest(ctx context.Context, callArgs CallTxArgs, traceDestroys bool, snap *snapshot.Snapshot) (*TraceResult, error) {
	from, msg := buildCallMsg(callArgs)
	// We're mutating so we need unique ownership
	snap = snap.Clone()
	callRes, debugPrints, err := snap.AddContractMessage(ctx, msg, from, t.s.maxAVMGas)
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
	if needsTopLevelCreate(frames) {
		fillInTopLevelCreate(ctx, frames, snap)
	}

	var destroyed *[]common.Address
	if traceDestroys {
		destroyedTmp, err := getDestroyedContracts(ctx, snap, frames)
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

func (t *Trace) Call(ctx context.Context, callArgs CallTxArgs, traceTypes []string, blockNum rpc.BlockNumberOrHash) (*TraceResult, error) {
	traceDestroys, err := authenticateTraceType(traceTypes)
	if err != nil {
		return nil, err
	}

	snap, err := t.s.getSnapshotForNumberOrHash(ctx, blockNum)
	if err != nil {
		return nil, err
	}
	return t.handleCallRequest(ctx, callArgs, traceDestroys, snap)
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

func (t *Trace) CallMany(ctx context.Context, calls []*CallTraceRequest, blockNum rpc.BlockNumberOrHash) ([]*TraceResult, error) {
	traceDestroys := make([]bool, 0, len(calls))
	for _, call := range calls {
		traceDestroy, err := authenticateTraceType(call.traceTypes)
		if err != nil {
			return nil, err
		}
		traceDestroys = append(traceDestroys, traceDestroy)
	}
	snap, err := t.s.getSnapshotForNumberOrHash(ctx, blockNum)
	if err != nil {
		return nil, err
	}

	traces := make([]*TraceResult, 0, len(calls))
	for i, call := range calls {
		frame, err := t.handleCallRequest(ctx, call.callArgs, traceDestroys[i], snap)
		if err != nil {
			return nil, err
		}
		traces = append(traces, frame)
	}
	return traces, nil
}

func (t *Trace) ReplayBlockTransactions(ctx context.Context, blockNum rpc.BlockNumberOrHash, traceTypes []string) ([]*TraceResult, error) {
	traceDestroys, err := authenticateTraceType(traceTypes)
	if err != nil {
		return nil, err
	}
	txTraces, blockInfo, _, err := t.block(ctx, blockNum, traceDestroys)
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
			Output:             txTrace.res.ReturnData,
			Trace:              txTrace.frames,
			DestroyedContracts: txTrace.destroyed,
		})
	}
	return results, nil
}

func (t *Trace) ReplayTransaction(ctx context.Context, txHash hexutil.Bytes, traceTypes []string) (*TraceResult, error) {
	traceDestroys, err := authenticateTraceType(traceTypes)
	if err != nil {
		return nil, err
	}
	txTrace, _, err := t.transaction(ctx, txHash, traceDestroys)
	if err != nil || txTrace.res == nil {
		return nil, err
	}

	return &TraceResult{
		Output:             txTrace.res.ReturnData,
		Trace:              txTrace.frames,
		DestroyedContracts: txTrace.destroyed,
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

func (t *Trace) Transaction(ctx context.Context, txHash hexutil.Bytes) ([]TraceFrame, error) {
	txTrace, blockInfo, err := t.transaction(ctx, txHash, false)
	if err != nil || txTrace == nil {
		return nil, err
	}
	chainContext := newChainContext(txTrace.res, blockInfo)
	for i := range txTrace.frames {
		addChainContext(&txTrace.frames[i], chainContext)
	}
	return txTrace.frames, nil
}

func (t *Trace) Get(ctx context.Context, txHash hexutil.Bytes, path []hexutil.Uint64) (*TraceFrame, error) {
	frames, err := t.Transaction(ctx, txHash)
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

func (t *Trace) Block(ctx context.Context, blockNum rpc.BlockNumberOrHash) ([]TraceFrame, error) {
	txTraces, blockInfo, _, err := t.block(ctx, blockNum, false)
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

type FilterRequest struct {
	FromBlock   *rpc.BlockNumberOrHash `json:"fromBlock"`
	ToBlock     *rpc.BlockNumberOrHash `json:"toBlock"`
	FromAddress *[]common.Address      `json:"fromAddress"`
	ToAddress   *[]common.Address      `json:"toAddress"`
	After       *uint64                `json:"after"`
	Count       *uint64                `json:"count"`
}

func (t *Trace) Filter(ctx context.Context, filter *FilterRequest) ([]TraceFrame, error) {
	fromBlock := rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber)
	if filter.FromBlock != nil {
		fromBlock = *filter.FromBlock
	}
	toBlock := rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber)
	if filter.ToBlock != nil {
		toBlock = *filter.ToBlock
	}
	fromBlockInfo, err := t.s.blockInfoForNumberOrHash(fromBlock)
	if err != nil {
		return nil, err
	}
	toBlockInfo, err := t.s.blockInfoForNumberOrHash(toBlock)
	if err != nil {
		return nil, err
	}
	if fromBlockInfo.LogCount > toBlockInfo.LogCount {
		return nil, nil
	}
	start := rpc.BlockNumber(fromBlockInfo.Header.Number.Int64())
	end := rpc.BlockNumber(toBlockInfo.Header.Number.Int64())

	fromAddrFilter := make(map[common.Address]struct{})
	if filter.FromAddress != nil {
		for _, addr := range *filter.FromAddress {
			fromAddrFilter[addr] = struct{}{}
		}
	}

	toAddrFilter := make(map[common.Address]struct{})
	if filter.ToAddress != nil {
		for _, addr := range *filter.ToAddress {
			toAddrFilter[addr] = struct{}{}
		}
	}

	totalTraces := uint64(0)
	if filter.After != nil {
		totalTraces += *filter.After
	}
	if filter.Count != nil {
		totalTraces += *filter.Count
	} else {
		totalTraces = math.MaxUint64
	}
	traces := make([]TraceFrame, 0)
blockLoop:
	for blockNum := start; blockNum <= end; blockNum++ {
		txTraces, blockInfo, _, err := t.block(ctx, rpc.BlockNumberOrHashWithNumber(blockNum), false)
		if err != nil {
			return nil, err
		}

		for _, txTrace := range txTraces {
			chainContext := newChainContext(txTrace.res, blockInfo)
			for _, frame := range txTrace.frames {
				if len(fromAddrFilter) != 0 {
					if _, found := fromAddrFilter[frame.Action.From]; !found {
						continue
					}
				}
				if len(toAddrFilter) != 0 {
					if frame.Action.To == nil {
						continue
					}
					if _, found := toAddrFilter[*frame.Action.To]; !found {
						continue
					}
				}

				addChainContext(&frame, chainContext)
				traces = append(traces, frame)
				if uint64(len(traces)) >= totalTraces {
					break blockLoop
				}
			}
		}
	}

	if filter.After != nil {
		traces = traces[*filter.After:]
	}
	if filter.Count != nil {
		return traces[:*filter.Count], nil
	}
	return traces, nil
}
