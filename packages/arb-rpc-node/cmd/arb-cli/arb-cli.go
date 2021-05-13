package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type Config struct {
	client ethutils.EthClient
	auth   *bind.TransactOpts
}

var config *Config

func switchFees(enabled bool) error {
	arbOwner, err := arboscontracts.NewArbOwner(arbos.ARB_OWNER_ADDRESS, config.client)
	if err != nil {
		return err
	}
	tx, err := arbOwner.SetFeesEnabled(config.auth, enabled)
	if err != nil {
		return err
	}
	fmt.Println("Waiting for receipt")
	_, err = ethbridge.WaitForReceiptWithResults(context.Background(), config.client, config.auth.From, tx, "SetFeesEnabled")
	if err != nil {
		return err
	}
	fmt.Println("Transaction completed successfully")
	return nil
}

func setDefaultAggregator(agg ethcommon.Address) error {
	arbAggregator, err := arboscontracts.NewArbAggregator(arbos.ARB_AGGREGATOR_ADDRESS, config.client)
	if err != nil {
		return err
	}
	tx, err := arbAggregator.SetDefaultAggregator(config.auth, agg)
	fmt.Println("Waiting for receipt")
	_, err = ethbridge.WaitForReceiptWithResults(context.Background(), config.client, config.auth.From, tx, "SetDefaultAggregator")
	if err != nil {
		return err
	}
	fmt.Println("Transaction completed successfully")
	return nil
}

func estimateTransferGas() error {
	dest := common.RandAddress().ToEthAddress()
	msg := ethereum.CallMsg{
		From: config.auth.From,
		To:   &dest,
	}
	gas, err := config.client.EstimateGas(context.Background(), msg)
	if err != nil {
		return err
	}
	fmt.Println("Gas estimate:", gas)
	return nil
}

func handleCommand(fields []string) error {
	switch fields[0] {
	case "enable-fees":
		if len(fields) != 2 {
			return errors.New("Expected a true or false argument")
		}
		enabled, err := strconv.ParseBool(fields[1])
		if err != nil {
			return err
		}
		return switchFees(enabled)
	case "estimate-transfer-gas":
		return estimateTransferGas()
	case "set-default-agg":
		if len(fields) != 2 {
			return errors.New("Expected address argument")
		}
		agg := ethcommon.HexToAddress(fields[1])
		return setDefaultAggregator(agg)
	default:
		fmt.Println("Unknown command")
	}
	return nil
}

func executor(t string) {
	if t == "exit" {
		os.Exit(0)
	}
	fields := strings.Fields(t)
	err := handleCommand(fields)
	if err != nil {
		fmt.Println("Error running command", err)
	}
}

func completer(t prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "enable-fees"},
		{Text: "exit"},
	}
}

func run(ctx context.Context) error {
	if len(os.Args) != 3 {
		fmt.Println("Expected: arb-cli rpcurl privkey")
	}
	arbUrl := os.Args[1]
	privKeystr := os.Args[2]

	client, err := ethutils.NewRPCEthClient(arbUrl)
	if err != nil {
		return err
	}
	chainId, err := client.ChainID(ctx)
	if err != nil {
		return err
	}
	privKey, err := crypto.HexToECDSA(privKeystr)
	if err != nil {
		return err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainId)
	if err != nil {
		return err
	}
	fmt.Println("Sending from address", auth.From)

	config = &Config{
		client: client,
		auth:   auth,
	}

	p := prompt.New(
		executor,
		completer,
	)
	p.Run()
	return nil
}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Println("Error running app", err)
	}
}
