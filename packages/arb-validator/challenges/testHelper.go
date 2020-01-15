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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"os"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/test"
)

type ChallengeFunc func(common.Address, *ethbridge.EthArbAuthClient, *structures.BlockId) (ChallengeState, error)

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

	auth1, err := test.SetupAuth("ffb2b26161e081f0cdf9db67200ee0ce25499d5ee683180a9781e6cceb791c39")
	if err != nil {
		return err
	}
	auth2, err := test.SetupAuth("979f020f6f6f71577c09db93ba944c89945f10fade64cfc7eb26137d5816fb76")
	if err != nil {
		return err
	}

	client1, err := ethbridge.NewEthAuthClient(ethURL, auth1)
	if err != nil {
		return err
	}
	client2, err := ethbridge.NewEthAuthClient(ethURL, auth2)
	if err != nil {
		return err
	}

	factory, err := client1.NewArbFactoryWatcher(connectionInfo.ArbFactoryAddress())
	if err != nil {
		return err
	}

	challengeFactoryAddress, err := factory.ChallengeFactoryAddress()
	if err != nil {
		return err
	}

	tester, err := client1.DeployChallengeTest()
	if err != nil {
		return err
	}

	challengeAddress, blockId, err := tester.StartChallenge(
		context.Background(),
		challengeFactoryAddress,
		client1.Address(),
		client2.Address(),
		common.TimeTicks{big.NewInt(13000 * 5)},
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
		cBlockId := blockId.MarshalToBuf().Unmarshal()
		tryCount := 0
		for {
			endState, err := asserterFunc(challengeAddress, client1, cBlockId)
			if err == nil {
				asserterEndChan <- endState
				return
			}
			if tryCount > 5 {
				asserterErrChan <- err
				return
			}
			tryCount += 1
			log.Println("Restarting asserter")
			cBlockId, err = client1.BlockIdForHeight(context.Background(), cBlockId.Height)
			if err != nil {
				asserterErrChan <- err
				return
			}
		}

	}()

	go func() {
		cBlockId := blockId.MarshalToBuf().Unmarshal()
		tryCount := 0
		for {
			endState, err := challengerFunc(challengeAddress, client2, cBlockId)
			if err == nil {
				asserterEndChan <- endState
				return
			}
			if tryCount > 5 {
				asserterErrChan <- err
				return
			}
			tryCount += 1
			log.Println("Restarting challenger")
			cBlockId, err = client1.BlockIdForHeight(context.Background(), cBlockId.Height)
			if err != nil {
				asserterErrChan <- err
				return
			}
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
