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

func (chain *ChainObserver) startConfirmThread(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				chain.RLock()
				if !chain.atHead {
					chain.RUnlock()
					break
				}
				confOpp := chain.nodeGraph.generateNextConfProof(common.TimeFromBlockNum(chain.latestBlockId.Height))
				if confOpp != nil {
					for _, listener := range chain.listeners {
						listener.NodesConfirmable(ctx, chain, confOpp)
					}
				}
				chain.RUnlock()
			}
		}
	}()
}
