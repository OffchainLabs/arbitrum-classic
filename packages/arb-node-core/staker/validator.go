package staker

import (
	"context"
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
	bridge         *ethbridge.BridgeWatcher
	validatorUtils *ethbridge.ValidatorUtils
	client         ethutils.EthClient
	lookup         core.ValidatorLookup
	wallet         *ethbridge.Validator
}

func NewValidator(
	ctx context.Context,
	lookup core.ValidatorLookup,
	client ethutils.EthClient,
	wallet *ethbridge.Validator,
	validatorUtilsAddress common.Address,
) (*Validator, error) {
	rollup, err := ethbridge.NewRollup(wallet.RollupAddress().ToEthAddress(), client)
	if err != nil {
		return nil, err
	}
	bridgeAddress, err := rollup.Bridge(ctx)
	if err != nil {
		return nil, err
	}
	bridge, err := ethbridge.NewBridgeWatcher(bridgeAddress.ToEthAddress(), client)
	if err != nil {
		return nil, err
	}
	validatorUtils, err := ethbridge.NewValidatorUtils(
		validatorUtilsAddress.ToEthAddress(),
		wallet.RollupAddress().ToEthAddress(),
		client,
	)
	if err != nil {
		return nil, err
	}
	return &Validator{
		rollup:         rollup,
		bridge:         bridge,
		validatorUtils: validatorUtils,
		client:         client,
		lookup:         lookup,
		wallet:         wallet,
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
	return v.wallet.ReturnOldDeposits(ctx, stakersToEliminate)
}

func (v *Validator) resolveNextNode(ctx context.Context) (*ethbridge.RawTransaction, error) {
	confirmType, successorWithStake, stakerAddress, err := v.validatorUtils.CheckDecidableNextNode(ctx)
	if err != nil {
		return nil, err
	}
	switch confirmType {
	case ethbridge.CONFIRM_TYPE_INVALID:
		return v.rollup.RejectNextNode(successorWithStake, stakerAddress)
	case ethbridge.CONFIRM_TYPE_VALID:
		unresolvedNodeIndex, err := v.rollup.FirstUnresolvedNode(ctx)
		if err != nil {
			return nil, err
		}
		nodeInfo, err := lookupNode(ctx, v.rollup.RollupWatcher, unresolvedNodeIndex)
		if err != nil {
			return nil, err
		}
		logAcc, err := v.lookup.GetLogAcc(common.Hash{}, nodeInfo.Assertion.Before.TotalLogCount, nodeInfo.Assertion.LogCount())
		if err != nil {
			return nil, err
		}
		sends, err := v.lookup.GetSends(nodeInfo.Assertion.Before.TotalSendCount, nodeInfo.Assertion.SendCount())
		if err != nil {
			return nil, err
		}
		return v.rollup.ConfirmNextNode(logAcc, sends)
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

func (v *Validator) generateNodeAction(ctx context.Context, base core.NodeID, maybeMakeNode bool) (nodeActionInfo, error) {
	startState, err := lookupNodeStartState(ctx, v.rollup.RollupWatcher, base)
	if err != nil {
		return nil, err
	}

	cursor, err := v.lookup.GetCursor(startState.TotalGasConsumed)
	if err != nil {
		return nil, err
	}
	if cursor.MachineHash() != startState.MachineHash {
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

	gasesUsed := make([]*big.Int, 0, len(successorsNodes))
	for _, nd := range successorsNodes {
		gasesUsed = append(gasesUsed, nd.Assertion.GasUsed())
	}

	currentBlockId, err := getBlockID(ctx, v.client, nil)
	if err != nil {
		return nil, err
	}

	minAssertionPeriod, err := v.rollup.MinimumAssertionPeriod(ctx)
	if err != nil {
		return nil, err
	}

	timeSinceProposed := new(big.Int).Sub(currentBlockId.Height.AsInt(), startState.ProposedBlock)
	if timeSinceProposed.Cmp(minAssertionPeriod) < 0 {
		// Too soon to assert
		return nil, nil
	}

	arbGasSpeedLimitPerBlock, err := v.rollup.ArbGasSpeedLimitPerBlock(ctx)
	if err != nil {
		return nil, err
	}

	minMessages := new(big.Int).Sub(startState.InboxMaxCount, startState.InboxIndex)
	minimumGasToConsume := new(big.Int).Mul(timeSinceProposed, arbGasSpeedLimitPerBlock)
	maximumGasToConsume := new(big.Int).Mul(minimumGasToConsume, big.NewInt(4))

	if maybeMakeNode {
		gasesUsed = append(gasesUsed, maximumGasToConsume)
	}

	execTracker := core.NewExecutionTracker(v.lookup, cursor, false, gasesUsed)

	for _, nd := range successorsNodes {
		chalType, err := core.JudgeAssertion(v.lookup, nd.Assertion, execTracker)
		if err != nil {
			return nil, err
		}
		if chalType == core.NO_CHALLENGE {
			return nd, nil
		}
		blockId, err := getBlockID(ctx, v.client, nd.Assertion.PrevProposedBlock)
		if err != nil {
			return nil, err
		}
		return &nodeMovementInfo{
			block:   blockId,
			nodeNum: nd.NodeNum,
		}, nil
	}

	if !maybeMakeNode {
		return nil, nil
	}

	execInfo, err := execTracker.GetExecutionInfo(gasesUsed[len(gasesUsed)-1])
	if err != nil {
		return nil, err
	}

	if execInfo.GasUsed().Cmp(minimumGasToConsume) < 0 && execInfo.InboxMessagesRead().Cmp(minMessages) < 0 {
		// Couldn't execute far enough
		return nil, nil
	}

	inboxDelta, err := v.lookup.GetInboxDelta(execInfo.Before.InboxIndex, execInfo.InboxMessagesRead())
	if err != nil {
		return nil, err
	}
	lastNodeCreated, err := v.rollup.LatestNodeCreated(ctx)
	if err != nil {
		return nil, err
	}

	msgBlock, err := v.bridge.LookupMessageBlock(ctx, execInfo.After.InboxIndex)
	if err != nil {
		return nil, err
	}
	newNodeID := new(big.Int).Add(lastNodeCreated, big.NewInt(1))
	return &nodeCreationInfo{
		assertion: &core.Assertion{
			PrevProposedBlock: startState.ProposedBlock,
			PrevInboxMaxCount: startState.InboxMaxCount,
			ExecutionInfo:     execInfo,
			InboxDelta:        inboxDelta,
		},
		block:     msgBlock,
		newNodeID: newNodeID,
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

func lookupNodeStartState(ctx context.Context, rollup *ethbridge.RollupWatcher, nodeNum *big.Int) (*core.NodeState, error) {
	if nodeNum.Cmp(big.NewInt(0)) == 0 {
		creationEvent, err := rollup.LookupCreation(ctx)
		if err != nil {
			return nil, err
		}
		return &core.NodeState{
			ProposedBlock: new(big.Int).SetUint64(creationEvent.Raw.BlockNumber),
			InboxMaxCount: big.NewInt(1),
			ExecutionState: &core.ExecutionState{
				TotalGasConsumed: big.NewInt(0),
				MachineHash:      creationEvent.MachineHash,
				InboxHash:        common.Hash{},
				InboxIndex:       big.NewInt(1),
				TotalSendCount:   big.NewInt(0),
				TotalLogCount:    big.NewInt(0),
			},
		}, nil
	}
	node, err := lookupNode(ctx, rollup, nodeNum)
	if err != nil {
		return nil, err
	}
	return node.AfterState(), nil
}
