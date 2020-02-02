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

	errors2 "github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/test"
)

type ChallengeFunc func(common.Address, *ethbridge.EthArbAuthClient, *structures.BlockID) (ChallengeState, error)

func testChallenge(
	challengeType structures.ChildType,
	challengeHash [32]byte,
	asserterKey, challengerKey string,
	asserterFunc, challengerFunc ChallengeFunc,
) error {
	bridgeEthAddresses := "../bridge_eth_addresses.json"
	ethURL := test.GetEthURL()
	seed := time.Now().UnixNano()
	// seed := int64(1559616168133477000)
	fmt.Println("seed", seed)
	rand.Seed(seed)
	jsonFile, err := os.Open(bridgeEthAddresses)
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

	auth1, err := test.SetupAuth(asserterKey)
	if err != nil {
		return err
	}
	auth2, err := test.SetupAuth(challengerKey)
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
		return errors2.Wrap(err, "Error getting challenge factory address")
	}

	tester, err := client1.DeployChallengeTest(context.Background(), challengeFactoryAddress)
	if err != nil {
		return errors2.Wrap(err, "Error deploying challenge")
	}

	challengeAddress, blockID, err := tester.StartChallenge(
		context.Background(),
		client1.Address(),
		client2.Address(),
		common.TimeFromBlockNum(common.NewTimeBlocksInt(5)),
		challengeHash,
		new(big.Int).SetUint64(uint64(challengeType)),
	)
	if err != nil {
		return errors2.Wrap(err, "Error starting challenge")
	}

	asserterEndChan := make(chan ChallengeState)
	asserterErrChan := make(chan error)
	challengerEndChan := make(chan ChallengeState)
	challengerErrChan := make(chan error)

	go func() {
		cBlockID := blockID.MarshalToBuf().Unmarshal()
		tryCount := 0
		for {
			endState, err := asserterFunc(challengeAddress, client1, cBlockID)
			if err == nil {
				asserterEndChan <- endState
				return
			}
			if tryCount > 20 {
				asserterErrChan <- err
				return
			}
			tryCount++
			log.Println("Restarting asserter", err)
			cBlockID, err = client1.BlockIDForHeight(context.Background(), cBlockID.Height)
			if err != nil {
				asserterErrChan <- err
				return
			}
		}
	}()

	go func() {
		cBlockID := blockID.MarshalToBuf().Unmarshal()
		tryCount := 0
		for {
			endState, err := challengerFunc(challengeAddress, client2, cBlockID)
			if err == nil {
				asserterEndChan <- endState
				return
			}
			if tryCount > 20 {
				asserterErrChan <- err
				return
			}
			tryCount++
			log.Println("Restarting challenger", err)
			cBlockID, err = client1.BlockIDForHeight(context.Background(), cBlockID.Height)
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
				return fmt.Errorf("asserter challenge ended with %v", challengeState)
			}
			doneCount++
			if doneCount == 2 {
				return nil
			}
		case challengeState := <-challengerEndChan:
			if challengeState != ChallengeAsserterWon {
				return fmt.Errorf("asserter challenge ended with %v", challengeState)
			}
			doneCount++
			if doneCount == 2 {
				return nil
			}
		case err := <-asserterErrChan:
			return errors2.Wrap(err, "asserter error")
		case err := <-challengerErrChan:
			return errors2.Wrap(err, "challenger error")
		case <-time.After(80 * time.Second):
			return errors.New("challenge never completed")
		}
	}
}
