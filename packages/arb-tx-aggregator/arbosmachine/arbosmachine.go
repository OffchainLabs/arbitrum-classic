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

package arbosmachine

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/rs/zerolog/log"
	"time"
)

var logger = log.With().Caller().Str("component", "arbosmachine").Logger()

type Machine struct {
	machine.Machine
}

func New(mach machine.Machine) *Machine {
	return &Machine{Machine: mach}
}

func (m *Machine) Clone() machine.Machine {
	return &Machine{Machine: m.Machine.Clone()}
}

func handleDebugPrint(value.Value) bool {
	return false
}

func handleDebugPrints(debugPrints []value.Value) {
	for _, d := range debugPrints {
		if !handleDebugPrint(d) {
			logger.Debug().Str("DebugPrint", d.String()).Send()
		}
	}
}

func (m *Machine) ExecuteAssertion(
	maxSteps uint64,
	messages []inbox.InboxMessage,
	maxWallTime time.Duration,
) (*protocol.ExecutionAssertion, []value.Value, uint64) {
	assertion, debugPrints, numSteps := m.Machine.ExecuteAssertion(maxSteps, messages, maxWallTime)
	handleDebugPrints(debugPrints)
	return assertion, debugPrints, numSteps
}

func (m *Machine) ExecuteCallServerAssertion(
	maxSteps uint64,
	inboxMessages []inbox.InboxMessage,
	fakeInboxPeekValue value.Value,
	maxWallTime time.Duration,
) (*protocol.ExecutionAssertion, []value.Value, uint64) {
	assertion, debugPrints, numSteps := m.Machine.ExecuteCallServerAssertion(maxSteps, inboxMessages, fakeInboxPeekValue, maxWallTime)
	handleDebugPrints(debugPrints)
	return assertion, debugPrints, numSteps
}
