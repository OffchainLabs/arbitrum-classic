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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func HandleBlockchainEvents(
	ctx context.Context,
	client ArbAuthClient,
	startBlockId *common.BlockId,
	startLogIndex uint,
	contract ContractWatcher,
) (context.Context, <-chan Event) {
	eventChan := make(chan Event, 1024)
	reorgCtx, cancelFunc := context.WithCancel(ctx)
	go func() {
		defer cancelFunc()
		defer close(eventChan)
		headersChan, err := client.SubscribeBlockHeaders(ctx, startBlockId)
		if err != nil {
			logger.Error().Stack().Err(err).Msg("Error subscribing to headers")
			return
		}
		for maybeBlockId := range headersChan {
			if maybeBlockId.Err != nil {
				logger.Error().Stack().Err(maybeBlockId.Err).Msg("Error getting header")
				return
			}

			blockId := maybeBlockId.BlockId

			events, err := contract.GetEvents(ctx, blockId, maybeBlockId.Timestamp)
			if err != nil {
				logger.Error().Stack().Err(err).Msg("Error getting events")
				return
			}

			if blockId.Height.Cmp(startBlockId.Height) == 0 {
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
