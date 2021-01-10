package validator

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/pkg/errors"
	"math/big"
)

type ValidatorLookup interface {
	GenerateLogAccumulator(startIndex *big.Int, count *big.Int) (common.Hash, error)
	GetSends(startIndex *big.Int, count *big.Int) ([][]byte, error)
	GetInboxAcc(index *big.Int) (common.Hash, error)
	GetMessages(startIndex *big.Int, count *big.Int) ([]inbox.InboxMessage, error)

	GetMachine(gasUsed *big.Int) (machine.Machine, error)
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
		logAcc, err := v.lookup.GenerateLogAccumulator(nodeInfo.Assertion.BeforeTotalLogCount, nodeInfo.Assertion.LogCount)
		if err != nil {
			return nil, err
		}
		sends, err := v.lookup.GetSends(nodeInfo.Assertion.BeforeTotalSendCount, nodeInfo.Assertion.SendCount)
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
		nodes, err := s.rollup.LookupNodes(ctx, successors)
		if err != nil {
			return err
		}

		for _, nd := range nodes {
			afterInboxHash, err := s.lookup.GetInboxAcc(nd.Assertion.AfterInboxCount())
			if err != nil {
				return err
			}
			if nd.Assertion.AfterInboxHash != afterInboxHash {
				// Failed inbox consistency
				continue
			}
			messages, err := s.lookup.GetMessages(nd.Assertion.BeforeInboxCount, nd.Assertion.InboxMessagesRead)
			if err != nil {
				return err
			}
			if nd.Assertion.InboxDelta != calculateInboxDeltaAcc(messages) {
				// Failed inbox delta
				continue
			}

		}
	}
	return nil
}

func calculateInboxDeltaAcc(messages []inbox.InboxMessage) common.Hash {
	acc := common.Hash{}
	for i := range messages {
		valHash := messages[len(messages)-1-i].AsValue().Hash()
		acc = hashing.SoliditySHA3(hashing.Bytes32(acc), hashing.Bytes32(valHash))
	}
	return acc
}
