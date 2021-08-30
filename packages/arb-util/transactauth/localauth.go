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

package transactauth

import (
	"context"
	"math/big"
	"sync"

	"github.com/offchainlabs/arbitrum/packages/arb-util/arbtransaction"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type LocalTransactAuth struct {
	sync.Mutex
	auth   *bind.TransactOpts
	signer bind.SignerFn
	client ethutils.EthClient
}

func NewTransactAuthAdvanced(
	ctx context.Context,
	client ethutils.EthClient,
	auth *bind.TransactOpts,
	usePendingNonce bool,
) (TransactAuth, error) {
	err := getNonce(ctx, client, auth, usePendingNonce)
	if err != nil {
		return nil, err
	}

	return &LocalTransactAuth{
		auth:   auth,
		signer: auth.Signer,
		client: client,
	}, nil
}

func NewTransactAuth(
	ctx context.Context,
	client ethutils.EthClient,
	auth *bind.TransactOpts,
) (TransactAuth, error) {
	return NewTransactAuthAdvanced(ctx, client, auth, true)
}
func (ta *LocalTransactAuth) TransactionReceipt(ctx context.Context, tx *arbtransaction.ArbTransaction) (*types.Receipt, error) {
	return ta.client.TransactionReceipt(ctx, tx.Hash())
}

func (ta *LocalTransactAuth) NonceAt(ctx context.Context, account ethcommon.Address, blockNumber *big.Int) (uint64, error) {
	return ta.client.NonceAt(ctx, account, blockNumber)
}

func (ta *LocalTransactAuth) SendTransaction(ctx context.Context, tx *types.Transaction, replaceTxByHash string) (*arbtransaction.ArbTransaction, error) {
	err := ta.client.SendTransaction(ctx, tx)
	if err != nil {
		logger.Error().Err(err).Hex("data", tx.Data()).Msg("error sending transaction")
		return nil, err
	}

	logger.Debug().Hex("data", tx.Data()).Msg("sent transaction")
	arbTx := arbtransaction.NewArbTransaction(tx)
	return arbTx, nil
}

func (ta *LocalTransactAuth) Sign(addr ethcommon.Address, tx *types.Transaction) (*types.Transaction, error) {
	return ta.signer(addr, tx)
}

func (ta *LocalTransactAuth) GetAuth() *bind.TransactOpts {
	return ta.auth
}

func (ta *LocalTransactAuth) From() ethcommon.Address {
	return ta.auth.From
}
