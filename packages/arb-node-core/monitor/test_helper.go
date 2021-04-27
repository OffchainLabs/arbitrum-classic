/*
 * Copyright 2021, Offchain Labs, Inc.
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

package monitor

import (
	"io/ioutil"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func PrepareArbCore(t *testing.T, messages []inbox.InboxMessage) (*Monitor, func()) {
	tmpDir, err := ioutil.TempDir("", "arbitrum")
	test.FailIfError(t, err)
	arbosPath, err := arbos.Path()
	test.FailIfError(t, err)
	monitor, err := NewMonitor(tmpDir, arbosPath)
	if err != nil {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Fatal(err)
		}
	}

	shutdown := func() {
		monitor.Close()
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Fatal(err)
		}
	}
	returning := false
	defer (func() {
		if !returning {
			shutdown()
		}
	})()

	if len(messages) > 0 {
		var batchItems []inbox.SequencerBatchItem
		var prevAcc common.Hash
		for i, msg := range messages {
			batchItem := inbox.SequencerBatchItem{
				LastSeqNum:        big.NewInt(int64(i)),
				TotalDelayedCount: big.NewInt(0),
				SequencerMessage:  msg.ToBytes(),
			}
			if err := batchItem.RecomputeAccumulator(prevAcc, big.NewInt(0), common.Hash{}); err != nil {
				t.Fatal(err)
			}
			batchItems = append(batchItems, batchItem)
			prevAcc = batchItem.Accumulator
		}

		_, err = core.DeliverMessagesAndWait(monitor.Core, common.Hash{}, batchItems, nil, nil)
		test.FailIfError(t, err)
	}
	for {
		msgCount, err := monitor.Core.GetMessageCount()
		test.FailIfError(t, err)
		if monitor.Core.MachineIdle() && msgCount.Cmp(big.NewInt(int64(len(messages)))) >= 0 {
			break
		}
		<-time.After(time.Millisecond * 200)
	}

	returning = true
	return monitor, shutdown
}
