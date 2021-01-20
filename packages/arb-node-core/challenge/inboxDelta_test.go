package challenge

import (
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"testing"
)

func TestInboxDeltaChallenge(t *testing.T) {
	mach, err := cmachine.New(arbos.Path())
	test.FailIfError(t, err)

	correctLookup := core.NewValidatorLookupMock(mach)
	for i := 0; i < 500; i++ {
		correctLookup.AddMessage(inbox.NewRandomInboxMessage())
	}
	otherLookup := core.NewValidatorLookupMock(mach)
	for i := 0; i < 500; i++ {
		otherLookup.AddMessage(inbox.NewRandomInboxMessage())
	}

	falseLookup := correctLookup.Clone()
	for i := 360; i < 370; i++ {
		falseLookup.Messages[i] = otherLookup.Messages[i]
	}

	inboxMessagesRead := big.NewInt(450)

	challengedNode := initializeChallengeData(t, correctLookup, inboxMessagesRead)

	messages, err := falseLookup.GetMessages(big.NewInt(0), inboxMessagesRead)
	test.FailIfError(t, err)
	challengedNode.Assertion.InboxDelta = core.CalculateInboxDeltaAcc(messages)

	asserterTime := big.NewInt(100000)
	challengerTime := big.NewInt(100000)

	rounds := executeChallenge(t, challengedNode, asserterTime, challengerTime, correctLookup, falseLookup)
	if rounds != 3 {
		t.Fatal("wrong round count", rounds)
	}
}
