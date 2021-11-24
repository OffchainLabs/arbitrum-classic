package challenge

import (
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

func distortHash(hash common.Hash) common.Hash {
	hash[0] = 0xde
	hash[1] = 0xad
	hash[2] = 0xbe
	hash[3] = 0xef
	return hash
}

type FaultConfig struct {
	DistortMachineAtGas *big.Int
	MessagesReadCap     *big.Int
	PhantomMessageAtGas *big.Int
	StallMachineAt      *big.Int
}

type FaultyExecutionCursor struct {
	config     FaultConfig
	phantomGas *big.Int
	core.ExecutionCursor
}

func (e FaultyExecutionCursor) Clone() core.ExecutionCursor {
	return FaultyExecutionCursor{
		config:          e.config,
		phantomGas:      new(big.Int).Set(e.phantomGas),
		ExecutionCursor: e.ExecutionCursor.Clone(),
	}
}

func (e FaultyExecutionCursor) MachineHash() common.Hash {
	hash := e.ExecutionCursor.MachineHash()
	if e.config.DistortMachineAtGas != nil && e.ExecutionCursor.TotalGasConsumed().Cmp(e.config.DistortMachineAtGas) >= 0 {
		hash = distortHash(hash)
	}
	return hash
}

func (e FaultyExecutionCursor) TotalMessagesRead() *big.Int {
	messages := e.ExecutionCursor.TotalMessagesRead()
	if e.config.PhantomMessageAtGas != nil && e.ExecutionCursor.TotalGasConsumed().Cmp(e.config.PhantomMessageAtGas) > 0 {
		messages = new(big.Int).Add(messages, big.NewInt(1))
	}
	if e.config.MessagesReadCap != nil && messages.Cmp(e.config.MessagesReadCap) > 0 {
		messages = new(big.Int).Set(e.config.MessagesReadCap)
	}
	return messages
}

func (e FaultyExecutionCursor) TotalGasConsumed() *big.Int {
	gas := e.ExecutionCursor.TotalGasConsumed()
	return new(big.Int).Add(gas, e.phantomGas)
}

type FaultyCore struct {
	config FaultConfig
	core.ArbCore
}

func NewFaultyCore(core core.ArbCore, config FaultConfig) FaultyCore {
	return FaultyCore{
		config:  config,
		ArbCore: core,
	}
}

func (c FaultyCore) GetExecutionCursor(totalGasUsed *big.Int, allowSlowLookup bool) (core.ExecutionCursor, error) {
	cursor, err := c.ArbCore.GetExecutionCursor(totalGasUsed, allowSlowLookup)
	if err != nil {
		return nil, err
	}
	return FaultyExecutionCursor{
		config:          c.config,
		phantomGas:      big.NewInt(0),
		ExecutionCursor: cursor,
	}, nil
}

func (c FaultyCore) prepareCursor(executionCursor core.ExecutionCursor, maxGas *big.Int) (FaultyExecutionCursor, *big.Int, error, bool) {
	faultyCursor := executionCursor.(FaultyExecutionCursor)
	targetGas := new(big.Int).Add(executionCursor.TotalGasConsumed(), maxGas)
	if c.config.StallMachineAt != nil && targetGas.Cmp(c.config.StallMachineAt) > 0 {
		phantomGas := new(big.Int).Sub(targetGas, c.config.StallMachineAt)
		maxGas = new(big.Int).Sub(targetGas, phantomGas)
		faultyCursor.phantomGas.Set(phantomGas)
		if maxGas.Cmp(big.NewInt(0)) <= 0 {
			return FaultyExecutionCursor{}, nil, nil, true
		}
	}
	return faultyCursor, maxGas, nil, false
}

func (c FaultyCore) AdvanceExecutionCursor(executionCursor core.ExecutionCursor, maxGas *big.Int, goOverGas bool, allowSlowLookup bool) error {
	faultyCursor, maxGas, err, done := c.prepareCursor(executionCursor, maxGas)
	if done {
		return err
	}
	return c.ArbCore.AdvanceExecutionCursor(faultyCursor.ExecutionCursor, maxGas, goOverGas, allowSlowLookup)
}

func (c FaultyCore) AdvanceExecutionCursorWithTracing(executionCursor core.ExecutionCursor, maxGas *big.Int, goOverGas bool, allowSlowLookup bool, logNumberStart, logNumberEnd *big.Int) ([]core.MachineEmission, error) {
	faultyCursor, maxGas, err, done := c.prepareCursor(executionCursor, maxGas)
	if done {
		return nil, err
	}
	return c.ArbCore.AdvanceExecutionCursorWithTracing(faultyCursor.ExecutionCursor, maxGas, goOverGas, allowSlowLookup, logNumberStart, logNumberEnd)
}

func (c FaultyCore) GetLastMachine() (machine.Machine, error) {
	return c.ArbCore.GetLastMachine()
}

func (c FaultyCore) TakeMachine(executionCursor core.ExecutionCursor) (machine.Machine, error) {
	faultyCursor := executionCursor.(FaultyExecutionCursor)
	return c.ArbCore.TakeMachine(faultyCursor.ExecutionCursor)
}
