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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/messageschallenge"
)

var messagesBisectedID ethcommon.Hash
var messagesOneStepProofCompletedID ethcommon.Hash

func init() {
	parsed, err := abi.JSON(strings.NewReader(messageschallenge.MessagesChallengeABI))
	if err != nil {
		panic(err)
	}
	messagesBisectedID = parsed.Events["Bisected"].ID()
	messagesOneStepProofCompletedID = parsed.Events["OneStepProofCompleted"].ID()
}

type messagesChallengeWatcher struct {
	*bisectionChallengeWatcher
	challengeInfo *challengeData
	client        *MockArbClient
	//address  ethcommon.Address
}

func newMessagesChallengeWatcher(address common.Address, client *MockArbClient) (*messagesChallengeWatcher, error) {
	bisectionChallenge, err := newBisectionChallengeWatcher(address, client)
	if err != nil {
		return nil, err
	}
	//messagesContract, err := messageschallenge.NewMessagesChallenge(address, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to messagesChallenge")
	//}
	chalData := client.MockEthClient.challenges[address]
	client.MockEthClient.challengeWatchers[chalData] = make(map[*structures.BlockId][]arbbridge.Event)

	return &messagesChallengeWatcher{bisectionChallengeWatcher: bisectionChallenge, challengeInfo: chalData, client: client}, nil
}

func (c *messagesChallengeWatcher) GetEvents(ctx context.Context, blockId *structures.BlockId) ([]arbbridge.Event, error) {
	fmt.Println("in messagesChallengeWatcher GetEvents")
	//bh := blockId.HeaderHash.ToEthHash()
	//logs, err := c.client.FilterLogs(ctx, ethereum.FilterQuery{
	//	BlockHash: &bh,
	//	Addresses: []ethcommon.Address{c.address},
	//	Topics:    c.topics,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//events := make([]arbbridge.Event, 0, len(logs))
	//for _, evmLog := range logs {
	//	event, err := c.parseMessagesEvent(getLogChainInfo(evmLog), evmLog)
	//	if err != nil {
	//		return nil, err
	//	}
	//	events = append(events, event)
	//}
	//return events, nil
	return c.client.MockEthClient.challengeWatchers[c.challengeInfo][blockId], nil
}

func (c *messagesChallengeWatcher) topics() []ethcommon.Hash {
	tops := []ethcommon.Hash{
		messagesBisectedID,
		messagesOneStepProofCompletedID,
	}
	return append(tops, c.bisectionChallengeWatcher.topics()...)
}
