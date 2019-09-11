package test

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
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
	"github.com/offchainlabs/arbitrum/packages/arb-provider-go"
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
func setupValidators(coordinatorKey string, followerKey string, t *testing.T) error {

	seed := time.Now().UnixNano()
	// seed := int64(1559616168133477000)
	fmt.Println("seed", seed)
	brand.Seed(seed)

	jsonFile, err := os.Open("bridge_eth_addresses.json")
	if err != nil {
		t.Errorf("bridge_eth_addresses error %v", err)
		return err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := jsonFile.Close(); err != nil {
		return err
	}

	var connectionInfo ethbridge.ArbAddresses
	if err := jsonenc.Unmarshal(byteValue, &connectionInfo); err != nil {
		return err
	}

	machine, err := loader.LoadMachineFromFile("contract.ao", true, "test")
	if err != nil {
		return err
	}

	key1, err := crypto.HexToECDSA(coordinatorKey)
	if err != nil {
		return err
	}
	key2, err := crypto.HexToECDSA(followerKey)
	if err != nil {
		return err
	}

	var vmID [32]byte
	_, err = rand.Read(vmID[:])
	if err != nil {
		return err
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
		return err
	}
	err = challenger.Run()
	if err != nil {
		return err
	}

	receiptChan, errChan := challenger.DepositFunds(context.Background(), escrowRequired)
	select {
	case receipt := <-receiptChan:
		if receipt.Status == 0 {
			return errors.New("Challenger could not deposit funds")
		}
	case err := <-errChan:
		return err
	}

	// start RPC server VM
	server := coordinator.StartRPCServerVM(coord)

	// Run server
	s := rpc.NewServer()

	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	if err := s.RegisterService(server, "Validator"); err != nil {
		return err
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
			t.Errorf("ListenAndServe error %v", err)
		}
	}()

	return nil

}

func _computePubKeyString(privKeyBytes []byte) (string, error) {
	privKey, err := crypto.ToECDSA(privKeyBytes)
	if err != nil {
		return "", err
	}
	pubKey := privKey.Public().(*ecdsa.PublicKey)
	buf := crypto.FromECDSAPub(pubKey)
	return hexutil.Encode(buf), nil
}

func RunValidators(t *testing.T) (*FibonacciSession, error) {
	coordinatorKey := "ffb2b26161e081f0cdf9db67200ee0ce25499d5ee683180a9781e6cceb791c39"
	followerKey := "979f020f6f6f71577c09db93ba944c89945f10fade64cfc7eb26137d5816fb76"
	err := setupValidators(coordinatorKey, followerKey, t)
	if err != nil {
		t.Errorf("Validator setup error %v", err)
		t.FailNow()
		return nil, err
	}
	//setupValidators()

	privateKeyBytes, _ := hex.DecodeString("ffb2b26161e081f0cdf9db67200ee0ce25499d5ee683180a9781e6cceb791c39")
	pubKey, err := _computePubKeyString(privateKeyBytes)
	if err != nil {
		t.Errorf("_computePubKeyString error %v", err)
		return nil, err
	}
	conn, dialerr := goarbitrum.Dial("", privateKeyBytes, pubKey)
	if dialerr != nil {
		t.Errorf("Dial error %v", dialerr)
		return nil, err
	}
	key1, err := crypto.HexToECDSA(coordinatorKey)
	if err != nil {
		t.Errorf("HexToECDSA error %v", err)
		return nil, err
	}
	auth := bind.NewKeyedTransactor(key1)
	auth.GasLimit = 100000000

	//conn, err = goarbitrum.Dial("", privateKeyBytes, _computePubKeyString(privateKeyBytes))
	var fibAddr common.Address
	fibAddr = common.HexToAddress("0x895521964D724c8362A36608AAf09A3D7d0A0445")
	fib, err := NewFibonacci(fibAddr, conn)
	if err != nil {
		t.Errorf("NewFibonacci error %v", err)
		return nil, err
	}

	//Wrap the Token contract instance into a session
	fibonacciSession := &FibonacciSession{
		Contract: fib,
		CallOpts: bind.CallOpts{
			From: auth.From,
		},
		TransactOpts: *auth,
	}

	return fibonacciSession, nil
}

func TestFib(t *testing.T) {
	session, err := RunValidators(t)
	if err != nil {
		t.Errorf("Validator setup error %v", err)
	}

	t.Run("TestFibResult", func(t *testing.T) {
		fibsize := 15
		fibnum := 11
		fmt.Println("generating fib")
		_, err := session.GenerateFib(big.NewInt(int64(fibsize)))
		if err != nil {
			t.Errorf("GenerateFib error %v", err)
			return
		}
		fmt.Println("getting fib")
		fibval, err := session.GetFib(big.NewInt(int64(fibnum)))
		if err != nil {
			t.Errorf("GetFib error %v", err)
			return
		}
		if fibval.Cmp(big.NewInt(144)) != 0 { // 11th fibanocci number
			t.Errorf("GetFib error - expected %v got %v", big.NewInt(int64(144)), fibval)
		}
		log.Printf("Fibonacci value number %v = %v", fibnum, fibval)

	})
	t.Run("TestEvent", func(t *testing.T) {
		eventChan := StartEventListeners(session.Contract)

		fibsize := 15
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Println("generating fib")
			_, err := session.GenerateFib(big.NewInt(int64(fibsize)))
			if err != nil {
				t.Errorf("GenerateFib error %v", err)
				return
			}
		}()

		fmt.Println("waiting for event")
	Loop:
		for ev := range eventChan {
			switch event := ev.(type) {
			case *FibonacciTestEvent:
				fmt.Printf("Received fibonacci test event %v", event.Number)
				break Loop
			case ListenerError:
				t.Errorf("errorEvent %v %v", event.ListenerName, event.Err)
			default:
				t.Error("eventLoop: unknown event type", ev)
			}
		}
		fmt.Println("end event test")

	})

}
