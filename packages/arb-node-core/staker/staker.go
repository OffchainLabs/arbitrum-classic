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
	"runtime"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
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
	activeChallenge     *challenge.Challenger
	strategy            Strategy
	fromBlock           int64
	baseCallOpts        bind.CallOpts
	auth                *ethbridge.TransactAuth
	config              configuration.Validator
	highGasBlocksBuffer *big.Int
	lastActCalledBlock  *big.Int
}

func NewStaker(
	ctx context.Context,
	lookup core.ArbCoreLookup,
	client ethutils.EthClient,
	wallet *ethbridge.ValidatorWallet,
	fromBlock int64,
	validatorUtilsAddress common.Address,
	strategy Strategy,
	callOpts bind.CallOpts,
	auth *ethbridge.TransactAuth,
	config configuration.Validator,
) (*Staker, *ethbridge.DelayedBridgeWatcher, error) {
	val, err := NewValidator(ctx, lookup, client, wallet, fromBlock, validatorUtilsAddress, callOpts)
	if err != nil {
		return nil, nil, err
	}
	return &Staker{
		Validator:           val,
		strategy:            strategy,
		fromBlock:           fromBlock,
		baseCallOpts:        callOpts,
		auth:                auth,
		config:              config,
		highGasBlocksBuffer: big.NewInt(config.L1PostingStrategy.HighGasDelayBlocks),
		lastActCalledBlock:  nil,
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
				_, err = ethbridge.WaitForReceiptWithResultsAndReplaceByFee(ctx, s.client, s.wallet.From().ToEthAddress(), tx, "for staking", s.auth)
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
			// Force a GC run to clean up any execution cursors while we wait
			delay := time.After(time.Minute)
			runtime.GC()
			<-delay
		}
	}()
	return done
}

func (s *Staker) shouldAct(ctx context.Context) bool {
	var gasPriceHigh = false
	var gasPriceFloat float64
	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		logger.Warn().Err(err).Msg("error getting gas price")
	} else {
		gasPriceFloat = float64(gasPrice.Int64()) / 1e9
		if gasPriceFloat >= s.config.L1PostingStrategy.HighGasThreshold {
			gasPriceHigh = true
		}
	}
	latestBlockInfo, err := s.client.BlockInfoByNumber(ctx, nil)
	if err != nil {
		logger.Warn().Err(err).Msg("error getting latest block")
		return true
	}
	latestBlockNum := latestBlockInfo.Number.ToInt()
	if s.lastActCalledBlock == nil {
		s.lastActCalledBlock = latestBlockNum
	}
	blocksSinceActCalled := new(big.Int).Sub(latestBlockNum, s.lastActCalledBlock)
	if gasPriceHigh {
		// We're eating into the high gas buffer to delay our tx
		s.highGasBlocksBuffer.Sub(s.highGasBlocksBuffer, blocksSinceActCalled)
	} else {
		// We'll make a tx if necessary, so we can add to the buffer for future high gas
		s.highGasBlocksBuffer.Add(s.highGasBlocksBuffer, blocksSinceActCalled)
	}
	// Clamp `s.highGasBlocksBuffer` to between 0 and HighGasDelayBlocks
	if s.highGasBlocksBuffer.Sign() < 0 {
		s.highGasBlocksBuffer.SetInt64(0)
	} else if s.highGasBlocksBuffer.Cmp(big.NewInt(s.config.L1PostingStrategy.HighGasDelayBlocks)) > 0 {
		s.highGasBlocksBuffer.SetInt64(s.config.L1PostingStrategy.HighGasDelayBlocks)
	}
	if gasPriceHigh && s.highGasBlocksBuffer.Sign() > 0 {
		logger.
			Info().
			Float64("gasPrice", gasPriceFloat).
			Float64("highGasPriceConfig", s.config.L1PostingStrategy.HighGasThreshold).
			Str("highGasBuffer", s.highGasBlocksBuffer.String()).
			Msg("not acting yet as gas price is high")
		return false
	} else {
		return true
	}
}

func (s *Staker) Act(ctx context.Context) (*types.Transaction, error) {
	if !s.shouldAct(ctx) {
		// The fact that we're delaying acting is alreay logged in `shouldAct`
		return nil, nil
	}
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
		// Keep the stake of this validator placed if we plan on staking further
		tx, err := s.removeOldStakers(ctx, effectiveStrategy >= StakeLatestStrategy)
		if err != nil || tx != nil {
			return tx, err
		}
		tx, err = s.resolveTimedOutChallenges(ctx)
		if err != nil || tx != nil {
			return tx, err
		}
		if err := s.resolveNextNode(ctx, rawInfo, s.fromBlock); err != nil {
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

		challengeCon, err := ethbridge.NewChallenge(info.CurrentChallenge.ToEthAddress(), s.fromBlock, s.client, s.builder, s.baseCallOpts)
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
	action, _, err := s.generateNodeAction(ctx, info, effectiveStrategy, s.fromBlock)
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
