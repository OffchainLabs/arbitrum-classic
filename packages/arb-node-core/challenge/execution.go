package challenge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/pkg/errors"
	"math/big"
)

type ExecutionImpl struct {
	initialCursor core.ExecutionCursor
	inboxDelta    *inboxDelta
}

func (i *ExecutionImpl) SegmentTarget() int {
	return 400
}

func (e *ExecutionImpl) GetCuts(lookup core.ArbCoreLookup, offsets []*big.Int) ([]core.Cut, error) {
	execTracker := core.NewExecutionTracker(lookup, e.initialCursor, true, offsets)
	cuts := make([]core.Cut, 0, len(offsets))
	for _, offset := range offsets {
		executionInfo, err := execTracker.GetExecutionInfo(offset)
		if err != nil {
			return nil, err
		}

		cuts = append(cuts, core.ExecutionCut{
			GasUsed:      executionInfo.GasUsed(),
			InboxDelta:   e.inboxDelta.inboxDeltaAccs[executionInfo.InboxMessagesRead().Uint64()],
			MachineState: executionInfo.After.MachineHash,
			SendAcc:      executionInfo.SendAcc,
			SendCount:    executionInfo.SendCount(),
			LogAcc:       executionInfo.LogAcc,
			LogCount:     executionInfo.LogCount(),
		})
	}
	return cuts, nil
}

func (e *ExecutionImpl) FindFirstDivergence(lookup core.ArbCoreLookup, offsets []*big.Int, cuts []core.Cut) (int, error) {
	execTracker := core.NewExecutionTracker(lookup, e.initialCursor, true, offsets)
	for i, offset := range offsets {
		executionInfo, err := execTracker.GetExecutionInfo(offset)
		if err != nil {
			return 0, err
		}
		cut := core.ExecutionCut{
			GasUsed:      executionInfo.GasUsed(),
			InboxDelta:   e.inboxDelta.inboxDeltaAccs[executionInfo.InboxMessagesRead().Uint64()],
			MachineState: executionInfo.After.MachineHash,
			SendAcc:      executionInfo.SendAcc,
			SendCount:    executionInfo.SendCount(),
			LogAcc:       executionInfo.LogAcc,
			LogCount:     executionInfo.LogCount(),
		}
		if cut.CutHash() != cuts[i].CutHash() {
			if i == 0 && len(offsets) > 2 {
				return 0, errors.New("first cut was already wrong")
			}
			return i, nil
		}

		cuts = append(cuts)
	}
	panic("implement me")
}

func (e *ExecutionImpl) Bisect(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	inconsistentSegment *core.ChallengeSegment,
	subCuts []core.Cut,
) error {
	return challenge.BisectExecution(
		ctx,
		prevBisection,
		segmentToChallenge,
		inconsistentSegment,
		subCuts,
	)
}

func (e *ExecutionImpl) OneStepProof(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	lookup core.ArbCoreLookup,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
) error {
	tracker := core.NewExecutionTracker(lookup, e.initialCursor, true, []*big.Int{challengedSegment.Start})
	execInfo, err := tracker.GetExecutionInfo(challengedSegment.Start)
	if err != nil {
		return err
	}

	beforeMachine, err := tracker.GetMachine(challengedSegment.Start)
	if err != nil {
		return err
	}

	proofData, bufferProofData, err := beforeMachine.MarshalForProof()
	if err != nil {
		return err
	}

	// TODO: Get this from proofData
	opcode := uint8(0)

	return challenge.OneStepProveExecution(
		ctx,
		prevBisection,
		segmentToChallenge,
		execInfo,
		e.inboxDelta.inboxDeltaAccs[execInfo.InboxMessagesRead().Uint64()],
		proofData,
		bufferProofData,
		opcode,
	)
}
