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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
)

func calculatePath(from common.Hash, proof []common.Hash) common.Hash {
	node := from
	if len(proof) > 0 {
		for _, val := range proof {
			//node = keccak256(abi.encodePacked(node, proof[i]));
			node = hashing.SoliditySHA3(hashing.Bytes32(node), hashing.Bytes32(val))
		}
	}
	return node
}
