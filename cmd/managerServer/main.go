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
	"crypto/rand"
	jsonenc "encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/gorilla/websocket"
	"github.com/offchainlabs/arb-validator/valmessage"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"net/http/pprof"
	"os"
	"time"

	"github.com/offchainlabs/arb-avm/evm"
	"github.com/offchainlabs/arb-avm/loader"
	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-avm/vm"

	"github.com/offchainlabs/arb-validator/ethvalidator"
)

import _ "net/http/pprof"

type TxInfo struct {
	found   bool
	success bool
	txIndex int
}

type TxRequest struct {
	txIndex    int
	resultChan chan<- TxInfo
}

type ManagerServer struct {
	coordinator *ethvalidator.ValidatorCoordinator
	follower    *ethvalidator.ValidatorFollower

	txRequests chan TxRequest
	txMetadata map[int]SendMessageArgs
}

func (m *ManagerServer) handleTxResults() {
	go func() {
		txIndex := 0
		txRequestIndex := 0
		transactions := make(map[int]TxInfo)
		for {
			select {
			case result := <-m.coordinator.CompletedCallChan:
				switch result := result.Result.(type) {
				case evm.Stop:
					transactions[txIndex] = TxInfo{true, true, txIndex}
					fmt.Println("EVMStop", result)
				case evm.Return:
					fmt.Println("EVMReturn", result)
					transactions[txIndex] = TxInfo{true, true, txIndex}
				case evm.Revert:
					transactions[txIndex] = TxInfo{true, false, txIndex}
					fmt.Println("EVMRevert", result)
				}
				txIndex++
			case request := <-m.txRequests:
				if request.txIndex == -1 {
					request.resultChan <- TxInfo{false, false, txRequestIndex}
					txRequestIndex += 1
				} else {
					tx, ok := transactions[request.txIndex]
					if ok {
						request.resultChan <- tx
					} else {
						request.resultChan <- TxInfo{false, false, request.txIndex}
					}
				}
			}
		}
	}()

}

func NewManagerServer(
	machine *vm.Machine,
	connectionInfo ethvalidator.ArbAddresses,
	ethURL string,
) *ManagerServer {
	key1, err := crypto.HexToECDSA("ffb2b26161e081f0cdf9db67200ee0ce25499d5ee683180a9781e6cceb791c39")
	if err != nil {
		log.Fatal(err)
	}
	key2, err := crypto.HexToECDSA("979f020f6f6f71577c09db93ba944c89945f10fade64cfc7eb26137d5816fb76")
	if err != nil {
		log.Fatal(err)
	}

	var vmId [32]byte
	_, err = rand.Read(vmId[:])
	if err != nil {
		log.Fatal(err)
	}

	auth1 := bind.NewKeyedTransactor(key1)
	auth2 := bind.NewKeyedTransactor(key2)

	// Commit all pending transactions in the simulator and print the names again

	validators := []common.Address{auth1.From, auth2.From}
	config := valmessage.NewVMConfiguration(10, big.NewInt(10), validators, 200000)

	man1, err := ethvalidator.NewValidatorCoordinator("Alice", machine.Clone(), key1, config, false, connectionInfo, ethURL)
	if err != nil {
		log.Fatal(err)
	}

	//_, err = man1.MintArbsToUser(auth1.From, big.NewInt(100000))
	//if err != nil {
	//	log.Fatalf("MintArbsToUser1 %v", err)
	//}
	//
	//_, err = man1.MintArbsToUser(auth2.From, big.NewInt(100000))
	//if err != nil {
	//	log.Fatalf("MintArbsToUser2 %v", err)
	//}

	go func() {
		err := man1.RunServer()
		fmt.Println("Running server", err)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = man1.Run()
	if err != nil {
		log.Fatal(err)
	}

	retChan, errChan := man1.CreateVM(time.Second * 10)

	man2, err := ethvalidator.NewValidatorFollower("Bob", machine, key2, config, false, connectionInfo, ethURL, "wss://127.0.0.1:1236/ws")
	if err != nil {
		log.Fatalf("Failed to create follower %v\n", err)
	}

	err = man2.Run()
	if err != nil {
		log.Fatal(err)
	}

	select {
	case <-retChan:
	case err := <-errChan:
		log.Fatalf("Failed to create vm: %v", err)
	}

	//time.Sleep(time.Second)

	//go func() {
	//	for {
	//		time.Sleep(2 * time.Second)
	//		err := man1.AdvanceBlockchain(1)
	//		if err != nil {
	//			panic(err)
	//		}
	//	}
	//}()

	time.Sleep(500 * time.Millisecond)
	txRequests := make(chan TxRequest, 100)
	txMetadata := make(map[int]SendMessageArgs)
	return &ManagerServer{man1, man2, txRequests, txMetadata}
}

type SendMessageArgs struct {
	Address   string `json:"address"`
	Data      string `json:"data"`
	TokenType string `json:"tokenType"`
	Amount    string `json:"amount"`
	Sender    string `json:"sender"`
}

type SendMessageReply struct {
	TxHash string `json:"hash"`
}

func (m *ManagerServer) SendMessage(r *http.Request, args *SendMessageArgs, reply *SendMessageReply) error {
	addressBytes, err := hexutil.Decode(args.Address)
	if err != nil {
		return err
	}
	addressInt := new(big.Int).SetBytes(addressBytes[:])

	dataBytes, err := hexutil.Decode(args.Data)
	if err != nil {
		return err
	}

	senderBytes, err := hexutil.Decode(args.Sender)
	if err != nil {
		return err
	}
	var sender common.Address
	copy(sender[:], senderBytes)

	data, err := evm.BytesToSizedByteArray(dataBytes)
	if err != nil {
		return err
	}

	amount, success := new(big.Int).SetString(args.Amount[2:], 16)
	if !success {
		return errors.New("amount wasn't a number")
	}

	tokenBytes, err := hexutil.Decode(args.TokenType)
	if err != nil {
		return err
	}
	var tokenType [21]byte
	copy(tokenType[:], tokenBytes)

	resultChan := make(chan TxInfo)
	m.txRequests <- TxRequest{-1, resultChan}

	txInfo := <-resultChan
	successChan, errChan := m.coordinator.InitiateUnanimousAssertion([]protocol.Message{ // retChan, errChan
		protocol.NewSimpleMessage(value.NewTuple2(value.NewIntValue(addressInt), data), tokenType, amount, sender),
	})
	select {
	case result := <-successChan:
		fmt.Println("Result", result)
	case err := <-errChan:
		return err
	}

	intBytes := value.NewInt64Value(int64(txInfo.txIndex)).ToBytes()
	reply.TxHash = hexutil.Encode(intBytes[:])
	m.txMetadata[txInfo.txIndex] = *args
	return nil
}

type GetMessageResultArgs struct {
	TxHash string `json:"txHash"`
}

type GetMessageResultReply struct {
	Found     bool   `json:"found"`
	Success   bool   `json:"success"`
	Address   string `json:"address"`
	Data      string `json:"data"`
	TokenType string `json:"tokenType"`
	Amount    string `json:"amount"`
	Sender    string `json:"sender"`
}

func (m *ManagerServer) GetMessageResult(r *http.Request, args *GetMessageResultArgs, reply *GetMessageResultReply) error {
	intBytes, err := hexutil.Decode(args.TxHash)
	if err != nil {
		return err
	}
	txIndex := new(big.Int).SetBytes(intBytes)

	resultChan := make(chan TxInfo)
	m.txRequests <- TxRequest{int(txIndex.Uint64()), resultChan}

	txInfo := <-resultChan
	reply.Found = txInfo.found
	reply.Success = txInfo.success

	txMetadata := m.txMetadata[int(txIndex.Uint64())]
	reply.Address = txMetadata.Address
	reply.Data = txMetadata.Data
	reply.TokenType = txMetadata.TokenType
	reply.Amount = txMetadata.Amount
	reply.Sender = txMetadata.Sender
	return nil
}

type CallMessageArgs struct {
	Address string `json:"address"`
	Data    string `json:"data"`
	Sender  string `json:"sender"`
}

type CallMessageReply struct {
	ReturnVal string
}

func (m *ManagerServer) CallMessage(r *http.Request, args *CallMessageArgs, reply *CallMessageReply) error {
	addressBytes, err := hexutil.Decode(args.Address)
	if err != nil {
		return err
	}
	addressInt := new(big.Int).SetBytes(addressBytes[:])

	dataBytes, err := hexutil.Decode(args.Data)
	if err != nil {
		return err
	}

	senderBytes, err := hexutil.Decode(args.Sender)
	if err != nil {
		return err
	}
	var sender common.Address
	copy(sender[:], senderBytes)

	data, err := evm.BytesToSizedByteArray(dataBytes)
	if err != nil {
		return err
	}

	resultChan := m.coordinator.Bot.RequestCall(protocol.NewSimpleMessage(value.NewTuple2(value.NewIntValue(addressInt), data), [21]byte{}, big.NewInt(0), sender), m.coordinator.Time)

	result := <-resultChan
	switch result := result.(type) {
	case evm.Stop:
		fmt.Println("EVMStop", result)
	case evm.Return:
		fmt.Println("EVMReturn", result)
		reply.ReturnVal = hexutil.Encode(result.ReturnVal)
		fmt.Println("Returned", len(result.ReturnVal))
	case evm.Revert:
		fmt.Println("EVMRevert", result)
	}
	return nil
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

func main() {

	jsonFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()

	var connectionInfo ethvalidator.ArbAddresses
	jsonenc.Unmarshal(byteValue, &connectionInfo)

	machine, _, err := loader.LoadMachineFromFile(os.Args[2], true)
	if err != nil {
		log.Fatal("Loader Error: ", err)
	}

	ethURL := os.Args[3]

	rpcInterface := NewManagerServer(machine, connectionInfo, ethURL)
	rpcInterface.handleTxResults()

	dataBytes, _ := hexutil.Decode("0xd0e30db0")
	data, _ := evm.BytesToSizedByteArray(dataBytes)
	addressInt, _ := new(big.Int).SetString("1221931544731715691820375137754886115565884929238", 10)
	_, err = rpcInterface.coordinator.SendEthMessage(
		value.NewTuple2(
			value.NewIntValue(addressInt),
			data,
		),
		big.NewInt(10000),
	)
	fmt.Println("Send error", err)
	time.Sleep(2000 * time.Millisecond)
	successChan, errChan := rpcInterface.coordinator.InitiateUnanimousAssertion(nil)
	select {
	case result := <-successChan:
		fmt.Println("Result", result)
	case err := <-errChan:
		panic(fmt.Sprintf("Error Running unan 1: %v", err))
	}
	successChan, errChan = rpcInterface.coordinator.InitiateUnanimousAssertion(nil)
	select {
	case result := <-successChan:
		fmt.Println("Result", result)
	case err := <-errChan:
		panic(fmt.Sprintf("Error Running unan 2: %v", err))
	}

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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// serveWs handles websocket requests from the peer.
func serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Made connection", conn)
	typ, message, err := conn.ReadMessage()
	fmt.Println("Got message", typ, message, err)
}
