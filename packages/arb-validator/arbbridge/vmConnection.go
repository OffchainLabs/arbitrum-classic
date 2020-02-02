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

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func HandleBlockchainEvents(
	ctx context.Context,
	client ArbAuthClient,
	startBlockID *structures.BlockID,
	startLogIndex uint,
	contract ContractWatcher,
) (context.Context, <-chan Event) {
	eventChan := make(chan Event, 1024)
	reorgCtx, cancelFunc := context.WithCancel(ctx)
	go func() {
		defer cancelFunc()
		defer close(eventChan)
		headersChan, err := client.SubscribeBlockHeaders(ctx, startBlockID)
		if err != nil {
			log.Println("error in challenge", err)
			return
		}
		for maybeBlockID := range headersChan {
			if maybeBlockID.Err != nil {
				log.Println("error in challenge", err)
				return
			}

			blockID := maybeBlockID.BlockID

			events, err := contract.GetEvents(ctx, blockID)
			if err != nil {
				log.Println("error in challenge", err)
				return
			}

			if blockID.Height.Cmp(startBlockID.Height) == 0 {
				for _, event := range events {
					if event.GetChainInfo().LogIndex >= startLogIndex {
						eventChan <- event
					}
				}
			} else {
				for _, event := range events {
					eventChan <- event
				}
			}
		}
	}()
	return reorgCtx, eventChan
}
