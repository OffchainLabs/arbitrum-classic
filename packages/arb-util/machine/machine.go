package machine

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Context interface {
	Send(message value.Value)
	GetTimeBounds() value.Value
	NotifyStep(uint64)
	LoggedValue(value.Value)

	OutMessageCount() int
}

type NoContext struct{}

func (m *NoContext) LoggedValue(data value.Value) {

}

func (m *NoContext) Send(message value.Value) {

}

func (m *NoContext) OutMessageCount() int {
	return 0
}

func (m *NoContext) GetTimeBounds() value.Value {
	return value.NewEmptyTuple()
}

func (m *NoContext) NotifyStep(_ uint64) {
}

type Status int

const (
	Extensive Status = iota
	ErrorStop
	Halt
)

type Machine interface {
	Hash() [32]byte
	Clone() Machine
	PrintState()

	CurrentStatus() Status
	LastBlockReason() BlockReason
	InboxHash() value.HashOnlyValue
	PendingMessageCount() uint64
	SendOnchainMessage(protocol.Message)
	DeliverOnchainMessage()
	SendOffchainMessages([]protocol.Message)

	ExecuteAssertion(maxSteps int32, timeBounds *protocol.TimeBounds) *protocol.Assertion
	MarshalForProof() ([]byte, error)

	Checkpoint(storage CheckpointStorage) bool
	RestoreCheckpoint(storage CheckpointStorage, checkpointName string) bool
}

func IsMachineBlocked(machine Machine, currentTime uint64) bool {
	lastReason := machine.LastBlockReason()
	if lastReason == nil {
		return false
	}
	return lastReason.IsBlocked(machine, currentTime)
}
