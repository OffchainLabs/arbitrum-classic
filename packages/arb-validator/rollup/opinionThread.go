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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type PreparedAssertion struct {
	prev        *structures.Node
	beforeState *valprotocol.VMProtoData
	params      *valprotocol.AssertionParams
	claim       *valprotocol.AssertionClaim
	assertion   *protocol.ExecutionAssertion
	machine     machine.Machine
}

func (pa *PreparedAssertion) Clone() *PreparedAssertion {
	return &PreparedAssertion{
		prev:        pa.prev,
		beforeState: pa.beforeState.Clone(),
		params:      pa.params.Clone(),
		claim:       pa.claim.Clone(),
		assertion:   pa.assertion,
		machine:     pa.machine,
	}
}

func (pa *PreparedAssertion) PossibleFutureNode(chainParams valprotocol.ChainParams) *structures.Node {
	node := structures.NewValidNodeFromPrev(
		pa.prev,
		valprotocol.NewDisputableNode(
			pa.params,
			pa.claim,
			common.Hash{},
			big.NewInt(0),
		),
		chainParams,
		common.BlocksFromSeconds(time.Now().Unix()),
		common.Hash{},
	)
	_ = node.UpdateValidOpinion(pa.machine, pa.assertion)
	return node
}

func (chain *ChainObserver) startOpinionUpdateThread(ctx context.Context) {
	go func() {
		assertionPreparedChan := make(chan *PreparedAssertion, 20)
		preparingAssertions := make(map[common.Hash]bool)
		preparedAssertions := make(map[common.Hash]*PreparedAssertion)

		updateCurrent := func() {
			currentOpinion := chain.calculatedValidNode
			currentHash := currentOpinion.Hash()
			log.Println("Building opinion on top of", currentHash)
			successorHashes := currentOpinion.SuccessorHashes()
			successor := func() *structures.Node {
				for _, successor := range successorHashes {
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

			disputable := successor.Disputable()

			if disputable == nil {
				panic("Node was created with disputable assertion")
			}

			if found &&
				prepped.params.Equals(disputable.AssertionParams) &&
				prepped.claim.Equals(disputable.AssertionClaim) {
				newOpinion = valprotocol.ValidChildType
				nextMachine = prepped.machine
				validExecution = prepped.assertion
				chain.RUnlock()
			} else {
				params := disputable.AssertionParams.Clone()
				claim := disputable.AssertionClaim.Clone()
				prevInboxCount := new(big.Int).Set(currentOpinion.VMProtoData().InboxCount)
				afterInboxTopHeight := new(big.Int).Add(prevInboxCount, params.ImportedMessageCount)
				afterInboxTopVal, err := chain.inbox.GetHashAtIndex(afterInboxTopHeight)
				var afterInboxTop *common.Hash
				if err == nil {
					afterInboxTop = &afterInboxTopVal
				}
				inbox, _ := chain.inbox.GenerateVMInbox(currentOpinion.VMProtoData().InboxTop, params.ImportedMessageCount.Uint64())
				messagesVal := inbox.AsValue()
				nextMachine = currentOpinion.Machine().Clone()

				chain.RUnlock()

				newOpinion, validExecution = getNodeOpinion(params, claim, afterInboxTop, inbox.Hash().Hash(), messagesVal, nextMachine)
			}
			// Reset prepared
			preparingAssertions = make(map[common.Hash]bool)
			preparedAssertions = make(map[common.Hash]*PreparedAssertion)

			chain.RLock()
			correctNode := chain.nodeGraph.GetSuccessor(currentOpinion, newOpinion)
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
				if correctNode.Depth() > chain.knownValidNode.Depth() {
					chain.knownValidNode = correctNode
				}
				chain.Unlock()
				chain.RLock()
				for _, listener := range chain.listeners {
					listener.AdvancedKnownNode(ctx, chain, correctNode)
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
			case prepped := <-assertionPreparedChan:
				preparedAssertions[prepped.prev.Hash()] = prepped
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
				_, isPreparing := preparingAssertions[chain.calculatedValidNode.Hash()]
				if !isPreparing {
					newMessages := chain.calculatedValidNode.VMProtoData().InboxTop != chain.inbox.GetTopHash()
					if chain.calculatedValidNode.Machine() != nil &&
						chain.calculatedValidNode.Machine().IsBlocked(chain.latestBlockId.Height, newMessages) == nil {
						preparingAssertions[chain.calculatedValidNode.Hash()] = true
						go func() {
							assertionPreparedChan <- chain.prepareAssertion()
						}()
					}
				} else {
					prepared, isPrepared := preparedAssertions[chain.calculatedValidNode.Hash()]
					if isPrepared && chain.nodeGraph.leaves.IsLeaf(chain.calculatedValidNode) {
						lowerBoundBlock := prepared.params.TimeBounds.LowerBoundBlock
						upperBoundBlock := prepared.params.TimeBounds.UpperBoundBlock
						endCushion := common.NewTimeBlocks(new(big.Int).Add(chain.latestBlockId.Height.AsInt(), big.NewInt(3)))
						if chain.latestBlockId.Height.Cmp(lowerBoundBlock) >= 0 && endCushion.Cmp(upperBoundBlock) <= 0 {
							chain.RUnlock()
							chain.Lock()
							chain.pendingState = prepared.machine
							chain.Unlock()
							chain.RLock()
							for _, lis := range chain.listeners {
								lis.AssertionPrepared(ctx, chain, prepared.Clone())
							}
						} else {
							log.Printf("Throwing out out of date assertion with bounds [%v, %v] at time %v\n", lowerBoundBlock.AsInt(), upperBoundBlock.AsInt(), chain.latestBlockId.Height.AsInt())
							// Prepared assertion is out of date
							delete(preparingAssertions, chain.calculatedValidNode.Hash())
							delete(preparedAssertions, chain.calculatedValidNode.Hash())
						}
					}
				}
				chain.RUnlock()

			}
		}
	}()
}

func (chain *ChainObserver) prepareAssertion() *PreparedAssertion {
	chain.RLock()
	currentOpinion := chain.calculatedValidNode
	beforeState := currentOpinion.VMProtoData().Clone()
	if !chain.nodeGraph.leaves.IsLeaf(currentOpinion) {
		return nil
	}
	afterInboxTop := chain.inbox.GetTopHash()
	beforeInboxTop := beforeState.InboxTop
	newMessageCount := new(big.Int).Sub(chain.inbox.TopCount(), beforeState.InboxCount)

	inbox, _ := chain.inbox.GenerateVMInbox(beforeInboxTop, newMessageCount.Uint64())
	messagesVal := inbox.AsValue()
	mach := currentOpinion.Machine().Clone()
	timeBounds := chain.currentTimeBounds()
	log.Println("timeBounds: ", timeBounds.LowerBoundBlock.String(), timeBounds.UpperBoundBlock.String())
	maxSteps := chain.nodeGraph.params.MaxExecutionSteps
	currentHeight := chain.latestBlockId.Height.Clone()
	timeBoundsLength := new(big.Int).Sub(timeBounds.UpperBoundBlock.AsInt(), timeBounds.LowerBoundBlock.AsInt())
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
		timeBounds.LowerBoundBlock.AsInt(),
		timeBounds.UpperBoundBlock.AsInt(),
		currentOpinion.Hash(),
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
			ImportedMessagesSlice: inbox.Hash().Hash(),
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
	return &PreparedAssertion{
		prev:        currentOpinion,
		beforeState: beforeState,
		params:      params,
		claim:       claim,
		assertion:   assertion,
		machine:     mach,
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
