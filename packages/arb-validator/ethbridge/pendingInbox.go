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
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethconnection"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type PendingInbox struct {
	contract *ethconnection.PendingInbox
	client   *ethclient.Client
}

func NewPendingInbox(address common.Address, client *ethclient.Client) (*PendingInbox, error) {
	pendingInbox, err := ethconnection.NewPendingInbox(address, client)
	return &PendingInbox{pendingInbox, client}, err
}

func (con *PendingInbox) SendMessage(
	auth *bind.TransactOpts,
	msg protocol.Message,
) (*types.Receipt, error) {
	tx, err := con.contract.SendMessage(
		auth,
		msg,
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
	tx, err := con.contract.ForwardMessage(
		auth,
		msg,
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
	tx, err := con.contract.SendEthMessage(
		auth,
		data,
		destination,
		amount,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, con.client, tx.Hash())
}

func (con *PendingInbox) DepositFunds(auth *bind.TransactOpts, amount *big.Int, dest common.Address) (*types.Receipt, error) {
	tx, err := con.contract.DepositFunds(
		auth,
		amount,
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
	return con.contract.GetTokenBalance(
		auth,
		tokenContract,
		user,
	)
}
