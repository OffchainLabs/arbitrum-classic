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
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-validator) -I. --go_out=paths=source_relative:. *.proto"

// Server provides an interface for interacting with a a running coordinator
type Server struct {
	rollupAddress common.Address
	tracker       *txTracker
}

// NewServer returns a new instance of the Server class
func NewServer(
	auth *bind.TransactOpts,
	client arbbridge.ArbClient,
	rollupAddress common.Address,
	codeFile string,
	config structures.ChainParams,
) (*Server, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	checkpointer := structures.NewRollupCheckpointer(ctx, rollupAddress, codeFile, 100)

	chainObserver, err := rollup.NewChain(ctx, rollupAddress, checkpointer, config, true, protocol.NewTimeBlocks(header.Number))
	if err != nil {
		return nil, err
	}

	err = rollup.RunObserver(ctx, chainObserver, client)
	if err != nil {
		return nil, err
	}

	validatorListener := rollup.NewValidatorChainListener(chainObserver)
	err = validatorListener.AddStaker(client, auth)
	if err != nil {
		return nil, err
	}
	completedAssertionChan := make(chan rollup.FinalizedAssertion)
	assertionListener := &rollup.AssertionListener{completedAssertionChan}
	chainObserver.AddListener(&rollup.AnnouncerListener{})
	chainObserver.AddListener(validatorListener)
	chainObserver.AddListener(assertionListener)

	tracker := newTxTracker(rollupAddress)
	go func() {
		tracker.handleTxResults(assertionListener.CompletedAssertionChan)
	}()

	return &Server{rollupAddress, tracker}, nil
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
	//	if !<-m.coordinator.ChannelVal.CanRun() {
	//		return nil, errors.New("Cannot call when machine can't run")
	//	}
	//	dataBytes, err := hexutil.Decode(args.Data)
	//	if err != nil {
	//		return nil, err
	//	}
	//	rd := bytes.NewReader(dataBytes)
	//	dataVal, err := value.UnmarshalValue(rd)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	senderBytes, err := hexutil.Decode(args.Sender)
	//	if err != nil {
	//		return nil, err
	//	}
	//	var sender common.Address
	//	copy(sender[:], senderBytes)
	//
	//	msg := protocol.NewSimpleMessage(dataVal, [21]byte{}, big.NewInt(0), sender)
	//	resultChan, errChan := m.coordinator.ChannelVal.RequestCall(msg)
	//
	//	select {
	//	case logVal := <-resultChan:
	//		var buf bytes.Buffer
	//		_ = value.MarshalValue(logVal, &buf) // error can only occur from writes and bytes.Buffer is safe
	//		return &CallMessageReply{
	//			RawVal: hexutil.Encode(buf.Bytes()),
	//		}, nil
	//	case err := <-errChan:
	//		return nil, err
	//	}
	return nil, nil
}
