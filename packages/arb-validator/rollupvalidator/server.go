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
	ethcommon "github.com/ethereum/go-ethereum/common"
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
	checkpointer := man.GetCheckpointer()
	tracker, err := newTxTracker(checkpointer.GetCheckpointDB(), checkpointer.GetConfirmedNodeStore(), man.RollupAddress)
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

	var logs []evm.FullLog
	if args.ToHeight == "latest" {
		logs, err = m.tracker.FindLogs(ctx, &fromHeight, nil, address, topics)
	} else {
		toHeight, err := strconv.ParseInt(args.ToHeight[2:], 16, 64)
		if err != nil {
			fmt.Println("FindLogs error4", err)
			return nil, err
		}
		logs, err = m.tracker.FindLogs(ctx, &fromHeight, &toHeight, address, topics)
	}
	if err != nil {
		return nil, err
	}

	logInfos := make([]*evm.FullLogBuf, 0, len(logs))
	for _, l := range logs {
		logInfos = append(logInfos, l.Marshal())
	}

	return &validatorserver.FindLogsReply{
		Logs: logInfos,
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
	outputValue, err := m.tracker.OutputMsgVal(ctx, assertionHash, msgIndex)

	if outputValue == nil || err != nil {
		return &validatorserver.GetOutputMessageReply{
			Found: false,
		}, err
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
	txHash := common.NewHashFromEth(ethcommon.HexToHash(args.TxHash))
	txInfo, err := m.tracker.TxInfo(ctx, txHash)
	if err != nil {
		return nil, err
	}

	return &validatorserver.GetMessageResultReply{
		Tx: txInfo.Marshal(),
	}, nil
}

// GetAssertionCount returns the total number of finalized assertions
func (m *Server) GetAssertionCount(ctx context.Context, _ *validatorserver.GetAssertionCountArgs) (*validatorserver.GetAssertionCountReply, error) {
	req, err := m.tracker.AssertionCount(ctx)
	if err != nil {
		return nil, err
	}
	return &validatorserver.GetAssertionCountReply{
		AssertionCount: int32(req),
	}, nil
}

// GetVMInfo returns current metadata about this VM
func (m *Server) GetVMInfo(_ context.Context, _ *validatorserver.GetVMInfoArgs) (*validatorserver.GetVMInfoReply, error) {
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

	callMsg := message.Call{
		To:   contractAddress,
		From: sender,
		Data: dataBytes,
	}

	deliveredMsg := message.Delivered{
		Message: callMsg,
		DeliveryInfo: message.DeliveryInfo{
			ChainTime: message.ChainTime{
				BlockNum:  m.man.CurrentBlockId().Height,
				Timestamp: big.NewInt(time.Now().Unix()),
			},
			TxId: big.NewInt(0),
		},
	}

	inbox := message.AddToPrev(value.NewEmptyTuple(), deliveredMsg)
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
	logHash := lastLog.GetEthMsg().TxHash()
	if logHash != deliveredMsg.ReceiptHash() {
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
