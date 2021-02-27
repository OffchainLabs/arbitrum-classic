package core

import (
	"math/big"

	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type ExecutionCursorMock struct {
	mach machine.Machine
}

func (e *ExecutionCursorMock) Clone() ExecutionCursor {
	return &ExecutionCursorMock{}
}

func (e *ExecutionCursorMock) MachineHash() common.Hash {
	return e.mach.Hash()
}

func (e *ExecutionCursorMock) TotalMessagesRead() *big.Int {
	return big.NewInt(0)
}

func (e *ExecutionCursorMock) InboxHash() common.Hash {
	return common.Hash{}
}

func (e *ExecutionCursorMock) TotalGasConsumed() *big.Int {
	return big.NewInt(0)
}

func (e *ExecutionCursorMock) TotalSendCount() *big.Int {
	return big.NewInt(0)
}

func (e *ExecutionCursorMock) TotalLogCount() *big.Int {
	return big.NewInt(0)
}

func (e *ExecutionCursorMock) TotalSteps() *big.Int {
	return big.NewInt(0)
}

func (e *ExecutionCursorMock) TakeMachine() (machine.Machine, error) {
	return e.mach, nil
}

type ValidatorLookupMock struct {
	Messages  []inbox.InboxMessage
	InboxAccs []common.Hash
	logs      []value.Value
	sends     [][]byte

	startMachine machine.Machine
}

func (v *ValidatorLookupMock) GetLogCount() (*big.Int, error) {
	return big.NewInt(int64(len(v.logs))), nil
}

func (v *ValidatorLookupMock) GetSendCount() (*big.Int, error) {
	return big.NewInt(int64(len(v.sends))), nil
}

func (v *ValidatorLookupMock) GetMessageCount() (*big.Int, error) {
	return big.NewInt(int64(len(v.Messages))), nil
}

func (v *ValidatorLookupMock) DeliverMessages(messages []inbox.InboxMessage, previousInboxHash common.Hash, lastBlockComplete bool) bool {
	panic("implement me")
}

func (v *ValidatorLookupMock) MessagesStatus() (MessageStatus, error) {
	return MessagesEmpty, nil
}

func (v *ValidatorLookupMock) StartThread() bool {
	panic("implement me")
}

func (v *ValidatorLookupMock) StopThread() {
	panic("implement me")
}

func NewValidatorLookupMock(mach machine.Machine) *ValidatorLookupMock {
	return &ValidatorLookupMock{
		InboxAccs:    []common.Hash{{}},
		startMachine: mach.Clone(),
	}
}

func (v *ValidatorLookupMock) Clone() *ValidatorLookupMock {
	messages := make([]inbox.InboxMessage, 0, len(v.Messages))
	for _, msg := range v.Messages {
		messages = append(messages, msg)
	}
	inboxAccs := make([]common.Hash, 0, len(v.InboxAccs))
	for _, inboxAcc := range v.InboxAccs {
		inboxAccs = append(inboxAccs, inboxAcc)
	}
	return &ValidatorLookupMock{
		Messages:     messages,
		InboxAccs:    inboxAccs,
		startMachine: v.startMachine.Clone(),
	}
}

func (v *ValidatorLookupMock) AdvanceExecutionCursor(cursor ExecutionCursor, maxGas *big.Int, goOverGas bool) error {
	panic("implement me")
}

func (v *ValidatorLookupMock) AddMessage(msg inbox.InboxMessage) {
	prevInboxAcc := v.InboxAccs[len(v.InboxAccs)-1]
	newInboxAcc := hashing.SoliditySHA3(hashing.Bytes32(prevInboxAcc), hashing.Bytes32(msg.CommitmentHash()))
	v.Messages = append(v.Messages, msg)
	v.InboxAccs = append(v.InboxAccs, newInboxAcc)
}

func (v *ValidatorLookupMock) GetSends(startIndex, count *big.Int) ([][]byte, error) {
	panic("implement me")
}

func (v *ValidatorLookupMock) GetLogs(startIndex, count *big.Int) ([]value.Value, error) {
	panic("implement me")
}

func (v *ValidatorLookupMock) GetMessages(startIndex *big.Int, count *big.Int) ([]inbox.InboxMessage, error) {
	if count.Cmp(big.NewInt(0)) == 0 {
		return nil, nil
	}
	start := startIndex.Uint64()
	c := count.Uint64()
	if start+c >= uint64(len(v.Messages)) {
		return nil, errors.Errorf("GetMessages: inbox index out of bounds (%v, %v)", startIndex, count)
	}
	return v.Messages[start : start+c], nil
}

func (v *ValidatorLookupMock) GetSendAcc(startAcc common.Hash, startIndex *big.Int, count *big.Int) (common.Hash, error) {
	panic("implement me")
}

func (v *ValidatorLookupMock) GetLogAcc(startAcc common.Hash, startIndex *big.Int, count *big.Int) (common.Hash, error) {
	panic("implement me")
}

func (v *ValidatorLookupMock) GetExecutionCursor(totalGasUsed *big.Int) (ExecutionCursor, error) {
	if totalGasUsed.Cmp(big.NewInt(0)) == 0 {
		return &ExecutionCursorMock{mach: v.startMachine}, nil
	}
	panic("implement me")
}

func (v *ValidatorLookupMock) GetMachineForSideload(blockNumber uint64) (machine.Machine, error) {
	panic("implement me")
}
