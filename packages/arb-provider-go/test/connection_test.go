package test

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	jsonenc "encoding/json"
	"errors"
	"io/ioutil"
	"math"
	"math/big"
	brand "math/rand"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/channel"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethconnection"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/coordinator"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethvalidator"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"

	goarbitrum "github.com/offchainlabs/arbitrum/packages/arb-provider-go"
)

/********************************************/
/*    Validators                            */
/********************************************/
func setupValidators(coordinatorKey string, followerKey string, t *testing.T) error {
	seed := time.Now().UnixNano()
	// seed := int64(1559616168133477000)
	brand.Seed(seed)

	jsonFile, err := os.Open("bridge_eth_addresses.json")
	if err != nil {
		t.Errorf("setupValidators Open error %v", err)
		return err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := jsonFile.Close(); err != nil {
		t.Errorf("setupValidators ReadAll error %v", err)
		return err
	}
	t.Log("bridge_eth_addresses.json loaded")
	var connectionInfo ethconnection.ArbAddresses
	if err := jsonenc.Unmarshal(byteValue, &connectionInfo); err != nil {
		t.Errorf("setupValidators Unmarshal error %v", err)
		return err
	}

	mach, err := loader.LoadMachineFromFile("contract.ao", true, "test")
	if err != nil {
		t.Errorf("setupValidators LoadMachineFromFile error %v", err)
		return err
	}

	key1, err := crypto.HexToECDSA(coordinatorKey)
	if err != nil {
		t.Errorf("setupValidators HexToECDSA error %v", err)
		return err
	}
	key2, err := crypto.HexToECDSA(followerKey)
	if err != nil {
		t.Errorf("setupValidators HexToECDSA error %v", err)
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

	t.Log("creating coordinator")
	// Validator creation
	val1, err := ethvalidator.NewValidator(key1, connectionInfo, ethURL)
	if err != nil {
		t.Error(err)
		return err
	}

	val2, err := ethvalidator.NewValidator(key2, connectionInfo, ethURL)
	if err != nil {
		t.Error(err)
		return err
	}

	address, err := val1.LaunchChannel(context.Background(), config, mach.Hash())
	if err != nil {
		t.Error(err)
		return err
	}

	server, err := coordinator.NewRPCServer(val1, address, mach, config)
	if err != nil {
		t.Error(err)
		return err
	}

	// follower/challenger creation
	challenger, err := channel.NewValidatorFollower(
		"Bob",
		val2,
		mach,
		config,
		false,
		math.MaxInt32, // maxCallSteps,
		math.MaxInt32, // maxUnanSteps,
		"wss://127.0.0.1:1236/ws",
	)
	if err != nil {
		t.Errorf("setupValidators NewValidatorFollower error %v", err)
		return err
	}

	if err := server.Run(context.Background()); err != nil {
		t.Errorf("setupValidators coordinator run error %v", err)
		return err
	}

	if err := challenger.Run(context.Background()); err != nil {
		t.Errorf("setupValidators challenger run error %v", err)
		return err
	}

	t.Log("challenger created")
	t.Log("starting RPCServerVM")

	// Run server
	s := rpc.NewServer()

	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	if err := s.RegisterService(server, "Validator"); err != nil {
		t.Errorf("setupValidators RegisterService error %v", err)
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
		return nil, err
	}

	privateKeyBytes, _ := hex.DecodeString(coordinatorKey)
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

type ListenerError struct {
	ListenerName string
	Err          error
}

func startFibTestEventListener(fibonacci *Fibonacci, ch chan interface{}, t *testing.T) {
	go func() {
		evCh := make(chan *FibonacciTestEvent, 2)
		start := uint64(0)
		watch := &bind.WatchOpts{
			Context: context.Background(),
			Start:   &start,
		}
		sub, err := fibonacci.WatchTestEvent(watch, evCh)
		if err != nil {
			t.Errorf("WatchTestEvent error %v", err)
			return
		}
		defer sub.Unsubscribe()
		errChan := sub.Err()
		for {
			select {
			case ev, ok := <-evCh:
				if ok {
					ch <- ev
				} else {
					ch <- &ListenerError{"FibonacciTestEvent ", errors.New("channel closed")}
					return
				}
			case err, ok := <-errChan:
				if ok {
					ch <- &ListenerError{"FibonacciTestEvent error:", err}
				} else {
					ch <- &ListenerError{"FibonacciTestEvent ", errors.New("error channel closed")}
					return
				}
			}
		}
	}()
}

func TestFib(t *testing.T) {
	session, err := RunValidators(t)
	if err != nil {
		t.Errorf("Validator setup error %v", err)
		t.FailNow()
	}

	t.Run("TestFibResult", func(t *testing.T) {
		fibsize := 15
		fibnum := 11
		_, err := session.GenerateFib(big.NewInt(int64(fibsize)))
		if err != nil {
			t.Errorf("GenerateFib error %v", err)
			return
		}
		fibval, err := session.GetFib(big.NewInt(int64(fibnum)))
		if err != nil {
			t.Errorf("GetFib error %v", err)
			return
		}
		if fibval.Cmp(big.NewInt(144)) != 0 { // 11th fibanocci number
			t.Errorf("GetFib error - expected %v got %v", big.NewInt(int64(144)), fibval)
		}
	})

	t.Run("TestEvent", func(t *testing.T) {
		eventChan := make(chan interface{}, 2)
		startFibTestEventListener(session.Contract, eventChan, t)
		testEventRcvd := false

		fibsize := 15
		time.Sleep(5 * time.Second)
		_, err := session.GenerateFib(big.NewInt(int64(fibsize)))
		if err != nil {
			t.Errorf("GenerateFib error %v", err)
			return
		}

	Loop:
		for ev := range eventChan {
			switch event := ev.(type) {
			case *FibonacciTestEvent:
				testEventRcvd = true
				break Loop
			case ListenerError:
				t.Errorf("errorEvent %v %v", event.ListenerName, event.Err)
				break Loop
			default:
				t.Error("eventLoop: unknown event type", ev)
				break Loop
			}
		}
		if testEventRcvd != true {
			t.Error("eventLoop: FibonacciTestEvent not received")
		}
	})

}
