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
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var initiatedChallengeID ethcommon.Hash
var timedOutAsserterID ethcommon.Hash
var timedOutChallengerID ethcommon.Hash

type challenge struct {
	client *GoArbAuthClient
	//auth   *bind.TransactOpts
	contractAddress common.Address
	state           int
	deadlineTicks   common.TimeTicks
}

func newChallenge(address common.Address, client *GoArbAuthClient) (*challenge, error) {
	//challengeContract, err := executionchallenge.NewChallenge(address, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to ChallengeManager")
	//}
	//ch := client.GoEthClient.challenges[address]
	//addr := client.GoEthClient.getNextAddress()
	chal := challenge{client: client, contractAddress: address}
	//client.GoEthClient.challenges[addr] =

	return &chal, nil
}

func (c *challenge) TimeoutChallenge(
	ctx context.Context,
) error {
	//c.auth.Context = ctx
	//tx, err := c.Challenge.TimeoutChallenge(c.auth)
	//if err != nil {
	//	return err
	//}
	//return c.waitForReceipt(ctx, tx, "TimeoutChallenge")
	return nil
}

type challengeWatcher struct {
	client    *GoArbClient
	Challenge common.Address
}

func newChallengeWatcher(address common.Address, client *GoArbClient) (*challengeWatcher, error) {
	//challengeContract, err := executionchallenge.NewChallenge(address, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to ChallengeManager")
	//}

	return &challengeWatcher{client: client, Challenge: address}, nil
}

func (c *challengeWatcher) topics() []ethcommon.Hash {
	return []ethcommon.Hash{
		initiatedChallengeID,
		timedOutAsserterID,
		timedOutChallengerID,
	}
}

//func (c *challengeWatcher) parseChallengeEvent(log types.Log) (arbbridge.Event, error) {
//	if log.Topics[0] == initiatedChallengeID {
//		eventVal, err := c.Challenge.ParseInitiatedChallenge(log)
//		if err != nil {
//			return nil, err
//		}
//		return arbbridge.InitiateChallengeEvent{
//			Deadline: common.TimeTicks{Val: eventVal.DeadlineTicks},
//		}, nil
//	} else if log.Topics[0] == timedOutAsserterID {
//		_, err := c.Challenge.ParseAsserterTimedOut(log)
//		if err != nil {
//			return nil, err
//		}
//		return arbbridge.AsserterTimeoutEvent{}, nil
//	} else if log.Topics[0] == timedOutChallengerID {
//		_, err := c.Challenge.ParseChallengerTimedOut(log)
//		if err != nil {
//			return nil, err
//		}
//		return arbbridge.ChallengerTimeoutEvent{}, nil
//	}
//	return nil, nil
//}
