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
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
	"math/big"
	"strconv"
	"time"

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
	tracker, err := newTxTracker(checkpointer.GetCheckpointDB(), checkpointer.GetConfirmedNodeStore())
	if err != nil {
		return nil, err
	}
	man.AddListener(tracker)
	return &Server{man.RollupAddress, tracker, man, maxCallTime}, nil
}

// FindLogs takes a set of parameters and return the list of all logs that match the query
func (m *Server) FindLogs(ctx context.Context, args *evm.FindLogsArgs) (*evm.FindLogsReply, error) {
	addresses := make([]common.Address, 0, 1)
	for _, addr := range args.Addresses {
		addresses = append(addresses, common.HexToAddress(addr))
	}

	topicGroups := make([][]common.Hash, 0, len(args.TopicGroups))
	for _, topicGroup := range args.TopicGroups {
		topics := make([]common.Hash, 0, len(topicGroup.Topics))
		for _, topic := range topicGroup.Topics {
			topics = append(topics, common.NewHashFromEth(ethcommon.HexToHash(topic)))
		}
		topicGroups = append(topicGroups, topics)
	}

	var fromHeight *uint64
	if len(args.FromHeight) != 0 {
		intVal, err := strconv.ParseUint(args.FromHeight[2:], 16, 64)
		if err != nil {
			fmt.Println("FindLogs error, bad fromHeight", err)
			return nil, err
		}
		fromHeight = &intVal
	}

	var toHeight *uint64
	if len(args.ToHeight) != 0 && args.ToHeight != "latest" {
		intVal, err := strconv.ParseUint(args.ToHeight[2:], 16, 64)
		if err != nil {
			fmt.Println("FindLogs error4", err)
			return nil, err
		}
		toHeight = &intVal
	}

	logs, err := m.tracker.FindLogs(ctx, fromHeight, toHeight, addresses, topicGroups)
	if err != nil {
		return nil, err
	}

	logInfos := make([]*evm.FullLogBuf, 0, len(logs))
	for _, l := range logs {
		logInfos = append(logInfos, l.Marshal())
	}

	return &evm.FindLogsReply{
		Logs: logInfos,
	}, nil
}

func (m *Server) GetOutputMessage(ctx context.Context, args *evm.GetOutputMessageArgs) (*evm.GetOutputMessageReply, error) {
	assertionHashBytes, err := hexutil.Decode(args.AssertionNodeHash)
	if err != nil {
		return nil, err
	}
	assertionHash := common.Hash{}
	copy(assertionHash[:], assertionHashBytes)
	msgIndex, err := strconv.ParseInt(args.MsgIndex, 16, 64)
	outputValue, err := m.tracker.OutputMsgVal(ctx, assertionHash, msgIndex)

	if outputValue == nil || err != nil {
		return &evm.GetOutputMessageReply{
			Found: false,
		}, err
	} else {
		var buf bytes.Buffer
		_ = value.MarshalValue(outputValue, &buf)
		return &evm.GetOutputMessageReply{
			Found:  true,
			RawVal: hexutil.Encode(buf.Bytes()),
		}, nil
	}
}

// GetMessageResult returns the value output by the VM in response to the message with the given hash
func (m *Server) GetMessageResult(ctx context.Context, args *evm.GetMessageResultArgs) (*evm.GetMessageResultReply, error) {
	txHash := common.NewHashFromEth(ethcommon.HexToHash(args.TxHash))
	txInfo, err := m.tracker.TxInfo(ctx, txHash)
	if err != nil {
		return nil, err
	}

	return &evm.GetMessageResultReply{
		Tx: txInfo.Marshal(),
	}, nil
}

// GetAssertionCount returns the total number of finalized assertions
func (m *Server) GetAssertionCount(ctx context.Context, _ *evm.GetAssertionCountArgs) (*evm.GetAssertionCountReply, error) {
	req, err := m.tracker.AssertionCount(ctx)
	if err != nil {
		return nil, err
	}
	return &evm.GetAssertionCountReply{
		AssertionCount: int32(req),
	}, nil
}

// GetVMInfo returns current metadata about this VM
func (m *Server) GetVMInfo(_ context.Context, _ *evm.GetVMInfoArgs) (*evm.GetVMInfoReply, error) {
	return &evm.GetVMInfoReply{
		VmID: hexutil.Encode(m.rollupAddress[:]),
	}, nil
}

func (m *Server) executeCall(mach machine.Machine, args *evm.CallMessageArgs) (*evm.CallMessageReply, error) {
	dataBytes, err := hexutil.Decode(args.Data)
	if err != nil {
		return nil, err
	}

	var sender common.Address
	if len(args.Sender) > 0 {
		senderBytes, err := hexutil.Decode(args.Sender)
		if err != nil {
			return nil, err
		}
		copy(sender[:], senderBytes)
	}

	latestBlock, err := m.man.CurrentBlockId()
	if err != nil {
		return nil, err
	}

	seq, _ := new(big.Int).SetString("999999999999999999999999", 10)

	callMsg := message.NewCallFromData(dataBytes)
	// Max gas 0 in a call is a request for infinite gas
	if callMsg.MaxGas.Cmp(big.NewInt(0)) == 0 {
		callMsg.MaxGas.SetUint64(1000000000000)
	}

	log.Println("message.L2Message------------------------x-----")
	log.Println("sender: ", sender)
	log.Println(message.L2Message{Msg: callMsg})
	inboxMsg := message.NewInboxMessage(
		message.L2Message{Msg: callMsg},
		sender,
		seq,
		message.ChainTime{
			BlockNum:  latestBlock.Height,
			Timestamp: big.NewInt(time.Now().Unix()),
		},
	)

	inbox := structures.NewVMInbox()
	inbox.DeliverMessage(inboxMsg)
	assertion, steps := mach.ExecuteAssertion(
		// Call execution is only limited by wall time, so use a massive max steps as an approximation to infinity
		10000000000000000,
		inbox.AsValue(),
		m.maxCallTime,
	)

	// If the machine wasn't able to run and it reports that it is currently
	// blocked, return the block reason to give the client more information
	// as opposed to just returning a general "call produced no output"
	if br := mach.IsBlocked(true); steps == 0 && br != nil {
		log.Println("can't produce solution since machine is blocked", br)
		return nil, fmt.Errorf("can't produce solution since machine is blocked %v", br)
	}

	log.Println("Executed call for", steps, "steps")

	results := assertion.ParseLogs()
	if len(results) == 0 {
		return nil, errors.New("call produced no output")
	}
	lastLogVal := results[len(results)-1]
	lastLog, err := evm.NewResultFromValue(lastLogVal)
	if err != nil {
		return nil, err
	}
	if lastLog.L1Message.MessageID() != inboxMsg.MessageID() {
		// Last produced log is not the call we sent
		return nil, errors.New("call took too long to execute")
	}

	result := results[len(results)-1]
	var buf bytes.Buffer
	_ = value.MarshalValue(result, &buf) // error can only occur from writes and bytes.Buffer is safe
	return &evm.CallMessageReply{
		RawVal: hexutil.Encode(buf.Bytes()),
	}, nil
}

// CallMessage takes a request from a client to process in a temporary context and return the result
func (m *Server) CallMessage(ctx context.Context, args *evm.CallMessageArgs) (*evm.CallMessageReply, error) {
	mach, err := m.man.GetLatestMachine()
	if err != nil {
		return nil, err
	}
	return m.executeCall(mach, args)
}

// PendingCall takes a request from a client to process in a temporary context and return the result
func (m *Server) PendingCall(ctx context.Context, args *evm.CallMessageArgs) (*evm.CallMessageReply, error) {
	mach, err := m.man.GetPendingMachine()
	if err != nil {
		return nil, err
	}
	return m.executeCall(mach, args)
}

func (m *Server) GetLatestNodeLocation(ctx context.Context, args *evm.GetLatestNodeLocationArgs,
) (*evm.GetLatestNodeLocationReply, error) {
	loc, err := m.tracker.GetLatestNodeLocation(ctx)
	if err != nil {
		return nil, err
	}
	return &evm.GetLatestNodeLocationReply{Location: loc}, nil
}

func (m *Server) GetLatestPendingNodeLocation(ctx context.Context, args *evm.GetLatestNodeLocationArgs,
) (*evm.GetLatestNodeLocationReply, error) {
	loc, err := m.tracker.GetLatestPendingNodeLocation(ctx)
	if err != nil {
		return nil, err
	}
	return &evm.GetLatestNodeLocationReply{Location: loc}, nil
}
