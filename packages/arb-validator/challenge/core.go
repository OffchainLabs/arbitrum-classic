/*
 * Copyright 2019, Offchain Labs, Inc.
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

package challenge

import (
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
)

type State interface {
	UpdateTime(uint64, bridge.ArbVMBridge) (State, error)
	UpdateState(ethbridge.Event, uint64, bridge.ArbVMBridge) (State, error)
}

type TimedOutChallenger struct {
	*core.Config
}

func (bot TimedOutChallenger) UpdateTime(time uint64, bridge bridge.ArbVMBridge) (State, error) {
	return bot, nil
}

func (bot TimedOutChallenger) UpdateState(ev ethbridge.Event, time uint64, brdg bridge.ArbVMBridge) (State, error) {
	switch ev.(type) {
	case ethbridge.ChallengerTimeoutEvent:
		return nil, nil
	default:
		return nil, &bridge.Error{nil, "ERROR: TimedOutChallenger: VM state got unsynchronized", false}
	}
}

type TimedOutAsserter struct {
	*core.Config
}

func (bot TimedOutAsserter) UpdateTime(time uint64, bridge bridge.ArbVMBridge) (State, error) {
	return bot, nil
}

func (bot TimedOutAsserter) UpdateState(ev ethbridge.Event, time uint64, brdg bridge.ArbVMBridge) (State, error) {
	switch ev.(type) {
	case ethbridge.AsserterTimeoutEvent:
		return nil, nil
	default:
		return nil, &bridge.Error{nil, "ERROR: TimedOutAsserter: VM state got unsynchronized", false}
	}
}
