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
	"testing"
)

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
