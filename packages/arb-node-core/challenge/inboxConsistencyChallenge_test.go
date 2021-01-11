package challenge

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func TestInboxConsistencyChallenge(t *testing.T) {
	mach, err := cmachine.New(arbos.Path())
	test.FailIfError(t, err)

	correctLookup := core.NewValidatorLookupMock(mach)
	for i := 0; i < 500; i++ {
		correctLookup.AddMessage(inbox.NewRandomInboxMessage())
	}
	otherLookup := core.NewValidatorLookupMock(mach)
	for i := 0; i < 500; i++ {
		otherLookup.AddMessage(inbox.NewRandomInboxMessage())
	}

	falseLookup := correctLookup.Clone()
	for i := 2; i < 6; i++ {
		falseLookup.InboxAccs[i] = otherLookup.InboxAccs[i]
	}

	inboxMessagesRead := big.NewInt(4)

	challengedNode := initializeChallengeData(t, correctLookup, inboxMessagesRead)

	inboxAcc, err := falseLookup.GetInboxAcc(new(big.Int).Add(challengedNode.Assertion.AfterInboxCount(), big.NewInt(1)))
	test.FailIfError(t, err)
	challengedNode.Assertion.AfterInboxHash = inboxAcc

	arbGasSpeedLimitPerBlock := big.NewInt(100000)
	challengePeriodBlocks := big.NewInt(100)

	rounds := executeChallenge(t, challengedNode, arbGasSpeedLimitPerBlock, challengePeriodBlocks, correctLookup, falseLookup)
	if rounds != 3 {
		t.Fatal("wrong round count", rounds)
	}
}

func executeChallenge(
	t *testing.T,
	challengedNode *core.NodeInfo,
	arbGasSpeedLimitPerBlock *big.Int,
	challengePeriodBlocks *big.Int,
	correctLookup *core.ValidatorLookupMock,
	falseLookup *core.ValidatorLookupMock,
) int {
	ctx := context.Background()

	client, tester, asserterAuth, challengerAuth, challengeAddress := initializeChallengeTest(t, challengedNode, arbGasSpeedLimitPerBlock, challengePeriodBlocks)

	challengerChallenge, err := ethbridge.NewChallenge(challengeAddress, client, ethbridge.NewTransactAuth(challengerAuth))
	test.FailIfError(t, err)
	asserterChallenge, err := ethbridge.NewChallenge(challengeAddress, client, ethbridge.NewTransactAuth(asserterAuth))
	test.FailIfError(t, err)

	challenger := NewChallenger(challengerChallenge, correctLookup, challengedNode)
	asserter := NewChallenger(asserterChallenge, falseLookup, challengedNode)

	turn := ethbridge.CHALLENGER_TURN
	rounds := 0
	for {
		checkTurn(t, challengerChallenge.ChallengeWatcher, turn)
		if turn == ethbridge.CHALLENGER_TURN {
			_, err := challenger.handleConflict(ctx)
			test.FailIfError(t, err)
			turn = ethbridge.ASSERTER_TURN
		} else {
			_, err := asserter.handleConflict(ctx)
			test.FailIfError(t, err)
			turn = ethbridge.CHALLENGER_TURN
		}
		rounds++

		client.Commit()

		completed, err := tester.ChallengeCompleted(&bind.CallOpts{Context: ctx})
		test.FailIfError(t, err)
		if completed {
			break
		}

		checkTurn(t, challengerChallenge.ChallengeWatcher, turn)
	}

	checkChallengeCompleted(t, tester, challengerAuth.From, asserterAuth.From)
	return rounds
}

func checkTurn(t *testing.T, challenge *ethbridge.ChallengeWatcher, turn ethbridge.ChallengeTurn) {
	t.Helper()
	ctx := context.Background()
	currentTurn, err := challenge.Turn(ctx)
	test.FailIfError(t, err)
	if currentTurn != turn {
		t.Fatal("wrong player's turn")
	}
}

func checkChallengeCompleted(t *testing.T, tester *ethbridgetestcontracts.ChallengeTester, correctWinner, correctLoser ethcommon.Address) {
	ctx := context.Background()
	completed, err := tester.ChallengeCompleted(&bind.CallOpts{Context: ctx})
	test.FailIfError(t, err)

	if !completed {
		t.Fatal("should be completed")
	}

	winner, err := tester.Winner(&bind.CallOpts{Context: ctx})
	test.FailIfError(t, err)

	if winner != correctWinner {
		t.Fatal("winner should be challenger")
	}

	loser, err := tester.Loser(&bind.CallOpts{Context: ctx})
	test.FailIfError(t, err)

	if loser != correctLoser {
		t.Fatal("loser should be challenger")
	}
}

func initializeChallengeData(
	t *testing.T,
	lookup *core.ValidatorLookupMock,
	inboxMessagesRead *big.Int,
) *core.NodeInfo {
	initialMach, err := lookup.GetMachine(big.NewInt(0))
	test.FailIfError(t, err)
	prevState := &core.NodeState{
		ProposedBlock:  big.NewInt(0),
		TotalGasUsed:   big.NewInt(0),
		MachineHash:    initialMach.Hash(),
		InboxHash:      common.Hash{},
		InboxCount:     big.NewInt(0),
		TotalSendCount: big.NewInt(0),
		TotalLogCount:  big.NewInt(0),
		InboxMaxCount:  big.NewInt(0),
	}

	messages, err := lookup.GetMessages(big.NewInt(0), inboxMessagesRead)
	test.FailIfError(t, err)
	afterInboxCount := new(big.Int).Add(prevState.InboxCount, inboxMessagesRead)

	inboxAcc, err := lookup.GetInboxAcc(afterInboxCount)
	test.FailIfError(t, err)
	assertionInfo := &core.AssertionInfo{
		InboxDelta: core.CalculateInboxDeltaAcc(messages),
		ExecInfo: &core.ExecutionInfo{
			BeforeMachineHash: common.Hash{},
			InboxMessagesRead: inboxMessagesRead,
			GasUsed:           big.NewInt(0),
			SendAcc:           common.Hash{},
			SendCount:         big.NewInt(0),
			LogAcc:            common.Hash{},
			LogCount:          big.NewInt(0),
			AfterMachineHash:  common.Hash{},
		},
		AfterInboxHash: inboxAcc,
	}

	assertion := &core.Assertion{
		PrevState:     prevState,
		AssertionInfo: assertionInfo,
	}

	inboxMaxCount := big.NewInt(int64(len(lookup.InboxAccs)) - 1)
	inboxTopAcc, err := lookup.GetInboxAcc(inboxMaxCount)
	test.FailIfError(t, err)
	return &core.NodeInfo{
		NodeNum: big.NewInt(1),
		BlockProposed: &common.BlockId{
			Height:     common.NewTimeBlocks(common.RandBigInt()),
			HeaderHash: common.RandHash(),
		},
		Assertion:     assertion,
		InboxMaxCount: inboxMaxCount,
		InboxMaxHash:  inboxTopAcc,
	}
}

func initializeChallengeTest(
	t *testing.T,
	nd *core.NodeInfo,
	arbGasLimitPerBlock *big.Int,
	challengePeriodBlocks *big.Int,
) (*ethutils.SimulatedEthClient, *ethbridgetestcontracts.ChallengeTester, *bind.TransactOpts, *bind.TransactOpts, ethcommon.Address) {
	clnt, pks := test.SimulatedBackend()
	deployer := bind.NewKeyedTransactor(pks[0])
	asserter := bind.NewKeyedTransactor(pks[1])
	challenger := bind.NewKeyedTransactor(pks[2])
	client := &ethutils.SimulatedEthClient{SimulatedBackend: clnt}
	osp1Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProof(deployer, client)
	test.FailIfError(t, err)
	osp2Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProof2(deployer, client)
	test.FailIfError(t, err)
	_, _, tester, err := ethbridgetestcontracts.DeployChallengeTester(deployer, client, osp1Addr, osp2Addr)
	test.FailIfError(t, err)
	_, err = tester.StartChallenge(
		deployer,
		nd.Assertion.InboxConsistencyHash(nd.InboxMaxHash, nd.InboxMaxCount),
		nd.Assertion.InboxDeltaHash(),
		nd.Assertion.ExecutionHash(),
		nd.Assertion.CheckTime(arbGasLimitPerBlock),
		asserter.From,
		challenger.From,
		challengePeriodBlocks,
	)
	test.FailIfError(t, err)
	client.Commit()
	challengeAddress, err := tester.Challenge(&bind.CallOpts{})
	test.FailIfError(t, err)
	return client, tester, asserter, challenger, challengeAddress
}
