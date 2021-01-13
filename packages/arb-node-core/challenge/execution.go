package challenge

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
)

type ExecutionImpl struct {
	initialGasUsed *big.Int
	finalGasUsed   *big.Int
}

func (e ExecutionImpl) GetCuts(lookup core.ValidatorLookup, offsets []*big.Int) ([]core.Cut, error) {
	panic("implement me")
}

func (e ExecutionImpl) FindFirstDivergence(lookup core.ValidatorLookup, offsets []*big.Int, cuts []core.Cut) (int, error) {
	panic("implement me")
}

func (e ExecutionImpl) Bisect(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	inconsistentSegment *core.ChallengeSegment,
	subCuts []core.Cut,
) (*types.Transaction, error) {
	return challenge.BisectExecution(
		ctx,
		prevBisection,
		segmentToChallenge,
		inconsistentSegment,
		subCuts,
	)
}

func (e ExecutionImpl) OneStepProof(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	lookup core.ValidatorLookup,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
) (*types.Transaction, error) {
	initalCursor, err := lookup.GetCursor(e.initialGasUsed)
	if err != nil {
		return nil, err
	}

	beforeCursor, err := lookup.MoveExecutionCursor(initalCursor, challengedSegment.Start, false)
	if err != nil {
		return nil, err
	}

	finalCursor, err := lookup.GetCursor(e.finalGasUsed)
	if err != nil {
		return nil, err
	}

	sendCount := new(big.Int).Sub(initalCursor.TotalSendCount(), beforeCursor.TotalSendCount())
	sendAcc, err := lookup.GetSendAcc(common.Hash{}, initalCursor.TotalSendCount(), sendCount)
	if err != nil {
		return nil, err
	}
	logCount := new(big.Int).Sub(beforeCursor.TotalLogCount(), beforeCursor.TotalLogCount())
	logAcc, err := lookup.GetLogAcc(common.Hash{}, initalCursor.TotalLogCount(), logCount)
	if err != nil {
		return nil, err
	}

	inboxRemaining := new(big.Int).Sub(finalCursor.NextInboxMessageIndex(), beforeCursor.NextInboxMessageIndex())
	inboxDelta, err := lookup.GetInboxDelta(beforeCursor.NextInboxMessageIndex(), inboxRemaining)

	beforeAssertion := &core.AssertionInfo{
		ExecutionInfo: &core.ExecutionInfo{
			SimpleExecutionInfo: &core.SimpleExecutionInfo{
				Before: core.NewExecutionState(initalCursor),
				After:  core.NewExecutionState(beforeCursor),
			},
			SendAcc: sendAcc,
			LogAcc:  logAcc,
		},
		InboxDelta: inboxDelta,
	}

	beforeMachine, err := lookup.GetMachine(beforeCursor)
	if err != nil {
		return nil, err
	}

	proofData, err := beforeMachine.MarshalForProof()
	if err != nil {
		return nil, err
	}

	bufferProofData, err := beforeMachine.MarshalBufferProof()
	if err != nil {
		return nil, err
	}

	return challenge.OneStepProveExecution(
		ctx,
		prevBisection,
		segmentToChallenge,
		beforeAssertion,
		proofData,
		bufferProofData,
	)
}
