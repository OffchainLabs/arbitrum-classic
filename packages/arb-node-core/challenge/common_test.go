package challenge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func executeChallenge(
	t *testing.T,
	challengedNode *core.NodeInfo,
	arbGasSpeedLimitPerBlock *big.Int,
	challengePeriodBlocks *big.Int,
	correctLookup *core.ValidatorLookupMock,
	falseLookup *core.ValidatorLookupMock,
) int {
	ctx := context.Background()

	client, tester, asserterWallet, challengerWallet, challengeAddress := initializeChallengeTest(t, challengedNode, arbGasSpeedLimitPerBlock, challengePeriodBlocks)

	challengeCon, err := ethbridge.NewChallenge(challengeAddress, client)
	test.FailIfError(t, err)

	challenger := NewChallenger(challengeCon, correctLookup, challengedNode, challengerWallet.Address())
	asserter := NewChallenger(challengeCon, falseLookup, challengedNode, asserterWallet.Address())

	turn := ethbridge.CHALLENGER_TURN
	rounds := 0
	for {
		checkTurn(t, challengeCon.ChallengeWatcher, turn)
		if turn == ethbridge.CHALLENGER_TURN {
			rawTx, err := challenger.HandleConflict(ctx)
			test.FailIfError(t, err)
			if rawTx == nil {
				t.Fatal("should be able to transact")
			}
			_, err = challengerWallet.ExecuteTransaction(ctx, rawTx)
			test.FailIfError(t, err)
			turn = ethbridge.ASSERTER_TURN
		} else {
			rawTx, err := asserter.HandleConflict(ctx)
			test.FailIfError(t, err)
			if rawTx == nil {
				t.Fatal("should be able to transact")
			}
			_, err = asserterWallet.ExecuteTransaction(ctx, rawTx)
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

		checkTurn(t, challengeCon.ChallengeWatcher, turn)
	}

	checkChallengeCompleted(t, tester, challengerWallet.Address().ToEthAddress(), asserterWallet.Address().ToEthAddress())
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
	initialMach, err := lookup.GetCursor(big.NewInt(0))
	test.FailIfError(t, err)
	prevState := &core.NodeState{
		ProposedBlock: big.NewInt(0),
		InboxMaxCount: big.NewInt(0),
		ExecutionState: &core.ExecutionState{
			TotalGasConsumed: big.NewInt(0),
			MachineHash:      initialMach.MachineHash(),
			InboxHash:        common.Hash{},
			InboxIndex:       big.NewInt(0),
			TotalSendCount:   big.NewInt(0),
			TotalLogCount:    big.NewInt(0),
		},
	}

	messages, err := lookup.GetMessages(big.NewInt(0), inboxMessagesRead)
	test.FailIfError(t, err)
	afterInboxCount := new(big.Int).Add(prevState.InboxIndex, inboxMessagesRead)

	inboxAcc, err := lookup.GetInboxAcc(afterInboxCount)
	test.FailIfError(t, err)

	assertion := &core.Assertion{
		PrevProposedBlock: prevState.ProposedBlock,
		PrevInboxMaxCount: prevState.InboxMaxCount,
		ExecutionInfo: &core.ExecutionInfo{
			Before: prevState.ExecutionState,
			After: &core.ExecutionState{
				MachineHash:      common.Hash{},
				InboxIndex:       inboxMessagesRead,
				InboxHash:        inboxAcc,
				TotalGasConsumed: big.NewInt(0),
				TotalSendCount:   big.NewInt(0),
				TotalLogCount:    big.NewInt(0),
			},
			SendAcc: common.Hash{},
			LogAcc:  common.Hash{},
		},
		InboxDelta: core.CalculateInboxDeltaAcc(messages),
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
) (*ethutils.SimulatedEthClient, *ethbridgetestcontracts.ChallengeTester, *ethbridge.Validator, *ethbridge.Validator, ethcommon.Address) {
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

	asserterWalletAddress, _, _, err := ethbridgecontracts.DeployValidator(asserter, client)
	test.FailIfError(t, err)
	challengerWalletAddress, _, _, err := ethbridgecontracts.DeployValidator(challenger, client)
	test.FailIfError(t, err)

	asserterWallet, err := ethbridge.NewValidator(asserterWalletAddress, client, ethbridge.NewTransactAuth(asserter))
	test.FailIfError(t, err)

	challengerWallet, err := ethbridge.NewValidator(challengerWalletAddress, client, ethbridge.NewTransactAuth(challenger))
	test.FailIfError(t, err)

	_, err = tester.StartChallenge(
		deployer,
		nd.Assertion.InboxConsistencyHash(nd.InboxMaxHash, nd.InboxMaxCount),
		nd.Assertion.InboxDeltaHash(),
		nd.Assertion.ExecutionHash(),
		nd.Assertion.CheckTime(arbGasLimitPerBlock),
		asserterWallet.Address().ToEthAddress(),
		challengerWallet.Address().ToEthAddress(),
		challengePeriodBlocks,
	)
	test.FailIfError(t, err)
	client.Commit()
	challengeAddress, err := tester.Challenge(&bind.CallOpts{})
	test.FailIfError(t, err)

	return client, tester, asserterWallet, challengerWallet, challengeAddress
}
