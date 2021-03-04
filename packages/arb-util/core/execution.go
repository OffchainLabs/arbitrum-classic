package core

import (
	"math/big"
	"sort"

	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
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

	cursors  []ExecutionCursor
	sendAccs []common.Hash
	logAccs  []common.Hash
}

func NewExecutionTracker(lookup ArbCoreLookup, cursor ExecutionCursor, goOverGas bool, stopPointsArg []*big.Int) *ExecutionTracker {
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
	cursors := make([]ExecutionCursor, 0, len(stopPoints)+1)
	cursors = append(cursors, cursor)
	sendAccs := make([]common.Hash, 0, len(stopPoints)+1)
	sendAccs = append(sendAccs, common.Hash{})
	logAccs := make([]common.Hash, 0, len(stopPoints)+1)
	logAccs = append(logAccs, common.Hash{})

	stopPointIndex := make(map[string]int)
	for i, stopPoint := range stopPoints {
		stopPointIndex[string(stopPoint.Bytes())] = i
	}
	return &ExecutionTracker{
		lookup:           lookup,
		goOverGas:        goOverGas,
		sortedStopPoints: stopPoints,
		stopPointIndex:   stopPointIndex,
		sendAccs:         sendAccs,
		logAccs:          logAccs,
		cursors:          cursors,
	}
}

func (e *ExecutionTracker) fillInCursors(max int) error {
	for i := len(e.cursors); i <= max; i++ {
		nextCursor := e.cursors[i-1].Clone()
		nextStopPoint := e.sortedStopPoints[i]
		if nextStopPoint.Cmp(nextCursor.TotalGasConsumed()) > 0 {
			gasToExecute := new(big.Int).Sub(nextStopPoint, nextCursor.TotalGasConsumed())
			err := e.lookup.AdvanceExecutionCursor(nextCursor, gasToExecute, e.goOverGas)
			if err != nil {
				return err
			}
		}
		e.cursors = append(e.cursors, nextCursor)
	}
	return nil
}

func (e *ExecutionTracker) fillInAccs(max int) error {
	if err := e.fillInCursors(max); err != nil {
		return err
	}

	for i := len(e.logAccs); i <= max; i++ {
		prevCursor := e.cursors[i-1]
		cursor := e.cursors[i]
		prevSendAcc := e.sendAccs[i-1]
		prevLogAcc := e.logAccs[i-1]
		sendCount := new(big.Int).Sub(cursor.TotalSendCount(), prevCursor.TotalSendCount())
		sendAcc, err := e.lookup.GetSendAcc(prevSendAcc, prevCursor.TotalSendCount(), sendCount)
		if err != nil {
			return err
		}
		logCount := new(big.Int).Sub(cursor.TotalLogCount(), prevCursor.TotalLogCount())
		logAcc, err := e.lookup.GetLogAcc(prevLogAcc, prevCursor.TotalLogCount(), logCount)
		if err != nil {
			return err
		}
		e.sendAccs = append(e.sendAccs, sendAcc)
		e.logAccs = append(e.logAccs, logAcc)
	}
	return nil
}

func (e *ExecutionTracker) GetExecutionInfo(gasUsed *big.Int) (*ExecutionInfo, *big.Int, error) {
	index, ok := e.stopPointIndex[string(gasUsed.Bytes())]
	if !ok {
		return nil, nil, errors.New("invalid gas used")
	}
	if err := e.fillInAccs(index); err != nil {
		return nil, nil, err
	}

	return &ExecutionInfo{
		Before:  NewExecutionState(e.cursors[0]),
		After:   NewExecutionState(e.cursors[index]),
		SendAcc: e.sendAccs[index],
		LogAcc:  e.logAccs[index],
	}, e.cursors[index].TotalSteps(), nil
}

func (e *ExecutionTracker) GetMachine(gasUsed *big.Int) (machine.Machine, error) {
	index, ok := e.stopPointIndex[string(gasUsed.Bytes())]
	if !ok {
		return nil, errors.New("invalid gas used")
	}
	if err := e.fillInCursors(index); err != nil {
		return nil, err
	}
	return e.cursors[index].Clone().TakeMachine()
}

func IsAssertionValid(assertion *Assertion, execTracker *ExecutionTracker, targetInboxAcc [32]byte) (bool, error) {
	localExecutionInfo, _, err := execTracker.GetExecutionInfo(assertion.After.TotalGasConsumed)
	if err != nil {
		return false, err
	}
	if localExecutionInfo.InboxMessagesRead().Cmp(assertion.InboxMessagesRead()) < 0 {
		// We didn't read enough messages.
		// This can either mean that our messages lasted longer, or that we are missing messages.
		if localExecutionInfo.After.TotalGasConsumed.Cmp(assertion.After.TotalGasConsumed) < 0 {
			// This means we stopped because we're missing messages,
			// but the on-chain rollup must've had these messages.
			// Error and try again when we have the messages.
			return false, errors.New("Missing messages to evaluate assertion")
		}
		actualEndAcc, expectedEndAcc, err := execTracker.lookup.GetInboxAccPair(localExecutionInfo.After.TotalMessagesRead, assertion.After.TotalMessagesRead)
		if err != nil {
			return false, err
		}
		if actualEndAcc != localExecutionInfo.After.InboxAcc || expectedEndAcc != targetInboxAcc {
			return false, errors.New("inbox reorg while evaluating assertion")
		}
	} else {
		if localExecutionInfo.After.InboxAcc != targetInboxAcc {
			return false, errors.New("inbox reorg while evaluating assertion")
		}
	}

	return assertion.ExecutionInfo.Equals(localExecutionInfo), nil
}
