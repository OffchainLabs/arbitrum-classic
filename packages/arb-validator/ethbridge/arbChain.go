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
	"context"
	"math/big"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/arbchain"
)

type ArbChain struct {
	*ArbBase
	contract *arbchain.ArbChain
}

func NewArbChain(address common.Address, client *ethclient.Client) (*ArbChain, error) {
	arbVM, err := NewArbBase(address, client)
	return &ArbChain{ArbBase: arbVM}, err
}

func (vm *ArbChain) StartConnection(ctx context.Context) error {
	if err := vm.ArbBase.StartConnection(ctx); err != nil {
		return err
	}
	trackerContract, err := arbchain.NewArbChain(vm.address, vm.Client)
	if err != nil {
		return errors2.Wrap(err, "Failed to connect to ArbChain")
	}
	vm.contract = trackerContract
	return nil
}

func (vm *ArbChain) IncreaseDeposit(
	auth *bind.TransactOpts,
	amount *big.Int,
) (*types.Receipt, error) {
	call := &bind.TransactOpts{
		From:     auth.From,
		Nonce:    auth.Nonce,
		Signer:   auth.Signer,
		Value:    amount,
		GasPrice: auth.GasPrice,
		GasLimit: 100000,
		Context:  auth.Context,
	}
	tx, err := vm.contract.IncreaseDeposit(call)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, auth, tx, "IncreaseDeposit")
}
