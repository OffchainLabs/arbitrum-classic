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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

type preparedAssertion struct {
	leafHash         common.Hash
	prevPrevLeafHash common.Hash
	prevDataHash     common.Hash
	prevDeadline     common.TimeTicks
	prevChildType    structures.ChildType

	beforeState *structures.VMProtoData
	params      *structures.AssertionParams
	claim       *structures.AssertionClaim
	assertion   *protocol.ExecutionAssertion
	machine     machine.Machine
}

func (pa *preparedAssertion) Clone() *preparedAssertion {
	return &preparedAssertion{
		leafHash:         pa.leafHash,
		prevPrevLeafHash: pa.prevPrevLeafHash,
		prevDataHash:     pa.prevDataHash,
		prevDeadline:     common.TimeTicks{Val: new(big.Int).Set(pa.prevDeadline.Val)},
		prevChildType:    pa.prevChildType,
		beforeState:      pa.beforeState.Clone(),
		params:           pa.params.Clone(),
		claim:            pa.claim.Clone(),
		assertion:        pa.assertion,
		machine:          pa.machine,
	}
}

func (co *ChainObserver) startOpinionUpdateThread(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(common.NewTimeBlocksInt(4).Duration())
		assertionPreparedChan := make(chan *preparedAssertion, 20)
		preparingAssertions := make(map[common.Hash]bool)
		preparedAssertions := make(map[common.Hash]*preparedAssertion)

		updateCurrent := func() {
			currentOpinion := co.calculatedValidNode
			currentHash := currentOpinion.hash
			log.Println("Building opinion on top of", currentHash)
			successorHashes := [4]common.Hash{}
			copy(successorHashes[:], currentOpinion.successorHashes[:])
			successor := func() *Node {
				for _, successor := range currentOpinion.successorHashes {
					if successor != zeroBytes32 {
						return co.nodeGraph.nodeFromHash[successor]
					}
				}
				return nil
			}()

			if successor == nil {
				panic("Node has no successor")
			}

			var newOpinion structures.ChildType
			var nextMachine machine.Machine
			var validExecution *protocol.ExecutionAssertion
			prepped, found := preparedAssertions[currentHash]

			if successor.disputable == nil {
				panic("Node was created with disputable assertion")
			}

			if found &&
				prepped.params.Equals(successor.disputable.AssertionParams) &&
				prepped.claim.Equals(successor.disputable.AssertionClaim) {
				newOpinion = structures.ValidChildType
				nextMachine = prepped.machine
				validExecution = prepped.assertion
				co.RUnlock()
			} else {
				params := successor.disputable.AssertionParams.Clone()
				claim := successor.disputable.AssertionClaim.Clone()
				prevPendingCount := new(big.Int).Set(currentOpinion.vmProtoData.PendingCount)
				afterPendingTopHeight := new(big.Int).Add(prevPendingCount, params.ImportedMessageCount)
				afterPendingTopVal, err := co.pendingInbox.GetHashAtIndex(afterPendingTopHeight)
				var afterPendingTop *common.Hash
				if err == nil {
					afterPendingTop = &afterPendingTopVal
				}
				inbox, _ := co.pendingInbox.GenerateInbox(currentOpinion.vmProtoData.PendingTop, params.ImportedMessageCount.Uint64())
				messagesVal := inbox.AsValue()
				nextMachine = currentOpinion.machine.Clone()

				co.RUnlock()

				newOpinion, validExecution = getNodeOpinion(params, claim, afterPendingTop, inbox.Hash(), messagesVal, nextMachine)
			}
			// Reset prepared
			preparingAssertions = make(map[common.Hash]bool)
			preparedAssertions = make(map[common.Hash]*preparedAssertion)

			co.RLock()
			correctNode, ok := co.nodeGraph.nodeFromHash[successorHashes[newOpinion]]
			if ok {
				co.RUnlock()
				co.Lock()
				if newOpinion == structures.ValidChildType {
					correctNode.machine = nextMachine
					correctNode.assertion = validExecution
				} else {
					correctNode.machine = currentOpinion.machine.Clone()
				}
				log.Println("Formed opinion that", newOpinion, successorHashes[newOpinion], "is the successor of", currentHash, "with after hash", correctNode.machine.Hash())
				co.calculatedValidNode = correctNode
				if correctNode.depth > co.knownValidNode.depth {
					co.knownValidNode = correctNode
				}
				co.Unlock()
				co.RLock()
				if newOpinion == structures.ValidChildType {
					for _, lis := range co.listeners {
						lis.AdvancedKnownAssertion(ctx, co, validExecution, correctNode.assertionTxHash)
					}
				}
				for _, listener := range co.listeners {
					listener.AdvancedCalculatedValidNode(ctx, co, correctNode.hash)
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
				co.RLock()
				// Catch up to current head
				for !co.nodeGraph.leaves.IsLeaf(co.calculatedValidNode) {
					updateCurrent()
					co.RUnlock()
					select {
					case <-ctx.Done():
						return
					default:
					}
					co.RLock()
				}
				if !co.atHead {
					co.RUnlock()
					break
				}
				// Prepare next assertion
				_, isPreparing := preparingAssertions[co.calculatedValidNode.hash]
				if !isPreparing {
					newMessages := co.calculatedValidNode.vmProtoData.PendingTop != co.pendingInbox.GetTopHash()
					if co.calculatedValidNode.machine != nil &&
						co.calculatedValidNode.machine.IsBlocked(co.latestBlockID.Height, newMessages) == nil {
						preparingAssertions[co.calculatedValidNode.hash] = true
						go func() {
							assertionPreparedChan <- co.prepareAssertion()
						}()
					}
				} else {
					prepared, isPrepared := preparedAssertions[co.calculatedValidNode.hash]
					if isPrepared && co.nodeGraph.leaves.IsLeaf(co.calculatedValidNode) {
						startTime := prepared.params.TimeBounds.Start
						endTime := prepared.params.TimeBounds.End
						endCushion := common.NewTimeBlocks(new(big.Int).Add(co.latestBlockID.Height.AsInt(), big.NewInt(3)))
						if co.latestBlockID.Height.Cmp(startTime) >= 0 && endCushion.Cmp(endTime) <= 0 {
							for _, lis := range co.listeners {
								lis.AssertionPrepared(ctx, co, prepared.Clone())
							}
						} else {
							log.Printf("Throwing out out of date assertion with bounds [%v, %v] at time %v\n", startTime.AsInt(), endTime.AsInt(), co.latestBlockID.Height.AsInt())
							// Prepared assertion is out of date
							delete(preparingAssertions, co.calculatedValidNode.hash)
							delete(preparedAssertions, co.calculatedValidNode.hash)
						}
					}
				}
				co.RUnlock()
			}
		}
	}()
}

func (co *ChainObserver) prepareAssertion() *preparedAssertion {
	co.RLock()
	currentOpinion := co.calculatedValidNode
	currentOpinionHash := currentOpinion.hash
	prevPrevLeafHash := currentOpinion.PrevHash()
	prevDataHash := currentOpinion.nodeDataHash
	prevDeadline := common.TimeTicks{Val: new(big.Int).Set(currentOpinion.deadline.Val)}
	prevChildType := currentOpinion.linkType
	beforeState := currentOpinion.vmProtoData.Clone()
	if !co.nodeGraph.leaves.IsLeaf(currentOpinion) {
		return nil
	}
	afterPendingTop := co.pendingInbox.GetTopHash()
	beforePendingTop := beforeState.PendingTop
	newMessageCount := new(big.Int).Sub(co.pendingInbox.TopCount(), beforeState.PendingCount)

	inbox, _ := co.pendingInbox.GenerateInbox(beforePendingTop, newMessageCount.Uint64())
	messagesVal := inbox.AsValue()
	mach := currentOpinion.machine.Clone()

	timeBounds := co.currentTimeBounds()
	maxSteps := co.nodeGraph.params.MaxExecutionSteps
	currentHeight := co.latestBlockID.Height.Clone()
	timeBoundsLength := new(big.Int).Sub(timeBounds.End.AsInt(), timeBounds.Start.AsInt())
	runBlocks := new(big.Int).Div(timeBoundsLength, big.NewInt(10))
	runTicks := common.TimeFromBlockNum(common.NewTimeBlocks(runBlocks))
	log.Println("Asserting for up to", runTicks.Duration().Seconds(), "seconds")
	co.RUnlock()

	beforeHash := mach.Hash()

	assertion, stepsRun := mach.ExecuteAssertion(maxSteps, timeBounds, messagesVal, runTicks.Duration())

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

	var params *structures.AssertionParams
	var claim *structures.AssertionClaim
	if assertion.DidInboxInsn {
		params = &structures.AssertionParams{
			NumSteps:             stepsRun,
			TimeBounds:           timeBounds,
			ImportedMessageCount: newMessageCount,
		}
		claim = &structures.AssertionClaim{
			AfterPendingTop:       afterPendingTop,
			ImportedMessagesSlice: inbox.Hash(),
			AssertionStub:         valprotocol.NewExecutionAssertionStubFromAssertion(assertion),
		}
	} else {
		params = &structures.AssertionParams{
			NumSteps:             stepsRun,
			TimeBounds:           timeBounds,
			ImportedMessageCount: big.NewInt(0),
		}
		claim = &structures.AssertionClaim{
			AfterPendingTop:       beforePendingTop,
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
	params *structures.AssertionParams,
	claim *structures.AssertionClaim,
	afterPendingTop *common.Hash,
	calculatedMessagesSlice common.Hash,
	messagesVal value.TupleValue,
	mach machine.Machine,
) (structures.ChildType, *protocol.ExecutionAssertion) {
	if afterPendingTop == nil || claim.AfterPendingTop != *afterPendingTop {
		return structures.InvalidPendingChildType, nil
	}
	if calculatedMessagesSlice != claim.ImportedMessagesSlice {
		return structures.InvalidMessagesChildType, nil
	}

	assertion, stepsRun := mach.ExecuteAssertion(
		params.NumSteps,
		params.TimeBounds,
		messagesVal,
		0,
	)
	if params.NumSteps != stepsRun || !claim.AssertionStub.Equals(valprotocol.NewExecutionAssertionStubFromAssertion(assertion)) {
		return structures.InvalidExecutionChildType, nil
	}

	return structures.ValidChildType, assertion
}
