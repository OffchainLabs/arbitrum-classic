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
	"math/big"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func PrepareArbCore(t *testing.T) (*Monitor, func()) {
	arbosPath, err := arbos.Path()
	test.FailIfError(t, err)
	return PrepareArbCoreWithMexe(t, arbosPath)
}

func PrepareArbCoreWithMexe(t *testing.T, mexe string) (*Monitor, func()) {
	coreConfig := configuration.DefaultCoreSettings()
	monitor, err := NewMonitor(t.TempDir(), mexe, coreConfig)
	test.FailIfError(t, err)

	shutdown := func() {
		monitor.Close()
	}
	returning := false
	defer (func() {
		if !returning {
			shutdown()
		}
	})()

	for {
		test.FailIfError(t, err)
		if monitor.Core.MachineIdle() {
			break
		}
		<-time.After(time.Millisecond * 200)
	}

	returning = true
	return monitor, shutdown
}

func DeliverMessagesToCore(t *testing.T, arbCore core.ArbCore, delayedCount *big.Int, prevCount *big.Int, prevAcc common.Hash, messages []inbox.InboxMessage) {
	startAcc := prevAcc
	var batchItems []inbox.SequencerBatchItem
	for _, msg := range messages {
		batchItem := inbox.NewSequencerItem(delayedCount, msg, prevAcc)
		batchItems = append(batchItems, batchItem)
		prevAcc = batchItem.Accumulator
	}

	beforeCount, err := arbCore.GetMessageCount()
	test.FailIfError(t, err)

	target := new(big.Int).Add(beforeCount, big.NewInt(int64(len(messages))))

	err = core.DeliverMessagesAndWait(arbCore, prevCount, startAcc, batchItems, nil, nil)
	test.FailIfError(t, err)

	for {
		msgCount, err := arbCore.GetMessageCount()
		test.FailIfError(t, err)
		if arbCore.MachineIdle() && msgCount.Cmp(target) == 0 {
			break
		}
		<-time.After(time.Millisecond * 200)
	}
}
