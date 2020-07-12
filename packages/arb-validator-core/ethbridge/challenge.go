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
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"strings"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge/executionchallenge"
)

var initiatedChallengeID ethcommon.Hash
var timedOutAsserterID ethcommon.Hash
var timedOutChallengerID ethcommon.Hash

func init() {
	parsed, err := abi.JSON(strings.NewReader(executionchallenge.ExecutionChallengeABI))
	if err != nil {
		panic(err)
	}
	initiatedChallengeID = parsed.Events["InitiatedChallenge"].ID
	timedOutAsserterID = parsed.Events["AsserterTimedOut"].ID
	timedOutChallengerID = parsed.Events["ChallengerTimedOut"].ID
}

type challenge struct {
	Challenge *executionchallenge.Challenge

	client          ethutils.EthClient
	auth            *TransactAuth
	contractAddress ethcommon.Address
}

func newChallenge(address ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*challenge, error) {
	challengeContract, err := executionchallenge.NewChallenge(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to ChallengeManager")
	}

	return &challenge{Challenge: challengeContract, client: client, auth: auth, contractAddress: address}, nil
}

func (c *challenge) TimeoutChallenge(ctx context.Context) error {
	c.auth.Lock()
	defer c.auth.Unlock()
	tx, err := c.Challenge.TimeoutChallenge(c.auth.getAuth(ctx))
	if err != nil {
		return c.Challenge.TimeoutChallengeCall(
			ctx,
			c.client,
			c.auth.auth.From,
			c.contractAddress,
		)
	}
	return c.waitForReceipt(ctx, tx, "TimeoutChallenge")
}

func (c *challenge) waitForReceipt(ctx context.Context, tx *types.Transaction, methodName string) error {
	return waitForReceipt(ctx, c.client, c.auth.auth.From, tx, methodName)
}

type challengeWatcher struct {
	Challenge *executionchallenge.Challenge
}

func newChallengeWatcher(address ethcommon.Address, client ethutils.EthClient) (*challengeWatcher, error) {
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

func (c *challengeWatcher) parseChallengeEvent(chainInfo arbbridge.ChainInfo, log types.Log) (arbbridge.Event, error) {
	if log.Topics[0] == initiatedChallengeID {
		eventVal, err := c.Challenge.ParseInitiatedChallenge(log)
		if err != nil {
			return nil, err
		}
		return arbbridge.InitiateChallengeEvent{
			ChainInfo: chainInfo,
			Deadline:  common.TimeTicks{Val: eventVal.DeadlineTicks},
		}, nil
	} else if log.Topics[0] == timedOutAsserterID {
		_, err := c.Challenge.ParseAsserterTimedOut(log)
		if err != nil {
			return nil, err
		}
		return arbbridge.AsserterTimeoutEvent{
			ChainInfo: chainInfo,
		}, nil
	} else if log.Topics[0] == timedOutChallengerID {
		_, err := c.Challenge.ParseChallengerTimedOut(log)
		if err != nil {
			return nil, err
		}
		return arbbridge.ChallengerTimeoutEvent{
			ChainInfo: chainInfo,
		}, nil
	}
	return nil, errors.New("unknown arbitrum event type")
}
