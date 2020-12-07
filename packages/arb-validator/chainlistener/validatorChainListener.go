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

package chainlistener

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenges"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/nodegraph"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"math/big"
	"sync"
	"time"
)

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

func NewValidatorChainListener(
	ctx context.Context,
	rollupAddress common.Address,
	actor arbbridge.ArbRollup,
) *ValidatorChainListener {
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

func stakeLatestValid(
	ctx context.Context,
	nodeGraph *nodegraph.StakedNodeGraph,
	location *structures.Node,
	stakingKey *StakingKey,
) error {
	proof1 := structures.GeneratePathProof(nodeGraph.LatestConfirmed(), location)
	proof2 := structures.GeneratePathProof(location, nodeGraph.GetLeaf(location))
	stakeAmount := nodeGraph.Params().StakeRequirement

	logger.Info().
		Hex("address", stakingKey.client.Address().Bytes()).
		Msg("Placing stake")
	_, err := stakingKey.contract.PlaceStake(ctx, stakeAmount, proof1, proof2)
	return err
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

func MakeAssertion(
	ctx context.Context,
	rollup arbbridge.ArbRollup,
	prepared *PreparedAssertion,
	proof []common.Hash,
) ([]arbbridge.Event, error) {
	return rollup.MakeAssertion(
		ctx,
		prepared.Prev.PrevHash(),
		prepared.Prev.NodeDataHash(),
		prepared.Prev.Deadline(),
		prepared.Prev.LinkType(),
		prepared.BeforeState,
		prepared.Params,
		prepared.AssertionStub,
		proof,
		prepared.ValidBlock,
	)
}

func (lis *ValidatorChainListener) AddedToChain(context.Context, []*structures.Node) {
}

func (lis *ValidatorChainListener) RestartingFromLatestValid(context.Context, *structures.Node) {
}

func (lis *ValidatorChainListener) AssertionPrepared(
	ctx context.Context,
	_ valprotocol.ChainParams,
	nodeGraph *nodegraph.StakedNodeGraph,
	nodeLocation *structures.Node,
	prepared *PreparedAssertion,
) {
	// Anyone confirm a node
	// No need to have your own stake
	lis.Lock()
	prevParams, alreadySent := lis.broadcastAssertions[prepared.Prev.Hash()]
	lis.Unlock()
	if alreadySent && prevParams.Equals(prepared.Params) {
		return
	}

	for stakingAddress, stakingKey := range lis.stakingKeys {
		stakerPos := nodeGraph.Stakers().Get(stakingAddress)
		if stakerPos == nil {
			// stakingKey is not staked
			continue
		}
		proof := structures.GeneratePathProof(stakerPos.Location(), prepared.Prev)
		if proof == nil {
			// staker can't move to new asertion
			continue
		}
		lis.Lock()
		lis.broadcastAssertions[prepared.Prev.Hash()] = prepared.Params
		lis.Unlock()
		logger.Info().
			Hex("address", stakingAddress.Bytes()).
			Msg("Making assertion")
		go func() {
			_, err := MakeAssertion(ctx, stakingKey.contract, prepared.Clone(), proof)
			if err != nil {
				logger.Warn().
					Stack().
					Err(err).
					Hex("address", stakingAddress.Bytes()).
					Msg("Error making assertion")
				lis.Lock()
				delete(lis.broadcastAssertions, prepared.Prev.Hash())
				lis.Unlock()
			} else {
				logger.Info().
					Hex("address", stakingAddress.Bytes()).
					Msg("Successfully made assertion")
			}
		}()
		return
	}

	logger.Info().Msg("Maybe putting down stake")
	for stakingAddress, stakingKey := range lis.stakingKeys {
		stakerPos := nodeGraph.Stakers().Get(stakingAddress)
		if stakerPos != nil {
			// stakingKey is already down
			continue
		}
		lis.Lock()
		currentTime, err := stakingKey.client.BlockIdForHeight(ctx, nil)
		if err != nil {
			logger.Error().
				Stack().
				Err(err).
				Msg("Validator couldn't get time")
			break
		}
		stakeTime, placedStake := lis.broadcastCreateStakes[stakingAddress]
		if placedStake {
			logger.Info().
				Str("currentHeight", currentTime.Height.AsInt().String()).
				Str("stakeHeight", new(big.Int).Add(stakeTime.AsInt(), big.NewInt(3)).String()).
				Msg("Thinking about placing stake")
		}
		if !placedStake || currentTime.Height.AsInt().Cmp(new(big.Int).Add(stakeTime.AsInt(), big.NewInt(3))) >= 0 {
			lis.broadcastCreateStakes[stakingAddress] = currentTime.Height
			logger.Info().Msg("No stake is currently down, so setting up a stake")
			lis.Unlock()
			// Put down new stake so that we can assert next time
			go func() {
				err := stakeLatestValid(ctx, nodeGraph, nodeLocation, stakingKey)
				if err != nil {
					lis.Lock()
					delete(lis.broadcastCreateStakes, stakingAddress)
					lis.Unlock()
					logger.Warn().Stack().Err(err).Msg("Error placing stake")
				}
			}()
			return
		} else {
			lis.Unlock()
		}
	}
}

func (lis *ValidatorChainListener) StakeCreated(
	ctx context.Context,
	nodeGraph *nodegraph.StakedNodeGraph,
	ev arbbridge.StakeCreatedEvent,
) {
	_, ok := lis.stakingKeys[ev.Staker]
	if ok {
		staker := nodeGraph.Stakers().Get(ev.Staker)
		if staker == nil {
			panic("Stake created but address is not in graph")
		}
		opp := nodeGraph.CheckChallengeOpportunityAny(staker)
		if opp != nil {
			_, err := InitiateChallenge(ctx, lis.actor, opp)
			if err != nil {
				logger.Warn().
					Stack().
					Err(err).
					Hex("staker", ev.Staker.Bytes()).
					Msg("Unable to initiate challenge")
			}
		}
	} else {
		opp := lis.challengeStakerIfPossible(nodeGraph, ev.Staker)
		if opp != nil {
			_, err := InitiateChallenge(ctx, lis.actor, opp)
			if err != nil {
				logger.Warn().
					Stack().
					Err(err).
					Hex("staker", ev.Staker.Bytes()).
					Msg("Unable to initiate challenge")
			}
		}
	}
}

func (lis *ValidatorChainListener) StakeMoved(
	ctx context.Context,
	nodeGraph *nodegraph.StakedNodeGraph,
	ev arbbridge.StakeMovedEvent,
) {
	opp := lis.challengeStakerIfPossible(nodeGraph, ev.Staker)

	if opp != nil {
		_, err := InitiateChallenge(ctx, lis.actor, opp)
		if err != nil {
			logger.Warn().
				Stack().
				Err(err).
				Hex("staker", ev.Staker.Bytes()).
				Msg("Unable to initiate challenge")
		}
	}
}

func (lis *ValidatorChainListener) challengeStakerIfPossible(nodeGraph *nodegraph.StakedNodeGraph, stakerAddr common.Address) *nodegraph.ChallengeOpportunity {
	_, ok := lis.stakingKeys[stakerAddr]
	if ok {
		// Don't challenge yourself
		return nil
	}

	newStaker := nodeGraph.Stakers().Get(stakerAddr)
	if newStaker == nil {
		logger.Fatal().
			Stack().
			Hex("staker", stakerAddr.Bytes()).
			Msg("Nonexistant staker moved")
	}

	// Search for an already staked staking key
	for myAddr := range lis.stakingKeys {
		meAsStaker := nodeGraph.Stakers().Get(myAddr)
		if meAsStaker == nil {
			continue
		}
		opp := nodeGraph.CheckChallengeOpportunityPair(newStaker, meAsStaker)
		if opp != nil {
			return opp
		}
		return nodeGraph.CheckChallengeOpportunityAny(newStaker)
	}
	return nil
}

// All functions below are either only called if you have a stake down, or don't require a stake

func (lis *ValidatorChainListener) StartedChallenge(
	ctx context.Context,
	msgStack *structures.MessageStack,
	chal *nodegraph.Challenge) {
	lis.launchChallenge(ctx, msgStack, chal)
}

func (lis *ValidatorChainListener) ResumedChallenge(
	ctx context.Context,
	msgStack *structures.MessageStack,
	chal *nodegraph.Challenge) {
	lis.launchChallenge(ctx, msgStack, chal)
}

func (lis *ValidatorChainListener) launchChallenge(
	ctx context.Context,
	msgStack *structures.MessageStack,
	chal *nodegraph.Challenge) {
	// Must already be staked to be challenged
	startBlockId := chal.BlockId()
	startLogIndex := chal.LogIndex() - 1
	asserterKey, ok := lis.stakingKeys[chal.Asserter()]
	if ok {
		switch chal.ConflictNode().LinkType() {
		case valprotocol.InvalidInboxTopChildType:
			go func() {
				res, err := challenges.DefendInboxTopClaim(
					ctx,
					asserterKey.client,
					chal.Contract(),
					startBlockId,
					startLogIndex,
					msgStack,
					chal.ConflictNode().Disputable().Assertion.AfterInboxHash,
					new(big.Int).Sub(
						chal.ConflictNode().Disputable().MaxInboxCount,
						new(big.Int).Add(chal.ConflictNode().Prev().VMProtoData().InboxCount, chal.ConflictNode().Disputable().AssertionParams.ImportedMessageCount),
					),
					100,
				)
				if err != nil {
					logger.Error().
						Stack().
						Err(err).
						Msg("Failed defending inbox top claim")
				} else {
					logger.Info().
						Str("challengeState", res.String()).
						Msg("Completed defending inbox top claim")
				}
			}()
		case valprotocol.InvalidExecutionChildType:
			go func() {
				res, err := challenges.DefendExecutionClaim(
					ctx,
					asserterKey.client,
					chal.Contract(),
					startBlockId,
					startLogIndex,
					chal.ConflictNode().Prev().Machine(),
					chal.ConflictNode().Disputable().Assertion,
					msgStack,
					chal.ConflictNode().Disputable().AssertionParams.NumSteps,
					50,
					challenges.StandardExecutionChallenge(),
				)
				if err != nil {
					logger.Error().
						Stack().
						Err(err).
						Msg("Failed defending execution claim")
				} else {
					logger.Info().
						Str("challengeState", res.String()).
						Msg("Completed defending execution claim")
				}
			}()
		default:
			logger.Fatal().
				Uint("challengeType", uint(chal.ConflictNode().LinkType())).
				Msg("Unexpected challenge type")
		}
	}

	challenger, ok := lis.stakingKeys[chal.Challenger()]
	if ok {
		switch chal.ConflictNode().LinkType() {
		case valprotocol.InvalidInboxTopChildType:
			go func() {
				res, err := challenges.ChallengeInboxTopClaim(
					ctx,
					challenger.client,
					chal.Contract(),
					startBlockId,
					startLogIndex,
					msgStack,
					false,
				)
				if err != nil {
					logger.Error().
						Stack().
						Err(err).
						Msg("Failed challenging inbox top claim")
				} else {
					logger.Info().
						Str("challengeState", res.String()).
						Msg("Completed challenging inbox top claim")
				}
			}()
		case valprotocol.InvalidExecutionChildType:
			go func() {
				res, err := challenges.ChallengeExecutionClaim(
					ctx,
					challenger.client,
					chal.Contract(),
					startBlockId,
					startLogIndex,
					msgStack,
					chal.ConflictNode().Disputable().AssertionParams.NumSteps,
					chal.ConflictNode().Prev().Machine(),
					chal.ConflictNode().VMProtoData().InboxTop,
					false,
					challenges.StandardExecutionChallenge(),
				)
				if err != nil {
					logger.Error().
						Stack().
						Err(err).
						Msg("Failed challenging execution claim")
				} else {
					logger.Info().
						Str("challengeState", res.String()).
						Msg("Completed challenging execution claim")
				}
			}()
		default:
			logger.Fatal().Stack().Msg("Unexpected challenge type")
		}
	}
}

func (lis *ValidatorChainListener) CompletedChallenge(
	ctx context.Context,
	nodeGraph *nodegraph.StakedNodeGraph,
	ev arbbridge.ChallengeCompletedEvent,
) {
	// Must be staked to have challenge completed
	_, ok := lis.stakingKeys[ev.Winner]
	if ok {
		lis.wonChallenge(ev)
	}
	_, ok = lis.stakingKeys[ev.Loser]
	if ok {
		lis.lostChallenge(ev)
	}
	opp := lis.challengeStakerIfPossible(nodeGraph, ev.Winner)
	if opp != nil {
		_, err := InitiateChallenge(ctx, lis.actor, opp)
		LogChallengeResult(err)
	}
}

func (lis *ValidatorChainListener) ConfirmableNodes(ctx context.Context, conf *valprotocol.ConfirmOpportunity) {
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
		_, err := lis.actor.Confirm(ctx, confClone)
		if err != nil {
			logger.Warn().Stack().Err(err).Msg("Failed to confirm valid node")
			lis.Lock()
			delete(lis.broadcastConfirmations, confClone.CurrentLatestConfirmed)
			lis.Unlock()
		}
	}()
}

func (lis *ValidatorChainListener) PrunableLeafs(ctx context.Context, params []valprotocol.PruneParams) {
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
		_, err := lis.actor.PruneLeaves(ctx, leavesToPrune)
		if err != nil {
			logger.Warn().Stack().Err(err).Msg("Failed pruning leaves")
			lis.Lock()
			for _, prune := range leavesToPrune {
				delete(lis.broadcastLeafPrunes, prune.LeafHash)
			}
			lis.Unlock()
		}
	}()
}

func (lis *ValidatorChainListener) MootableStakes(ctx context.Context, params []nodegraph.RecoverStakeMootedParams) {
	// Anyone can moot any stake
	for _, moot := range params {
		mootCopy := moot
		go func() {
			_, err := lis.actor.RecoverStakeMooted(
				ctx,
				mootCopy.AncestorHash,
				mootCopy.Addr,
				mootCopy.LcProof,
				mootCopy.StProof,
			)
			if err != nil {
				logger.Warn().
					Stack().
					Err(err).
					Hex("address", mootCopy.Addr.Bytes()).
					Msg("Unable to recover mooted stake")
			}
		}()
	}
}

func (lis *ValidatorChainListener) OldStakes(ctx context.Context, params []nodegraph.RecoverStakeOldParams) {
	// Anyone can remove an old stake
	for _, old := range params {
		oldCopy := old
		go func() {
			_, err := lis.actor.RecoverStakeOld(
				ctx,
				oldCopy.Addr,
				oldCopy.Proof,
			)
			if err != nil {
				logger.Warn().
					Stack().
					Err(err).
					Hex("address", oldCopy.Addr.Bytes()).
					Msg("Unable to recover old stake")
			}
		}()
	}
}

func (lis *ValidatorChainListener) AdvancedKnownNode(
	ctx context.Context,
	nodeGraph *nodegraph.StakedNodeGraph,
	node *structures.Node) {
	// TODO: It would be better to rate limit how often the stake can be moved
	// and just move to the latest position at the end of a delay period
	for stakingAddress := range lis.stakingKeys {
		staker := nodeGraph.Stakers().Get(stakingAddress)
		if staker == nil {
			continue
		}
		if node.Depth() <= staker.Location().Depth() {
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
				prevMoveNode := nodeGraph.NodeFromHash(prevMove.nodeHash)
				if prevMoveNode != nil {
					return prevMoveNode
				}
			}
			return staker.Location()
		}()

		move := attemptedMove{
			nodeHeight: node.Depth(),
			nodeHash:   node.Hash(),
		}

		lis.broadcastMovedStakes[stakingAddress] = move
		lis.Unlock()

		proof1 := structures.GeneratePathProof(stakerLocation, node)
		proof2 := structures.GeneratePathProof(node, nodeGraph.GetLeaf(node))
		stakingAddr := stakingAddress
		go func() {
			_, err := lis.actor.MoveStake(ctx, proof1, proof2)
			lis.Lock()
			if err != nil {
				logger.Warn().Stack().Err(err).Msg("Failed moving stake")
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

func (lis *ValidatorChainListener) StakeRemoved(context.Context, arbbridge.StakeRefundedEvent) {
}
func (lis *ValidatorChainListener) lostChallenge(arbbridge.ChallengeCompletedEvent) {}
func (lis *ValidatorChainListener) wonChallenge(arbbridge.ChallengeCompletedEvent)  {}
func (lis *ValidatorChainListener) SawAssertion(context.Context, arbbridge.AssertedEvent) {
}
func (lis *ValidatorChainListener) ConfirmedNode(context.Context, arbbridge.ConfirmedEvent) {
}
func (lis *ValidatorChainListener) PrunedLeaf(context.Context, arbbridge.PrunedEvent) {
}
func (lis *ValidatorChainListener) MessageDelivered(context.Context, arbbridge.MessageDeliveredEvent) {
}
