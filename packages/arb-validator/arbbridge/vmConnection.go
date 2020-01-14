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
	"log"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func HandleBlockchainNotifications(ctx context.Context, startHeight *common.TimeBlocks, startLogIndex uint, contract ContractWatcher) <-chan Event {
	rawEventChan := make(chan Event, 1024)
	errChan := make(chan error, 1024)
	if err := contract.StartConnection(ctx, startHeight, startLogIndex, rawEventChan, errChan); err != nil {
		close(rawEventChan)
		close(errChan)
		return nil
	}

	eventChan := make(chan Event, 1024)
	go func() {
		defer close(rawEventChan)
		defer close(errChan)
		defer close(eventChan)
		for {
			hitError := false
			select {
			case <-ctx.Done():
				break
			case notification, ok := <-rawEventChan:
				if !ok {
					hitError = true
					break
				}
				eventChan <- notification
			case <-errChan:
				hitError = true
			}

			if hitError {
				// Ignore error and try to reset connection
				for {
					err := contract.StartConnection(ctx, startHeight, 0, rawEventChan, errChan)
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
	return eventChan
}
