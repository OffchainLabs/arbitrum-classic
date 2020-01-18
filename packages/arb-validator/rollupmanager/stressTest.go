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
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type ReorgStressTestWatcher struct {
	watcher       arbbridge.ArbRollupWatcher
	reorgInterval time.Duration
}

func NewStressTestWatcher(watcher arbbridge.ArbRollupWatcher, reorgInterval time.Duration) arbbridge.ArbRollupWatcher {
	return &ReorgStressTestWatcher{watcher, reorgInterval}
}

func (w *ReorgStressTestWatcher) GetEvents(ctx context.Context, blockId *structures.BlockId) ([]arbbridge.Event, error) {
	return nil, nil
}

func (con *ReorgStressTestWatcher) GetCreationHeight(ctx context.Context) (*structures.BlockId, error) {
	return con.watcher.GetCreationHeight(ctx)
}

func (w *ReorgStressTestWatcher) StartConnection(
	ctx context.Context,
	startHeight *common.TimeBlocks,
	startLogIndex uint,
) (<-chan arbbridge.MaybeEvent, error) {
	//ch, err := w.watcher.StartConnection(ctx, startHeight, startLogIndex)
	//if err != nil {
	//	return ch, err
	//}
	newCh := make(chan arbbridge.MaybeEvent)
	//go func() {
	//	defer close(newCh)
	//	ticker := time.NewTicker(w.reorgInterval)
	//	defer ticker.Stop()
	//	fakeBlockId := &structures.BlockId{}
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			return
	//		case <-ticker.C:
	//			log.Println("Stress tester triggering fake reorg")
	//			newCh <- arbbridge.MaybeEvent{
	//				&arbbridge.NewTimeEvent{
	//					arbbridge.ChainInfo{
	//						fakeBlockId,
	//						0,
	//						common.Hash{},
	//					},
	//				},
	//				nil,
	//			}
	//		case maybeEvent, ok := <-ch:
	//			if !ok {
	//				return
	//			}
	//			bid := maybeEvent.Event.GetChainInfo().BlockId
	//			fakeBlockId = &structures.BlockId{bid.Height, common.Hash{}}
	//			newCh <- maybeEvent
	//		}
	//	}
	//}()
	return newCh, nil
}

func (w *ReorgStressTestWatcher) GetParams(ctx context.Context) (structures.ChainParams, error) {
	return w.watcher.GetParams(ctx)
}

func (w *ReorgStressTestWatcher) InboxAddress(ctx context.Context) (common.Address, error) {
	return w.watcher.InboxAddress(ctx)
}
