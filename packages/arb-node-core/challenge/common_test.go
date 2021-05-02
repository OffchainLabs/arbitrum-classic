package challenge

import (
	"context"
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
)

func executeChallenge(
	t *testing.T,
	init inbox.InboxMessage,
	challengedNode *core.NodeInfo,
	asserterTime *big.Int,
	challengerTime *big.Int,
	correctLookup core.ArbCoreLookup,
	falseLookup core.ArbCoreLookup,
	asserterMayFail bool,
) int {
	ctx := context.Background()

	client, tester, seqInbox, asserterWallet, challengerWallet, challengeAddress := initializeChallengeTest(t, challengedNode, asserterTime, challengerTime, init)

	asserterBackend, err := ethbridge.NewBuilderBackend(asserterWallet)
	test.FailIfError(t, err)
	challengerBackend, err := ethbridge.NewBuilderBackend(challengerWallet)
	test.FailIfError(t, err)

	asserterChallengeCon, err := ethbridge.NewChallenge(challengeAddress, client, asserterBackend)
	test.FailIfError(t, err)
	challengerChallengeCon, err := ethbridge.NewChallenge(challengeAddress, client, challengerBackend)
	test.FailIfError(t, err)

	challenge, err := ethbridge.NewChallengeWatcher(challengeAddress, client)
	test.FailIfError(t, err)

	challenger := NewChallenger(challengerChallengeCon, seqInbox, correctLookup, challengedNode, challengerWallet.Address())
	asserter := NewChallenger(asserterChallengeCon, seqInbox, falseLookup, challengedNode, asserterWallet.Address())

	turn := ethbridge.CHALLENGER_TURN
	rounds := 0
	for {
		t.Logf("executing challenge round %v", rounds)
		checkTurn(t, challenge, turn)
		if turn == ethbridge.CHALLENGER_TURN {
			err := challenger.HandleConflict(ctx)
			test.FailIfError(t, err)

			if challengerBackend.TransactionCount() == 0 {
				t.Fatal("should be able to transact")
			}
			tx, err := challengerWallet.ExecuteTransactions(ctx, challengerBackend)
			test.FailIfError(t, err)
			client.Commit()
			receipt, err := client.TransactionReceipt(ctx, tx.Hash())
			test.FailIfError(t, err)
			t.Log("Challenger Used", receipt.GasUsed, "gas")
			turn = ethbridge.ASSERTER_TURN
		} else {
			err := asserter.HandleConflict(ctx)
			if asserterMayFail && err != nil {
				t.Logf("Asserter failed challenge: %v", err.Error())
				return rounds
			}
			test.FailIfError(t, err)
			if asserterBackend.TransactionCount() == 0 {
				t.Fatal("should be able to transact")
			}
			tx, err := asserterWallet.ExecuteTransactions(ctx, asserterBackend)
			test.FailIfError(t, err)
			client.Commit()
			receipt, err := client.TransactionReceipt(ctx, tx.Hash())
			test.FailIfError(t, err)
			t.Log("Asserter Used", receipt.GasUsed, "gas")
			turn = ethbridge.CHALLENGER_TURN
		}
		rounds++

		completed, err := tester.ChallengeCompleted(&bind.CallOpts{Context: ctx})
		test.FailIfError(t, err)
		if completed {
			break
		}

		checkTurn(t, challenge, turn)
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

func initializeChallengeData(t *testing.T, lookup core.ArbCoreLookup, startGas *big.Int, endGas *big.Int) (*core.NodeInfo, error) {
	cursor, err := lookup.GetExecutionCursor(startGas)
	test.FailIfError(t, err)
	inboxMaxCount, err := lookup.GetMessageCount()
	test.FailIfError(t, err)
	prevExecState, err := core.NewExecutionState(cursor)
	test.FailIfError(t, err)
	prevState := &core.NodeState{
		ProposedBlock:  big.NewInt(0),
		InboxMaxCount:  inboxMaxCount,
		ExecutionState: prevExecState,
	}

	err = lookup.AdvanceExecutionCursor(cursor, endGas, true)
	test.FailIfError(t, err)
	after, err := core.NewExecutionState(cursor)
	test.FailIfError(t, err)
	if err != nil {
		return nil, err
	}
	assertion := &core.Assertion{
		Before: prevState.ExecutionState,
		After:  after,
	}

	return &core.NodeInfo{
		NodeNum: big.NewInt(1),
		BlockProposed: &common.BlockId{
			Height:     common.NewTimeBlocks(common.RandBigInt()),
			HeaderHash: common.RandHash(),
		},
		Assertion:          assertion,
		InboxMaxCount:      inboxMaxCount,
		NodeHash:           common.RandHash(),
		AfterInboxBatchAcc: [32]byte{},
	}, nil
}

func initializeChallengeTest(
	t *testing.T,
	nd *core.NodeInfo,
	asserterTime *big.Int,
	challengerTime *big.Int,
	init inbox.InboxMessage,
) (*ethutils.SimulatedEthClient, *ethbridgetestcontracts.ChallengeTester, *ethbridge.SequencerInboxWatcher, *ethbridge.ValidatorWallet, *ethbridge.ValidatorWallet, ethcommon.Address) {
	ctx := context.Background()
	clnt, pks := test.SimulatedBackend(t)
	deployer := bind.NewKeyedTransactor(pks[0])
	asserter := bind.NewKeyedTransactor(pks[1])
	challenger := bind.NewKeyedTransactor(pks[2])
	sequencer := bind.NewKeyedTransactor(pks[3])
	client := &ethutils.SimulatedEthClient{SimulatedBackend: clnt}
	osp1Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProof(deployer, client)
	test.FailIfError(t, err)
	osp2Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProof2(deployer, client)
	test.FailIfError(t, err)
	osp3Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProofHash(deployer, client)
	test.FailIfError(t, err)
	_, _, tester, err := ethbridgetestcontracts.DeployChallengeTester(deployer, client, []ethcommon.Address{osp1Addr, osp2Addr, osp3Addr})
	test.FailIfError(t, err)

	delayedBridgeAddr, _, delayedBridge, err := ethbridgecontracts.DeployBridge(deployer, client)
	test.FailIfError(t, err)
	client.Commit()

	_, err = delayedBridge.SetInbox(deployer, deployer.From, true)
	test.FailIfError(t, err)
	client.Commit()

	_, err = delayedBridge.DeliverMessageToInbox(deployer, uint8(init.Kind), init.Sender.ToEthAddress(), hashing.SoliditySHA3(init.Data))
	test.FailIfError(t, err)

	maxDelayBlocks := big.NewInt(60)
	maxDelaySeconds := big.NewInt(900)
	sequencerBridgeAddr, _, sequencerBridge, err := ethbridgecontracts.DeploySequencerInbox(deployer, client, delayedBridgeAddr, sequencer.From, maxDelayBlocks, maxDelaySeconds)
	test.FailIfError(t, err)
	client.Commit()

	delayedAcc, err := delayedBridge.InboxAccs(&bind.CallOpts{}, big.NewInt(0))
	test.FailIfError(t, err)

	latestHeader, err := client.HeaderByNumber(context.Background(), nil)
	test.FailIfError(t, err)
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocks(latestHeader.Number),
		Timestamp: new(big.Int).SetUint64(latestHeader.Time),
	}

	delayedItem := inbox.NewDelayedItem(big.NewInt(0), big.NewInt(1), common.Hash{}, big.NewInt(0), delayedAcc)
	endOfBlockMessage := message.NewInboxMessage(
		message.EndBlockMessage{},
		common.Address{},
		big.NewInt(1),
		big.NewInt(0),
		chainTime,
	)
	endOfBlockItem := inbox.NewSequencerItem(big.NewInt(1), endOfBlockMessage, delayedItem.Accumulator)
	_, err = sequencerBridge.AddSequencerL2BatchFromOrigin(sequencer, nil, nil, chainTime.BlockNum.AsInt(), chainTime.Timestamp, big.NewInt(1), endOfBlockItem.Accumulator)
	test.FailIfError(t, err)

	asserterWalletAddress, _, _, err := ethbridgecontracts.DeployValidator(asserter, client)
	test.FailIfError(t, err)
	challengerWalletAddress, _, _, err := ethbridgecontracts.DeployValidator(challenger, client)
	test.FailIfError(t, err)

	asserterAuth, err := ethbridge.NewTransactAuth(ctx, client, asserter)
	test.FailIfError(t, err)
	asserterWallet, err := ethbridge.NewValidator(asserterWalletAddress, ethcommon.Address{}, client, asserterAuth)
	test.FailIfError(t, err)

	challengerAuth, err := ethbridge.NewTransactAuth(ctx, client, challenger)
	test.FailIfError(t, err)
	challengerWallet, err := ethbridge.NewValidator(challengerWalletAddress, ethcommon.Address{}, client, challengerAuth)
	test.FailIfError(t, err)

	_, err = tester.StartChallenge(
		deployer,
		nd.Assertion.ExecutionHash(),
		nd.Assertion.After.TotalMessagesRead,
		asserterWallet.Address().ToEthAddress(),
		challengerWallet.Address().ToEthAddress(),
		asserterTime,
		challengerTime,
		sequencerBridgeAddr,
		delayedBridgeAddr,
	)
	test.FailIfError(t, err)
	client.Commit()
	challengeAddress, err := tester.Challenge(&bind.CallOpts{})
	test.FailIfError(t, err)

	sequencerInboxWatcher, err := ethbridge.NewSequencerInboxWatcher(sequencerBridgeAddr, client)
	test.FailIfError(t, err)

	return client, tester, sequencerInboxWatcher, asserterWallet, challengerWallet, challengeAddress
}
