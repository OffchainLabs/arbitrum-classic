package core

import (
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
)

type NodeState struct {
	ProposedBlock  *big.Int
	TotalGasUsed   *big.Int
	MachineHash    common.Hash
	InboxHash      common.Hash
	InboxCount     *big.Int
	TotalSendCount *big.Int
	TotalLogCount  *big.Int
	InboxMaxCount  *big.Int
}

type Assertion struct {
	PrevState *NodeState
	*AssertionInfo
}

func NewAssertionFromFields(a [7][32]byte, b [10]*big.Int) *Assertion {
	prevState := &NodeState{
		ProposedBlock:  b[0],
		TotalGasUsed:   b[1],
		MachineHash:    a[0],
		InboxHash:      a[1],
		InboxCount:     b[2],
		TotalSendCount: b[3],
		TotalLogCount:  b[4],
		InboxMaxCount:  b[5],
	}
	return &Assertion{
		PrevState: prevState,
		AssertionInfo: &AssertionInfo{
			InboxDelta: a[2],
			ExecInfo: &ExecutionInfo{
				BeforeMachineHash: prevState.MachineHash,
				InboxMessagesRead: b[6],
				GasUsed:           b[7],
				SendAcc:           a[4],
				SendCount:         b[8],
				LogAcc:            a[5],
				LogCount:          b[9],
				AfterMachineHash:  a[6],
			},
			AfterInboxHash: a[3],
		},
	}
}

func (a *Assertion) BytesFields() [7][32]byte {
	return [7][32]byte{
		a.PrevState.MachineHash,
		a.PrevState.InboxHash,
		a.InboxDelta,
		a.ExecInfo.SendAcc,
		a.ExecInfo.LogAcc,
		a.AfterInboxHash,
		a.ExecInfo.AfterMachineHash,
	}
}

func (a *Assertion) IntFields() [10]*big.Int {
	return [10]*big.Int{
		a.PrevState.ProposedBlock,
		a.PrevState.TotalGasUsed,
		a.PrevState.InboxCount,
		a.PrevState.TotalSendCount,
		a.PrevState.TotalLogCount,
		a.PrevState.InboxMaxCount,
		a.ExecInfo.InboxMessagesRead,
		a.ExecInfo.GasUsed,
		a.ExecInfo.SendCount,
		a.ExecInfo.LogCount,
	}
}

func (a *Assertion) AfterInboxCount() *big.Int {
	return new(big.Int).Add(a.PrevState.InboxCount, a.ExecInfo.InboxMessagesRead)
}

func (a *Assertion) AfterTotalGasUsed() *big.Int {
	return new(big.Int).Add(a.PrevState.TotalGasUsed, a.ExecInfo.GasUsed)
}

func (a *Assertion) AfterTotalSendCount() *big.Int {
	return new(big.Int).Add(a.PrevState.TotalSendCount, a.ExecInfo.SendCount)
}

func (a *Assertion) AfterTotalLogCount() *big.Int {
	return new(big.Int).Add(a.PrevState.TotalLogCount, a.ExecInfo.LogCount)
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
	messagesAfterCount := new(big.Int).Sub(inboxTopCount, a.PrevState.InboxCount)
	messagesAfterCount = messagesAfterCount.Sub(messagesAfterCount, a.ExecInfo.InboxMessagesRead)
	return BisectionChunkHash(messagesAfterCount, messagesAfterCount, inboxTopHash, a.AfterInboxHash)
}

func (a *Assertion) InboxDeltaHash() common.Hash {
	return BisectionChunkHash(
		a.ExecInfo.InboxMessagesRead,
		a.ExecInfo.InboxMessagesRead,
		InboxDeltaHash(a.AfterInboxHash, common.Hash{}),
		InboxDeltaHash(a.PrevState.InboxHash, a.InboxDelta),
	)
}

func (a *Assertion) BeforeExecutionHash() common.Hash {
	restBefore := assertionRestHash(
		a.InboxDelta,
		a.PrevState.MachineHash,
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
		a.ExecInfo.AfterMachineHash,
		a.ExecInfo.SendAcc,
		a.ExecInfo.SendCount,
		a.ExecInfo.LogAcc,
		a.ExecInfo.LogCount,
	)
	return assertionHash(a.ExecInfo.GasUsed, restAfter)
}

func (a *Assertion) ExecutionHash() common.Hash {
	return BisectionChunkHash(
		a.ExecInfo.GasUsed,
		a.ExecInfo.GasUsed,
		a.BeforeExecutionHash(),
		a.AfterExecutionHash(),
	)
}

func (a *Assertion) CheckTime(arbGasSpeedLimitPerBlock *big.Int) *big.Int {
	return new(big.Int).Div(a.ExecInfo.GasUsed, arbGasSpeedLimitPerBlock)
}
