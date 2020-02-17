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

package gobridge

import (
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"math/big"
)

type staker struct {
	location           common.Hash
	creationTimeBlocks *common.TimeBlocks
	inChallenge        bool
	balance            *big.Int
}

type nodeGraph struct {
	stakeRequirement *big.Int
	stakers          map[common.Address]*staker
	leaves           map[common.Hash]bool
	lastConfirmed    common.Hash
}

func newNodeGraph(auth *TransOpts, ru *rollupData) *nodeGraph {
	vmProto := hashing.SoliditySHA3(
		hashing.Bytes32(ru.vmState),
		hashing.Bytes32(value.NewEmptyTuple().Hash()),
	)
	innerHash := hashing.SoliditySHA3(
		hashing.Bytes32(vmProto),
		hashing.Uint32(0),
		hashing.Uint32(0),
		hashing.Uint32(0),
	)
	initialNode := hashing.SoliditySHA3(
		hashing.Uint32(0),
		hashing.Bytes32(innerHash),
	)

	//		register for inbox
	//
	//        // VM parameters
	//        vmParams.gracePeriodTicks = _gracePeriodTicks;
	//        vmParams.arbGasSpeedLimitPerTick = _arbGasSpeedLimitPerTick;
	//        vmParams.maxExecutionSteps = _maxExecutionSteps;
	ng := &nodeGraph{
		stakeRequirement: auth.GasPrice,
		stakers:          make(map[common.Address]*staker),
		leaves:           make(map[common.Hash]bool),
		lastConfirmed:    initialNode,
	}
	ng.leaves[initialNode] = true

	return ng
}

func (ng *nodeGraph) pruneLeaf(from common.Hash, leafProof []common.Hash, latestConfirmedProof []common.Hash) error {
	leaf := calculatePath(from, leafProof)
	if !ng.leaves[leaf] {
		return errors.New("PRUNE_LEAF invalid")
	}
	if (leafProof[0] == latestConfirmedProof[0]) ||
		(calculatePath(from, latestConfirmedProof) != ng.lastConfirmed) {
		return errors.New("PRUNE_CONFLICT")
	}
	delete(ng.leaves, leaf)
	// emit RollupPruned(leaf)

	return nil
}

type MakeAssertionData struct {
	beforeHash         common.Hash
	beforePendingTop   common.Hash
	beforePendingCount common.Hash

	prevPrevLeafHash common.Hash
	prevDeadlinTicks *big.Int
	prevDataHash     common.Hash
	prevChildType    uint32

	numSteps         uint32
	timeBoundBlocks  [2]*big.Int
	importedMsgCount *big.Int

	afterPendingTop common.Hash

	importedMessagesSlice common.Hash

	afterVMHash     common.Hash
	didInboxInsn    bool
	numArbGas       uint64
	messagesAccHash common.Hash
	logsAccHash     common.Hash
}
