/*
 * Copyright 2021, Offchain Labs, Inc.
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

package ethbridge

import (
	"context"
	"crypto/rsa"
	"encoding/hex"
	"math/big"
	"strings"
	"sync"

	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/fireblocks"
	"github.com/offchainlabs/arbitrum/packages/arb-util/fireblocks/accounttype"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var logger = log.With().Caller().Stack().Str("component", "ethbridge").Logger()

const (
	smallNonceRepeatCount = 100
	smallNonceError       = "Try increasing the gas price or incrementing the nonce."
)

type TransactAuth struct {
	sync.Mutex
	auth   *bind.TransactOpts
	SendTx func(ctx context.Context, tx *types.Transaction) (*ArbTransaction, error)
	Signer bind.SignerFn
}

func getNonce(ctx context.Context, client ethutils.EthClient, auth *bind.TransactOpts, usePendingNonce bool) error {
	if auth.Nonce == nil {
		var nonce uint64
		var err error
		if usePendingNonce {
			nonce, err = client.PendingNonceAt(ctx, auth.From)
		} else {
			blockNum := big.NewInt(int64(rpc.LatestBlockNumber))
			nonce, err = client.NonceAt(ctx, auth.From, blockNum)
		}
		if err != nil {
			return errors.Wrap(err, "failed to get nonce")
		}
		auth.Nonce = new(big.Int).SetUint64(nonce)
	}

	return nil
}

func NewTransactAuthAdvanced(
	ctx context.Context,
	client ethutils.EthClient,
	auth *bind.TransactOpts,
	config *configuration.Config,
	walletConfig *configuration.Wallet,
	usePendingNonce bool,
) (*TransactAuth, error) {
	err := getNonce(ctx, client, auth, usePendingNonce)
	if err != nil {
		return nil, err
	}

	// Send transaction normally
	sendTx := func(ctx context.Context, tx *types.Transaction) (*ArbTransaction, error) {
		err := client.SendTransaction(ctx, tx)
		if err != nil {
			logger.Error().Err(err).Hex("data", tx.Data()).Msg("error sending transaction")
			return nil, err
		}

		logger.Debug().Hex("data", tx.Data()).Msg("sent transaction")
		arbTx := NewArbTransaction(tx)
		return arbTx, nil
	}

	return &TransactAuth{
		auth:   auth,
		SendTx: sendTx,
		Signer: auth.Signer,
	}, nil
}

func NewFireblocksTransactAuthAdvanced(
	ctx context.Context,
	client ethutils.EthClient,
	auth *bind.TransactOpts,
	config *configuration.Config,
	walletConfig *configuration.Wallet,
	usePendingNonce bool,
) (*TransactAuth, *fireblocks.Fireblocks, error) {
	err := getNonce(ctx, client, auth, usePendingNonce)
	if err != nil {
		return nil, nil, err
	}

	var signKey *rsa.PrivateKey
	if len(walletConfig.FireblocksSSLKeyPassword) != 0 {
		signKey, err = jwt.ParseRSAPrivateKeyFromPEMWithPassword([]byte(walletConfig.FireblocksSSLKey), walletConfig.FireblocksSSLKeyPassword)
	} else {
		signKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(walletConfig.FireblocksSSLKey))
	}
	if err != nil {
		return nil, nil, errors.Wrap(err, "problem with fireblocks privatekey")
	}
	sourceType, err := accounttype.New(config.Fireblocks.SourceType)
	if err != nil {
		return nil, nil, errors.Wrap(err, "problem with fireblocks source-type")
	}
	fb := fireblocks.New(config.Fireblocks.AssetId, config.Fireblocks.BaseURL, *sourceType, config.Fireblocks.SourceId, config.Fireblocks.APIKey, signKey)
	sendTx := func(ctx context.Context, tx *types.Transaction) (*ArbTransaction, error) {
		return fireblocksSendTransaction(ctx, fb, tx)
	}

	return &TransactAuth{
		auth:   auth,
		SendTx: sendTx,
		Signer: func(addr ethcommon.Address, tx *types.Transaction) (*types.Transaction, error) {
			// Fireblocks handles signing
			return tx, nil
		},
	}, fb, nil
}

func fireblocksSendTransaction(ctx context.Context, fb *fireblocks.Fireblocks, tx *types.Transaction) (*ArbTransaction, error) {
	txResponse, err := fb.CreateContractCall(
		accounttype.OneTimeAddress,
		tx.To().Hex(),
		"",
		tx.Value(),
		tx.GasTipCap(),
		tx.GasFeeCap(),
		ethcommon.Bytes2Hex(tx.Data()),
	)
	if err != nil {
		return nil, err
	}
	if fb.IsTransactionStatusFailed(txResponse.Status) {
		logger.
			Error().
			Hex("data", tx.Data()).
			Str("id", txResponse.Id).
			Str("status", txResponse.Status).
			Msg("fireblocks transaction failed")
		return nil, errors.New("fireblocks transaction failed")
	}
	logger.Debug().Hex("data", tx.Data()).Msg("sent transaction")

	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("ctx done")
		default:
		}

		details, err := fb.GetTransaction(txResponse.Id)
		if err != nil && strings.Contains(err.Error(), "not found") {
			logger.
				Warn().
				Err(err).
				Hex("To", tx.To().Bytes()).
				Str("id", txResponse.Id).
				Msg("fireblocks transaction not found")
			return nil, errors.Wrap(err, "fireblocks transaction not found")
		}

		if fb.IsTransactionStatusFailed(details.Status) {
			logger.
				Error().
				Err(err).
				Str("To", details.DestinationAddress).
				Str("From", details.SourceAddress).
				Str("status", details.Status).
				Str("txhash", details.TxHash).
				Msg("fireblocks transaction failed")
			return nil, errors.Wrapf(err, "fireblocks transaction failed: %s", details.Status)
		}

		if len(details.TxHash) > 0 {
			signer := NewDummyHashSigner(big.NewInt(99))
			txHash, err := hex.DecodeString(details.TxHash)
			if err != nil || len(txHash) < 32 {
				return nil, errors.Wrap(err, "error decoding txHash from fireblocks response")
			}
			tx, err := tx.WithSignature(signer, txHash)
			if err != nil {
				return nil, errors.Wrap(err, "error encoding fireblocks transaction hash into signature")
			}
			arbTx := NewFireblocksArbTransaction(tx, ethcommon.BytesToHash(txHash))
			return arbTx, nil
		}

		// Hash not returned, keep trying
	}
}

func NewTransactAuth(
	ctx context.Context,
	client ethutils.EthClient,
	auth *bind.TransactOpts,
	config *configuration.Config,
	walletConfig *configuration.Wallet,
) (*TransactAuth, error) {
	return NewTransactAuthAdvanced(ctx, client, auth, config, walletConfig, true)
}

func NewFireblocksTransactAuth(
	ctx context.Context,
	client ethutils.EthClient,
	auth *bind.TransactOpts,
	config *configuration.Config,
	walletConfig *configuration.Wallet,
) (*TransactAuth, *fireblocks.Fireblocks, error) {
	return NewFireblocksTransactAuthAdvanced(ctx, client, auth, config, walletConfig, true)
}

func (t *TransactAuth) makeContract(ctx context.Context, contractFunc func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error)) (ethcommon.Address, *ArbTransaction, error) {
	auth := t.getAuth(ctx)

	// Form transaction without sending it
	auth.NoSend = true
	addr, tx, _, err := contractFunc(auth)
	err = errors.WithStack(err)
	if err != nil {
		// Error occurred before sending, so don't need retry logic below
		logger.Error().Err(err).Msg("error forming transaction")
		return addr, nil, err
	}

	// Actually send transaction
	arbTx, err := t.SendTx(ctx, tx)

	if err != nil {
		logger.Error().Err(err).Str("nonce", auth.Nonce.String()).Hex("sender", t.auth.From.Bytes()).Hex("to", arbTx.To().Bytes()).Hex("data", arbTx.Data()).Str("nonce", auth.Nonce.String()).Msg("unable to send transaction")
		return addr, nil, err
	}

	logger.Info().Str("nonce", auth.Nonce.String()).Hex("sender", t.auth.From.Bytes()).Msg("transaction sent")

	// Transaction successful, increment nonce for next time
	t.auth.Nonce = t.auth.Nonce.Add(t.auth.Nonce, big.NewInt(1))
	return addr, arbTx, err
}

func (t *TransactAuth) makeTx(ctx context.Context, txFunc func(auth *bind.TransactOpts) (*types.Transaction, error)) (*ArbTransaction, error) {
	_, arbTx, err := t.makeContract(ctx, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
		tx, err := txFunc(auth)
		return ethcommon.BigToAddress(big.NewInt(0)), tx, nil, err
	})

	return arbTx, err
}

func (t *TransactAuth) getAuth(ctx context.Context) *bind.TransactOpts {
	opts := *t.auth
	opts.Context = ctx
	return &opts
}
