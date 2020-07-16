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

package rollupvalidator

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func TestProcessNode(t *testing.T) {
	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		t.Fatal(err)
	}
	chain := common.RandAddress()
	initialNode := structures.NewInitialNode(mach, common.Hash{})
	_, err = processNode(initialNode, chain)
	if err != nil {
		t.Fatal(err)
	}
}

func logListMatches(a []evm.FullLog, b []evm.FullLog) bool {
	if len(a) != len(b) {
		return false
	}
	for i, l := range a {
		if !l.Equals(b[i]) {
			return false
		}
	}
	return true
}

func TestFindLogs(t *testing.T) {
	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		t.Fatal(err)
	}
	chain := common.RandAddress()

	results := make([]*evm.Result, 0, 5)
	for i := int32(0); i < 5; i++ {
		stop := evm.NewRandomResult(message.NewRandomEth(), 2)
		results = append(results, stop)
	}

	initialNode := structures.NewInitialNode(mach, common.Hash{})
	nextNode := structures.NewRandomNodeFromValidPrev(initialNode, results)
	info, err := processNode(nextNode, chain)
	if err != nil {
		t.Fatal(err)
	}
	flatLogs := info.fullLogs()

	if !logListMatches(info.FindLogs(nil, nil), flatLogs) {
		t.Error("empty query should match everything")
	}

	if !logListMatches(info.FindLogs(nil, nil), flatLogs) {
		t.Error("empty query should match everything")
	}

	if !logListMatches(
		info.FindLogs(
			[]common.Address{flatLogs[0].Log.Address},
			nil,
		),
		flatLogs[:1],
	) {
		t.Error("query result wrong")
	}

	if !logListMatches(
		info.FindLogs(
			[]common.Address{
				flatLogs[0].Log.Address,
				flatLogs[1].Log.Address,
				flatLogs[2].Log.Address,
			},
			nil,
		),
		flatLogs[:3],
	) {
		t.Error("query result wrong")
	}

}
