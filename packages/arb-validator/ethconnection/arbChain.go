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

package ethconnection

import (
	"context"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethconnection/arblauncher"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ArbChain struct {
	*ArbitrumVM
	Tracker *arblauncher.ArbChain
}

func NewArbChain(address common.Address, client *ethclient.Client) (*ArbChain, error) {
	arbVM, err := NewArbitrumVM(address, client)
	return &ArbChain{ArbitrumVM: arbVM}, err
}

func (vm *ArbChain) StartConnection(ctx context.Context) error {
	if err := vm.ArbitrumVM.StartConnection(ctx); err != nil {
		return err
	}
	trackerContract, err := arblauncher.NewArbChain(vm.address, vm.Client)
	if err != nil {
		return errors2.Wrap(err, "Failed to connect to ArbChannel")
	}
	vm.Tracker = trackerContract
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
	tx, err := vm.Tracker.IncreaseDeposit(call)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, tx.Hash())
}
