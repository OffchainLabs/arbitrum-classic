package bridge

import (
	"github.com/offchainlabs/arb-avm/protocol"
)

type Bridge interface {
	FinalizedAssertion(assertion *protocol.Assertion, newLogCount int)
	FinalUnanimousAssert(newInboxHash [32]byte, timeBounds protocol.TimeBounds, assertion *protocol.Assertion, signatures [][]byte)
	UnanimousAssert(newInboxHash [32]byte, timeBounds protocol.TimeBounds, assertion *protocol.Assertion, sequenceNum uint64, signatures [][]byte)
	ConfirmUnanimousAssertion(newInboxHash [32]byte, assertion *protocol.Assertion)
	DisputableAssert(precondition *protocol.Precondition, assertion *protocol.Assertion)
	ConfirmDisputableAssertion(precondition *protocol.Precondition, assertion *protocol.Assertion)
	InitiateChallenge(precondition *protocol.Precondition, assertion *protocol.AssertionStub)
	BisectAssertion(precondition *protocol.Precondition, assertions []*protocol.Assertion, deadline uint64)
	ContinueChallenge(assertionToChallenge uint16, preconditions []*protocol.Precondition, assertions []*protocol.AssertionStub, deadline uint64)
	OneStepProof(precondition *protocol.Precondition, assertion *protocol.Assertion, proof []byte, deadline uint64)
	TimeoutAsserter(precondition *protocol.Precondition, assertion *protocol.AssertionStub, deadline uint64)
	TimeoutChallenger(preconditions []*protocol.Precondition, assertions []*protocol.AssertionStub, deadline uint64)
}
