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

func LogMatchTest(
	t *testing.T,
	checkFunc func(addresses []common.Address, topics [][]common.Hash) bool,
	l Log,
) {
	if !checkFunc([]common.Address{}, [][]common.Hash{}) {
		t.Error("empty query should match anything")
	}

	if checkFunc([]common.Address{common.RandAddress()}, [][]common.Hash{}) {
		t.Error("query with different address shouldn't match")
	}

	if !checkFunc([]common.Address{}, [][]common.Hash{nil}) {
		t.Error("query with empty topic list should match")
	}

	if !checkFunc([]common.Address{}, [][]common.Hash{{l.Topics[0]}}) {
		t.Error("query with with correct topic should match")
	}

	if checkFunc([]common.Address{}, [][]common.Hash{{common.RandHash()}}) {
		t.Error("query with with different topic shouldn't match")
	}

	if !checkFunc([]common.Address{}, [][]common.Hash{{common.RandHash(), l.Topics[0]}}) {
		t.Error("query with with correct and incorrect topics should match")
	}

	if !checkFunc([]common.Address{}, [][]common.Hash{nil, {common.RandHash(), l.Topics[1]}}) {
		t.Error("query with with correct and incorrect second topics should match")
	}

	if checkFunc([]common.Address{}, [][]common.Hash{{common.RandHash()}, {l.Topics[1]}}) {
		t.Error("query with with incorrect first topic shouldn't match")
	}

	if checkFunc([]common.Address{}, [][]common.Hash{{l.Topics[0]}, {common.RandHash()}}) {
		t.Error("query with with incorrect second topic shouldn't match")
	}
}
