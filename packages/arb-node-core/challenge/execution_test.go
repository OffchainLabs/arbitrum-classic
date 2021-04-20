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

func runExecutionTest(t *testing.T, messages []inbox.InboxMessage, startGas *big.Int, endGas *big.Int, faultConfig FaultConfig, asserterMayFail bool) int {
	mon, shutdown := monitor.PrepareArbCore(t, messages)
	defer shutdown()
	faultyCore := NewFaultyCore(mon.Core, faultConfig)

	challengedNode, err := initializeChallengeData(t, faultyCore, startGas, endGas)
	if err != nil {
		t.Fatal("Error with initializeChallengeData")
	}

	time := big.NewInt(100)
	return executeChallenge(
		t,
		challengedNode,
		time,
		time,
		mon.Core,
		faultyCore,
		asserterMayFail,
	)
}

func TestChallengeToOSP(t *testing.T) {
	runExecutionTest(t, []inbox.InboxMessage{}, big.NewInt(0), big.NewInt(400*2), FaultConfig{DistortMachineAtGas: big.NewInt(1)}, false)
}

func makeInitMsg() inbox.InboxMessage {
	owner := common.RandAddress()
	chain := common.RandAddress()
	return message.NewInboxMessage(
		message.Init{
			ChainParams: protocol.ChainParams{
				StakeRequirement:          big.NewInt(0),
				StakeToken:                common.Address{},
				GracePeriod:               common.NewTimeBlocks(big.NewInt(3)),
				MaxExecutionSteps:         0,
				ArbGasSpeedLimitPerSecond: 0,
			},
			Owner:       owner,
			ExtraConfig: []byte{},
		},
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
	runExecutionTest(t, []inbox.InboxMessage{makeInitMsg()}, start, end, FaultConfig{DistortMachineAtGas: inboxGas}, false)
}

func TestChallengeToUnreachable(t *testing.T) {
	inboxGas := calculateGasToFirstInbox(t)
	start := new(big.Int).Sub(inboxGas, big.NewInt(50000))
	end := new(big.Int).Add(inboxGas, big.NewInt(50000))
	rounds := runExecutionTest(t, []inbox.InboxMessage{makeInitMsg()}, start, end, FaultConfig{MessagesReadCap: big.NewInt(0)}, true)
	if rounds < 2 {
		t.Fatal("TestChallengeToUnreachable failed too early")
	}
}

func calculateGasToFirstInbox(t *testing.T) *big.Int {
	mon, shutdown := monitor.PrepareArbCore(t, nil)
	defer shutdown()
	cursor, err := mon.Core.GetExecutionCursor(big.NewInt(100000000))
	test.FailIfError(t, err)
	return cursor.TotalGasConsumed()
}

func TestChallengeToUnreachableSmall(t *testing.T) {
	messages := []inbox.InboxMessage{makeInitMsg()}
	mon, shutdown := monitor.PrepareArbCore(t, messages)
	defer shutdown()
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
		challengedNode,
		time,
		time,
		mon.Core,
		faultyCore,
		true,
	)
}
