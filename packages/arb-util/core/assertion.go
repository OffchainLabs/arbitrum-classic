package core

import (
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
)

type NodeState struct {
	ProposedBlock *big.Int
	InboxMaxCount *big.Int
	*ExecutionState
}

type Assertion struct {
	Before *ExecutionState `json:"beforeState"`
	After  *ExecutionState `json:"afterState"`
}

func newExecutionStateFromFields(a [3][32]byte, b [4]*big.Int) *ExecutionState {
	return &ExecutionState{
		TotalGasConsumed:  b[0],
		MachineHash:       a[0],
		TotalMessagesRead: b[1],
		TotalSendCount:    b[2],
		TotalLogCount:     b[3],
		SendAcc:           a[1],
		LogAcc:            a[2],
	}
}

func NewAssertionFromFields(a [2][3][32]byte, b [2][4]*big.Int) *Assertion {
	return &Assertion{
		Before: newExecutionStateFromFields(a[0], b[0]),
		After:  newExecutionStateFromFields(a[1], b[1]),
	}
}

func (a *Assertion) InitialExecutionBisection() *Bisection {
	return &Bisection{
		ChallengedSegment: &ChallengeSegment{
			Start:  a.Before.TotalGasConsumed,
			Length: new(big.Int).Sub(a.After.TotalGasConsumed, a.Before.TotalGasConsumed),
		},
		Cuts: []common.Hash{a.Before.CutHash(), a.After.CutHash()},
	}
}

func stateByteFields(s *ExecutionState) [3][32]byte {
	return [3][32]byte{
		s.MachineHash,
		s.SendAcc,
		s.LogAcc,
	}
}

func (a *Assertion) BytesFields() [2][3][32]byte {
	return [2][3][32]byte{
		stateByteFields(a.Before),
		stateByteFields(a.After),
	}
}

func stateIntFields(s *ExecutionState) [4]*big.Int {
	return [4]*big.Int{
		s.TotalGasConsumed,
		s.TotalMessagesRead,
		s.TotalSendCount,
		s.TotalLogCount,
	}
}

func (a *Assertion) IntFields() [2][4]*big.Int {
	return [2][4]*big.Int{
		stateIntFields(a.Before),
		stateIntFields(a.After),
	}
}

func BisectionChunkHash(
	segmentStart *big.Int,
	segmentLength *big.Int,
	startHash common.Hash,
	endHash common.Hash,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint256(segmentStart),
		hashing.Uint256(segmentLength),
		hashing.Bytes32(startHash),
		hashing.Bytes32(endHash),
	)
}

func (a *Assertion) BeforeExecutionHash() common.Hash {
	return a.Before.CutHash()
}

func (a *Assertion) AfterExecutionHash() common.Hash {
	return a.After.CutHash()
}

func (a *Assertion) ExecutionHash() common.Hash {
	return BisectionChunkHash(
		a.Before.TotalGasConsumed,
		new(big.Int).Sub(a.After.TotalGasConsumed, a.Before.TotalGasConsumed),
		a.BeforeExecutionHash(),
		a.AfterExecutionHash(),
	)
}

func (a *Assertion) GasUsed() *big.Int {
	return new(big.Int).Sub(a.After.TotalGasConsumed, a.Before.TotalGasConsumed)
}

func (a *Assertion) CheckTime(arbGasSpeedLimitPerBlock *big.Int) *big.Int {
	return new(big.Int).Div(a.GasUsed(), arbGasSpeedLimitPerBlock)
}
