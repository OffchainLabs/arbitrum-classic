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
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var logger = log.With().Caller().Stack().Str("component", "ethbridge").Logger()

type TransactAuth interface {
	SendTransaction(ctx context.Context, tx *types.Transaction, replaceTxByHash string) (*ArbTransaction, error)
	TransactionReceipt(ctx context.Context, tx *ArbTransaction) (*types.Receipt, error)
	NonceAt(ctx context.Context, account ethcommon.Address, blockNumber *big.Int) (uint64, error)
	Sign(ethcommon.Address, *types.Transaction) (*types.Transaction, error)
	From() ethcommon.Address
	getAuth() *bind.TransactOpts
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

func makeContract(
	ctx context.Context,
	t TransactAuth,
	contractFunc func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error),
) (ethcommon.Address, *ArbTransaction, error) {
	auth := t.getAuth()

	addr, arbTx, err := makeContractImpl(ctx, t, auth, contractFunc)
	if err != nil {
		return ethcommon.Address{}, nil, err
	}

	// Transaction successful, increment nonce for next time
	auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
	return addr, arbTx, err
}

func makeContractCustomNonce(
	ctx context.Context,
	t TransactAuth,
	contractFunc func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error),
	customNonce *big.Int,
) (ethcommon.Address, *ArbTransaction, error) {
	auth := t.getAuth()
	origNonce := auth.Nonce
	defer func(auth *bind.TransactOpts) {
		auth.Nonce = origNonce
	}(auth)

	auth.Nonce = customNonce

	addr, arbTx, err := makeContractImpl(ctx, t, auth, contractFunc)
	if err != nil {
		return ethcommon.Address{}, nil, err
	}

	return addr, arbTx, err
}

func makeContractImpl(
	ctx context.Context,
	t TransactAuth,
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
	arbTx, err := t.SendTransaction(ctx, tx, "")
	if err != nil {
		logger.
			Error().
			Err(err).
			Str("nonce", auth.Nonce.String()).
			Hex("sender", auth.From.Bytes()).
			Hex("to", tx.To().Bytes()).
			Hex("data", tx.Data()).
			Str("nonce", auth.Nonce.String()).
			Msg("unable to send transaction")
		return addr, nil, err
	}

	logger.Info().Str("nonce", auth.Nonce.String()).Hex("sender", auth.From.Bytes()).Msg("transaction sent")

	return addr, arbTx, nil
}

func makeTx(
	ctx context.Context,
	t TransactAuth,
	txFunc func(auth *bind.TransactOpts) (*types.Transaction, error),
) (*ArbTransaction, error) {
	_, arbTx, err := makeContract(ctx, t, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
		tx, err := txFunc(auth)
		return ethcommon.BigToAddress(big.NewInt(0)), tx, nil, err
	})

	return arbTx, err
}

func makeTxCustomNonce(
	ctx context.Context,
	t TransactAuth,
	txFunc func(auth *bind.TransactOpts) (*types.Transaction, error),
	customNonce *big.Int,
) (*ArbTransaction, error) {
	_, arbTx, err := makeContractCustomNonce(ctx, t, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
		tx, err := txFunc(auth)
		return ethcommon.BigToAddress(big.NewInt(0)), tx, nil, err
	}, customNonce)

	return arbTx, err
}
