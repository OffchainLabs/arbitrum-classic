package challenge

import (
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

func runExecutionTest(t *testing.T, init message.Init, messages []inbox.InboxMessage, startGas *big.Int, endGas *big.Int, faultConfig FaultConfig, asserterMayFail bool) int {
	mon, shutdown := monitor.PrepareArbCore(t)
	defer shutdown()

	monitor.DeliverMessagesToCore(t, mon.Core, big.NewInt(0), common.Hash{}, messages)
	faultyCore := NewFaultyCore(mon.Core, faultConfig)

	challengedNode, err := initializeChallengeData(t, faultyCore, startGas, endGas)
	if err != nil {
		t.Fatal("Error with initializeChallengeData")
	}

	time := big.NewInt(100)
	return executeChallenge(
		t,
		init,
		challengedNode,
		time,
		time,
		mon.Core,
		faultyCore,
		asserterMayFail,
	)
}

func TestChallengeToOSP(t *testing.T) {
	runExecutionTest(t, makeInit(), nil, big.NewInt(0), big.NewInt(400*2), FaultConfig{DistortMachineAtGas: big.NewInt(1)}, false)
}

func makeInit() message.Init {
	return message.Init{
		ChainParams: protocol.ChainParams{
			StakeRequirement:          big.NewInt(0),
			StakeToken:                common.Address{},
			GracePeriod:               common.NewTimeBlocks(big.NewInt(3)),
			MaxExecutionSteps:         0,
			ArbGasSpeedLimitPerSecond: 0,
		},
		Owner:       common.RandAddress(),
		ExtraConfig: []byte{},
	}
}

func makeInitMsg() inbox.InboxMessage {
	chain := common.RandAddress()
	return message.NewInboxMessage(
		makeInit(),
		chain,
		big.NewInt(0),
		big.NewInt(0),
		inbox.ChainTime{
			BlockNum:  common.NewTimeBlocksInt(0),
			Timestamp: big.NewInt(0),
		},
	)
}

func TestChallengeToOSPWithMessage(t *testing.T) {
	inboxGas := calculateGasToFirstInbox(t)
	start := new(big.Int).Sub(inboxGas, big.NewInt(50000))
	end := new(big.Int).Add(inboxGas, big.NewInt(50000))
	runExecutionTest(t, makeInit(), []inbox.InboxMessage{makeInitMsg()}, start, end, FaultConfig{DistortMachineAtGas: inboxGas}, false)
}

func TestChallengeToUnreachable(t *testing.T) {
	inboxGas := calculateGasToFirstInbox(t)
	start := new(big.Int).Sub(inboxGas, big.NewInt(50000))
	end := new(big.Int).Add(inboxGas, big.NewInt(50000))
	rounds := runExecutionTest(t, makeInit(), []inbox.InboxMessage{makeInitMsg()}, start, end, FaultConfig{MessagesReadCap: big.NewInt(0)}, true)
	if rounds < 2 {
		t.Fatal("TestChallengeToUnreachable failed too early")
	}
}

func calculateGasToFirstInbox(t *testing.T) *big.Int {
	mon, shutdown := monitor.PrepareArbCore(t)
	defer shutdown()
	cursor, err := mon.Core.GetExecutionCursor(big.NewInt(100000000))
	test.FailIfError(t, err)
	inboxGas := new(big.Int).Add(cursor.TotalGasConsumed(), big.NewInt(1))
	t.Logf("Found first inbox instruction starting at %v", inboxGas)
	return inboxGas
}

func TestChallengeToUnreachableSmall(t *testing.T) {
	messages := []inbox.InboxMessage{makeInitMsg()}
	mon, shutdown := monitor.PrepareArbCore(t)
	defer shutdown()
	monitor.DeliverMessagesToCore(t, mon.Core, big.NewInt(0), common.Hash{}, messages)
	cursor, err := mon.Core.GetExecutionCursor(big.NewInt(1 << 30))
	test.FailIfError(t, err)
	startGas := cursor.TotalGasConsumed()
	endGas := new(big.Int).Add(startGas, big.NewInt(1))

	faultConfig := FaultConfig{StallMachineAt: startGas}
	faultyCore := NewFaultyCore(mon.Core, faultConfig)

	challengedNode, _ := initializeChallengeData(t, faultyCore, startGas, endGas)

	time := big.NewInt(100)
	executeChallenge(
		t,
		makeInit(),
		challengedNode,
		time,
		time,
		mon.Core,
		faultyCore,
		true,
	)
}
