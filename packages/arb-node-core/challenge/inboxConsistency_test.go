package challenge

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"math/big"
	"os"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

var tmpDir = "./tmp"

type InvalidInboxArbCore struct {
	core.ArbCore
	fakeInboxAccs map[uint64]common.Hash
}

func NewInvalidInboxArbCore(realCore core.ArbCore) *InvalidInboxArbCore {
	fakeInboxAccs := make(map[uint64]common.Hash)
	for i := uint64(200); i < 206; i++ {
		fakeInboxAccs[i] = common.RandHash()
	}
	return &InvalidInboxArbCore{
		ArbCore:       realCore,
		fakeInboxAccs: fakeInboxAccs,
	}
}

func (c *InvalidInboxArbCore) GetInboxAcc(index *big.Int) (common.Hash, error) {
	h, ok := c.fakeInboxAccs[index.Uint64()]
	if ok {
		return h, nil
	}
	return c.ArbCore.GetInboxAcc(index)
}

func generateRandomValidMessages(msgCount int) []inbox.InboxMessage {
	messages := make([]inbox.InboxMessage, 0)
	for i := 0; i < msgCount; i++ {
		msg := inbox.InboxMessage{
			Kind:        inbox.Type(2),
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
	return messages
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

	_, err = core.DeliverMessagesAndWait(correctLookup, generateRandomValidMessages(10000), common.Hash{}, false)
	test.FailIfError(t, err)

	falseLookup := NewInvalidInboxArbCore(correctLookup)

	inboxMessagesRead := big.NewInt(203)

	challengedNode := initializeChallengeData(t, correctLookup, inboxMessagesRead)

	inboxAcc, err := falseLookup.GetInboxAcc(new(big.Int).Add(challengedNode.Assertion.After.TotalMessagesRead, big.NewInt(1)))
	test.FailIfError(t, err)
	challengedNode.Assertion.After.InboxHash = inboxAcc

	asserterTime := big.NewInt(100000)
	challengerTime := big.NewInt(100000)

	rounds := executeChallenge(t, challengedNode, asserterTime, challengerTime, correctLookup, falseLookup)
	if rounds != 3 {
		t.Fatal("wrong round count", rounds)
	}
}

func TestInboxHashing(t *testing.T) {
	messages := generateRandomValidMessages(100)
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

	lookup := storage.GetArbCore()
	started := lookup.StartThread()
	if !started {
		t.Fatal("failed to start thread")
	}

	_, err = core.DeliverMessagesAndWait(lookup, messages, common.Hash{}, false)
	test.FailIfError(t, err)

	acc := common.Hash{}
	for i, msg := range messages {
		acc = hashing.SoliditySHA3(hashing.Bytes32(acc), hashing.Bytes32(msg.CommitmentHash()))
		dbAcc, err := lookup.GetInboxAcc(big.NewInt(int64(i)))
		test.FailIfError(t, err)
		if acc != dbAcc {
			t.Fatal("bad acc", i, new(big.Int).SetBytes(acc.Bytes()), new(big.Int).SetBytes(dbAcc.Bytes()))
		}

		msgs, err := lookup.GetMessages(big.NewInt(int64(i)), big.NewInt(1))
		test.FailIfError(t, err)
		if !msg.Equals(msgs[0]) {
			t.Log(msg)
			t.Log(msgs[0])
			t.Fatal("unequal messages")
		}
	}

	msgHashes, err := lookup.GetMessageHashes(big.NewInt(0), big.NewInt(int64(len(messages))))
	test.FailIfError(t, err)
	if len(msgHashes) != len(messages) {
		t.Fatal("wrong msg hash count", len(msgHashes), len(messages))
	}

	delta := common.Hash{}
	for i := range msgHashes {
		delta = hashing.SoliditySHA3(hashing.Bytes32(delta), hashing.Bytes32(msgHashes[len(msgHashes)-1-i]))
	}

	dbDelta, err := lookup.GetInboxDelta(big.NewInt(0), big.NewInt(int64(len(messages))))
	test.FailIfError(t, err)

	if delta != dbDelta {
		t.Log(delta)
		t.Log(dbDelta)
		t.Fatal("wrong delta")
	}
}
