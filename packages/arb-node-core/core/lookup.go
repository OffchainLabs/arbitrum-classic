package core

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"math/big"
)

type ExecutionCursor interface {
	Clone() ExecutionCursor
	MachineHash() common.Hash
	NextInboxMessageIndex() *big.Int
	InboxHash() common.Hash
	TotalGasConsumed() *big.Int
	TotalSendCount() *big.Int
	TotalLogCount() *big.Int

	// Advance executes as much as it can without going over maxGas or
	// optionally until it goes over maxGas
	Advance(
		maxGas *big.Int,
		goOverGas bool,
	) error

	// TakeMachine takes ownership of machine such that ExecutionCursor will
	// no longer be able to advance.
	TakeMachine() (machine.Machine, error)
}

type ValidatorLookup interface {
	GetSends(startIndex *big.Int, count *big.Int) ([][]byte, error)
	GetMessages(startIndex *big.Int, count *big.Int) ([]inbox.InboxMessage, error)
	GetInboxDelta(startIndex *big.Int, count *big.Int) (common.Hash, error)

	GetInboxAcc(index *big.Int) (common.Hash, error)
	GetSendAcc(startAcc common.Hash, startIndex *big.Int, count *big.Int) (common.Hash, error)
	GetLogAcc(startAcc common.Hash, startIndex *big.Int, count *big.Int) (common.Hash, error)

	// GetCursor returns a cursor containing the machine after executing totalGasUsed
	// from the original machine
	GetCursor(totalGasUsed *big.Int) (ExecutionCursor, error)
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
