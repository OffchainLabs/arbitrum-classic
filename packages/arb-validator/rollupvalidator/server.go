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

package rollupvalidator

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/validatorserver"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupmanager"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
)

// Server provides an interface for interacting with a a running coordinator
type Server struct {
	rollupAddress common.Address
	tracker       *txTracker
	man           *rollupmanager.Manager
	maxCallTime   time.Duration
}

// NewServer returns a new instance of the Server class
func NewServer(man *rollupmanager.Manager, maxCallTime time.Duration) (*Server, error) {
	tracker, err := newTxTracker(man.GetCheckpointer().GetCheckpointDB(), man.GetCheckpointer().GetConfirmedNodeStore(), man.RollupAddress)
	if err != nil {
		return nil, err
	}
	man.AddListener(tracker)
	return &Server{man.RollupAddress, tracker, man, maxCallTime}, nil
}

// FindLogs takes a set of parameters and return the list of all logs that match the query
func (m *Server) FindLogs(ctx context.Context, args *validatorserver.FindLogsArgs) (*validatorserver.FindLogsReply, error) {
	var address *common.Address
	if len(args.Address) > 0 {
		addr := common.HexToAddress(args.Address)
		address = &addr
	}

	topics := make([]common.Hash, 0, len(args.Topics))
	for _, topic := range args.Topics {
		topicBytes, err := hexutil.Decode(topic)
		if err == nil {
			var topic common.Hash
			copy(topic[:], topicBytes)
			topics = append(topics, topic)
		}
	}

	fromHeight, err := strconv.ParseInt(args.FromHeight[2:], 16, 64)
	if err != nil {
		fmt.Println("FindLogs error, bad fromHeight", err)
		return nil, err
	}

	var logs []*validatorserver.LogInfo
	if args.ToHeight == "latest" {
		logs = m.tracker.FindLogs(&fromHeight, nil, address, topics)
	} else {
		toHeight, err := strconv.ParseInt(args.ToHeight[2:], 16, 64)
		if err != nil {
			fmt.Println("FindLogs error4", err)
			return nil, err
		}
		logs = m.tracker.FindLogs(&fromHeight, &toHeight, address, topics)
	}

	return &validatorserver.FindLogsReply{
		Logs: logs,
	}, nil
}

func (m *Server) GetOutputMessage(ctx context.Context, args *validatorserver.GetOutputMessageArgs) (*validatorserver.GetOutputMessageReply, error) {
	assertionHashBytes, err := hexutil.Decode(args.AssertionNodeHash)
	if err != nil {
		return nil, err
	}
	assertionHash := common.Hash{}
	copy(assertionHash[:], assertionHashBytes)
	msgIndex, err := strconv.ParseInt(args.MsgIndex, 16, 64)
	outputValue := m.tracker.OutputMsgVal(assertionHash, msgIndex)

	if outputValue == nil {
		return &validatorserver.GetOutputMessageReply{
			Found: false,
		}, nil
	} else {
		var buf bytes.Buffer
		_ = value.MarshalValue(outputValue, &buf)
		return &validatorserver.GetOutputMessageReply{
			Found:  true,
			RawVal: hexutil.Encode(buf.Bytes()),
		}, nil
	}
}

// GetMessageResult returns the value output by the VM in response to the message with the given hash
func (m *Server) GetMessageResult(ctx context.Context, args *validatorserver.GetMessageResultArgs) (*validatorserver.GetMessageResultReply, error) {
	txHashBytes, err := hexutil.Decode(args.TxHash)
	if err != nil {
		return nil, err
	}
	txHash := common.Hash{}
	copy(txHash[:], txHashBytes)
	txInfo := m.tracker.TxInfo(txHash)

	if !txInfo.Found {
		return &validatorserver.GetMessageResultReply{
			Found: false,
		}, nil
	}

	var buf bytes.Buffer
	_ = value.MarshalValue(txInfo.RawVal, &buf) // error can only occur from writes and bytes.Buffer is safe
	return &validatorserver.GetMessageResultReply{
		Found:         true,
		RawVal:        hexutil.Encode(buf.Bytes()),
		LogPreHash:    txInfo.LogsPreHash,
		LogPostHash:   txInfo.LogsPostHash,
		LogValHashes:  txInfo.LogsValHashes,
		OnChainTxHash: txInfo.OnChainTxHash,
	}, nil
}

// GetAssertionCount returns the total number of finalized assertions
func (m *Server) GetAssertionCount(ctx context.Context, args *validatorserver.GetAssertionCountArgs) (*validatorserver.GetAssertionCountReply, error) {
	req := m.tracker.AssertionCount()
	return &validatorserver.GetAssertionCountReply{
		AssertionCount: int32(req),
	}, nil
}

// GetVMInfo returns current metadata about this VM
func (m *Server) GetVMInfo(ctx context.Context, args *validatorserver.GetVMInfoArgs) (*validatorserver.GetVMInfoReply, error) {
	return &validatorserver.GetVMInfoReply{
		VmID: hexutil.Encode(m.rollupAddress[:]),
	}, nil
}

// CallMessage takes a request from a client to process in a temporary context and return the result
func (m *Server) CallMessage(ctx context.Context, args *validatorserver.CallMessageArgs) (*validatorserver.CallMessageReply, error) {
	log.Println("CallMessage", args.Data)
	dataBytes, err := hexutil.Decode(args.Data)
	if err != nil {
		return nil, err
	}

	contractAddressBytes, err := hexutil.Decode(args.ContractAddress)
	if err != nil {
		return nil, err
	}
	var contractAddress common.Address
	copy(contractAddress[:], contractAddressBytes)

	senderBytes, err := hexutil.Decode(args.Sender)
	if err != nil {
		return nil, err
	}
	var sender common.Address
	copy(sender[:], senderBytes)

	msg := message.Call{
		To:        contractAddress,
		From:      sender,
		Data:      dataBytes,
		BlockNum:  m.man.CurrentBlockId().Height,
		Timestamp: big.NewInt(time.Now().Unix()),
	}

	inbox := message.AddToPrev(value.NewEmptyTuple(), msg)
	assertion, steps := m.man.ExecuteCall(inbox, m.maxCallTime)

	log.Println("Executed call for", steps, "steps")

	results := assertion.Logs
	if len(results) == 0 {
		return nil, errors.New("call produced no output")
	}
	lastLogVal := results[len(results)-1]
	lastLog, err := evm.ProcessLog(lastLogVal, m.rollupAddress)
	if err != nil {
		return nil, err
	}
	logHash := lastLog.GetEthMsg().TxHash
	if logHash != msg.ReceiptHash() {
		// Last produced log is not the call we sent
		return nil, errors.New("call took too long to execute")
	}

	result := results[len(results)-1]
	var buf bytes.Buffer
	_ = value.MarshalValue(result, &buf) // error can only occur from writes and bytes.Buffer is safe
	return &validatorserver.CallMessageReply{
		RawVal: hexutil.Encode(buf.Bytes()),
	}, nil
}
