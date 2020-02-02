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

package rollupmanager

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type ArbClientStressTest struct {
	arbbridge.ArbClient
	reorgInterval time.Duration
}

func NewStressTestClient(client arbbridge.ArbClient, reorgInterval time.Duration) *ArbClientStressTest {
	return &ArbClientStressTest{client, reorgInterval}
}

var errReorg = errors.New("reorg occurred")

func (st *ArbClientStressTest) SubscribeBlockHeaders(ctx context.Context, startBlockID *structures.BlockID) (<-chan arbbridge.MaybeBlockID, error) {
	rawHeadersChan, err := st.ArbClient.SubscribeBlockHeaders(ctx, startBlockID)
	if err != nil {
		return nil, err
	}
	ticker := time.NewTicker(st.reorgInterval)
	headerChan := make(chan arbbridge.MaybeBlockID, 10)
	go func() {
		defer close(headerChan)
		for {
			select {
			case maybeHeader, ok := <-rawHeadersChan:
				if !ok {
					return
				}
				headerChan <- maybeHeader
				if maybeHeader.Err != nil {
					return
				}

			case <-ticker.C:
				log.Println("Manually triggering reorg")
				headerChan <- arbbridge.MaybeBlockID{Err: errReorg}
				return
			}
		}
	}()
	return headerChan, nil
}
