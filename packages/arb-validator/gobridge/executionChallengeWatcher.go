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

package gobridge

import (
	"context"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"strings"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/executionchallenge"
)

var bisectedAssertionID ethcommon.Hash
var oneStepProofCompletedID ethcommon.Hash

func init() {
	parsed, err := abi.JSON(strings.NewReader(executionchallenge.ExecutionChallengeABI))
	if err != nil {
		panic(err)
	}
	bisectedAssertionID = parsed.Events["BisectedAssertion"].ID()
	oneStepProofCompletedID = parsed.Events["OneStepProofCompleted"].ID()
}

type executionChallengeWatcher struct {
	*bisectionChallengeWatcher
	challengeInfo *challengeData
	client        *GoArbClient
}

func newExecutionChallengeWatcher(address common.Address, client *GoArbClient) (*executionChallengeWatcher, error) {
	fmt.Println("in newExecutionChallengeWatcher")
	bisectionChallenge, err := newBisectionChallengeWatcher(address, client)
	if err != nil {
		return nil, err
	}

	chalData := client.GoEthClient.challenges[address]
	client.GoEthClient.challengeWatchersMutex.Lock()
	if _, ok := client.GoEthClient.challengeWatcherEvents[chalData]; !ok {
		client.GoEthClient.challengeWatcherEvents[chalData] = make(map[*structures.BlockId][]arbbridge.Event)
	}
	client.GoEthClient.challengeWatchersMutex.Unlock()

	return &executionChallengeWatcher{bisectionChallengeWatcher: bisectionChallenge, challengeInfo: chalData, client: client}, nil
}

func (c *executionChallengeWatcher) GetEvents(ctx context.Context, blockId *structures.BlockId) ([]arbbridge.Event, error) {
	c.client.GoEthClient.challengeWatchersMutex.Lock()
	cw := c.client.GoEthClient.challengeWatcherEvents[c.challengeInfo][blockId]
	c.client.GoEthClient.challengeWatchersMutex.Unlock()
	return cw, nil
}

//func (c *executionChallengeWatcher) parseExecutionEvent(chainInfo arbbridge.ChainInfo, log types.Log) (arbbridge.Event, error) {
//	if log.Topics[0] == bisectedAssertionID {
//		bisectChal, err := c.challenge.ParseBisectedAssertion(log)
//		if err != nil {
//			return nil, err
//		}
//		bisectionCount := len(bisectChal.MachineHashes) - 1
//		assertions := make([]*valprotocol.ExecutionAssertionStub, 0, bisectionCount)
//		for i := 0; i < bisectionCount; i++ {
//			assertion := &valprotocol.ExecutionAssertionStub{
//				AfterHash:        bisectChal.MachineHashes[i+1],
//				DidInboxInsn:     bisectChal.DidInboxInsns[i],
//				NumGas:           bisectChal.Gases[i],
//				FirstMessageHash: bisectChal.MessageAccs[i],
//				LastMessageHash:  bisectChal.MessageAccs[i+1],
//				FirstLogHash:     bisectChal.LogAccs[i],
//				LastLogHash:      bisectChal.LogAccs[i+1],
//			}
//			assertions = append(assertions, assertion)
//		}
//		return arbbridge.ExecutionBisectionEvent{
//			ChainInfo:  chainInfo,
//			Assertions: assertions,
//			TotalSteps: bisectChal.TotalSteps,
//			Deadline:   common.TimeTicks{Val: bisectChal.DeadlineTicks},
//		}, nil
//	} else if log.Topics[0] == oneStepProofCompletedID {
//		_, err := c.challenge.ParseOneStepProofCompleted(log)
//		if err != nil {
//			return nil, err
//		}
//		return arbbridge.OneStepProofEvent{
//			ChainInfo: chainInfo,
//		}, nil
//	}
//	return c.bisectionChallengeWatcher.parseBisectionEvent(chainInfo, log)
//}
