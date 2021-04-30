package staker

import (
	"context"
	"encoding/hex"
	"math/big"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
)

type Validator struct {
	rollup         *ethbridge.Rollup
	delayedBridge  *ethbridge.DelayedBridgeWatcher
	sequencerInbox *ethbridge.SequencerInboxWatcher
	validatorUtils *ethbridge.ValidatorUtils
	client         ethutils.EthClient
	lookup         core.ArbCoreLookup
	builder        *ethbridge.BuilderBackend
	wallet         *ethbridge.ValidatorWallet
}

func NewValidator(
	ctx context.Context,
	lookup core.ArbCoreLookup,
	client ethutils.EthClient,
	wallet *ethbridge.ValidatorWallet,
	validatorUtilsAddress common.Address,
) (*Validator, error) {
	builder, err := ethbridge.NewBuilderBackend(wallet)
	if err != nil {
		return nil, err
	}
	rollup, err := ethbridge.NewRollup(wallet.RollupAddress().ToEthAddress(), client, builder)
	if err != nil {
		return nil, err
	}
	delayedBridgeAddress, err := rollup.DelayedBridge(context.Background())
	if err != nil {
		return nil, err
	}
	delayedBridge, err := ethbridge.NewDelayedBridgeWatcher(delayedBridgeAddress.ToEthAddress(), client)
	if err != nil {
		return nil, err
	}
	sequencerBridgeAddress, err := rollup.SequencerBridge(context.Background())
	if err != nil {
		return nil, err
	}
	sequencerInbox, err := ethbridge.NewSequencerInboxWatcher(sequencerBridgeAddress.ToEthAddress(), client)
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
		delayedBridge:  delayedBridge,
		sequencerInbox: sequencerInbox,
		validatorUtils: validatorUtils,
		client:         client,
		lookup:         lookup,
		builder:        builder,
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
	logger.Info().Int("count", len(stakersToEliminate)).Msg("Removing old stakers")
	return v.wallet.ReturnOldDeposits(ctx, stakersToEliminate)
}

func (v *Validator) resolveTimedOutChallenges(ctx context.Context) (*types.Transaction, error) {
	challengesToEliminate, err := v.validatorUtils.TimedOutChallenges(ctx, 10)
	if err != nil {
		return nil, err
	}
	if len(challengesToEliminate) == 0 {
		return nil, nil
	}
	logger.Info().Int("count", len(challengesToEliminate)).Msg("Timing out challenges")
	return v.wallet.TimeoutChallenges(ctx, challengesToEliminate)
}

func (v *Validator) resolveNextNode(ctx context.Context, info *ethbridge.StakerInfo) error {
	confirmType, err := v.validatorUtils.CheckDecidableNextNode(ctx)
	if err != nil {
		return err
	}
	unresolvedNodeIndex, err := v.rollup.FirstUnresolvedNode(ctx)
	if err != nil {
		return err
	}
	switch confirmType {
	case ethbridge.CONFIRM_TYPE_INVALID:
		if info == nil || info.LatestStakedNode.Cmp(unresolvedNodeIndex) <= 0 {
			// We aren't an example of someone staked on a competitor
			return nil
		}
		logger.Info().Int("node", int(unresolvedNodeIndex.Int64())).Msg("Rejecting node")
		return v.rollup.RejectNextNode(ctx, v.wallet.Address())
	case ethbridge.CONFIRM_TYPE_VALID:
		nodeInfo, err := v.rollup.RollupWatcher.LookupNode(ctx, unresolvedNodeIndex)
		if err != nil {
			return err
		}
		sendCount := new(big.Int).Sub(nodeInfo.Assertion.After.TotalSendCount, nodeInfo.Assertion.Before.TotalSendCount)
		sends, err := v.lookup.GetSends(nodeInfo.Assertion.Before.TotalSendCount, sendCount)
		if err != nil {
			return errors.Wrap(err, "catching up to chain")
		}
		logger.Info().Int("node", int(unresolvedNodeIndex.Int64())).Msg("Confirming node")
		return v.rollup.ConfirmNextNode(ctx, nodeInfo.Assertion, sends)
	default:
		return nil
	}
}

func (v *Validator) isRequiredStakeElevated(ctx context.Context) (bool, error) {
	requiredStake, err := v.rollup.CurrentRequiredStake(ctx)
	if err != nil {
		return false, err
	}
	baseStake, err := v.rollup.BaseStake(ctx)
	if err != nil {
		return false, err
	}
	return requiredStake.Cmp(baseStake) > 0, nil
}

type createNodeAction struct {
	assertion           *core.Assertion
	prevProposedBlock   *big.Int
	prevInboxMaxCount   *big.Int
	hash                [32]byte
	sequencerBatchProof []byte
}

type existingNodeAction struct {
	number core.NodeID
	hash   [32]byte
}

type nodeAction interface{}

type OurStakerInfo struct {
	LatestStakedNode      *big.Int
	LatestStakedNodeHash  [32]byte
	CanProgress           bool
	latestExecutionCursor core.ExecutionCursor
	*ethbridge.StakerInfo
}

func (v *Validator) generateNodeAction(ctx context.Context, stakerInfo *OurStakerInfo, strategy Strategy) (nodeAction, bool, error) {
	startState, err := lookupNodeStartState(ctx, v.rollup.RollupWatcher, stakerInfo.LatestStakedNode, stakerInfo.LatestStakedNodeHash)
	if err != nil {
		return nil, false, err
	}

	coreMessageCount := v.lookup.MachineMessagesRead()
	if coreMessageCount.Cmp(startState.TotalMessagesRead) < 0 {
		logger.Info().
			Str("localcount", coreMessageCount.String()).
			Str("target", startState.TotalMessagesRead.String()).
			Msg("catching up to chain")
		return nil, false, nil
	}
	cursor := stakerInfo.latestExecutionCursor
	if cursor == nil || startState.TotalGasConsumed.Cmp(cursor.TotalGasConsumed()) < 0 {
		cursor, err = v.lookup.GetExecutionCursor(startState.TotalGasConsumed)
		if err != nil {
			return nil, false, err
		}
	} else {
		err = v.lookup.AdvanceExecutionCursor(cursor, new(big.Int).Sub(startState.TotalGasConsumed, cursor.TotalGasConsumed()), false)
		if err != nil {
			return nil, false, err
		}
	}
	cursorHash, err := cursor.MachineHash()
	if err != nil {
		return nil, false, err
	}
	if cursorHash != startState.MachineHash {
		return nil, false, errors.Errorf("local machine doesn't match chain %v %v", cursor.TotalGasConsumed(), startState.TotalGasConsumed)
	}

	// Not necessarily successors
	successorNodes, err := v.rollup.LookupNodeChildren(ctx, stakerInfo.LatestStakedNodeHash, startState.ProposedBlock)
	if err != nil {
		return nil, false, err
	}

	gasesUsed := make([]*big.Int, 0, len(successorNodes)+1)
	for _, nd := range successorNodes {
		gasesUsed = append(gasesUsed, nd.Assertion.After.TotalGasConsumed)
	}

	currentBlock, err := getBlockID(ctx, v.client, nil)
	if err != nil {
		return nil, false, err
	}

	minAssertionPeriod, err := v.rollup.MinimumAssertionPeriod(ctx)
	if err != nil {
		return nil, false, err
	}

	timeSinceProposed := new(big.Int).Sub(currentBlock.Height.AsInt(), startState.ProposedBlock)
	if timeSinceProposed.Cmp(minAssertionPeriod) < 0 {
		// Too soon to assert
		return nil, false, nil
	}

	arbGasSpeedLimitPerBlock, err := v.rollup.ArbGasSpeedLimitPerBlock(ctx)
	if err != nil {
		return nil, false, err
	}

	minimumGasToConsume := new(big.Int).Mul(timeSinceProposed, arbGasSpeedLimitPerBlock)
	maximumGasTarget := new(big.Int).Mul(minimumGasToConsume, big.NewInt(4))
	maximumGasTarget = maximumGasTarget.Add(maximumGasTarget, startState.TotalGasConsumed)

	if strategy > WatchtowerStrategy {
		gasesUsed = append(gasesUsed, maximumGasTarget)
	}

	execTracker := core.NewExecutionTrackerWithInitialCursor(v.lookup, false, gasesUsed, cursor)

	var correctNode nodeAction
	wrongNodesExist := false
	if len(successorNodes) > 0 {
		logger.Info().Int("count", len(successorNodes)).Msg("Examining existing potential successors")
	}
	for _, nd := range successorNodes {
		if correctNode != nil && wrongNodesExist {
			// We've found everything we could hope to find
			break
		}
		if correctNode == nil {
			var batchItemEndAcc common.Hash
			if nd.Assertion.After.TotalMessagesRead.Cmp(nd.AfterInboxBatchEndCount) == 0 {
				batchItemEndAcc = nd.AfterInboxBatchAcc
			} else if nd.Assertion.After.TotalMessagesRead.Cmp(big.NewInt(0)) > 0 {
				var haveBatchEndAcc common.Hash
				index1 := new(big.Int).Sub(nd.Assertion.After.TotalMessagesRead, big.NewInt(1))
				index2 := new(big.Int).Sub(nd.AfterInboxBatchEndCount, big.NewInt(1))
				batchItemEndAcc, haveBatchEndAcc, err = v.lookup.GetInboxAccPair(index1, index2)
				if err != nil {
					return nil, false, err
				}
				if haveBatchEndAcc != nd.AfterInboxBatchAcc {
					return nil, false, errors.New("inbox reorg detected by batch end acc mismatch")
				}
			}
			valid, err := core.IsAssertionValid(nd.Assertion, execTracker, batchItemEndAcc)
			if err != nil {
				return nil, false, err
			}
			if valid {
				logger.Info().Int("node", int((*big.Int)(nd.NodeNum).Int64())).Msg("Found correct node")
				correctNode = existingNodeAction{
					number: nd.NodeNum,
					hash:   nd.NodeHash,
				}
				stakerInfo.latestExecutionCursor, err = execTracker.GetExecutionCursor(nd.AfterState().TotalGasConsumed)
				if err != nil {
					return nil, false, err
				}
				continue
			} else {
				logger.Warn().Int("node", int((*big.Int)(nd.NodeNum).Int64())).Msg("Found node with incorrect assertion")
			}
		} else {
			logger.Warn().Int("node", int((*big.Int)(nd.NodeNum).Int64())).Msg("Found younger sibling to correct node")
		}
		// If we've hit this point, the node is "wrong"
		wrongNodesExist = true
	}

	if strategy == WatchtowerStrategy || correctNode != nil || (strategy < MakeNodesStrategy && !wrongNodesExist) {
		return correctNode, wrongNodesExist, nil
	}

	execState, _, err := execTracker.GetExecutionState(maximumGasTarget)
	if err != nil {
		return nil, false, err
	}
	stakerInfo.latestExecutionCursor, err = execTracker.GetExecutionCursor(maximumGasTarget)
	if err != nil {
		return nil, false, err
	}

	if new(big.Int).Sub(execState.TotalGasConsumed, startState.TotalGasConsumed).Cmp(minimumGasToConsume) < 0 && execState.TotalMessagesRead.Cmp(startState.InboxMaxCount) < 0 {
		// Couldn't execute far enough
		return nil, wrongNodesExist, nil
	}

	inboxAcc := execState.InboxAcc
	hasSiblingByte := [1]byte{0}
	lastNum := stakerInfo.LatestStakedNode
	lastHash := stakerInfo.LatestStakedNodeHash
	if len(successorNodes) > 0 {
		lastSuccessor := successorNodes[len(successorNodes)-1]
		lastNum = lastSuccessor.NodeNum
		lastHash = lastSuccessor.NodeHash
		hasSiblingByte[0] = 1
	}
	assertion := &core.Assertion{
		Before: startState.ExecutionState,
		After:  execState,
	}
	executionHash := assertion.ExecutionHash()
	newNodeHash := hashing.SoliditySHA3(hasSiblingByte[:], lastHash[:], executionHash[:], inboxAcc[:])

	var seqBatchProof []byte
	if execState.TotalMessagesRead.Cmp(big.NewInt(0)) > 0 {
		batch, err := challenge.LookupBatchContaining(ctx, v.lookup, v.sequencerInbox, new(big.Int).Sub(execState.TotalMessagesRead, big.NewInt(1)))
		if err != nil {
			return nil, false, err
		}
		if batch == nil {
			return nil, false, errors.New("Failed to lookup batch containing message")
		}
		seqBatchProof = append(seqBatchProof, math.U256Bytes(batch.GetBatchIndex())...)
		proofPart, err := v.generateBatchEndProof(batch.GetBeforeCount())
		if err != nil {
			return nil, false, err
		}
		seqBatchProof = append(seqBatchProof, proofPart...)
		proofPart, err = v.generateBatchEndProof(batch.GetAfterCount())
		if err != nil {
			return nil, false, err
		}
		seqBatchProof = append(seqBatchProof, proofPart...)
	}

	action := createNodeAction{
		assertion:           assertion,
		hash:                newNodeHash,
		prevProposedBlock:   startState.ProposedBlock,
		prevInboxMaxCount:   startState.InboxMaxCount,
		sequencerBatchProof: seqBatchProof,
	}
	logger.Info().Str("hash", hex.EncodeToString(newNodeHash[:])).Int("lastNode", int(lastNum.Int64())).Int("parentNode", int(stakerInfo.LatestStakedNode.Int64())).Msg("Creating node")
	return action, wrongNodesExist, nil
}

func (v *Validator) generateBatchEndProof(count *big.Int) ([]byte, error) {
	if count.Cmp(big.NewInt(0)) == 0 {
		return []byte{}, nil
	}
	var beforeAcc common.Hash
	var err error
	if count.Cmp(big.NewInt(2)) >= 0 {
		beforeAcc, err = v.lookup.GetInboxAcc(new(big.Int).Sub(count, big.NewInt(2)))
		if err != nil {
			return nil, err
		}
	}
	seqNum := new(big.Int).Sub(count, big.NewInt(1))
	message, err := core.GetSingleMessage(v.lookup, seqNum)
	if err != nil {
		return nil, err
	}
	var proof []byte
	proof = append(proof, beforeAcc.Bytes()...)
	proof = append(proof, math.U256Bytes(seqNum)...)
	prefixHash := hashing.SoliditySHA3(
		hashing.Address(message.Sender),
		hashing.Uint256(message.ChainTime.BlockNum.AsInt()),
		hashing.Uint256(message.ChainTime.Timestamp),
	)
	proof = append(proof, prefixHash.Bytes()...)
	proof = append(proof, hashing.SoliditySHA3(message.Data).Bytes()...)
	return proof, nil
}

func (v *Validator) GetInitialMachineHash(ctx context.Context) ([32]byte, error) {
	creationEvent, err := v.rollup.LookupCreation(ctx)
	if err != nil {
		return [32]byte{}, err
	}
	return creationEvent.MachineHash, nil
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

func lookupNodeStartState(ctx context.Context, rollup *ethbridge.RollupWatcher, nodeNum *big.Int, nodeHash [32]byte) (*core.NodeState, error) {
	if nodeNum.Cmp(big.NewInt(0)) == 0 {
		creationEvent, err := rollup.LookupCreation(ctx)
		if err != nil {
			return nil, err
		}
		return &core.NodeState{
			ProposedBlock: new(big.Int).SetUint64(creationEvent.Raw.BlockNumber),
			InboxMaxCount: big.NewInt(1),
			ExecutionState: &core.ExecutionState{
				TotalGasConsumed:  big.NewInt(0),
				MachineHash:       creationEvent.MachineHash,
				TotalMessagesRead: big.NewInt(0),
				TotalSendCount:    big.NewInt(0),
				TotalLogCount:     big.NewInt(0),
			},
		}, nil
	}
	node, err := rollup.LookupNode(ctx, nodeNum)
	if err != nil {
		return nil, err
	}
	if node.NodeHash != nodeHash {
		return nil, errors.New("Looked up starting node but found wrong hash")
	}
	return node.AfterState(), nil
}
