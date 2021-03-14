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
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var validatorABI abi.ABI

func init() {
	parsedValidator, err := abi.JSON(strings.NewReader(ethbridgecontracts.ValidatorABI))
	if err != nil {
		panic(err)
	}
	validatorABI = parsedValidator
}

type ValidatorWallet struct {
	con           *ethbridgecontracts.Validator
	address       ethcommon.Address
	client        ethutils.EthClient
	auth          *TransactAuth
	rollupAddress ethcommon.Address
}

func NewValidator(address, rollupAddress ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*ValidatorWallet, error) {
	con, err := ethbridgecontracts.NewValidator(address, client)
	if err != nil {
		return nil, err
	}
	return &ValidatorWallet{
		con:           con,
		address:       address,
		client:        client,
		auth:          auth,
		rollupAddress: rollupAddress,
	}, nil
}

func (v *ValidatorWallet) Address() common.Address {
	return common.NewAddressFromEth(v.address)
}

func (v *ValidatorWallet) From() common.Address {
	return common.NewAddressFromEth(v.auth.auth.From)
}

func (v *ValidatorWallet) RollupAddress() common.Address {
	return common.NewAddressFromEth(v.rollupAddress)
}

func (v *ValidatorWallet) executeTransaction(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	return v.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		auth.Value = tx.Value()
		return v.con.ExecuteTransaction(auth, tx.Data(), *tx.To(), tx.Value())
	})
}

func combineTxes(txes []*types.Transaction) ([][]byte, []ethcommon.Address, []*big.Int, *big.Int) {
	totalAmount := big.NewInt(0)
	data := make([][]byte, 0, len(txes))
	dest := make([]ethcommon.Address, 0, len(txes))
	amount := make([]*big.Int, 0, len(txes))

	for _, tx := range txes {
		data = append(data, tx.Data())
		dest = append(dest, *tx.To())
		amount = append(amount, tx.Value())
		totalAmount = totalAmount.Add(totalAmount, tx.Value())
	}
	return data, dest, amount, totalAmount
}

func (v *ValidatorWallet) ExecuteTransactions(ctx context.Context, builder *BuilderBackend) (*types.Transaction, error) {
	txes := builder.transactions
	if len(txes) == 0 {
		return nil, nil
	}

	if len(txes) == 1 {
		tx, err := v.executeTransaction(ctx, txes[0])
		if err != nil {
			return nil, err
		}
		builder.transactions = nil
		return tx, nil
	}

	totalAmount := big.NewInt(0)
	data := make([][]byte, 0, len(txes))
	dest := make([]ethcommon.Address, 0, len(txes))
	amount := make([]*big.Int, 0, len(txes))

	for _, tx := range txes {
		data = append(data, tx.Data())
		dest = append(dest, *tx.To())
		amount = append(amount, tx.Value())
		totalAmount = totalAmount.Add(totalAmount, tx.Value())
	}

	tx, err := v.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		auth.Value = totalAmount
		return v.con.ExecuteTransactions(auth, data, dest, amount)
	})
	if err != nil {
		return nil, err
	}
	builder.transactions = nil
	return tx, nil
}

func (v *ValidatorWallet) ReturnOldDeposits(ctx context.Context, stakers []common.Address) (*types.Transaction, error) {
	return v.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return v.con.ReturnOldDeposits(auth, v.rollupAddress, common.AddressArrayToEth(stakers))
	})
}

func (v *ValidatorWallet) TimeoutChallenges(ctx context.Context, challenges []common.Address) (*types.Transaction, error) {
	return v.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return v.con.TimeoutChallenges(auth, common.AddressArrayToEth(challenges))
	})
}
