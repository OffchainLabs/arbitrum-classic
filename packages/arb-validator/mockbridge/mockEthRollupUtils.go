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
)

func calculatePath(from common.Hash, proof []common.Hash) common.Hash {
	return calculatePathOffset(from, proof, 0, len(proof))
}

func calculatePathOffset(from common.Hash, proof []common.Hash, start int, end int) common.Hash {
	node := from
	for i := start; i <= end; i++ {

	}
	return node
}
