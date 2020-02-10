/*
 * Copyright 2019, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	errors2 "github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/ethereum/go-ethereum/accounts"

	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupmanager"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupvalidator"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"
)

func main() {
	// Check number of args
	flag.Parse()
	switch os.Args[1] {
	case "create":
		if err := createRollupChain(); err != nil {
			log.Fatal(err)
		}
	case "validate":
		if err := validateRollupChain(); err != nil {
			log.Fatal(err)
		}
	default:
	}
}

func getKeystore(validatorFolder string, pass *string, flags *flag.FlagSet) (*bind.TransactOpts, error) {
	ks := keystore.NewKeyStore(filepath.Join(validatorFolder, "wallets"), keystore.StandardScryptN, keystore.StandardScryptP)

	found := false
	flags.Visit(func(f *flag.Flag) {
		if f.Name == "password" {
			found = true
		}
	})

	var passphrase string
	if !found {
		if len(ks.Accounts()) == 0 {
			fmt.Print("Enter new account password: ")
		} else {
			fmt.Print("Enter account password: ")
		}

		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return nil, err
		}
		passphrase = string(bytePassword)

		passphrase = strings.TrimSpace(passphrase)
	} else {
		passphrase = *pass
	}

	var account accounts.Account
	if len(ks.Accounts()) == 0 {
		var err error
		account, err = ks.NewAccount(passphrase)
		if err != nil {
			return nil, err
		}
	} else {
		account = ks.Accounts()[0]
	}
	err := ks.Unlock(account, passphrase)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyStoreTransactor(ks, account)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func createRollupChain() error {
	createCmd := flag.NewFlagSet("validate", flag.ExitOnError)
	passphrase := createCmd.String("password", "", "password=pass")
	gasPrice := createCmd.Float64("gasprice", 4.5, "gasprice=FloatInGwei")
	err := createCmd.Parse(os.Args[2:])
	if err != nil {
		return err
	}

	if createCmd.NArg() != 3 {
		return errors.New("usage: rollupServer create [--password=pass] [--gasprice==FloatInGwei] <validator_folder> <ethURL> <factoryAddress>")
	}

	validatorFolder := createCmd.Arg(0)
	ethURL := createCmd.Arg(1)
	addressString := createCmd.Arg(2)
	factoryAddress := common.HexToAddress(addressString)
	contractFile := filepath.Join(validatorFolder, "contract.ao")

	// 1) Compiled Arbitrum bytecode
	mach, err := loader.LoadMachineFromFile(contractFile, true, "cpp")
	if err != nil {
		return errors2.Wrap(err, "loader error")
	}

	auth, err := getKeystore(validatorFolder, passphrase, createCmd)
	if err != nil {
		return err
	}
	gasPriceAsFloat := 1e9 * (*gasPrice)
	if gasPriceAsFloat < math.MaxInt64 {
		auth.GasPrice = big.NewInt(int64(gasPriceAsFloat))
	}

	// Rollup creation
	client, err := ethbridge.NewEthAuthClient(ethURL, auth)
	if err != nil {
		return err
	}

	if err := arbbridge.WaitForNonZeroBalance(context.Background(), client, common.NewAddressFromEth(auth.From)); err != nil {
		return err
	}

	factory, err := client.NewArbFactory(factoryAddress)
	if err != nil {
		return err
	}

	address, err := factory.CreateRollup(
		context.Background(),
		mach.Hash(),
		rollup.DefaultChainParams(),
		common.Address{},
	)
	if err != nil {
		return err
	}
	fmt.Println(address.Hex())
	return nil
}

func validateRollupChain() error {
	// Check number of args

	validateCmd := flag.NewFlagSet("validate", flag.ExitOnError)
	passphrase := validateCmd.String("password", "", "password=pass")
	rpcEnable := validateCmd.Bool("rpc", false, "rpc")
	blocktime := validateCmd.Int64("blocktime", 2, "blocktime=NumSeconds")
	gasPrice := validateCmd.Float64("gasprice", 4.5, "gasprice=FloatInGwei")
	err := validateCmd.Parse(os.Args[2:])
	if err != nil {
		return err
	}

	if validateCmd.NArg() != 3 {
		return errors.New("usage: rollupServer validate [--password=pass] [--rpc] [--blocktime=NumSeconds] [--gasprice==FloatInGwei] <validator_folder> <ethURL> <rollup_address>")
	}

	common.SetDurationPerBlock(time.Duration(*blocktime) * time.Second)

	validatorFolder := validateCmd.Arg(0)
	ethURL := validateCmd.Arg(1)
	addressString := validateCmd.Arg(2)
	address := common.HexToAddress(addressString)

	auth, err := getKeystore(validatorFolder, passphrase, validateCmd)
	if err != nil {
		return err
	}

	// Rollup creation
	gasPriceAsFloat := 1e9 * (*gasPrice)
	if gasPriceAsFloat < math.MaxInt64 {
		auth.GasPrice = big.NewInt(int64(gasPriceAsFloat))
	}
	client, err := ethbridge.NewEthAuthClient(ethURL, auth)
	if err != nil {
		return err
	}

	if err := arbbridge.WaitForNonZeroBalance(context.Background(), client, common.NewAddressFromEth(auth.From)); err != nil {
		return err
	}

	rollupActor, err := client.NewRollup(address)
	if err != nil {
		return err
	}

	validatorListener := rollup.NewValidatorChainListener(context.Background(), address, rollupActor)
	err = validatorListener.AddStaker(client)
	if err != nil {
		return err
	}

	contractFile := filepath.Join(validatorFolder, "contract.ao")
	dbPath := filepath.Join(validatorFolder, "checkpoint_db")
	manager, err := rollupmanager.CreateManager(address, client, contractFile, dbPath)

	if err != nil {
		return err
	}
	manager.AddListener(&rollup.AnnouncerListener{})
	manager.AddListener(validatorListener)

	if *rpcEnable {
		if err := rollupvalidator.LaunchRPC(manager, "1235"); err != nil {
			log.Fatal(err)
		}
	} else {
		wait := make(chan bool)
		<-wait
	}
	return nil
}
