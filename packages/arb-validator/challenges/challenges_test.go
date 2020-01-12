/*
 * Copyright 2019, Offchain Labs, Inc.
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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/challengetester"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/test"
)

type ChallengeFunc func(common.Address, *ethbridge.EthArbAuthClient) (ChallengeState, error)

func testChallenge(
	challengeType structures.ChildType,
	challengeHash [32]byte,
	asserterFunc, challengerFunc ChallengeFunc,
) error {
	bridge_eth_addresses := "../bridge_eth_addresses.json"
	ethURL := test.GetEthUrl()
	seed := time.Now().UnixNano()
	// seed := int64(1559616168133477000)
	fmt.Println("seed", seed)
	rand.Seed(seed)
	jsonFile, err := os.Open(bridge_eth_addresses)
	if err != nil {
		return err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := jsonFile.Close(); err != nil {
		return err
	}

	var connectionInfo ethbridge.ArbAddresses
	if err := json.Unmarshal(byteValue, &connectionInfo); err != nil {
		return err
	}

	key1, err := crypto.HexToECDSA("ffb2b26161e081f0cdf9db67200ee0ce25499d5ee683180a9781e6cceb791c39")
	if err != nil {
		return err
	}
	key2, err := crypto.HexToECDSA("979f020f6f6f71577c09db93ba944c89945f10fade64cfc7eb26137d5816fb76")
	if err != nil {
		return err
	}

	auth1 := bind.NewKeyedTransactor(key1)
	auth2 := bind.NewKeyedTransactor(key2)

	client1, err := ethbridge.NewEthAuthClient(ethURL, auth1)
	if err != nil {
		return err
	}
	client2, err := ethbridge.NewEthAuthClient(ethURL, auth2)
	if err != nil {
		return err
	}

	factory, err := client1.NewArbFactoryWatcher(common.HexToAddress(connectionInfo.ArbFactory))
	if err != nil {
		return err
	}

	challengeFactoryAddress, err := factory.ChallengeFactoryAddress()
	if err != nil {
		return err
	}

	testerAddress, tx, _, err := challengetester.DeployChallengeTester(auth1, client1.Client)
	if err != nil {
		return err
	}

	_, err = ethbridge.WaitForReceiptWithResults(
		context.Background(),
		client1.Client,
		auth1.From,
		tx,
		"DeployChallengeTester",
	)
	if err != nil {
		return err
	}
	tester, err := ethbridge.NewChallengeTester(testerAddress, client1.Client, auth1)
	if err != nil {
		return err
	}

	challengeAddress, err := tester.StartChallenge(
		context.Background(),
		challengeFactoryAddress,
		client1.Address(),
		client2.Address(),
		structures.TimeTicks{big.NewInt(13000 * 5)},
		challengeHash,
		new(big.Int).SetUint64(uint64(challengeType)),
	)
	if err != nil {
		return err
	}

	asserterEndChan := make(chan ChallengeState)
	asserterErrChan := make(chan error)
	challengerEndChan := make(chan ChallengeState)
	challengerErrChan := make(chan error)

	go func() {
		endState, err := asserterFunc(challengeAddress, client1)
		if err != nil {
			asserterErrChan <- err
		} else {
			asserterEndChan <- endState
		}
	}()

	go func() {
		endState, err := challengerFunc(challengeAddress, client2)
		if err != nil {
			challengerErrChan <- err
		} else {
			challengerEndChan <- endState
		}
	}()

	doneCount := 0
	for {
		select {
		case challengeState := <-asserterEndChan:
			if challengeState != ChallengeAsserterWon {
				return fmt.Errorf("Asserter challenge ended with %v", challengeState)
			}
			doneCount++
			if doneCount == 2 {
				return nil
			}
		case challengeState := <-challengerEndChan:
			if challengeState != ChallengeAsserterWon {
				return fmt.Errorf("Asserter challenge ended with %v", challengeState)
			}
			doneCount++
			if doneCount == 2 {
				return nil
			}
		case err := <-asserterErrChan:
			return err
		case err := <-challengerErrChan:
			return err
		case <-time.After(60 * time.Second):
			return errors.New("Challenge never completed")
		}
	}
}

func TestPendingTopChallenge(t *testing.T) {
	messageStack := structures.NewMessageStack()
	messageStack.DeliverMessage(value.NewInt64Value(0))
	messageStack.DeliverMessage(value.NewInt64Value(1))
	messageStack.DeliverMessage(value.NewInt64Value(2))
	messageStack.DeliverMessage(value.NewInt64Value(3))

	bottomHash, err := messageStack.GetHashAtIndex(big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	topHash, err := messageStack.GetHashAtIndex(big.NewInt(3))
	if err != nil {
		t.Fatal(err)
	}
	challengeHash := structures.PendingTopChallengeDataHash(bottomHash, topHash, big.NewInt(3))

	if err := testChallenge(
		structures.InvalidPendingChildType,
		challengeHash,
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient) (ChallengeState, error) {
			return DefendPendingTopClaim(
				client,
				challengeAddress,
				2,
				messageStack,
				bottomHash,
				topHash,
			)
		},
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient) (ChallengeState, error) {
			return ChallengePendingTopClaim(client, challengeAddress, messageStack)
		},
	); err != nil {
		t.Fatal(err)
	}
}

func TestMessagesChallenge(t *testing.T) {
	messageStack := structures.NewMessageStack()
	for i := int64(0); i < 8; i++ {
		messageStack.DeliverMessage(value.NewInt64Value(i))
	}
	beforePending, err := messageStack.GetHashAtIndex(big.NewInt(2))
	if err != nil {
		t.Fatal(err)
	}
	afterPending, err := messageStack.GetHashAtIndex(big.NewInt(6))
	if err != nil {
		t.Fatal(err)
	}

	substack, err := messageStack.Substack(beforePending, afterPending)
	if err != nil {
		t.Fatal(err)
	}

	importedMessages := substack.GetTopHash()
	challengeHash := structures.MessageChallengeDataHash(
		beforePending,
		afterPending,
		value.NewEmptyTuple().Hash(),
		importedMessages,
		big.NewInt(4),
	)

	if err := testChallenge(
		structures.InvalidMessagesChildType,
		challengeHash,
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient) (ChallengeState, error) {
			return DefendMessagesClaim(
				client,
				challengeAddress,
				2,
				messageStack,
				beforePending,
				afterPending,
				importedMessages,
			)
		},
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient) (ChallengeState, error) {
			return ChallengeMessagesClaim(
				client,
				challengeAddress,
				messageStack,
				beforePending,
				afterPending,
			)
		},
	); err != nil {
		t.Fatal(err)
	}
}

func TestExecution(t *testing.T) {
	contract := "../contract.ao"

	mach, err := loader.LoadMachineFromFile(contract, true, "test")
	if err != nil {
		t.Fatal("Loader Error: ", err)
	}

	timeBounds := &protocol.TimeBoundsBlocks{
		protocol.NewTimeBlocks(big.NewInt(100)),
		protocol.NewTimeBlocks(big.NewInt(200)),
	}
	afterMachine := mach.Clone()
	precondition := valprotocol.NewPrecondition(mach.Hash(), timeBounds, value.NewEmptyTuple())
	assertion, numSteps := afterMachine.ExecuteAssertion(1000, timeBounds, value.NewEmptyTuple())

	challengeHash := structures.ExecutionDataHash(
		numSteps,
		precondition.Hash(),
		valprotocol.NewExecutionAssertionStubFromAssertion(assertion).Hash(),
	)

	if err := testChallenge(
		structures.InvalidExecutionChildType,
		challengeHash,
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient) (ChallengeState, error) {
			return DefendExecutionClaim(
				client,
				challengeAddress,
				2,
				precondition,
				numSteps,
				mach.Clone(),
			)
		},
		func(challengeAddress common.Address, client *ethbridge.EthArbAuthClient) (ChallengeState, error) {
			return ChallengeExecutionClaim(
				client,
				challengeAddress,
				precondition,
				mach.Clone(),
				true,
			)
		},
	); err != nil {
		t.Fatal(err)
	}
}
