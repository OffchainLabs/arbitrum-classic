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
	"bytes"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/gotest"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"io/ioutil"
	"testing"
)

func TestArbOSCases(t *testing.T) {
	arbosTests := gotest.ArbOSTestFiles()
	t.Log("Running", len(arbosTests), "test cases")
	for _, testFile := range arbosTests {
		data, err := ioutil.ReadFile(testFile)
		failIfError(t, err)

		t.Run(testFile, func(t *testing.T) {
			inboxMessages, avmLogs, avmSends, err := inbox.LoadTestVector(data)
			failIfError(t, err)

			calcLogs, calcSends, _, _ := runAssertion(t, inboxMessages, len(avmLogs), len(avmSends))

			for i, calcLog := range calcLogs {
				if !value.Eq(calcLog, avmLogs[i]) {
					calcRes, err := evm.NewResultFromValue(calcLog)
					res, err2 := evm.NewResultFromValue(avmLogs[i])
					if err == nil && err2 == nil {
						calcTxRes, ok1 := calcRes.(*evm.TxResult)
						txRes, ok2 := res.(*evm.TxResult)
						if ok1 && ok2 {
							for _, difference := range evm.CompareResults(calcTxRes, txRes) {
								t.Log(difference)
							}
						} else {
							t.Log("Calculated:", calcRes)
							t.Log("Correct:", res)
						}

					} else {
						if err != nil {
							t.Log("Error generating result", err)
						}
						if err2 != nil {
							t.Log("Error generating result2", err2)
						}
					}
					t.Error("wrong log")
				}
			}

			for i, calcSend := range calcSends {
				if !bytes.Equal(calcSend, avmSends[i]) {
					t.Error("wrong send")
				}
			}
		})
	}
}
