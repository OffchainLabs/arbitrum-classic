package ethbridge

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"math/big"
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

type ExecutionInfo struct {
	BeforeMachineHash common.Hash
	GasUsed           *big.Int
	SendAcc           common.Hash
	SendCount         *big.Int
	LogAcc            common.Hash
	LogCount          *big.Int
	AfterMachineHash  common.Hash
}

func (e *ExecutionInfo) Equals(o *ExecutionInfo) bool {
	return e.BeforeMachineHash == o.BeforeMachineHash &&
		e.GasUsed.Cmp(o.GasUsed) == 0 &&
		e.SendAcc == o.SendAcc &&
		e.SendCount.Cmp(o.SendCount) == 0 &&
		e.LogAcc == o.LogAcc &&
		e.LogCount.Cmp(o.LogCount) == 0 &&
		e.AfterMachineHash == o.AfterMachineHash
}

type Assertion struct {
	PrevState *NodeState
	*AssertionInfo
}

type AssertionInfo struct {
	InboxDelta        common.Hash
	InboxMessagesRead *big.Int
	ExecInfo          *ExecutionInfo
	AfterInboxHash    common.Hash
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
			InboxDelta:        a[2],
			InboxMessagesRead: b[6],
			ExecInfo: &ExecutionInfo{
				BeforeMachineHash: prevState.MachineHash,
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
		a.InboxMessagesRead,
		a.ExecInfo.GasUsed,
		a.ExecInfo.SendCount,
		a.ExecInfo.LogCount,
	}
}

func (a *Assertion) AfterInboxCount() *big.Int {
	return new(big.Int).Add(a.PrevState.InboxCount, a.InboxMessagesRead)
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

func bisectionChunkHash(
	length *big.Int,
	segmentEnd *big.Int,
	startHash common.Hash,
	endHash common.Hash,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint256(length),
		hashing.Uint256(segmentEnd),
		hashing.Bytes32(startHash),
		hashing.Bytes32(endHash),
	)
}

func inboxDeltaHash(inboxAcc, deltaAcc common.Hash) common.Hash {
	return hashing.SoliditySHA3(hashing.Bytes32(inboxAcc), hashing.Bytes32(deltaAcc))
}

func assertionHash(
	inboxDelta common.Hash,
	gasUsed *big.Int,
	outputAcc common.Hash,
	machineState common.Hash,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(inboxDelta),
		hashing.Uint256(gasUsed),
		hashing.Bytes32(outputAcc),
		hashing.Bytes32(machineState),
	)
}

func outputAccHash(
	sendAcc common.Hash,
	sendCount *big.Int,
	logAcc common.Hash,
	logCount *big.Int,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(sendAcc),
		hashing.Uint256(sendCount),
		hashing.Bytes32(logAcc),
		hashing.Uint256(logCount),
	)
}

func (a *Assertion) InboxConsistencyHash(inboxTopHash common.Hash, inboxTopCount *big.Int) common.Hash {
	messagesAfterCount := new(big.Int).Sub(inboxTopCount, a.PrevState.InboxCount)
	messagesAfterCount = messagesAfterCount.Sub(messagesAfterCount, a.InboxMessagesRead)
	return bisectionChunkHash(messagesAfterCount, messagesAfterCount, inboxTopHash, a.AfterInboxHash)
}

func (a *Assertion) InboxDeltaHash() common.Hash {
	return bisectionChunkHash(
		a.InboxMessagesRead,
		a.InboxMessagesRead,
		inboxDeltaHash(a.AfterInboxHash, common.Hash{}),
		inboxDeltaHash(a.PrevState.InboxHash, a.InboxDelta),
	)
}

func (a *Assertion) ExecutionHash() common.Hash {
	return bisectionChunkHash(
		a.ExecInfo.GasUsed,
		a.ExecInfo.GasUsed,
		assertionHash(
			a.InboxDelta,
			big.NewInt(0),
			outputAccHash(common.Hash{}, big.NewInt(0), common.Hash{}, big.NewInt(0)),
			a.PrevState.MachineHash,
		),
		assertionHash(
			common.Hash{},
			a.ExecInfo.GasUsed,
			outputAccHash(
				a.ExecInfo.SendAcc,
				a.ExecInfo.SendCount,
				a.ExecInfo.LogAcc,
				a.ExecInfo.LogCount,
			),
			a.ExecInfo.AfterMachineHash,
		),
	)
}

func (a *Assertion) CheckTime(arbGasSpeedLimitPerBlock *big.Int) *big.Int {
	return new(big.Int).Div(a.ExecInfo.GasUsed, arbGasSpeedLimitPerBlock)
}
