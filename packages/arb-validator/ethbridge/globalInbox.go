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
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/globalinbox"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type globalInbox struct {
	GlobalInbox *globalinbox.GlobalInbox
	client      *ethclient.Client
	auth        *TransactAuth
}

func newGlobalInbox(address ethcommon.Address, client *ethclient.Client, auth *TransactAuth) (*globalInbox, error) {
	globalInboxContract, err := globalinbox.NewGlobalInbox(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to GlobalInbox")
	}
	return &globalInbox{globalInboxContract, client, auth}, nil
}

func (con *globalInbox) SendTransactionMessage(ctx context.Context, data []byte, vmAddress common.Address, contactAddress common.Address, amount *big.Int, seqNumber *big.Int) error {
	con.auth.Lock()
	defer con.auth.Unlock()
	tx, err := con.GlobalInbox.SendTransactionMessage(
		con.auth.getAuth(ctx),
		vmAddress.ToEthAddress(),
		contactAddress.ToEthAddress(),
		seqNumber,
		amount,
		data,
	)
	if err != nil {
		return err
	}
	return con.waitForReceipt(ctx, tx, "SendTransactionMessage")
}

func (con *globalInbox) DepositEthMessage(
	ctx context.Context,
	vmAddress common.Address,
	destination common.Address,
	value *big.Int,
) error {

	tx, err := con.GlobalInbox.DepositEthMessage(
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

func (con *globalInbox) DepositERC20Message(
	ctx context.Context,
	vmAddress common.Address,
	tokenAddress common.Address,
	destination common.Address,
	value *big.Int,
) error {
	con.auth.Lock()
	defer con.auth.Unlock()
	tx, err := con.GlobalInbox.DepositERC20Message(
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

func (con *globalInbox) DepositERC721Message(
	ctx context.Context,
	vmAddress common.Address,
	tokenAddress common.Address,
	destination common.Address,
	value *big.Int,
) error {
	con.auth.Lock()
	defer con.auth.Unlock()
	tx, err := con.GlobalInbox.DepositERC721Message(
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

func (con *globalInbox) GetTokenBalance(
	ctx context.Context,
	user common.Address,
	tokenContract common.Address,
) (*big.Int, error) {
	return con.GlobalInbox.GetERC20Balance(
		&bind.CallOpts{Context: ctx},
		tokenContract.ToEthAddress(),
		user.ToEthAddress(),
	)
}

func (con *globalInbox) waitForReceipt(ctx context.Context, tx *types.Transaction, methodName string) error {
	return waitForReceipt(ctx, con.client, con.auth.auth.From, tx, methodName)
}
