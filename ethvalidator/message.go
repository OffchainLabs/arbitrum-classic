package ethvalidator

import "github.com/offchainlabs/arb-avm/protocol"

type sendAssertMessage struct {
	Precondition *protocol.Precondition
	Assertion    *protocol.Assertion
}

func (sendAssertMessage) IsOutgoingMessage() {}

type sendUnanimousAssertMessage struct {
	NewInboxHash [32]byte
	TimeBounds   protocol.TimeBounds
	Assertion    *protocol.Assertion
	Signatures   [][]byte
}

func (sendUnanimousAssertMessage) IsOutgoingMessage() {}

type sendProposeUnanimousAssertMessage struct {
	NewInboxHash [32]byte
	TimeBounds   protocol.TimeBounds
	Assertion    *protocol.Assertion
	SequenceNum  uint64
	Signatures   [][]byte
}

func (sendProposeUnanimousAssertMessage) IsOutgoingMessage() {}

type sendConfirmUnanimousAssertedMessage struct {
	NewInboxHash [32]byte
	Assertion    *protocol.Assertion
}

func (sendConfirmUnanimousAssertedMessage) IsOutgoingMessage() {}

type sendInitiateChallengeMessage struct {
	Precondition *protocol.Precondition
	Assertion    *protocol.AssertionStub
}

func (sendInitiateChallengeMessage) IsOutgoingMessage() {}

type sendBisectionMessage struct {
	Deadline     uint64
	Precondition *protocol.Precondition
	Assertions   []*protocol.Assertion
}

func (sendBisectionMessage) IsOutgoingMessage() {}

type sendContinueChallengeMessage struct {
	AssertionToChallenge uint16
	Deadline             uint64
	Preconditions        []*protocol.Precondition
	Assertions           []*protocol.AssertionStub
}

func (sendContinueChallengeMessage) IsOutgoingMessage() {}

type sendOneStepProofMessage struct {
	Precondition *protocol.Precondition
	Assertion    *protocol.Assertion
	Proof        []byte
	Deadline     uint64
}

func (sendOneStepProofMessage) IsOutgoingMessage() {}

type sendConfirmedAssertMessage struct {
	Precondition *protocol.Precondition
	Assertion    *protocol.Assertion
}

func (sendConfirmedAssertMessage) IsOutgoingMessage() {}

type sendAsserterTimedOutChallengeMessage struct {
	Deadline     uint64
	Precondition *protocol.Precondition
	Assertion    *protocol.AssertionStub
}

func (sendAsserterTimedOutChallengeMessage) IsOutgoingMessage() {}

type sendChallengerTimedOutChallengeMessage struct {
	Deadline      uint64
	Preconditions []*protocol.Precondition
	Assertions    []*protocol.AssertionStub
}

func (sendChallengerTimedOutChallengeMessage) IsOutgoingMessage() {}