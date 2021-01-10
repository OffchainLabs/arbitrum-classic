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
	GetExecutionInfo(startMachine machine.Machine, gas *big.Int, maxMessages *big.Int) (*ethbridge.ExecutionInfo, error)
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

func (v *Validator) generateNodeAction(ctx context.Context, base ethbridge.NodeID) (nodeActionInfo, error) {
	lastNodeCreated, err := v.rollup.LatestNodeCreated(ctx)
	if err != nil {
		return nil, err
	}

	validChild, err := v.selectValidChild(ctx, base)
	if err != nil {
		return nil, err
	}
	if validChild != nil {
		blockId, err := GetBlockID(ctx, v.client, validChild.Assertion.PrevState.ProposedBlock)
		if err != nil {
			return nil, err
		}
		return &nodeMovementInfo{
			block:   blockId,
			nodeNum: validChild.NodeNum,
		}, nil
	}

	currentNode, err := v.lookupNode(ctx, base)
	if err != nil {
		return nil, err
	}
	minAssertionPeriod, err := v.rollup.MinimumAssertionPeriod(ctx)
	if err != nil {
		return nil, err
	}
	arbGasSpeedLimitPerBlock, err := v.rollup.ArbGasSpeedLimitPerBlock(ctx)
	if err != nil {
		return nil, err
	}

	currentBlockId, err := GetBlockID(ctx, v.client, nil)
	if err != nil {
		return nil, err
	}
	timeSinceProposed := new(big.Int).Sub(currentBlockId.Height.AsInt(), currentNode.BlockProposed.Height.AsInt())
	if timeSinceProposed.Cmp(minAssertionPeriod) < 0 {
		// Too soon to assert
		return nil, nil
	}

	mach, err := v.lookup.GetMachine(currentNode.Assertion.AfterTotalGasUsed())
	if err != nil {
		return nil, err
	}
	if mach.Hash() != currentNode.Assertion.ExecInfo.AfterMachineHash {
		return nil, errors.New("local machine doesn't match chain")
	}

	minimumGasToConsume := new(big.Int).Mul(timeSinceProposed, arbGasSpeedLimitPerBlock)
	maximumGasToConsume := new(big.Int).Mul(minimumGasToConsume, big.NewInt(4))

	assertionInfo, err := v.lookup.GetExecutionInfoInRange(mach, minimumGasToConsume, maximumGasToConsume)
	if err != nil {
		return nil, err
	}

	assertion := &ethbridge.Assertion{
		PrevState:     currentNode.AfterState(),
		AssertionInfo: assertionInfo,
	}
	msg, err := v.lookupMessageByNum(ctx, assertion.AfterInboxCount())
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

func (v *Validator) lookupNode(ctx context.Context, node ethbridge.NodeID) (*ethbridge.NodeInfo, error) {
	currentNodes, err := v.rollup.LookupNodes(ctx, []*big.Int{node})
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

func (v *Validator) lookupMessageByNum(ctx context.Context, messageNum *big.Int) (*ethbridge.DeliveredInboxMessage, error) {
	messages, err := v.rollup.LookupMessagesByNum(ctx, []*big.Int{messageNum})
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

func (v *Validator) selectValidChild(ctx context.Context, node ethbridge.NodeID) (*ethbridge.NodeInfo, error) {
	successors, err := v.validatorUtils.SuccessorNodes(ctx, node)
	if err != nil {
		return nil, err
	}
	nodes, err := v.rollup.LookupNodes(ctx, successors)
	if err != nil {
		return nil, err
	}

	if len(nodes) == 0 {
		return nil, nil
	}
	mach, err := v.lookup.GetMachine(nodes[0].Assertion.PrevState.TotalGasUsed)
	if err != nil {
		return nil, err
	}
	if mach.Hash() != nodes[0].Assertion.PrevState.MachineHash {
		return nil, errors.New("invalid machine state in start node")
	}

	for _, nd := range nodes {
		chalType, err := v.judgeNode(nd, mach)
		if err != nil {
			return nil, err
		}
		if chalType == ethbridge.NO_CHALLENGE {
			return nd, nil
		}
	}
	return nil, nil
}

func (v *Validator) judgeNode(nd *ethbridge.NodeInfo, mach machine.Machine) (ethbridge.ChallengeKind, error) {
	afterInboxHash, err := v.lookup.GetInboxAcc(nd.Assertion.AfterInboxCount())
	if err != nil {
		return 0, err
	}
	if nd.Assertion.AfterInboxHash != afterInboxHash {
		// Failed inbox consistency
		return ethbridge.INBOX_CONSISTENCY, nil
	}
	messages, err := v.lookup.GetMessages(nd.Assertion.PrevState.InboxCount, nd.Assertion.ExecInfo.InboxMessagesRead)
	if err != nil {
		return 0, err
	}
	if nd.Assertion.InboxDelta != calculateInboxDeltaAcc(messages) {
		// Failed inbox delta
		return ethbridge.INBOX_DELTA, nil
	}
	if mach == nil {
		mach, err = v.lookup.GetMachine(nd.Assertion.PrevState.TotalGasUsed)
		if err != nil {
			return 0, err
		}
	}
	localExecutionInfo, err := v.lookup.GetExecutionInfo(mach, nd.Assertion.ExecInfo.GasUsed, nd.Assertion.ExecInfo.InboxMessagesRead)
	if err != nil {
		return 0, err
	}

	if localExecutionInfo.GasUsed.Cmp(nd.Assertion.ExecInfo.GasUsed) < 0 {
		return ethbridge.STOPPED_SHORT, nil
	}

	if !nd.Assertion.ExecInfo.Equals(localExecutionInfo) {
		return ethbridge.EXECUTION, nil
	}
	return ethbridge.NO_CHALLENGE, nil
}

type Staker struct {
	*Validator
	address      common.Address
	makeNewNodes bool
}

func (s *Staker) Act(ctx context.Context) error {
	_, err := s.resolveNextNode(ctx)
	if err != nil {
		return err
	}
	info, err := s.rollup.StakerInfo(ctx, s.address)
	if err != nil {
		return err
	}
	if info != nil {
		_, err := s.advanceStake(ctx, info)
		if err != nil {
			return err
		}
	} else {
		_, err := s.placeStake(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Staker) handleConflict(ctx context.Context, info *ethbridge.StakerInfo) (*types.Transaction, error) {
	if info.CurrentChallenge == nil {
		return nil, nil
	}
	challenge, err := ethbridge.NewChallengeWatcher(info.CurrentChallenge.ToEthAddress(), s.client)
	if err != nil {
		return nil, err
	}
	responder, err := challenge.CurrentResponder(ctx)
	if err != nil {
		return nil, err
	}
	if responder != s.address {
		// Not our turn
		return nil, nil
	}
	kind, err := challenge.Kind(ctx)
	if err != nil {
		return nil, err
	}

	challengedNodeNum, err := challenge.ChallengedNodeNum(ctx)
	if err != nil {
		return nil, err
	}

	nodeInfo, err := s.lookupNode(ctx, challengedNodeNum)
	if err != nil {
		return nil, err
	}

	switch kind {
	case ethbridge.UNINITIALIZED:
		judgment, err := s.judgeNode(nodeInfo, nil)
		if err != nil {
			return nil, err
		}
		switch judgment {
		case ethbridge.INBOX_CONSISTENCY:
			return s.handleInboxConsistencyChallenge()
		case ethbridge.INBOX_DELTA:
			return s.handleInboxDeltaChallenge()
		case ethbridge.EXECUTION:
			return s.handleExecutionChallenge()
		case ethbridge.STOPPED_SHORT:
			return s.handleStoppedShortChallenge()
		default:
			return nil, errors.New("can't handle challenge")
		}
	case ethbridge.INBOX_CONSISTENCY:
		return s.handleInboxConsistencyChallenge()
	case ethbridge.INBOX_DELTA:
		return s.handleInboxDeltaChallenge()
	case ethbridge.EXECUTION:
		return s.handleExecutionChallenge()
	case ethbridge.STOPPED_SHORT:
		return s.handleStoppedShortChallenge()
	default:
		return nil, errors.New("can't handle challenge")
	}
}

func (s *Staker) handleInboxConsistencyChallenge() (*types.Transaction, error) {
	return nil, nil
}

func (s *Staker) handleInboxDeltaChallenge() (*types.Transaction, error) {
	return nil, nil
}

func (s *Staker) handleExecutionChallenge() (*types.Transaction, error) {
	return nil, nil
}

func (s *Staker) handleStoppedShortChallenge() (*types.Transaction, error) {
	return nil, nil
}

func (s *Staker) advanceStake(ctx context.Context, info *ethbridge.StakerInfo) (*types.Transaction, error) {
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
		if !s.makeNewNodes {
			return nil, nil
		}
		return s.rollup.AddStakeOnNewNode(ctx, action.block, action.newNodeID, action.assertion)
	case nodeMovementInfo:
		return s.rollup.AddStakeOnExistingNode(ctx, action.block, action.nodeNum)
	default:
		panic("invalid type")
	}
}

func (s *Staker) placeStake(ctx context.Context) (*types.Transaction, error) {
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
		if !s.makeNewNodes {
			return nil, nil
		}
		return s.rollup.NewStakeOnNewNode(ctx, action.block, action.newNodeID, latestConfirmedNode, action.assertion)
	case nodeMovementInfo:
		return s.rollup.NewStakeOnExistingNode(ctx, action.block, action.nodeNum)
	default:
		panic("invalid type")
	}
}

func (s *Staker) createConflict(ctx context.Context) (*types.Transaction, error) {
	info, err := s.rollup.StakerInfo(ctx, s.address)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, errors.New("not staked")
	}
	if info.CurrentChallenge != nil {
		return nil, nil
	}

	stakers, err := s.rollup.GetStakers(ctx)
	if err != nil {
		return nil, err
	}
	for _, staker := range stakers {
		conflictType, node1, node2, err := s.validatorUtils.FindStakerConflict(ctx, s.address, staker)
		if err != nil {
			return nil, err
		}
		if conflictType != ethbridge.CONFLICT_TYPE_FOUND {
			continue
		}
		staker1 := s.address
		staker2 := staker
		if node2.Cmp(node1) < 0 {
			staker1, staker2 = staker2, staker1
			node1, node2 = node1, node1
		}

		nodeInfo, err := s.lookupNode(ctx, node1)
		if err != nil {
			return nil, err
		}
		maxInboxHash, err := s.lookup.GetInboxAcc(nodeInfo.InboxMaxCount)
		if err != nil {
			return nil, err
		}
		return s.rollup.CreateChallenge(
			ctx,
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

func calculateInboxDeltaAcc(messages []inbox.InboxMessage) common.Hash {
	acc := common.Hash{}
	for i := range messages {
		valHash := messages[len(messages)-1-i].AsValue().Hash()
		acc = hashing.SoliditySHA3(hashing.Bytes32(acc), hashing.Bytes32(valHash))
	}
	return acc
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
