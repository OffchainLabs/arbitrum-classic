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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type ethRollupWatcher struct {
	client *MockArbClient

	address common.Address
}

func newRollupWatcher(address common.Address, client *MockArbClient) (*ethRollupWatcher, error) {
	vm := &ethRollupWatcher{client: client, address: address}
	//err := vm.setupContracts()
	return vm, nil
	//arbitrumRollupContract, err := rollup.NewArbRollup(rollupAddress, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to arbRollup")
	//}
	//
	//globalPendingInboxAddress, err := arbitrumRollupContract.GlobalInbox(&bind.CallOpts{
	//	Pending: false,
	//	Context: context.Background(),
	//})
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to get GlobalPendingInbox address")
	//}
	//globalPendingContract, err := rollup.NewIGlobalPendingInbox(globalPendingInboxAddress, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to GlobalPendingInbox")
	//}
	//
	//return &ethRollupWatcher{
	//	ArbRollup:          arbitrumRollupContract,
	//	GlobalPendingInbox: globalPendingContract,
	//	rollupAddress:      rollupAddress,
	//	inboxAddress:       globalPendingInboxAddress,
	//	client:             client,
	//}, nil
}

//func (vm *ethRollupWatcher) setupContracts() error {
//	return nil
//}

func (vm *ethRollupWatcher) GetEvents(ctx context.Context, blockId *structures.BlockId) ([]arbbridge.Event, error) {
	// copy events
	//vm.client.MockEthClient.Lock()
	//var blockEvents []arbbridge.Event
	//for i, b := range vm.client.MockEthClient.rollups[vm.address].events[blockId] {
	//	blockEvents[i] = b
	//}
	//return TupleValue{newContents, tv.itemCount, tv.cachedHash, tv.size}

	// delete events
	// return events
	return vm.client.MockEthClient.rollups[vm.address].events[blockId], nil
}

func (vm *ethRollupWatcher) GetParams(ctx context.Context) (structures.ChainParams, error) {
	return structures.ChainParams{
		StakeRequirement:        vm.client.MockEthClient.rollups[vm.address].escrowRequired,
		GracePeriod:             vm.client.MockEthClient.rollups[vm.address].gracePeriod,
		MaxExecutionSteps:       vm.client.MockEthClient.rollups[vm.address].maxSteps,
		ArbGasSpeedLimitPerTick: 200000,
	}, nil
}

func (vm *ethRollupWatcher) InboxAddress(ctx context.Context) (common.Address, error) {
	return common.Address{}, nil
}

func (vm *ethRollupWatcher) GetCreationHeight(ctx context.Context) (*structures.BlockId, error) {
	return vm.client.MockEthClient.rollups[vm.address].creation, nil
}
