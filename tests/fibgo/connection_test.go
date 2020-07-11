package test

import (
	"context"
	jsonenc "encoding/json"
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/gotest"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup/chainlistener"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"net"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	goarbitrum "github.com/offchainlabs/arbitrum/packages/arb-provider-go"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupmanager"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupvalidator"
)

var db = "testman"

/********************************************/
/*    Validators                            */
/********************************************/
func setupValidators(
	hexKeys []string,
	t *testing.T,
) ([]arbbridge.ArbAuthClient, error) {
	if len(hexKeys) == 0 {
		panic("must supply at least one key")
	}
	seed := time.Now().UnixNano()
	// seed := int64(1559616168133477000)
	rand.Seed(seed)

	ethURL := test.GetEthUrl()

	jsonFile, err := os.Open("bridge_eth_addresses.json")

	if err != nil {
		t.Errorf("setupValidators Open error %v", err)
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := jsonFile.Close(); err != nil {
		t.Errorf("setupValidators ReadAll error %v", err)
		return nil, err
	}
	var connectionInfo ethbridge.ArbAddresses
	if err := jsonenc.Unmarshal(byteValue, &connectionInfo); err != nil {
		t.Errorf("setupValidators Unmarshal error %v", err)
		return nil, err
	}

	clients := make([]arbbridge.ArbAuthClient, 0, len(hexKeys))
	for _, hexKey := range hexKeys {
		key, err := crypto.HexToECDSA(hexKey)
		if err != nil {
			t.Errorf("setupValidators HexToECDSA error %v", err)
			return nil, err
		}

		auth := bind.NewKeyedTransactor(key)
		ethclint, err := ethclient.Dial(ethURL)
		if err != nil {
			return nil, err
		}

		client := ethbridge.NewEthAuthClient(ethclint, auth)
		clients = append(clients, client)
	}

	config := valprotocol.ChainParams{
		StakeRequirement:        big.NewInt(10),
		GracePeriod:             common.TimeTicks{Val: big.NewInt(13000 * 2)},
		MaxExecutionSteps:       1000000000,
		ArbGasSpeedLimitPerTick: 200000,
	}

	factory, err := clients[0].NewArbFactory(connectionInfo.ArbFactoryAddress())
	if err != nil {
		return nil, err
	}

	contract := gotest.TestMachinePath()

	mach, err := loader.LoadMachineFromFile(contract, false, "cpp")
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	rollupAddress, _, err := factory.CreateRollup(
		ctx,
		mach.Hash(),
		config,
		clients[0].Address(),
	)

	managers := make([]*rollupmanager.Manager, 0, len(clients))
	for _, client := range clients {
		rollupActor, err := client.NewRollup(rollupAddress)
		if err != nil {
			return nil, err
		}

		dbName := db + client.Address().String()

		if err := os.RemoveAll(dbName); err != nil {
			log.Fatal(err)
		}

		manager, err := rollupmanager.CreateManager(
			ctx,
			rollupAddress,
			rollupmanager.NewStressTestClient(client, time.Second*10),
			contract,
			dbName,
		)
		if err != nil {
			return nil, err
		}

		manager.AddListener(&chainlistener.AnnouncerListener{Prefix: "validator " + client.Address().String() + ": "})

		validatorListener := chainlistener.NewValidatorChainListener(
			context.Background(),
			rollupAddress,
			rollupActor,
		)
		err = validatorListener.AddStaker(client)
		if err != nil {
			return nil, err
		}
		manager.AddListener(validatorListener)
		managers = append(managers, manager)
	}

	go func() {
		server, err := rollupvalidator.NewRPCServer(
			managers[0],
			time.Second*60,
		)
		if err != nil {
			t.Fatal(err)
		}
		s := rpc.NewServer()
		s.RegisterCodec(
			json.NewCodec(),
			"application/json",
		)
		s.RegisterCodec(
			json.NewCodec(),
			"application/json;charset=UTF-8",
		)

		if err := s.RegisterService(server, "Validator"); err != nil {
			t.Fatal(err)
		}

		if err := utils.LaunchRPC(s, "1235", utils.RPCFlags{}); err != nil {
			t.Fatal(err)
		}
	}()

	ticker := time.NewTicker(time.Second)
waitloop:
	for {
		select {
		case <-ticker.C:
			conn, err := net.DialTimeout(
				"tcp",
				net.JoinHostPort("127.0.0.1", "1235"),
				time.Second,
			)
			if err != nil || conn == nil {
				continue
			}
			if err := conn.Close(); err != nil {
				t.Fatal(err)
			}
			// Wait for the validator to catch up to head
			time.Sleep(time.Second * 2)
			break waitloop
		case <-time.After(time.Second * 5):
			t.Fatal("Couldn't connect to rpc")
		}
	}

	return clients, nil

}

func SetupProvider(
	t *testing.T,
) (*goarbitrum.ArbConnection, *bind.TransactOpts, error) {
	ethURL := test.GetEthUrl()
	key3 := "d26a199ae5b6bed1992439d1840f7cb400d0a55a0c9f796fa67d7c571fbb180e"

	userKey, err := crypto.HexToECDSA(key3)
	if err != nil {
		t.Errorf("HexToECDSA error %v", err)
		return nil, nil, err
	}
	auth := bind.NewKeyedTransactor(userKey)

	ethclint, err := ethclient.Dial(ethURL)
	if err != nil {
		return nil, nil, err
	}

	conn, dialerr := goarbitrum.Dial(
		"http://localhost:1235",
		auth,
		ethclint,
	)
	if dialerr != nil {
		t.Errorf("Dial error %v", dialerr)
		return nil, nil, err
	}
	return conn, auth, nil
}

type ListenerError struct {
	ListenerName string
	Err          error
}

func startFibTestEventListener(
	fibonacci *Fibonacci,
	ch chan interface{},
	t *testing.T,
) {
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
					ch <- &ListenerError{
						"FibonacciTestEvent ",
						errors.New("channel closed"),
					}
					return
				}
			case err, ok := <-errChan:
				if ok {
					ch <- &ListenerError{
						"FibonacciTestEvent error:",
						err,
					}
				} else {
					ch <- &ListenerError{
						"FibonacciTestEvent ",
						errors.New("error channel closed"),
					}
					return
				}
			}
		}
	}()
}

func waitForReceipt(
	client *goarbitrum.ArbConnection,
	tx *types.Transaction,
	sender common.Address,
	timeout time.Duration,
) (*types.Receipt, error) {
	txhash := client.TxHash(tx, sender)
	ticker := time.NewTicker(timeout)
	for {
		select {
		case <-ticker.C:
			return nil, errors.New("timed out waiting for receipt")
		default:
		}
		receipt, err := client.TransactionReceipt(
			context.Background(),
			txhash.ToEthHash(),
		)
		if err != nil {
			if err.Error() == "not found" {
				continue
			}
			log.Println("GetMessageResult error:", err)
			return nil, err
		}
		return receipt, nil
	}
}

func TestFib(t *testing.T) {
	key1 := "ffb2b26161e081f0cdf9db67200ee0ce25499d5ee683180a9781e6cceb791c39"
	//key2 := "979f020f6f6f71577c09db93ba944c89945f10fade64cfc7eb26137d5816fb76"
	validatorClients, err := setupValidators([]string{key1}, t)
	if err != nil {
		t.Fatalf("Validator setup error %v", err)
	}

	client, auth, err := SetupProvider(t)
	if err != nil {
		t.Fatalf("Validator setup error %v", err)
	}
	_, tx, fib, err := DeployFibonacci(auth, client)
	if err != nil {
		t.Fatal("DeployFibonacci failed", err)
	}

	receipt, err := waitForReceipt(
		client,
		tx,
		common.NewAddressFromEth(auth.From),
		time.Second*60,
	)
	if err != nil {
		t.Fatal("DeployFibonacci receipt error", err)
	}
	if receipt.Status != 1 {
		t.Fatal("tx deploying fib failed")
	}

	//Wrap the Token contract instance into a session
	session := &FibonacciSession{
		Contract: fib,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Pending: true,
		},
		TransactOpts: *auth,
	}

	fibsize := 15
	fibnum := 11

	tx, err = session.GenerateFib(big.NewInt(int64(fibsize)))
	if err != nil {
		t.Fatal("GenerateFib error", err)
	}
	receipt, err = waitForReceipt(
		client,
		tx,
		common.NewAddressFromEth(session.TransactOpts.From),
		time.Second*60,
	)
	if err != nil {
		t.Fatal("GenerateFib receipt error", err)
	}
	if receipt.Status != 1 {
		t.Fatal("tx generating numbers failed")
	}

	fibval, err := session.GetFib(big.NewInt(int64(fibnum)))
	if err != nil {
		t.Fatal("GetFib error", err)
	}
	if fibval.Cmp(big.NewInt(144)) != 0 { // 11th fibanocci number
		t.Fatalf(
			"GetFib error - expected %v got %v",
			big.NewInt(int64(144)),
			fibval,
		)
	}

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

	for _, client := range validatorClients {
		if err := os.RemoveAll(db + client.Address().String()); err != nil {
			log.Fatal(err)
		}
	}
}
