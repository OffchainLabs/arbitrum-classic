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
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"math/big"

	errors2 "github.com/pkg/errors"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type arbFactory struct {
	*arbFactoryWatcher
	auth *TransactAuth
}

func newArbFactory(address ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*arbFactory, error) {
	watcher, err := newArbFactoryWatcher(address, client)
	if err != nil {
		return nil, err
	}
	return &arbFactory{arbFactoryWatcher: watcher, auth: auth}, nil
}

func DeployRollupFactory(auth *bind.TransactOpts, client ethutils.EthClient) (ethcommon.Address, error) {
	rollupAddr, _, _, err := ethbridgecontracts.DeployArbRollup(auth, client)
	if err != nil {
		return ethcommon.Address{}, err
	}
	inbox, _, _, err := ethbridgecontracts.DeployGlobalInbox(auth, client)
	if err != nil {
		return ethcommon.Address{}, err
	}
	chalFactory, err := DeployChallengeFactory(auth, client)
	if err != nil {
		return ethcommon.Address{}, err
	}
	factoryAddr, _, _, err := ethbridgecontracts.DeployArbFactory(auth, client, rollupAddr, inbox, chalFactory)
	return factoryAddr, err
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
		params.StakeToken.ToEthAddress(),
		owner.ToEthAddress(),
		[]byte{},
	)
	if err != nil {
		return common.Address{}, nil, errors2.Wrap(err, "Failed to call to ChainFactory.CreateChain")
	}
	receipt, err := WaitForReceiptWithResults(ctx, con.client, con.auth.auth.From, tx, "CreateChain")
	if err != nil {
		return common.Address{}, nil, err
	}
	if len(receipt.Logs) != 3 {
		return common.Address{}, nil, fmt.Errorf("wrong receipt count %v instead of 2", len(receipt.Logs))
	}
	event, err := con.contract.ParseRollupCreated(*receipt.Logs[2])
	if err != nil {
		return common.Address{}, nil, err
	}
	return common.NewAddressFromEth(event.RollupAddress), GetReceiptBlockID(receipt), nil
}

type arbFactoryWatcher struct {
	contract *ethbridgecontracts.ArbFactory
	client   ethutils.EthClient
	address  ethcommon.Address
}

func newArbFactoryWatcher(address ethcommon.Address, client ethutils.EthClient) (*arbFactoryWatcher, error) {
	vmCreatorContract, err := ethbridgecontracts.NewArbFactory(address, client)
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
