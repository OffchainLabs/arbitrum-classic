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

package rollup

import (
	"context"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type pruneParams struct {
	leafHash     common.Hash
	ancestorHash common.Hash
	leafProof    []common.Hash
	ancProof     []common.Hash
}

func (pp pruneParams) Clone() pruneParams {
	return pruneParams{
		leafHash:     pp.leafHash,
		ancestorHash: pp.ancestorHash,
		leafProof:    append(make([]common.Hash, 0), pp.leafProof...),
		ancProof:     append(make([]common.Hash, 0), pp.ancProof...),
	}
}

type recoverStakeOldParams struct {
	addr  common.Address
	proof []common.Hash
}

type recoverStakeMootedParams struct {
	addr         common.Address
	ancestorHash common.Hash
	lcProof      []common.Hash
	stProof      []common.Hash
}

func (chain *ChainObserver) startCleanupThread(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(time.Second * 5)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				chain.RLock()
				prunesToDo := chain.nodeGraph.generateNodePruneInfo()
				mootedToDo, oldToDo := chain.nodeGraph.generateStakerPruneInfo()
				chain.RUnlock()

				if len(prunesToDo) > 0 {
					for _, listener := range chain.listeners {
						listener.PrunableLeafs(chain, prunesToDo)
					}
				}
				if len(mootedToDo) > 0 {
					for _, listener := range chain.listeners {
						listener.MootableStakes(chain, mootedToDo)
					}
				}
				if len(oldToDo) > 0 {
					for _, listener := range chain.listeners {
						listener.OldStakes(chain, oldToDo)
					}
				}

			}
		}
	}()
}
