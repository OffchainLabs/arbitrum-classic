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

	"github.com/pkg/errors"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

var inboxTopBisectedID ethcommon.Hash
var inboxTopOneStepProofCompletedID ethcommon.Hash

func init() {
	parsed, err := abi.JSON(strings.NewReader(ethbridgecontracts.InboxTopChallengeABI))
	if err != nil {
		panic(err)
	}
	inboxTopBisectedID = parsed.Events["Bisected"].ID
	inboxTopOneStepProofCompletedID = parsed.Events["OneStepProofCompleted"].ID
}

type inboxTopChallengeWatcher struct {
	*bisectionChallengeWatcher
	contract *ethbridgecontracts.InboxTopChallenge
	client   ethutils.EthClient
	address  ethcommon.Address
	topics   [][]ethcommon.Hash
}

func newInboxTopChallengeWatcher(address ethcommon.Address, client ethutils.EthClient) (*inboxTopChallengeWatcher, error) {
	bisectionChallenge, err := newBisectionChallengeWatcher(address, client)
	if err != nil {
		return nil, err
	}
	inboxTopContract, err := ethbridgecontracts.NewInboxTopChallenge(address, client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to InboxTopChallenge")
	}
	tops := []ethcommon.Hash{
		inboxTopBisectedID,
		inboxTopOneStepProofCompletedID,
	}
	tops = append(tops, bisectionChallenge.topics()...)

	return &inboxTopChallengeWatcher{
		bisectionChallengeWatcher: bisectionChallenge,
		contract:                  inboxTopContract,
		client:                    client,
		address:                   address,
		topics:                    [][]ethcommon.Hash{tops},
	}, nil
}

func (c *inboxTopChallengeWatcher) GetEvents(ctx context.Context, blockId *common.BlockId, timestamp *big.Int) ([]arbbridge.Event, error) {
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
		event, err := c.parseInboxTopEvent(getLogChainInfo(evmLog), evmLog)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (c *inboxTopChallengeWatcher) parseInboxTopEvent(chainInfo arbbridge.ChainInfo, log types.Log) (arbbridge.Event, error) {
	if log.Topics[0] == inboxTopBisectedID {
		eventVal, err := c.contract.ParseBisected(log)
		if err != nil {
			return nil, err
		}
		return arbbridge.InboxTopBisectionEvent{
			ChainInfo:   chainInfo,
			ChainHashes: hashSliceToHashes(eventVal.ChainHashes),
			TotalLength: eventVal.TotalLength,
			Deadline:    common.TimeTicks{Val: eventVal.DeadlineTicks},
		}, nil
	} else if log.Topics[0] == inboxTopOneStepProofCompletedID {
		_, err := c.contract.ParseOneStepProofCompleted(log)
		if err != nil {
			return nil, err
		}
		return arbbridge.OneStepProofEvent{
			ChainInfo: chainInfo,
		}, nil
	}
	return c.bisectionChallengeWatcher.parseBisectionEvent(chainInfo, log)
}
