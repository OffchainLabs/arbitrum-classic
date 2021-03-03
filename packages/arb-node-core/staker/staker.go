package staker

import (
	"context"

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
) (*Staker, error) {
	val, err := NewValidator(ctx, lookup, client, wallet, validatorUtilsAddress)
	if err != nil {
		return nil, err
	}
	return &Staker{
		Validator: val,
		strategy:  strategy,
	}, nil
}

func (s *Staker) Act(ctx context.Context) (*types.Transaction, error) {
	s.builder.ClearTransactions()
	info, err := s.rollup.StakerInfo(ctx, s.wallet.Address())
	if err != nil {
		return nil, err
	}
	creatingNewStake := false
	if info == nil && s.strategy >= StakeLatestStrategy {
		creatingNewStake = true
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
	if txCount == 0 && info != nil {
		tx, err := s.removeOldStakers(ctx)
		if err != nil || tx == nil {
			return tx, err
		}
		return nil, nil
	}
	return s.wallet.ExecuteTransactions(ctx, s.builder)
}

func (s *Staker) handleConflict(ctx context.Context, info *ethbridge.StakerInfo) error {
	if info.CurrentChallenge == nil {
		s.activeChallenge = nil
		return nil
	}

	if s.activeChallenge == nil || s.activeChallenge.ChallengeAddress() != *info.CurrentChallenge {
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
	action, wrongNodesExist, err := s.generateNodeAction(ctx, s.wallet.Address(), active, s.strategy == MakeNodesStrategy)
	if err != nil {
		return err
	}
	// TODO raise an alert if wrongNodesExist (esp for watchtower strategy)
	if action == nil || !active {
		return nil
	}

	switch action := action.(type) {
	case createNodeAction:
		if !wrongNodesExist && s.strategy < StakeLatestStrategy {
			return nil
		}
		return s.rollup.StakeOnNewNode(ctx, action.hash, action.assertion)
	case existingNodeAction:
		if !wrongNodesExist && s.strategy < StakeLatestStrategy {
			return nil
		}
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
		if node1.Cmp(latestNode) < 0 {
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
