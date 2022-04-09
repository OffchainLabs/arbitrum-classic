package staker

import (
	"context"
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbtransaction"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
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
	GasThreshold   *big.Int
	SendThreshold  *big.Int
	BlockThreshold *big.Int
}

func NewValidator(
	ctx context.Context,
	lookup core.ArbCoreLookup,
	client ethutils.EthClient,
	wallet *ethbridge.ValidatorWallet,
	fromBlock int64,
	validatorUtilsAddress common.Address,
	callOpts bind.CallOpts,
) (*Validator, error) {
	builder, err := ethbridge.NewBuilderBackend(wallet)
	if err != nil {
		return nil, err
	}
	rollup, err := ethbridge.NewRollup(wallet.RollupAddress().ToEthAddress(), fromBlock, client, builder, callOpts)
	_ = rollup
	if err != nil {
		return nil, err
	}
	delayedBridgeAddress, err := rollup.DelayedBridge(ctx)
	if err != nil {
		return nil, err
	}
	delayedBridge, err := ethbridge.NewDelayedBridgeWatcher(delayedBridgeAddress.ToEthAddress(), fromBlock, client)
	if err != nil {
		return nil, err
	}
	sequencerBridgeAddress, err := rollup.SequencerBridge(ctx)
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
		callOpts,
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
		GasThreshold:   big.NewInt(100_000_000_000),
		SendThreshold:  big.NewInt(5),
		BlockThreshold: big.NewInt(960),
	}, nil
}

// removeOldStakers removes the stakes of all currently staked validators except
// its own if dontRemoveSelf is true
func (v *Validator) removeOldStakers(ctx context.Context, dontRemoveSelf bool) (*arbtransaction.ArbTransaction, error) {
	stakersToEliminate, err := v.validatorUtils.RefundableStakers(ctx)
	if err != nil {
		return nil, err
	}
	walletAddr := v.wallet.Address()
	if dontRemoveSelf && walletAddr != nil {
		for i, staker := range stakersToEliminate {
			if staker.ToEthAddress() == *walletAddr {
				stakersToEliminate[i] = stakersToEliminate[len(stakersToEliminate)-1]
				stakersToEliminate = stakersToEliminate[:len(stakersToEliminate)-1]
				break
			}
		}
	}

	if len(stakersToEliminate) == 0 {
		return nil, nil
	}
	logger.Info().Int("count", len(stakersToEliminate)).Msg("Removing old stakers")
	return v.wallet.ReturnOldDeposits(ctx, stakersToEliminate)
}

func (v *Validator) resolveTimedOutChallenges(ctx context.Context) (*arbtransaction.ArbTransaction, error) {
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

func (v *Validator) resolveNextNode(ctx context.Context, info *ethbridge.StakerInfo, fromBlock int64) error {
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
		addr := v.wallet.Address()
		if info == nil || addr == nil || info.LatestStakedNode.Cmp(unresolvedNodeIndex) <= 0 {
			// We aren't an example of someone staked on a competitor
			return nil
		}
		logger.Info().Int("node", int(unresolvedNodeIndex.Int64())).Msg("Rejecting node")
		return v.rollup.RejectNextNode(ctx, *addr)
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

var maxAssertionSendCount *big.Int = big.NewInt(100) // From MAX_SEND_COUNT in RollupCore.sol

func (v *Validator) generateNodeAction(ctx context.Context, stakerInfo *OurStakerInfo, strategy configuration.ValidatorStrategy, fromBlock int64) (nodeAction, bool, error) {
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

	cursor := stakerInfo.latestExecutionCursor
	if cursor == nil || startState.TotalGasConsumed.Cmp(cursor.TotalGasConsumed()) < 0 {
		cursor, err = v.lookup.GetExecutionCursor(startState.TotalGasConsumed, true)
		if err != nil {
			return nil, false, err
		}
	} else {
		err = v.lookup.AdvanceExecutionCursor(cursor, new(big.Int).Sub(startState.TotalGasConsumed, cursor.TotalGasConsumed()), false, true)
		if err != nil {
			return nil, false, err
		}
	}
	cursorHash := cursor.MachineHash()
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

	// If there are no successor nodes, and there isn't much activity to process, don't do anything yet
	if len(successorNodes) == 0 {
		coreGasExecuted, err := v.lookup.GetLastMachineTotalGas()
		if err != nil {
			return nil, false, err
		}
		coreSendCount, err := v.lookup.GetSendCount()
		if err != nil {
			return nil, false, err
		}
		gasExecuted := new(big.Int).Sub(coreGasExecuted, startState.TotalGasConsumed)
		sendCount := new(big.Int).Sub(coreSendCount, startState.TotalSendCount)
		if sendCount.Cmp(v.SendThreshold) < 0 &&
			gasExecuted.Cmp(v.GasThreshold) < 0 &&
			timeSinceProposed.Cmp(v.BlockThreshold) < 0 {
			return nil, false, nil
		}
	}

	gasesUsed := make([]*big.Int, 0, len(successorNodes)+2)
	gasesUsed = append(gasesUsed, startState.TotalGasConsumed)
	for _, nd := range successorNodes {
		gasesUsed = append(gasesUsed, nd.Assertion.After.TotalGasConsumed)
	}

	arbGasSpeedLimitPerBlock, err := v.rollup.ArbGasSpeedLimitPerBlock(ctx)
	if err != nil {
		return nil, false, err
	}

	minimumGasToConsume := new(big.Int).Mul(timeSinceProposed, arbGasSpeedLimitPerBlock)
	maximumGasTarget := new(big.Int).Mul(minimumGasToConsume, big.NewInt(4))
	maximumGasTarget = maximumGasTarget.Add(maximumGasTarget, startState.TotalGasConsumed)
	maxTotalSendCount := new(big.Int).Add(startState.TotalSendCount, maxAssertionSendCount)

	if strategy.IsActive() || strategy == configuration.DefensiveStrategy {
		gasesUsed = append(gasesUsed, maximumGasTarget)
	}

	execTracker := core.NewExecutionTrackerWithInitialCursor(v.lookup, false, gasesUsed, cursor, false)

	var correctNode nodeAction
	wrongNodesExist := false
	if len(successorNodes) > 0 {
		logger.Info().Int("count", len(successorNodes)).Msg("examining existing potential successors")
	}
	for nodeI, nd := range successorNodes {
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
				logger.Info().Int("node", int((*big.Int)(nd.NodeNum).Int64())).Msg("found correct node")
				correctNode = existingNodeAction{
					number: nd.NodeNum,
					hash:   nd.NodeHash,
				}
				stakerInfo.latestExecutionCursor, err = execTracker.GetExecutionCursor(nd.AfterState().TotalGasConsumed, true)
				if err != nil {
					return nil, false, err
				}
				if nodeI != len(successorNodes)-1 && stakerInfo.latestExecutionCursor != nil {
					// We will need to use this execution tracker more, so we need to clone this cursor
					stakerInfo.latestExecutionCursor = stakerInfo.latestExecutionCursor.Clone()
				}
				continue
			} else {
				logger.Warn().Int("node", int((*big.Int)(nd.NodeNum).Int64())).Msg("found node with incorrect assertion")
			}
		} else {
			logger.Warn().Int("node", int((*big.Int)(nd.NodeNum).Int64())).Msg("found younger sibling to correct node")
		}
		// If we've hit this point, the node is "wrong"
		wrongNodesExist = true
	}

	if strategy == configuration.WatchtowerStrategy || correctNode != nil || (strategy != configuration.MakeNodesStrategy && !wrongNodesExist) {
		return correctNode, wrongNodesExist, nil
	}

	for i := len(gasesUsed) - 1; i >= 0; i-- {
		requestingGas := gasesUsed[i]
		if requestingGas.Cmp(maximumGasTarget) > 0 {
			continue
		}
		cursor, err = execTracker.GetExecutionCursor(requestingGas, true)
		if err != nil {
			return nil, false, err
		}
		if cursor.TotalSendCount().Cmp(maxTotalSendCount) > 0 {
			maximumGasTarget = new(big.Int).Sub(cursor.TotalGasConsumed(), big.NewInt(1))
			continue
		}
		// Binary search to find the maximum gas target that doesn't exceed the maxTotalSendCount
		for cursor.TotalGasConsumed().Cmp(maximumGasTarget) < 0 {
			advance := new(big.Int).Sub(maximumGasTarget, cursor.TotalGasConsumed())
			advance.Add(advance, big.NewInt(1))
			advance.Div(advance, big.NewInt(2))
			// At this point, advance = (maximumGasTarget - cursor.TotalGasConsumed() + 1) / 2
			// In other words, half the distance from cursor.TotalGasConsumed() to maximumGasTarget, rounding up.
			newCursor := cursor.Clone()
			v.lookup.AdvanceExecutionCursor(newCursor, advance, false, true)
			if cursor.TotalGasConsumed().Cmp(newCursor.TotalGasConsumed()) == 0 {
				// We've binary searched down to one instruction. This is good enough.
				break
			}
			if newCursor.TotalSendCount().Cmp(maxTotalSendCount) > 0 {
				maximumGasTarget = new(big.Int).Sub(newCursor.TotalGasConsumed(), big.NewInt(1))
			} else {
				cursor = newCursor
			}
		}
		break
	}

	stakerInfo.latestExecutionCursor = cursor
	execState, err := core.NewExecutionState(cursor)
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
		batch, err := v.sequencerInbox.LookupBatchContaining(ctx, v.lookup, new(big.Int).Sub(execState.TotalMessagesRead, big.NewInt(1)))
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
		return nil, errors.New("looked up starting node but found wrong hash")
	}
	return node.AfterState(), nil
}
