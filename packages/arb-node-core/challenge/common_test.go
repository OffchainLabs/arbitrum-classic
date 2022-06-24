/*
 * Copyright 2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package challenge

import (
	"context"
	"math/big"
	"math/rand"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/transactauth"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
)

func executeChallenge(
	t *testing.T,
	challengedAssertion *core.Assertion,
	correctLookup core.ArbCoreLookup,
	falseLookup core.ArbCoreLookup,
	client *ethutils.SimulatedEthClient,
	tester *ethbridgetestcontracts.ChallengeTester,
	seqInboxAddr ethcommon.Address,
	asserterWallet *ethbridge.ValidatorWallet,
	challengerWallet *ethbridge.ValidatorWallet,
) ([]Move, error) {
	ctx := context.Background()

	challengeAddress, err := tester.Challenge(&bind.CallOpts{})
	test.FailIfError(t, err)

	asserterBackend, err := ethbridge.NewBuilderBackend(asserterWallet)
	test.FailIfError(t, err)
	challengerBackend, err := ethbridge.NewBuilderBackend(challengerWallet)
	test.FailIfError(t, err)

	asserterChallengeCon, err := ethbridge.NewChallenge(challengeAddress, 0, client, asserterBackend, bind.CallOpts{})
	test.FailIfError(t, err)
	challengerChallengeCon, err := ethbridge.NewChallenge(challengeAddress, 0, client, challengerBackend, bind.CallOpts{})
	test.FailIfError(t, err)

	challenge, err := ethbridge.NewChallengeWatcher(challengeAddress, 0, client, bind.CallOpts{})
	test.FailIfError(t, err)

	seqInbox, err := ethbridge.NewSequencerInboxWatcher(seqInboxAddr, client)
	test.FailIfError(t, err)

	challenger := NewChallenger(challengerChallengeCon, seqInbox, correctLookup, challengedAssertion, common.NewAddressFromEth(*challengerWallet.Address()))
	asserter := NewChallenger(asserterChallengeCon, seqInbox, falseLookup, challengedAssertion, common.NewAddressFromEth(*asserterWallet.Address()))

	var moves []Move

	turn := ethbridge.CHALLENGER_TURN
	rounds := 0
	for {
		t.Logf("executing challenge round %v", rounds)
		checkTurn(t, challenge, turn)
		if turn == ethbridge.CHALLENGER_TURN {
			move, err := challenger.HandleConflict(ctx)
			moves = append(moves, move)
			test.FailIfError(t, err)
			arbTx, err := challengerWallet.ExecuteTransactions(ctx, challengerBackend)
			test.FailIfError(t, err)
			client.Commit()
			if arbTx != nil {
				receipt, err := client.TransactionReceipt(ctx, arbTx.Hash())
				test.FailIfError(t, err)
				t.Log("Challenger Used", receipt.GasUsed, "gas")
				turn = ethbridge.ASSERTER_TURN
			}
		} else {
			move, err := asserter.HandleConflict(ctx)
			moves = append(moves, move)
			if err != nil {
				t.Logf("Asserter failed challenge: %v", err.Error())
				return moves, err
			}
			arbTx, err := asserterWallet.ExecuteTransactions(ctx, asserterBackend)
			test.FailIfError(t, err)
			client.Commit()
			if arbTx != nil {
				receipt, err := client.TransactionReceipt(ctx, arbTx.Hash())
				test.FailIfError(t, err)
				t.Log("Asserter Used", receipt.GasUsed, "gas")
				turn = ethbridge.CHALLENGER_TURN
			}
		}
		rounds++

		completed, err := tester.ChallengeCompleted(&bind.CallOpts{Context: ctx})
		test.FailIfError(t, err)
		if completed {
			break
		}

		checkTurn(t, challenge, turn)
	}

	checkChallengeCompleted(t, tester, *challengerWallet.Address(), *asserterWallet.Address())
	return moves, nil
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

func initializeChallengeData(t *testing.T, lookup core.ArbCoreLookup, startGas *big.Int, endGas *big.Int) (*core.Assertion, error) {
	cursor, err := lookup.GetExecutionCursor(startGas, true)
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

	err = lookup.AdvanceExecutionCursor(cursor, endGas, true, true)
	test.FailIfError(t, err)
	after, err := core.NewExecutionState(cursor)
	test.FailIfError(t, err)
	if err != nil {
		return nil, err
	}
	return &core.Assertion{
		Before: prevState.ExecutionState,
		After:  after,
	}, nil
}

func gasPrice(tx *types.Transaction, baseFee *big.Int) *big.Int {
	if baseFee == nil {
		return tx.GasPrice()
	}
	return math.BigMin(new(big.Int).Add(tx.GasTipCap(), baseFee), tx.GasFeeCap())
}

func initializeChallengeTest(
	t *testing.T,
	asserterTime *big.Int,
	challengerTime *big.Int,
	arbCore core.ArbCore,
) (*ethutils.SimulatedEthClient, *ethbridgetestcontracts.ChallengeTester, ethcommon.Address, *ethbridge.ValidatorWallet, *ethbridge.ValidatorWallet, func(*core.Assertion), []inbox.InboxMessage) {
	rand.Seed(100000)
	ctx := context.Background()
	clnt, auths := test.SimulatedBackend(t)
	deployer := auths[0]
	rollupAddr := deployer.From
	asserter := auths[1]
	challenger := auths[2]
	sequencer := auths[3]
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
	sequencerBridgeAddr, _, sequencerBridge, err := ethbridgecontracts.DeploySequencerInbox(deployer, client)
	test.FailIfError(t, err)
	client.Commit()

	_, err = delayedBridge.Initialize(deployer)
	test.FailIfError(t, err)
	_, err = sequencerBridge.Initialize(deployer, delayedBridgeAddr, sequencer.From, rollupAddr)
	test.FailIfError(t, err)
	client.Commit()

	_, err = delayedBridge.SetInbox(deployer, deployer.From, true)
	test.FailIfError(t, err)

	_, err = sequencerBridge.SetMaxDelay(deployer, big.NewInt(60), big.NewInt(900))
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
		gasPrice(tx, initBlock.BaseFee()),
		inbox.ChainTime{
			BlockNum:  common.NewTimeBlocks(initBlock.Number()),
			Timestamp: new(big.Int).SetUint64(initBlock.Time()),
		},
	)

	acc, err := delayedBridge.InboxAccs(&bind.CallOpts{}, big.NewInt(0))
	test.FailIfError(t, err)
	delayed := inbox.NewDelayedMessage(common.Hash{}, initMsg)

	if acc != delayed.DelayedAccumulator {
		t.Fatal("unexpected acc in inbox")
	}

	latestHeader, err := client.HeaderByNumber(context.Background(), nil)
	test.FailIfError(t, err)
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocks(latestHeader.Number),
		Timestamp: new(big.Int).SetUint64(latestHeader.Time),
	}

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

	err = core.DeliverMessagesAndWait(ctx, arbCore, big.NewInt(0), common.Hash{}, []inbox.SequencerBatchItem{delayedItem, endOfBlockItem}, []inbox.DelayedMessage{delayed}, nil)
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

	asserterAuth, err := transactauth.NewTransactAuth(ctx, client, asserter)
	test.FailIfError(t, err)
	asserterWallet, err := ethbridge.NewValidator(&asserterWalletAddress, ethcommon.Address{}, ethcommon.Address{}, client, asserterAuth, 0, 1000, nil)
	test.FailIfError(t, err)

	challengerAuth, err := transactauth.NewTransactAuth(ctx, client, challenger)
	test.FailIfError(t, err)
	challengerWallet, err := ethbridge.NewValidator(&challengerWalletAddress, ethcommon.Address{}, ethcommon.Address{}, client, challengerAuth, 0, 1000, nil)
	test.FailIfError(t, err)

	startChallenge := func(assertion *core.Assertion) {
		_, err = tester.StartChallenge(
			deployer,
			assertion.ExecutionHash(),
			assertion.After.TotalMessagesRead,
			*asserterWallet.Address(),
			*challengerWallet.Address(),
			asserterTime,
			challengerTime,
			sequencerBridgeAddr,
			delayedBridgeAddr,
		)
		test.FailIfError(t, err)
		client.Commit()
	}

	messages := []inbox.InboxMessage{initMsg}
	return client, tester, sequencerBridgeAddr, asserterWallet, challengerWallet, startChallenge, messages
}
