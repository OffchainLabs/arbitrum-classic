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

package structures

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
)

func CalculateBisectionStepCount(chunkIndex, segmentCount, totalSteps uint64) uint64 {
	if chunkIndex == 0 {
		return totalSteps/segmentCount + totalSteps%segmentCount
	} else {
		return totalSteps / segmentCount
	}
}

func PendingTopChallengeDataHash(
	lowerPending common.Hash,
	upperPending common.Hash,
	messageCount *big.Int,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(lowerPending),
		hashing.Bytes32(upperPending),
		hashing.Uint256(messageCount),
	)
}

func MessageChallengeDataHash(
	lowerPending common.Hash,
	upperPending common.Hash,
	lowerMessages common.Hash,
	upperMessages common.Hash,
	messageCount *big.Int,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(lowerPending),
		hashing.Bytes32(upperPending),
		hashing.Bytes32(lowerMessages),
		hashing.Bytes32(upperMessages),
		hashing.Uint256(messageCount),
	)
}

func ExecutionPreconditionHash(machineHash common.Hash, timeBounds *protocol.TimeBoundsBlocks, msgSlices common.Hash) common.Hash {
	pre := &valprotocol.Precondition{
		BeforeHash:  machineHash,
		TimeBounds:  timeBounds,
		BeforeInbox: value.NewHashOnlyValue(msgSlices, 0),
	}
	return pre.Hash()
}

func ExecutionDataHash(
	numSteps uint64,
	preconditionHash common.Hash,
	assertionHash common.Hash,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint64(numSteps),
		hashing.Bytes32(preconditionHash),
		hashing.Bytes32(assertionHash),
	)
}

func NodeHash(prevHash common.Hash,
	protoHash common.Hash,
	deadline common.TimeTicks,
	nodeDataHash common.Hash,
	linkType ChildType) (common.Hash, common.Hash) {
	innerHash := hashing.SoliditySHA3(
		hashing.Bytes32(protoHash),
		hashing.TimeTicks(deadline),
		hashing.Bytes32(nodeDataHash),
		hashing.Uint256(new(big.Int).SetUint64(uint64(linkType))),
	)
	hash := hashing.SoliditySHA3(
		hashing.Bytes32(prevHash),
		hashing.Bytes32(innerHash),
	)
	return hash, innerHash
}
