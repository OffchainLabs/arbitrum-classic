/*
 * Copyright 2021, Offchain Labs, Inc.
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

package staker

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
)

var logger = log.With().Caller().Stack().Str("component", "staker").Logger()

type Strategy uint8

const (
	WatchtowerStrategy Strategy = iota
	DefensiveStrategy
	StakeLatestStrategy
	MakeNodesStrategy
)

type Staker struct {
	*Validator
	activeChallenge *challenge.Challenger
	strategy        Strategy
}

func NewStaker(
	ctx context.Context,
	lookup core.ArbCoreLookup,
	client ethutils.EthClient,
	wallet *ethbridge.ValidatorWallet,
	validatorUtilsAddress common.Address,
	strategy Strategy,
) (*Staker, *ethbridge.DelayedBridgeWatcher, error) {
	val, err := NewValidator(ctx, lookup, client, wallet, validatorUtilsAddress)
	if err != nil {
		return nil, nil, err
	}
	return &Staker{
		Validator: val,
		strategy:  strategy,
	}, val.delayedBridge, nil
}

func (s *Staker) RunInBackground(ctx context.Context) chan bool {
	done := make(chan bool)
	go func() {
		defer func() {
			done <- true
		}()
		backoff := time.Second
		for {
			tx, err := s.Act(ctx)
			if err == nil && tx != nil {
				// Note: methodName isn't accurate, it's just used for logging
				_, err = ethbridge.WaitForReceiptWithResults(ctx, s.client, s.wallet.From().ToEthAddress(), tx, "for staking")
				err = errors.Wrap(err, "error waiting for tx receipt")
				if err == nil {
					logger.Info().Str("hash", tx.Hash().String()).Msg("Successfully executed transaction")
				}
			}
			if err != nil {
				logger.Warn().Err(err).Send()
				<-time.After(backoff)
				if backoff < 60*time.Second {
					backoff *= 2
				}
				continue
			} else {
				backoff = time.Second
			}
			<-time.After(time.Minute)
		}
	}()
	return done
}

func (s *Staker) Act(ctx context.Context) (*types.Transaction, error) {
	s.builder.ClearTransactions()
	rawInfo, err := s.rollup.StakerInfo(ctx, s.wallet.Address())
	if err != nil {
		return nil, err
	}
	latestStakedNode, latestStakedNodeHash, err := s.validatorUtils.LatestStaked(ctx, s.wallet.Address())
	if err != nil {
		return nil, err
	}
	if rawInfo != nil {
		rawInfo.LatestStakedNode = latestStakedNode
	}
	info := OurStakerInfo{
		CanProgress:          true,
		LatestStakedNode:     latestStakedNode,
		LatestStakedNodeHash: latestStakedNodeHash,
		StakerInfo:           rawInfo,
	}

	effectiveStrategy := s.strategy
	nodesLinear, err := s.validatorUtils.AreUnresolvedNodesLinear(ctx)
	if err != nil {
		return nil, err
	}
	if !nodesLinear {
		logger.Warn().Msg("Fork detected")
		if effectiveStrategy == DefensiveStrategy {
			effectiveStrategy = StakeLatestStrategy
		}
	}

	// Resolve nodes if either we're on the make nodes strategy,
	// or we're on the stake latest strategy but don't have a stake
	// (attempt to reduce the current required stake).
	shouldResolveNodes := effectiveStrategy >= MakeNodesStrategy
	if !shouldResolveNodes && effectiveStrategy >= StakeLatestStrategy && rawInfo == nil {
		shouldResolveNodes, err = s.isRequiredStakeElevated(ctx)
		if err != nil {
			return nil, err
		}
	}
	if shouldResolveNodes {
		tx, err := s.removeOldStakers(ctx)
		if err != nil || tx != nil {
			return tx, err
		}
		tx, err = s.resolveTimedOutChallenges(ctx)
		if err != nil || tx != nil {
			return tx, err
		}
		if err := s.resolveNextNode(ctx, rawInfo); err != nil {
			return nil, err
		}
	}

	// Don't attempt to create a new stake if we're resolving a node,
	// as that might affect the current required stake.
	creatingNewStake := rawInfo == nil && s.builder.TransactionCount() == 0
	if creatingNewStake {
		if err := s.newStake(ctx); err != nil {
			return nil, err
		}
	}

	if rawInfo != nil {
		if err = s.handleConflict(ctx, rawInfo); err != nil {
			return nil, err
		}
	}
	if rawInfo != nil || creatingNewStake {
		// Advance stake up to 20 times in one transaction
		for i := 0; info.CanProgress && i < 20; i++ {
			if err := s.advanceStake(ctx, &info, effectiveStrategy); err != nil {
				return nil, err
			}
		}
	}
	if rawInfo != nil && s.builder.TransactionCount() == 0 {
		if err := s.createConflict(ctx, rawInfo); err != nil {
			return nil, err
		}
	}
	txCount := s.builder.TransactionCount()
	if creatingNewStake {
		// Ignore our stake creation, as it's useless by itself
		txCount--
	}
	if txCount == 0 {
		return nil, nil
	}
	if creatingNewStake {
		logger.Info().Msg("Staking to execute transactions")
	}
	return s.wallet.ExecuteTransactions(ctx, s.builder)
}

func (s *Staker) handleConflict(ctx context.Context, info *ethbridge.StakerInfo) error {
	if info.CurrentChallenge == nil {
		s.activeChallenge = nil
		return nil
	}

	if s.activeChallenge == nil || s.activeChallenge.ChallengeAddress() != *info.CurrentChallenge {
		logger.Warn().Str("challenge", info.CurrentChallenge.String()).Msg("Entered challenge")

		challengeCon, err := ethbridge.NewChallenge(info.CurrentChallenge.ToEthAddress(), s.client, s.builder)
		if err != nil {
			return err
		}

		challengedNode, err := s.rollup.LookupChallengedNode(ctx, *info.CurrentChallenge)
		if err != nil {
			return err
		}

		nodeInfo, err := s.rollup.RollupWatcher.LookupNode(ctx, challengedNode)
		if err != nil {
			return err
		}

		s.activeChallenge = challenge.NewChallenger(challengeCon, s.sequencerInbox, s.lookup, nodeInfo, s.wallet.Address())
	}

	return s.activeChallenge.HandleConflict(ctx)
}

func (s *Staker) newStake(ctx context.Context) error {
	info, err := s.rollup.StakerInfo(ctx, s.wallet.Address())
	if err != nil {
		return err
	}
	if info != nil {
		return nil
	}
	stakeAmount, err := s.rollup.CurrentRequiredStake(ctx)
	if err != nil {
		return err
	}
	return s.rollup.NewStake(ctx, stakeAmount)
}

func (s *Staker) advanceStake(ctx context.Context, info *OurStakerInfo, effectiveStrategy Strategy) error {
	active := effectiveStrategy > WatchtowerStrategy
	action, _, err := s.generateNodeAction(ctx, info, effectiveStrategy)
	if err != nil {
		return err
	}
	// TODO raise an alert if wrongNodesExist (esp for watchtower strategy)
	if action == nil || !active {
		info.CanProgress = false
		return nil
	}

	switch action := action.(type) {
	case createNodeAction:
		// Already logged with more details in generateNodeAction
		info.CanProgress = false
		info.LatestStakedNode = nil
		info.LatestStakedNodeHash = action.hash
		return s.rollup.StakeOnNewNode(ctx, action.hash, action.assertion, action.prevProposedBlock, action.prevInboxMaxCount, action.sequencerBatchProof)
	case existingNodeAction:
		logger.Info().Int("node", int((*big.Int)(action.number).Int64())).Msg("Staking on existing node")
		info.LatestStakedNode = action.number
		info.LatestStakedNodeHash = action.hash
		return s.rollup.StakeOnExistingNode(ctx, action.number, action.hash)
	default:
		panic("invalid action type")
	}
}

func (s *Staker) createConflict(ctx context.Context, info *ethbridge.StakerInfo) error {
	if info.CurrentChallenge != nil {
		return nil
	}

	stakers, err := s.validatorUtils.GetStakers(ctx)
	if err != nil {
		return err
	}
	latestNode, err := s.rollup.LatestConfirmedNode(ctx)
	if err != nil {
		return err
	}
	for _, staker := range stakers {
		stakerInfo, err := s.rollup.StakerInfo(ctx, staker)
		if err != nil {
			return err
		}
		if stakerInfo.CurrentChallenge != nil {
			continue
		}
		conflictType, node1, node2, err := s.validatorUtils.FindStakerConflict(ctx, s.wallet.Address(), staker)
		if err != nil {
			return err
		}
		if conflictType != ethbridge.CONFLICT_TYPE_FOUND {
			continue
		}
		staker1 := s.wallet.Address()
		staker2 := staker
		if node2.Cmp(node1) < 0 {
			staker1, staker2 = staker2, staker1
			node1, node2 = node2, node1
		}
		if node1.Cmp(latestNode) <= 0 {
			// removeOldStakers will take care of them
			continue
		}

		node1Info, err := s.rollup.RollupWatcher.LookupNode(ctx, node1)
		if err != nil {
			return err
		}
		node2Info, err := s.rollup.RollupWatcher.LookupNode(ctx, node2)
		if err != nil {
			return err
		}
		logger.Warn().Int("ourNode", int(node1.Int64())).Int("otherNode", int(node2.Int64())).Str("otherStaker", staker2.String()).Msg("Creating challenge")
		return s.rollup.CreateChallenge(
			ctx,
			staker1,
			node1Info,
			staker2,
			node2Info,
		)
	}
	// No conflicts exist
	return nil
}
