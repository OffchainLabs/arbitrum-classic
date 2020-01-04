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

import "time"

type pruneParams struct {
	leaf      *Node
	ancestor  *Node
	leafProof [][32]byte
	ancProof  [][32]byte
}

func (chain *StakedNodeGraph) startCleanupThread(doneChan chan interface{}) {
	if doneChan == nil {
		doneChan = make(chan interface{})
	}
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-doneChan:
				return
			case <-ticker.C:
				prunesToDo := []pruneParams{}
				chain.RLock()
				chain.leaves.forall(func(leaf *Node) {
					ancestor, _, err := chain.GetConflictAncestor(leaf, chain.latestConfirmed)
					if err == nil && ancestor != nil {
						prunesToDo = append(prunesToDo, pruneParams{
							leaf,
							ancestor,
							GeneratePathProof(ancestor, leaf),
							GeneratePathProof(ancestor, chain.latestConfirmed),
						})
					}
				})
				chain.Unlock()
				for _, prune := range prunesToDo {
					_ = prune
					//TODO: call contract's PruneLeaf method with params from prune
				}
			}
		}
	}()
}
