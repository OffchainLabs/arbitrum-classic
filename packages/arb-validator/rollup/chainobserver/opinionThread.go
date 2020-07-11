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

package chainobserver

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/nodegraph"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup/chainlistener"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

func (chain *ChainObserver) startOpinionUpdateThread(ctx context.Context) {
	go func() {
		preparingAssertions := make(map[common.Hash]bool)
		preparedAssertions := make(map[common.Hash]*chainlistener.PreparedAssertion)
		preparedAssertionsMut := new(sync.Mutex)

		updateCurrent := func() {
			currentOpinion := chain.calculatedValidNode
			currentHash := currentOpinion.Hash()
			log.Println("Building opinion on top of", currentHash)
			successorHashes := currentOpinion.SuccessorHashes()
			successor := func() *structures.Node {
				for _, successor := range successorHashes {
					if successor != nodegraph.ZeroBytes32 {
						return chain.NodeGraph.NodeFromHash(successor)
					}
				}
				return nil
			}()

			if successor == nil {
				panic("Node has no successor")
			}

			var newOpinion valprotocol.ChildType
			var nextMachine machine.Machine
			var validExecution *protocol.ExecutionAssertion
			preparedAssertionsMut.Lock()
			prepped, found := preparedAssertions[currentHash]
			preparedAssertionsMut.Unlock()
			disputable := successor.Disputable()

			if disputable == nil {
				panic("Node was created with disputable assertion")
			}

			if found &&
				prepped.Params.Equals(disputable.AssertionParams) &&
				prepped.Claim.Equals(disputable.AssertionClaim) {
				newOpinion = valprotocol.ValidChildType
				nextMachine = prepped.Machine
				validExecution = prepped.Assertion
				chain.RUnlock()
			} else {
				params := disputable.AssertionParams.Clone()
				claim := disputable.AssertionClaim.Clone()
				prevInboxCount := new(big.Int).Set(currentOpinion.VMProtoData().InboxCount)
				afterInboxTopHeight := new(big.Int).Add(prevInboxCount, params.ImportedMessageCount)
				afterInboxTopVal, err := chain.Inbox.GetHashAtIndex(afterInboxTopHeight)
				var afterInboxTop *common.Hash
				if err == nil {
					afterInboxTop = &afterInboxTopVal
				}
				inbox, _ := chain.Inbox.GenerateVMInbox(currentOpinion.VMProtoData().InboxTop, params.ImportedMessageCount.Uint64())
				messages, _ := chain.Inbox.GetMessages(currentOpinion.VMProtoData().InboxTop, params.ImportedMessageCount.Uint64())
				messagesVal := inbox.AsValue()
				nextMachine = currentOpinion.Machine().Clone()
				log.Println("Forming opinion on", successor.Hash().ShortString(), "which imported", messages, "messages")

				chain.RUnlock()

				newOpinion, validExecution = getNodeOpinion(params, claim, afterInboxTop, inbox.Hash().Hash(), messagesVal, nextMachine)
			}
			// Reset prepared
			log.Println("Resetting prepped assertions")
			preparingAssertions = make(map[common.Hash]bool)
			preparedAssertionsMut.Lock()
			preparedAssertions = make(map[common.Hash]*chainlistener.PreparedAssertion)
			preparedAssertionsMut.Unlock()
			chain.RLock()
			correctNode := chain.NodeGraph.GetSuccessor(currentOpinion, newOpinion)
			if correctNode != nil {
				chain.RUnlock()
				chain.Lock()
				if newOpinion == valprotocol.ValidChildType {
					_ = correctNode.UpdateValidOpinion(nextMachine, validExecution)
				} else {
					_ = correctNode.UpdateInvalidOpinion()
				}
				log.Println("Formed opinion that", newOpinion, successorHashes[newOpinion], "is the successor of", currentHash, "with after hash", correctNode.Machine().Hash())
				chain.calculatedValidNode = correctNode
				if correctNode.Depth() > chain.KnownValidNode.Depth() {
					chain.KnownValidNode = correctNode
				}
				chain.Unlock()
				chain.RLock()
				for _, listener := range chain.listeners {
					listener.AdvancedKnownNode(ctx, chain.NodeGraph, correctNode)
				}
			} else {
				log.Println("Formed opinion on nonexistant node", successorHashes[newOpinion])
			}
		}

		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				chain.RLock()
				// Catch up to current head
				for !chain.NodeGraph.Leaves().IsLeaf(chain.calculatedValidNode) {
					updateCurrent()
					chain.RUnlock()
					select {
					case <-ctx.Done():
						return
					default:
					}
					chain.RLock()
				}
				if !chain.atHead {
					chain.RUnlock()
					break
				}
				// Prepare next assertion
				_, isPreparing := preparingAssertions[chain.calculatedValidNode.Hash()]
				log.Println("opinion1", chain.calculatedValidNode.Hash().ShortString(), isPreparing)
				if !isPreparing {
					newMessages := chain.calculatedValidNode.VMProtoData().InboxTop != chain.Inbox.GetTopHash()
					if chain.calculatedValidNode.Machine() != nil &&
						chain.calculatedValidNode.Machine().IsBlocked(newMessages) == nil {
						preparingAssertions[chain.calculatedValidNode.Hash()] = true
						go func() {
							prepped := chain.PrepareAssertion()
							log.Println("prepped assertion", prepped.Prev.Hash().ShortString())
							preparedAssertionsMut.Lock()
							preparedAssertions[prepped.Prev.Hash()] = prepped
							preparedAssertionsMut.Unlock()
							log.Println("saved prepped assertion", prepped.Prev.Hash().ShortString())
						}()
					}
				} else {
					log.Println("Checking for prepped assertion on", chain.calculatedValidNode.Hash().ShortString())
					preparedAssertionsMut.Lock()
					prepared, isPrepared := preparedAssertions[chain.calculatedValidNode.Hash()]
					preparedAssertionsMut.Unlock()
					log.Println("opinion2", isPrepared, chain.NodeGraph.Leaves().IsLeaf(chain.calculatedValidNode))
					if isPrepared && chain.NodeGraph.Leaves().IsLeaf(chain.calculatedValidNode) {
						chain.RUnlock()
						chain.Lock()
						chain.pendingState = prepared.Machine
						chain.Unlock()
						chain.RLock()
						log.Println("opinion3", len(chain.listeners))
						for _, lis := range chain.listeners {
							lis.AssertionPrepared(
								ctx,
								chain.GetChainParams(),
								chain.NodeGraph,
								chain.KnownValidNode,
								chain.LatestBlockId,
								prepared.Clone())
						}
					}
				}
				chain.RUnlock()

			}
		}
	}()
}

func (chain *ChainObserver) PrepareAssertion() *chainlistener.PreparedAssertion {
	chain.RLock()
	currentOpinion := chain.calculatedValidNode
	beforeState := currentOpinion.VMProtoData().Clone()
	if !chain.NodeGraph.Leaves().IsLeaf(currentOpinion) {
		return nil
	}
	afterInboxTop := chain.Inbox.GetTopHash()
	beforeInboxTop := beforeState.InboxTop
	newMessageCount := new(big.Int).Sub(chain.Inbox.TopCount(), beforeState.InboxCount)

	inbox, _ := chain.Inbox.GenerateVMInbox(beforeInboxTop, newMessageCount.Uint64())
	messagesVal := inbox.AsValue()
	mach := currentOpinion.Machine().Clone()
	maxSteps := chain.NodeGraph.Params().MaxExecutionSteps
	chain.RUnlock()

	beforeHash := mach.Hash()

	assertion, stepsRun := mach.ExecuteAssertion(maxSteps, messagesVal, 0)

	afterHash := mach.Hash()

	blockReason := mach.IsBlocked(false)

	log.Printf(
		"Prepared assertion of %v steps, from %v to %v with block reason %v on top of leaf %v\n",
		stepsRun,
		beforeHash,
		afterHash,
		blockReason,
		currentOpinion.Hash(),
	)

	var params *valprotocol.AssertionParams
	var claim *valprotocol.AssertionClaim
	stub := valprotocol.NewExecutionAssertionStubFromAssertion(assertion)
	if assertion.DidInboxInsn {
		params = &valprotocol.AssertionParams{
			NumSteps:             stepsRun,
			ImportedMessageCount: newMessageCount,
		}
		claim = &valprotocol.AssertionClaim{
			AfterInboxTop:         afterInboxTop,
			ImportedMessagesSlice: inbox.Hash().Hash(),
			AssertionStub:         stub,
		}
	} else {
		params = &valprotocol.AssertionParams{
			NumSteps:             stepsRun,
			ImportedMessageCount: big.NewInt(0),
		}
		claim = &valprotocol.AssertionClaim{
			AfterInboxTop:         beforeInboxTop,
			ImportedMessagesSlice: value.NewEmptyTuple().Hash(),
			AssertionStub:         stub,
		}
	}
	return &chainlistener.PreparedAssertion{
		Prev:        currentOpinion,
		BeforeState: beforeState,
		Params:      params,
		Claim:       claim,
		Assertion:   assertion,
		Machine:     mach,
	}
}

func getNodeOpinion(
	params *valprotocol.AssertionParams,
	claim *valprotocol.AssertionClaim,
	afterInboxTop *common.Hash,
	calculatedMessagesSlice common.Hash,
	messagesVal value.TupleValue,
	mach machine.Machine,
) (valprotocol.ChildType, *protocol.ExecutionAssertion) {
	if afterInboxTop == nil || claim.AfterInboxTop != *afterInboxTop {
		log.Println("Saw node with invalid after inbox top claim", claim.AfterInboxTop)
		return valprotocol.InvalidInboxTopChildType, nil
	}
	if calculatedMessagesSlice != claim.ImportedMessagesSlice {
		log.Println("Saw node with invalid imported messages claim", claim.ImportedMessagesSlice)
		return valprotocol.InvalidMessagesChildType, nil
	}

	assertion, stepsRun := mach.ExecuteAssertion(
		params.NumSteps,
		messagesVal,
		0,
	)
	if params.NumSteps != stepsRun || !claim.AssertionStub.Equals(valprotocol.NewExecutionAssertionStubFromAssertion(assertion)) {
		log.Println("Saw node with invalid execution claim")
		return valprotocol.InvalidExecutionChildType, nil
	}

	return valprotocol.ValidChildType, assertion
}
