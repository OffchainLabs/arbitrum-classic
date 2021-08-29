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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/arbtransaction"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/fireblocks"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
)

type DelayedInboxFireblocksTransactAuth struct {
	inner        *FireblocksTransactAuth
	auth         *bind.TransactOpts // has remapped from address
	l1Client     ethutils.EthClient
	l2Client     ethutils.EthClient
	l2ChainId    *big.Int
	delayedInbox *ethbridgecontracts.Inbox
}

func l2RemapAccount(account ethcommon.Address) ethcommon.Address {
	magic, _ := new(big.Int).SetString("1111000000000000000000000000000000001111", 16)
	overflow := new(big.Int).Exp(big.NewInt(2), big.NewInt(20*8), nil)

	translated := new(big.Int).SetBytes(account.Bytes())
	translated.Add(translated, magic)
	if translated.Cmp(overflow) == 1 {
		translated.Sub(translated, overflow)
	}

	return ethcommon.BigToAddress(translated)
}

func NewDelayedInboxFireblocksTransactAuthAdvanced(
	ctx context.Context,
	l1Client ethutils.EthClient,
	l2Client ethutils.EthClient,
	auth *bind.TransactOpts,
	walletConfig *configuration.Wallet,
	delayedInboxAddress ethcommon.Address,
	remapFromAddress bool,
	usePendingNonce bool,
) (TransactAuth, *fireblocks.Fireblocks, error) {
	// The inner TransactAuth doesn't need an EthClient
	inner, fb, err := NewFireblocksTransactAuthAdvanced(ctx, l1Client, auth, walletConfig, usePendingNonce)
	if err != nil {
		return nil, nil, err
	}
	delayedInbox, err := ethbridgecontracts.NewInbox(delayedInboxAddress, l1Client)
	if err != nil {
		return nil, nil, err
	}
	l2ChainId, err := l2Client.ChainID(ctx)
	if err != nil {
		return nil, nil, err
	}
	l2Auth := *auth
	l2Auth.From = l2RemapAccount(auth.From)

	wrapped := &DelayedInboxFireblocksTransactAuth{
		inner:        inner,
		auth:         &l2Auth,
		l1Client:     l1Client,
		l2Client:     l2Client,
		l2ChainId:    l2ChainId,
		delayedInbox: delayedInbox,
	}
	return wrapped, fb, nil
}

func (ta *DelayedInboxFireblocksTransactAuth) TransactionReceipt(ctx context.Context, tx *arbtransaction.ArbTransaction) (*types.Receipt, error) {
	err := ta.inner.checkIfFailed(tx)
	if err != nil {
		return nil, err
	}

	return ta.l2Client.TransactionReceipt(ctx, tx.Hash())
}

func (ta *DelayedInboxFireblocksTransactAuth) NonceAt(ctx context.Context, account ethcommon.Address, blockNumber *big.Int) (uint64, error) {
	return ta.l2Client.NonceAt(ctx, account, blockNumber)
}

func (ta *DelayedInboxFireblocksTransactAuth) SendTransaction(ctx context.Context, tx *types.Transaction, replaceTxByHash string) (*arbtransaction.ArbTransaction, error) {
	var hashData []byte
	hashData = append(hashData, 0) // L2MessageType_unsignedEOATx
	hashData = append(hashData, math.U256Bytes(new(big.Int).SetUint64(tx.Gas()))...)
	hashData = append(hashData, math.U256Bytes(tx.GasPrice())...)
	hashData = append(hashData, math.U256Bytes(new(big.Int).SetUint64(tx.Nonce()))...)
	hashData = append(hashData, tx.To().Hash().Bytes()...)
	hashData = append(hashData, math.U256Bytes(tx.Value())...)
	hashData = append(hashData, tx.Data()...)
	innermostHash := hashing.SoliditySHA3(hashData)

	hashData = []byte{}
	hashData = append(hashData, math.U256Bytes(ta.l2ChainId)...)
	hashData = append(hashData, innermostHash.Bytes()...)
	hashWithChainId := hashing.SoliditySHA3(hashData)

	hashData = []byte{}
	hashData = append(hashData, ta.From().Hash().Bytes()...)
	hashData = append(hashData, hashWithChainId.Bytes()...)
	l2TxHash := hashing.SoliditySHA3(hashData)

	wrappedTx, err := ta.delayedInbox.SendL1FundedUnsignedTransaction(
		ta.inner.GetAuth(),
		new(big.Int).SetUint64(tx.Gas()),
		tx.GasPrice(),
		new(big.Int).SetUint64(tx.Nonce()),
		*tx.To(),
		tx.Data(),
	)
	if err != nil {
		return nil, err
	}

	arbTx, err := ta.inner.SendTransaction(ctx, wrappedTx, replaceTxByHash)
	if arbTx != nil {
		arbTx.OverrideHash(l2TxHash.ToEthHash())
	}
	return arbTx, err
}

func (ta *DelayedInboxFireblocksTransactAuth) Sign(addr ethcommon.Address, tx *types.Transaction) (*types.Transaction, error) {
	return ta.inner.Sign(addr, tx)
}

func (ta *DelayedInboxFireblocksTransactAuth) GetAuth() *bind.TransactOpts {
	return ta.auth
}

func (ta *DelayedInboxFireblocksTransactAuth) From() ethcommon.Address {
	return ta.auth.From
}
