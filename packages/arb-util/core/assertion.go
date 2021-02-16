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
	PrevProposedBlock *big.Int
	PrevInboxMaxCount *big.Int
	*ExecutionInfo
}

func NewAssertionFromFields(a [4][32]byte, b [10]*big.Int) *Assertion {
	beforeState := &ExecutionState{
		MachineHash:       a[0],
		TotalMessagesRead: b[2],
		TotalGasConsumed:  b[1],
		TotalSendCount:    b[3],
		TotalLogCount:     b[4],
	}

	return &Assertion{
		PrevProposedBlock: b[0],
		PrevInboxMaxCount: b[5],
		ExecutionInfo: &ExecutionInfo{
			Before: beforeState,
			After: &ExecutionState{
				MachineHash:       a[3],
				TotalMessagesRead: new(big.Int).Add(beforeState.TotalMessagesRead, b[6]),
				TotalGasConsumed:  new(big.Int).Add(beforeState.TotalGasConsumed, b[7]),
				TotalSendCount:    new(big.Int).Add(beforeState.TotalSendCount, b[8]),
				TotalLogCount:     new(big.Int).Add(beforeState.TotalLogCount, b[9]),
			},
			SendAcc: a[1],
			LogAcc:  a[2],
		},
	}
}

func (a *Assertion) BytesFields() [4][32]byte {
	return [4][32]byte{
		a.Before.MachineHash,
		a.SendAcc,
		a.LogAcc,
		a.After.MachineHash,
	}
}

func (a *Assertion) IntFields() [10]*big.Int {
	return [10]*big.Int{
		a.PrevProposedBlock,
		a.Before.TotalGasConsumed,
		a.Before.TotalMessagesRead,
		a.Before.TotalSendCount,
		a.Before.TotalLogCount,
		a.PrevInboxMaxCount,
		a.InboxMessagesRead(),
		a.GasUsed(),
		a.SendCount(),
		a.LogCount(),
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

func assertionHash(
	gasUsed *big.Int,
	assertionRest common.Hash,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint256(gasUsed),
		hashing.Bytes32(assertionRest),
	)
}

func assertionRestHash(
	totalMessagesRead *big.Int,
	machineState common.Hash,
	sendAcc common.Hash,
	sendCount *big.Int,
	logAcc common.Hash,
	logCount *big.Int,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint256(totalMessagesRead),
		hashing.Bytes32(machineState),
		hashing.Bytes32(sendAcc),
		hashing.Uint256(sendCount),
		hashing.Bytes32(logAcc),
		hashing.Uint256(logCount),
	)
}

func (a *Assertion) BeforeExecutionHash() common.Hash {
	restBefore := assertionRestHash(
		a.Before.TotalMessagesRead,
		a.Before.MachineHash,
		common.Hash{},
		big.NewInt(0),
		common.Hash{},
		big.NewInt(0),
	)
	return assertionHash(big.NewInt(0), restBefore)
}

func (a *Assertion) AfterExecutionHash() common.Hash {
	restAfter := assertionRestHash(
		a.After.TotalMessagesRead,
		a.After.MachineHash,
		a.SendAcc,
		a.SendCount(),
		a.LogAcc,
		a.LogCount(),
	)
	return assertionHash(a.GasUsed(), restAfter)
}

func (a *Assertion) ExecutionHash() common.Hash {
	return BisectionChunkHash(
		big.NewInt(0),
		a.GasUsed(),
		a.BeforeExecutionHash(),
		a.AfterExecutionHash(),
	)
}

func (a *Assertion) CheckTime(arbGasSpeedLimitPerBlock *big.Int) *big.Int {
	return new(big.Int).Div(a.GasUsed(), arbGasSpeedLimitPerBlock)
}
