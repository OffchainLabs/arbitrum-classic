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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethconnection"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type VMCreator struct {
	contract *ethconnection.VMCreator
	client   *ethclient.Client
}

func NewVMCreator(address common.Address, client *ethclient.Client) (*VMCreator, error) {
	vm, err := ethconnection.NewVMCreator(address, client)
	return &VMCreator{vm, client}, err
}

func (con *VMCreator) LaunchVM(
	auth *bind.TransactOpts,
	config *valmessage.VMConfiguration,
	vmState [32]byte,
) (*types.Receipt, error) {
	tx, err := con.contract.LaunchVM(
		auth,
		config,
		vmState,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, con.client, tx.Hash())
}

func (con *VMCreator) ParseVMCreated(log *types.Log) (common.Address, [32]byte, *valmessage.VMConfiguration, error) {
	return con.contract.ParseVMCreated(log)
}
