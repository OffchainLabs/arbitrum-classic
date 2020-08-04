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
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

var bisectedAssertionID ethcommon.Hash
var oneStepProofCompletedID ethcommon.Hash

func init() {
	parsed, err := abi.JSON(strings.NewReader(ethbridgecontracts.ExecutionChallengeABI))
	if err != nil {
		panic(err)
	}
	bisectedAssertionID = parsed.Events["BisectedAssertion"].ID
	oneStepProofCompletedID = parsed.Events["OneStepProofCompleted"].ID
}

type executionChallengeWatcher struct {
	*bisectionChallengeWatcher
	challenge *ethbridgecontracts.ExecutionChallenge
	client    ethutils.EthClient
	address   ethcommon.Address
	topics    [][]ethcommon.Hash
}

func newExecutionChallengeWatcher(address ethcommon.Address, client ethutils.EthClient) (*executionChallengeWatcher, error) {
	bisectionChallenge, err := newBisectionChallengeWatcher(address, client)
	if err != nil {
		return nil, err
	}
	executionContract, err := ethbridgecontracts.NewExecutionChallenge(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to ChallengeManager")
	}
	tops := []ethcommon.Hash{
		bisectedAssertionID,
		oneStepProofCompletedID,
	}
	tops = append(tops, bisectionChallenge.topics()...)
	return &executionChallengeWatcher{
		bisectionChallengeWatcher: bisectionChallenge,
		challenge:                 executionContract,
		client:                    client,
		address:                   address,
		topics:                    [][]ethcommon.Hash{tops},
	}, nil
}

func (c *executionChallengeWatcher) GetEvents(ctx context.Context, blockId *common.BlockId, timestamp *big.Int) ([]arbbridge.Event, error) {
	bh := blockId.HeaderHash.ToEthHash()
	logs, err := c.client.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &bh,
		Addresses: []ethcommon.Address{c.address},
		Topics:    c.topics,
	})
	if err != nil {
		return nil, err
	}
	events := make([]arbbridge.Event, 0, len(logs))
	for _, evmLog := range logs {
		event, err := c.parseExecutionEvent(getLogChainInfo(evmLog), evmLog)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (c *executionChallengeWatcher) parseExecutionEvent(chainInfo arbbridge.ChainInfo, log types.Log) (arbbridge.Event, error) {
	if log.Topics[0] == bisectedAssertionID {
		bisectChal, err := c.challenge.ParseBisectedAssertion(log)
		if err != nil {
			return nil, err
		}
		bisectionCount := len(bisectChal.MachineHashes) - 1
		assertions := make([]*valprotocol.ExecutionAssertionStub, 0, bisectionCount)
		for i := 0; i < bisectionCount; i++ {
			assertion := &valprotocol.ExecutionAssertionStub{
				NumGas:            bisectChal.Gases[i],
				BeforeMachineHash: bisectChal.MachineHashes[i],
				AfterMachineHash:  bisectChal.MachineHashes[i+1],
				BeforeInboxHash:   bisectChal.InboxAccs[i],
				AfterInboxHash:    bisectChal.InboxAccs[i+1],
				FirstMessageHash:  bisectChal.MessageAccs[i],
				LastMessageHash:   bisectChal.MessageAccs[i+1],
				MessageCount:      bisectChal.OutCounts[i],
				FirstLogHash:      bisectChal.LogAccs[i],
				LastLogHash:       bisectChal.LogAccs[i+1],
				LogCount:          bisectChal.OutCounts[bisectionCount+i],
			}
			assertions = append(assertions, assertion)
		}
		return arbbridge.ExecutionBisectionEvent{
			ChainInfo:  chainInfo,
			Assertions: assertions,
			TotalSteps: bisectChal.TotalSteps,
			Deadline:   common.TimeTicks{Val: bisectChal.DeadlineTicks},
		}, nil
	} else if log.Topics[0] == oneStepProofCompletedID {
		_, err := c.challenge.ParseOneStepProofCompleted(log)
		if err != nil {
			return nil, err
		}
		return arbbridge.OneStepProofEvent{
			ChainInfo: chainInfo,
		}, nil
	}
	return c.bisectionChallengeWatcher.parseBisectionEvent(chainInfo, log)
}
