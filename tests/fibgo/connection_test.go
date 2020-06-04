package test

import (
	"context"
	jsonenc "encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"net"
	"os"
	"testing"
	"time"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/utils"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	goarbitrum "github.com/offchainlabs/arbitrum/packages/arb-provider-go"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupmanager"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupvalidator"
)

var db1 = "testman1db"
var db2 = "testman2db"

/********************************************/
/*    Validators                            */
/********************************************/
func setupValidators(
	coordinatorKey string,
	followerKey string,
	t *testing.T,
) error {
	seed := time.Now().UnixNano()
	// seed := int64(1559616168133477000)
	rand.Seed(seed)

	ethURL := test.GetEthUrl()
	contract := "contract.ao"

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
	var connectionInfo ethbridge.ArbAddresses
	if err := jsonenc.Unmarshal(byteValue, &connectionInfo); err != nil {
		t.Errorf("setupValidators Unmarshal error %v", err)
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

	ethclint1, err := ethclient.Dial(ethURL)
	if err != nil {
		return err
	}

	ethclint2, err := ethclient.Dial(ethURL)
	if err != nil {
		return err
	}

	client1 := ethbridge.NewEthAuthClient(ethclint1, auth1)
	client2 := ethbridge.NewEthAuthClient(ethclint2, auth2)

	ckpFac := checkpointing.NewDummyCheckpointerFactory(contract)

	checkpointer1 := ckpFac.New(context.TODO())
	config := valprotocol.ChainParams{
		StakeRequirement:        big.NewInt(10),
		GracePeriod:             common.TimeTicks{big.NewInt(13000 * 2)},
		MaxExecutionSteps:       250000,
		MaxBlockBoundsWidth:     20,
		MaxTimestampBoundsWidth: 600,
		ArbGasSpeedLimitPerTick: 200000,
	}

	factory, err := client1.NewArbFactory(connectionInfo.ArbFactoryAddress())
	if err != nil {
		return err
	}

	mach, err := checkpointer1.GetInitialMachine()
	if err != nil {
		return err
	}

	ctx := context.Background()

	rollupAddress, err := factory.CreateRollup(
		ctx,
		mach.Hash(),
		config,
		client1.Address(),
	)

	rollupActor1, err := client1.NewRollup(rollupAddress)
	if err != nil {
		return err
	}
	rollupActor2, err := client2.NewRollup(rollupAddress)
	if err != nil {
		return err
	}

	if err := os.RemoveAll(db1); err != nil {
		log.Fatal(err)
	}

	if err := os.RemoveAll(db2); err != nil {
		log.Fatal(err)
	}

	manager1, err := rollupmanager.CreateManagerAdvanced(
		ctx,
		rollupAddress,
		true,
		client1,
		ckpFac,
	)
	if err != nil {
		return err
	}
	manager1.AddListener(&rollup.AnnouncerListener{"chainObserver1: "})

	validatorListener1 := rollup.NewValidatorChainListener(
		context.Background(),
		rollupAddress,
		rollupActor1,
	)
	err = validatorListener1.AddStaker(client1)
	if err != nil {
		return err
	}
	manager1.AddListener(validatorListener1)

	manager2, err := rollupmanager.CreateManagerAdvanced(
		ctx,
		rollupAddress,
		true,
		client2,
		ckpFac,
	)
	if err != nil {
		return err
	}
	manager2.AddListener(&rollup.AnnouncerListener{"chainObserver2: "})

	validatorListener2 := rollup.NewValidatorChainListener(
		context.Background(),
		rollupAddress,
		rollupActor2,
	)
	err = validatorListener2.AddStaker(client2)
	if err != nil {
		return err
	}
	manager2.AddListener(validatorListener2)

	go func() {
		server := rollupvalidator.NewRPCServer(
			manager2,
			time.Second*60,
		)
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

		err := utils.LaunchRPC(s, "1235")
		if err != nil {
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
			break waitloop
		case <-time.After(time.Second * 5):
			t.Fatal("Couldn't connect to rpc")
		}
	}

	return nil

}

func RunValidators(
	t *testing.T,
) (*FibonacciSession, *goarbitrum.ArbConnection, error) {
	ethURL := test.GetEthUrl()
	key1 := "ffb2b26161e081f0cdf9db67200ee0ce25499d5ee683180a9781e6cceb791c39"
	key2 := "979f020f6f6f71577c09db93ba944c89945f10fade64cfc7eb26137d5816fb76"
	key3 := "d26a199ae5b6bed1992439d1840f7cb400d0a55a0c9f796fa67d7c571fbb180e"
	fibAddrHex := "0x895521964D724c8362A36608AAf09A3D7d0A0445"

	err := setupValidators(key1, key2, t)
	if err != nil {
		t.Errorf("Validator setup error %v", err)
		return nil, nil, err
	}

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

	var fibAddr common.Address
	fibAddr = common.HexToAddress(fibAddrHex)
	fib, err := NewFibonacci(fibAddr.ToEthAddress(), conn)
	if err != nil {
		t.Errorf("NewFibonacci error %v", err)
		return nil, nil, err
	}

	//Wrap the Token contract instance into a session
	fibonacciSession := &FibonacciSession{
		Contract: fib,
		CallOpts: bind.CallOpts{
			From: auth.From,
		},
		TransactOpts: *auth,
	}

	return fibonacciSession, conn, nil
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
			return nil, errors.New("Timed out waiting for receipt")
		default:
		}
		receipt, err := client.TransactionReceipt(
			context.Background(),
			txhash.ToEthHash(),
		)
		if err == nil {
			return receipt, nil
		}
		if err.Error() == "not found" {
			continue
		}
		log.Println("GetMessageResult error:", err)
		return nil, err
	}
}

func TestFib(t *testing.T) {
	session, client, err := RunValidators(t)
	if err != nil {
		t.Errorf("Validator setup error %v", err)
		t.FailNow()
	}

	t.Run("TestFibResult", func(t *testing.T) {
		fibsize := 15
		fibnum := 11

		tx, err := session.GenerateFib(big.NewInt(int64(fibsize)))
		if err != nil {
			t.Errorf("GenerateFib error %v", err)
			return
		}
		_, err = waitForReceipt(
			client,
			tx,
			common.NewAddressFromEth(session.TransactOpts.From),
			time.Second*60,
		)
		if err != nil {
			t.Errorf("GenerateFib receipt error %v", err)
			return
		}
		fibval, err := session.GetFib(big.NewInt(int64(fibnum)))
		if err != nil {
			t.Errorf("GetFib error %v", err)
			return
		}
		if fibval.Cmp(big.NewInt(144)) != 0 { // 11th fibanocci number
			t.Errorf(
				"GetFib error - expected %v got %v",
				big.NewInt(int64(144)),
				fibval,
			)
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

	if err := os.RemoveAll(db1); err != nil {
		log.Fatal(err)
	}

	if err := os.RemoveAll(db2); err != nil {
		log.Fatal(err)
	}
}
