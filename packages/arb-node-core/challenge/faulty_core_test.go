package challenge

import (
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
)

func distortHash(hash common.Hash) common.Hash {
	hash[0] = 0xde
	hash[1] = 0xad
	hash[2] = 0xbe
	hash[3] = 0xef
	return hash
}

type faultConfig struct {
	distortMachineAtGas *big.Int
	messagesReadCap     *big.Int
	phantomMessageAtGas *big.Int
}

type faultyExecutionCursor struct {
	config faultConfig
	core.ExecutionCursor
}

func (e faultyExecutionCursor) Clone() core.ExecutionCursor {
	return faultyExecutionCursor{
		config:          e.config,
		ExecutionCursor: e.ExecutionCursor.Clone(),
	}
}

func (e faultyExecutionCursor) MachineHash() common.Hash {
	hash := e.ExecutionCursor.MachineHash()
	if e.config.distortMachineAtGas != nil && e.ExecutionCursor.TotalGasConsumed().Cmp(e.config.distortMachineAtGas) >= 0 {
		hash = distortHash(hash)
	}
	return hash
}

func (e faultyExecutionCursor) TotalMessagesRead() *big.Int {
	messages := e.ExecutionCursor.TotalMessagesRead()
	if e.config.phantomMessageAtGas != nil && e.ExecutionCursor.TotalGasConsumed().Cmp(e.config.phantomMessageAtGas) > 0 {
		messages = new(big.Int).Add(messages, big.NewInt(1))
	}
	if e.config.messagesReadCap != nil && messages.Cmp(e.config.messagesReadCap) > 0 {
		messages = new(big.Int).Set(e.config.messagesReadCap)
	}
	return messages
}

type faultyCore struct {
	config faultConfig
	core.ArbCore
}

func newFaultyCore(core core.ArbCore, config faultConfig) faultyCore {
	return faultyCore{
		config:  config,
		ArbCore: core,
	}
}

func (c faultyCore) GetExecutionCursor(totalGasUsed *big.Int) (core.ExecutionCursor, error) {
	cursor, err := c.ArbCore.GetExecutionCursor(totalGasUsed)
	if err != nil {
		return nil, err
	}
	return faultyExecutionCursor{
		config:          c.config,
		ExecutionCursor: cursor,
	}, nil
}

func (c faultyCore) AdvanceExecutionCursor(executionCursor core.ExecutionCursor, maxGas *big.Int, goOverGas bool) error {
	return c.ArbCore.AdvanceExecutionCursor(executionCursor.(faultyExecutionCursor).ExecutionCursor, maxGas, goOverGas)
}
