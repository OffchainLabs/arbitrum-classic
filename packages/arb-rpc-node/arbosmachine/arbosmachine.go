/*
* Copyright 2020-2021, Offchain Labs, Inc.
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
	"context"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

var logger = log.With().Stack().Str("component", "arbosmachine").Logger()

type Machine struct {
	machine.Machine
}

func New(mach machine.Machine) *Machine {
	return &Machine{Machine: mach}
}

func (m *Machine) Clone() machine.Machine {
	return &Machine{Machine: m.Machine.Clone()}
}

func (m *Machine) ExecuteAssertion(
	ctx context.Context,
	maxGas uint64,
	goOverGas bool,
	messages []inbox.InboxMessage,
) (*protocol.ExecutionAssertion, []value.Value, uint64, error) {
	assertion, debugPrints, numSteps, err := m.Machine.ExecuteAssertion(ctx, maxGas, goOverGas, messages)
	if err != nil {
		return nil, nil, 0, err
	}
	for _, d := range debugPrints {
		parsed, err := evm.NewLogLineFromValue(d)
		if err != nil {
			logger.Debug().Str("raw", d.String()).Msg("debugprint")
		} else {
			logger.Debug().EmbedObject(parsed).Msg("debugprint")
		}
	}
	return assertion, debugPrints, numSteps, nil
}
