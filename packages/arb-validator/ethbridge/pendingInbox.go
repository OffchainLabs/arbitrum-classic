/*
 * Copyright 2019, Offchain Labs, Inc.
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
	"math/big"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/globalpendinginbox"
)

type PendingInbox struct {
	GlobalPendingInbox *globalpendinginbox.GlobalPendingInbox
	client             *ethclient.Client
}

func NewPendingInbox(address common.Address, client *ethclient.Client) (*PendingInbox, error) {
	globalPendingInboxContract, err := globalpendinginbox.NewGlobalPendingInbox(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to GlobalPendingInbox")
	}
	return &PendingInbox{globalPendingInboxContract, client}, nil
}

func (con *PendingInbox) SendMessage(
	auth *bind.TransactOpts,
	msg protocol.Message,
) (*types.Receipt, error) {
	var dataBuf bytes.Buffer
	if err := value.MarshalValue(msg.Data, &dataBuf); err != nil {
		return nil, err
	}
	tx, err := con.GlobalPendingInbox.SendMessage(
		auth,
		msg.Destination,
		msg.TokenType,
		msg.Currency,
		dataBuf.Bytes(),
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, con.client, tx.Hash())
}

func (con *PendingInbox) ForwardMessage(
	auth *bind.TransactOpts,
	msg protocol.Message,
	sig []byte,
) (*types.Receipt, error) {
	var dataBuf bytes.Buffer
	if err := value.MarshalValue(msg.Data, &dataBuf); err != nil {
		return nil, err
	}
	tx, err := con.GlobalPendingInbox.ForwardMessage(
		auth,
		msg.Destination,
		msg.TokenType,
		msg.Currency,
		dataBuf.Bytes(),
		sig,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, con.client, tx.Hash())
}

func (con *PendingInbox) SendEthMessage(
	auth *bind.TransactOpts,
	data value.Value,
	destination common.Address,
	amount *big.Int,
) (*types.Receipt, error) {
	var dataBuf bytes.Buffer
	if err := value.MarshalValue(data, &dataBuf); err != nil {
		return nil, err
	}
	tx, err := con.GlobalPendingInbox.SendEthMessage(
		&bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: auth.GasLimit,
			Value:    amount,
		},
		destination,
		dataBuf.Bytes(),
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, con.client, tx.Hash())
}

func (con *PendingInbox) DepositFunds(auth *bind.TransactOpts, amount *big.Int, dest common.Address) (*types.Receipt, error) {
	tx, err := con.GlobalPendingInbox.DepositEth(
		&bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: auth.GasLimit,
			Value:    amount,
		},
		dest,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, con.client, tx.Hash())
}

func (con *PendingInbox) GetTokenBalance(
	auth *bind.CallOpts,
	user common.Address,
	tokenContract common.Address,
) (*big.Int, error) {
	return con.GlobalPendingInbox.GetTokenBalance(
		auth,
		tokenContract,
		user,
	)
}
