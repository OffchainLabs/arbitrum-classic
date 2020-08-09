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

package challenges

import (
	"context"
	errors2 "github.com/pkg/errors"
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
)

type ChallengeFunc func(common.Address, *ethbridge.EthArbAuthClient, *common.BlockId) (ChallengeState, error)

func testChallengerCatchUp(
	t *testing.T,
	client ethutils.EthClient,
	asserter *bind.TransactOpts,
	challenger *bind.TransactOpts,
	challengeType valprotocol.ChildType,
	challengeHash [32]byte,
	asserterFunc ChallengeFunc,
	asserterFuncStop ChallengeFunc,
	challengerFunc ChallengeFunc,
	challengerFuncStop ChallengeFunc,
	testerAddress ethcommon.Address,
) {
	current, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	blockId := &common.BlockId{
		Height:     common.NewTimeBlocks(current.Number),
		HeaderHash: common.NewHashFromEth(current.Hash()),
	}
	asserterClient, challengerClient, challengeAddress, _, err := getChallengeInfo(
		client,
		asserter,
		challenger,
		challengeType,
		challengeHash,
		testerAddress,
	)
	if err != nil {
		t.Fatal("Error starting challenge", err)
	}

	asserterEndChan := make(chan ChallengeState)
	asserterErrChan := make(chan error)
	challengerEndChan := make(chan ChallengeState)
	challengerErrChan := make(chan error)

	go func() {
		cBlockId := blockId.MarshalToBuf().Unmarshal()
		tryCount := 0
		for {
			endState, err := asserterFuncStop(challengeAddress, asserterClient, cBlockId)
			if endState == DefenderDiscontinued {
				break
			}
			if tryCount > 20 {
				asserterErrChan <- err
				return
			}
			tryCount += 1
			log.Println("Restarting asserter", err)
			cBlockId, err = asserterClient.BlockIdForHeight(context.Background(), cBlockId.Height)
			if err != nil {
				asserterErrChan <- err
				return
			}
		}
		for {
			endState, err := asserterFunc(challengeAddress, asserterClient, cBlockId)
			if err == nil {
				asserterEndChan <- endState
				return
			}
			if tryCount > 20 {
				asserterErrChan <- err
				return
			}
			tryCount += 1
			log.Println("Restarting asserter", err)
			cBlockId, err = asserterClient.BlockIdForHeight(context.Background(), cBlockId.Height)
			if err != nil {
				asserterErrChan <- err
				return
			}
		}
	}()

	go func() {
		cBlockId := blockId.MarshalToBuf().Unmarshal()
		for {
			endState, err := challengerFuncStop(challengeAddress, challengerClient, cBlockId)
			if endState == ChallengerDiscontinued {
				break
			}
			asserterErrChan <- err
			return
		}
		for {
			endState, err := challengerFunc(challengeAddress, challengerClient, cBlockId)
			if err == nil {
				asserterEndChan <- endState
				return
			}
			asserterErrChan <- err
			return
		}
	}()

	resolveChallenge(t, asserterEndChan, asserterErrChan, challengerEndChan, challengerErrChan)
}

func testChallenge(
	t *testing.T,
	client ethutils.EthClient,
	asserter *bind.TransactOpts,
	challenger *bind.TransactOpts,
	challengeType valprotocol.ChildType,
	challengeHash [32]byte,
	asserterFunc ChallengeFunc,
	challengerFunc ChallengeFunc,
	testerAddress ethcommon.Address,
) {
	current, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	blockId := &common.BlockId{
		Height:     common.NewTimeBlocks(current.Number),
		HeaderHash: common.NewHashFromEth(current.Hash()),
	}
	asserterClient, challengerClient, challengeAddress, _, err := getChallengeInfo(
		client,
		asserter,
		challenger,
		challengeType,
		challengeHash,
		testerAddress,
	)
	if err != nil {
		t.Fatal("Error starting challenge", err)
	}

	asserterEndChan := make(chan ChallengeState)
	asserterErrChan := make(chan error)
	challengerEndChan := make(chan ChallengeState)
	challengerErrChan := make(chan error)

	go func() {
		for {
			endState, err := asserterFunc(challengeAddress, asserterClient, blockId)
			if err == nil {
				asserterEndChan <- endState
				return
			}
			asserterErrChan <- err
			return
		}
	}()

	go func() {
		cBlockId := blockId.MarshalToBuf().Unmarshal()
		for {
			endState, err := challengerFunc(challengeAddress, challengerClient, cBlockId)
			if err == nil {
				asserterEndChan <- endState
				return
			}
			asserterErrChan <- err
			return
		}
	}()

	resolveChallenge(t, asserterEndChan, asserterErrChan, challengerEndChan, challengerErrChan)
}

func resolveChallenge(
	t *testing.T,
	asserterEndChan chan ChallengeState,
	asserterErrChan chan error,
	challengerEndChan chan ChallengeState,
	challengerErrChan chan error) {
	doneCount := 0
	for {
		select {
		case challengeState := <-asserterEndChan:
			if challengeState != ChallengeAsserterWon {
				t.Fatalf("Asserter Ended: Asserter challenge ended with %v", challengeState)
			}
			doneCount++
			if doneCount == 2 {
				return
			}
		case challengeState := <-challengerEndChan:
			if challengeState != ChallengeAsserterWon {
				t.Fatalf("Challenger Ended: Asserter challenge ended with %v", challengeState)
			}
			doneCount++
			if doneCount == 2 {
				return
			}
		case err := <-asserterErrChan:
			t.Fatal("Asserter error", err)
		case err := <-challengerErrChan:
			t.Fatal("Challenger error", err)
		case <-time.After(80 * time.Second):
			t.Fatal("Challenge never completed")
		}
	}
}

func getTestMachine(t *testing.T) machine.Machine {
	mach, err := loader.LoadMachineFromFile(arbos.Path(), true, "cpp")
	if err != nil {
		t.Fatal("Loader Error: ", err)
	}

	return mach
}

func getChallengeInfo(
	client ethutils.EthClient,
	asserter *bind.TransactOpts,
	challenger *bind.TransactOpts,
	challengeType valprotocol.ChildType,
	challengeHash [32]byte,
	testerAddress ethcommon.Address,
) (*ethbridge.EthArbAuthClient, *ethbridge.EthArbAuthClient, common.Address, *common.BlockId, error) {
	asserterClient := ethbridge.NewEthAuthClient(client, asserter)
	challengerClient := ethbridge.NewEthAuthClient(client, challenger)

	tester, err := ethbridgetestcontracts.NewChallengeTester(testerAddress, client)
	if err != nil {
		return nil, nil, common.Address{}, nil, err
	}

	tx, err := tester.StartChallenge(
		asserter,
		asserterClient.Address().ToEthAddress(),
		challengerClient.Address().ToEthAddress(),
		common.TicksFromBlockNum(common.NewTimeBlocksInt(5)).Val,
		challengeHash,
		new(big.Int).SetUint64(uint64(challengeType)),
	)
	if err != nil {
		return nil, nil, common.Address{}, nil, errors2.Wrap(err, "Error starting challenge")
	}

	receipt, err := ethbridge.WaitForReceiptWithResults(context.Background(), client, asserter.From, tx, "StartChallenge")
	if err != nil {
		return nil, nil, common.Address{}, nil, errors2.Wrap(err, "Error starting challenge")
	}

	if len(receipt.Logs) != 1 {
		return nil, nil, common.Address{}, nil, errors2.Wrap(err, "Error starting challenge")
	}

	challengeAddress := common.NewAddressFromEth(receipt.Logs[0].Address)
	blockId := ethbridge.GetReceiptBlockID(receipt)

	log.Println("Started challenge at block", blockId)

	return asserterClient, challengerClient, challengeAddress, blockId, nil
}
