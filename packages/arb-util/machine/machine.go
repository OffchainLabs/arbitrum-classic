package machine

import (
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Context interface {
	Send(message protocol.Message)
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

func (m *NoContext) Send(message protocol.Message) {

}

func (m *NoContext) OutMessageCount() int {
	return 0
}

func (m *NoContext) GetTimeBounds() value.Value {
	return value.NewEmptyTuple()
}

func (m *NoContext) NotifyStep() {
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

	CurrentStatus() Status
	LastBlockReason() BlockReason
	CanSpend(tokenType protocol.TokenType, currency *big.Int) bool
	InboxHash() value.HashOnlyValue
	PendingMessageCount() uint64
	SendOnchainMessage(protocol.Message)
	DeliverOnchainMessage()
	SendOffchainMessages([]protocol.Message)

	ExecuteAssertion(maxSteps int32, timeBounds protocol.TimeBounds) *protocol.Assertion
	MarshalForProof() ([]byte, error)
}

func IsMachineBlocked(machine Machine, currentTime uint64) bool {
	lastReason := machine.LastBlockReason()
	if lastReason == nil {
		return false
	}
	return lastReason.IsBlocked(machine, currentTime)
}
