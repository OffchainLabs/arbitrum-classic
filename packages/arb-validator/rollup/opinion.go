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
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func (chain *ChainObserver) startOpinionUpdateThread() {
	go func() {
		for {
			chain.RLock()
			for chain.nodeGraph.leaves.IsLeaf(chain.knownValidNode) {
				chain.assertionMadeCond.Wait()
			}

			currentOpinion := chain.knownValidNode
			successorHashes := [4][32]byte{}
			copy(successorHashes[:], currentOpinion.successorHashes[:])
			successor := func() *Node {
				for _, successor := range currentOpinion.successorHashes {
					if successor != zeroBytes32 {
						return chain.nodeGraph.nodeFromHash[successor]
					}
				}
				return nil
			}()

			params := successor.disputable.AssertionParams.Clone()
			claim := successor.disputable.AssertionClaim.Clone()
			claimHeight, found := chain.pendingInbox.GetHeight(claim.AfterPendingTop)
			var claimHeightCopy *big.Int
			if found {
				claimHeightCopy = new(big.Int).Set(claimHeight)
			}
			messageStack, _ := chain.pendingInbox.Substack(currentOpinion.vmProtoData.PendingTop, claim.AfterPendingTop)
			messagesVal := chain.pendingInbox.ValueForSubseq(currentOpinion.vmProtoData.PendingTop, claim.AfterPendingTop)
			prevMach := currentOpinion.machine.Clone()
			prevPendingCount := new(big.Int).Set(currentOpinion.vmProtoData.PendingCount)
			chain.RUnlock()

			newOpinion := getNodeOpinion(params, claim, prevPendingCount, claimHeightCopy, messageStack, messagesVal, prevMach)

			chain.Lock()
			correctNode, ok := chain.nodeGraph.nodeFromHash[successorHashes[newOpinion]]
			if ok {
				chain.knownValidNode = correctNode
			}
			chain.Unlock()
		}
	}()
}

func getNodeOpinion(
	params *structures.AssertionParams,
	claim *structures.AssertionClaim,
	prevPendingCount *big.Int,
	claimHeight *big.Int,
	messageStack *structures.MessageStack,
	messagesVal value.TupleValue,
	prevMach machine.Machine,
) structures.ChildType {
	correctAfterPendingTopHeight := new(big.Int).Add(prevPendingCount, params.ImportedMessageCount)
	if claimHeight == nil || correctAfterPendingTopHeight.Cmp(claimHeight) != 0 {
		return structures.InvalidPendingChildType
	}
	if messageStack.GetTopHash() != claim.ImportedMessagesSlice {
		return structures.InvalidMessagesChildType
	}

	mach := prevMach
	mach.DeliverMessages(messagesVal)
	assertion, stepsRun := mach.ExecuteAssertion(params.NumSteps, params.TimeBounds)
	if params.NumSteps != stepsRun || !claim.AssertionStub.Equals(assertion.Stub()) {
		return structures.InvalidExecutionChildType
	}

	return structures.ValidChildType
}
