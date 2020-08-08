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

package arbostest

import (
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/gotest"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/testvector"
	"testing"
)

func TestArbOSCases(t *testing.T) {
	arbosTests := gotest.ArbOSTestFiles()
	for _, testFile := range arbosTests {
		t.Run(testFile, func(t *testing.T) {
			if err := testvector.Check(testFile); err != nil {
				t.Error(err)
			}
		})
	}
}
