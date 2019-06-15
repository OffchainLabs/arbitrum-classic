package coordinator

import (
	"bytes"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"
	"fmt"
	"errors"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arb-avm/evm"
	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-avm/vm"

	"github.com/offchainlabs/arb-validator/ethvalidator"
	"github.com/offchainlabs/arb-validator/valmessage"
)

type CoordinatorServer struct {
	coordinator *ethvalidator.ValidatorCoordinator

	requests chan ValidatorRequest
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

	go func() {
		tracker := NewTxTracker(man.Val.VmId)
		tracker.HandleTxResults(man.Val.CompletedCallChan, requests)
	}()

	return &CoordinatorServer{man, requests}
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


type FindLogsArgs struct {
	FromHeight string   `json:"fromHeight"`
	ToHeight   string   `json:"toHeight"`
	Address    string   `json:"address"`
	Topics     []string `json:"topics"`
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
	reply.Found = txInfo.Found
	if txInfo.Found {
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