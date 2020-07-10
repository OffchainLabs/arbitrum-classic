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
	"context"
	"math/big"

	errors2 "github.com/pkg/errors"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge/arbfactory"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type arbFactory struct {
	contract *arbfactory.ArbFactory
	client   *ethclient.Client
	auth     *TransactAuth
}

func newArbFactory(address ethcommon.Address, client *ethclient.Client, auth *TransactAuth) (*arbFactory, error) {
	vmCreatorContract, err := arbfactory.NewArbFactory(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to arbFactory")
	}
	return &arbFactory{vmCreatorContract, client, auth}, nil
}

func (con *arbFactory) CreateRollup(
	ctx context.Context,
	vmState common.Hash,
	params valprotocol.ChainParams,
	owner common.Address,
) (common.Address, *common.BlockId, error) {
	con.auth.Lock()
	defer con.auth.Unlock()
	tx, err := con.contract.CreateRollup(
		con.auth.getAuth(ctx),
		vmState,
		params.GracePeriod.Val,
		new(big.Int).SetUint64(params.ArbGasSpeedLimitPerTick),
		params.MaxExecutionSteps,
		params.StakeRequirement,
		owner.ToEthAddress(),
	)
	if err != nil {
		return common.Address{}, nil, errors2.Wrap(err, "Failed to call to ChainFactory.CreateChain")
	}
	receipt, err := WaitForReceiptWithResults(ctx, con.client, con.auth.auth.From, tx, "CreateChain")
	if err != nil {
		return common.Address{}, nil, err
	}
	if len(receipt.Logs) != 2 {
		return common.Address{}, nil, errors2.New("Wrong receipt count")
	}
	event, err := con.contract.ParseRollupCreated(*receipt.Logs[1])
	if err != nil {
		return common.Address{}, nil, err
	}
	return common.NewAddressFromEth(event.VmAddress), GetReceiptBlockID(receipt), nil
}

type arbFactoryWatcher struct {
	contract *arbfactory.ArbFactory
	client   *ethclient.Client
	address  ethcommon.Address
}

func newArbFactoryWatcher(address ethcommon.Address, client *ethclient.Client) (*arbFactoryWatcher, error) {
	vmCreatorContract, err := arbfactory.NewArbFactory(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to arbFactory")
	}
	return &arbFactoryWatcher{contract: vmCreatorContract, client: client, address: address}, nil
}

func (con *arbFactoryWatcher) GlobalInboxAddress() (common.Address, error) {
	addr, err := con.contract.GlobalInboxAddress(nil)
	return common.NewAddressFromEth(addr), err
}

func (con *arbFactoryWatcher) ChallengeFactoryAddress() (common.Address, error) {
	addr, err := con.contract.ChallengeFactoryAddress(nil)
	return common.NewAddressFromEth(addr), err
}
