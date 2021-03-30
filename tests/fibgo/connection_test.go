package test

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/rpc"
	utils2 "github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	golog "log"
	"math/big"
	"math/rand"
	"net"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	gethlog "github.com/ethereum/go-ethereum/log"

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

var logger zerolog.Logger

var db = "./testman"
var contract = arbos.Path()

func setupRollup(ctx context.Context, authClient *ethbridge.EthArbAuthClient) (common.Address, error) {
	config := valprotocol.ChainParams{
		StakeRequirement:        big.NewInt(10),
		StakeToken:              common.Address{},
		GracePeriod:             common.TimeTicks{Val: big.NewInt(13000 * 2)},
		MaxExecutionSteps:       10000000000,
		ArbGasSpeedLimitPerTick: 200000,
	}

	factoryAddr, err := ethbridge.DeployRollupFactory(ctx, authClient)
	if err != nil {
		return common.Address{}, err
	}

	factory, err := authClient.NewArbFactory(common.NewAddressFromEth(factoryAddr))
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
func setupValidators(ctx context.Context, rollupAddress common.Address, authClients []*ethbridge.EthArbAuthClient) error {
	if len(authClients) < 1 {
		panic("must have at least 1 authClient")
	}
	seed := time.Now().UnixNano()
	// seed := int64(1559616168133477000)
	rand.Seed(seed)

	managers := make([]*rollupmanager.Manager, 0, len(authClients))
	for _, authClient := range authClients {
		rollupActor, err := authClient.NewRollup(rollupAddress)
		if err != nil {
			return err
		}

		dbName := db + "/" + authClient.Address().String()
		manager, err := rollupmanager.CreateManager(
			ctx,
			rollupAddress,
			arbbridge.NewStressTestClient(authClient, time.Second*15),
			contract,
			dbName,
		)
		if err != nil {
			return err
		}

		manager.AddListener(ctx, chainlistener.NewAnnouncerListener(authClient.Address().String()))

		validatorListener := chainlistener.NewValidatorChainListener(
			ctx,
			rollupAddress,
			rollupActor,
		)
		err = validatorListener.AddStaker(authClient)
		if err != nil {
			return err
		}
		manager.AddListener(ctx, validatorListener)
		managers = append(managers, manager)
	}

	_ = managers

	return nil
}

func launchAggregator(client ethutils.EthClient, auth *bind.TransactOpts, rollupAddress common.Address) error {
	go func() {
		if err := rpc.LaunchAggregator(
			context.Background(),
			client,
			rollupAddress,
			contract,
			db+"/aggregator",
			"9546",
			"9547",
			utils2.RPCFlags{},
			time.Second,
			rpc.StatelessBatcherMode{Auth: auth},
		); err != nil {
			logger.Fatal().Stack().Err(err).Msg("LaunchAggregator failed")
		}
	}()

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			conn, err := net.DialTimeout(
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

func waitForReceipt(
	client bind.DeployBackend,
	tx *types.Transaction,
	timeout time.Duration,
) (*types.Receipt, error) {
	ticker := time.NewTicker(timeout)
	for {
		select {
		case <-ticker.C:
			return nil, errors.Errorf("timed out waiting for receipt for tx %v", tx.Hash().Hex())
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
			logger.Error().Stack().Err(err).Msg("Failure getting TransactionReceipt")
			return nil, err
		}
		return receipt, nil
	}
}

func TestFib(t *testing.T) {
	// TODO
	return
	if err := os.RemoveAll(db); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.RemoveAll(db); err != nil {
			t.Fatal(err)
		}
	}()

	gethlog.Root().SetHandler(gethlog.LvlFilterHandler(gethlog.LvlInfo, gethlog.StreamHandler(os.Stderr, gethlog.TerminalFormat(true))))

	// Enable line numbers in logging
	golog.SetFlags(golog.LstdFlags | golog.Lshortfile)

	// Print stack trace when `.Error().Stack().Err(err).` is added to zerolog call
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	// Print line number that log was created on
	logger = log.With().Caller().Str("component", "connection-test").Logger()

	ctx := context.Background()
	l1Backend, pks := test.SimulatedBackend()
	l1Client := &ethutils.SimulatedEthClient{SimulatedBackend: l1Backend}

	// pks[0]:     setupRollup     (L1)
	// pks[1,2]:   setupValidators (L1)

	// pks[3]:     launchAggregator            (not tied to client)
	// pks[4]:     DeployFibonacci and session (L2, not tied to client)

	auths := make([]*bind.TransactOpts, 0)
	authClients := make([]*ethbridge.EthArbAuthClient, 0)
	// 0-3 are on L1
	for _, pk := range pks[0:3] {
		auth := bind.NewKeyedTransactor(pk)
		auths = append(auths, auth)

		authClient, err := ethbridge.NewEthAuthClient(ctx, l1Client, auth)
		if err != nil {
			t.Fatal(err)
		}
		authClients = append(authClients, authClient)
	}
	// 3 just uses auth, authClient created inside launchAggregator
	auths = append(auths, bind.NewKeyedTransactor(pks[3]))

	// 4 is on L2, doesn't use ethbridge
	auths = append(auths, bind.NewKeyedTransactor(pks[4]))

	go func() {
		t := time.NewTicker(time.Second * 2)
		for range t.C {
			logger.Info().Msg("Commit")
			l1Client.Commit()
		}
	}()

	if err := os.RemoveAll(db); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir(db, 0700); err != nil {
		t.Fatal(err)
	}

	rollupAddress, err := setupRollup(ctx, authClients[0])
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Created rollup chain", rollupAddress)

	if err := setupValidators(ctx, rollupAddress, authClients[1:3]); err != nil {
		t.Fatalf("Validator setup error %v", err)
	}

	logger.Info().Msg("Validators setup, launching aggregator")

	if err := launchAggregator(
		l1Client,
		auths[3],
		rollupAddress,
	); err != nil {
		t.Fatal(err)
	}

	logger.Info().Msg("Launched aggregator, connecting to RPC")

	l2Client, err := ethclient.Dial("http://localhost:9546")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Connected to aggregator")

	logger.Info().Hex("account4", auths[4].From.Bytes()).Msg("Account being used to deploy fibonacci")

	// Do not wrap with MakeContract because auth is wrapped in session below
	auths[4].Nonce = big.NewInt(0)
	_, tx, _, err := arbostestcontracts.DeployFibonacci(auths[4], l2Client)
	if err != nil {
		t.Fatal("DeployFibonacci failed", err)
	}
	auths[4].Nonce = auths[4].Nonce.Add(auths[4].Nonce, big.NewInt(1))

	logger.Info().Hex("tx", tx.Hash().Bytes()).Msg("Fibonacci deployed")

	receipt, err := waitForReceipt(
		l2Client,
		tx,
		time.Second*20,
	)
	if err != nil {
		t.Fatal("DeployFibonacci receipt error", err)
	}
	if receipt.Status != 1 {
		t.Fatal("tx deploying fib failed")
	}

	logger.Info().Hex("address", receipt.ContractAddress.Bytes()).Msg("Contract address found")

	t.Log("Fib contract is at", receipt.ContractAddress.Hex())

	fib, err := arbostestcontracts.NewFibonacci(receipt.ContractAddress, l2Client)
	if err != nil {
		t.Fatal("connect fib failed", err)
	}

	// Wrap the Token contract instance into a session
	session := &arbostestcontracts.FibonacciSession{
		Contract: fib,
		CallOpts: bind.CallOpts{
			From:    auths[4].From,
			Pending: true,
		},
		TransactOpts: *auths[4],
	}

	fibsize := 15
	fibnum := 11

	tx, err = session.GenerateFib(big.NewInt(int64(fibsize)))
	if err != nil {
		t.Fatal("GenerateFib error", err)
	}
	receipt, err = waitForReceipt(
		l2Client,
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

	start := uint64(0)

Loop:
	for {
		select {
		case <-time.After(time.Second * 20):
			return
		default:
		}

		filter := &bind.FilterOpts{
			Start:   start,
			End:     nil,
			Context: ctx,
		}

		it, err := session.Contract.FilterTestEvent(filter)
		if err != nil {
			t.Fatalf("FilterTestEvent error %v", err)
			return
		}

		for it.Next() {
			if it.Event.Number.Cmp(big.NewInt(int64(fibsize))) != 0 {
				t.Error("test event had wrong number")
			}
			break Loop
		}
	}
}
