package core

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/pkg/errors"
	"math/big"
)

type ValidatorLookupMock struct {
	Messages  []inbox.InboxMessage
	InboxAccs []common.Hash

	startMachine machine.Machine
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

func (v *ValidatorLookupMock) AddMessage(msg inbox.InboxMessage) {
	prevInboxAcc := v.InboxAccs[len(v.InboxAccs)-1]
	newInboxAcc := hashing.SoliditySHA3(hashing.Bytes32(prevInboxAcc), hashing.Bytes32(msg.CommitmentHash()))
	v.Messages = append(v.Messages, msg)
	v.InboxAccs = append(v.InboxAccs, newInboxAcc)
}

func (v *ValidatorLookupMock) GetSends(startIndex *big.Int, count *big.Int) ([][]byte, error) {
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

func (v *ValidatorLookupMock) GetInboxDelta(startIndex *big.Int, count *big.Int) (common.Hash, error) {
	panic("implement me")
}

func (v *ValidatorLookupMock) GetInboxAcc(index *big.Int) (common.Hash, error) {
	i := index.Uint64()
	if i >= uint64(len(v.InboxAccs)) {
		return common.Hash{}, errors.New("GetInboxAcc: inbox index out of bounds")
	}
	return v.InboxAccs[i], nil
}

func (v *ValidatorLookupMock) GetSendAcc(startAcc common.Hash, startIndex *big.Int, count *big.Int) (common.Hash, error) {
	panic("implement me")
}

func (v *ValidatorLookupMock) GetLogAcc(startAcc common.Hash, startIndex *big.Int, count *big.Int) (common.Hash, error) {
	panic("implement me")
}

func (v *ValidatorLookupMock) GetCursor(totalGasUsed *big.Int) (ExecutionCursor, error) {
	panic("implement me")
}

func (v *ValidatorLookupMock) MoveExecutionCursor(start ExecutionCursor, maxGas *big.Int, goOverGas bool) error {
	panic("implement me")
}

func (v *ValidatorLookupMock) GetMachine(cursor ExecutionCursor) (machine.Machine, error) {
	panic("implement me")
}
