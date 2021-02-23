package core

import (
	"math/big"
	"time"

	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type MessageStatus uint8

const (
	MessagesEmpty MessageStatus = iota
	MessagesReady
	MessagesSuccess
	MessagesNeedOlder
	MessagesError
)

type ExecutionCursor interface {
	Clone() ExecutionCursor
	MachineHash() common.Hash
	TotalMessagesRead() *big.Int
	InboxHash() common.Hash
	TotalGasConsumed() *big.Int
	TotalSteps() *big.Int
	TotalSendCount() *big.Int
	TotalLogCount() *big.Int

	// TakeMachine takes ownership of machine such that ExecutionCursor will
	// no longer be able to advance.
	TakeMachine() (machine.Machine, error)
}

type ArbCoreLookup interface {
	GetLogCount() (*big.Int, error)
	GetLogs(startIndex, count *big.Int) ([]value.Value, error)

	GetSendCount() (*big.Int, error)
	GetSends(startIndex, count *big.Int) ([][]byte, error)

	GetMessageCount() (*big.Int, error)
	GetMessages(startIndex, count *big.Int) ([]inbox.InboxMessage, error)

	GetSendAcc(startAcc common.Hash, startIndex, count *big.Int) (common.Hash, error)
	GetLogAcc(startAcc common.Hash, startIndex, count *big.Int) (common.Hash, error)

	// GetExecutionCursor returns a cursor containing the machine after executing totalGasUsed
	// from the original machine
	GetExecutionCursor(totalGasUsed *big.Int) (ExecutionCursor, error)

	// Advance executes as much as it can without going over maxGas or
	// optionally until it reaches or goes over maxGas
	AdvanceExecutionCursor(executionCursor ExecutionCursor, maxGas *big.Int, goOverGas bool) error
}

type ArbCoreInbox interface {
	DeliverMessages(messages []inbox.InboxMessage, previousInboxHash common.Hash, lastBlockComplete bool) bool
	MessagesStatus() (MessageStatus, error)
}

func DeliverMessagesAndWait(db ArbCoreInbox, messages []inbox.InboxMessage, previousInboxHash common.Hash, lastBlockComplete bool) (bool, error) {
	if !db.DeliverMessages(messages, previousInboxHash, lastBlockComplete) {
		return false, errors.New("unable to deliver messages")
	}

	start := time.Now()
	var status MessageStatus
	var err error
	for {
		status, err = db.MessagesStatus()
		if err != nil {
			return false, err
		}

		if status == MessagesEmpty {
			return false, errors.New("should have messages")
		}
		if status != MessagesReady {
			break
		}
		if time.Since(start) > time.Second*30 {
			return false, errors.New("timed out adding messages")
		}
		<-time.After(time.Millisecond * 200)
	}
	if status == MessagesSuccess {
		return true, nil
	}
	if status == MessagesNeedOlder {
		return false, nil
	}
	return false, errors.New("Unexpected status")
}

type ArbCore interface {
	ArbCoreLookup
	ArbCoreInbox
	StartThread() bool
	StopThread()
	GetMachineForSideload(uint64) (machine.Machine, error)
	MachineIdle() bool
}

func GetSingleMessage(lookup ArbCoreLookup, index *big.Int) (inbox.InboxMessage, error) {
	messages, err := lookup.GetMessages(index, big.NewInt(1))
	if err != nil {
		return inbox.InboxMessage{}, err
	}
	if len(messages) == 0 {
		return inbox.InboxMessage{}, errors.New("no send found")
	}
	if len(messages) > 1 {
		return inbox.InboxMessage{}, errors.New("too many sends")
	}
	return messages[0], nil
}

func GetSingleSend(lookup ArbCoreLookup, index *big.Int) ([]byte, error) {
	sends, err := lookup.GetSends(index, big.NewInt(1))
	if err != nil {
		return nil, err
	}
	if len(sends) == 0 {
		return nil, errors.New("no send found")
	}
	if len(sends) > 1 {
		return nil, errors.New("too many sends")
	}
	return sends[0], nil
}

func GetSingleLog(lookup ArbCoreLookup, index *big.Int) (value.Value, error) {
	logs, err := lookup.GetLogs(index, big.NewInt(1))
	if err != nil {
		return nil, err
	}
	if len(logs) == 0 {
		return nil, errors.New("no log found")
	}
	if len(logs) > 1 {
		return nil, errors.New("too many logs")
	}
	return logs[0], nil
}

type ExecutionState struct {
	MachineHash       common.Hash
	TotalMessagesRead *big.Int
	TotalGasConsumed  *big.Int
	TotalSendCount    *big.Int
	TotalLogCount     *big.Int
}

func NewExecutionState(c ExecutionCursor) *ExecutionState {
	return &ExecutionState{
		MachineHash:       c.MachineHash(),
		TotalMessagesRead: c.TotalMessagesRead(),
		TotalGasConsumed:  c.TotalGasConsumed(),
		TotalSendCount:    c.TotalSendCount(),
		TotalLogCount:     c.TotalLogCount(),
	}
}

func (e *ExecutionState) Equals(o *ExecutionState) bool {
	return e.MachineHash == o.MachineHash &&
		e.TotalMessagesRead.Cmp(o.TotalMessagesRead) == 0 &&
		e.TotalGasConsumed.Cmp(o.TotalGasConsumed) == 0 &&
		e.TotalSendCount.Cmp(o.TotalSendCount) == 0 &&
		e.TotalLogCount.Cmp(o.TotalLogCount) == 0
}

type ExecutionInfo struct {
	Before  *ExecutionState
	After   *ExecutionState
	SendAcc common.Hash
	LogAcc  common.Hash
}

func (e *ExecutionInfo) Equals(o *ExecutionInfo) bool {
	return e.Before.Equals(o.Before) &&
		e.After.Equals(o.After) &&
		e.SendAcc == o.SendAcc &&
		e.LogAcc == o.SendAcc
}

func (e *ExecutionInfo) GasUsed() *big.Int {
	return new(big.Int).Sub(e.After.TotalGasConsumed, e.Before.TotalGasConsumed)
}

func (e *ExecutionInfo) SendCount() *big.Int {
	return new(big.Int).Sub(e.After.TotalSendCount, e.Before.TotalSendCount)
}

func (e *ExecutionInfo) LogCount() *big.Int {
	return new(big.Int).Sub(e.After.TotalLogCount, e.Before.TotalLogCount)
}

func (e *ExecutionInfo) InboxMessagesRead() *big.Int {
	return new(big.Int).Sub(e.After.TotalMessagesRead, e.Before.TotalMessagesRead)
}

type LogConsumer interface {
	AddLogs(avmLogs []value.Value) error
	DeleteLogs(avmLogs []value.Value) error
}

type LogsCursor interface {
	LogsCursorRequest(cursorIndex *big.Int, count *big.Int) error
	LogsCursorGetLogs(cursorIndex *big.Int) ([]value.Value, error)
	LogsCursorGetDeletedLogs(cursorIndex *big.Int) ([]value.Value, error)
	LogsCursorClearError(cursorIndex *big.Int) error
	LogsCursorConfirmReceived(cursorIndex *big.Int) (bool, error)
}
