/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

package cmdhelper

import (
	"context"
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

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupmanager"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupvalidator"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func GetKeystore(validatorFolder string, pass *string, flags *flag.FlagSet) (*bind.TransactOpts, error) {
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

func ValidateRollupChain(execName string, managerCreationFunc func(rollupAddress common.Address, client arbbridge.ArbAuthClient, contractFile string, dbPath string) (*rollupmanager.Manager, error)) error {
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
		return fmt.Errorf("usage: %v validate [--password=pass] [--rpc] [--blocktime=NumSeconds] [--gasprice==FloatInGwei] <validator_folder> <ethURL> <rollup_address>", execName)
	}

	common.SetDurationPerBlock(time.Duration(*blocktime) * time.Second)

	validatorFolder := validateCmd.Arg(0)
	ethURL := validateCmd.Arg(1)
	addressString := validateCmd.Arg(2)
	address := common.HexToAddress(addressString)

	auth, err := GetKeystore(validatorFolder, passphrase, validateCmd)
	if err != nil {
		return err
	}

	// Rollup creation
	gasPriceAsFloat := 1e9 * (*gasPrice)
	if gasPriceAsFloat < math.MaxInt64 {
		auth.GasPrice = big.NewInt(int64(gasPriceAsFloat))
	}
	ethclint, err := ethclient.Dial(ethURL)
	if err != nil {
		return err
	}
	client := ethbridge.NewEthAuthClient(ethclint, auth)

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

	manager, err := managerCreationFunc(address, client, contractFile, dbPath)

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
