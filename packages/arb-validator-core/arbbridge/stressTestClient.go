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

package arbbridge

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type ArbClientStressTest struct {
	ArbClient
	reorgInterval time.Duration
}

func NewStressTestClient(client ArbClient, reorgInterval time.Duration) *ArbClientStressTest {
	return &ArbClientStressTest{client, reorgInterval}
}

var reorgError = errors.New("reorg occured")

func (st *ArbClientStressTest) SubscribeBlockHeadersAfter(ctx context.Context, prevBlockId *common.BlockId) (<-chan MaybeBlockId, error) {
	headers, err := st.ArbClient.SubscribeBlockHeadersAfter(ctx, prevBlockId)
	if err != nil {
		return nil, err
	}
	return st.addReorgs(headers), nil
}

func (st *ArbClientStressTest) SubscribeBlockHeaders(ctx context.Context, startBlockId *common.BlockId) (<-chan MaybeBlockId, error) {
	headers, err := st.ArbClient.SubscribeBlockHeaders(ctx, startBlockId)
	if err != nil {
		return nil, err
	}
	return st.addReorgs(headers), nil
}

func (st *ArbClientStressTest) addReorgs(blockIdChan <-chan MaybeBlockId) <-chan MaybeBlockId {
	ticker := time.NewTicker(st.reorgInterval)
	headerChan := make(chan MaybeBlockId, 10)
	go func() {
		defer close(headerChan)
		for {
			select {
			case maybeHeader, ok := <-blockIdChan:
				if !ok {
					return
				}
				headerChan <- maybeHeader
				if maybeHeader.Err != nil {
					return
				}

			case <-ticker.C:
				log.Println("Manually triggering reorg")
				headerChan <- MaybeBlockId{Err: reorgError}
				return
			}
		}
	}()
	return headerChan
}
