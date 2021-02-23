package challenge

import (
	"context"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/pkg/errors"
)

type ExecutionImpl struct {
	initialCursor core.ExecutionCursor
}

func (e *ExecutionImpl) SegmentTarget() int {
	return 400
}

var unreachableCut core.SimpleCut = core.NewSimpleCut([32]byte{})

// Note: it's vital that execTracker's first gas offset was the beginning of the original assertion.
// All fields except machine state and total messages read are relative to the beginning of the assertion.
func getCut(execTracker *core.ExecutionTracker, maxTotalMessagesRead *big.Int, gasTarget *big.Int) (core.Cut, *big.Int, error) {
	executionInfo, steps, err := execTracker.GetExecutionInfo(gasTarget)
	if err != nil {
		return nil, nil, err
	}
	if executionInfo.After.TotalMessagesRead.Cmp(maxTotalMessagesRead) > 0 || executionInfo.After.TotalGasConsumed.Cmp(gasTarget) < 0 {
		// Execution read more messages than provided so assertion should have
		// stopped short
		return unreachableCut, steps, nil
	}
	return core.ExecutionCut{
		GasUsed:           executionInfo.GasUsed(),
		TotalMessagesRead: executionInfo.After.TotalMessagesRead,
		MachineState:      executionInfo.After.MachineHash,
		SendAcc:           executionInfo.SendAcc,
		SendCount:         executionInfo.SendCount(),
		LogAcc:            executionInfo.LogAcc,
		LogCount:          executionInfo.LogCount(),
	}, steps, nil
}

func (e *ExecutionImpl) GetCuts(lookup core.ArbCoreLookup, assertion *core.Assertion, offsets []*big.Int) ([]core.Cut, error) {
	fullOffsets := append(offsets, assertion.Before.TotalGasConsumed)
	execTracker := core.NewExecutionTracker(lookup, e.initialCursor, true, fullOffsets)
	cuts := make([]core.Cut, 0, len(offsets))
	for i, offset := range offsets {
		cut, _, err := getCut(execTracker, assertion.After.TotalMessagesRead, offset)
		if err != nil {
			return nil, err
		}
		if i == 0 {
			_, ok := cut.(core.ExecutionCut)
			if !ok {
				return nil, errors.New("first cut is unreachable")
			}
		}

		cuts = append(cuts, cut)
	}
	return cuts, nil
}

type DivergenceInfo struct {
	DifferentIndex   int
	SegmentSteps     *big.Int
	EndIsUnreachable bool
}

func (e *ExecutionImpl) FindFirstDivergence(lookup core.ArbCoreLookup, assertion *core.Assertion, offsets []*big.Int, cuts []core.Cut) (DivergenceInfo, error) {
	errRes := DivergenceInfo{
		DifferentIndex:   0,
		SegmentSteps:     big.NewInt(0),
		EndIsUnreachable: false,
	}
	fullOffsets := append(offsets, assertion.Before.TotalGasConsumed)
	execTracker := core.NewExecutionTracker(lookup, e.initialCursor, true, fullOffsets)
	lastSteps := big.NewInt(0)
	for i, offset := range offsets {
		localCut, newSteps, err := getCut(execTracker, assertion.After.TotalMessagesRead, offset)
		if err != nil {
			return errRes, err
		}
		if localCut.CutHash() != cuts[i].CutHash() {
			return DivergenceInfo{
				DifferentIndex:   i,
				SegmentSteps:     new(big.Int).Sub(newSteps, lastSteps),
				EndIsUnreachable: localCut == unreachableCut,
			}, nil
		}
		lastSteps = newSteps
	}
	return errRes, errors.New("no divergence found in cuts")
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

func (e *ExecutionImpl) getSegmentStartInfo(lookup core.ArbCoreLookup, assertion *core.Assertion, segment *core.ChallengeSegment) (core.ExecutionCut, machine.Machine, error) {
	execTracker := core.NewExecutionTracker(lookup, e.initialCursor, true, []*big.Int{assertion.Before.TotalGasConsumed, segment.Start})
	cut, _, err := getCut(execTracker, assertion.After.TotalMessagesRead, segment.Start)
	if err != nil {
		return core.ExecutionCut{}, nil, err
	}
	execCut, ok := cut.(core.ExecutionCut)
	if !ok {
		return core.ExecutionCut{}, nil, errors.New("attempted to one step prove blocked machine")
	}

	beforeMachine, err := execTracker.GetMachine(segment.Start)
	if err != nil {
		return core.ExecutionCut{}, nil, err
	}

	return execCut, beforeMachine, nil
}

func (e *ExecutionImpl) OneStepProof(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	lookup core.ArbCoreLookup,
	assertion *core.Assertion,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
) error {
	previousCut, previousMachine, err := e.getSegmentStartInfo(lookup, assertion, challengedSegment)
	if err != nil {
		return err
	}

	proofData, bufferProofData, err := previousMachine.MarshalForProof()
	if err != nil {
		return err
	}

	opcode := proofData[0]

	return challenge.OneStepProveExecution(
		ctx,
		prevBisection,
		segmentToChallenge,
		challengedSegment,
		previousCut,
		proofData,
		bufferProofData,
		opcode,
	)
}

func (e *ExecutionImpl) ProveContinuedExecution(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	lookup core.ArbCoreLookup,
	assertion *core.Assertion,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
) error {
	previousCut, _, err := e.getSegmentStartInfo(lookup, assertion, challengedSegment)
	if err != nil {
		return err
	}

	return challenge.ProveContinuedExecution(
		ctx,
		prevBisection,
		segmentToChallenge,
		challengedSegment,
		previousCut,
	)
}
