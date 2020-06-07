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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/rand"
	"testing"
)

func TestLog(t *testing.T) {
	rand.Seed(43242)
	logs := make([]Log, 0)
	for i := 0; i < 10; i++ {
		l := NewRandomLog(3)
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

func TestLogFilter(t *testing.T) {
	l := NewRandomLog(3)

	LogMatchTest(
		t,
		func(addresses []common.Address, topics [][]common.Hash) bool {
			return l.MatchesQuery(addresses, topics)
		},
		l,
	)

	if l.MatchesQuery([]common.Address{}, [][]common.Hash{nil, nil, nil, nil}) {
		t.Error("query with too many topics shouldn't match")
	}
}
