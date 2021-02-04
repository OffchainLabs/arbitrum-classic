package challenge

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
	"math/rand"
	"os"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

var tmpDir = "./tmp"

type InvalidArbCore struct {
	core.ArbCore
}

func TestInboxConsistencyChallenge(t *testing.T) {
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			panic(err)
		}
	}()
	storage, err := cmachine.NewArbStorage(tmpDir)
	test.FailIfError(t, err)
	defer storage.CloseArbStorage()

	err = storage.Initialize(arbos.Path())
	test.FailIfError(t, err)

	correctLookup := storage.GetArbCore()
	started := correctLookup.StartThread()
	if !started {
		t.Fatal("failed to start thread")
	}
	messages := make([]inbox.InboxMessage, 0)
	for i := 0; i < 10000; i++ {
		msg := inbox.InboxMessage{
			Kind:        inbox.Type(rand.Uint32()),
			Sender:      common.RandAddress(),
			InboxSeqNum: big.NewInt(int64(i)),
			Data:        common.RandBytes(200),
			ChainTime: inbox.ChainTime{
				BlockNum:  common.NewTimeBlocksInt(int64(i)),
				Timestamp: big.NewInt(int64(i)),
			},
		}
		messages = append(messages, msg)
	}
	_, err = core.DeliverMessagesAndWait(correctLookup, messages, common.Hash{}, false)
	test.FailIfError(t, err)

	falseLookup := InvalidArbCore{ArbCore: correctLookup}

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
