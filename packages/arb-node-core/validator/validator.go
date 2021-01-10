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

	GetMachine(totalGasUsed *big.Int) (machine.Machine, error)
	GetExecutionInfo(startMachine machine.Machine, gas *big.Int) (*ethbridge.ExecutionInfo, error)
	GetExecutionInfoInRange(startMachine machine.Machine, minGas, maxGas *big.Int) (*ethbridge.AssertionInfo, error)
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

type Staker struct {
	address        common.Address
	rollup         *ethbridge.Rollup
	validatorUtils *ethbridge.ValidatorUtils
	client         ethutils.EthClient
	lookup         ValidatorLookup
}

func (s *Staker) advanceStake(ctx context.Context) (*types.Transaction, error) {
	info, err := s.rollup.StakerInfo(ctx, s.address)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, errors.New("no stake placed")
	}

	action, err := s.generateNodeAction(ctx, info.LatestStakedNode)
	if err != nil {
		return nil, err
	}

	switch action := action.(type) {
	case nodeCreationInfo:
		return s.rollup.AddStakeOnNewNode(ctx, action.block, action.newNodeID, action.assertion)
	case nodeMovementInfo:
		return s.rollup.AddStakeOnExistingNode(ctx, action.block, action.nodeNum)
	default:
		panic("invalid type")
	}
}

func (s *Staker) placeStake(ctx context.Context) (*types.Transaction, error) {
	info, err := s.rollup.StakerInfo(ctx, s.address)
	if err != nil {
		return nil, err
	}
	if info != nil {
		return nil, errors.New("stake already placed")
	}

	latestConfirmedNode, err := s.rollup.LatestConfirmedNode(ctx)
	if err != nil {
		return nil, err
	}

	action, err := s.generateNodeAction(ctx, latestConfirmedNode)
	if err != nil {
		return nil, err
	}

	switch action := action.(type) {
	case nodeCreationInfo:
		return s.rollup.NewStakeOnNewNode(ctx, action.block, action.newNodeID, latestConfirmedNode, action.assertion)
	case nodeMovementInfo:
		return s.rollup.NewStakeOnExistingNode(ctx, action.block, action.nodeNum)
	default:
		panic("invalid type")
	}
}

type nodeCreationInfo struct {
	assertion *ethbridge.Assertion
	block     *common.BlockId
	newNodeID ethbridge.NodeID
}

type nodeMovementInfo struct {
	block   *common.BlockId
	nodeNum ethbridge.NodeID
}

type nodeActionInfo interface {
}

func (s *Staker) generateNodeAction(ctx context.Context, base ethbridge.NodeID) (nodeActionInfo, error) {
	lastNodeCreated, err := s.rollup.LatestNodeCreated(ctx)
	if err != nil {
		return nil, err
	}

	validChild, err := s.selectValidChild(ctx, base)
	if err != nil {
		return nil, err
	}
	if validChild != nil {
		blockId, err := GetBlockID(ctx, s.client, validChild.Assertion.PrevState.ProposedBlock)
		if err != nil {
			return nil, err
		}
		return &nodeMovementInfo{
			block:   blockId,
			nodeNum: validChild.NodeNum,
		}, nil
	}

	currentNode, err := s.lookupNode(ctx, base)
	if err != nil {
		return nil, err
	}
	minAssertionPeriod, err := s.rollup.MinimumAssertionPeriod(ctx)
	if err != nil {
		return nil, err
	}
	arbGasSpeedLimitPerBlock, err := s.rollup.ArbGasSpeedLimitPerBlock(ctx)
	if err != nil {
		return nil, err
	}

	currentBlockId, err := GetBlockID(ctx, s.client, nil)
	if err != nil {
		return nil, err
	}
	timeSinceProposed := new(big.Int).Sub(currentBlockId.Height.AsInt(), currentNode.BlockProposed.Height.AsInt())
	if timeSinceProposed.Cmp(minAssertionPeriod) < 0 {
		// Too soon to assert
		return nil, nil
	}

	mach, err := s.lookup.GetMachine(currentNode.Assertion.AfterTotalGasUsed())
	if err != nil {
		return nil, err
	}
	if mach.Hash() != currentNode.Assertion.ExecInfo.AfterMachineHash {
		return nil, errors.New("local machine doesn't match chain")
	}

	minimumGasToConsume := new(big.Int).Mul(timeSinceProposed, arbGasSpeedLimitPerBlock)
	maximumGasToConsume := new(big.Int).Mul(minimumGasToConsume, big.NewInt(4))

	assertionInfo, err := s.lookup.GetExecutionInfoInRange(mach, minimumGasToConsume, maximumGasToConsume)
	if err != nil {
		return nil, err
	}

	assertion := &ethbridge.Assertion{
		PrevState:     currentNode.AfterState(),
		AssertionInfo: assertionInfo,
	}
	msg, err := s.lookupMessageByNum(ctx, assertion.AfterInboxCount())
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

func GetBlockID(ctx context.Context, client ethutils.EthClient, number *big.Int) (*common.BlockId, error) {
	blockInfo, err := client.BlockInfoByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	return &common.BlockId{
		Height:     common.NewTimeBlocks((*big.Int)(blockInfo.Number)),
		HeaderHash: common.NewHashFromEth(blockInfo.Hash),
	}, nil
}

func (s *Staker) lookupNode(ctx context.Context, node ethbridge.NodeID) (*ethbridge.NodeInfo, error) {
	currentNodes, err := s.rollup.LookupNodes(ctx, []*big.Int{node})
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

func (s *Staker) lookupMessageByNum(ctx context.Context, messageNum *big.Int) (*ethbridge.DeliveredInboxMessage, error) {
	messages, err := s.rollup.LookupMessagesByNum(ctx, []*big.Int{messageNum})
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

func (s *Staker) selectValidChild(ctx context.Context, node ethbridge.NodeID) (*ethbridge.NodeInfo, error) {
	successors, err := s.validatorUtils.SuccessorNodes(ctx, node)
	if err != nil {
		return nil, err
	}
	nodes, err := s.rollup.LookupNodes(ctx, successors)
	if err != nil {
		return nil, err
	}

	if len(nodes) == 0 {
		return nil, nil
	}
	mach, err := s.lookup.GetMachine(nodes[0].Assertion.PrevState.TotalGasUsed)
	if err != nil {
		return nil, err
	}
	if mach.Hash() != nodes[0].Assertion.PrevState.MachineHash {
		return nil, errors.New("invalid machine state in start node")
	}

	for _, nd := range nodes {
		afterInboxHash, err := s.lookup.GetInboxAcc(nd.Assertion.AfterInboxCount())
		if err != nil {
			return nil, err
		}
		if nd.Assertion.AfterInboxHash != afterInboxHash {
			// Failed inbox consistency
			continue
		}
		messages, err := s.lookup.GetMessages(nd.Assertion.PrevState.InboxCount, nd.Assertion.InboxMessagesRead)
		if err != nil {
			return nil, err
		}
		if nd.Assertion.InboxDelta != calculateInboxDeltaAcc(messages) {
			// Failed inbox delta
			continue
		}
		localExecutionInfo, err := s.lookup.GetExecutionInfo(mach, nd.Assertion.ExecInfo.GasUsed)
		if err != nil {
			return nil, err
		}
		if !nd.Assertion.ExecInfo.Equals(localExecutionInfo) {
			// Failed execution
			continue
		}
		return nd, nil
	}
	return nil, nil
}

func calculateInboxDeltaAcc(messages []inbox.InboxMessage) common.Hash {
	acc := common.Hash{}
	for i := range messages {
		valHash := messages[len(messages)-1-i].AsValue().Hash()
		acc = hashing.SoliditySHA3(hashing.Bytes32(acc), hashing.Bytes32(valHash))
	}
	return acc
}
