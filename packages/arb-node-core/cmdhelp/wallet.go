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
	"fmt"
	"math"
	"math/big"
	"path/filepath"
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
) (*bind.TransactOpts, func([]byte) ([]byte, error), error) {
	var account accounts.Account
	var signer = func(data []byte) ([]byte, error) { return nil, errors.New("undefined signer") }
	var auth *bind.TransactOpts

	if len(config.Wallet.FeedPrivateKey) != 0 {
		privateKey, err := crypto.HexToECDSA(config.Wallet.PrivateKey)
		if err != nil {
			return nil, nil, errors.Wrap(err, "error loading feed private key")
		}
		auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainId)
		if err != nil {
			return nil, nil, errors.Wrap(err, "error creating feed auth object")
		}

		signer = func(data []byte) ([]byte, error) {
			return crypto.Sign(data, privateKey)
		}
	}

	if len(walletConfig.FireblocksSSLKey) != 0 {
		fromAddress := ethcommon.HexToAddress(config.Fireblocks.SourceAddress)
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

		if len(config.Wallet.FeedPrivateKey) == 0 {
			// Separate key for signing feed not provided
			signer = func(data []byte) ([]byte, error) {
				// Fireblocks cannot sign arbitrary data
				return make([]byte, 32), nil
			}
		}
	} else if len(config.Wallet.PrivateKey) != 0 {
		privateKey, err := crypto.HexToECDSA(config.Wallet.PrivateKey)
		if err != nil {
			return nil, nil, err
		}
		auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainId)
		if err != nil {
			return nil, nil, err
		}

		signer = func(data []byte) ([]byte, error) {
			return crypto.Sign(data, privateKey)
		}
	} else {
		ks := keystore.NewKeyStore(
			filepath.Join(config.Persistent.Chain, "wallets"),
			keystore.StandardScryptN,
			keystore.StandardScryptP,
		)

		if len(ks.Accounts()) > 0 {
			account = ks.Accounts()[0]
		}

		if ks.Unlock(account, walletConfig.Password) != nil {
			if len(walletConfig.Password) == 0 {
				// Wallet doesn't exist and no password provided
				if len(ks.Accounts()) == 0 {
					fmt.Print("Enter new account password: ")
				} else {
					fmt.Print("Enter account password: ")
				}

				bytePassword, err := terminal.ReadPassword(syscall.Stdin)
				if err != nil {
					return nil, nil, err
				}
				passphrase := string(bytePassword)

				walletConfig.Password = strings.TrimSpace(passphrase)
			}

			if len(ks.Accounts()) == 0 {
				var err error
				account, err = ks.NewAccount(walletConfig.Password)
				if err != nil {
					return nil, nil, err
				}
			}
			err := ks.Unlock(account, walletConfig.Password)
			if err != nil {
				return nil, nil, err
			}

			logger.Info().Hex("address", account.Address.Bytes()).Msg("created new wallet")
		}

		if len(config.Wallet.FeedPrivateKey) == 0 {
			// Separate key for signing feed not provided
			signer = func(data []byte) ([]byte, error) {
				return ks.SignHash(account, data)
			}
		}

		var err error
		auth, err = bind.NewKeyStoreTransactorWithChainID(ks, account, chainId)
		if err != nil {
			return nil, nil, err
		}
	}

	gasPriceAsFloat := 1e9 * config.GasPrice
	if gasPriceAsFloat < math.MaxInt64 && gasPriceAsFloat > 0 {
		auth.GasPrice = big.NewInt(int64(gasPriceAsFloat))
	}

	return auth, signer, nil
}

const WalletArgsString = "[--wallet.password=pass] [--wallet.gasprice==FloatInGwei]"
