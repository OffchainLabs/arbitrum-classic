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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

type pendingInbox struct {
	GlobalPendingInbox *globalpendinginbox.GlobalPendingInbox
	client             *ethclient.Client
	auth               *bind.TransactOpts
}

func newPendingInbox(address ethcommon.Address, client *ethclient.Client, auth *bind.TransactOpts) (*pendingInbox, error) {
	globalPendingInboxContract, err := globalpendinginbox.NewGlobalPendingInbox(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to GlobalPendingInbox")
	}
	return &pendingInbox{globalPendingInboxContract, client, auth}, nil
}

func (con *pendingInbox) SendMessage(
	ctx context.Context,
	msg valprotocol.Message,
) error {
	var dataBuf bytes.Buffer
	if err := value.MarshalValue(msg.Data, &dataBuf); err != nil {
		return err
	}
	con.auth.Context = ctx
	tx, err := con.GlobalPendingInbox.SendMessage(
		con.auth,
		msg.Destination.ToEthAddress(),
		msg.TokenType,
		msg.Currency,
		dataBuf.Bytes(),
	)
	if err != nil {
		return err
	}
	return con.waitForReceipt(ctx, tx, "SendMessage")
}

func (con *pendingInbox) ForwardMessage(
	ctx context.Context,
	msg valprotocol.Message,
	sig []byte,
) error {
	var dataBuf bytes.Buffer
	if err := value.MarshalValue(msg.Data, &dataBuf); err != nil {
		return err
	}
	con.auth.Context = ctx
	tx, err := con.GlobalPendingInbox.ForwardMessage(
		con.auth,
		msg.Destination.ToEthAddress(),
		msg.TokenType,
		msg.Currency,
		dataBuf.Bytes(),
		sig,
	)
	if err != nil {
		return err
	}
	return con.waitForReceipt(ctx, tx, "ForwardMessage")
}

func (con *pendingInbox) SendEthMessage(
	ctx context.Context,
	data value.Value,
	destination common.Address,
	amount *big.Int,
) error {
	var dataBuf bytes.Buffer
	if err := value.MarshalValue(data, &dataBuf); err != nil {
		return err
	}
	tx, err := con.GlobalPendingInbox.SendEthMessage(
		&bind.TransactOpts{
			From:     con.auth.From,
			Signer:   con.auth.Signer,
			GasLimit: con.auth.GasLimit,
			Value:    amount,
			Context:  ctx,
		},
		destination.ToEthAddress(),
		dataBuf.Bytes(),
	)
	if err != nil {
		return err
	}
	return con.waitForReceipt(ctx, tx, "SendEthMessage")
}

func (con *pendingInbox) DepositFunds(ctx context.Context, amount *big.Int, dest common.Address) error {
	tx, err := con.GlobalPendingInbox.DepositEth(
		&bind.TransactOpts{
			From:     con.auth.From,
			Signer:   con.auth.Signer,
			GasLimit: con.auth.GasLimit,
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
	return waitForReceipt(ctx, con.client, con.auth.From, tx, methodName)
}
