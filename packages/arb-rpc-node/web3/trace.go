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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	arbcommon "github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type TraceAction struct {
	CallType string         `json:"callType,omitempty"`
	From     common.Address `json:"from"`
	Gas      hexutil.Uint64 `json:"gas"`
	Input    hexutil.Bytes  `json:"input,omitempty"`
	Init     hexutil.Bytes  `json:"init,omitempty"`
	To       hexutil.Bytes  `json:"to"`
	Value    *hexutil.Big   `json:"value"`
}

type TraceCallResult struct {
	Address *common.Address `json:"address,omitempty"`
	Code    hexutil.Bytes   `json:"code,omitempty"`
	GasUsed hexutil.Uint64  `json:"gasUsed"`
	Output  hexutil.Bytes   `json:"output,omitempty"`
}

type TraceFrame struct {
	Action       TraceAction      `json:"action"`
	Result       *TraceCallResult `json:"result,omitempty"`
	Error        *string          `json:"error,omitempty"`
	Subtraces    int              `json:"subtraces"`
	TraceAddress []int            `json:"traceAddress"`
	Type         string           `json:"type"`
}

type TraceResult struct {
	Output    hexutil.Bytes
	StateDiff *int         `json:"stateDiff"`
	Trace     []TraceFrame `json:"trace"`
	VmTrace   *int         `json:"vmTrace"`
}

type Trace struct {
	s *Server
}

func NewTracer(s *Server) *Trace {
	return &Trace{s: s}
}

func (t *Trace) Call(callArgs CallTxArgs, traceTypes []string, blockNum rpc.BlockNumberOrHash) (*TraceResult, error) {
	for _, typ := range traceTypes {
		if typ != "trace" {
			return nil, errors.Errorf("unsupported trace type: %v", typ)
		}
	}
	snap, err := t.s.getSnapshotForNumberOrHash(blockNum)
	if err != nil {
		return nil, err
	}
	from, msg := buildCallMsg(callArgs, t.s.maxCallGas)

	callRes, debugPrints, err := snap.Call(msg, from)
	if err != nil {
		return nil, err
	}
	if callRes.ResultCode != evm.ReturnCode {
		return nil, evm.HandleCallError(callRes, t.s.ganacheMode)
	}
	receipt := callRes.ToEthReceipt(arbcommon.Hash{})
	trace, err := evm.GetTrace(debugPrints)
	if err != nil {
		return nil, err
	}
	frame, err := trace.FrameTree()
	if err != nil {
		return nil, err
	}

	type trackedFrame struct {
		f            evm.Frame
		traceAddress []int
	}

	frames := []trackedFrame{{f: frame, traceAddress: make([]int, 0)}}
	res := &TraceResult{
		Output: callRes.ReturnData,
	}
	for len(frames) > 0 {
		frame := frames[0]
		frames = frames[1:]
		var emptyAddress common.Address
		var createdContractAddress *common.Address
		switch frame := frame.f.(type) {
		case *evm.CallFrame:
			if len(res.Trace) == 0 && receipt.ContractAddress != emptyAddress {
				createdContractAddress = &receipt.ContractAddress
			}
		case *evm.CreateFrame:
			tmp := frame.Create.ContractAddress.ToEthAddress()
			createdContractAddress = &tmp
		case *evm.Create2Frame:
			tmp := frame.Create.ContractAddress.ToEthAddress()
			createdContractAddress = &tmp
		}

		callFrame := frame.f.GetCallFrame()
		action := TraceAction{
			From:  callFrame.Call.From.ToEthAddress(),
			Gas:   hexutil.Uint64(callFrame.Call.Gas.Uint64()),
			Value: (*hexutil.Big)(callFrame.Call.Value),
		}

		if createdContractAddress != nil {
			action.Init = callFrame.Call.Data
		} else {
			if callFrame.Call.To == nil {
				return nil, errors.New("expected call to have destination")
			}
			action.Input = callFrame.Call.Data
			action.To = callFrame.Call.To.Bytes()
			action.CallType = callFrame.Call.Type.RPCString()
		}

		var result *TraceCallResult
		var callErr *string
		if callFrame.Return.Result == evm.ReturnCode {
			result = &TraceCallResult{
				GasUsed: hexutil.Uint64(callFrame.Return.GasUsed.Uint64()),
			}
			if createdContractAddress != nil {
				result.Address = createdContractAddress
				result.Code = callFrame.Return.ReturnData
			} else {
				result.Output = callFrame.Return.ReturnData
			}
		} else {
			tmp := callFrame.Return.Result.String()
			callErr = &tmp
		}
		frameType := "call"
		if createdContractAddress != nil {
			frameType = "create"
		}
		res.Trace = append(res.Trace, TraceFrame{
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
	return res, nil
}
