/*
 * Copyright 2019, Offchain Labs, Inc.
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

package coordinator

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethvalidator"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-validator) -I. --go_out=paths=source_relative:. *.proto"

// Server provides an interface for interacting with a a running coordinator
type Server struct {
	coordinator *ethvalidator.ValidatorCoordinator

	tracker *txTracker
}

// NewServer returns a new instance of the Server class
func NewServer(
	machine machine.Machine,
	key *ecdsa.PrivateKey,
	validators []common.Address,
	connectionInfo ethbridge.ArbAddresses,
	ethURL string,
) *Server {
	// Commit all pending transactions in the simulator and print the names again
	escrowRequired := big.NewInt(10)
	config := valmessage.NewVMConfiguration(
		10,
		escrowRequired,
		common.Address{}, // Address 0 is eth
		validators,
		200000,
		common.Address{}, // Address 0 means no owner
	)

	man, err := ethvalidator.NewCoordinator(
		"Alice",
		machine.Clone(),
		key,
		config,
		false,
		math.MaxInt32, // maxCallSteps
		connectionInfo,
		ethURL,
		math.MaxInt32, // maxUnanSteps
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := man.Run(); err != nil {
		log.Fatalln(err)
	}

	receiptChan, errChan := man.Val.DepositFunds(context.Background(), escrowRequired)
	select {
	case receipt := <-receiptChan:
		if receipt.Status == 0 {
			log.Fatalln("Coordinator could not deposit funds")
		}
	case err := <-errChan:
		log.Fatal(err)
	}

	log.Println("Coordinator is trying to create the VM")

	receiptChan, errChan = man.CreateVM(time.Second * 60)

	select {
	case receipt := <-receiptChan:
		if receipt.Status == 0 {
			log.Fatalln("Coordinator failed to create VM")
		}
		log.Println("Coordinator created VM")
	case err := <-errChan:
		log.Fatalf("Failed to create vm: %v", err)
	}

	time.Sleep(500 * time.Millisecond)

	tracker := newTxTracker(man.Val.VMID, <-man.Val.VMCreatedTxHashChan)
	go func() {

		tracker.handleTxResults(man.Val.CompletedCallChan)
	}()

	return &Server{man, tracker}
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

// SendMessage takes a request from a client and sends it to the VM
func (m *Server) SendMessage(ctx context.Context, args *SendMessageArgs) (*SendMessageReply, error) {
	if !<-m.coordinator.Val.Bot.CanRun() {
		return nil, errors.New("Cannot send message when machine can't run")
	}
	sigBytes, err := hexutil.Decode(args.Signature)
	if err != nil {
		log.Printf("SendMessage: Failed to decode signature, %v\n", err)
		return nil, err
	}
	if len(sigBytes) != 65 {
		return nil, errors.New("SendMessage: Signature of wrong length")
	}
	// Convert sig with normalized v
	if sigBytes[64] == 27 {
		sigBytes[64] = 0
	} else if sigBytes[64] == 28 {
		sigBytes[64] = 1
	}

	dataBytes, err := hexutil.Decode(args.Data)
	if err != nil {
		return nil, err
	}
	rd := bytes.NewReader(dataBytes)
	dataVal, err := value.UnmarshalValue(rd)
	if err != nil {
		return nil, err
	}

	amount := big.NewInt(0)
	tokenType := [21]byte{}

	messageHash := solsha3.SoliditySHA3(
		solsha3.Bytes32(m.coordinator.Val.VMID),
		solsha3.Bytes32(dataVal.Hash()),
		solsha3.Uint256(amount),
		tokenType[:],
	)

	pubkey, err := hexutil.Decode(args.Pubkey)
	if err != nil {
		return nil, err
	}
	pub, err := crypto.UnmarshalPubkey(pubkey)
	if err != nil {
		return nil, err
	}

	go func() {
		signedMsg := solsha3.SoliditySHA3WithPrefix(solsha3.Bytes32(messageHash))
		if !crypto.VerifySignature(pubkey, signedMsg, sigBytes[:len(sigBytes)-1]) {
			return
		}

		senderArr := [32]byte{}
		copy(senderArr[12:], crypto.PubkeyToAddress(*pub).Bytes())

		m.coordinator.SendMessage(ethvalidator.OffchainMessage{
			Message: protocol.Message{
				Data:        dataVal,
				TokenType:   tokenType,
				Currency:    amount,
				Destination: senderArr,
			},
			Hash:      messageHash,
			Signature: sigBytes,
		})
	}()

	return &SendMessageReply{
		TxHash: hexutil.Encode(messageHash),
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
		ValidatorSigs: txInfo.ValidatorSigs,
		PartialHash:   txInfo.PartialHash,
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

// GetVMCreatedTxHash returns the txHash containing the CreateVM Event
func (m *Server) GetVMCreatedTxHash(ctx context.Context, args *GetVMCreatedTxHashArgs) (*GetVMCreatedTxHashReply, error) {
	res := <-m.tracker.VMCreatedTxHash()
	return &GetVMCreatedTxHashReply{
		VmCreatedTxHash: hexutil.Encode(res[:]),
	}, nil
}

// GetVMInfo returns current metadata about this VM
func (m *Server) GetVMInfo(ctx context.Context, args *GetVMInfoArgs) (*GetVMInfoReply, error) {
	return &GetVMInfoReply{
		VmID: hexutil.Encode(m.coordinator.Val.VMID[:]),
	}, nil
}

// CallMessage takes a request from a client to process in a temporary context and return the result
func (m *Server) CallMessage(ctx context.Context, args *CallMessageArgs) (*CallMessageReply, error) {
	if !<-m.coordinator.Val.Bot.CanRun() {
		return nil, errors.New("Cannot call when machine can't run")
	}
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
	copy(sender[12:], senderBytes)

	msg := protocol.NewSimpleMessage(dataVal, [21]byte{}, big.NewInt(0), sender)
	resultChan, errChan := m.coordinator.Val.Bot.RequestCall(msg)

	select {
	case logVal := <-resultChan:
		var buf bytes.Buffer
		_ = value.MarshalValue(logVal, &buf) // error can only occur from writes and bytes.Buffer is safe
		return &CallMessageReply{
			RawVal: hexutil.Encode(buf.Bytes()),
		}, nil
	case err := <-errChan:
		return nil, err
	}
}
