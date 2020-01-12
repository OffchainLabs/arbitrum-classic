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

	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-validator) -I. --go_out=paths=source_relative:. *.proto"

// Server provides an interface for interacting with a a running coordinator
type Server struct {
	rollupAddress common.Address
	tracker       *txTracker
	chain         *rollup.ChainObserver
	maxCallSteps  uint32
}

// NewServer returns a new instance of the Server class
func NewServer(chainObserver *rollup.ChainObserver, maxCallSteps uint32) (*Server, error) {
	completedAssertionChan := make(chan rollup.FinalizedAssertion)
	assertionListener := &rollup.AssertionListener{completedAssertionChan}
	chainObserver.AddListener(assertionListener)

	rollupAddress := chainObserver.ContractAddress()
	tracker := newTxTracker(rollupAddress)
	go func() {
		tracker.handleTxResults(assertionListener.CompletedAssertionChan)
	}()

	return &Server{rollupAddress, tracker, chainObserver, maxCallSteps}, nil
}

// FindLogs takes a set of parameters and return the list of all logs that match the query
func (m *Server) FindLogs(ctx context.Context, args *FindLogsArgs) (*FindLogsReply, error) {
	addressBytes, err := hexutil.Decode(args.Address)
	if err != nil {
		fmt.Println("FindLogs error1", err)
		return nil, err
	}
	addressInt := new(big.Int).SetBytes(addressBytes[:])

	topics := make([][32]byte, 0, len(args.Topics))
	for _, topic := range args.Topics {
		topicBytes, err := hexutil.Decode(topic)
		if err == nil {
			var topic [32]byte
			copy(topic[:], topicBytes)
			topics = append(topics, topic)
		}
	}

	fromHeight, err := strconv.ParseInt(args.FromHeight[2:], 16, 64)
	if err != nil {
		fmt.Println("FindLogs error, bad fromHeight", err)
		return nil, err
	}

	var logsChan <-chan []*LogInfo
	if args.ToHeight == "latest" {
		logsChan = m.tracker.FindLogs(&fromHeight, nil, addressInt, topics)
	} else {
		toHeight, err := strconv.ParseInt(args.ToHeight[2:], 16, 64)
		if err != nil {
			fmt.Println("FindLogs error4", err)
			return nil, err
		}
		logsChan = m.tracker.FindLogs(&fromHeight, &toHeight, addressInt, topics)
	}

	ret := <-logsChan
	return &FindLogsReply{
		Logs: ret,
	}, nil
}

// GetMessageResult returns the value output by the VM in response to the message with the given hash
func (m *Server) GetMessageResult(ctx context.Context, args *GetMessageResultArgs) (*GetMessageResultReply, error) {
	txHashBytes, err := hexutil.Decode(args.TxHash)
	if err != nil {
		return nil, err
	}
	txHash := [32]byte{}
	copy(txHash[:], txHashBytes)
	resultChan := m.tracker.TxInfo(txHash)

	txInfo := <-resultChan
	if !txInfo.Found {
		return &GetMessageResultReply{
			Found: false,
		}, nil
	}

	var buf bytes.Buffer
	_ = value.MarshalValue(txInfo.RawVal, &buf) // error can only occur from writes and bytes.Buffer is safe
	return &GetMessageResultReply{
		Found:         true,
		RawVal:        hexutil.Encode(buf.Bytes()),
		LogPreHash:    txInfo.LogsPreHash,
		LogPostHash:   txInfo.LogsPostHash,
		LogValHashes:  txInfo.LogsValHashes,
		OnChainTxHash: txInfo.OnChainTxHash,
	}, nil
}

// GetAssertionCount returns the total number of finalized assertions
func (m *Server) GetAssertionCount(ctx context.Context, args *GetAssertionCountArgs) (*GetAssertionCountReply, error) {
	req := m.tracker.AssertionCount()
	return &GetAssertionCountReply{
		AssertionCount: int32(<-req),
	}, nil
}

// GetVMInfo returns current metadata about this VM
func (m *Server) GetVMInfo(ctx context.Context, args *GetVMInfoArgs) (*GetVMInfoReply, error) {
	return &GetVMInfoReply{
		VmID: hexutil.Encode(m.rollupAddress[:]),
	}, nil
}

// CallMessage takes a request from a client to process in a temporary context and return the result
func (m *Server) CallMessage(ctx context.Context, args *CallMessageArgs) (*CallMessageReply, error) {
	dataBytes, err := hexutil.Decode(args.Data)
	if err != nil {
		return nil, err
	}
	rd := bytes.NewReader(dataBytes)
	dataVal, err := value.UnmarshalValue(rd)
	if err != nil {
		return nil, err
	}

	senderBytes, err := hexutil.Decode(args.Sender)
	if err != nil {
		return nil, err
	}
	var sender common.Address
	copy(sender[:], senderBytes)

	msg := valprotocol.NewSimpleMessage(dataVal, [21]byte{}, big.NewInt(0), sender)
	messageHash := solsha3.SoliditySHA3(
		solsha3.Bytes32(msg.Destination),
		solsha3.Bytes32(msg.Data.Hash()),
		solsha3.Uint256(msg.Currency),
		msg.TokenType[:],
	)
	msgHashInt := new(big.Int).SetBytes(messageHash[:])
	val, _ := value.NewTupleFromSlice([]value.Value{
		msg.Data,
		value.NewIntValue(m.chain.CurrentTime().AsInt()),
		value.NewIntValue(msgHashInt),
	})
	callingMessage := valprotocol.Message{
		Data:        val.Clone(),
		TokenType:   msg.TokenType,
		Currency:    msg.Currency,
		Destination: msg.Destination,
	}
	messageStack := protocol.NewMessageStack()
	messageStack.AddMessage(callingMessage.AsValue())

	assertion, steps := m.chain.ExecuteCall(messageStack.GetValue(), m.maxCallSteps)

	log.Println("Executed call for", steps, "steps")

	results := assertion.Logs
	if len(results) == 0 {
		return nil, errors.New("call produced no output")
	}
	lastLogVal := results[len(results)-1]
	lastLog, err := evm.ProcessLog(lastLogVal)
	if err != nil {
		return nil, err
	}
	logHash := lastLog.GetEthMsg().Data.TxHash
	if !bytes.Equal(logHash[:], messageHash) {
		// Last produced log is not the call we sent
		return nil, errors.New("call took too long to execute")
	}

	result := results[len(results)-1]
	var buf bytes.Buffer
	_ = value.MarshalValue(result, &buf) // error can only occur from writes and bytes.Buffer is safe
	return &CallMessageReply{
		RawVal: hexutil.Encode(buf.Bytes()),
	}, nil
}
