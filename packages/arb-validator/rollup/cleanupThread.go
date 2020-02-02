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

func (co *ChainObserver) startCleanupThread(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(common.NewTimeBlocksInt(2).Duration())
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				co.RLock()
				if !co.atHead {
					co.RUnlock()
					break
				}
				prunesToDo := co.nodeGraph.generateNodePruneInfo()
				mootedToDo, oldToDo := co.nodeGraph.generateStakerPruneInfo()
				co.RUnlock()

				if len(prunesToDo) > 0 {
					for _, listener := range co.listeners {
						listener.PrunableLeafs(ctx, co, prunesToDo)
					}
				}
				if len(mootedToDo) > 0 {
					for _, listener := range co.listeners {
						listener.MootableStakes(ctx, co, mootedToDo)
					}
				}
				if len(oldToDo) > 0 {
					for _, listener := range co.listeners {
						listener.OldStakes(ctx, co, oldToDo)
					}
				}
			}
		}
	}()
}
