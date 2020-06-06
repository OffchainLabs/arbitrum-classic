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

package evm

import (
	"math/rand"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func generateSampleLog() (Log, error) {
	var address common.Address
	var topic1 common.Hash
	var topic2 common.Hash
	var topic3 common.Hash
	data := make([]byte, 200)

	slicesToFill := [][]byte{address[:], topic1[:], topic2[:], topic3[:], data}
	for _, sl := range slicesToFill {
		_, err := rand.Read(sl)
		if err != nil {
			return Log{}, err
		}
	}
	return Log{
		Address: address,
		Topics: []common.Hash{
			topic1,
			topic2,
			topic3,
		},
		Data: data,
	}, nil
}

func TestLog(t *testing.T) {
	rand.Seed(43242)
	logs := make([]Log, 0)
	for i := 0; i < 10; i++ {
		l, err := generateSampleLog()
		if err != nil {
			t.Fatal(err)
		}

		logVal := l.AsValue()
		l2, err := NewLogFromValue(logVal)
		if err != nil {
			t.Fatal(err)
		}
		if !l.Equals(l2) {
			t.Error("unmarshaled log not equals to original")
		}

		logs = append(logs, l)
	}

	logVal := LogsToLogStack(logs)
	logs2, err := LogStackToLogs(logVal)

	if err != nil {
		t.Fatal(err)
	}

	if len(logs) != len(logs2) {
		t.Fatal("wrong long count")
	}

	for i, l1 := range logs {
		if !l1.Equals(logs2[i]) {
			t.Fatalf("logs not equal: %v and %v", l1, logs2[i])
		}
	}
}
