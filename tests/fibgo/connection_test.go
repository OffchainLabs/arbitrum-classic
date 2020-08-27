package test

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	goarbitrum "github.com/offchainlabs/arbitrum/packages/arb-provider-go"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/rpc"
	utils2 "github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/utils"
	"log"
	"math/big"
	"math/rand"
	"net"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/chainlistener"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupmanager"
)

var db = "./testman"
var contract = arbos.Path()

func setupRollup(ctx context.Context, client ethutils.EthClient, auth *bind.TransactOpts) (common.Address, error) {
	config := valprotocol.ChainParams{
		StakeRequirement:        big.NewInt(10),
		GracePeriod:             common.TimeTicks{Val: big.NewInt(13000 * 2)},
		MaxExecutionSteps:       10000000000,
		ArbGasSpeedLimitPerTick: 200000,
	}

	factoryAddr, err := ethbridge.DeployRollupFactory(auth, client)
	if err != nil {
		return common.Address{}, err
	}

	arbClient := ethbridge.NewEthAuthClient(client, auth)

	factory, err := arbClient.NewArbFactory(common.NewAddressFromEth(factoryAddr))
	if err != nil {
		return common.Address{}, err
	}

	mach, err := loader.LoadMachineFromFile(contract, false, "cpp")
	if err != nil {
		return common.Address{}, err
	}

	rollupAddress, _, err := factory.CreateRollup(
		ctx,
		mach.Hash(),
		config,
		common.Address{},
	)
	return rollupAddress, err
}

/********************************************/
/*    Validators                            */
/********************************************/
func setupValidators(
	rollupAddress common.Address,
	client ethutils.EthClient,
	pks []*ecdsa.PrivateKey,
) error {
	if len(pks) < 1 {
		panic("must have at least 1 pks")
	}
	seed := time.Now().UnixNano()
	// seed := int64(1559616168133477000)
	rand.Seed(seed)

	clients := make([]arbbridge.ArbAuthClient, 0, len(pks))
	for _, pk := range pks {
		clients = append(clients, ethbridge.NewEthAuthClient(client, bind.NewKeyedTransactor(pk)))
	}

	ctx := context.Background()

	managers := make([]*rollupmanager.Manager, 0, len(clients))
	for _, client := range clients {
		rollupActor, err := client.NewRollup(rollupAddress)
		if err != nil {
			return err
		}

		dbName := db + "/" + client.Address().String()
		manager, err := rollupmanager.CreateManager(
			ctx,
			rollupAddress,
			rollupmanager.NewStressTestClient(client, time.Second*15),
			contract,
			dbName,
		)
		if err != nil {
			return err
		}

		manager.AddListener(&chainlistener.AnnouncerListener{Prefix: "validator " + client.Address().String() + ": "})

		validatorListener := chainlistener.NewValidatorChainListener(
			context.Background(),
			rollupAddress,
			rollupActor,
		)
		err = validatorListener.AddStaker(client)
		if err != nil {
			return err
		}
		manager.AddListener(validatorListener)
		managers = append(managers, manager)
	}

	return nil
}

func launchAggregator(client ethutils.EthClient, auth *bind.TransactOpts, rollupAddress common.Address) error {
	go func() {
		if err := rpc.LaunchAggregator(
			context.Background(),
			client,
			auth,
			rollupAddress,
			contract,
			db+"/aggregator",
			"2235",
			"9546",
			utils2.RPCFlags{},
			time.Second,
		); err != nil {
			log.Fatal(err)
		}
	}()

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			conn, err := net.DialTimeout(
				"tcp",
				net.JoinHostPort("127.0.0.1", "2235"),
				time.Second,
			)
			if err != nil || conn == nil {
				break
			}
			if err := conn.Close(); err != nil {
				return err
			}

			conn, err = net.DialTimeout(
				"tcp",
				net.JoinHostPort("127.0.0.1", "9546"),
				time.Second,
			)
			if err != nil || conn == nil {
				break
			}
			if err := conn.Close(); err != nil {
				return err
			}
			// Wait for the validator to catch up to head
			time.Sleep(time.Second * 2)
			return nil
		case <-time.After(time.Second * 5):
			return errors.New("couldn't connect to rpc")
		}
	}
}

type ListenerError struct {
	ListenerName string
	Err          error
}

func startFibTestEventListener(
	t *testing.T,
	fibonacci *Fibonacci,
	ch chan interface{},
	timeLimit time.Duration,
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
			case <-time.After(timeLimit):
				ch <- &ListenerError{
					ListenerName: "FibonacciTestEvent ",
					Err:          errors.New("timed out"),
				}
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
	client bind.DeployBackend,
	tx *types.Transaction,
	timeout time.Duration,
) (*types.Receipt, error) {
	ticker := time.NewTicker(timeout)
	for {
		select {
		case <-ticker.C:
			return nil, fmt.Errorf("timed out waiting for receipt for tx %v", tx.Hash().Hex())
		default:
		}
		receipt, err := client.TransactionReceipt(
			context.Background(),
			tx.Hash(),
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
	if err := os.RemoveAll(db); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.RemoveAll(db); err != nil {
			t.Fatal(err)
		}
	}()

	l1Client, pks := test.SimulatedBackend()
	go func() {
		t := time.NewTicker(time.Second * 2)
		for range t.C {
			l1Client.Commit()
		}
	}()

	if err := os.RemoveAll(db); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir(db, 0700); err != nil {
		t.Fatal(err)
	}

	rollupAddress, err := setupRollup(context.Background(), l1Client, bind.NewKeyedTransactor(pks[2]))
	if err != nil {
		t.Fatal(err)
	}

	if err := setupValidators(rollupAddress, l1Client, pks[3:5]); err != nil {
		t.Fatalf("Validator setup error %v", err)
	}

	if err := launchAggregator(
		l1Client,
		bind.NewKeyedTransactor(pks[1]),
		rollupAddress,
	); err != nil {
		t.Fatal(err)
	}
	pk := pks[0]

	//client, err := ethclient.Dial("http://localhost:8546")
	//if err != nil {
	//	t.Fatal(err)
	//}

	client := goarbitrum.Dial(
		"http://localhost:2235",
		pk,
		rollupAddress,
	)

	auth := bind.NewKeyedTransactor(pk)
	_, tx, _, err := DeployFibonacci(auth, client)
	if err != nil {
		t.Fatal("DeployFibonacci failed", err)
	}

	receipt, err := waitForReceipt(
		client,
		tx,
		time.Second*20,
	)
	if err != nil {
		t.Fatal("DeployFibonacci receipt error", err)
	}
	if receipt.Status != 1 {
		t.Fatal("tx deploying fib failed")
	}

	t.Log("Fib contract is at", receipt.ContractAddress.Hex())

	fib, err := NewFibonacci(receipt.ContractAddress, client)
	if err != nil {
		t.Fatal("connect fib failed", err)
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
		time.Second*20,
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
		startFibTestEventListener(t, session.Contract, eventChan, time.Second*20)
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
