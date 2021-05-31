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

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/pkg/errors"
)

type BuilderBackend struct {
	transactions []*types.Transaction
	builderAuth  *bind.TransactOpts
	realSender   common.Address
	wallet       common.Address

	realClient ethutils.EthClient
}

func NewBuilderBackend(wallet *ValidatorWallet) (*BuilderBackend, error) {
	randKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return &BuilderBackend{
		builderAuth: bind.NewKeyedTransactor(randKey),
		realSender:  wallet.From().ToEthAddress(),
		wallet:      wallet.Address().ToEthAddress(),
		realClient:  wallet.client,
	}, nil
}

func (b *BuilderBackend) TransactionCount() int {
	return len(b.transactions)
}

func (b *BuilderBackend) ClearTransactions() {
	b.transactions = nil
}

func (b *BuilderBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return []byte{1}, nil
}

func (b *BuilderBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	panic("implement me")
}

func (b *BuilderBackend) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return []byte{1}, nil
}

func (b *BuilderBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return 0, nil
}

func (b *BuilderBackend) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return 0, nil
}

func (b *BuilderBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (b *BuilderBackend) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	return 0, nil
}

func (b *BuilderBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	b.transactions = append(b.transactions, tx)
	data, dest, amount, totalAmount := combineTxes(b.transactions)
	realData, err := validatorABI.Pack("executeTransactions", data, dest, amount)
	if err != nil {
		return err
	}
	msg := ethereum.CallMsg{
		From:     b.realSender,
		To:       &b.wallet,
		GasPrice: big.NewInt(1),
		Value:    totalAmount,
		Data:     realData,
	}
	_, err = b.realClient.EstimateGas(ctx, msg)
	return errors.WithStack(err)
}

func (b *BuilderBackend) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	panic("implement me")
}

func (b *BuilderBackend) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	panic("implement me")
}

func authWithContextAndAmount(ctx context.Context, auth *bind.TransactOpts, amount *big.Int) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:     auth.From,
		Nonce:    auth.Nonce,
		Signer:   auth.Signer,
		Value:    amount,
		GasPrice: auth.GasPrice,
		GasLimit: auth.GasLimit,
		Context:  ctx,
	}
}

func authWithContext(ctx context.Context, auth *bind.TransactOpts) *bind.TransactOpts {
	return authWithContextAndAmount(ctx, auth, big.NewInt(0))
}
