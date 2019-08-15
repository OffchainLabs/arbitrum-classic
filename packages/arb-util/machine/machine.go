package machine

import (
	"errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Context interface {
	Send(data value.Value, tokenType value.IntValue, currency value.IntValue, dest value.IntValue) error
	GetTimeBounds() value.Value
	NotifyStep()
	LoggedValue(value.Value)

	OutMessageCount() int
}

type NoContext struct{}

func (m *NoContext) LoggedValue(data value.Value) {

}

func (m *NoContext) CanSpend(tokenType value.IntValue, currency value.IntValue) bool {
	return false
}

func (m *NoContext) Send(data value.Value, tokenType value.IntValue, currency value.IntValue, dest value.IntValue) error {
	return errors.New("can't send message outside of assertion mode")
}

func (m *NoContext) OutMessageCount() int {
	return 0
}

func (m *NoContext) GetTimeBounds() value.Value {
	return value.NewEmptyTuple()
}

func (m *NoContext) NotifyStep() {
}

type Machine interface {
	Hash() [32]byte
	Clone() Machine

	InboxHash() value.HashOnlyValue
	PendingMessageCount() uint64
	SendOnchainMessage(protocol.Message)
	DeliverOnchainMessage()
	SendOffchainMessages([]protocol.Message)

	ExecuteAssertion(maxSteps int32, timeBounds protocol.TimeBounds) *protocol.Assertion
	MarshalForProof() ([]byte, error)
}
