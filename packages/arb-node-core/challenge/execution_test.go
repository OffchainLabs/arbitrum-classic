package challenge

import (
	"io/ioutil"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

func prepareArbCore(t *testing.T, messages []inbox.InboxMessage) (core.ArbCore, func()) {
	tmpDir, err := ioutil.TempDir("", "arbitrum")
	test.FailIfError(t, err)
	storage, err := cmachine.NewArbStorage(tmpDir)
	if err != nil {
		os.RemoveAll(tmpDir)
	}
	test.FailIfError(t, err)
	shutdown := func() {
		storage.CloseArbStorage()
		if err := os.RemoveAll(tmpDir); err != nil {
			panic(err)
		}
	}
	returning := false
	defer (func() {
		if !returning {
			shutdown()
		}
	})()

	err = storage.Initialize(arbos.Path())
	test.FailIfError(t, err)

	arbCore := storage.GetArbCore()
	started := arbCore.StartThread()
	if !started {
		t.Fatal("failed to start thread")
	}

	if len(messages) > 0 {
		_, err = core.DeliverMessagesAndWait(arbCore, messages, common.Hash{}, false)
		test.FailIfError(t, err)
	}
	for {
		if arbCore.MachineIdle() {
			break
		}
		<-time.After(time.Millisecond * 200)
	}

	returning = true
	return arbCore, shutdown
}

func runExecutionTest(t *testing.T, messages []inbox.InboxMessage, startGas *big.Int, endGas *big.Int, faultConfig faultConfig, asserterMayFail bool) int {
	arbCore, shutdown := prepareArbCore(t, messages)
	defer shutdown()
	faultyCore := newFaultyCore(arbCore, faultConfig)

	challengedNode := initializeChallengeData(t, faultyCore, startGas, endGas)

	time := big.NewInt(100)
	return executeChallenge(
		t,
		challengedNode,
		time,
		time,
		arbCore,
		faultyCore,
		asserterMayFail,
	)
}

func TestChallengeToOSP(t *testing.T) {
	runExecutionTest(t, []inbox.InboxMessage{}, big.NewInt(0), big.NewInt(200000), faultConfig{distortMachineAtGas: big.NewInt(100000)}, false)
}

func makeInitMsg() inbox.InboxMessage {
	owner := common.RandAddress()
	chain := common.RandAddress()
	return message.NewInboxMessage(
		message.Init{
			ChainParams: protocol.ChainParams{
				StakeRequirement:        big.NewInt(0),
				StakeToken:              common.Address{},
				GracePeriod:             common.NewTimeBlocks(big.NewInt(3)),
				MaxExecutionSteps:       0,
				ArbGasSpeedLimitPerTick: 0,
			},
			Owner:       owner,
			ExtraConfig: []byte{},
		},
		chain,
		big.NewInt(0),
		inbox.ChainTime{
			BlockNum:  common.NewTimeBlocksInt(0),
			Timestamp: big.NewInt(0),
		},
	)
}

func TestChallengeToOSPWithMessage(t *testing.T) {
	runExecutionTest(t, []inbox.InboxMessage{makeInitMsg()}, big.NewInt(1200000), big.NewInt(1300000), faultConfig{distortMachineAtGas: big.NewInt(1250000)}, false)
}

func TestChallengeToUnreachable(t *testing.T) {
	rounds := runExecutionTest(t, []inbox.InboxMessage{makeInitMsg()}, big.NewInt(1200000), big.NewInt(1300000), faultConfig{messagesReadCap: big.NewInt(0)}, true)
	if rounds < 2 {
		t.Fatal("TestChallengeToUnreachable failed too early")
	}
}

func TestChallengeToUnreachableSmall(t *testing.T) {
	messages := []inbox.InboxMessage{makeInitMsg()}
	arbCore, shutdown := prepareArbCore(t, messages)
	defer shutdown()
	cursor, err := arbCore.GetExecutionCursor(big.NewInt(1 << 30))
	test.FailIfError(t, err)
	startGas := cursor.TotalGasConsumed()
	endGas := new(big.Int).Add(startGas, big.NewInt(1))

	faultConfig := faultConfig{stallMachineAt: startGas}
	faultyCore := newFaultyCore(arbCore, faultConfig)

	challengedNode := initializeChallengeData(t, faultyCore, startGas, endGas)

	time := big.NewInt(100)
	executeChallenge(
		t,
		challengedNode,
		time,
		time,
		arbCore,
		faultyCore,
		true,
	)
}
