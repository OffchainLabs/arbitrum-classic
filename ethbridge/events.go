package ethbridge

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arb-avm/protocol"
	"math/big"
)

type Event interface {

}

type Notification struct {
	Header *types.Header
	VmID   [32]byte
	Event  Event
}

type FinalUnanimousAssertEvent struct {
	UnanHash [32]byte
}

type ProposedUnanimousAssertEvent struct {
	UnanHash    [32]byte
	SequenceNum uint64
}

type ConfirmedUnanimousAssertEvent struct {
	SequenceNum uint64
}

type DisputableAssertionEvent struct {
	Precondition *protocol.Precondition
	Assertion    *protocol.AssertionStub
	Asserter     common.Address
}

type ConfirmedAssertEvent struct {}

type InitiateChallengeEvent struct {
	Challenger common.Address
}

type BisectionEvent struct {
	Assertions []*protocol.AssertionStub
}

type ContinueChallengeEvent struct {
	ChallengedAssertion uint16
}

type ChallengerTimeoutEvent struct {}

type AsserterTimeoutEvent struct {}

type OneStepProofEvent struct {}

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

type MessageDeliveredEvent struct {
	Msg protocol.Message
}

type NewTimeEvent struct {}