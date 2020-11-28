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

package ethbridge

import (
	"context"
	"errors"
	errors2 "github.com/pkg/errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
)

type globalInbox struct {
	*globalInboxWatcher
	auth *TransactAuth
}

func newGlobalInbox(address ethcommon.Address, chain ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*globalInbox, error) {
	watcher, err := newGlobalInboxWatcher(address, chain, client)
	if err != nil {
		return nil, errors2.WithStack(errors2.Wrap(err, "Failed to connect to GlobalInbox"))
	}
	return &globalInbox{watcher, auth}, nil
}

func (con *globalInbox) SendL2Message(ctx context.Context, data []byte) (arbbridge.MessageDeliveredEvent, error) {
	con.auth.Lock()
	defer con.auth.Unlock()
	tx, err := con.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return con.GlobalInbox.SendL2MessageFromOrigin(
			auth,
			con.rollupAddress,
			data,
		)
	})
	if err != nil {
		return arbbridge.MessageDeliveredEvent{}, err
	}
	receipt, err := WaitForReceiptWithResults(ctx, con.client, con.auth.auth.From, tx, "SendL2MessageFromOrigin")
	if err != nil {
		return arbbridge.MessageDeliveredEvent{}, err
	}
	for _, evmLog := range receipt.Logs {
		if receipt.Logs[0].Topics[0] != messageDeliveredFromOriginID {
			continue
		}
		blockHeader, err := con.client.HeaderByHash(ctx, evmLog.BlockHash)
		if err != nil {
			return arbbridge.MessageDeliveredEvent{}, err
		}
		timestamp := new(big.Int).SetUint64(blockHeader.Time)
		return con.parseMessageFromOrigin(*evmLog, timestamp, data)
	}
	return arbbridge.MessageDeliveredEvent{}, errors.New("didn't output l2message delivered event")
}

func (con *globalInbox) SendL2MessageNoWait(ctx context.Context, data []byte) (common.Hash, error) {
	con.auth.Lock()
	defer con.auth.Unlock()
	tx, err := con.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return con.GlobalInbox.SendL2MessageFromOrigin(
			auth,
			con.rollupAddress,
			data,
		)
	})
	if err != nil {
		return common.Hash{}, err
	}
	return common.NewHashFromEth(tx.Hash()), nil
}

func (con *globalInbox) DepositEthMessage(
	ctx context.Context,
	destination common.Address,
	value *big.Int,
) error {
	con.auth.Lock()
	defer con.auth.Unlock()
	tx, err := con.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return con.GlobalInbox.DepositEthMessage(
			&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasLimit: auth.GasLimit,
				Nonce:    auth.Nonce,
				Value:    value,
				Context:  ctx,
			},
			con.rollupAddress,
			destination.ToEthAddress(),
		)
	})

	if err != nil {
		return err
	}

	return con.waitForReceipt(ctx, tx, "DepositEthMessage")
}

func (con *globalInbox) DepositERC20Message(
	ctx context.Context,
	tokenAddress common.Address,
	destination common.Address,
	value *big.Int,
) error {
	con.auth.Lock()
	defer con.auth.Unlock()
	tx, err := con.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return con.GlobalInbox.DepositERC20Message(
			auth,
			con.rollupAddress,
			tokenAddress.ToEthAddress(),
			destination.ToEthAddress(),
			value,
		)
	})

	if err != nil {
		return err
	}

	return con.waitForReceipt(ctx, tx, "DepositERC20Message")
}

func (con *globalInbox) DepositERC721Message(
	ctx context.Context,
	tokenAddress common.Address,
	destination common.Address,
	value *big.Int,
) error {
	con.auth.Lock()
	defer con.auth.Unlock()
	tx, err := con.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return con.GlobalInbox.DepositERC721Message(
			auth,
			con.rollupAddress,
			tokenAddress.ToEthAddress(),
			destination.ToEthAddress(),
			value,
		)
	})

	if err != nil {
		return err
	}

	return con.waitForReceipt(ctx, tx, "DepositERC721Message")
}

func (con *globalInbox) waitForReceipt(ctx context.Context, tx *types.Transaction, methodName string) error {
	return waitForReceipt(ctx, con.client, con.auth.auth.From, tx, methodName)
}
