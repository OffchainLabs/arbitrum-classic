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
	"strings"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/pendingtopchallenge"
)

var pendingTopBisectedID ethcommon.Hash
var pendingTopOneStepProofCompletedID ethcommon.Hash

func init() {
	parsed, err := abi.JSON(strings.NewReader(pendingtopchallenge.PendingTopChallengeABI))
	if err != nil {
		panic(err)
	}
	pendingTopBisectedID = parsed.Events["Bisected"].ID()
	pendingTopOneStepProofCompletedID = parsed.Events["OneStepProofCompleted"].ID()
}

type pendingTopChallengeWatcher struct {
	*bisectionChallengeWatcher
	contract *pendingtopchallenge.PendingTopChallenge
	client   *ethclient.Client
	address  ethcommon.Address
}

func newPendingTopChallengeWatcher(address ethcommon.Address, client *ethclient.Client) (*pendingTopChallengeWatcher, error) {
	bisectionChallenge, err := newBisectionChallengeWatcher(address, client)
	if err != nil {
		return nil, err
	}
	pendingTopContract, err := pendingtopchallenge.NewPendingTopChallenge(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to PendingTopChallenge")
	}
	return &pendingTopChallengeWatcher{
		bisectionChallengeWatcher: bisectionChallenge,
		contract:                  pendingTopContract,
		client:                    client,
		address:                   address,
	}, nil
}

func (c *pendingTopChallengeWatcher) topics() []ethcommon.Hash {
	tops := []ethcommon.Hash{
		pendingTopBisectedID,
		pendingTopOneStepProofCompletedID,
	}
	return append(tops, c.bisectionChallengeWatcher.topics()...)
}

func (c *pendingTopChallengeWatcher) StartConnection(ctx context.Context, startHeight *common.TimeBlocks, startLogIndex uint, outChan chan arbbridge.Event, errChan chan error) error {
	headers := make(chan *types.Header)
	headersSub, err := c.client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return err
	}

	filter := ethereum.FilterQuery{
		Addresses: []ethcommon.Address{c.address},
		Topics:    [][]ethcommon.Hash{c.topics()},
	}

	logChan := make(chan types.Log, 1024)
	logErrChan := make(chan error, 10)

	if err := getLogs(ctx, c.client, filter, startHeight, startLogIndex, logChan, logErrChan); err != nil {
		return err
	}

	go func() {
		defer headersSub.Unsubscribe()

		for {
			select {
			case <-ctx.Done():
				break
			case evmLog, ok := <-logChan:
				if !ok {
					errChan <- errors.New("logChan terminated early")
					return
				}
				header, err := c.client.HeaderByHash(ctx, evmLog.BlockHash)
				if err != nil {
					errChan <- err
					return
				}
				chainInfo := getChainInfo(evmLog, header)
				event, err := c.parsePendingTopEvent(chainInfo, evmLog)
				if err != nil {
					errChan <- err
					return
				}
				outChan <- event
			case err := <-logErrChan:
				errChan <- err
				return
			case err := <-headersSub.Err():
				errChan <- err
				return
			}
		}
	}()
	return nil
}

func (c *pendingTopChallengeWatcher) parsePendingTopEvent(chainInfo arbbridge.ChainInfo, log types.Log) (arbbridge.Event, error) {
	if log.Topics[0] == pendingTopBisectedID {
		eventVal, err := c.contract.ParseBisected(log)
		if err != nil {
			return nil, err
		}
		return arbbridge.PendingTopBisectionEvent{
			ChainInfo:   chainInfo,
			ChainHashes: hashSliceToHashes(eventVal.ChainHashes),
			TotalLength: eventVal.TotalLength,
			Deadline:    common.TimeTicks{Val: eventVal.DeadlineTicks},
		}, nil
	} else if log.Topics[0] == pendingTopOneStepProofCompletedID {
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
