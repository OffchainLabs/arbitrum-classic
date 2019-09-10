package test

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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/offchainlabs/arbitrum/packages/arb-provider-go"
	"github.com/offchainlabs/arbitrum/packages/arb-provider-go/chain"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/coordinator"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethvalidator"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	brand "math/rand"
	"net/http"
	"os"
	"testing"
	"time"
)

/********************************************/
/*    Validators                            */
/********************************************/
func setupValidators(coordinatorKey string, followerKey string) error {

	seed := time.Now().UnixNano()
	// seed := int64(1559616168133477000)
	fmt.Println("seed", seed)
	brand.Seed(seed)

	jsonFile, err := os.Open("bridge_eth_addresses.json")
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

	key1, err := crypto.HexToECDSA(coordinatorKey)
	if err != nil {
		log.Fatal(err)
	}
	key2, err := crypto.HexToECDSA(followerKey)
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
	coord := coordinator.NewCoordinator(machine, key1, validators, connectionInfo, ethURL)

	// follower/challenger creation
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

	// start RPC server VM
	server := coordinator.StartRPCServerVM(coord)

	// Run server
	s := rpc.NewServer()

	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

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
			log.Fatal(err)
		}
	}()

	return err

}

func _computePubKeyString(privKeyBytes []byte) string {
	privKey, err := crypto.ToECDSA(privKeyBytes)
	if err != nil {
		log.Fatal(err)
	}
	pubKey := privKey.Public().(*ecdsa.PublicKey)
	buf := crypto.FromECDSAPub(pubKey)
	return hexutil.Encode(buf)
}

func eventLoop(session *chain.FibonacciSession, eventChan chan interface{}) {
	for ev := range eventChan {
		switch event := ev.(type) {
		case *chain.FibonacciTestEvent:
			fmt.Printf("Received fibonacci test event %v", event.Number)
			return
		case chain.ListenerError:
			log.Fatalln("errorEvent", event.ListenerName, event.Err)
		default:
			log.Fatalln("eventLoop: unknown event type", ev)
		}
	}
}

func RunValidators(t *testing.T) *chain.FibonacciSession {
	coordinatorKey := "ffb2b26161e081f0cdf9db67200ee0ce25499d5ee683180a9781e6cceb791c39"
	followerKey := "979f020f6f6f71577c09db93ba944c89945f10fade64cfc7eb26137d5816fb76"
	err := setupValidators(coordinatorKey, followerKey)
	if err != nil {
		t.Errorf("Validator setup error %v", err)
	}
	//setupValidators()

	privateKeyBytes, _ := hex.DecodeString("ffb2b26161e081f0cdf9db67200ee0ce25499d5ee683180a9781e6cceb791c39")
	conn, dialerr := goarbitrum.Dial("", privateKeyBytes, _computePubKeyString(privateKeyBytes))
	if dialerr != nil {
		t.Errorf("Dial error %v", dialerr)
	}
	key1, _ := crypto.HexToECDSA(coordinatorKey)
	auth := bind.NewKeyedTransactor(key1)
	auth.GasLimit = 100000000

	//conn, err = goarbitrum.Dial("", privateKeyBytes, _computePubKeyString(privateKeyBytes))
	var fibAddr common.Address
	fibAddr = common.HexToAddress("0x895521964D724c8362A36608AAf09A3D7d0A0445")
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
	return fibonacciSession
}

func TestFib(t *testing.T) {
	session := RunValidators(t)
	fibsize := 15
	fibnum := 11
	fmt.Println("generating fib")
	_, _ = session.GenerateFib(big.NewInt(int64(fibsize)))
	fmt.Println("getting fib")
	fibval, _ := session.GetFib(big.NewInt(int64(fibnum)))
	log.Printf("Fibonacci value number %v = %v", fibnum, fibval)
}
