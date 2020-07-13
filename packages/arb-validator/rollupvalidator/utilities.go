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
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
)

func valueSlicesEqual(a []value.Value, b []value.Value) bool {
	if len(a) != len(b) {
		return false
	}
	for i, t := range a {
		if !value.Eq(t, b[i]) {
			return false
		}
	}
	return true
}

func stringSlicesEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, t := range a {
		if t != b[i] {
			return false
		}
	}
	return true
}

func logSlicesEqual(a []evm.Log, b []evm.Log) bool {
	if len(a) != len(b) {
		return false
	}
	for i, t := range a {
		if !t.Equals(b[i]) {
			return false
		}
	}
	return true
}

func nestedLogSlicesEqual(a [][]evm.Log, b [][]evm.Log) bool {
	if len(a) != len(b) {
		return false
	}
	for i, t := range a {
		if !logSlicesEqual(t, b[i]) {
			return false
		}
	}
	return true
}

func hashSlicesEqual(a []common.Hash, b []common.Hash) bool {
	if len(a) != len(b) {
		return false
	}
	for i, t := range a {
		if t != b[i] {
			return false
		}
	}
	return true
}
