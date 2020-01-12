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

	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func PendingTopChallengeDataHash(
	lowerPending [32]byte,
	upperPending [32]byte,
	messageCount *big.Int,
) [32]byte {
	ret := [32]byte{}
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(lowerPending),
		solsha3.Bytes32(upperPending),
		solsha3.Uint256(messageCount),
	))
	return ret
}

func MessageChallengeDataHash(
	lowerPending [32]byte,
	upperPending [32]byte,
	lowerMessages [32]byte,
	upperMessages [32]byte,
	messageCount *big.Int,
) [32]byte {
	ret := [32]byte{}
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(lowerPending),
		solsha3.Bytes32(upperPending),
		solsha3.Bytes32(lowerMessages),
		solsha3.Bytes32(upperMessages),
		solsha3.Uint256(messageCount),
	))
	return ret
}

func ExecutionDataHash(
	numSteps uint32,
	preconditionHash [32]byte,
	assertionHash [32]byte,
) [32]byte {
	ret := [32]byte{}
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Uint32(numSteps),
		solsha3.Bytes32(preconditionHash),
		solsha3.Bytes32(assertionHash),
	))
	return ret
}
