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
	messages  []inbox.InboxMessage
	inboxAccs []common.Hash

	startMachine machine.Machine
}

func NewValidatorLookupMock(mach machine.Machine) *ValidatorLookupMock {
	return &ValidatorLookupMock{
		inboxAccs:    []common.Hash{{}},
		startMachine: mach.Clone(),
	}
}

func (v *ValidatorLookupMock) AddMessage(msg inbox.InboxMessage) {
	prevInboxAcc := v.inboxAccs[len(v.inboxAccs)-1]
	newInboxAcc := hashing.SoliditySHA3(hashing.Bytes32(prevInboxAcc), hashing.Bytes32(msg.CommitmentHash()))
	v.messages = append(v.messages, msg)
	v.inboxAccs = append(v.inboxAccs, newInboxAcc)
}

func (v *ValidatorLookupMock) GenerateLogAccumulator(startIndex *big.Int, count *big.Int) (common.Hash, error) {
	panic("implement me")
}

func (v *ValidatorLookupMock) GetSends(startIndex *big.Int, count *big.Int) ([][]byte, error) {
	panic("implement me")
}

func (v *ValidatorLookupMock) GetInboxAcc(index *big.Int) (common.Hash, error) {
	i := index.Uint64()
	if i >= uint64(len(v.inboxAccs)) {
		return common.Hash{}, errors.New("inbox index out of bounds")
	}
	return v.inboxAccs[i], nil
}

func (v *ValidatorLookupMock) GetMessages(startIndex *big.Int, count *big.Int) ([]inbox.InboxMessage, error) {
	if count.Cmp(big.NewInt(0)) == 0 {
		return nil, nil
	}
	start := startIndex.Uint64()
	c := count.Uint64()
	if start+c >= uint64(len(v.messages)) {
		return nil, errors.New("inbox index out of bounds")
	}
	return v.messages[start : start+c], nil
}

func (v *ValidatorLookupMock) GetMachine(totalGasUsed *big.Int) (machine.Machine, error) {
	if totalGasUsed.Cmp(big.NewInt(0)) == 0 {
		return v.startMachine, nil
	}
	return nil, errors.New("GetMachine not yet supported")
}

func (v *ValidatorLookupMock) GetExecutionInfoWithMaxMessages(startMachine machine.Machine, targetGas *big.Int, maxMessages *big.Int) (*ExecutionInfo, error) {
	if targetGas.Cmp(big.NewInt(0)) == 0 {
		return &ExecutionInfo{
			BeforeMachineHash: startMachine.Hash(),
			InboxMessagesRead: big.NewInt(0),
			GasUsed:           big.NewInt(0),
			SendAcc:           common.Hash{},
			SendCount:         big.NewInt(0),
			LogAcc:            common.Hash{},
			LogCount:          big.NewInt(0),
			AfterMachineHash:  startMachine.Hash(),
		}, nil
	}
	return nil, errors.New("GetExecutionInfoWithMaxMessages not yet supported")
}

func (v *ValidatorLookupMock) GetExecutionInfo(startMachine machine.Machine, maxGas *big.Int) (*AssertionInfo, error) {
	panic("implement me")
}
