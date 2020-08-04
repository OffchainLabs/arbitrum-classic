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

package challenges

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type ExecutionChallengeInfo struct {
	isDiscontinueType bool
	challengeRounds   int
	currentRound      int
}

func StandardExecutionChallenge() ExecutionChallengeInfo {
	return ExecutionChallengeInfo{
		false,
		0,
		0,
	}
}

func ContinueChallenge(typeReq ExecutionChallengeInfo) bool {
	if !typeReq.isDiscontinueType {
		return true
	} else {
		if typeReq.challengeRounds == typeReq.currentRound {
			return false
		} else {
			return true
		}
	}
}

func challengeEnded(state ChallengeState, err error) bool {
	if err != nil || state != ChallengeContinuing {
		return true
	} else {
		return false
	}
}

func findSegmentToChallenge(
	validatorHashes []common.Hash,
	chainHashes []common.Hash) (uint64, bool) {
	// If any inbox segment is wrong, we can easily win
	for i := uint64(1); i < uint64(len(validatorHashes)); i++ {
		if validatorHashes[i] != chainHashes[i] {
			return i - 1, true
		}
	}

	return 0, false
}
