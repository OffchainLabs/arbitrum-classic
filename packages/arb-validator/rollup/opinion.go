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
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"math/big"
)

func (chain *ChainObserver) startOpinionUpdateThread() {
	go func() {
		for {
			chain.RLock()
			for chain.nodeGraph.leaves.IsLeaf(chain.knownValidNode) {
				chain.assertionMadeCond.Wait()
			}

			currentOpinion := chain.knownValidNode
			claim := currentOpinion.disputable.AssertionClaim
			claimHeight, found := chain.pendingInbox.GetHeight(claim.AfterPendingTop)
			var successors [structures.MaxChildType + 1]*Node
			for i, sh := range currentOpinion.successorHashes {
				successors[i] = chain.nodeGraph.nodeFromHash[sh]
			}
			messageStack, _ := chain.pendingInbox.Substack(currentOpinion.prev.vmProtoData.PendingTop, claim.AfterPendingTop)
			topHash := messageStack.GetTopHash()
			messagesVal := chain.pendingInbox.ValueForSubseq(currentOpinion.prev.vmProtoData.PendingTop, claim.AfterPendingTop)
			prevMach := currentOpinion.prev.machine.Clone()
			chain.RUnlock()

			newOpinion := updateNodeOpinion(currentOpinion, claimHeight, found, successors, topHash, messagesVal, prevMach)

			chain.Lock()
			if newOpinion.depth > chain.nodeGraph.latestConfirmed.depth {
				chain.knownValidNode = newOpinion
			}
			chain.Unlock()
		}
	}()
}

func updateNodeOpinion(
	currentOpinion *Node,
	claimHeight *big.Int,
	found bool,
	successors [structures.MaxChildType + 1]*Node,
	topHash [32]byte,
	messagesVal value.TupleValue,
	prevMach machine.Machine,
) *Node {
	params := currentOpinion.disputable.AssertionParams
	claim := currentOpinion.disputable.AssertionClaim
	correctAfterPendingTopHeight := new(big.Int).Add(currentOpinion.prev.vmProtoData.PendingCount, params.ImportedMessageCount)
	if !found || correctAfterPendingTopHeight.Cmp(claimHeight) != 0 {
		return successors[structures.InvalidPendingChildType]
	}
	if topHash != claim.ImportedMessagesSlice {
		return successors[structures.InvalidMessagesChildType]
	}

	mach := currentOpinion.prev.machine.Clone()
	mach.DeliverMessages(messagesVal)
	assertion, stepsRun := mach.ExecuteAssertion(params.NumSteps, params.TimeBounds)
	if params.NumSteps != stepsRun || !claim.AssertionStub.Equals(assertion.Stub()) {
		return successors[structures.InvalidExecutionChildType]
	}

	return successors[structures.ValidChildType]
}
