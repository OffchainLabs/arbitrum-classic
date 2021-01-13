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
	*AssertionInfo
}

func NewAssertionFromFields(a [7][32]byte, b [10]*big.Int) *Assertion {
	beforeState := &ExecutionState{
		MachineHash:      a[0],
		InboxIndex:       b[2],
		InboxHash:        a[1],
		TotalGasConsumed: b[1],
		TotalSendCount:   b[3],
		TotalLogCount:    b[4],
	}

	return &Assertion{
		PrevProposedBlock: b[0],
		PrevInboxMaxCount: b[5],
		AssertionInfo: &AssertionInfo{
			ExecutionInfo: &ExecutionInfo{
				Before: beforeState,
				After: &ExecutionState{
					MachineHash:      a[6],
					InboxIndex:       new(big.Int).Add(beforeState.InboxIndex, b[6]),
					InboxHash:        a[3],
					TotalGasConsumed: new(big.Int).Add(beforeState.TotalGasConsumed, b[7]),
					TotalSendCount:   new(big.Int).Add(beforeState.TotalSendCount, b[8]),
					TotalLogCount:    new(big.Int).Add(beforeState.TotalLogCount, b[9]),
				},
				SendAcc: a[4],
				LogAcc:  a[5],
			},
			InboxDelta: a[2],
		},
	}
}

func (a *Assertion) BytesFields() [7][32]byte {
	return [7][32]byte{
		a.Before.MachineHash,
		a.Before.InboxHash,
		a.InboxDelta,
		a.SendAcc,
		a.LogAcc,
		a.After.InboxHash,
		a.After.MachineHash,
	}
}

func (a *Assertion) IntFields() [10]*big.Int {
	return [10]*big.Int{
		a.PrevProposedBlock,
		a.Before.TotalGasConsumed,
		a.Before.InboxIndex,
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

func InboxDeltaHash(inboxAcc, deltaAcc common.Hash) common.Hash {
	return hashing.SoliditySHA3(hashing.Bytes32(inboxAcc), hashing.Bytes32(deltaAcc))
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
	inboxDelta common.Hash,
	machineState common.Hash,
	sendAcc common.Hash,
	sendCount *big.Int,
	logAcc common.Hash,
	logCount *big.Int,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(inboxDelta),
		hashing.Bytes32(machineState),
		hashing.Bytes32(sendAcc),
		hashing.Uint256(sendCount),
		hashing.Bytes32(logAcc),
		hashing.Uint256(logCount),
	)
}

func (a *Assertion) InboxConsistencyHash(inboxTopHash common.Hash, inboxTopCount *big.Int) common.Hash {
	messagesAfterCount := new(big.Int).Sub(inboxTopCount, a.After.InboxIndex)
	return BisectionChunkHash(big.NewInt(0), messagesAfterCount, inboxTopHash, a.After.InboxHash)
}

func (a *AssertionInfo) InboxDeltaHash() common.Hash {
	return BisectionChunkHash(
		big.NewInt(0),
		a.InboxMessagesRead(),
		InboxDeltaHash(a.After.InboxHash, common.Hash{}),
		InboxDeltaHash(a.Before.InboxHash, a.InboxDelta),
	)
}

func (a *Assertion) BeforeExecutionHash() common.Hash {
	restBefore := assertionRestHash(
		a.InboxDelta,
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
		common.Hash{},
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
