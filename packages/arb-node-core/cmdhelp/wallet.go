/*
* Copyright 2020-2021, Offchain Labs, Inc.
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

package cmdhelp

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arblog"
	"math"
	"math/big"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/pkg/errors"
)

var logger = arblog.Logger.With().Str("component", "cmdhelp").Logger()

func readPass() (string, error) {
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	passphrase := string(bytePassword)
	passphrase = strings.TrimSpace(passphrase)
	return passphrase, nil
}

// GetKeystore returns a transaction authorization based on an existing ethereum
// keystore located in validatorFolder/wallets or creates one if it does not
// exist. It accepts a password using the "password" command line argument or
// via an interactive prompt. It also sets the gas price of the auth via an
// optional "gasprice" argument.
func GetKeystore(
	config *configuration.Config,
	walletConfig *configuration.Wallet,
	chainId *big.Int,
	signerRequired bool,
) (*bind.TransactOpts, func([]byte) ([]byte, error), string, error) {
	var signer = func(data []byte) ([]byte, error) { return nil, errors.New("undefined signer") }
	var auth *bind.TransactOpts

	if len(walletConfig.Fireblocks.SSLKey) != 0 {
		if walletConfig.Local.OnlyCreateKey {
			return nil, nil, "using fireblocks keystore, remove --wallet.local.only-create-key to run normally", nil
		}
		fromAddress := ethcommon.HexToAddress(walletConfig.Fireblocks.SourceAddress)
		logger.Info().Hex("address", fromAddress.Bytes()).Msg("fireblocks enabled")
		auth = &bind.TransactOpts{
			From: fromAddress,
			Signer: func(address ethcommon.Address, tx *types.Transaction) (*types.Transaction, error) {
				if address != fromAddress {
					logger.Error().Hex("currentaddress", address.Bytes()).Hex("expectedaddress", fromAddress.Bytes()).Msg("incorrect from address provided")
					return nil, bind.ErrNotAuthorized
				}
				// Just return original unsigned transaction because fireblocks will handle signing
				return tx, nil
			},
		}

		if len(walletConfig.Fireblocks.FeedSigner.PrivateKey) != 0 {
			privateKey, err := crypto.HexToECDSA(walletConfig.Fireblocks.FeedSigner.PrivateKey)
			if err != nil {
				return nil, nil, "", errors.Wrap(err, "error loading feed private key")
			}

			publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
			if !ok {
				return nil, nil, "", errors.Wrap(err, "error generating public address of feed private key")
			}
			logger.
				Info().
				Hex("signer", crypto.PubkeyToAddress(*publicKeyECDSA).Bytes()).
				Msg("feed private key used as signer")
			signer = func(data []byte) ([]byte, error) {
				return crypto.Sign(data, privateKey)
			}
		} else if signerRequired {
			if len(walletConfig.Fireblocks.FeedSigner.Pathname) == 0 {
				return nil, nil, "", errors.New("missing feed signer private key")
			}
			ks, account, _, err := openKeystore("feed signer", walletConfig.Fireblocks.FeedSigner.Pathname, walletConfig.Fireblocks.FeedSigner.Password(), false)
			if err != nil {
				return nil, nil, "", err
			}

			logger.
				Info().
				Hex("signer", account.Address.Bytes()).
				Msg("feed signer wallet used as signer")
			signer = func(data []byte) ([]byte, error) {
				return ks.SignHash(*account, data)
			}
		}
	} else if len(walletConfig.Local.PrivateKey) != 0 {
		if walletConfig.Local.OnlyCreateKey {
			return nil, nil, "wallet key already exists, remove --wallet.local.only-create-key to run normally", nil
		}
		privateKey, err := crypto.HexToECDSA(walletConfig.Local.PrivateKey)
		if err != nil {
			return nil, nil, "", err
		}
		auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainId)
		if err != nil {
			return nil, nil, "", err
		}

		logger.
			Info().
			Hex("signer", auth.From.Bytes()).
			Msg("private key used as signer")
		signer = func(data []byte) ([]byte, error) {
			return crypto.Sign(data, privateKey)
		}
	} else {
		ks, account, newKeystoreCreated, err := openKeystore("account", walletConfig.Local.Pathname, walletConfig.Local.Password(), walletConfig.Local.OnlyCreateKey)
		if err != nil {
			return nil, nil, "", err
		}

		if newKeystoreCreated {
			return nil, nil, "wallet key created, backup key (" + walletConfig.Local.Pathname + ") and remove --wallet.local.only-create-key to start normally", nil
		}

		if walletConfig.Local.OnlyCreateKey {
			return nil, nil, "wallet key already created, backup key (" + walletConfig.Local.Pathname + ") and remove --wallet.local.only-create-key to run normally", nil
		}

		auth, err = bind.NewKeyStoreTransactorWithChainID(ks, *account, chainId)
		if err != nil {
			return nil, nil, "", err
		}

		logger.
			Info().
			Hex("signer", account.Address.Bytes()).
			Msg("wallet used as signer")
		signer = func(data []byte) ([]byte, error) {
			return ks.SignHash(*account, data)
		}
	}

	gasPriceAsFloat := 1e9 * config.GasPrice
	if gasPriceAsFloat < math.MaxInt64 && gasPriceAsFloat > 0 {
		auth.GasPrice = big.NewInt(int64(gasPriceAsFloat))
	}

	return auth, signer, "", nil
}

func openKeystore(description string, walletPath string, walletPassword *string, createNewKey bool) (*keystore.KeyStore, *accounts.Account, bool, error) {
	ks := keystore.NewKeyStore(
		walletPath,
		keystore.StandardScryptN,
		keystore.StandardScryptP,
	)
	logger.Info().
		Str("location", walletPath).
		Int("accounts", len(ks.Accounts())).
		Msg("loading wallet")

	creatingNew := len(ks.Accounts()) == 0
	if creatingNew && !createNewKey {
		return nil, nil, false, errors.New("No wallet exists, re-run with --wallet.local.only-create-key to create a wallet")
	}
	passOpt := walletPassword
	var password string
	if passOpt != nil {
		password = *passOpt
	} else {
		if creatingNew {
			fmt.Print("Enter new account password: ")
		} else {
			fmt.Print("Enter account password: ")
		}
		var err error
		password, err = readPass()
		if err != nil {
			return nil, nil, false, err
		}
	}

	var account accounts.Account
	if creatingNew {
		var err error
		account, err = ks.NewAccount(password)
		if err != nil {
			return nil, &accounts.Account{}, false, err
		}

		logger.Info().Hex("address", account.Address.Bytes()).Str("description", description).Msg("created new wallet")
	} else {
		account = ks.Accounts()[0]

		logger.Info().Hex("address", account.Address.Bytes()).Str("description", description).Msg("used existing wallet")
	}

	err := ks.Unlock(account, password)
	if err != nil {
		return nil, nil, false, err
	}

	return ks, &account, creatingNew, nil
}

const WalletArgsString = "[--wallet.password=pass] [--wallet.gasprice==FloatInGwei]"
