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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
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
	client   *GoArbClient
	address  common.Address
}

func newPendingTopChallengeWatcher(address common.Address, client *GoArbClient) (*pendingTopChallengeWatcher, error) {
	bisectionChallenge, err := newBisectionChallengeWatcher(address, client)
	if err != nil {
		return nil, err
	}
	//pendingTopContract, err := pendingtopchallenge.NewPendingTopChallenge(address, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to pendingTopChallenge")
	//}
	return &pendingTopChallengeWatcher{
		bisectionChallengeWatcher: bisectionChallenge,
		contract:                  nil,
		client:                    client,
		address:                   address,
	}, nil
}

func (c *pendingTopChallengeWatcher) GetEvents(ctx context.Context, blockId *structures.BlockId) ([]arbbridge.Event, error) {
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
	//	event, err := c.parsePendingTopEvent(getLogChainInfo(evmLog), evmLog)
	//	if err != nil {
	//		return nil, err
	//	}
	//	events = append(events, event)
	//}
	//return events, nil
	return nil, nil
}

func (c *pendingTopChallengeWatcher) topics() []ethcommon.Hash {
	tops := []ethcommon.Hash{
		pendingTopBisectedID,
		pendingTopOneStepProofCompletedID,
	}
	return append(tops, c.bisectionChallengeWatcher.topics()...)
}

func (c *pendingTopChallengeWatcher) StartConnection(ctx context.Context, startHeight *common.TimeBlocks, startLogIndex uint) (<-chan arbbridge.MaybeEvent, error) {
	headers := make(chan arbbridge.MaybeEvent)
	c.client.GoEthClient.registerOutChan(headers)
	//headers := make(chan *types.Header)
	//headersSub, err := c.client.SubscribeNewHead(ctx, headers)
	//if err != nil {
	//	return err
	//}

	//filter := ethereum.FilterQuery{
	//	Addresses: []ethcommon.Address{c.address},
	//	Topics:    [][]ethcommon.Hash{c.topics()},
	//}

	//logChan := make(chan types.Log, 1024)
	//logErrChan := make(chan error, 10)

	//if err := getLogs(ctx, c.client, filter, big.NewInt(0), logChan, logErrChan); err != nil {
	//	return err
	//}

	go func() {
		//defer headersSub.Unsubscribe()

		for {
			select {
			case <-ctx.Done():
				break
				//case evmLog, ok := <-logChan:
				//	if !ok {
				//		errChan <- errors.New("logChan terminated early")
				//		return
				//	}
				//	if err := c.processEvents(ctx, evmLog, outChan); err != nil {
				//		errChan <- err
				//		return
				//	}
				//case err := <-logErrChan:
				//	errChan <- err
				//	return
				//	//case err := <-headersSub.Err():
				//	//	errChan <- err
				//	return
			}
		}
	}()
	return headers, nil
}

//func (c *pendingTopChallengeWatcher) processEvents(ctx context.Context, log types.Log, outChan chan arbbridge.Notification) error {
//	event, err := func() (arbbridge.Event, error) {
//		if log.Topics[0] == pendingTopBisectedID {
//			eventVal, err := c.contract.ParseBisected(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.PendingTopBisectionEvent{
//				//ChainHashes: hashSliceToHashes(eventVal.ChainHashes),
//				ChainHashes: nil,
//				TotalLength: eventVal.TotalLength,
//				Deadline:    common.TimeTicks{Val: eventVal.DeadlineTicks},
//			}, nil
//		} else if log.Topics[0] == pendingTopOneStepProofCompletedID {
//			_, err := c.contract.ParseOneStepProofCompleted(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.OneStepProofEvent{}, nil
//		} else {
//			event, err := c.bisectionChallengeWatcher.parseBisectionEvent(log)
//			if event != nil || err != nil {
//				return event, err
//			}
//		}
//		return nil, errors2.New("unknown arbitrum event type")
//	}()
//
//	if err != nil {
//		return err
//	}
//
//	//header, err := c.client.HeaderByHash(ctx, log.BlockHash)
//	//if err != nil {
//	//	return err
//	//}
//	outChan <- arbbridge.Notification{
//		//BlockHeader: common.Hash{},
//		//BlockHeight: nil,
//		//VMID:        common.NewAddressFromEth(c.address),
//		Event:  event,
//		TxHash: log.TxHash,
//	}
//	return nil
//}
