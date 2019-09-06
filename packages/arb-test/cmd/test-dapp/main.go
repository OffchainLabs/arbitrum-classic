package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	jsonenc "encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	goarbitrum "github.com/offchainlabs/arbitrum/packages/arb-provider-go"
	"github.com/offchainlabs/arbitrum/packages/arb-test/chain"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/coordinator"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethvalidator"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
	"math"
	"math/big"
	"net/http"
	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	//"github.com/offchainlabs/arbitrum/packages/arb-validator/coordinator"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"io/ioutil"
	"log"
	brand "math/rand"
	//"net/http"
	"os"
	"time"
)

func main() {
	conn, auth, _ := setupValidators()
	//setupValidators()

	//conn, err = goarbitrum.Dial("", privateKeyBytes, _computePubKeyString(privateKeyBytes))
	var fibAddr common.Address
	fibAddr = common.HexToAddress("0x2EEBB8EE9c377caBC476654ca4aba016ECA1B9fc")
	fib, _ := chain.NewFibonacci(fibAddr, conn)
	eventChan := chain.StartEventListeners(fib)

	//Wrap the Token contract instance into a session
	fibonacciSession := &chain.FibonacciSession{
		Contract: fib,
		CallOpts: bind.CallOpts{
			From: auth.From,
		},
		TransactOpts: *auth,
	}

	go func() {
		eventLoop(fibonacciSession, eventChan)
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("generating fib")
	fibonacciSession.GenerateFib(big.NewInt(int64(5)))
	fmt.Println("getting fib")
	fibonacciSession.GetFib(big.NewInt(int64(2)))
	time.Sleep(30 * time.Second)
}

/********************************************/
/*    Validators                            */
/********************************************/
func setupValidators() (bind.ContractBackend, *bind.TransactOpts, error) {

	seed := time.Now().UnixNano()
	// seed := int64(1559616168133477000)
	fmt.Println("seed", seed)
	brand.Seed(seed)

	jsonFile, err := os.Open("bridge_eth_addresses.json")
	//jsonFile, err := os.Open("/Users/tobryan/work/arbitrum/validator-states/validator0/bridge_eth_addresses.json")
	if err != nil {
		log.Fatalln(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := jsonFile.Close(); err != nil {
		log.Fatalln(err)
	}

	var connectionInfo ethbridge.ArbAddresses
	if err := jsonenc.Unmarshal(byteValue, &connectionInfo); err != nil {
		log.Fatalln(err)
	}

	machine, err := loader.LoadMachineFromFile("contract.ao", true, "cpp")
	if err != nil {
		log.Fatal("Loader Error: ", err)
	}

	key1, err := crypto.HexToECDSA("ffb2b26161e081f0cdf9db67200ee0ce25499d5ee683180a9781e6cceb791c39")
	if err != nil {
		log.Fatal(err)
	}
	key2, err := crypto.HexToECDSA("979f020f6f6f71577c09db93ba944c89945f10fade64cfc7eb26137d5816fb76")
	if err != nil {
		log.Fatal(err)
	}

	var vmID [32]byte
	_, err = rand.Read(vmID[:])
	if err != nil {
		log.Fatal(err)
	}

	auth1 := bind.NewKeyedTransactor(key1)
	auth2 := bind.NewKeyedTransactor(key2)

	validators := []common.Address{auth1.From, auth2.From}
	escrowRequired := big.NewInt(10)
	config := valmessage.NewVMConfiguration(
		10,
		escrowRequired,
		common.Address{}, // Address 0 is eth
		validators,
		200000,
		common.Address{}, // Address 0 means no owner
	)
	ethURL := "ws://127.0.0.1:7545"

	// Validator creation
	var server *coordinator.RPCServer
	go func() {
		server = coordinator.NewRPCServer(machine, key1, validators, connectionInfo, ethURL)
	}()
	time.Sleep(500 * time.Millisecond)

	challenger, err := ethvalidator.NewValidatorFollower(
		"Bob",
		machine,
		key2,
		config,
		true,
		math.MaxInt32, // maxCallSteps
		connectionInfo,
		ethURL,
		"wss://127.0.0.1:1236/ws",
		math.MaxInt32, // maxUnanSteps
	)
	if err != nil {
		log.Fatalf("Failed to create follower %v\n", err)
	}
	err = challenger.Run()
	if err != nil {
		log.Fatal(err)
	}

	receiptChan, errChan := challenger.DepositFunds(context.Background(), escrowRequired)
	select {
	case receipt := <-receiptChan:
		if receipt.Status == 0 {
			log.Fatalln("Challenger could not deposit funds")
		}
	case err := <-errChan:
		log.Fatal(err)
	}

	// Run server
	s := rpc.NewServer()

	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	time.Sleep(5000 * time.Millisecond)

	if err := s.RegisterService(server, "Validator"); err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.Handle("/", s).Methods("GET", "POST", "OPTIONS")
	//attachProfiler(r)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	go func() {
		err = http.ListenAndServe(":1235", handlers.CORS(headersOk, originsOk, methodsOk)(r))
		if err != nil {
			panic(err)
		}
	}()

	privateKeyBytes, err := hex.DecodeString("ffb2b26161e081f0cdf9db67200ee0ce25499d5ee683180a9781e6cceb791c39")
	contract, err := goarbitrum.Dial("", privateKeyBytes, _computePubKeyString(privateKeyBytes))
	auth := bind.NewKeyedTransactor(key1)
	auth.GasLimit = 100000000

	return contract, auth, err

}

func _computePubKeyString(privKeyBytes []byte) string {
	privKey, err := crypto.ToECDSA(privKeyBytes)
	if err != nil {
		panic(err)
	}
	pubKey := privKey.Public().(*ecdsa.PublicKey)
	buf := crypto.FromECDSAPub(pubKey)
	return hexutil.Encode(buf)
}

func eventLoop(session *chain.FibonacciSession, eventChan chan interface{}) {
	fmt.Println("event loop")
	for ev := range eventChan {

		switch event := ev.(type) {
		case *chain.FibonacciTestEvent:
			fmt.Println("test event")
			fmt.Print(event.Number)
			return
		case chain.ListenerError:
			log.Println("errorEvent", event.ListenerName, event.Err)
		default:
			log.Println("eventLoop: unknown event type", ev)
		}
	}
	fmt.Println("exiting eventLoop")
}
