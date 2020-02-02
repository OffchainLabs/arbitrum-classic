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

package rolluptest

import (
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

// evil machine is like a regular machine, except it returns a wrong hash w/ probability 1/8, repeatably
// this is useful for testing challenge functionality
type EvilMachine struct {
	*cmachine.Machine
}

func NewEvilMachine(machine *cmachine.Machine) *EvilMachine {
	return &EvilMachine{Machine: machine}
}

func (e EvilMachine) Clone() machine.Machine {
	return NewEvilMachine(e.Machine.Clone().(*cmachine.Machine))
}

func (e EvilMachine) Hash() common.Hash {
	return _tweakHash(e.Machine.Hash())
}

func _tweakHash(h common.Hash) common.Hash {
	// tweak the hash with probability 1/64; don't modify all-zero hash (it's special)
	// this is idempotent (calling it more than once has same effect as calling it once)
	if uint(h[0]) < 33 {
		h2 := h
		h2[0] = 0
		return h2
	} else {
		return h
	}
}

func (e EvilMachine) ExecuteAssertion(
	maxSteps uint64,
	timeBounds *protocol.TimeBoundsBlocks,
	inbox value.TupleValue,
	maxWallTime time.Duration,
) (*protocol.ExecutionAssertion, uint64) {
	assn, numSteps := e.Machine.ExecuteAssertion(maxSteps, timeBounds, inbox, maxWallTime)
	assn.AfterHash = _tweakHash(assn.AfterHash)
	return assn, numSteps
}
