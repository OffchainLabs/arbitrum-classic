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

package testmachine

import "C"
import (
	"fmt"
	"log"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Machine struct {
	cppmachine *cmachine.Machine
}

func New(codeFile string, warnMode bool) (*Machine, error) {
	cm, cmerr := cmachine.New(codeFile)

	if cmerr != nil {
		err := fmt.Errorf("cpp machine error: %v ", cmerr)
		return nil, err
	} else {
		return &Machine{
			cm,
		}, nil
	}
}

func (m *Machine) Hash() common.Hash {
	h1 := m.cppmachine.Hash()
	return h1
}

func (m *Machine) PrintState() {
	log.Println("Cpp state")
	m.cppmachine.PrintState()
}

func (m *Machine) Clone() machine.Machine {
	return &Machine{m.cppmachine.Clone().(*cmachine.Machine)}
}

func (m *Machine) CurrentStatus() machine.Status {
	return m.cppmachine.CurrentStatus()
}

func (m *Machine) IsBlocked(currentTime *common.TimeBlocks, newMessages bool) machine.BlockReason {
	return m.cppmachine.IsBlocked(currentTime, newMessages)
}

func (m *Machine) ExecuteAssertion(
	maxSteps uint64,
	timeBounds *protocol.TimeBounds,
	inbox value.TupleValue,
	maxWallTime time.Duration,
) (*protocol.ExecutionAssertion, uint64) {
	hasTimeLimit := maxWallTime.Nanoseconds() != 0
	startTime := time.Now()
	timeLeft := maxWallTime
	a := &protocol.ExecutionAssertion{
		AfterHash:    m.cppmachine.Hash(),
		DidInboxInsn: false,
		NumGas:       0,
		OutMsgs:      nil,
		Logs:         nil,
	}
	totalSteps := uint64(0)
	stepIncrease := uint64(5000)
	for i := uint64(0); i < maxSteps; i += stepIncrease {
		steps := stepIncrease
		if i+stepIncrease > maxSteps {
			steps = maxSteps - i
		}

		a1, ranSteps1 := m.cppmachine.ExecuteAssertion(steps, timeBounds, inbox, timeLeft)

		a.AfterHash = a1.AfterHash
		totalSteps += ranSteps1
		a.NumGas += a1.NumGas
		a.Logs = append(a.Logs, a1.Logs...)
		a.OutMsgs = append(a.OutMsgs, a1.OutMsgs...)
		a.DidInboxInsn = a.DidInboxInsn || a1.DidInboxInsn
		if a1.DidInboxInsn {
			inbox = value.NewEmptyTuple()
		}
		if ranSteps1 < steps {
			break
		}
		if hasTimeLimit {
			elapsedTime := time.Now().Sub(startTime)
			if elapsedTime > maxWallTime {
				break
			}
			timeLeft = maxWallTime - elapsedTime
		}

	}
	fmt.Println("Ran", totalSteps, "steps")
	return a, totalSteps
}

func (m *Machine) MarshalForProof() ([]byte, error) {
	h1, err1 := m.cppmachine.MarshalForProof()

	if err1 != nil {
		return nil, err1
	}
	return h1, nil
}

func (m *Machine) MarshalState() ([]byte, error) {
	h1, err1 := m.cppmachine.MarshalState()

	if err1 != nil {
		return nil, err1
	}
	return h1, nil
}

func (m *Machine) Checkpoint(storage machine.CheckpointStorage) bool {
	return m.cppmachine.Checkpoint(storage)
}
