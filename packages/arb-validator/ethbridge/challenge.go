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

package ethbridge

import (
	"context"
	"strings"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/executionchallenge"
)

var initiatedChallengeID ethcommon.Hash
var timedOutAsserterID ethcommon.Hash
var timedOutChallengerID ethcommon.Hash

func init() {
	parsed, err := abi.JSON(strings.NewReader(executionchallenge.ExecutionChallengeABI))
	if err != nil {
		panic(err)
	}
	initiatedChallengeID = parsed.Events["InitiatedChallenge"].ID()
	timedOutAsserterID = parsed.Events["AsserterTimedOut"].ID()
	timedOutChallengerID = parsed.Events["ChallengerTimedOut"].ID()
}

type challenge struct {
	Challenge *executionchallenge.Challenge

	client *ethclient.Client
	auth   *bind.TransactOpts
}

func newChallenge(address ethcommon.Address, client *ethclient.Client, auth *bind.TransactOpts) (*challenge, error) {
	challengeContract, err := executionchallenge.NewChallenge(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to ChallengeManager")
	}

	return &challenge{Challenge: challengeContract, client: client, auth: auth}, nil
}

func (c *challenge) TimeoutChallenge(
	ctx context.Context,
) error {
	c.auth.Context = ctx
	tx, err := c.Challenge.TimeoutChallenge(c.auth)
	if err != nil {
		return err
	}
	return c.waitForReceipt(ctx, tx, "TimeoutChallenge")
}

func (c *challenge) waitForReceipt(ctx context.Context, tx *types.Transaction, methodName string) error {
	return waitForReceipt(ctx, c.client, c.auth.From, tx, methodName)
}

type challengeWatcher struct {
	Challenge *executionchallenge.Challenge
}

func newChallengeWatcher(address ethcommon.Address, client *ethclient.Client) (*challengeWatcher, error) {
	challengeContract, err := executionchallenge.NewChallenge(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to ChallengeManager")
	}

	return &challengeWatcher{Challenge: challengeContract}, nil
}

func (c *challengeWatcher) topics() []ethcommon.Hash {
	return []ethcommon.Hash{
		initiatedChallengeID,
		timedOutAsserterID,
		timedOutChallengerID,
	}
}

func (c *challengeWatcher) parseChallengeEvent(log types.Log) (arbbridge.Event, error) {
	if log.Topics[0] == initiatedChallengeID {
		eventVal, err := c.Challenge.ParseInitiatedChallenge(log)
		if err != nil {
			return nil, err
		}
		return arbbridge.InitiateChallengeEvent{
			Deadline: common.TimeTicks{Val: eventVal.DeadlineTicks},
		}, nil
	} else if log.Topics[0] == timedOutAsserterID {
		_, err := c.Challenge.ParseAsserterTimedOut(log)
		if err != nil {
			return nil, err
		}
		return arbbridge.AsserterTimeoutEvent{}, nil
	} else if log.Topics[0] == timedOutChallengerID {
		_, err := c.Challenge.ParseChallengerTimedOut(log)
		if err != nil {
			return nil, err
		}
		return arbbridge.ChallengerTimeoutEvent{}, nil
	}
	return nil, nil
}
