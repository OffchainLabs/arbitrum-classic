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

package checkpointing

import (
	"context"
	"sync"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type asyncCheckpointWriter struct {
	sync.Mutex
	checkpointer *RollupCheckpointerImpl
	notifyChan   chan interface{}
	nextJob      func()
	doneChans    []chan struct{}
}

func NewAsyncCheckpointWriter(ctx context.Context, cp *RollupCheckpointerImpl) *asyncCheckpointWriter {
	ret := &asyncCheckpointWriter{sync.Mutex{}, cp, make(chan interface{}, 1), nil, nil}
	go func() {
		deleteTicker := time.NewTicker(common.NewTimeBlocksInt(25).Duration())
		defer deleteTicker.Stop()
		for {
			select {
			case <-ret.notifyChan:
				ret.Lock()
				job := ret.nextJob
				if job != nil {
					ret.nextJob = nil
				}
				doneChansCopy := append([]chan struct{}{}, ret.doneChans...)
				ret.Unlock()
				if job != nil {
					job()
				}
				ret.Lock()
				for _, dc := range doneChansCopy {
					if dc != nil {
						close(dc)
					}
				}
				ret.Unlock()
			case <-deleteTicker.C:
				ret.Lock()
				ret.checkpointer.deleteSomeOldCheckpoints()
				ret.Unlock()
			case <-ctx.Done():
				ret.checkpointer.Close() //BUGBUG: must ensure this finishes before allowing db to be reopened
				return
			}
		}
	}()
	return ret
}

func (acw *asyncCheckpointWriter) SubmitJob(job func(), doneChan chan struct{}) {
	acw.Lock()
	defer acw.Unlock()
	acw.nextJob = job
	acw.doneChans = append(acw.doneChans, doneChan)
	select {
	case acw.notifyChan <- nil: // do nothing; only purpose was to send on the channel
	default: // no need to do anything, because channel already has something in it
	}
}
