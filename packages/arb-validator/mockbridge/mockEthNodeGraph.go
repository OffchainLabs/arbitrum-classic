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

package mockbridge

import (
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
)

type NodeGraph struct {
	//vmParams
	leaves        map[common.Hash]bool
	lastConfirmed common.Hash
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

func initNodeGraph() {
	//		register for inbox
	//		init protocol state
	//		create initial node
	//		  latestConfirmedPriv = initialNode;
	//        leaves[initialNode] = true;
	//
	//        // VM parameters
	//        vmParams.gracePeriodTicks = _gracePeriodTicks;
	//        vmParams.arbGasSpeedLimitPerTick = _arbGasSpeedLimitPerTick;
	//        vmParams.maxExecutionSteps = _maxExecutionSteps;

}

func (ng *NodeGraph) pruneLeaf(from common.Hash, leafProof []common.Hash, latestConfirmedProof []common.Hash) error {
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
