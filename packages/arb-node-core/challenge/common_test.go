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
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
)

func executeChallenge(
	t *testing.T,
	challengedNode *core.NodeInfo,
	correctLookup core.ArbCoreLookup,
	falseLookup core.ArbCoreLookup,
	asserterMayFail bool,
	client *ethutils.SimulatedEthClient,
	tester *ethbridgetestcontracts.ChallengeTester,
	seqInboxAddr ethcommon.Address,
	asserterWallet *ethbridge.ValidatorWallet,
	challengerWallet *ethbridge.ValidatorWallet,
) int {
	ctx := context.Background()

	challengeAddress, err := tester.Challenge(&bind.CallOpts{})
	test.FailIfError(t, err)

	asserterBackend, err := ethbridge.NewBuilderBackend(asserterWallet)
	test.FailIfError(t, err)
	challengerBackend, err := ethbridge.NewBuilderBackend(challengerWallet)
	test.FailIfError(t, err)

	asserterChallengeCon, err := ethbridge.NewChallenge(challengeAddress, 0, client, asserterBackend)
	test.FailIfError(t, err)
	challengerChallengeCon, err := ethbridge.NewChallenge(challengeAddress, 0, client, challengerBackend)
	test.FailIfError(t, err)

	challenge, err := ethbridge.NewChallengeWatcher(challengeAddress, 0, client)
	test.FailIfError(t, err)

	seqInbox, err := ethbridge.NewSequencerInboxWatcher(seqInboxAddr, client)
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
	asserterTime *big.Int,
	challengerTime *big.Int,
	arbCore core.ArbCore,
) (*ethutils.SimulatedEthClient, *ethbridgetestcontracts.ChallengeTester, ethcommon.Address, *ethbridge.ValidatorWallet, *ethbridge.ValidatorWallet, func(nd *core.NodeInfo)) {
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
	rollupAddr, _, rollup, err := ethbridgetestcontracts.DeployRollupMock(deployer, client)
	test.FailIfError(t, err)

	delayedBridgeAddr, _, delayedBridge, err := ethbridgecontracts.DeployBridge(deployer, client)
	test.FailIfError(t, err)
	client.Commit()
	_, err = delayedBridge.Initialize(deployer)
	test.FailIfError(t, err)
	client.Commit()

	_, err = delayedBridge.SetInbox(deployer, deployer.From, true)
	test.FailIfError(t, err)

	maxDelayBlocks := big.NewInt(60)
	maxDelaySeconds := big.NewInt(900)
	_, err = rollup.SetMock(deployer, maxDelayBlocks, maxDelaySeconds)
	test.FailIfError(t, err)
	client.Commit()

	init := makeInit()

	tx, err := delayedBridge.DeliverMessageToInbox(deployer, uint8(init.Type()), rollupAddr, hashing.SoliditySHA3(init.AsData()))
	test.FailIfError(t, err)
	client.Commit()
	initReceipt, err := clnt.TransactionReceipt(context.Background(), tx.Hash())
	test.FailIfError(t, err)
	initBlock, err := clnt.BlockByHash(context.Background(), initReceipt.BlockHash)
	test.FailIfError(t, err)
	initMsg := message.NewInboxMessage(
		init,
		common.NewAddressFromEth(rollupAddr),
		big.NewInt(0),
		tx.GasPrice(),
		inbox.ChainTime{
			BlockNum:  common.NewTimeBlocks(initBlock.Number()),
			Timestamp: new(big.Int).SetUint64(initBlock.Time()),
		},
	)

	sequencerBridgeAddr, _, sequencerBridge, err := ethbridgecontracts.DeploySequencerInbox(deployer, client)
	test.FailIfError(t, err)
	client.Commit()
	_, err = sequencerBridge.Initialize(deployer, delayedBridgeAddr, sequencer.From, rollupAddr)
	test.FailIfError(t, err)
	client.Commit()
	latestHeader, err := client.HeaderByNumber(context.Background(), nil)
	test.FailIfError(t, err)
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocks(latestHeader.Number),
		Timestamp: new(big.Int).SetUint64(latestHeader.Time),
	}

	delayed := inbox.NewDelayedMessage(common.Hash{}, initMsg)
	delayedItem := inbox.NewDelayedItem(big.NewInt(0), big.NewInt(1), common.Hash{}, big.NewInt(0), delayed.DelayedAccumulator)
	endOfBlockMessage := message.NewInboxMessage(
		message.EndBlockMessage{},
		common.Address{},
		big.NewInt(1),
		big.NewInt(0),
		chainTime,
	)
	endOfBlockItem := inbox.NewSequencerItem(big.NewInt(1), endOfBlockMessage, delayedItem.Accumulator)
	delayedAccInt := new(big.Int).SetBytes(delayed.DelayedAccumulator.Bytes())
	batchMetadata := []*big.Int{big.NewInt(0), chainTime.BlockNum.AsInt(), chainTime.Timestamp, big.NewInt(1), delayedAccInt}
	_, err = sequencerBridge.AddSequencerL2BatchFromOrigin(sequencer, nil, nil, batchMetadata, endOfBlockItem.Accumulator)
	test.FailIfError(t, err)

	err = core.DeliverMessagesAndWait(arbCore, big.NewInt(0), common.Hash{}, []inbox.SequencerBatchItem{delayedItem, endOfBlockItem}, []inbox.DelayedMessage{delayed}, nil)
	test.FailIfError(t, err)

	asserterWalletAddress, _, validatorCon, err := ethbridgecontracts.DeployValidator(asserter, client)
	test.FailIfError(t, err)
	client.Commit()
	_, err = validatorCon.Initialize(asserter)
	test.FailIfError(t, err)

	challengerWalletAddress, _, validatorCon2, err := ethbridgecontracts.DeployValidator(challenger, client)
	test.FailIfError(t, err)
	client.Commit()
	_, err = validatorCon2.Initialize(challenger)
	test.FailIfError(t, err)

	asserterAuth, err := ethbridge.NewTransactAuth(ctx, client, asserter, "")
	test.FailIfError(t, err)
	asserterWallet, err := ethbridge.NewValidator(asserterWalletAddress, ethcommon.Address{}, client, asserterAuth)
	test.FailIfError(t, err)

	challengerAuth, err := ethbridge.NewTransactAuth(ctx, client, challenger, "")
	test.FailIfError(t, err)
	challengerWallet, err := ethbridge.NewValidator(challengerWalletAddress, ethcommon.Address{}, client, challengerAuth)
	test.FailIfError(t, err)

	startChallenge := func(nd *core.NodeInfo) {
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
	}

	return client, tester, sequencerBridgeAddr, asserterWallet, challengerWallet, startChallenge
}
