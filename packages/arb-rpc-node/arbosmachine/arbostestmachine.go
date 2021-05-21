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

package arbosmachine

import (
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type TestMachine struct {
	t *testing.T
	machine.Machine
}

func NewTestMachine(t *testing.T, mach machine.Machine) *TestMachine {
	return &TestMachine{t: t, Machine: mach}
}

func (m *TestMachine) Clone() machine.Machine {
	return &TestMachine{t: m.t, Machine: m.Machine.Clone()}
}

func (m *TestMachine) ExecuteAssertion(
	maxGas uint64,
	goOverGas bool,
	messages []inbox.InboxMessage,
	finalMessageOfBlock bool,
) (*protocol.ExecutionAssertion, []value.Value, uint64, error) {
	assertion, debugPrints, numSteps, err := m.Machine.ExecuteAssertion(maxGas, goOverGas, messages, finalMessageOfBlock)
	if err != nil {
		return nil, nil, 0, err
	}
	for _, d := range debugPrints {
		parsed, err := handleDebugPrint(d)
		if err != nil {
			m.t.Log("raw debugprint", d)
		} else {
			m.t.Log("debugprint", parsed)
		}
	}
	return assertion, debugPrints, numSteps, nil
}
