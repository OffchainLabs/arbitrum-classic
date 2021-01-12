package staker

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/pkg/errors"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type Validator struct {
	rollup         *ethbridge.Rollup
	validatorUtils *ethbridge.ValidatorUtils
	client         ethutils.EthClient
	lookup         core.ValidatorLookup
}

func NewValidator(
	lookup core.ValidatorLookup,
	client ethutils.EthClient,
	auth *ethbridge.TransactAuth,
	rollupAddress,
	validatorUtilsAddress common.Address,
) (*Validator, error) {
	rollup, err := ethbridge.NewRollup(rollupAddress.ToEthAddress(), client, auth)
	if err != nil {
		return nil, err
	}
	validatorUtils, err := ethbridge.NewValidatorUtils(
		validatorUtilsAddress.ToEthAddress(),
		rollupAddress.ToEthAddress(),
		client,
		auth,
	)
	if err != nil {
		return nil, err
	}
	return &Validator{
		rollup:         rollup,
		validatorUtils: validatorUtils,
		client:         client,
		lookup:         lookup,
	}, nil
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
		logAcc, err := v.lookup.GenerateLogAccumulator(nodeInfo.Assertion.PrevState.TotalLogCount, nodeInfo.Assertion.ExecInfo.LogCount)
		if err != nil {
			return nil, err
		}
		sends, err := v.lookup.GetSends(nodeInfo.Assertion.PrevState.TotalSendCount, nodeInfo.Assertion.ExecInfo.SendCount)
		if err != nil {
			return nil, err
		}
		return v.rollup.ConfirmNextNode(ctx, logAcc, sends)
	default:
		return nil, nil
	}
}

type nodeCreationInfo struct {
	assertion *core.Assertion
	block     *common.BlockId
	newNodeID core.NodeID
}

type nodeMovementInfo struct {
	block   *common.BlockId
	nodeNum core.NodeID
}

type nodeActionInfo interface {
}

func (v *Validator) generateNodeAction(ctx context.Context, base core.NodeID) (nodeActionInfo, error) {
	currentNode, err := lookupNode(ctx, v.rollup.RollupWatcher, base)
	if err != nil {
		return nil, err
	}
	mach, err := v.lookup.GetMachine(currentNode.Assertion.AfterTotalGasUsed())
	if err != nil {
		return nil, err
	}
	if mach.Hash() != currentNode.Assertion.ExecInfo.AfterMachineHash {
		return nil, errors.New("local machine doesn't match chain")
	}

	successorsIndexes, err := v.validatorUtils.SuccessorNodes(ctx, base)
	if err != nil {
		return nil, err
	}
	successorsNodes, err := v.rollup.LookupNodes(ctx, successorsIndexes)
	if err != nil {
		return nil, err
	}
	validChild, err := selectValidNode(v.lookup, successorsNodes, mach)
	if err != nil {
		return nil, err
	}
	if validChild != nil {
		blockId, err := getBlockID(ctx, v.client, validChild.Assertion.PrevState.ProposedBlock)
		if err != nil {
			return nil, err
		}
		return &nodeMovementInfo{
			block:   blockId,
			nodeNum: validChild.NodeNum,
		}, nil
	}

	minAssertionPeriod, err := v.rollup.MinimumAssertionPeriod(ctx)
	if err != nil {
		return nil, err
	}
	arbGasSpeedLimitPerBlock, err := v.rollup.ArbGasSpeedLimitPerBlock(ctx)
	if err != nil {
		return nil, err
	}

	currentBlockId, err := getBlockID(ctx, v.client, nil)
	if err != nil {
		return nil, err
	}

	assertion, err := createAssertion(v.lookup, currentNode, currentBlockId.Height.AsInt(), minAssertionPeriod, arbGasSpeedLimitPerBlock)
	if err != nil {
		return nil, err
	}

	lastNodeCreated, err := v.rollup.LatestNodeCreated(ctx)
	if err != nil {
		return nil, err
	}

	msg, err := lookupMessageByNum(ctx, v.rollup.RollupWatcher, assertion.AfterInboxCount())
	if err != nil {
		return nil, err
	}
	newNodeID := new(big.Int).Add(lastNodeCreated, big.NewInt(1))
	return &nodeCreationInfo{
		assertion: assertion,
		block:     msg.Block(),
		newNodeID: newNodeID,
	}, nil
}

func selectValidNode(lookup core.ValidatorLookup, nodes []*core.NodeInfo, startMach machine.Machine) (*core.NodeInfo, error) {
	for _, nd := range nodes {
		chalType, err := core.JudgeAssertion(lookup, nd.Assertion, startMach)
		if err != nil {
			return nil, err
		}
		if chalType == core.NO_CHALLENGE {
			return nd, nil
		}
	}
	return nil, nil
}

func createAssertion(lookup core.ValidatorLookup, prevNode *core.NodeInfo, currentBlock, minAssertionPeriod, arbGasSpeedLimitPerBlock *big.Int) (*core.Assertion, error) {
	timeSinceProposed := new(big.Int).Sub(currentBlock, prevNode.BlockProposed.Height.AsInt())
	if timeSinceProposed.Cmp(minAssertionPeriod) < 0 {
		// Too soon to assert
		return nil, nil
	}

	mach, err := lookup.GetMachine(prevNode.Assertion.AfterTotalGasUsed())
	if err != nil {
		return nil, err
	}
	if mach.Hash() != prevNode.Assertion.ExecInfo.AfterMachineHash {
		return nil, errors.New("local machine doesn't match chain")
	}

	minimumGasToConsume := new(big.Int).Mul(timeSinceProposed, arbGasSpeedLimitPerBlock)
	minMessages := new(big.Int).Sub(prevNode.InboxMaxCount, prevNode.Assertion.AfterInboxCount())

	maximumGasToConsume := new(big.Int).Mul(minimumGasToConsume, big.NewInt(4))

	assertionInfo, err := lookup.GetExecutionInfo(mach, maximumGasToConsume)
	if err != nil {
		return nil, err
	}

	if assertionInfo.ExecInfo.GasUsed.Cmp(minimumGasToConsume) < 0 && assertionInfo.ExecInfo.InboxMessagesRead.Cmp(minMessages) < 0 {
		// Couldn't execute far enough
		return nil, nil
	}

	return &core.Assertion{
		PrevState:     prevNode.AfterState(),
		AssertionInfo: assertionInfo,
	}, nil
}

func getBlockID(ctx context.Context, client ethutils.EthClient, number *big.Int) (*common.BlockId, error) {
	blockInfo, err := client.BlockInfoByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	return &common.BlockId{
		Height:     common.NewTimeBlocks((*big.Int)(blockInfo.Number)),
		HeaderHash: common.NewHashFromEth(blockInfo.Hash),
	}, nil
}

func lookupNode(ctx context.Context, rollup *ethbridge.RollupWatcher, node core.NodeID) (*core.NodeInfo, error) {
	currentNodes, err := rollup.LookupNodes(ctx, []*big.Int{node})
	if err != nil {
		return nil, err
	}
	if len(currentNodes) == 0 {
		return nil, errors.New("no matching node")
	}
	if len(currentNodes) > 1 {
		return nil, errors.New("too many matching nodes")
	}
	return currentNodes[0], nil
}

func lookupMessageByNum(ctx context.Context, rollup *ethbridge.RollupWatcher, messageNum *big.Int) (*ethbridge.DeliveredInboxMessage, error) {
	messages, err := rollup.LookupMessagesByNum(ctx, []*big.Int{messageNum})
	if err != nil {
		return nil, err
	}
	if len(messages) == 0 {
		return nil, errors.New("no matching message")
	}
	if len(messages) > 1 {
		return nil, errors.New("too many matching messages")
	}
	return messages[0], nil
}
