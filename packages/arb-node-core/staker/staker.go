package staker

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
)

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
) (*Staker, *ethbridge.BridgeWatcher, error) {
	val, err := NewValidator(ctx, lookup, client, wallet, validatorUtilsAddress)
	if err != nil {
		return nil, nil, err
	}
	return &Staker{
		Validator: val,
		strategy:  strategy,
	}, val.bridge, nil
}

func (s *Staker) RunInBackground(ctx context.Context) chan bool {
	done := make(chan bool)
	go func() {
		defer func() {
			done <- true
		}()
		for {
			tx, err := s.Act(ctx)
			backoff := time.Second
			if tx != nil {
				// Note: methodName isn't accurate, it's just used for logging
				_, err = ethbridge.WaitForReceiptWithResults(ctx, s.client, s.wallet.From().ToEthAddress(), tx, "for staking")
				if err == nil {
					logger.Info().Str("hash", tx.Hash().String()).Msg("Successfully executed transaction")
				}
			}
			if err != nil {
				logger.Warn().Stack().Err(err).Msg("Staking error (possible reorg?)")
				<-time.After(backoff)
				if backoff < 60*time.Second {
					backoff *= 2
				}
				continue
			} else {
				backoff = time.Second
			}
			if tx != nil {
				// We did something, there's probably something else to do
				<-time.After(time.Second)
			} else {
				// Nothing to do for now
				<-time.After(time.Minute)
			}
		}
	}()
	return done
}

func (s *Staker) Act(ctx context.Context) (*types.Transaction, error) {
	s.builder.ClearTransactions()
	info, err := s.rollup.StakerInfo(ctx, s.wallet.Address())
	if err != nil {
		return nil, err
	}
	creatingNewStake := info == nil
	if creatingNewStake {
		if err := s.newStake(ctx); err != nil {
			return nil, err
		}
	}

	if err := s.resolveNextNode(ctx); err != nil {
		return nil, err
	}

	if info != nil {
		if err = s.handleConflict(ctx, info); err != nil {
			return nil, err
		}
	}
	if err := s.advanceStake(ctx); err != nil {
		return nil, err
	}
	if info != nil && s.builder.TransactionCount() == 0 {
		if err := s.createConflict(ctx, info); err != nil {
			return nil, err
		}
	}
	txCount := s.builder.TransactionCount()
	if creatingNewStake {
		// Ignore our stake creation, as it's useless by itself
		txCount--
	}
	if txCount == 0 {
		if info != nil {
			tx, err := s.removeOldStakers(ctx)
			if err != nil || tx != nil {
				return tx, err
			}
			tx, err = s.resolveTimedOutChallenges(ctx)
			if err != nil || tx != nil {
				return tx, err
			}
		}
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

		s.activeChallenge = challenge.NewChallenger(challengeCon, s.lookup, nodeInfo, s.wallet.Address())
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

func (s *Staker) advanceStake(ctx context.Context) error {
	active := s.strategy > WatchtowerStrategy
	action, _, err := s.generateNodeAction(ctx, s.wallet.Address(), s.strategy)
	if err != nil {
		return err
	}
	// TODO raise an alert if wrongNodesExist (esp for watchtower strategy)
	if action == nil || !active {
		return nil
	}

	switch action := action.(type) {
	case createNodeAction:
		// Already logged with more details in generateNodeAction
		return s.rollup.StakeOnNewNode(ctx, action.hash, action.assertion)
	case existingNodeAction:
		logger.Info().Int("node", int((*big.Int)(action.number).Int64())).Msg("Staking on existing node")
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
