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
	"log"
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

type preparedAssertion struct {
	leafHash         common.Hash
	prevPrevLeafHash common.Hash
	prevDataHash     common.Hash
	prevDeadline     common.TimeTicks
	prevChildType    valprotocol.ChildType

	beforeState *valprotocol.VMProtoData
	params      *valprotocol.AssertionParams
	claim       *valprotocol.AssertionClaim
	assertion   *protocol.ExecutionAssertion
	machine     machine.Machine
}

func (pa *preparedAssertion) Clone() *preparedAssertion {
	return &preparedAssertion{
		leafHash:         pa.leafHash,
		prevPrevLeafHash: pa.prevPrevLeafHash,
		prevDataHash:     pa.prevDataHash,
		prevDeadline:     pa.prevDeadline.Clone(),
		prevChildType:    pa.prevChildType,
		beforeState:      pa.beforeState.Clone(),
		params:           pa.params.Clone(),
		claim:            pa.claim.Clone(),
		assertion:        pa.assertion,
		machine:          pa.machine,
	}
}

func (chain *ChainObserver) startOpinionUpdateThread(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(common.NewTimeBlocksInt(2).Duration())
		assertionPreparedChan := make(chan *preparedAssertion, 20)
		preparingAssertions := make(map[common.Hash]bool)
		preparedAssertions := make(map[common.Hash]*preparedAssertion)

		updateCurrent := func() {
			currentOpinion := chain.calculatedValidNode
			currentHash := currentOpinion.hash
			log.Println("Building opinion on top of", currentHash)
			successorHashes := [4]common.Hash{}
			copy(successorHashes[:], currentOpinion.successorHashes[:])
			successor := func() *Node {
				for _, successor := range currentOpinion.successorHashes {
					if successor != zeroBytes32 {
						return chain.nodeGraph.nodeFromHash[successor]
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
			prepped, found := preparedAssertions[currentHash]

			if successor.disputable == nil {
				panic("Node was created with disputable assertion")
			}

			if found &&
				prepped.params.Equals(successor.disputable.AssertionParams) &&
				prepped.claim.Equals(successor.disputable.AssertionClaim) {
				newOpinion = valprotocol.ValidChildType
				nextMachine = prepped.machine
				validExecution = prepped.assertion
				chain.RUnlock()
			} else {
				params := successor.disputable.AssertionParams.Clone()
				claim := successor.disputable.AssertionClaim.Clone()
				prevInboxCount := new(big.Int).Set(currentOpinion.vmProtoData.InboxCount)
				afterInboxTopHeight := new(big.Int).Add(prevInboxCount, params.ImportedMessageCount)
				afterInboxTopVal, err := chain.inbox.GetHashAtIndex(afterInboxTopHeight)
				var afterInboxTop *common.Hash
				if err == nil {
					afterInboxTop = &afterInboxTopVal
				}
				inbox, _ := chain.inbox.GenerateVMInbox(currentOpinion.vmProtoData.InboxTop, params.ImportedMessageCount.Uint64())
				messagesVal := inbox.AsValue()
				nextMachine = currentOpinion.machine.Clone()

				chain.RUnlock()

				newOpinion, validExecution = getNodeOpinion(params, claim, afterInboxTop, inbox.Hash(), messagesVal, nextMachine)
			}
			// Reset prepared
			preparingAssertions = make(map[common.Hash]bool)
			preparedAssertions = make(map[common.Hash]*preparedAssertion)

			chain.RLock()
			correctNode, ok := chain.nodeGraph.nodeFromHash[successorHashes[newOpinion]]
			if ok {
				chain.RUnlock()
				chain.Lock()
				if newOpinion == valprotocol.ValidChildType {
					correctNode.machine = nextMachine
					correctNode.assertion = validExecution
				} else {
					correctNode.machine = currentOpinion.machine.Clone()
				}
				log.Println("Formed opinion that", newOpinion, successorHashes[newOpinion], "is the successor of", currentHash, "with after hash", correctNode.machine.Hash())
				chain.calculatedValidNode = correctNode
				if correctNode.depth > chain.knownValidNode.depth {
					chain.knownValidNode = correctNode
				}
				chain.Unlock()
				chain.RLock()
				if newOpinion == valprotocol.ValidChildType {
					for _, lis := range chain.listeners {
						lis.AdvancedKnownAssertion(ctx, chain, validExecution, correctNode.assertionTxHash)
					}
				}
				for _, listener := range chain.listeners {
					listener.AdvancedCalculatedValidNode(ctx, chain, correctNode.hash)
				}
			} else {
				log.Println("Formed opinion on nonexistant node", successorHashes[newOpinion])
			}
		}

		for {
			select {
			case <-ctx.Done():
				return
			case prepped := <-assertionPreparedChan:
				preparedAssertions[prepped.leafHash] = prepped
			case <-ticker.C:
				chain.RLock()
				// Catch up to current head
				for !chain.nodeGraph.leaves.IsLeaf(chain.calculatedValidNode) {
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
				_, isPreparing := preparingAssertions[chain.calculatedValidNode.hash]
				if !isPreparing {
					newMessages := chain.calculatedValidNode.vmProtoData.InboxTop != chain.inbox.GetTopHash()
					if chain.calculatedValidNode.machine != nil &&
						chain.calculatedValidNode.machine.IsBlocked(chain.latestBlockId.Height, newMessages) == nil {
						preparingAssertions[chain.calculatedValidNode.hash] = true
						go func() {
							assertionPreparedChan <- chain.prepareAssertion()
						}()
					}
				} else {
					prepared, isPrepared := preparedAssertions[chain.calculatedValidNode.hash]
					if isPrepared && chain.nodeGraph.leaves.IsLeaf(chain.calculatedValidNode) {
						startTime := prepared.params.TimeBounds.Start
						endTime := prepared.params.TimeBounds.End
						endCushion := common.NewTimeBlocks(new(big.Int).Add(chain.latestBlockId.Height.AsInt(), big.NewInt(3)))
						if chain.latestBlockId.Height.Cmp(startTime) >= 0 && endCushion.Cmp(endTime) <= 0 {
							for _, lis := range chain.listeners {
								lis.AssertionPrepared(ctx, chain, prepared.Clone())
							}
						} else {
							log.Printf("Throwing out out of date assertion with bounds [%v, %v] at time %v\n", startTime.AsInt(), endTime.AsInt(), chain.latestBlockId.Height.AsInt())
							// Prepared assertion is out of date
							delete(preparingAssertions, chain.calculatedValidNode.hash)
							delete(preparedAssertions, chain.calculatedValidNode.hash)
						}
					}
				}
				chain.RUnlock()

			}
		}
	}()
}

func (chain *ChainObserver) prepareAssertion() *preparedAssertion {
	chain.RLock()
	currentOpinion := chain.calculatedValidNode
	currentOpinionHash := currentOpinion.hash
	prevPrevLeafHash := currentOpinion.PrevHash()
	prevDataHash := currentOpinion.nodeDataHash
	prevDeadline := common.TimeTicks{new(big.Int).Set(currentOpinion.deadline.Val)}
	prevChildType := currentOpinion.linkType
	beforeState := currentOpinion.vmProtoData.Clone()
	if !chain.nodeGraph.leaves.IsLeaf(currentOpinion) {
		return nil
	}
	afterInboxTop := chain.inbox.GetTopHash()
	beforeInboxTop := beforeState.InboxTop
	newMessageCount := new(big.Int).Sub(chain.inbox.TopCount(), beforeState.InboxCount)

	inbox, _ := chain.inbox.GenerateVMInbox(beforeInboxTop, newMessageCount.Uint64())
	messagesVal := inbox.AsValue()
	mach := currentOpinion.machine.Clone()
	timeBounds := chain.currentTimeBounds()
	log.Println("timeBounds: ", timeBounds.Start.String(), timeBounds.End.String())
	maxSteps := chain.nodeGraph.params.MaxExecutionSteps
	currentHeight := chain.latestBlockId.Height.Clone()
	timeBoundsLength := new(big.Int).Sub(timeBounds.End.AsInt(), timeBounds.Start.AsInt())
	runBlocks := new(big.Int).Div(timeBoundsLength, big.NewInt(10))
	runDuration := common.NewTimeBlocks(runBlocks).Duration()
	log.Println("Asserting for up to", runBlocks, " blocks")
	chain.RUnlock()

	beforeHash := mach.Hash()

	assertion, stepsRun := mach.ExecuteAssertion(maxSteps, timeBounds, messagesVal, runDuration)

	afterHash := mach.Hash()

	blockReason := mach.IsBlocked(currentHeight, false)

	log.Printf(
		"Prepared assertion of %v steps, from %v to %v with block reason %v and timebounds [%v, %v] on top of leaf %v\n",
		stepsRun,
		beforeHash,
		afterHash,
		blockReason,
		timeBounds.Start.AsInt(),
		timeBounds.End.AsInt(),
		currentOpinionHash,
	)

	var params *valprotocol.AssertionParams
	var claim *valprotocol.AssertionClaim
	if assertion.DidInboxInsn {
		params = &valprotocol.AssertionParams{
			NumSteps:             stepsRun,
			TimeBounds:           timeBounds,
			ImportedMessageCount: newMessageCount,
		}
		claim = &valprotocol.AssertionClaim{
			AfterInboxTop:         afterInboxTop,
			ImportedMessagesSlice: inbox.Hash(),
			AssertionStub:         valprotocol.NewExecutionAssertionStubFromAssertion(assertion),
		}
	} else {
		params = &valprotocol.AssertionParams{
			NumSteps:             stepsRun,
			TimeBounds:           timeBounds,
			ImportedMessageCount: big.NewInt(0),
		}
		claim = &valprotocol.AssertionClaim{
			AfterInboxTop:         beforeInboxTop,
			ImportedMessagesSlice: value.NewEmptyTuple().Hash(),
			AssertionStub:         valprotocol.NewExecutionAssertionStubFromAssertion(assertion),
		}
	}
	return &preparedAssertion{
		leafHash:         currentOpinionHash,
		prevPrevLeafHash: prevPrevLeafHash,
		prevDataHash:     prevDataHash,
		prevDeadline:     prevDeadline,
		prevChildType:    prevChildType,
		beforeState:      beforeState,
		params:           params,
		claim:            claim,
		assertion:        assertion,
		machine:          mach,
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
		return valprotocol.InvalidInboxTopChildType, nil
	}
	if calculatedMessagesSlice != claim.ImportedMessagesSlice {
		return valprotocol.InvalidMessagesChildType, nil
	}

	assertion, stepsRun := mach.ExecuteAssertion(
		params.NumSteps,
		params.TimeBounds,
		messagesVal,
		0,
	)
	if params.NumSteps != stepsRun || !claim.AssertionStub.Equals(valprotocol.NewExecutionAssertionStubFromAssertion(assertion)) {
		return valprotocol.InvalidExecutionChildType, nil
	}

	return valprotocol.ValidChildType, assertion
}
