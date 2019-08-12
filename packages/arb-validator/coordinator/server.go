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
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethvalidator"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

// Server provides an interface for interacting with a a running coordinator
type Server struct {
	coordinator *ethvalidator.ValidatorCoordinator

	requests chan validatorRequest
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

	tx, err := man.Val.DepositFunds(escrowRequired)
	if err != nil {
		log.Fatal(err, tx)
	}

	if err := man.Run(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Coordinator is trying to create the VM")

	retChan, errChan := man.CreateVM(time.Second * 60)

	select {
	case <-retChan:
		log.Println("Coordinator created VM")
	case err := <-errChan:
		log.Fatalf("Failed to create vm: %v", err)
	}

	time.Sleep(500 * time.Millisecond)
	requests := make(chan validatorRequest, 100)

	go func() {
		tracker := newTxTracker(man.Val.VMID, <-man.Val.VMCreatedTxHashChan)
		tracker.handleTxResults(man.Val.CompletedCallChan, requests)
	}()

	return &Server{man, requests}
}

func (m *Server) requestAssertionCount() <-chan int {
	req := make(chan int, 1)
	m.requests <- assertionCountRequest{req}
	return req
}

func (m *Server) requestVMCreatedTxHashChan() <-chan [32]byte {
	req := make(chan [32]byte, 1)
	m.requests <- vmCreatedTxHashRequest{req}
	return req
}

func (m *Server) requestTxInfo(txHash [32]byte) <-chan txInfo {
	req := make(chan txInfo, 1)
	m.requests <- txRequest{txHash, req}
	return req
}

func (m *Server) requestFindLogs(
	fromHeight *int64,
	toHeight *int64,
	address *big.Int,
	topics [][32]byte,
) <-chan []LogInfo {
	req := make(chan []LogInfo, 1)
	m.requests <- findLogsRequest{fromHeight, toHeight, address, topics, req}
	return req
}

// FindLogsArgs contains input data for FindLogs
type FindLogsArgs struct {
	FromHeight string   `json:"fromHeight"`
	ToHeight   string   `json:"toHeight"`
	Address    string   `json:"address"`
	Topics     []string `json:"topics"`
}

// FindLogsReply contains output data for FindLogs
type FindLogsReply struct {
	Logs []LogInfo `json:"logs"`
}

// FindLogs takes a set of parameters and return the list of all logs that match the query
func (m *Server) FindLogs(r *http.Request, args *FindLogsArgs, reply *FindLogsReply) error {
	addressBytes, err := hexutil.Decode(args.Address)
	if err != nil {
		fmt.Println("FindLogs error1", err)
		return err
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
		return err
	}

	var logsChan <-chan []LogInfo
	if args.ToHeight == "latest" {
		logsChan = m.requestFindLogs(&fromHeight, nil, addressInt, topics)
	} else {
		toHeight, err := strconv.ParseInt(args.ToHeight[2:], 16, 64)
		if err != nil {
			fmt.Println("FindLogs error4", err)
			return err
		}
		logsChan = m.requestFindLogs(&fromHeight, &toHeight, addressInt, topics)
	}

	ret := <-logsChan
	reply.Logs = ret
	return nil
}

// SendMessageArgs contains input data for SendMessage
type SendMessageArgs struct {
	Data      string `json:"data"`
	Pubkey    string `json:"pubkey"`
	Signature string `json:"signature"`
}

// SendMessageReply contains output data for SendMessage
type SendMessageReply struct {
	TxHash string `json:"hash"`
}

// SendMessage takes a request from a client and sends it to the VM
func (m *Server) SendMessage(r *http.Request, args *SendMessageArgs, reply *SendMessageReply) error {
	sigBytes, err := hexutil.Decode(args.Signature)
	if err != nil {
		log.Printf("SendMessage: Failed to decode signature, %v\n", err)
		return err
	}
	if len(sigBytes) != 65 {
		return errors.New("SendMessage: Signature of wrong length")
	}
	// Convert sig with normalized v
	if sigBytes[64] == 27 {
		sigBytes[64] = 0
	} else if sigBytes[64] == 28 {
		sigBytes[64] = 1
	}

	dataBytes, err := hexutil.Decode(args.Data)
	if err != nil {
		return err
	}
	rd := bytes.NewReader(dataBytes)
	dataVal, err := value.UnmarshalValue(rd)
	if err != nil {
		return err
	}

	amount := big.NewInt(0)
	tokenType := [21]byte{}

	messageHash := solsha3.SoliditySHA3(
		solsha3.Bytes32(m.coordinator.Val.VMID),
		solsha3.Bytes32(dataVal.Hash()),
		solsha3.Uint256(amount),
		tokenType[:],
	)
	reply.TxHash = hexutil.Encode(messageHash)

	pubkey, err := hexutil.Decode(args.Pubkey)
	if err != nil {
		return err
	}
	pub, err := crypto.UnmarshalPubkey(pubkey)
	if err != nil {
		return err
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

	return nil
}

// GetMessageResultArgs contains input data for GetMessageResult
type GetMessageResultArgs struct {
	TxHash string `json:"txHash"`
}

// GetMessageResultReply contains output data for GetMessageResult
type GetMessageResultReply struct {
	Found         bool     `json:"found"`
	RawVal        string   `json:"rawVal"`
	LogPreHash    string   `json:"logPreHash"`    // Acc hash before RawVal
	LogPostHash   string   `json:"logPostHash"`   // Acc hash to prove
	LogValHashes  []string `json:"logValHashes"`  // Intermediate value hashes
	ValidatorSigs []string `json:"validatorSigs"` // Unanimous signatures
	PartialHash   string   `json:"partialHash"`   // Unanimous partial hash
	OnChainTxHash string   `json:"onChainTxHash"` // Disputable Tx hash
}

// GetMessageResult returns the value output by the VM in response to the message with the given hash
func (m *Server) GetMessageResult(r *http.Request, args *GetMessageResultArgs, reply *GetMessageResultReply) error {
	txHashBytes, err := hexutil.Decode(args.TxHash)
	if err != nil {
		return err
	}
	txHash := [32]byte{}
	copy(txHash[:], txHashBytes)
	resultChan := m.requestTxInfo(txHash)

	txInfo := <-resultChan
	reply.Found = txInfo.Found
	if txInfo.Found {
		var buf bytes.Buffer
		_ = value.MarshalValue(txInfo.RawVal, &buf) // error can only occur from writes and bytes.Buffer is safe
		reply.RawVal = hexutil.Encode(buf.Bytes())

		// Log Proof pieces
		reply.LogPreHash = txInfo.LogsPreHash
		reply.LogPostHash = txInfo.LogsPostHash
		reply.LogValHashes = txInfo.LogsValHashes

		// Unanimous or Disputable assertion proof info
		reply.ValidatorSigs = txInfo.ValidatorSigs
		reply.PartialHash = txInfo.PartialHash
		reply.OnChainTxHash = txInfo.OnChainTxHash
	}
	return nil
}

// GetAssertionCountReply contains output data for GetAssertionCount
type GetAssertionCountReply struct {
	AssertionCount int `json:"assertionCount"`
}

// GetAssertionCount returns the total number of finalized assertions
func (m *Server) GetAssertionCount(r *http.Request, _ *struct{}, reply *GetAssertionCountReply) error {
	req := m.requestAssertionCount()
	reply.AssertionCount = <-req
	return nil
}

// GetVMCreatedTxHashReply contains output data for GetVMCreatedTxHash
type GetVMCreatedTxHashReply struct {
	VMCreatedTxHash string `json:"vmCreatedTxHash"`
}

// GetVMCreatedTxHash returns the txHash containing the CreateVM Event
func (m *Server) GetVMCreatedTxHash(
	r *http.Request,
	_ *struct{},
	reply *GetVMCreatedTxHashReply,
) error {
	res := <-m.requestVMCreatedTxHashChan()
	reply.VMCreatedTxHash = hexutil.Encode(res[:])
	return nil
}

// GetVMInfoReply contains output data for GetVMInfo
type GetVMInfoReply struct {
	VMId string `json:"vmID"`
}

// GetVMInfo returns current metadata about this VM
func (m *Server) GetVMInfo(r *http.Request, _ *struct{}, reply *GetVMInfoReply) error {
	reply.VMId = hexutil.Encode(m.coordinator.Val.VMID[:])
	return nil
}

// CallMessageArgs contains input data for CallMessage
type CallMessageArgs struct {
	Data   string `json:"data"`
	Sender string `json:"sender"`
}

// CallMessageReply contains output data for CallMessage
type CallMessageReply struct {
	ReturnVal string
	Success   bool
}

// CallMessage takes a request from a client to process in a temporary context and return the result
func (m *Server) CallMessage(r *http.Request, args *CallMessageArgs, reply *CallMessageReply) error {
	dataBytes, err := hexutil.Decode(args.Data)
	if err != nil {
		return err
	}
	rd := bytes.NewReader(dataBytes)
	dataVal, err := value.UnmarshalValue(rd)
	if err != nil {
		return err
	}

	senderBytes, err := hexutil.Decode(args.Sender)
	if err != nil {
		return err
	}
	var sender common.Address
	copy(sender[:], senderBytes)

	msg := protocol.NewSimpleMessage(dataVal, [21]byte{}, big.NewInt(0), sender)
	resultChan, errChan := m.coordinator.Val.Bot.RequestCall(msg)

	select {
	case logVal := <-resultChan:
		result, err := evm.ProcessLog(logVal)
		if err != nil {
			log.Printf("Error %v while responding to message %v\n", err, msg)
		}
		switch result := result.(type) {
		case evm.Stop:
			reply.Success = true
		case evm.Return:
			reply.ReturnVal = hexutil.Encode(result.ReturnVal)
			reply.Success = true
		case evm.Revert:
			reply.ReturnVal = hexutil.Encode(result.ReturnVal)
			reply.Success = false
		}
		return nil
	case err := <-errChan:
		fmt.Println("Call failed")
		return err
	}
}
