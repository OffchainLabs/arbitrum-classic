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

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"

	"github.com/ethereum/go-ethereum/common"
)

type pruneParams struct {
	leaf      *Node
	ancestor  *Node
	leafProof [][32]byte
	ancProof  [][32]byte
}

type recoverStakeOldParams struct {
	addr  common.Address
	proof [][32]byte
}

type recoverStakeMootedParams struct {
	addr     common.Address
	ancestor *Node
	lcProof  [][32]byte
	stProof  [][32]byte
}

func (chain *ChainObserver) startCleanupThread(client *ethbridge.ArbRollup, doneChan chan interface{}) {
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
				chain.RLock()
				prunesToDo := chain.nodeGraph.generateNodePruneInfo()
				mootedToDo, oldToDo := chain.nodeGraph.generateStakerPruneInfo()
				chain.RUnlock()

				for _, prune := range prunesToDo {
					client.PruneLeaf(
						context.TODO(),
						prune.ancestor.hash,
						prune.leafProof,
						prune.ancProof,
					)
				}
				for _, moot := range mootedToDo {
					client.RecoverStakeMooted(
						context.TODO(),
						moot.ancestor.hash,
						moot.addr,
						moot.lcProof,
						moot.stProof,
					)
				}
				for _, old := range oldToDo {
					client.RecoverStakeOld(
						context.TODO(),
						old.addr,
						old.proof,
					)
				}
			}
		}
	}()
}
