package challenge

import (
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

func runExecutionTest(t *testing.T, messages []inbox.InboxMessage, startGas *big.Int, endGas *big.Int, faultConfig FaultConfig, asserterMayFail bool) int {
	arbCore, shutdown := test.PrepareArbCore(t, messages)
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
	runExecutionTest(t, []inbox.InboxMessage{makeInitMsg()}, big.NewInt(1200000), big.NewInt(1300000), FaultConfig{DistortMachineAtGas: big.NewInt(1250000)}, false)
}

func TestChallengeToUnreachable(t *testing.T) {
	rounds := runExecutionTest(t, []inbox.InboxMessage{makeInitMsg()}, big.NewInt(1200000), big.NewInt(1300000), FaultConfig{MessagesReadCap: big.NewInt(0)}, true)
	if rounds < 2 {
		t.Fatal("TestChallengeToUnreachable failed too early")
	}
}

func TestChallengeToUnreachableSmall(t *testing.T) {
	messages := []inbox.InboxMessage{makeInitMsg()}
	arbCore, shutdown := test.PrepareArbCore(t, messages)
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
