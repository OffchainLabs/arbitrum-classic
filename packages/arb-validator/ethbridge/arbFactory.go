/*
 * Copyright 2020, Offchain Labs, Inc.
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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/arbfactory"
	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ArbFactory struct {
	contract *arbfactory.ArbFactory
	client   *ethclient.Client
}

func NewArbFactory(address common.Address, client arbbridge.ArbClient) (*ArbFactory, error) {
	vmCreatorContract, err := arbfactory.NewArbFactory(address, client.(*EthArbClient).client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to ArbFactory")
	}
	return &ArbFactory{vmCreatorContract, client.(*EthArbClient).client}, nil
}

func (con *ArbFactory) CreateRollup(
	auth *bind.TransactOpts,
	vmState [32]byte,
	params structures.ChainParams,
	owner common.Address,
) (common.Address, error) {
	tx, err := con.contract.CreateRollup(
		auth,
		vmState,
		params.GracePeriod.Val,
		new(big.Int).SetUint64(params.ArbGasSpeedLimitPerTick),
		params.MaxExecutionSteps,
		params.StakeRequirement,
		owner,
	)
	if err != nil {
		return common.Address{}, errors2.Wrap(err, "Failed to call to ChainFactory.CreateChain")
	}
	receipt, err := waitForReceiptWithResults(auth.Context, con.client, auth.From, tx, "CreateChain")
	if err != nil {
		return common.Address{}, err
	}
	if len(receipt.Logs) != 1 {
		return common.Address{}, errors2.New("Wrong receipt count")
	}
	event, err := con.contract.ParseRollupCreated(*receipt.Logs[0])
	if err != nil {
		return common.Address{}, err
	}
	return event.VmAddress, nil
}
