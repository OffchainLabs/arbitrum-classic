package challenge

import (
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"math/big"
	"os"
	"testing"
)

type InvalidDeltaArbCore struct {
	core.ArbCore
	fakeInboxAccs map[uint64]common.Hash
}

func NewInvalidDeltaArbCore(realCore core.ArbCore) *InvalidDeltaArbCore {
	fakeInboxAccs := make(map[uint64]common.Hash)
	for i := uint64(200); i < 206; i++ {
		fakeInboxAccs[i] = common.RandHash()
	}
	return &InvalidDeltaArbCore{
		ArbCore:       realCore,
		fakeInboxAccs: fakeInboxAccs,
	}
}

func (c *InvalidDeltaArbCore) GetInboxDelta(index, count *big.Int) (common.Hash, error) {
	//messageHashes, err := c.ArbCore.GetMessageHashes(index, count)
	//if err != nil {
	//	return common.Hash{}, err
	//}
	h, ok := c.fakeInboxAccs[index.Uint64()]
	if ok {
		return h, nil
	}
	return c.ArbCore.GetInboxAcc(index)
}

func TestInboxDeltaChallenge(t *testing.T) {
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

	_, err = core.DeliverMessagesAndWait(correctLookup, generateRandomValidMessages(500), common.Hash{}, false)
	test.FailIfError(t, err)

	falseLookup := NewInvalidInboxArbCore(correctLookup)

	inboxMessagesRead := big.NewInt(450)

	challengedNode := initializeChallengeData(t, correctLookup, inboxMessagesRead)

	inboxDeltaHash, err := falseLookup.GetInboxDelta(big.NewInt(0), inboxMessagesRead)
	test.FailIfError(t, err)
	challengedNode.Assertion.InboxDelta = inboxDeltaHash

	asserterTime := big.NewInt(100000)
	challengerTime := big.NewInt(100000)

	rounds := executeChallenge(t, challengedNode, asserterTime, challengerTime, correctLookup, falseLookup)
	if rounds != 3 {
		t.Fatal("wrong round count", rounds)
	}
}
