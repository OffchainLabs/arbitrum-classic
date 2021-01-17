package staker

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/pkg/errors"
)

type Staker struct {
	*Validator
	makeNewNodes bool

	activeChallenge *challenge.Challenger
}

func NewStaker(
	ctx context.Context,
	lookup core.ValidatorLookup,
	client ethutils.EthClient,
	wallet *ethbridge.Validator,
	validatorUtilsAddress common.Address,
) (*Staker, error) {
	val, err := NewValidator(ctx, lookup, client, wallet, validatorUtilsAddress)
	if err != nil {
		return nil, err
	}
	return &Staker{
		Validator:    val,
		makeNewNodes: true,
	}, nil
}

func (s *Staker) Act(ctx context.Context) (*types.Transaction, error) {
	txes := make([]*ethbridge.RawTransaction, 0)
	info, err := s.rollup.StakerInfo(ctx, s.wallet.Address())
	if err != nil {
		return nil, err
	}
	rawTx, err := s.handleConflict(ctx, info)
	if err != nil {
		return nil, err
	}
	if rawTx != nil {
		txes = append(txes, rawTx)
	}

	rawTx, err = s.resolveNextNode(ctx)
	if err != nil {
		return nil, err
	}
	if rawTx != nil {
		txes = append(txes, rawTx)
	}

	if info != nil {
		rawTx, err := s.advanceStake(ctx, info)
		if err != nil {
			return nil, err
		}
		if rawTx != nil {
			txes = append(txes, rawTx)
		}
	}
	return s.wallet.ExecuteTransactions(ctx, txes)
}

func (s *Staker) handleConflict(ctx context.Context, info *ethbridge.StakerInfo) (*ethbridge.RawTransaction, error) {
	if info.CurrentChallenge == nil {
		s.activeChallenge = nil
		return nil, nil
	}

	if s.activeChallenge != nil || s.activeChallenge.ChallengeAddress() != *info.CurrentChallenge {
		challengeCon, err := ethbridge.NewChallenge(info.CurrentChallenge.ToEthAddress(), s.client)
		if err != nil {
			return nil, err
		}

		challengedNode, err := s.rollup.LookupChallengedNode(ctx, *info.CurrentChallenge)
		if err != nil {
			return nil, err
		}

		nodeInfo, err := lookupNode(ctx, s.rollup.RollupWatcher, challengedNode)
		if err != nil {
			return nil, err
		}

		s.activeChallenge = challenge.NewChallenger(challengeCon, s.lookup, nodeInfo, s.wallet.Address())
	}

	return s.activeChallenge.HandleConflict(ctx)
}

func (s *Staker) newStake(ctx context.Context) (*ethbridge.RawTransaction, error) {
	info, err := s.rollup.StakerInfo(ctx, s.wallet.Address())
	if err != nil {
		return nil, err
	}
	if info != nil {
		return nil, nil
	}
	stakeAmount, err := s.rollup.CurrentRequiredStake(ctx)
	if err != nil {
		return nil, err
	}
	return s.rollup.NewStake(ctx, stakeAmount)
}

func (s *Staker) advanceStake(ctx context.Context, info *ethbridge.StakerInfo) (*ethbridge.RawTransaction, error) {
	info, err := s.rollup.StakerInfo(ctx, s.wallet.Address())
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, errors.New("no stake placed")
	}

	action, err := s.generateNodeAction(ctx, info.LatestStakedNode, true)
	if err != nil || action == nil {
		return nil, err
	}

	switch action := action.(type) {
	case *nodeCreationInfo:
		if !s.makeNewNodes {
			return nil, nil
		}
		return s.rollup.StakeOnNewNode(action.block, action.newNodeID, action.assertion)
	case *nodeMovementInfo:
		return s.rollup.StakeOnExistingNode(action.block, action.nodeNum)
	default:
		panic("invalid type")
	}
}

func (s *Staker) createConflict(ctx context.Context) (*ethbridge.RawTransaction, error) {
	info, err := s.rollup.StakerInfo(ctx, s.wallet.Address())
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, errors.New("not staked")
	}
	if info.CurrentChallenge != nil {
		return nil, nil
	}

	stakers, err := s.validatorUtils.GetStakers(ctx)
	if err != nil {
		return nil, err
	}
	for _, staker := range stakers {
		conflictType, node1, node2, err := s.validatorUtils.FindStakerConflict(ctx, s.wallet.Address(), staker)
		if err != nil {
			return nil, err
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

		nodeInfo, err := lookupNode(ctx, s.rollup.RollupWatcher, node1)
		if err != nil {
			return nil, err
		}
		maxInboxHash, err := s.lookup.GetInboxAcc(nodeInfo.InboxMaxCount)
		if err != nil {
			return nil, err
		}
		return s.rollup.CreateChallenge(
			staker1,
			node1,
			staker2,
			node2,
			nodeInfo.Assertion,
			maxInboxHash,
			nodeInfo.InboxMaxCount,
		)
	}
	// No conflicts exist
	return nil, nil
}
