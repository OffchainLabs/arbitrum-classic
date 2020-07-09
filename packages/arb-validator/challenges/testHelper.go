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
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/gotest"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge/arbfactory"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetest"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"

	errors2 "github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
)

type ChallengeFunc func(common.Address, *ethbridge.EthArbAuthClient, *common.BlockId) (ChallengeState, error)

func testChallenge(
	challengeType valprotocol.ChildType,
	challengeHash [32]byte,
	asserterKey string,
	challengerKey string,
	asserterFunc ChallengeFunc,
	challengerFunc ChallengeFunc,
	testerAddress common.Address,
) error {
	asserterClient, challengerClient, challengeAddress, blockId, err := getChallengeInfo(
		asserterKey,
		challengerKey,
		challengeType,
		challengeHash,
		testerAddress,
	)
	if err != nil {
		return errors2.Wrap(err, "Error starting challenge")
	}

	asserterEndChan := make(chan ChallengeState)
	asserterErrChan := make(chan error)
	challengerEndChan := make(chan ChallengeState)
	challengerErrChan := make(chan error)

	go func() {
		cBlockId := blockId.MarshalToBuf().Unmarshal()
		tryCount := 0
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
		tryCount := 0
		for {
			endState, err := challengerFunc(challengeAddress, challengerClient, cBlockId)
			if err == nil {
				asserterEndChan <- endState
				return
			}
			if tryCount > 20 {
				asserterErrChan <- err
				return
			}
			tryCount += 1
			log.Println("Restarting challenger", err)
			cBlockId, err = asserterClient.BlockIdForHeight(context.Background(), cBlockId.Height)
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
				return fmt.Errorf("Asserter Ended: Asserter challenge ended with %v", challengeState)
			}
			doneCount++
			if doneCount == 2 {
				return nil
			}
		case challengeState := <-challengerEndChan:
			if challengeState != ChallengeAsserterWon {
				return fmt.Errorf("Challenger Ended: Asserter challenge ended with %v", challengeState)
			}
			doneCount++
			if doneCount == 2 {
				return nil
			}
		case err := <-asserterErrChan:
			return errors2.Wrap(err, "Asserter error")
		case err := <-challengerErrChan:
			return errors2.Wrap(err, "Challenger error")
		case <-time.After(80 * time.Second):
			return errors.New("Challenge never completed")
		}
	}
}

func launchChallengeTester(
	key string,
) (common.Address, error) {
	auth, err := test.SetupAuth(key)
	if err != nil {
		return common.Address{}, err
	}
	ethclint, err := ethclient.Dial(test.GetEthUrl())
	if err != nil {
		return common.Address{}, err
	}

	connectionInfo, err := getConnectionInfo()
	if err != nil {
		return common.Address{}, err
	}

	factory, err := arbfactory.NewArbFactory(connectionInfo.ArbFactoryAddress().ToEthAddress(), ethclint)
	if err != nil {
		return common.Address{}, err
	}

	challengeFactoryAddress, err := factory.ChallengeFactoryAddress(nil)
	if err != nil {
		return common.Address{}, errors2.Wrap(err, "Error getting challenge factory address")
	}

	tester, err := ethbridgetest.DeployChallengeTest(context.Background(), ethclint, auth, common.NewAddressFromEth(challengeFactoryAddress))
	if err != nil {
		return common.Address{}, errors2.Wrap(err, "Error deploying challenge")
	}

	return common.NewAddressFromEth(tester.Address), nil
}

func getConnectionInfo() (ethbridge.ArbAddresses, error) {
	bridge_eth_addresses := "../bridge_eth_addresses.json"
	var connectionInfo ethbridge.ArbAddresses

	jsonFile, err := os.Open(bridge_eth_addresses)
	if err != nil {
		return connectionInfo, err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := jsonFile.Close(); err != nil {
		return connectionInfo, err
	}

	if err := json.Unmarshal(byteValue, &connectionInfo); err != nil {
		return connectionInfo, err
	}

	return connectionInfo, nil
}

func getAuth() (*ethclient.Client, *ethclient.Client, error) {
	ethURL := test.GetEthUrl()
	seed := time.Now().UnixNano()

	fmt.Println("seed", seed)
	rand.Seed(seed)

	ethclint1, err := ethclient.Dial(ethURL)
	if err != nil {
		return nil, nil, err
	}

	ethclint2, err := ethclient.Dial(ethURL)
	if err != nil {
		return nil, nil, err
	}

	return ethclint1, ethclint2, nil
}

func getTestMachine(t *testing.T) machine.Machine {
	mach, err := loader.LoadMachineFromFile(gotest.TestMachinePath(), true, "cpp")
	if err != nil {
		t.Fatal("Loader Error: ", err)
	}

	return mach
}

func getChallengeInfo(
	asserterKey string,
	challengerKey string,
	challengeType valprotocol.ChildType,
	challengeHash [32]byte,
	testerAddress common.Address,
) (*ethbridge.EthArbAuthClient, *ethbridge.EthArbAuthClient, common.Address, *common.BlockId, error) {
	auth1, err := test.SetupAuth(asserterKey)
	if err != nil {
		return nil, nil, common.Address{}, nil, err
	}
	auth2, err := test.SetupAuth(challengerKey)
	if err != nil {
		return nil, nil, common.Address{}, nil, err
	}

	ethclint1, ethclint2, err := getAuth()

	asserterClient := ethbridge.NewEthAuthClient(ethclint1, auth1)
	challengerClient := ethbridge.NewEthAuthClient(ethclint2, auth2)

	tester, err := ethbridgetest.NewChallengeTester(testerAddress.ToEthAddress(), ethclint1, auth1)
	if err != nil {
		return nil, nil, common.Address{}, nil, err
	}

	challengeAddress, blockId, err := tester.StartChallenge(
		context.Background(),
		asserterClient.Address(),
		challengerClient.Address(),
		common.TicksFromBlockNum(common.NewTimeBlocksInt(5)),
		challengeHash,
		new(big.Int).SetUint64(uint64(challengeType)),
	)
	if err != nil {
		return nil, nil, common.Address{}, nil, errors2.Wrap(err, "Error starting challenge")
	}

	return asserterClient, challengerClient, challengeAddress, blockId, nil
}
