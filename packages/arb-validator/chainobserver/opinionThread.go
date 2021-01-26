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
	"github.com/pkg/errors"
	"math/big"
	"sync"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/chainlistener"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/nodegraph"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func (chain *ChainObserver) startOpinionUpdateThread(ctx context.Context) {
	go func() {
		logger.Info().Msg("Launching opinion thread")
		preparingAssertions := make(map[common.Hash]struct{})
		preparedAssertions := make(map[common.Hash]*chainlistener.PreparedAssertion)
		// This mutex protects all access to preparingAssertions and preparedAssertions
		assertionsMut := new(sync.Mutex)

		updateCurrent := func() {
			currentOpinion := chain.calculatedValidNode
			currentHash := currentOpinion.Hash()
			logger.Info().Hex("current", currentHash.Bytes()).Msg("Building opinion on top")
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
			assertionsMut.Lock()
			prepped, found := preparedAssertions[currentHash]
			assertionsMut.Unlock()
			disputable := successor.Disputable()

			if disputable == nil {
				panic("Node was created with disputable assertion")
			}

			if found &&
				prepped.Params.Equals(disputable.AssertionParams) &&
				prepped.AssertionStub.Equals(disputable.Assertion) {
				newOpinion = valprotocol.ValidChildType
				nextMachine = prepped.Machine
				validExecution = prepped.Assertion
				chain.RUnlock()
			} else {
				params := disputable.AssertionParams.Clone()
				claim := disputable.Assertion.Clone()
				prevInboxCount := new(big.Int).Set(currentOpinion.VMProtoData().InboxCount)
				afterInboxTopHeight := new(big.Int).Add(prevInboxCount, params.ImportedMessageCount)
				afterInboxTopVal, err := chain.Inbox.GetHashAtIndex(afterInboxTopHeight)
				var afterInboxTop *common.Hash
				if err == nil {
					afterInboxTop = &afterInboxTopVal
				}
				nextMachine = currentOpinion.Machine().Clone()
				logger.Info().Hex("successor", successor.Hash().Bytes()).Msg("Forming opinion on")

				chain.RUnlock()

				newOpinion, validExecution = chain.getNodeOpinion(params, claim, afterInboxTop, nextMachine)
			}
			// Reset prepared
			assertionsMut.Lock()
			preparingAssertions = make(map[common.Hash]struct{})
			preparedAssertions = make(map[common.Hash]*chainlistener.PreparedAssertion)
			assertionsMut.Unlock()
			chain.RLock()
			correctNode := chain.NodeGraph.GetSuccessor(currentOpinion, newOpinion)
			if correctNode != nil {
				chain.RUnlock()
				chain.Lock()
				if newOpinion == valprotocol.ValidChildType {
					// Already confirmed node is valid, so error can be ignored
					_ = correctNode.UpdateValidOpinion(nextMachine, validExecution)
				} else {
					// Already confirmed node is invalid, so error can be ignored
					_ = correctNode.UpdateInvalidOpinion()
				}
				logger.Info().
					Str("opinion", newOpinion.String()).
					Hex("successor", successorHashes[newOpinion].Bytes()).
					Hex("current", currentHash.Bytes()).
					Hex("machine", correctNode.Machine().Hash().Bytes()).
					Msg("Formed new opinion")
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
				logger.Info().
					Hex("successor", successorHashes[newOpinion].Bytes()).
					Msg("Formed opinion on nonexistant node")
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
				if !chain.atHead || chain.calculatedValidNode.Machine() == nil {
					chain.RUnlock()
					break
				}
				// Prepare next assertion
				assertionsMut.Lock()
				prevNode := chain.calculatedValidNode.Hash()
				_, isPreparing := preparingAssertions[prevNode]
				preparingAssertions[prevNode] = struct{}{}
				assertionsMut.Unlock()
				if !isPreparing {
					go func() {
						prepped, err := chain.prepareAssertion(chain.assumedValidBlock)
						assertionsMut.Lock()
						if err != nil {
							delete(preparingAssertions, prevNode)
							assertionsMut.Unlock()
							return
						}
						preparedAssertions[prevNode] = prepped
						assertionsMut.Unlock()
					}()
				} else {
					assertionsMut.Lock()
					prepared, isPrepared := preparedAssertions[chain.calculatedValidNode.Hash()]
					assertionsMut.Unlock()
					if isPrepared && chain.NodeGraph.Leaves().IsLeaf(chain.calculatedValidNode) {
						if new(big.Int).Sub(chain.assumedValidBlock.Height.AsInt(), prepared.ValidBlock.Height.AsInt()).Cmp(big.NewInt(200)) < 0 {
							for _, lis := range chain.listeners {
								lis.AssertionPrepared(
									ctx,
									chain.NodeGraph.Params(),
									chain.NodeGraph,
									chain.KnownValidNode,
									prepared.Clone(),
								)
							}
						} else {
							assertionsMut.Lock()
							// Prepared assertion is out of date
							logger.Info().Hex("node", chain.calculatedValidNode.Hash().Bytes()).Msg("Throwing out old assertion")
							delete(preparingAssertions, chain.calculatedValidNode.Hash())
							delete(preparedAssertions, chain.calculatedValidNode.Hash())
							assertionsMut.Unlock()
						}
					}
				}
				chain.RUnlock()

			}
		}
	}()
}

func (chain *ChainObserver) prepareAssertion(maxValidBlock *common.BlockId) (*chainlistener.PreparedAssertion, error) {
	chain.RLock()
	currentOpinion := chain.calculatedValidNode

	if !chain.NodeGraph.Leaves().IsLeaf(currentOpinion) {
		chain.RUnlock()
		return nil, errors.New("current opinion is not a leaf")
	}

	beforeState := currentOpinion.VMProtoData().Clone()

	var newMessages bool
	var maxMessageCount *big.Int
	var found bool
	found, maxMessageCount = chain.Inbox.GetMaxAtHeight(maxValidBlock.Height)
	if !found {
		maxMessageCount = beforeState.InboxCount
	}

	newMessages = maxMessageCount.Cmp(beforeState.InboxCount) > 0

	if currentOpinion.Machine().IsBlocked(newMessages) != nil {
		chain.RUnlock()
		return nil, errors.New("machine blocked")
	}

	beforeInboxTop := beforeState.InboxTop
	newMessageCount := new(big.Int).Sub(maxMessageCount, beforeState.InboxCount)

	messages, err := chain.Inbox.GetMessages(beforeInboxTop, newMessageCount.Uint64())
	if err != nil {
		logger.Warn().Stack().Err(err).Msg("Nonfatal error getting messages")
	}

	mach := currentOpinion.Machine().Clone()
	maxSteps := chain.NodeGraph.Params().MaxExecutionSteps
	chain.RUnlock()

	beforeHash := mach.Hash()

	assertion, _, stepsRun := mach.ExecuteAssertion(maxSteps, true, messages, true)

	afterHash := mach.Hash()

	blockReason := mach.IsBlocked(false)

	logger.Info().
		Uint64("stepsRun", stepsRun).
		Hex("before", beforeHash.Bytes()).
		Hex("after", afterHash.Bytes()).
		Hex("currentOpinion", currentOpinion.Hash().Bytes()).
		Msgf("Prepared assertion with block reason: %v", blockReason)

	chain.RLock()
	defer chain.RUnlock()
	stub := structures.NewExecutionAssertionStubFromWholeAssertion(assertion, beforeInboxTop, chain.Inbox.MessageStack)
	params := &valprotocol.AssertionParams{
		NumSteps:             stepsRun,
		ImportedMessageCount: new(big.Int).SetUint64(assertion.InboxMessagesConsumed),
	}
	return &chainlistener.PreparedAssertion{
		Prev:          currentOpinion,
		BeforeState:   beforeState,
		Params:        params,
		AssertionStub: stub,
		Assertion:     assertion,
		Machine:       mach,
		ValidBlock:    maxValidBlock,
	}, nil
}

func (chain *ChainObserver) getNodeOpinion(
	params *valprotocol.AssertionParams,
	assertionStub *valprotocol.ExecutionAssertionStub,
	afterInboxTop *common.Hash,
	mach machine.Machine,
) (valprotocol.ChildType, *protocol.ExecutionAssertion) {
	if afterInboxTop == nil || assertionStub.AfterInboxHash != *afterInboxTop {
		logger.Info().
			Hex("address", assertionStub.AfterInboxHash.Bytes()).
			Msg("Saw node with invalid after inbox top claim")
		return valprotocol.InvalidInboxTopChildType, nil
	}

	chain.RLock()
	messages, err := chain.Inbox.GetMessages(assertionStub.BeforeInboxHash, params.ImportedMessageCount.Uint64())
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("accepted assertion can't overrun the inbox")
	}
	chain.RUnlock()

	assertion, _, stepsRun := mach.ExecuteAssertion(
		params.NumSteps,
		true,
		messages,
		true,
	)
	chain.RLock()
	defer chain.RUnlock()
	if params.NumSteps != stepsRun || !assertionStub.Equals(structures.NewExecutionAssertionStubFromWholeAssertion(assertion, assertionStub.BeforeInboxHash, chain.Inbox.MessageStack)) {
		logger.Warn().Msg("Saw node with invalid execution claim")
		return valprotocol.InvalidExecutionChildType, nil
	}

	return valprotocol.ValidChildType, assertion
}
