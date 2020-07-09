/*
 * Copyright 2020, Offchain Labs, Inc.
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

package rollup

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetest/rolluptester"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup/chainlistener"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup/chainobserver"
	"log"
	"math/big"
	"math/rand"
	"os"
	"testing"
	"time"
)

var dbPath = "./testdb"

var rollupTester *rolluptester.RollupTester
var ethclnt *ethclient.Client
var auth *bind.TransactOpts

func TestMain(m *testing.M) {
	var err error
	ethclnt, err = ethclient.Dial(test.GetEthUrl())
	if err != nil {
		log.Fatal(err)
	}

	auth, err = test.SetupAuth("8285795ed740b32384e72554128f80714ed6c93d50aeb18aa655547431a7a3cb")
	if err != nil {
		log.Fatal(err)
	}

	_, tx, deployedTester, err := rolluptester.DeployRollupTester(
		auth,
		ethclnt,
	)

	if err != nil {
		log.Fatal(err)
	}
	_, err = ethbridge.WaitForReceiptWithResults(
		context.Background(),
		ethclnt,
		auth.From,
		tx,
		"DeployRollupTester",
	)
	rollupTester = deployedTester

	code := m.Run()
	if err := os.RemoveAll(dbPath); err != nil {
		log.Fatal(err)
	}
	os.Exit(code)
}

func TestConfirmAssertion(t *testing.T) {
	clnt := ethbridge.NewEthAuthClient(ethclnt, auth)

	chainParams := valprotocol.ChainParams{
		StakeRequirement:        big.NewInt(0),
		GracePeriod:             common.TicksFromSeconds(1),
		MaxExecutionSteps:       100000,
		MaxBlockBoundsWidth:     10000,
		ArbGasSpeedLimitPerTick: 100000,
	}

	arbFactoryAddress, err := test.GetFactoryAddress()
	if err != nil {
		t.Fatal(err)
	}

	factory, err := clnt.NewArbFactory(arbFactoryAddress)
	if err != nil {
		t.Fatal(err)
	}

	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	rollupAddress, blockCreated, err := factory.CreateRollup(
		context.Background(),
		mach.Hash(),
		chainParams,
		common.Address{},
	)
	if err != nil {
		t.Fatal(err)
	}

	rollupContract, err := clnt.NewRollup(rollupAddress)
	if err != nil {
		t.Fatal(err)
	}

	inboxAddress, err := rollupContract.InboxAddress(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	globalInbox, err := clnt.NewGlobalInbox(inboxAddress, rollupAddress)
	if err != nil {
		t.Fatal(err)
	}

	if err := globalInbox.DepositEthMessage(
		context.Background(),
		rollupAddress,
		common.NewAddressFromEth(auth.From),
		big.NewInt(100),
	); err != nil {
		t.Fatal(err)
	}

	checkBalance := func(address common.Address, amount *big.Int) {
		balance, err := globalInbox.GetEthBalance(context.Background(), address)
		if err != nil {
			t.Fatal(err)
		}

		if balance.Cmp(amount) != 0 {
			t.Fatal("failed to deposit balance")
		}
	}

	checkBalance(rollupAddress, big.NewInt(100))

	checkpointer := checkpointing.NewIndexedCheckpointer(
		rollupAddress,
		dbPath,
		big.NewInt(100000),
		true,
	)

	if err := checkpointer.Initialize(contractPath); err != nil {
		t.Fatal(err)
	}

	chain, err := chainobserver.NewChain(
		rollupAddress,
		checkpointer,
		chainParams,
		false,
		blockCreated,
		common.Hash{},
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("place stake", func(t *testing.T) {
		events, err := rollupContract.PlaceStake(
			context.Background(),
			big.NewInt(0),
			[]common.Hash{},
			[]common.Hash{},
		)
		if err != nil {
			t.Fatal(err)
		}
		for _, ev := range events {
			if err := chain.HandleNotification(context.Background(), ev); err != nil {
				t.Fatal(err)
			}
		}
	})

	rand.Seed(time.Now().Unix())
	dest := common.RandAddress()
	results := make([]evm.Result, 0, 5)
	messages := make([]value.Value, 0)
	messages = append(messages, message.Eth{
		To:    dest,
		From:  common.NewAddressFromEth(auth.From),
		Value: big.NewInt(75),
	}.AsInboxValue())
	for i := int32(0); i < 5; i++ {
		stop := evm.NewRandomStop(message.NewRandomEth(), 2)
		results = append(results, stop)
		messages = append(messages, message.NewRandomEth().AsInboxValue())
	}

	assertion := evm.NewRandomEVMAssertion(results, messages)
	assertion.NumGas = 100
	prepared := chain.PrepareAssertion()
	prepared.Assertion = assertion
	prepared.Claim.AssertionStub = valprotocol.NewExecutionAssertionStubFromAssertion(assertion)
	t.Run("make assertion", func(t *testing.T) {
		var stakerProof []common.Hash
		events, err := chainlistener.MakeAssertion(context.Background(), rollupContract, prepared, stakerProof)
		if err != nil {
			t.Fatal(err)
		}
		for _, ev := range events {
			if err := chain.HandleNotification(context.Background(), ev); err != nil {
				t.Fatal(err)
			}
		}
	})

	latestConf := chain.NodeGraph.LatestConfirmed()
	validNode := chain.NodeGraph.NodeFromHash(latestConf.SuccessorHashes()[3])
	if err := validNode.UpdateValidOpinion(prepared.Machine, prepared.Assertion); err != nil {
		t.Fatal(err)
	}

	time.Sleep(2 * time.Second)

	currentTime, err := clnt.CurrentBlockId(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	confTime := new(big.Int).Add(currentTime.Height.AsInt(), big.NewInt(1))

	t.Run("confirm assertion", func(t *testing.T) {
		opp, nodes := chain.NodeGraph.GenerateNextConfProof(common.TicksFromBlockNum(common.NewTimeBlocks(confTime)))
		if opp == nil {
			t.Fatal("should have had opp")
		}
		proof := opp.PrepareProof()
		offset := big.NewInt(0)
		validCount := int64(0)
		for i, nodeOpp := range opp.Nodes {
			nd := nodes[i]
			nodeOpp, ok := nodeOpp.(valprotocol.ConfirmValidOpportunity)
			if !ok {
				continue
			}
			if nd.Disputable().AssertionClaim.AssertionStub.LastLogHash != nodeOpp.LogsAcc {
				t.Fatal("incorrect logs acc in proof")
			}

			if nd.Disputable().AssertionClaim.AssertionStub.LastMessageHash != valprotocol.BytesArrayAccumHash(nodeOpp.MessagesData, nodeOpp.MessageCount) {
				t.Fatal("incorrect messages acc in proof")
			}
			messageAccHash, nextOffset, err := rollupTester.GenerateLastMessageHash(
				nil,
				proof.Messages,
				offset,
				proof.MessageCounts[validCount],
			)
			if err != nil {
				t.Fatal(err)
			}
			if messageAccHash != nd.Disputable().AssertionClaim.AssertionStub.LastMessageHash {
				t.Fatal("generated incorrect messages acc")
			}

			_, nodeDataHash, vmProtoStateHash, err := rollupTester.ProcessValidNode(
				nil,
				proof.InitalProtoStateHash,
				proof.BranchesNums,
				proof.DeadlineTicks,
				proof.ChallengeNodeData,
				proof.LogsAcc,
				proof.VMProtoStateHashes,
				proof.MessageCounts,
				proof.Messages,
				big.NewInt(validCount),
				offset,
			)

			if vmProtoStateHash != nodeOpp.VMProtoStateHash {
				t.Error("incorrect state hash")
			}

			if nodeDataHash != nd.NodeDataHash() {
				t.Error("incorrect data hash")
			}

			offset = nextOffset
			validCount++
		}

		ret, err := rollupTester.Confirm(
			nil,
			latestConf.Hash().ToEthHash(),
			proof.InitalProtoStateHash,
			proof.BranchesNums,
			proof.DeadlineTicks,
			proof.ChallengeNodeData,
			proof.LogsAcc,
			proof.VMProtoStateHashes,
			proof.MessageCounts,
			proof.Messages,
		)

		if err != nil {
			t.Fatal(err)
		}
		if len(ret.ValidNodeHashes) != 1 {
			t.Fatal("wrong valid node count")
		}
		if ret.LastNode != validNode.Hash() {
			t.Fatalf("incorrect last node hash: was %v but should have been %v", hexutil.Encode(ret.LastNode[:]), validNode.Hash())
		}
		if ret.ValidNodeHashes[0] != validNode.Hash() {
			t.Fatal("wrong node hash")
		}
		events, err := rollupContract.Confirm(context.Background(), opp)
		if err != nil {
			t.Fatal(err)
		}
		for _, ev := range events {
			if err := chain.HandleNotification(context.Background(), ev); err != nil {
				t.Fatal(err)
			}
		}
	})

	checkBalance(rollupAddress, big.NewInt(25))
	checkBalance(dest, big.NewInt(75))
}
