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
	"sync"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenges"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

const (
	PruneSizeLimit = 120
)

type ChainListener interface {
	// This function is called when a ChainListener is added to a ChainObserver
	AddedToChain(context.Context, *ChainObserver)

	// This function is called every time ChainObserver starts running. This
	// includes both the initial run, and after a reorg. The third parameter
	// is the current calculated valid node
	RestartingFromLatestValid(context.Context, *ChainObserver, *structures.Node)

	StakeCreated(context.Context, *ChainObserver, arbbridge.StakeCreatedEvent)
	StakeRemoved(context.Context, *ChainObserver, arbbridge.StakeRefundedEvent)
	StakeMoved(context.Context, *ChainObserver, arbbridge.StakeMovedEvent)
	StartedChallenge(context.Context, *ChainObserver, *Challenge)
	ResumedChallenge(context.Context, *ChainObserver, *Challenge)
	CompletedChallenge(context.Context, *ChainObserver, arbbridge.ChallengeCompletedEvent)
	SawAssertion(context.Context, *ChainObserver, arbbridge.AssertedEvent)
	ConfirmedNode(context.Context, *ChainObserver, arbbridge.ConfirmedEvent)
	PrunedLeaf(context.Context, *ChainObserver, arbbridge.PrunedEvent)
	MessageDelivered(context.Context, *ChainObserver, arbbridge.MessageDeliveredEvent)

	AssertionPrepared(context.Context, *ChainObserver, *PreparedAssertion)
	ConfirmableNodes(context.Context, *ChainObserver, *valprotocol.ConfirmOpportunity)
	PrunableLeafs(context.Context, *ChainObserver, []valprotocol.PruneParams)
	MootableStakes(context.Context, *ChainObserver, []RecoverStakeMootedParams)
	OldStakes(context.Context, *ChainObserver, []RecoverStakeOldParams)

	AdvancedKnownNode(context.Context, *ChainObserver, *structures.Node)
}

type NoopListener struct{}

func (nl *NoopListener) AddedToChain(context.Context, *ChainObserver, *structures.Node) {
}

func (nl *NoopListener) RestartingFromLatestValid(context.Context, *ChainObserver, *structures.Node) {
}

func (NoopListener) StakeCreated(context.Context, *ChainObserver, arbbridge.StakeCreatedEvent) {
}
func (NoopListener) StakeRemoved(context.Context, *ChainObserver, arbbridge.StakeRefundedEvent) {
}
func (NoopListener) StakeMoved(context.Context, *ChainObserver, arbbridge.StakeMovedEvent) {
}
func (NoopListener) StartedChallenge(context.Context, *ChainObserver, *Challenge) {
}
func (NoopListener) ResumedChallenge(context.Context, *ChainObserver, *Challenge) {

}
func (NoopListener) CompletedChallenge(context.Context, *ChainObserver, arbbridge.ChallengeCompletedEvent) {
}
func (NoopListener) SawAssertion(context.Context, *ChainObserver, arbbridge.AssertedEvent) {
}
func (NoopListener) ConfirmedNode(context.Context, *ChainObserver, arbbridge.ConfirmedEvent) {
}
func (NoopListener) PrunedLeaf(context.Context, *ChainObserver, arbbridge.PrunedEvent) {
}
func (NoopListener) MessageDelivered(context.Context, *ChainObserver, arbbridge.MessageDeliveredEvent) {
}

func (NoopListener) AssertionPrepared(context.Context, *ChainObserver, *PreparedAssertion) {
}
func (NoopListener) ConfirmableNodes(context.Context, *ChainObserver, *valprotocol.ConfirmOpportunity) {
}
func (NoopListener) PrunableLeafs(context.Context, *ChainObserver, []valprotocol.PruneParams) {
}
func (NoopListener) MootableStakes(context.Context, *ChainObserver, []RecoverStakeMootedParams) {
}
func (NoopListener) OldStakes(context.Context, *ChainObserver, []RecoverStakeOldParams) {
}

func (NoopListener) AdvancedKnownNode(context.Context, *ChainObserver, *structures.Node) {}

type attemptedMove struct {
	nodeHeight uint64
	nodeHash   common.Hash
}

type StakingKey struct {
	client   arbbridge.ArbAuthClient
	contract arbbridge.ArbRollup
}

type ValidatorChainListener struct {
	sync.Mutex
	actor                  arbbridge.ArbRollup
	rollupAddress          common.Address
	stakingKeys            map[common.Address]*StakingKey
	broadcastAssertions    map[common.Hash]*valprotocol.AssertionParams
	broadcastConfirmations map[common.Hash]bool
	broadcastLeafPrunes    map[common.Hash]bool
	broadcastCreateStakes  map[common.Address]*common.TimeBlocks
	broadcastMovedStakes   map[common.Address]attemptedMove
}

func NewValidatorChainListener(ctx context.Context, rollupAddress common.Address, actor arbbridge.ArbRollup) *ValidatorChainListener {
	ret := &ValidatorChainListener{
		actor:         actor,
		rollupAddress: rollupAddress,
		stakingKeys:   make(map[common.Address]*StakingKey),
	}
	ret.resetBroadcastCache()
	go func() {
		ticker := time.NewTicker(common.NewTimeBlocksInt(30).Duration())
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				ret.Lock()
				ret.resetBroadcastCache()
				ret.Unlock()
			}
		}
	}()
	return ret
}

func (lis *ValidatorChainListener) resetBroadcastCache() {
	lis.broadcastAssertions = make(map[common.Hash]*valprotocol.AssertionParams)
	lis.broadcastConfirmations = make(map[common.Hash]bool)
	lis.broadcastLeafPrunes = make(map[common.Hash]bool)
	lis.broadcastCreateStakes = make(map[common.Address]*common.TimeBlocks)
	lis.broadcastMovedStakes = make(map[common.Address]attemptedMove)
}

func stakeLatestValid(ctx context.Context, chain *ChainObserver, stakingKey *StakingKey) error {
	location := chain.knownValidNode
	proof1 := structures.GeneratePathProof(chain.nodeGraph.latestConfirmed, location)
	proof2 := structures.GeneratePathProof(location, chain.nodeGraph.getLeaf(location))
	stakeAmount := chain.nodeGraph.params.StakeRequirement

	log.Println("Placing stake for", stakingKey.client.Address())
	return stakingKey.contract.PlaceStake(ctx, stakeAmount, proof1, proof2)
}

func (lis *ValidatorChainListener) AddStaker(client arbbridge.ArbAuthClient) error {
	contract, err := client.NewRollup(lis.rollupAddress)
	if err != nil {
		return err
	}

	address := client.Address()
	lis.stakingKeys[address] = &StakingKey{
		client:   client,
		contract: contract,
	}
	return nil
}

func makeAssertion(ctx context.Context, rollup arbbridge.ArbRollup, prepared *PreparedAssertion, proof []common.Hash) error {
	return rollup.MakeAssertion(
		ctx,
		prepared.prev.PrevHash(),
		prepared.prev.NodeDataHash(),
		prepared.prev.Deadline(),
		prepared.prev.LinkType(),
		prepared.beforeState,
		prepared.params,
		prepared.claim,
		proof,
	)
}

func (lis *ValidatorChainListener) AddedToChain(context.Context, *ChainObserver) {
}

func (lis *ValidatorChainListener) RestartingFromLatestValid(context.Context, *ChainObserver, *structures.Node) {
}

func (lis *ValidatorChainListener) AssertionPrepared(ctx context.Context, chain *ChainObserver, prepared *PreparedAssertion) {
	// Anyone confirm a node
	// No need to have your own stake
	lis.Lock()
	prevParams, alreadySent := lis.broadcastAssertions[prepared.prev.Hash()]
	lis.Unlock()
	if alreadySent && prevParams.Equals(prepared.params) {
		return
	}

	for stakingAddress, stakingKey := range lis.stakingKeys {
		stakerPos := chain.nodeGraph.stakers.Get(stakingAddress)
		if stakerPos == nil {
			// stakingKey is not staked
			continue
		}
		proof := structures.GeneratePathProof(stakerPos.location, prepared.prev)
		if proof == nil {
			// staker can't move to new asertion
			continue
		}
		lis.Lock()
		lis.broadcastAssertions[prepared.prev.Hash()] = prepared.params
		lis.Unlock()
		log.Printf("%v is making an assertion\n", stakingAddress)
		go func() {
			err := makeAssertion(ctx, stakingKey.contract, prepared.Clone(), proof)
			if err != nil {
				log.Println("Error making assertion", err)
				lis.Lock()
				delete(lis.broadcastAssertions, prepared.prev.Hash())
				lis.Unlock()
			} else {
				log.Println("Successfully made assertion")
			}
		}()
		return
	}

	log.Println("Maybe putting down stake")
	for stakingAddress, stakingKey := range lis.stakingKeys {
		stakerPos := chain.nodeGraph.stakers.Get(stakingAddress)
		if stakerPos != nil {
			// stakingKey is already down
			continue
		}
		lis.Lock()
		stakeTime, placedStake := lis.broadcastCreateStakes[stakingAddress]
		if placedStake {
			log.Println("Thinking about placing stake", chain.latestBlockId.Height.AsInt(), new(big.Int).Add(stakeTime.AsInt(), big.NewInt(3)))
		}
		if !placedStake || chain.latestBlockId.Height.AsInt().Cmp(new(big.Int).Add(stakeTime.AsInt(), big.NewInt(3))) >= 0 {
			lis.broadcastCreateStakes[stakingAddress] = chain.latestBlockId.Height
			log.Println("No stake is currently down, so setting up a stake")
			lis.Unlock()
			// Put down new stake so that we can assert next time
			go func() {
				err := stakeLatestValid(ctx, chain, stakingKey)
				if err != nil {
					lis.Lock()
					delete(lis.broadcastCreateStakes, stakingAddress)
					lis.Unlock()
					log.Println("Error placing stake", err)
				}
			}()
			return
		} else {
			lis.Unlock()
		}
	}
}

func (lis *ValidatorChainListener) initiateChallenge(ctx context.Context, opp *challengeOpportunity) error {
	return lis.actor.StartChallenge(
		ctx,
		opp.asserter,
		opp.challenger,
		opp.prevNodeHash,
		opp.deadlineTicks.Val,
		opp.asserterNodeType,
		opp.challengerNodeType,
		opp.asserterVMProtoHash,
		opp.challengerVMProtoHash,
		opp.asserterProof,
		opp.challengerProof,
		opp.asserterNodeHash,
		opp.challengerDataHash,
		opp.challengerPeriodTicks,
	)
}

func (lis *ValidatorChainListener) StakeCreated(ctx context.Context, chain *ChainObserver, ev arbbridge.StakeCreatedEvent) {
	_, ok := lis.stakingKeys[ev.Staker]
	if ok {
		staker := chain.nodeGraph.stakers.Get(ev.Staker)
		if staker == nil {
			panic("Stake created but address is not in graph")
		}
		opp := chain.nodeGraph.checkChallengeOpportunityAny(staker)
		if opp != nil {
			go func() {
				err := lis.initiateChallenge(ctx, opp)
				if err != nil {
					log.Println("Failed to initiate challenge", err)
				} else {
					log.Println("Successfully initiated challenge")
				}
			}()
		}
	} else {
		lis.challengeStakerIfPossible(ctx, chain, ev.Staker)
	}
}

func (lis *ValidatorChainListener) StakeMoved(ctx context.Context, chain *ChainObserver, ev arbbridge.StakeMovedEvent) {
	lis.challengeStakerIfPossible(ctx, chain, ev.Staker)
}

func (lis *ValidatorChainListener) challengeStakerIfPossible(ctx context.Context, chain *ChainObserver, stakerAddr common.Address) {
	_, ok := lis.stakingKeys[stakerAddr]
	if ok {
		// Don't challenge yourself
		return
	}

	newStaker := chain.nodeGraph.stakers.Get(stakerAddr)
	if newStaker == nil {
		log.Fatalf("Nonexistant staker moved %v", stakerAddr)
	}

	// Search for an already staked staking key
	for myAddr := range lis.stakingKeys {
		meAsStaker := chain.nodeGraph.stakers.Get(myAddr)
		if meAsStaker == nil {
			continue
		}
		opp := chain.nodeGraph.checkChallengeOpportunityPair(newStaker, meAsStaker)
		if opp != nil {
			go func() {
				err := lis.initiateChallenge(ctx, opp)
				if err != nil {
					log.Println("Failed to initiate challenge", err)
				} else {
					log.Println("Successfully initiated challenge")
				}
			}()
			return
		}
	}
	opp := chain.nodeGraph.checkChallengeOpportunityAny(newStaker)
	if opp != nil {
		go func() {
			err := lis.initiateChallenge(ctx, opp)
			if err != nil {
				log.Println("Failed to initiate challenge", err)
			} else {
				log.Println("Successfully initiated challenge")
			}
		}()
		return
	}
}

// All functions below are either only called if you have a stake down, or don't require a stake

func (lis *ValidatorChainListener) StartedChallenge(ctx context.Context, chain *ChainObserver, chal *Challenge) {
	lis.launchChallenge(ctx, chain, chal)
}

func (lis *ValidatorChainListener) ResumedChallenge(ctx context.Context, chain *ChainObserver, chal *Challenge) {
	lis.launchChallenge(ctx, chain, chal)
}

func (lis *ValidatorChainListener) launchChallenge(ctx context.Context, chain *ChainObserver, chal *Challenge) {
	// Must already be staked to be challenged
	startBlockId := chal.blockId
	startLogIndex := chal.logIndex - 1
	asserterKey, ok := lis.stakingKeys[chal.asserter]
	if ok {
		switch chal.conflictNode.LinkType() {
		case valprotocol.InvalidInboxTopChildType:
			go func() {
				res, err := challenges.DefendInboxTopClaim(
					ctx,
					asserterKey.client,
					chal.contract,
					startBlockId,
					startLogIndex,
					chain.inbox.MessageStack,
					chal.conflictNode.Disputable().AssertionClaim.AfterInboxTop,
					new(big.Int).Sub(
						chal.conflictNode.Disputable().MaxInboxCount,
						new(big.Int).Add(chal.conflictNode.Prev().VMProtoData().InboxCount, chal.conflictNode.Disputable().AssertionParams.ImportedMessageCount),
					),
					100,
				)
				if err != nil {
					log.Println("Failed defending inbox top claim", err)
				} else {
					log.Println("Completed defending inbox top claim", res)
				}
			}()
		case valprotocol.InvalidMessagesChildType:
			go func() {
				res, err := challenges.DefendMessagesClaim(
					ctx,
					asserterKey.client,
					chal.contract,
					startBlockId,
					startLogIndex,
					chain.inbox.MessageStack,
					chal.conflictNode.VMProtoData().InboxTop,
					chal.conflictNode.Disputable().AssertionParams.ImportedMessageCount,
					100,
				)
				if err != nil {
					log.Println("Failed defending messages claim", err)
				} else {
					log.Println("Completed defending messages claim", res)
				}
			}()
		case valprotocol.InvalidExecutionChildType:
			go func() {
				res, err := challenges.DefendExecutionClaim(
					ctx,
					asserterKey.client,
					chal.contract,
					startBlockId,
					startLogIndex,
					chain.executionPrecondition(chal.conflictNode),
					chal.conflictNode.Prev().Machine(),
					chal.conflictNode.Disputable().AssertionParams.NumSteps,
					50,
				)
				if err != nil {
					log.Println("Failed defending execution claim", err)
				} else {
					log.Println("Completed defending execution claim", res)
				}
			}()
		default:
			log.Fatal("unexpected challenge type")
		}
	}

	challenger, ok := lis.stakingKeys[chal.challenger]
	if ok {
		switch chal.conflictNode.LinkType() {
		case valprotocol.InvalidInboxTopChildType:
			go func() {
				res, err := challenges.ChallengeInboxTopClaim(
					ctx,
					challenger.client,
					chal.contract,
					startBlockId,
					startLogIndex,
					chain.inbox.MessageStack,
					false,
				)
				if err != nil {
					log.Println("Failed challenging inbox top claim", err)
				} else {
					log.Println("Completed challenging inbox top claim", res)
				}
			}()
		case valprotocol.InvalidMessagesChildType:
			go func() {
				res, err := challenges.ChallengeMessagesClaim(
					ctx,
					challenger.client,
					chal.contract,
					startBlockId,
					startLogIndex,
					chain.inbox.MessageStack,
					chal.conflictNode.VMProtoData().InboxTop,
					chal.conflictNode.Disputable().AssertionParams.ImportedMessageCount,
					false,
				)
				if err != nil {
					log.Println("Failed challenging messages claim", err)
				} else {
					log.Println("Completed challenging messages claim", res)
				}
			}()
		case valprotocol.InvalidExecutionChildType:
			go func() {
				res, err := challenges.ChallengeExecutionClaim(
					ctx,
					challenger.client,
					chal.contract,
					startBlockId,
					startLogIndex,
					chain.executionPrecondition(chal.conflictNode),
					chal.conflictNode.Prev().Machine(),
					false,
				)
				if err != nil {
					log.Println("Failed challenging execution claim", err)
				} else {
					log.Println("Completed challenging execution claim", res)
				}
			}()
		default:
			log.Fatal("unexpected challenge type")
		}
	}
}

func (lis *ValidatorChainListener) CompletedChallenge(ctx context.Context, chain *ChainObserver, ev arbbridge.ChallengeCompletedEvent) {
	// Must be staked to have challenge completed
	_, ok := lis.stakingKeys[ev.Winner]
	if ok {
		lis.wonChallenge(ev)
	}

	_, ok = lis.stakingKeys[ev.Loser]
	if ok {
		lis.lostChallenge(ev)
	}
	lis.challengeStakerIfPossible(ctx, chain, ev.Winner)
}

func (lis *ValidatorChainListener) ConfirmableNodes(ctx context.Context, _ *ChainObserver, conf *valprotocol.ConfirmOpportunity) {
	// Anyone confirm a node
	// No need to have your own stake
	lis.Lock()
	_, alreadySent := lis.broadcastConfirmations[conf.CurrentLatestConfirmed]
	if alreadySent {
		lis.Unlock()
		return
	}
	lis.broadcastConfirmations[conf.CurrentLatestConfirmed] = true
	lis.Unlock()
	confClone := conf.Clone()

	go func() {
		err := lis.actor.Confirm(ctx, confClone)
		if err != nil {
			log.Println("Failed to confirm valid node", err)
			lis.Lock()
			delete(lis.broadcastConfirmations, confClone.CurrentLatestConfirmed)
			lis.Unlock()
		}
	}()
}

func (lis *ValidatorChainListener) PrunableLeafs(ctx context.Context, _ *ChainObserver, params []valprotocol.PruneParams) {
	// Anyone can prune a leaf
	leavesToPrune := make([]valprotocol.PruneParams, 0, len(params))
	lis.Lock()
	totalSize := 0
	for _, prune := range params {
		_, alreadySent := lis.broadcastLeafPrunes[prune.LeafHash]
		if alreadySent {
			continue
		}
		leavesToPrune = append(leavesToPrune, prune)
		lis.broadcastLeafPrunes[prune.LeafHash] = true
		totalSize += len(prune.LeafProof) + len(prune.AncProof) + 1
		if totalSize > PruneSizeLimit {
			break
		}
	}
	lis.Unlock()
	go func() {
		err := lis.actor.PruneLeaves(ctx, leavesToPrune)
		if err != nil {
			log.Println("Failed pruning leaves", err)
			lis.Lock()
			for _, prune := range leavesToPrune {
				delete(lis.broadcastLeafPrunes, prune.LeafHash)
			}
			lis.Unlock()
		}
	}()
}

func (lis *ValidatorChainListener) MootableStakes(ctx context.Context, _ *ChainObserver, params []RecoverStakeMootedParams) {
	// Anyone can moot any stake
	for _, moot := range params {
		mootCopy := moot
		go func() {
			lis.actor.RecoverStakeMooted(
				ctx,
				mootCopy.ancestorHash,
				mootCopy.addr,
				mootCopy.lcProof,
				mootCopy.stProof,
			)
		}()
	}
}

func (lis *ValidatorChainListener) OldStakes(ctx context.Context, _ *ChainObserver, params []RecoverStakeOldParams) {
	// Anyone can remove an old stake
	for _, old := range params {
		oldCopy := old
		go func() {
			lis.actor.RecoverStakeOld(
				ctx,
				oldCopy.addr,
				oldCopy.proof,
			)
		}()
	}
}

func (lis *ValidatorChainListener) AdvancedKnownNode(ctx context.Context, chain *ChainObserver, node *structures.Node) {
	// TODO: It would be better to rate limit how often the stake can be moved
	// and just move to the latest position at the end of a delay period
	for stakingAddress := range lis.stakingKeys {
		staker := chain.nodeGraph.stakers.idx[stakingAddress]
		if staker == nil {
			continue
		}
		if node.Depth() <= staker.location.Depth() {
			continue
		}

		lis.Lock()
		prevMove, alreadySent := lis.broadcastMovedStakes[stakingAddress]

		if alreadySent && node.Depth() <= prevMove.nodeHeight {
			lis.Unlock()
			continue
		}

		// If there's already an outstanding transaction moving the stake to an existing
		// node, make the new move transaction initiate from the position after
		// that move. Otherwise start the move from the staker's current location
		stakerLocation := func() *structures.Node {
			if alreadySent {
				prevMoveNode, found := chain.nodeGraph.nodeFromHash[prevMove.nodeHash]
				if found {
					return prevMoveNode
				}
			}
			return staker.location
		}()

		move := attemptedMove{
			nodeHeight: node.Depth(),
			nodeHash:   node.Hash(),
		}

		lis.broadcastMovedStakes[stakingAddress] = move
		lis.Unlock()

		proof1 := structures.GeneratePathProof(stakerLocation, node)
		proof2 := structures.GeneratePathProof(node, chain.nodeGraph.getLeaf(node))
		stakingAddr := stakingAddress
		go func() {
			err := lis.actor.MoveStake(ctx, proof1, proof2)
			lis.Lock()
			if err != nil {
				log.Println("Failed moving stake", err)
				delete(lis.broadcastMovedStakes, stakingAddr)
			} else {
				prevMove, alreadySent := lis.broadcastMovedStakes[stakingAddr]
				if alreadySent {
					if prevMove.nodeHeight <= move.nodeHeight {
						delete(lis.broadcastMovedStakes, stakingAddr)
					}
				}
			}
			lis.Unlock()
		}()
	}
}

func (lis *ValidatorChainListener) StakeRemoved(context.Context, *ChainObserver, arbbridge.StakeRefundedEvent) {
}
func (lis *ValidatorChainListener) lostChallenge(arbbridge.ChallengeCompletedEvent) {}
func (lis *ValidatorChainListener) wonChallenge(arbbridge.ChallengeCompletedEvent)  {}
func (lis *ValidatorChainListener) SawAssertion(context.Context, *ChainObserver, arbbridge.AssertedEvent) {
}
func (lis *ValidatorChainListener) ConfirmedNode(context.Context, *ChainObserver, arbbridge.ConfirmedEvent) {
}
func (lis *ValidatorChainListener) PrunedLeaf(context.Context, *ChainObserver, arbbridge.PrunedEvent) {
}
func (lis *ValidatorChainListener) MessageDelivered(context.Context, *ChainObserver, arbbridge.MessageDeliveredEvent) {
}
