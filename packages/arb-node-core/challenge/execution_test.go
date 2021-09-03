package challenge

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

type ChallengeTestData struct {
	ChallengedAssertion *core.Assertion
	Messages            []inbox.InboxMessage
	Moves               []Move
}

func runExecutionTest(t *testing.T, startGas *big.Int, endGas *big.Int, faultConfig FaultConfig, asserterMayFail bool) int {
	mon, shutdown := monitor.PrepareArbCore(t)
	defer shutdown()

	client, tester, seqInboxAddr, asserterWallet, challengerWallet, startChallenge, messages := initializeChallengeTest(t, big.NewInt(10), big.NewInt(10), mon.Core)

	faultyCore := NewFaultyCore(mon.Core, faultConfig)

	challengedAssertion, err := initializeChallengeData(t, faultyCore, startGas, endGas)
	if err != nil {
		t.Fatal("Error with initializeChallengeData")
	}

	startChallenge(challengedAssertion)
	rounds, moves := executeChallenge(
		t,
		challengedAssertion,
		mon.Core,
		faultyCore,
		asserterMayFail,
		client,
		tester,
		seqInboxAddr,
		asserterWallet,
		challengerWallet,
	)

	challengeData := ChallengeTestData{
		ChallengedAssertion: challengedAssertion,
		Messages:            messages,
		Moves:               moves,
	}
	o, err := json.Marshal(challengeData)
	test.FailIfError(t, err)
	err = ioutil.WriteFile(t.Name()+".json", o, 0777)
	test.FailIfError(t, err)
	return rounds
}

func TestChallengeToOSP(t *testing.T) {
	runExecutionTest(t, big.NewInt(0), big.NewInt(400*2), FaultConfig{DistortMachineAtGas: big.NewInt(1)}, false)
}

func makeInit() message.Init {
	return message.Init{
		ChainParams: protocol.ChainParams{
			GracePeriod:               common.NewTimeBlocks(big.NewInt(3)),
			ArbGasSpeedLimitPerSecond: 0,
		},
		Owner:       common.RandAddress(),
		ExtraConfig: []byte{},
	}
}

func TestChallengeToOSPWithMessage(t *testing.T) {
	inboxGas := calculateGasToFirstInbox(t)
	start := new(big.Int).Sub(inboxGas, big.NewInt(50000))
	end := new(big.Int).Add(inboxGas, big.NewInt(50000))
	runExecutionTest(t, start, end, FaultConfig{DistortMachineAtGas: inboxGas}, false)
}

func TestChallengeToUnreachable(t *testing.T) {
	inboxGas := calculateGasToFirstInbox(t)
	start := new(big.Int).Sub(inboxGas, big.NewInt(50000))
	end := new(big.Int).Add(inboxGas, big.NewInt(50000))
	rounds := runExecutionTest(t, start, end, FaultConfig{MessagesReadCap: big.NewInt(0)}, true)
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
	mon, shutdown := monitor.PrepareArbCore(t)
	defer shutdown()
	client, tester, seqInboxAddr, asserterWallet, challengerWallet, startChallenge, _ := initializeChallengeTest(t, big.NewInt(10), big.NewInt(10), mon.Core)
	cursor, err := mon.Core.GetExecutionCursor(big.NewInt(1 << 30))
	test.FailIfError(t, err)
	startGas := cursor.TotalGasConsumed()
	endGas := new(big.Int).Add(startGas, big.NewInt(1))

	faultConfig := FaultConfig{StallMachineAt: startGas}
	faultyCore := NewFaultyCore(mon.Core, faultConfig)

	challengedAssertion, _ := initializeChallengeData(t, faultyCore, startGas, endGas)

	startChallenge(challengedAssertion)

	executeChallenge(
		t,
		challengedAssertion,
		mon.Core,
		faultyCore,
		true,
		client,
		tester,
		seqInboxAddr,
		asserterWallet,
		challengerWallet,
	)
}
