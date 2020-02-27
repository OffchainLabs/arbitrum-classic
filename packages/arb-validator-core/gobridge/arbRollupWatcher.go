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

func (rw *ethRollupWatcher) GetEvents(ctx context.Context, blockId *common.BlockId) ([]arbbridge.Event, error) {
	return rw.client.GoEthClient.rollups[rw.rollupAddress].events[blockId], nil
}

func (rw *ethRollupWatcher) GetParams(ctx context.Context) (valprotocol.ChainParams, error) {
	return valprotocol.ChainParams{
		StakeRequirement:        rw.client.GoEthClient.rollups[rw.rollupAddress].escrowRequired,
		GracePeriod:             rw.client.GoEthClient.rollups[rw.rollupAddress].gracePeriod,
		MaxExecutionSteps:       rw.client.GoEthClient.rollups[rw.rollupAddress].maxSteps,
		ArbGasSpeedLimitPerTick: rw.client.GoEthClient.rollups[rw.rollupAddress].arbGasSpeedLimitPerTick,
		MaxTimeBoundsWidth:      rw.client.GoEthClient.rollups[rw.rollupAddress].maxTimeBoundsWidth,
	}, nil
}

func (rw *ethRollupWatcher) GetCreationInfo(ctx context.Context) (*common.BlockId, common.Hash, error) {
	return rw.client.GoEthClient.rollups[rw.rollupAddress].creation, rw.client.GoEthClient.rollups[rw.rollupAddress].initVMHash, nil
}

func (rw *ethRollupWatcher) GetVersion(ctx context.Context) (string, error) {
	return string("1"), nil
}

func (rw *ethRollupWatcher) InboxAddress(ctx context.Context) (common.Address, error) {
	return rw.client.GoEthClient.getNextAddress(), nil
}

func (rw *ethRollupWatcher) GetCreationHeight(ctx context.Context) (*common.BlockId, error) {
	return rw.client.GoEthClient.rollups[rw.rollupAddress].creation, nil
}
