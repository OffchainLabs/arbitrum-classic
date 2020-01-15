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
	"bytes"
	"context"
	"math/big"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/globalpendinginbox"
)

type pendingInbox struct {
	GlobalPendingInbox *globalpendinginbox.GlobalPendingInbox
	client             *ethclient.Client
	auth               *TransactAuth
}

func newPendingInbox(address ethcommon.Address, client *ethclient.Client, auth *TransactAuth) (*pendingInbox, error) {
	globalPendingInboxContract, err := globalpendinginbox.NewGlobalPendingInbox(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to GlobalPendingInbox")
	}
	return &pendingInbox{globalPendingInboxContract, client, auth}, nil
}

func (con *pendingInbox) SendTransactionMessage(
	ctx context.Context,
	data value.Value,
	vmAddress common.Address,
	amount *big.Int,
	seqNumber *big.Int,
) error {
	var dataBuf bytes.Buffer

	if err := value.MarshalValue(data, &dataBuf); err != nil {
		return err
	}

	con.auth.Lock()
	defer con.auth.Unlock()
	tx, err := con.GlobalPendingInbox.SendTransactionMessage(
		con.auth.getAuth(ctx),
		vmAddress.ToEthAddress(),
		seqNumber,
		amount,
		dataBuf.Bytes(),
	)
	if err != nil {
		return err
	}
	return con.waitForReceipt(ctx, tx, "SendTransactionMessage")
}

func (con *pendingInbox) DepositEthMessage(
	ctx context.Context,
	vmAddress common.Address,
	destination common.Address,
	value *big.Int,
) error {

	tx, err := con.GlobalPendingInbox.DepositEthMessage(
		&bind.TransactOpts{
			From:     con.auth.auth.From,
			Signer:   con.auth.auth.Signer,
			GasLimit: con.auth.auth.GasLimit,
			Value:    value,
			Context:  ctx,
		},
		vmAddress.ToEthAddress(),
		destination.ToEthAddress(),
	)

	if err != nil {
		return err
	}

	return con.waitForReceipt(ctx, tx, "DepositEthMessage")
}

func (con *pendingInbox) DepositERC20Message(
	ctx context.Context,
	vmAddress common.Address,
	tokenAddress common.Address,
	destination common.Address,
	value *big.Int,
) error {
	con.auth.Lock()
	defer con.auth.Unlock()
	tx, err := con.GlobalPendingInbox.DepositERC20Message(
		con.auth.getAuth(ctx),
		vmAddress.ToEthAddress(),
		tokenAddress.ToEthAddress(),
		destination.ToEthAddress(),
		value,
	)

	if err != nil {
		return err
	}

	return con.waitForReceipt(ctx, tx, "DepositERC20Message")
}

func (con *pendingInbox) DepositERC721Message(
	ctx context.Context,
	vmAddress common.Address,
	tokenAddress common.Address,
	destination common.Address,
	value *big.Int,
) error {
	con.auth.Lock()
	defer con.auth.Unlock()
	tx, err := con.GlobalPendingInbox.DepositERC721Message(
		con.auth.getAuth(ctx),
		vmAddress.ToEthAddress(),
		tokenAddress.ToEthAddress(),
		destination.ToEthAddress(),
		value,
	)

	if err != nil {
		return err
	}

	return con.waitForReceipt(ctx, tx, "DepositERC721Message")
}

func (con *pendingInbox) DepositFunds(ctx context.Context, amount *big.Int, dest common.Address) error {
	con.auth.Lock()
	defer con.auth.Unlock()
	tx, err := con.GlobalPendingInbox.DepositEth(
		&bind.TransactOpts{
			From:     con.auth.auth.From,
			Signer:   con.auth.auth.Signer,
			GasLimit: con.auth.auth.GasLimit,
			Value:    amount,
			Context:  ctx,
		},
		dest.ToEthAddress(),
	)
	if err != nil {
		return err
	}
	return con.waitForReceipt(ctx, tx, "DepositFunds")
}

func (con *pendingInbox) GetTokenBalance(
	ctx context.Context,
	user common.Address,
	tokenContract common.Address,
) (*big.Int, error) {
	return con.GlobalPendingInbox.GetTokenBalance(
		&bind.CallOpts{Context: ctx},
		tokenContract.ToEthAddress(),
		user.ToEthAddress(),
	)
}

func (con *pendingInbox) waitForReceipt(ctx context.Context, tx *types.Transaction, methodName string) error {
	return waitForReceipt(ctx, con.client, con.auth.auth.From, tx, methodName)
}
