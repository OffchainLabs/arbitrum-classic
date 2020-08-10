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

package chainobserver

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
	"math/big"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/chainlistener"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
)

var dbPath = "./testdb"

var rollupTester *ethbridgetestcontracts.RollupTester
var ethclnt *backends.SimulatedBackend
var auth *bind.TransactOpts

func ethTransfer(dest common.Address, amount *big.Int) value.Value {
	ethData := make([]byte, 0)
	ethData = append(ethData, math.U256Bytes(inbox.NewIntFromAddress(dest).BigInt())...)
	ethData = append(ethData, math.U256Bytes(amount)...)
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(0), // ETH type
		inbox.NewIntFromAddress(common.NewAddressFromEth(auth.From)),
		inbox.BytesToByteStack(ethData),
	})
	return tup
}

func checkBalance(t *testing.T, globalInbox arbbridge.GlobalInbox, address common.Address, amount *big.Int) {
	balance, err := globalInbox.GetEthBalance(context.Background(), address)
	if err != nil {
		t.Fatal(err)
	}

	if balance.Cmp(amount) != 0 {
		t.Fatalf("failed checking balance, expected %v but saw %v", amount, balance)
	}
}

func TestMain(m *testing.M) {
	var pks []*ecdsa.PrivateKey
	ethclnt, pks = test.SimulatedBackend()
	auth = bind.NewKeyedTransactor(pks[0])

	go func() {
		t := time.NewTicker(time.Second * 1)
		for range t.C {
			ethclnt.Commit()
		}
	}()

	_, tx, deployedTester, err := ethbridgetestcontracts.DeployRollupTester(
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
		ArbGasSpeedLimitPerTick: 100000,
	}

	arbFactoryAddress, err := ethbridge.DeployRollupFactory(auth, ethclnt)
	if err != nil {
		t.Fatal(err)
	}

	factory, err := clnt.NewArbFactory(common.NewAddressFromEth(arbFactoryAddress))
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

	checkBalance(t, globalInbox, rollupAddress, big.NewInt(100))

	checkpointer, err := checkpointing.NewIndexedCheckpointer(
		rollupAddress,
		dbPath,
		big.NewInt(100000),
		true,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := checkpointer.Initialize(contractPath); err != nil {
		t.Fatal(err)
	}

	chain, err := newChain(
		rollupAddress,
		checkpointer,
		chainParams,
		blockCreated,
	)
	if err != nil {
		t.Fatal(err)
	}
	chain.Inbox = &structures.Inbox{MessageStack: structures.NewRandomMessageStack(100)}

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

	rand.Seed(time.Now().Unix())
	dest := common.RandAddress()
	sends := make([]value.Value, 0)
	sends = append(sends, ethTransfer(dest, big.NewInt(75)))

	assertion := protocol.NewExecutionAssertionFromValues(
		chain.calculatedValidNode.VMProtoData().MachineHash,
		common.RandHash(),
		100,
		6,
		sends,
		[]value.Value{},
	)

	prepared, err := chain.prepareAssertion(chain.latestBlockId)
	if err != nil {
		t.Fatal(err)
	}
	prepared.Assertion = assertion
	prepared.AssertionStub = structures.NewExecutionAssertionStubFromWholeAssertion(assertion, chain.calculatedValidNode.VMProtoData().InboxTop, chain.Inbox.MessageStack)
	var stakerProof []common.Hash
	events, err = chainlistener.MakeAssertion(context.Background(), rollupContract, prepared, stakerProof)
	if err != nil {
		t.Fatal(err)
	}
	for _, ev := range events {
		if err := chain.HandleNotification(context.Background(), ev); err != nil {
			t.Fatal(err)
		}
	}

	latestConf := chain.NodeGraph.LatestConfirmed()
	validNode := chain.NodeGraph.NodeFromHash(latestConf.SuccessorHashes()[valprotocol.ValidChildType])
	if validNode == nil {
		t.Fatal("valid node was nil")
	}
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
			if nd.Disputable().Assertion.LastLogHash != nodeOpp.LogsAcc {
				t.Fatal("incorrect logs acc in proof")
			}

			lastMessageHashOpp := valprotocol.BytesArrayAccumHash(common.Hash{}, nodeOpp.MessagesData, nodeOpp.MessageCount)
			if nd.Disputable().Assertion.LastMessageHash != lastMessageHashOpp {
				t.Log("Assertion", nd.Disputable().Assertion)
				t.Log("nodeOpp.MessagesData", hexutil.Encode(nodeOpp.MessagesData))
				t.Log("nodeOpp.MessageCount", nodeOpp.MessageCount)
				t.Fatal("incorrect messages acc in proof", lastMessageHashOpp, nd.Disputable().Assertion.LastMessageHash)
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
			if messageAccHash != nd.Disputable().Assertion.LastMessageHash {
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

		t.Log(
			latestConf.Hash(),
			proof.InitalProtoStateHash,
			proof.BranchesNums,
			proof.DeadlineTicks,
			proof.ChallengeNodeData,
			proof.LogsAcc,
			proof.VMProtoStateHashes,
			proof.MessageCounts,
			proof.Messages,
		)

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

	checkBalance(t, globalInbox, rollupAddress, big.NewInt(25))
	checkBalance(t, globalInbox, dest, big.NewInt(75))
}
