package ethbridge

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"math/big"
)

type Assertion struct {
	BeforeProposedBlock *big.Int
	BeforeStepsRun      *big.Int
	BeforeMachineHash   common.Hash
	BeforeInboxHash     common.Hash
	BeforeInboxCount    *big.Int
	BeforeSendCount     *big.Int
	BeforeLogCount      *big.Int
	BeforeInboxMaxCount *big.Int
	StepsExecuted       *big.Int
	InboxDelta          common.Hash
	InboxMessagesRead   *big.Int
	GasUsed             *big.Int
	SendAcc             common.Hash
	SendCount           *big.Int
	LogAcc              common.Hash
	LogCount            *big.Int
	AfterInboxHash      common.Hash
	AfterMachineHash    common.Hash
}

func NewAssertionFromFields(a [7][32]byte, b [11]*big.Int) *Assertion {
	return &Assertion{
		BeforeProposedBlock: b[0],
		BeforeStepsRun:      b[1],
		BeforeMachineHash:   a[0],
		BeforeInboxHash:     a[1],
		BeforeInboxCount:    b[2],
		BeforeSendCount:     b[3],
		BeforeLogCount:      b[4],
		BeforeInboxMaxCount: b[5],
		StepsExecuted:       b[6],
		InboxDelta:          a[2],
		InboxMessagesRead:   b[7],
		GasUsed:             b[8],
		SendAcc:             a[3],
		SendCount:           b[9],
		LogAcc:              a[4],
		LogCount:            b[10],
		AfterInboxHash:      a[5],
		AfterMachineHash:    a[6],
	}
}

func (a *Assertion) BytesFields() [7][32]byte {
	return [7][32]byte{
		a.BeforeMachineHash,
		a.BeforeInboxHash,
		a.InboxDelta,
		a.SendAcc,
		a.LogAcc,
		a.AfterInboxHash,
		a.AfterMachineHash,
	}
}

func (a *Assertion) IntFields() [11]*big.Int {
	return [11]*big.Int{
		a.BeforeProposedBlock,
		a.BeforeStepsRun,
		a.BeforeInboxCount,
		a.BeforeSendCount,
		a.BeforeLogCount,
		a.BeforeInboxMaxCount,
		a.StepsExecuted,
		a.InboxMessagesRead,
		a.GasUsed,
		a.SendCount,
		a.LogCount,
	}
}

func (a *Assertion) AfterInboxCount() *big.Int {
	return new(big.Int).Add(a.BeforeInboxCount, a.InboxMessagesRead)
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
	messagesAfterCount := new(big.Int).Sub(inboxTopCount, a.BeforeInboxCount)
	messagesAfterCount = messagesAfterCount.Sub(messagesAfterCount, a.InboxMessagesRead)
	return bisectionChunkHash(messagesAfterCount, messagesAfterCount, inboxTopHash, a.AfterInboxHash)
}

func (a *Assertion) InboxDeltaHash() common.Hash {
	return bisectionChunkHash(
		a.InboxMessagesRead,
		a.InboxMessagesRead,
		inboxDeltaHash(a.AfterInboxHash, common.Hash{}),
		inboxDeltaHash(a.BeforeInboxHash, a.InboxDelta),
	)
}

func (a *Assertion) ExecutionHash() common.Hash {
	return bisectionChunkHash(
		a.GasUsed,
		a.GasUsed,
		assertionHash(
			a.InboxDelta,
			big.NewInt(0),
			outputAccHash(common.Hash{}, big.NewInt(0), common.Hash{}, big.NewInt(0)),
			a.BeforeMachineHash,
		),
		assertionHash(
			common.Hash{},
			a.GasUsed,
			outputAccHash(
				a.SendAcc,
				a.SendCount,
				a.LogAcc,
				a.LogCount,
			),
			a.AfterMachineHash,
		),
	)
}

func (a *Assertion) CheckTime(arbGasSpeedLimitPerBlock *big.Int) *big.Int {
	return new(big.Int).Div(a.GasUsed, arbGasSpeedLimitPerBlock)
}
