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
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/pkg/errors"
)

var logger = log.With().Caller().Stack().Str("component", "cmdhelp").Logger()

// GetKeystore returns a transaction authorization based on an existing ethereum
// keystore located in validatorFolder/wallets or creates one if it does not
// exist. It accepts a password using the "password" command line argument or
// via an interactive prompt. It also sets the gas price of the auth via an
// optional "gasprice" arguement.
func GetKeystore(
	config *configuration.Config,
	walletConfig *configuration.Wallet,
	chainId *big.Int,
	signerRequired bool,
) (*bind.TransactOpts, func([]byte) ([]byte, error), error) {
	var signer = func(data []byte) ([]byte, error) { return nil, errors.New("undefined signer") }
	var auth *bind.TransactOpts

	if len(walletConfig.Fireblocks.SSLKey) != 0 {
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
			privateKey, err := crypto.HexToECDSA(config.Wallet.Local.PrivateKey)
			if err != nil {
				return nil, nil, errors.Wrap(err, "error loading feed private key")
			}

			publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
			if !ok {
				return nil, nil, errors.Wrap(err, "error generating public address of feed private key")
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
				return nil, nil, errors.New("missing feed signer private key")
			}
			ks, account, err := openKeystore("feed signer", walletConfig.Fireblocks.FeedSigner.Pathname, walletConfig.Fireblocks.FeedSigner.Password)
			if err != nil {
				return nil, nil, err
			}

			logger.
				Info().
				Hex("signer", account.Address.Bytes()).
				Msg("feed signer wallet used as signer")
			signer = func(data []byte) ([]byte, error) {
				return ks.SignHash(account, data)
			}
		}
	} else if len(config.Wallet.Local.PrivateKey) != 0 {
		privateKey, err := crypto.HexToECDSA(config.Wallet.Local.PrivateKey)
		if err != nil {
			return nil, nil, err
		}
		auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainId)
		if err != nil {
			return nil, nil, err
		}

		logger.
			Info().
			Hex("signer", auth.From.Bytes()).
			Msg("private key used as signer")
		signer = func(data []byte) ([]byte, error) {
			return crypto.Sign(data, privateKey)
		}
	} else {
		ks, account, err := openKeystore("account", walletConfig.Local.Pathname, walletConfig.Local.Password)
		if err != nil {
			return nil, nil, err
		}

		auth, err = bind.NewKeyStoreTransactorWithChainID(ks, account, chainId)
		if err != nil {
			return nil, nil, err
		}

		logger.
			Info().
			Hex("signer", account.Address.Bytes()).
			Msg("wallet used as signer")
		signer = func(data []byte) ([]byte, error) {
			return ks.SignHash(account, data)
		}
	}

	gasPriceAsFloat := 1e9 * config.GasPrice
	if gasPriceAsFloat < math.MaxInt64 && gasPriceAsFloat > 0 {
		auth.GasPrice = big.NewInt(int64(gasPriceAsFloat))
	}

	return auth, signer, nil
}

func openKeystore(description string, walletPath string, walletPassword string) (*keystore.KeyStore, accounts.Account, error) {
	ks := keystore.NewKeyStore(
		walletPath,
		keystore.StandardScryptN,
		keystore.StandardScryptP,
	)

	var account accounts.Account
	if len(ks.Accounts()) > 0 {
		account = ks.Accounts()[0]
	}

	if ks.Unlock(account, walletPassword) != nil {
		if len(walletPassword) == 0 {
			if len(ks.Accounts()) == 0 {
				// Wallet doesn't exist and no password provided
				fmt.Printf("Enter new %s password: ", description)
			} else {
				// Wallet exists and no password provided
				fmt.Printf("Enter %s password: ", description)
			}

			bytePassword, err := terminal.ReadPassword(syscall.Stdin)
			if err != nil {
				return nil, accounts.Account{}, err
			}
			passphrase := string(bytePassword)

			walletPassword = strings.TrimSpace(passphrase)
		}

		if len(ks.Accounts()) == 0 {
			var err error
			account, err = ks.NewAccount(walletPassword)
			if err != nil {
				return nil, accounts.Account{}, err
			}
		}
		err := ks.Unlock(account, walletPassword)
		if err != nil {
			return nil, accounts.Account{}, err
		}

		logger.Info().Hex("address", account.Address.Bytes()).Str("description", description).Msg("created new wallet")
	}
	return ks, account, nil
}

const WalletArgsString = "[--wallet.password=pass] [--wallet.gasprice==FloatInGwei]"
