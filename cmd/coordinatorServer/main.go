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

package main

import (
	"bytes"
	"crypto/ecdsa"
	jsonenc "encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"net/http/pprof"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arb-validator/valmessage"

	"github.com/offchainlabs/arb-avm/evm"
	"github.com/offchainlabs/arb-avm/loader"
	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-avm/vm"

	"github.com/offchainlabs/arb-validator/ethvalidator"
)

type TxInfo struct {
	found          bool
	assertionIndex int
	RawVal         value.Value
}

type ValidatorRequest interface {
}

type AssertionCountRequest struct {
	resultChan chan<- int
}

type TxRequest struct {
	txHash     [32]byte
	resultChan chan<- TxInfo
}

type FindLogsRequest struct {
	fromHeight *int64
	toHeight   *int64
	address    *big.Int
	topics     [][32]byte

	resultChan chan<- []LogInfo
}

type CoordinatorServer struct {
	coordinator *ethvalidator.ValidatorCoordinator

	requests chan ValidatorRequest
}

func (m *CoordinatorServer) requestAssertionCount() <-chan int {
	req := make(chan int, 1)
	m.requests <- AssertionCountRequest{req}
	return req
}

func (m *CoordinatorServer) requestTxInfo(txHash [32]byte) <-chan TxInfo {
	req := make(chan TxInfo, 1)
	m.requests <- TxRequest{txHash, req}
	return req
}

func (m *CoordinatorServer) requestFindLogs(
	fromHeight *int64,
	toHeight *int64,
	address *big.Int,
	topics [][32]byte,
) <-chan []LogInfo {
	req := make(chan []LogInfo, 1)
	m.requests <- FindLogsRequest{fromHeight, toHeight, address, topics, req}
	return req
}

type LogsInfo struct {
	msg  evm.EthMsg
	Logs []evm.Log
}

type AssertionInfo struct {
	TxLogs []LogsInfo
}

type LogResponse struct {
	Log evm.Log
	Msg evm.EthMsg
}

func (a *AssertionInfo) FindLogs(address *big.Int, topics [][32]byte) []LogResponse {
	logs := make([]LogResponse, 0)
	for _, txLogs := range a.TxLogs {
		for _, evmLog := range txLogs.Logs {
			if address != nil && !value.NewIntValue(address).Equal(evmLog.ContractId) {
				continue
			}

			if len(topics) > len(evmLog.Topics) {
				continue
			}

			for i, topic := range topics {
				if topic != evmLog.Topics[i] {
					continue
				}
			}
			logs = append(logs, LogResponse{evmLog, txLogs.msg})
		}
	}
	return logs
}

func NewAssertionInfo() *AssertionInfo {
	logs := make([]LogsInfo, 0)
	return &AssertionInfo{logs}
}

type TxTracker struct {
	txRequestIndex int
	transactions   map[[32]byte]TxInfo
	assertionInfo  []*AssertionInfo
	accountNonces  map[common.Address]uint64
	vmId           [32]byte
}

func NewTxTracker(vmId [32]byte) *TxTracker {
	return &TxTracker{
		txRequestIndex: 0,
		transactions:   make(map[[32]byte]TxInfo),
		assertionInfo:  make([]*AssertionInfo, 0),
		accountNonces:  make(map[common.Address]uint64),
		vmId:           vmId,
	}
}

func (tr *TxTracker) processFinalizedAssertion(assertion valmessage.FinalizedAssertion) {
	log.Println("Coordinator produced finalized assertion")
	info := NewAssertionInfo()
	for _, res := range assertion.NewLogs() {
		evmVal, err := evm.ProcessLog(res)
		if err != nil {
			log.Printf("VM produced invalid evm result: %v\n", err)
		}

		msg := evmVal.GetEthMsg()
		msgHash := msg.MsgHash(tr.vmId)

		log.Println("Coordinator got response for", hexutil.Encode(msgHash[:]))
		txInfo := TxInfo{
			found:          true,
			assertionIndex: 0,
			RawVal:         res,
		}
		txInfo.assertionIndex = len(tr.assertionInfo)
		switch evmVal := evmVal.(type) {
		case evm.Stop:
			info.TxLogs = append(info.TxLogs, LogsInfo{evmVal.Msg, evmVal.Logs})
		case evm.Return:
			info.TxLogs = append(info.TxLogs, LogsInfo{evmVal.Msg, evmVal.Logs})
		case evm.Revert:
		}
		tr.transactions[msgHash] = txInfo
	}
	tr.assertionInfo = append(tr.assertionInfo, info)
}

func (tr *TxTracker) processRequest(request ValidatorRequest) {
	switch request := request.(type) {
	case AssertionCountRequest:
		request.resultChan <- len(tr.assertionInfo) - 1
	case TxRequest:
		tx, ok := tr.transactions[request.txHash]
		if ok {
			request.resultChan <- tx
		} else {
			request.resultChan <- TxInfo{found: false}
		}
	case FindLogsRequest:
		startHeight := int64(0)
		endHeight := int64(len(tr.assertionInfo))
		if request.fromHeight != nil && *request.fromHeight > int64(0) {
			startHeight = *request.fromHeight
		}
		if request.toHeight != nil {
			endHeight = *request.toHeight + 1
		}
		logs := make([]LogInfo, 0)
		if startHeight >= int64(len(tr.assertionInfo)) {
			request.resultChan <- logs
			break
		}
		assertions := tr.assertionInfo[startHeight:endHeight]

		for i, assertion := range assertions {
			assertionLogs := assertion.FindLogs(request.address, request.topics)
			for j, evmLog := range assertionLogs {
				addressBytes := evmLog.Log.ContractId.ToBytes()
				topicStrings := make([]string, 0, len(evmLog.Log.Topics))
				for _, topic := range evmLog.Log.Topics {
					topicStrings = append(topicStrings, hexutil.Encode(topic[:]))
				}
				txHash := evmLog.Msg.MsgHash(tr.vmId)
				logs = append(logs, LogInfo{
					Address:          hexutil.Encode(addressBytes[12:]),
					BlockHash:        hexutil.Encode(txHash[:]),
					BlockNumber:      "0x" + strconv.FormatInt(int64(i), 16),
					Data:             hexutil.Encode(evmLog.Log.Data[:]),
					LogIndex:         "0x" + strconv.FormatInt(int64(j), 16),
					Topics:           topicStrings,
					TransactionIndex: "0x00",
					TransactionHash:  hexutil.Encode(txHash[:]),
				})
			}
		}
		request.resultChan <- logs
	}
}

func (tr *TxTracker) HandleTxResults(completedCalls chan valmessage.FinalizedAssertion, requests chan ValidatorRequest) {
	for {
		select {
		case finalizedAssertion := <-completedCalls:
			tr.processFinalizedAssertion(finalizedAssertion)
		case request := <-requests:
			tr.processRequest(request)

		}
	}
}

func NewCoordinatorServer(
	machine *vm.Machine,
	key *ecdsa.PrivateKey,
	validators []common.Address,
	connectionInfo ethvalidator.ArbAddresses,
	ethURL string,
) *CoordinatorServer {
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

	man, err := ethvalidator.NewValidatorCoordinator("Alice", machine.Clone(), key, config, false, connectionInfo, ethURL)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := man.Val.DepositEth(escrowRequired)
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
	requests := make(chan ValidatorRequest, 100)
	return &CoordinatorServer{man, requests}
}

type FindLogsArgs struct {
	FromHeight string   `json:"fromHeight"`
	ToHeight   string   `json:"toHeight"`
	Address    string   `json:"address"`
	Topics     []string `json:"topics"`
}

type LogInfo struct {
	Address          string   `json:"address"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Data             string   `json:"data"`
	LogIndex         string   `json:"logIndex"`
	Topics           []string `json:"topics"`
	TransactionIndex string   `json:"transactionIndex"`
	TransactionHash  string   `json:"transactionHash"`
}

type FindLogsReply struct {
	Logs []LogInfo `json:"logs"`
}

func (m *CoordinatorServer) FindLogs(r *http.Request, args *FindLogsArgs, reply *FindLogsReply) error {
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

type SendMessageArgs struct {
	Data      string `json:"data"`
	Signature string `json:"signature"`
}

type SendMessageReply struct {
	TxHash string `json:"hash"`
}

func (m *CoordinatorServer) SendMessage(r *http.Request, args *SendMessageArgs, reply *SendMessageReply) error {
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
		solsha3.Bytes32(m.coordinator.Val.VmId),
		solsha3.Bytes32(dataVal.Hash()),
		solsha3.Uint256(amount),
		tokenType[:],
	)

	signedMsg := solsha3.SoliditySHA3WithPrefix(solsha3.Bytes32(messageHash))
	pubkey, err := crypto.SigToPub(signedMsg, sigBytes)
	if err != nil {
		log.Printf("SendMessage: Failed to convert signature to pubkey, %v\n", err)
		return err
	}
	sender := crypto.PubkeyToAddress(*pubkey)
	log.Printf("Coordinator recieved transaction from %v\n", hexutil.Encode(sender[:]))
	senderArr := [32]byte{}
	copy(senderArr[12:], sender.Bytes())

	msg := protocol.Message{
		Data:        dataVal,
		TokenType:   tokenType,
		Currency:    amount,
		Destination: senderArr,
	}
	m.coordinator.SendMessage(ethvalidator.OffchainMessage{
		Message:   msg,
		Signature: sigBytes,
	})
	reply.TxHash = hexutil.Encode(messageHash)
	return nil
}

type GetMessageResultArgs struct {
	TxHash string `json:"txHash"`
}

type GetMessageResultReply struct {
	Found  bool   `json:"found"`
	RawVal string `json:"rawVal"`
}

func (m *CoordinatorServer) GetMessageResult(r *http.Request, args *GetMessageResultArgs, reply *GetMessageResultReply) error {
	txHashBytes, err := hexutil.Decode(args.TxHash)
	if err != nil {
		return err
	}
	txHash := [32]byte{}
	copy(txHash[:], txHashBytes)
	resultChan := m.requestTxInfo(txHash)

	txInfo := <-resultChan
	reply.Found = txInfo.found
	if txInfo.found {
		var buf bytes.Buffer
		_ = value.MarshalValue(txInfo.RawVal, &buf) // error can only occur from writes and bytes.Buffer is safe
		reply.RawVal = hexutil.Encode(buf.Bytes())
	}
	return nil
}

type GetAssertionCountReply struct {
	AssertionCount int `json:"assertionCount"`
}

func (m *CoordinatorServer) GetAssertionCount(r *http.Request, _ *struct{}, reply *GetAssertionCountReply) error {
	req := m.requestAssertionCount()
	reply.AssertionCount = <-req
	return nil
}

type GetVMInfoReply struct {
	VMId string `json:"vmId"`
}

func (m *CoordinatorServer) GetVMInfo(r *http.Request, _ *struct{}, reply *GetVMInfoReply) error {
	reply.VMId = hexutil.Encode(m.coordinator.Val.VmId[:])
	return nil
}

func (m *CoordinatorServer) TranslateToValue(r *http.Request, arg *string, reply *string) error {
	rawBytes, err := hexutil.Decode(*arg)
	if err != nil {
		return err
	}
	data, err := evm.BytesToSizedByteArray(rawBytes)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := value.MarshalValue(data, &buf); err != nil {
		return err
	}
	*reply = hexutil.Encode(buf.Bytes())
	return nil
}

type CallMessageArgs struct {
	Data   string `json:"data"`
	Sender string `json:"sender"`
}

type CallMessageReply struct {
	ReturnVal string
	Success   bool
}

func (m *CoordinatorServer) CallMessage(r *http.Request, args *CallMessageArgs, reply *CallMessageReply) error {
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
	log.Printf("Coordinator recieved call from %v\n", hexutil.Encode(sender[:]))

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

func AttachProfiler(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// Manually add support for paths linked to by index page at /debug/pprof/
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))
	router.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
}

// Launches the Coordinator validator with the following command line arguments:
// 1) Compiled Arbitrum bytecode file
// 2) private key file
// 3) public addresses file (newline separated)
// 4) Global EthBridge addresses json file
// 5) ethURL
func main() {
	// Check number of args
	if len(os.Args)-1 != 5 {
		log.Fatalln("Expected five arguments")
	}

	// 1) Compiled Arbitrum bytecode
	machine, err := loader.LoadMachineFromFile(os.Args[1], true)
	if err != nil {
		log.Fatal("Loader Error: ", err)
	}

	// 2) Private key
	keyFile, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	byteValue, err := ioutil.ReadAll(keyFile)
	if err != nil {
		log.Fatalln(err)
	}
	if err := keyFile.Close(); err != nil {
		log.Fatalln(err)
	}
	rawKey := strings.TrimSpace(string(byteValue))
	key, err := crypto.HexToECDSA(rawKey)
	if err != nil {
		log.Fatal("HexToECDSA private key error: ", err)
	}

	// 3) All public key addresses
	addrFile, err := os.Open(os.Args[3])
	if err != nil {
		log.Fatalln(err)
	}
	byteValue, err = ioutil.ReadAll(addrFile)
	if err != nil {
		log.Fatalln(err)
	}
	if err := addrFile.Close(); err != nil {
		log.Fatalln(err)
	}
	validatorHexAddrs := strings.Split(strings.TrimSpace(string(byteValue)), "\n")
	validators := make([]common.Address, len(validatorHexAddrs))
	for i, v := range validatorHexAddrs {
		validators[i] = common.HexToAddress(v)
	}

	// 4) Global EthBridge addresses json
	jsonFile, err := os.Open(os.Args[4])
	if err != nil {
		log.Fatalln(err)
	}
	byteValue, _ = ioutil.ReadAll(jsonFile)
	if err := jsonFile.Close(); err != nil {
		log.Fatalln(err)
	}

	var connectionInfo ethvalidator.ArbAddresses
	if err := jsonenc.Unmarshal(byteValue, &connectionInfo); err != nil {
		log.Fatalln(err)
	}

	// 5) URL
	ethURL := os.Args[5]

	// Validator creation
	rpcInterface := NewCoordinatorServer(machine, key, validators, connectionInfo, ethURL)

	go func() {
		tracker := NewTxTracker(rpcInterface.coordinator.Val.VmId)
		tracker.HandleTxResults(rpcInterface.coordinator.Val.CompletedCallChan, rpcInterface.requests)
	}()

	// Run server
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	if err := s.RegisterService(rpcInterface, "Validator"); err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.Handle("/", s).Methods("GET", "POST", "OPTIONS")
	AttachProfiler(r)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	err = http.ListenAndServe(":1235", handlers.CORS(headersOk, originsOk, methodsOk)(r))
	if err != nil {
		panic(err)
	}
}
