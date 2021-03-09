package core

import (
	"math/big"
	"sort"

	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

type BigIntList []*big.Int

func (l BigIntList) Len() int {
	return len(l)
}

func (l BigIntList) Less(i, j int) bool {
	return l[i].Cmp(l[j]) < 0
}

func (l BigIntList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type ExecutionTracker struct {
	lookup    ArbCoreLookup
	goOverGas bool

	sortedStopPoints []*big.Int
	stopPointIndex   map[string]int

	cursors []ExecutionCursor
}

func NewExecutionTracker(lookup ArbCoreLookup, goOverGas bool, stopPointsArg []*big.Int) *ExecutionTracker {
	sort.Sort(BigIntList(stopPointsArg))
	// Deduplicate stop points
	stopPoints := make([]*big.Int, 0, len(stopPointsArg))
	var lastStopPoint *big.Int = nil
	for _, stopPoint := range stopPointsArg {
		if lastStopPoint != nil && lastStopPoint.Cmp(stopPoint) == 0 {
			continue
		}
		stopPoints = append(stopPoints, stopPoint)
		lastStopPoint = stopPoint
	}
	cursors := make([]ExecutionCursor, 0, len(stopPoints))

	stopPointIndex := make(map[string]int)
	for i, stopPoint := range stopPoints {
		stopPointIndex[string(stopPoint.Bytes())] = i
	}
	return &ExecutionTracker{
		lookup:           lookup,
		goOverGas:        goOverGas,
		sortedStopPoints: stopPoints,
		stopPointIndex:   stopPointIndex,
		cursors:          cursors,
	}
}

func (e *ExecutionTracker) fillInCursors(max int) error {
	for i := len(e.cursors); i <= max; i++ {
		var nextCursor ExecutionCursor
		var err error
		if i > 0 {
			nextCursor = e.cursors[i-1].Clone()
		} else {
			nextCursor, err = e.lookup.GetExecutionCursor(e.sortedStopPoints[i])
			// Note: we still might need to advance since we can't set goOverGas here
		}
		if err != nil {
			return err
		}
		nextStopPoint := e.sortedStopPoints[i]
		if nextStopPoint.Cmp(nextCursor.TotalGasConsumed()) > 0 {
			gasToExecute := new(big.Int).Sub(nextStopPoint, nextCursor.TotalGasConsumed())
			err = e.lookup.AdvanceExecutionCursor(nextCursor, gasToExecute, e.goOverGas)
		}
		if err != nil {
			return err
		}
		e.cursors = append(e.cursors, nextCursor)
	}
	return nil
}

func (e *ExecutionTracker) GetExecutionState(gasUsed *big.Int) (*ExecutionState, *big.Int, error) {
	index, ok := e.stopPointIndex[string(gasUsed.Bytes())]
	if !ok {
		return nil, nil, errors.New("invalid gas used")
	}
	if err := e.fillInCursors(index); err != nil {
		return nil, nil, err
	}

	cursor := e.cursors[index]
	return NewExecutionState(cursor), cursor.TotalSteps(), nil
}

func (e *ExecutionTracker) GetMachine(gasUsed *big.Int) (machine.Machine, error) {
	index, ok := e.stopPointIndex[string(gasUsed.Bytes())]
	if !ok {
		return nil, errors.New("invalid gas used")
	}
	if err := e.fillInCursors(index); err != nil {
		return nil, err
	}
	return e.lookup.TakeMachine(e.cursors[index].Clone())
}

func IsAssertionValid(assertion *Assertion, execTracker *ExecutionTracker, targetInboxAcc [32]byte) (bool, error) {
	localExecutionState, _, err := execTracker.GetExecutionState(assertion.After.TotalGasConsumed)
	if err != nil {
		return false, err
	}
	if localExecutionState.TotalMessagesRead.Cmp(assertion.After.TotalMessagesRead) < 0 {
		// We didn't read enough messages.
		// This can either mean that our messages lasted longer, or that we are missing messages.
		if localExecutionState.TotalGasConsumed.Cmp(assertion.After.TotalGasConsumed) < 0 && !localExecutionState.IsPermanentlyBlocked() {
			// This means we stopped because we're missing messages,
			// but the on-chain rollup must've had these messages.
			// Error and try again when we have the messages.
			return false, errors.New("Missing messages to evaluate assertion")
		}
		actualEndAcc, expectedEndAcc, err := execTracker.lookup.GetInboxAccPair(localExecutionState.TotalMessagesRead, assertion.After.TotalMessagesRead)
		if err != nil {
			return false, err
		}
		if actualEndAcc != localExecutionState.InboxAcc || expectedEndAcc != targetInboxAcc {
			return false, errors.New("inbox reorg while evaluating assertion")
		}
	} else {
		if localExecutionState.InboxAcc != targetInboxAcc {
			return false, errors.New("inbox reorg while evaluating assertion")
		}
	}

	return assertion.After.Equals(localExecutionState), nil
}
