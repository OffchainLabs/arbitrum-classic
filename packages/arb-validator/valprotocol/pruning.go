/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

import "github.com/offchainlabs/arbitrum/packages/arb-util/common"

type PruneParams struct {
	LeafHash     common.Hash
	AncestorHash common.Hash
	LeafProof    []common.Hash
	AncProof     []common.Hash
}

func (pp PruneParams) Clone() PruneParams {
	return PruneParams{
		LeafHash:     pp.LeafHash,
		AncestorHash: pp.AncestorHash,
		LeafProof:    append(make([]common.Hash, 0), pp.LeafProof...),
		AncProof:     append(make([]common.Hash, 0), pp.AncProof...),
	}
}
