package core

import (
	"github.com/pkg/errors"
	"math/big"
	"sort"

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

func NewExecutionTracker(lookup ArbCoreLookup, cursor ExecutionCursor, goOverGas bool, stopPoints []*big.Int) *ExecutionTracker {
	sort.Sort(BigIntList(stopPoints))
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
	for i := len(e.cursors) - 1; i < max; i++ {
		nextCursor := e.cursors[len(e.cursors)-1].Clone()
		nextStopPoint := e.sortedStopPoints[i]
		gasToExecute := new(big.Int).Sub(nextStopPoint, nextCursor.TotalGasConsumed())
		err := e.lookup.AdvanceExecutionCursor(nextCursor, gasToExecute, e.goOverGas)
		if err != nil {
			return err
		}
		e.cursors = append(e.cursors, nextCursor)
	}
	return nil
}

func (e *ExecutionTracker) fillInAccs(max int) error {
	if err := e.fillInCursors(max); err != nil {
		return err
	}
	if len(e.logAccs) < 2 {
		// Nothing to fill in
		return nil
	}

	for i := len(e.logAccs) - 1; i < max; i++ {
		prevCursor := e.cursors[i-1]
		cursor := e.cursors[i]
		prevSendAcc := e.sendAccs[i-1]
		prevLogAcc := e.logAccs[i-1]
		sendCount := new(big.Int).Sub(prevCursor.TotalSendCount(), cursor.TotalSendCount())
		sendAcc, err := e.lookup.GetSendAcc(prevSendAcc, cursor.TotalSendCount(), sendCount)
		if err != nil {
			return err
		}
		logCount := new(big.Int).Sub(prevCursor.TotalLogCount(), cursor.TotalLogCount())
		logAcc, err := e.lookup.GetLogAcc(prevLogAcc, cursor.TotalLogCount(), logCount)
		if err != nil {
			return err
		}
		e.sendAccs = append(e.sendAccs, sendAcc)
		e.logAccs = append(e.logAccs, logAcc)
	}
	return nil
}

func (e *ExecutionTracker) GetExecutionInfo(gasUsed *big.Int) (*ExecutionInfo, error) {
	index, ok := e.stopPointIndex[string(gasUsed.Bytes())]
	if !ok {
		return nil, errors.New("invalid gas used")
	}
	if err := e.fillInAccs(index); err != nil {
		return nil, err
	}

	return &ExecutionInfo{
		Before:  NewExecutionState(e.cursors[0]),
		After:   NewExecutionState(e.cursors[index]),
		SendAcc: e.sendAccs[index],
		LogAcc:  e.logAccs[index],
	}, nil
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

func JudgeAssertion(lookup ArbCoreLookup, assertion *Assertion, execTracker *ExecutionTracker) (ChallengeKind, error) {
	var afterInboxIndex big.Int
	if assertion.After.TotalMessagesRead.Cmp(big.NewInt(0)) != 0 {
		afterInboxIndex = *assertion.After.TotalMessagesRead
		//afterInboxIndex.Sub(&afterInboxIndex, big.NewInt(1))
	}
	afterInboxHash, err := lookup.GetInboxAcc(&afterInboxIndex)
	if err != nil {
		return 0, err
	}
	if assertion.After.InboxHash != afterInboxHash {
		// Failed inbox consistency
		return INBOX_CONSISTENCY, nil
	}
	var beforeInboxIndex big.Int
	if assertion.Before.TotalMessagesRead.Cmp(big.NewInt(0)) != 0 {
		beforeInboxIndex = *assertion.Before.TotalMessagesRead
		//beforeInboxIndex.Sub(&beforeInboxIndex, big.NewInt(1))
	}
	inboxDelta, err := lookup.GetInboxDelta(&beforeInboxIndex, assertion.InboxMessagesRead())
	if err != nil {
		return 0, err
	}
	if assertion.InboxDelta != inboxDelta {
		// Failed inbox delta
		return INBOX_DELTA, nil
	}

	localExecutionInfo, err := execTracker.GetExecutionInfo(assertion.GasUsed())
	if err != nil {
		return 0, err
	}
	if localExecutionInfo.InboxMessagesRead().Cmp(assertion.InboxMessagesRead()) > 0 {
		// Execution read more messages than provided so assertion should have
		// stopped short
		return STOPPED_SHORT, nil
	}

	if !assertion.ExecutionInfo.Equals(localExecutionInfo) {
		return EXECUTION, nil
	}
	return NO_CHALLENGE, nil
}
