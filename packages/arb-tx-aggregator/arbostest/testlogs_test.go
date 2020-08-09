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
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/gotest"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"io/ioutil"
	"testing"
)

func TestArbOSCases(t *testing.T) {
	arbosTests := gotest.ArbOSTestFiles()
	for _, testFile := range arbosTests {
		data, err := ioutil.ReadFile(testFile)
		if err != nil {
			t.Fatal(err)
		}
		t.Run(testFile, func(t *testing.T) {
			inboxMessages, avmLogs, avmSends, err := inbox.LoadTestVector(data)
			if err != nil {
				t.Fatal(err)
			}
			mach, err := cmachine.New(arbos.Path())
			if err != nil {
				t.Fatal(err)
			}
			assertion, _ := mach.ExecuteAssertion(100000000000, inboxMessages, 0)
			calcLogs := assertion.ParseLogs()
			calcSends := assertion.ParseOutMessages()

			commonLogCount := len(avmLogs)
			if len(calcLogs) < commonLogCount {
				commonLogCount = len(calcLogs)
			}

			commonSendCount := len(avmSends)
			if len(calcSends) < commonSendCount {
				commonSendCount = len(calcSends)
			}

			for i := 0; i < commonLogCount; i++ {
				calcRes, err := evm.NewTxResultFromValue(calcLogs[i])
				if err != nil {
					t.Fatal(err)
				}
				res, err := evm.NewTxResultFromValue(avmLogs[i])
				if err != nil {
					t.Fatal(err)
				}
				if !value.Eq(calcRes.AsValue(), res.AsValue()) {
					t.Log("Calculated:", calcRes)
					t.Log("Correct", res)
					t.Error("wrong log")
				}
			}

			for i := 0; i < commonSendCount; i++ {
				if !value.Eq(calcSends[i], avmSends[i]) {
					t.Error("wrong send")
				}
			}

			if len(calcLogs) != len(avmLogs) {
				t.Error("wrong log count")
			}
			if len(calcSends) != len(avmSends) {
				t.Error("wrong send count")
			}
		})

	}
}
