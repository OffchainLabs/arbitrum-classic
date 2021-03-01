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

package valprotocol

import (
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

func InboxTopChallengeDataHash(
	lowerInbox common.Hash,
	upperInbox common.Hash,
	messageCount *big.Int,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(lowerInbox),
		hashing.Bytes32(upperInbox),
		hashing.Uint256(messageCount),
	)
}

func MessageChallengeDataHash(
	lowerInbox common.Hash,
	upperInbox common.Hash,
	lowerMessages common.Hash,
	upperMessages common.Hash,
	messageCount *big.Int,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(lowerInbox),
		hashing.Bytes32(upperInbox),
		hashing.Bytes32(lowerMessages),
		hashing.Bytes32(upperMessages),
		hashing.Uint256(messageCount),
	)
}

func ExecutionDataHash(
	numSteps uint64,
	assertion *ExecutionAssertionStub,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint64(numSteps),
		hashing.Uint64(assertion.NumGas),
		hashing.Bytes32(assertion.BeforeMachineHash),
		hashing.Bytes32(assertion.AfterMachineHash),
		hashing.Bytes32(assertion.BeforeInboxAcc),
		hashing.Bytes32(assertion.AfterInboxAcc),
		hashing.Bytes32(assertion.FirstMessageHash),
		hashing.Bytes32(assertion.LastMessageHash),
		hashing.Uint64(assertion.MessageCount),
		hashing.Bytes32(assertion.FirstLogHash),
		hashing.Bytes32(assertion.LastLogHash),
		hashing.Uint64(assertion.LogCount),
	)
}
