package machine

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Context interface {
	Send(message value.Value)
	GetTimeBounds() value.TupleValue
	NotifyStep(uint64)
	LoggedValue(value.Value)
	GetInbox() value.TupleValue
	ReadInbox()

	OutMessageCount() int
}

type NoContext struct{}

func (m *NoContext) LoggedValue(data value.Value) {

}

func (m *NoContext) GetInbox() value.TupleValue {
	return value.NewEmptyTuple()
}

func (m *NoContext) ReadInbox() {

}

func (m *NoContext) Send(message value.Value) {

}

func (m *NoContext) OutMessageCount() int {
	return 0
}

func (m *NoContext) GetTimeBounds() value.TupleValue {
	return value.NewEmptyTuple()
}

func (m *NoContext) NotifyStep(uint64) {
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

	ExecuteAssertion(maxSteps uint32, timeBounds *protocol.TimeBoundsBlocks, inbox value.TupleValue) (*protocol.ExecutionAssertion, uint32)
	MarshalForProof() ([]byte, error)

	Checkpoint(storage CheckpointStorage) bool
	RestoreCheckpoint(storage CheckpointStorage, checkpointName string) bool
}

func IsMachineBlocked(machine Machine, currentTime *protocol.TimeBlocks) bool {
	lastReason := machine.LastBlockReason()
	if lastReason == nil {
		return false
	}
	return lastReason.IsBlocked(machine, currentTime)
}
