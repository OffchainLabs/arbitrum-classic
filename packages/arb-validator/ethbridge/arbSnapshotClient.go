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
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/rollup"
	errors2 "github.com/pkg/errors"
	"math/big"
)

type arbSnapshotClient struct {
	Client    *ethclient.Client
	ArbRollup *rollup.ArbRollup
	auth      *TransactAuth
}

func newSnapshotClient(ru *arbRollup) (*arbSnapshotClient, error) {
	return &arbSnapshotClient{ru.Client, ru.ArbRollup, ru.auth}, nil
}

func (vm *arbSnapshotClient) SnapshotLatestConfirmed(
	ctx context.Context,
	idx *big.Int,
) (common.Hash, common.Hash, error) {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.ArbRollup.SnapshotLatestConfirmed(
		vm.auth.getAuth(ctx),
		idx,
	)
	if err != nil {
		return common.Hash{}, common.Hash{}, err
	}
	receipt, err := WaitForReceiptWithResults(ctx, vm.Client, vm.auth.auth.From, tx, "SnapshotLatestConfirmed")
	if err != nil {
		return common.Hash{}, common.Hash{}, err
	}
	if len(receipt.Logs) != 1 {
		return common.Hash{}, common.Hash{}, errors2.New("Wrong receipt count")
	}
	event, err := vm.ArbRollup.ParseSavedLatestConfirmedSnapshot(*receipt.Logs[0])
	if err != nil {
		return common.Hash{}, common.Hash{}, err
	}
	return event.LatestConfirmed, event.Snapshot, nil
}

func (vm *arbSnapshotClient) SnapshotLeafNodeExists(
	ctx context.Context,
	idx *big.Int,
	nodeHash common.Hash,
) (common.Hash, error) {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.ArbRollup.SnapshotLeafNodeExists(
		vm.auth.getAuth(ctx),
		idx,
		nodeHash,
	)
	if err != nil {
		return common.Hash{}, err
	}
	receipt, err := WaitForReceiptWithResults(ctx, vm.Client, vm.auth.auth.From, tx, "SnapshotLeafNodeExists")
	if err != nil {
		return common.Hash{}, err
	}
	if len(receipt.Logs) != 1 {
		return common.Hash{}, errors2.New("Wrong receipt count")
	}
	event, err := vm.ArbRollup.ParseSavedNodeExistsSnapshot(*receipt.Logs[0])
	if err != nil {
		return common.Hash{}, err
	}
	return event.Snapshot, nil
}

func (vm *arbSnapshotClient) SnapshotStakerNodeExists(
	ctx context.Context,
	idx *big.Int,
	stakerAddr common.Address,
) (common.Hash, common.Hash, error) {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.ArbRollup.SnapshotStakerNodeExists(
		vm.auth.getAuth(ctx),
		idx,
		stakerAddr.ToEthAddress(),
	)
	if err != nil {
		return common.Hash{}, common.Hash{}, err
	}
	receipt, err := WaitForReceiptWithResults(ctx, vm.Client, vm.auth.auth.From, tx, "SnapshotStakerNodeExists")
	if err != nil {
		return common.Hash{}, common.Hash{}, err
	}
	if len(receipt.Logs) != 1 {
		return common.Hash{}, common.Hash{}, errors2.New("Wrong receipt count")
	}
	event, err := vm.ArbRollup.ParseSavedNodeExistsSnapshot(*receipt.Logs[0])
	if err != nil {
		return common.Hash{}, common.Hash{}, err
	}
	return event.NodeHash, event.Snapshot, nil
}

func (vm *arbSnapshotClient) SnapshotTwoStakers(
	ctx context.Context,
	idx *big.Int,
	stakerAddr1 common.Address,
	stakerAddr2 common.Address,
) (common.Hash, common.Hash, common.Hash, error) {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.ArbRollup.SnapshotTwoStakers(
		vm.auth.getAuth(ctx),
		idx,
		stakerAddr1.ToEthAddress(),
		stakerAddr2.ToEthAddress(),
	)
	if err != nil {
		return common.Hash{}, common.Hash{}, common.Hash{}, err
	}
	receipt, err := WaitForReceiptWithResults(ctx, vm.Client, vm.auth.auth.From, tx, "SnapshotTwoStakers")
	if err != nil {
		return common.Hash{}, common.Hash{}, common.Hash{}, err
	}
	if len(receipt.Logs) != 1 {
		return common.Hash{}, common.Hash{}, common.Hash{}, errors2.New("Wrong receipt count")
	}
	event, err := vm.ArbRollup.ParseSavedTwoStakersSnapshot(*receipt.Logs[0])
	if err != nil {
		return common.Hash{}, common.Hash{}, common.Hash{}, err
	}
	return event.Location1, event.Location2, event.Snapshot, nil
}

func (vm *arbSnapshotClient) SnapshotDeadlineStakers(
	ctx context.Context,
	idx *big.Int,
	deadline common.TimeTicks,
	beforeDeadlineAddrs []common.Address,
	atOrAfterDeadlineAddrs []common.Address,
) ([]common.Hash, error) {
	vm.auth.Lock()
	defer vm.auth.Unlock()
	tx, err := vm.ArbRollup.SnapshotDeadlineStakers(
		vm.auth.getAuth(ctx),
		idx,
		deadline.Val,
		addressSliceToRaw(beforeDeadlineAddrs),
		addressSliceToRaw(atOrAfterDeadlineAddrs),
	)
	if err != nil {
		return nil, err
	}
	receipt, err := WaitForReceiptWithResults(ctx, vm.Client, vm.auth.auth.From, tx, "SnapshotDeadlineStakers")
	if err != nil {
		return nil, err
	}
	if len(receipt.Logs) != 1 {
		return nil, errors2.New("Wrong receipt count")
	}
	event, err := vm.ArbRollup.ParseSavedDeadlineStakersSnapshot(*receipt.Logs[0])
	if err != nil {
		return nil, err
	}
	return hashSliceToHashes(event.StakerLocations), nil
}
