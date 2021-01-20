package challenge

import (
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func TestInboxConsistencyChallenge(t *testing.T) {
	mach, err := cmachine.New(arbos.Path())
	test.FailIfError(t, err)

	correctLookup := core.NewValidatorLookupMock(mach)
	for i := 0; i < 10000; i++ {
		correctLookup.AddMessage(inbox.NewRandomInboxMessage())
	}
	otherLookup := core.NewValidatorLookupMock(mach)
	for i := 0; i < 10000; i++ {
		otherLookup.AddMessage(inbox.NewRandomInboxMessage())
	}

	falseLookup := correctLookup.Clone()
	for i := 200; i < 206; i++ {
		falseLookup.InboxAccs[i] = otherLookup.InboxAccs[i]
	}

	inboxMessagesRead := big.NewInt(203)

	challengedNode := initializeChallengeData(t, correctLookup, inboxMessagesRead)

	inboxAcc, err := falseLookup.GetInboxAcc(new(big.Int).Add(challengedNode.Assertion.After.InboxIndex, big.NewInt(1)))
	test.FailIfError(t, err)
	challengedNode.Assertion.After.InboxHash = inboxAcc

	asserterTime := big.NewInt(100000)
	challengerTime := big.NewInt(100000)

	rounds := executeChallenge(t, challengedNode, asserterTime, challengerTime, correctLookup, falseLookup)
	if rounds != 3 {
		t.Fatal("wrong round count", rounds)
	}
}
