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
	"math/big"
	"sync"
	"time"

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
	auth               *bind.TransactOpts
	SendTx             func(ctx context.Context, tx *types.Transaction, replaceTxByHash string) (*ArbTransaction, error)
	Signer             bind.SignerFn
	transactionReceipt func(ctx context.Context, tx *ArbTransaction) (*types.Receipt, error)
	nonceAt            func(ctx context.Context, account ethcommon.Address, blockNumber *big.Int) (uint64, error)
}

func (ta *TransactAuth) TransactionReceipt(ctx context.Context, tx *ArbTransaction) (*types.Receipt, error) {
	return ta.transactionReceipt(ctx, tx)
}

func (ta *TransactAuth) NonceAt(ctx context.Context, account ethcommon.Address, blockNumber *big.Int) (uint64, error) {
	return ta.nonceAt(ctx, account, blockNumber)
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
	usePendingNonce bool,
) (*TransactAuth, error) {
	err := getNonce(ctx, client, auth, usePendingNonce)
	if err != nil {
		return nil, err
	}

	// Send transaction normally
	sendTx := func(ctx context.Context, tx *types.Transaction, replaceTxByHash string) (*ArbTransaction, error) {
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
		transactionReceipt: func(ctx context.Context, tx *ArbTransaction) (*types.Receipt, error) {
			return client.TransactionReceipt(ctx, tx.Hash())
		},
		nonceAt: client.NonceAt,
	}, nil
}

func NewFireblocksTransactAuthAdvanced(
	ctx context.Context,
	client ethutils.EthClient,
	auth *bind.TransactOpts,
	walletConfig *configuration.Wallet,
	usePendingNonce bool,
) (*TransactAuth, *fireblocks.Fireblocks, error) {
	err := getNonce(ctx, client, auth, usePendingNonce)
	if err != nil {
		return nil, nil, err
	}

	var signKey *rsa.PrivateKey
	if len(walletConfig.Fireblocks.SSLKeyPassword) != 0 {
		signKey, err = jwt.ParseRSAPrivateKeyFromPEMWithPassword([]byte(walletConfig.Fireblocks.SSLKey), walletConfig.Fireblocks.SSLKeyPassword)
	} else {
		signKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(walletConfig.Fireblocks.SSLKey))
	}
	if err != nil {
		return nil, nil, errors.Wrap(err, "problem with fireblocks privatekey")
	}
	sourceType, err := accounttype.New(walletConfig.Fireblocks.SourceType)
	if err != nil {
		return nil, nil, errors.Wrap(err, "problem with fireblocks source-type")
	}
	fb := fireblocks.New(
		walletConfig.Fireblocks.AssetId,
		walletConfig.Fireblocks.BaseURL,
		*sourceType,
		walletConfig.Fireblocks.SourceId,
		walletConfig.Fireblocks.APIKey,
		signKey,
	)
	sendTx := func(ctx context.Context, tx *types.Transaction, replaceTxByHash string) (*ArbTransaction, error) {
		return fireblocksSendTransaction(ctx, fb, tx, replaceTxByHash)
	}
	transactionReceipt := func(ctx context.Context, tx *ArbTransaction) (*types.Receipt, error) {
		details, err := fb.GetTransaction(tx.Id())
		if err != nil {
			logger.
				Warn().
				Err(err).
				Str("hash", tx.Hash().String()).
				Msg("error getting fireblocks transaction for receipt")
			return nil, errors.Wrapf(err, "error getting fireblocks transaction for receipt: %s", details.Status)
		}
		if fb.IsTransactionStatusFailed(details.Status) {
			logger.
				Error().
				Err(err).
				Str("To", details.DestinationAddress).
				Str("From", details.SourceAddress).
				Str("status", details.Status).
				Str("txhash", details.TxHash).
				Msg("fireblocks transaction failed when getting receipt")
			return nil, errors.Wrapf(err, "fireblocks transaction failed when getting receipt: %s", details.Status)
		}

		return client.TransactionReceipt(ctx, tx.Hash())
	}

	transactAuth := &TransactAuth{
		auth:   auth,
		SendTx: sendTx,
		Signer: func(addr ethcommon.Address, tx *types.Transaction) (*types.Transaction, error) {
			// Fireblocks handles signing, so nothing to do here
			return tx, nil
		},
		transactionReceipt: transactionReceipt,
		nonceAt:            client.NonceAt,
	}

	// Handle any pending transactions left from last time
	/* TODO
	err = waitForPendingTransactions(ctx, client, transactAuth, fb)
	if err != nil {
		return nil, nil, err
	}
	*/

	return transactAuth, fb, nil
}

func waitForPendingTransactions(
	ctx context.Context,
	client ethutils.EthClient,
	transactAuth *TransactAuth,
	fb *fireblocks.Fireblocks,
) error {
	for {
		pendingTx, err := fb.ListPendingTransactions()
		if err != nil {
			logger.Error().Err(err).Msg("error listing pending transactions")
			return err
		}

		if len(*pendingTx) == 0 {
			break
		}

		// Get updated fees to use
		networkFees, err := fb.EstimateNetworkFees()
		if err != nil {
			return err
		}
		// TODO
		_ = networkFees

		for _, details := range *pendingTx {
			if details.Status == fireblocks.Broadcasting {
				logger.
					Info().
					Str("id", details.Id).
					Str("status", details.Status).
					Str("destination", details.DestinationAddress).
					Msg("retrying pending fireblocks transaction")
				// Existing transaction is stuck
				destinationAddress := ethcommon.HexToAddress(details.DestinationAddress)
				baseTx := &types.DynamicFeeTx{
					ChainID:   big.NewInt(0), // Fireblocks ignore chain id
					Nonce:     0,             // Fireblocks ignores nonce
					GasTipCap: big.NewInt(0),
					GasFeeCap: big.NewInt(0),
					Gas:       0,
					To:        &destinationAddress,
					Value:     big.NewInt(details.Amount),
					Data:      []byte(details.ExtraParameters.ContractCallData),
				}
				rawTx := types.NewTx(baseTx)
				arbTx, err := NewFireblocksArbTransaction(rawTx, &details)
				if err != nil {
					logger.Error().Err(err).Msg("unable to wait for pending transactions")
					return err
				}
				_, err = WaitForReceiptWithResultsAndReplaceByFee(
					ctx,
					client,
					transactAuth.auth.From,
					arbTx,
					"CreateWallet",
					transactAuth,
					transactAuth,
				)
				if err != nil {
					return err
				}
			} else {
				logger.
					Info().
					Str("id", details.Id).
					Str("status", details.Status).
					Str("destination", details.DestinationAddress).
					Msg("waiting on pending fireblocks transaction")
			}
		}

		select {
		case <-ctx.Done():
			return nil
		case <-time.After(5 * time.Second):
		}
	}

	return nil
}

func fireblocksSendTransaction(ctx context.Context, fb *fireblocks.Fireblocks, tx *types.Transaction, replaceTxByHash string) (*ArbTransaction, error) {
	txResponse, err := fb.CreateContractCall(
		accounttype.OneTimeAddress,
		tx.To().Hex(),
		"",
		tx.Value(),
		tx.GasTipCap(),
		tx.GasFeeCap(),
		replaceTxByHash,
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
		case <-time.After(2 * time.Second):
		}

		details, err := fb.GetTransaction(txResponse.Id)
		if err != nil {
			logger.
				Warn().
				Err(err).
				Hex("to", tx.To().Bytes()).
				Str("id", txResponse.Id).
				Str("status", details.Status).
				Msg("error getting fireblocks transaction")
			return nil, errors.Wrapf(err, "error getting fireblocks transaction: %s", details.Status)
		}

		if fb.IsTransactionStatusFailed(details.Status) {
			logger.
				Error().
				Err(err).
				Str("To", details.DestinationAddress).
				Str("From", details.SourceAddress).
				Str("status", details.Status).
				Str("txhash", details.TxHash).
				Str("replaceTxByHash", replaceTxByHash).
				Msg("fireblocks transaction failed")
			return nil, errors.Wrapf(err, "fireblocks transaction failed: %s", details.Status)
		}

		if len(details.TxHash) > 0 {
			return NewFireblocksArbTransaction(tx, details)
		}

		// Hash not returned, keep trying
	}
}

func NewTransactAuth(
	ctx context.Context,
	client ethutils.EthClient,
	auth *bind.TransactOpts,
) (*TransactAuth, error) {
	return NewTransactAuthAdvanced(ctx, client, auth, true)
}

func NewFireblocksTransactAuth(
	ctx context.Context,
	client ethutils.EthClient,
	auth *bind.TransactOpts,
	walletConfig *configuration.Wallet,
) (*TransactAuth, *fireblocks.Fireblocks, error) {
	return NewFireblocksTransactAuthAdvanced(ctx, client, auth, walletConfig, true)
}

func (t *TransactAuth) makeContract(
	ctx context.Context,
	contractFunc func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error),
) (ethcommon.Address, *ArbTransaction, error) {
	auth := t.getAuth(ctx)

	addr, arbTx, err := t.makeContractImpl(ctx, auth, contractFunc)
	if err != nil {
		return ethcommon.Address{}, nil, err
	}

	// Transaction successful, increment nonce for next time
	t.auth.Nonce = t.auth.Nonce.Add(t.auth.Nonce, big.NewInt(1))
	return addr, arbTx, err
}

func (t *TransactAuth) makeContractCustomNonce(
	ctx context.Context,
	contractFunc func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error),
	customNonce *big.Int,
) (ethcommon.Address, *ArbTransaction, error) {
	auth := t.getAuth(ctx)
	origNonce := auth.Nonce
	defer func(auth *bind.TransactOpts) {
		auth.Nonce = origNonce
	}(auth)

	auth.Nonce = customNonce

	addr, arbTx, err := t.makeContractImpl(ctx, auth, contractFunc)
	if err != nil {
		return ethcommon.Address{}, nil, err
	}

	return addr, arbTx, err
}

func (t *TransactAuth) makeContractImpl(
	ctx context.Context,
	auth *bind.TransactOpts,
	contractFunc func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error),
) (ethcommon.Address, *ArbTransaction, error) {
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
	arbTx, err := t.SendTx(ctx, tx, "")
	if err != nil {
		logger.
			Error().
			Err(err).
			Str("nonce", auth.Nonce.String()).
			Hex("sender", t.auth.From.Bytes()).
			Hex("to", tx.To().Bytes()).
			Hex("data", tx.Data()).
			Str("nonce", auth.Nonce.String()).
			Msg("unable to send transaction")
		return addr, nil, err
	}

	logger.Info().Str("nonce", auth.Nonce.String()).Hex("sender", t.auth.From.Bytes()).Msg("transaction sent")

	return addr, arbTx, nil
}

func (t *TransactAuth) makeTx(
	ctx context.Context,
	txFunc func(auth *bind.TransactOpts) (*types.Transaction, error),
) (*ArbTransaction, error) {
	_, arbTx, err := t.makeContract(ctx, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
		tx, err := txFunc(auth)
		return ethcommon.BigToAddress(big.NewInt(0)), tx, nil, err
	})

	return arbTx, err
}

func (t *TransactAuth) makeTxCustomNonce(
	ctx context.Context,
	txFunc func(auth *bind.TransactOpts) (*types.Transaction, error),
	customNonce *big.Int,
) (*ArbTransaction, error) {
	_, arbTx, err := t.makeContractCustomNonce(ctx, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
		tx, err := txFunc(auth)
		return ethcommon.BigToAddress(big.NewInt(0)), tx, nil, err
	}, customNonce)

	return arbTx, err
}

func (t *TransactAuth) getAuth(ctx context.Context) *bind.TransactOpts {
	opts := *t.auth
	opts.Context = ctx
	return &opts
}
