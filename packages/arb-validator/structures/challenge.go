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
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func CalculateBisectionStepCount(chunkIndex, segmentCount, totalSteps uint32) uint32 {
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
	ret := common.Hash{}
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(lowerPending.Bytes()),
		solsha3.Bytes32(upperPending.Bytes()),
		solsha3.Uint256(messageCount),
	))
	return ret
}

func MessageChallengeDataHash(
	lowerPending common.Hash,
	upperPending common.Hash,
	lowerMessages common.Hash,
	upperMessages common.Hash,
	messageCount *big.Int,
) common.Hash {
	ret := common.Hash{}
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(lowerPending.Bytes()),
		solsha3.Bytes32(upperPending.Bytes()),
		solsha3.Bytes32(lowerMessages.Bytes()),
		solsha3.Bytes32(upperMessages.Bytes()),
		solsha3.Uint256(messageCount),
	))
	return ret
}

func ExecutionDataHash(
	numSteps uint32,
	preconditionHash common.Hash,
	assertionHash common.Hash,
) common.Hash {
	ret := common.Hash{}
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Uint32(numSteps),
		solsha3.Bytes32(preconditionHash.Bytes()),
		solsha3.Bytes32(assertionHash.Bytes()),
	))
	return ret
}
