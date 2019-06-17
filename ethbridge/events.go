package ethbridge

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arb-avm/protocol"
	"math/big"
)


type IncomingMessageType int

const (
	CommonMessage IncomingMessageType = iota
	ChallengeMessage
)

type Event interface {

}

type VMEvent interface {
	GetIncomingMessageType() IncomingMessageType
}

type Notification struct {
	Header *types.Header
	VmID   [32]byte
	Event  Event
}

type FinalUnanimousAssertEvent struct {
	UnanHash [32]byte
}

func (FinalUnanimousAssertEvent) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type ProposedUnanimousAssertEvent struct {
	UnanHash    [32]byte
	SequenceNum uint64
}

func (ProposedUnanimousAssertEvent) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type ConfirmedUnanimousAssertEvent struct {
	SequenceNum uint64
}

func (ConfirmedUnanimousAssertEvent) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type DisputableAssertionEvent struct {
	Precondition *protocol.Precondition
	Assertion    *protocol.AssertionStub
	Asserter     common.Address
}

func (DisputableAssertionEvent) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type ConfirmedAssertEvent struct {}

func (ConfirmedAssertEvent) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type InitiateChallengeEvent struct {
	Challenger common.Address
}

func (InitiateChallengeEvent) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type BisectionEvent struct {
	Assertions []*protocol.AssertionStub
}

func (BisectionEvent) GetIncomingMessageType() IncomingMessageType {
	return ChallengeMessage
}

type ContinueChallengeEvent struct {
	ChallengedAssertion uint16
}

func (ContinueChallengeEvent) GetIncomingMessageType() IncomingMessageType {
	return ChallengeMessage
}

type ChallengerTimeoutEvent struct {}

func (ChallengerTimeoutEvent) GetIncomingMessageType() IncomingMessageType {
	return ChallengeMessage
}

type AsserterTimeoutEvent struct {}

func (AsserterTimeoutEvent) GetIncomingMessageType() IncomingMessageType {
	return ChallengeMessage
}

type OneStepProofEvent struct {}

func (OneStepProofEvent) GetIncomingMessageType() IncomingMessageType {
	return ChallengeMessage
}

type VMCreatedEvent struct {
	GracePeriod         uint32
	EscrowRequired      *big.Int
	EscrowCurrency      common.Address
	MaxExecutionSteps   uint32
	VmId                [32]byte
	VmState             [32]byte
	ChallengeManagerNum uint16
	Owner               common.Address
	Validators          []common.Address
}

func (VMCreatedEvent) GetIncomingMessageType() IncomingMessageType {
	return CommonMessage
}

type MessageDeliveredEvent struct {
	Msg protocol.Message
}

type NewTimeEvent struct {}
