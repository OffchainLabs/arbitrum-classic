package machine

import (
	"errors"
	"github.com/offchainlabs/arb-util/protocol"
	"github.com/offchainlabs/arb-util/value"
)

type MachineContext interface {
	Send(data value.Value, tokenType value.IntValue, currency value.IntValue, dest value.IntValue) error
	GetTimeBounds() value.Value
	NotifyStep()
	LoggedValue(value.Value) error

	OutMessageCount() int
}

type MachineNoContext struct{}

func (m *MachineNoContext) LoggedValue(data value.Value) error {
	return errors.New("can't log values outside of assertion mode")
}

func (m *MachineNoContext) CanSpend(tokenType value.IntValue, currency value.IntValue) bool {
	return false
}

func (m *MachineNoContext) Send(data value.Value, tokenType value.IntValue, currency value.IntValue, dest value.IntValue) error {
	return errors.New("can't send message outside of assertion mode")
}

func (m *MachineNoContext) OutMessageCount() int {
	return 0
}

func (m *MachineNoContext) GetTimeBounds() value.Value {
	return value.NewEmptyTuple()
}

func (m *MachineNoContext) NotifyStep() {
}

type Machine interface {
	Hash() [32]byte
	Clone() Machine

	InboxHash() value.HashOnlyValue
	HasPendingMessages() bool
	SendOnchainMessage(protocol.Message)
	DeliverOnchainMessage()
	SendOffchainMessages([]protocol.Message)

	ExecuteAssertion(maxSteps int32, timeBounds protocol.TimeBounds) *protocol.Assertion
	MarshalForProof() ([]byte, error)
}
