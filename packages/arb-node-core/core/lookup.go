package core

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/pkg/errors"
	"math/big"
)

type ExecutionCursor interface {
	Clone() (ExecutionCursor, error)
	MachineHash() common.Hash
	NextInboxMessageIndex() *big.Int
	InboxHash() common.Hash
	TotalGasConsumed() *big.Int
	TotalSendCount() *big.Int
	TotalLogCount() *big.Int

	// TakeMachine takes ownership of machine such that ExecutionCursor will
	// no longer be able to advance.
	TakeMachine() (machine.Machine, error)
}

type ValidatorLookup interface {
	GetLogs(startIndex, count *big.Int) ([]value.Value, error)
	GetSends(startIndex, count *big.Int) ([][]byte, error)
	GetMessages(startIndex, count *big.Int) ([]inbox.InboxMessage, error)
	GetInboxDelta(startIndex, count *big.Int) (common.Hash, error)

	GetInboxAcc(index *big.Int) (common.Hash, error)
	GetSendAcc(startAcc common.Hash, startIndex, count *big.Int) (common.Hash, error)
	GetLogAcc(startAcc common.Hash, startIndex, count *big.Int) (common.Hash, error)

	// GetExecutionCursor returns a cursor containing the machine after executing totalGasUsed
	// from the original machine
	GetExecutionCursor(totalGasUsed *big.Int) (ExecutionCursor, error)

	// Advance executes as much as it can without going over maxGas or
	// optionally until it reaches or goes over maxGas
	Advance(executionCursor ExecutionCursor, maxGas *big.Int, goOverGas bool) error
}

func GetSingleSend(lookup ValidatorLookup, index *big.Int) ([]byte, error) {
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

func GetSingleLog(lookup ValidatorLookup, index *big.Int) (value.Value, error) {
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
	return new(big.Int).Sub(e.After.InboxIndex, e.Before.InboxIndex)
}
