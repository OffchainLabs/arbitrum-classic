/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

package arbbridge

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func HandleBlockchainNotifications(ctx context.Context, startBlockId *structures.BlockId, startLogIndex uint, contract ContractWatcher) (context.Context, <-chan Event, error) {
	rawEventChan, err := contract.StartConnection(ctx, startBlockId.Height, startLogIndex)
	if err != nil {
		return nil, nil, err
	}

	reorgCtx, cancelFunc := context.WithCancel(ctx)

	eventChan := make(chan Event, 1024)
	go func() {
		defer cancelFunc()
		defer close(eventChan)

		latestBlockId := startBlockId
		latestLogIndex := startLogIndex

		for {
			var err error
			select {
			case <-ctx.Done():
				log.Println("Event monitor canceled")
				return
			case maybeEvent, ok := <-rawEventChan:
				if !ok {
					log.Println("rawEventChan channel closed")
					err = errors.New("rawEventChan closed")
					break
				}
				if maybeEvent.Err != nil {
					log.Println("HandleBlockchainNotifications rawEventChan had error", maybeEvent.Err)
					err = maybeEvent.Err
					break
				}

				chainInfo := maybeEvent.Event.GetChainInfo()
				switch chainInfo.BlockId.Height.Cmp(latestBlockId.Height) {
				case -1:
					// reorg
					log.Println("Reorg occured")
					return
				case 0:
					if !chainInfo.BlockId.HeaderHash.Equals(latestBlockId.HeaderHash) {
						// reorg
						log.Println("Reorg occured")
						return
					}
					if chainInfo.LogIndex >= latestLogIndex {
						latestLogIndex = chainInfo.LogIndex
						eventChan <- maybeEvent.Event
					}
				case 1:
					latestBlockId = chainInfo.BlockId
					latestLogIndex = chainInfo.LogIndex
					eventChan <- maybeEvent.Event
				}
			}

			if err != nil {
				// Ignore error and try to reset connection
				log.Println("Restarting connection due to error", err)
				for {
					rawEventChan, err = contract.StartConnection(ctx, latestBlockId.Height, latestLogIndex+1)
					if err == nil {
						break
					}
					select {
					case <-ctx.Done():
						return
					default:
					}
					log.Println("Error: Can't connect to blockchain", err)
					time.Sleep(5 * time.Second)
				}
			}
		}
	}()
	return reorgCtx, eventChan, nil
}
