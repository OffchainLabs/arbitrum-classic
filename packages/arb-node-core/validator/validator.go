package validator

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/pkg/errors"
	"math/big"
)

type ValidatorLookup interface {
	GenerateLogAccumulator(startIndex *big.Int, count *big.Int) (common.Hash, error)
	GetSends(startIndex *big.Int, count *big.Int) ([][]byte, error)
}

type Validator struct {
	rollup         *ethbridge.Rollup
	validatorUtils *ethbridge.ValidatorUtils
	client         ethutils.EthClient
	lookup         ValidatorLookup
}

func (v *Validator) removeOldStakers(ctx context.Context) (*types.Transaction, error) {
	stakersToEliminate, err := v.validatorUtils.RefundableStakers(ctx)
	if err != nil {
		return nil, err
	}
	if len(stakersToEliminate) == 0 {
		return nil, nil
	}
	return v.validatorUtils.RefundStakers(ctx, stakersToEliminate)
}

func (v *Validator) resolveNextNode(ctx context.Context) (*types.Transaction, error) {
	confirmType, successorWithStake, stakerAddress, err := v.validatorUtils.CheckDecidableNextNode(ctx)
	if err != nil {
		return nil, err
	}
	switch confirmType {
	case ethbridge.CONFIRM_TYPE_OUT_OF_ORDER:
		return v.rollup.RejectNextNodeOutOfOrder(ctx)
	case ethbridge.CONFIRM_TYPE_INVALID:
		return v.rollup.RejectNextNode(ctx, successorWithStake, stakerAddress)
	case ethbridge.CONFIRM_TYPE_VALID:
		unresolvedNodeIndex, err := v.rollup.FirstUnresolvedNode(ctx)
		if err != nil {
			return nil, err
		}
		nodesInfo, err := v.rollup.LookupNodes(ctx, []*big.Int{unresolvedNodeIndex})
		if err != nil {
			return nil, err
		}
		if len(nodesInfo) != 1 {
			return nil, errors.New("bad node query")
		}
		nodeInfo := nodesInfo[0]
		logAcc, err := v.lookup.GenerateLogAccumulator(nodeInfo.Assertion.BeforeLogCount, nodeInfo.Assertion.LogCount)
		if err != nil {
			return nil, err
		}
		sends, err := v.lookup.GetSends(nodeInfo.Assertion.BeforeSendCount, nodeInfo.Assertion.SendCount)
		if err != nil {
			return nil, err
		}
		return v.rollup.ConfirmNextNode(ctx, logAcc, sends)
	default:
		return nil, nil
	}
}

type Staker struct {
	address        common.Address
	rollup         *ethbridge.Rollup
	validatorUtils *ethbridge.ValidatorUtils
	client         ethutils.EthClient
	lookup         ValidatorLookup
}

func (s *Staker) act(ctx context.Context) error {
	info, err := s.rollup.StakerInfo(ctx, s.address)
	if err != nil {
		return nil
	}
	if info != nil {
		successors, err := s.validatorUtils.SuccessorNodes(ctx, info.LatestStakedNode)
		if err != nil {
			return err
		}
		_, err = s.rollup.LookupNodes(ctx, successors)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Staker) validateNode(ctx context.Context, info *ethbridgecontracts.RollupNodeCreated) error {
	return nil
}
