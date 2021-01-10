package core

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"math/big"
)

type ValidatorLookup interface {
	GenerateLogAccumulator(startIndex *big.Int, count *big.Int) (common.Hash, error)
	GetSends(startIndex *big.Int, count *big.Int) ([][]byte, error)
	GetInboxAcc(index *big.Int) (common.Hash, error)
	GetMessages(startIndex *big.Int, count *big.Int) ([]inbox.InboxMessage, error)

	// GetMachine returns the image of the machine after executing totalGasUsed
	// from the original machine
	GetMachine(totalGasUsed *big.Int) (machine.Machine, error)

	// GetExecutionInfo tries to execute targetGas steps, but only reads up to
	// maxMessages messages and stops short if it runs out of messages and needs
	// more to continue
	GetExecutionInfoWithMaxMessages(startMachine machine.Machine, targetGas *big.Int, maxMessages *big.Int) (*ExecutionInfo, error)

	// GetExecutionInfo executes as much as it can not over maxGas
	GetExecutionInfo(startMachine machine.Machine, maxGas *big.Int) (*AssertionInfo, error)
}

type ExecutionInfo struct {
	BeforeMachineHash common.Hash
	InboxMessagesRead *big.Int
	GasUsed           *big.Int
	SendAcc           common.Hash
	SendCount         *big.Int
	LogAcc            common.Hash
	LogCount          *big.Int
	AfterMachineHash  common.Hash
}

func (e *ExecutionInfo) Equals(o *ExecutionInfo) bool {
	return e.BeforeMachineHash == o.BeforeMachineHash &&
		e.InboxMessagesRead.Cmp(o.InboxMessagesRead) == 0 &&
		e.GasUsed.Cmp(o.GasUsed) == 0 &&
		e.SendAcc == o.SendAcc &&
		e.SendCount.Cmp(o.SendCount) == 0 &&
		e.LogAcc == o.LogAcc &&
		e.LogCount.Cmp(o.LogCount) == 0 &&
		e.AfterMachineHash == o.AfterMachineHash
}

type AssertionInfo struct {
	InboxDelta     common.Hash
	ExecInfo       *ExecutionInfo
	AfterInboxHash common.Hash
}
