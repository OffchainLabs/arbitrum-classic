package challenge

import (
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func runExecutionTest(t *testing.T, messages []inbox.InboxMessage, startGas *big.Int, endGas *big.Int, faultConfig FaultConfig, asserterMayFail bool) int {
	arbCore, shutdown := PrepareTestArbCore(t, messages)
	defer shutdown()
	faultyCore := NewFaultyCore(arbCore, faultConfig)

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
	runExecutionTest(t, []inbox.InboxMessage{}, big.NewInt(0), big.NewInt(200000), FaultConfig{DistortMachineAtGas: big.NewInt(100000)}, false)
}

func TestChallengeToOSPWithMessage(t *testing.T) {
	runExecutionTest(t, []inbox.InboxMessage{MakeTestInitMsg()}, big.NewInt(1200000), big.NewInt(1300000), FaultConfig{DistortMachineAtGas: big.NewInt(1250000)}, false)
}

func TestChallengeToUnreachable(t *testing.T) {
	rounds := runExecutionTest(t, []inbox.InboxMessage{MakeTestInitMsg()}, big.NewInt(1200000), big.NewInt(1300000), FaultConfig{MessagesReadCap: big.NewInt(0)}, true)
	if rounds < 2 {
		t.Fatal("TestChallengeToUnreachable failed too early")
	}
}

func TestChallengeToUnreachableSmall(t *testing.T) {
	messages := []inbox.InboxMessage{MakeTestInitMsg()}
	arbCore, shutdown := PrepareTestArbCore(t, messages)
	defer shutdown()
	cursor, err := arbCore.GetExecutionCursor(big.NewInt(1 << 30))
	test.FailIfError(t, err)
	startGas := cursor.TotalGasConsumed()
	endGas := new(big.Int).Add(startGas, big.NewInt(1))

	faultConfig := FaultConfig{StallMachineAt: startGas}
	faultyCore := NewFaultyCore(arbCore, faultConfig)

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
