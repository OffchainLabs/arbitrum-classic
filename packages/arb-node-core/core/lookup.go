package core

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/pkg/errors"
	"math/big"
	"sort"
)

type ExecutionCursor interface {
	MachineHash() common.Hash
	NextInboxMessageIndex() *big.Int
	InboxHash() common.Hash
	TotalGasConsumed() *big.Int
	TotalSendCount() *big.Int
	TotalLogCount() *big.Int
}

type ExecutionState struct {
	MachineHash      common.Hash
	InboxIndex       *big.Int
	InboxHash        common.Hash
	TotalGasConsumed *big.Int
	TotalSendCount   *big.Int
	TotalLogCount    *big.Int
}

func NewExecutionState(c ExecutionCursor) *ExecutionState {
	return &ExecutionState{
		MachineHash:      c.MachineHash(),
		InboxIndex:       c.NextInboxMessageIndex(),
		InboxHash:        c.InboxHash(),
		TotalGasConsumed: c.TotalGasConsumed(),
		TotalSendCount:   c.TotalSendCount(),
		TotalLogCount:    c.TotalLogCount(),
	}
}

func (e *ExecutionState) Equals(o *ExecutionState) bool {
	return e.MachineHash == o.MachineHash &&
		e.InboxIndex.Cmp(o.InboxIndex) == 0 &&
		e.InboxHash == o.InboxHash &&
		e.TotalGasConsumed.Cmp(o.TotalGasConsumed) == 0 &&
		e.TotalSendCount.Cmp(o.TotalSendCount) == 0 &&
		e.TotalLogCount.Cmp(o.TotalLogCount) == 0
}

type ValidatorLookup interface {
	GetSends(startIndex *big.Int, count *big.Int) ([][]byte, error)
	GetMessages(startIndex *big.Int, count *big.Int) ([]inbox.InboxMessage, error)
	GetInboxDelta(startIndex *big.Int, count *big.Int) (common.Hash, error)

	GetInboxAcc(index *big.Int) (common.Hash, error)
	GetSendAcc(startAcc common.Hash, startIndex *big.Int, count *big.Int) (common.Hash, error)
	GetLogAcc(startAcc common.Hash, startIndex *big.Int, count *big.Int) (common.Hash, error)

	// GetMachine returns the image of the machine after executing totalGasUsed
	// from the original machine
	GetCursor(totalGasUsed *big.Int) (ExecutionCursor, error)

	// GetExecutionInfo executes as much as it can not over maxGas
	MoveExecutionCursor(
		start ExecutionCursor,
		maxGas *big.Int,
		goOverGas bool,
	) (ExecutionCursor, error)

	GetMachine(cursor ExecutionCursor) (machine.Machine, error)
}

type SimpleExecutionInfo struct {
	Before *ExecutionState
	After  *ExecutionState
}

func (e *SimpleExecutionInfo) Equals(o *SimpleExecutionInfo) bool {
	return e.Before.Equals(o.Before) && e.After.Equals(o.After)
}

func (e *SimpleExecutionInfo) GasUsed() *big.Int {
	return new(big.Int).Sub(e.After.TotalGasConsumed, e.Before.TotalGasConsumed)
}

func (e *SimpleExecutionInfo) SendCount() *big.Int {
	return new(big.Int).Sub(e.After.TotalSendCount, e.Before.TotalSendCount)
}

func (e *SimpleExecutionInfo) LogCount() *big.Int {
	return new(big.Int).Sub(e.After.TotalLogCount, e.Before.TotalLogCount)
}

func (e *SimpleExecutionInfo) InboxMessagesRead() *big.Int {
	return new(big.Int).Sub(e.After.InboxIndex, e.Before.InboxIndex)
}

type ExecutionInfo struct {
	*SimpleExecutionInfo
	SendAcc common.Hash
	LogAcc  common.Hash
}

func (e *ExecutionInfo) Equals(o *ExecutionInfo) bool {
	return e.SimpleExecutionInfo.Equals(o.SimpleExecutionInfo) &&
		e.SendAcc == o.SendAcc &&
		e.LogAcc == o.SendAcc
}

type AssertionInfo struct {
	*ExecutionInfo
	InboxDelta common.Hash
}

func (a *AssertionInfo) Equals(o *AssertionInfo) bool {
	return a.ExecutionInfo.Equals(o.ExecutionInfo) &&
		a.InboxDelta == o.InboxDelta
}

type ExecutionTracker struct {
	lookup    ValidatorLookup
	goOverGas bool

	sortedStopPoints []*big.Int
	stopPointIndex   map[string]int

	cursors  []ExecutionCursor
	sendAccs []common.Hash
	logAccs  []common.Hash
}

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

func NewExecutionTracker(lookup ValidatorLookup, cursor ExecutionCursor, goOverGas bool, stopPoints []*big.Int) *ExecutionTracker {
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

func (e *ExecutionTracker) fillInCursorSnapshots(max int) error {
	for i := len(e.cursors) - 1; i < max; i++ {
		cursor := e.cursors[len(e.cursors)-1]
		nextStopPoint := e.sortedStopPoints[i]
		gasToExecute := new(big.Int).Sub(nextStopPoint, cursor.TotalGasConsumed())
		newCursor, err := e.lookup.MoveExecutionCursor(cursor, gasToExecute, e.goOverGas)
		if err != nil {
			return err
		}
		e.cursors = append(e.cursors, newCursor)
	}
	return nil
}

func (e *ExecutionTracker) fillInAccs(max int) error {
	if err := e.fillInCursorSnapshots(max); err != nil {
		return err
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

func (e *ExecutionTracker) GenerateExecutionInfo(gasUsed *big.Int) (*ExecutionInfo, error) {
	index, ok := e.stopPointIndex[string(gasUsed.Bytes())]
	if !ok {
		return nil, errors.New("invalid gas used")
	}
	if err := e.fillInAccs(index); err != nil {
		return nil, err
	}

	return &ExecutionInfo{
		SimpleExecutionInfo: &SimpleExecutionInfo{
			Before: NewExecutionState(e.cursors[0]),
			After:  NewExecutionState(e.cursors[index]),
		},
		SendAcc: e.sendAccs[index],
		LogAcc:  e.logAccs[index],
	}, nil
}

func JudgeAssertion(lookup ValidatorLookup, assertion *Assertion, execTracker *ExecutionTracker) (ChallengeKind, error) {
	afterInboxHash, err := lookup.GetInboxAcc(assertion.After.InboxIndex)
	if err != nil {
		return 0, err
	}
	if assertion.After.InboxHash != afterInboxHash {
		// Failed inbox consistency
		return INBOX_CONSISTENCY, nil
	}
	inboxDelta, err := lookup.GetInboxDelta(assertion.Before.InboxIndex, assertion.InboxMessagesRead())
	if err != nil {
		return 0, err
	}
	if assertion.InboxDelta != inboxDelta {
		// Failed inbox delta
		return INBOX_DELTA, nil
	}

	localExecutionInfo, err := execTracker.GenerateExecutionInfo(assertion.GasUsed())
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
