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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

type messagesChallengeWatcher struct {
	*bisectionChallengeWatcher
	challengeInfo *challengeData
	client        *goEthdata
}

func newMessagesChallengeWatcher(address common.Address, client *goEthdata) (*messagesChallengeWatcher, error) {
	bisectionChallenge, err := newBisectionChallengeWatcher(address, client)
	if err != nil {
		return nil, err
	}

	chalData := client.challenges[address]
	client.challengeWatchersMutex.Lock()
	if _, ok := client.challengeWatcherEvents[chalData]; !ok {
		client.challengeWatcherEvents[chalData] = make(map[*common.BlockId][]arbbridge.Event)
	}
	client.challengeWatchersMutex.Unlock()

	return &messagesChallengeWatcher{bisectionChallengeWatcher: bisectionChallenge, challengeInfo: chalData, client: client}, nil
}

func (c *messagesChallengeWatcher) GetEvents(ctx context.Context, blockID *common.BlockId) ([]arbbridge.Event, error) {
	c.client.challengeWatchersMutex.Lock()
	cw := c.client.challengeWatcherEvents[c.challengeInfo][blockID]
	c.client.challengeWatchersMutex.Unlock()
	return cw, nil
}
