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

package gobridge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type ethRollupWatcher struct {
	inboxAddress common.Address
	client       *GoArbClient

	rollupAddress common.Address
}

func newRollupWatcher(address common.Address, client *GoArbClient) (*ethRollupWatcher, error) {
	vm := &ethRollupWatcher{client: client, rollupAddress: address}
	return vm, nil
}

func (vm *ethRollupWatcher) GetEvents(ctx context.Context, blockId *common.BlockId) ([]arbbridge.Event, error) {
	return vm.client.GoEthClient.rollups[vm.rollupAddress].events[blockId], nil
}

func (vm *ethRollupWatcher) GetParams(ctx context.Context) (valprotocol.ChainParams, error) {
	return valprotocol.ChainParams{
		StakeRequirement:        vm.client.GoEthClient.rollups[vm.rollupAddress].escrowRequired,
		GracePeriod:             vm.client.GoEthClient.rollups[vm.rollupAddress].gracePeriod,
		MaxExecutionSteps:       vm.client.GoEthClient.rollups[vm.rollupAddress].maxSteps,
		ArbGasSpeedLimitPerTick: vm.client.GoEthClient.rollups[vm.rollupAddress].arbGasSpeedLimitPerTick,
		MaxTimeBoundsWidth:      vm.client.GoEthClient.rollups[vm.rollupAddress].maxTimeBoundsWidth,
	}, nil
}

func (con *ethRollupWatcher) GetCreationInfo(ctx context.Context) (*common.BlockId, common.Hash, error) {
	return con.client.GoEthClient.rollups[con.rollupAddress].creation, con.client.GoEthClient.rollups[con.rollupAddress].initVMHash, nil
}

func (con *ethRollupWatcher) GetVersion(ctx context.Context) (string, error) {
	return string("1"), nil
}

func (vm *ethRollupWatcher) InboxAddress(ctx context.Context) (common.Address, error) {
	return vm.client.GoEthClient.getNextAddress(), nil
}

func (vm *ethRollupWatcher) GetCreationHeight(ctx context.Context) (*common.BlockId, error) {
	return vm.client.GoEthClient.rollups[vm.rollupAddress].creation, nil
}
