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
	"github.com/offchainlabs/arbitrum/packages/arb-avm-go/goloader"
	gomachine "github.com/offchainlabs/arbitrum/packages/arb-avm-go/vm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Machine struct {
	cppmachine *cmachine.Machine
	gomachine  *gomachine.Machine
}

func New(codeFile string, warnMode bool) (*Machine, error) {
	gm, gmerr := goloader.LoadMachineFromFile(codeFile, warnMode)
	cm, cmerr := cmachine.New(codeFile)
	var err error
	if gmerr != nil {
		if cmerr != nil {
			err = fmt.Errorf("Go machine error: %v, cpp machine error: %v ", gmerr, cmerr)
		} else {
			err = fmt.Errorf("Go machine error: %v", gmerr)
		}
	} else if cmerr != nil {
		err = fmt.Errorf("cpp machine error: %v ", cmerr)
	}
	return &Machine{
		cm,
		gm,
	}, err
}

func (m *Machine) Hash() common.Hash {
	h1 := m.cppmachine.Hash()
	h2 := m.gomachine.Hash()
	if h1 != h2 {
		log.Fatalln("Hash error at pc", m.gomachine.GetPC())
	}
	return h1
}

func (m *Machine) PrintState() {
	log.Println("Cpp state")
	m.cppmachine.PrintState()
	//log.Println("Go state")
	//m.gomachine.PrintState()
}

func (m *Machine) Clone() machine.Machine {
	return &Machine{m.cppmachine.Clone().(*cmachine.Machine), m.gomachine.Clone().(*gomachine.Machine)}
}

func (m *Machine) CurrentStatus() machine.Status {
	b1 := m.cppmachine.CurrentStatus()
	//b2 := m.gomachine.CurrentStatus()
	//if b1 != b2 {
	//	log.Fatalln("CurrentStatus error at pc", m.gomachine.GetPC())
	//}
	return b1
}

func (m *Machine) IsBlocked(currentTime *common.TimeBlocks, newMessages bool) machine.BlockReason {
	b1 := m.cppmachine.IsBlocked(currentTime, newMessages)
	//b2 := m.gomachine.IsBlocked(currentTime, newMessages)
	//if b1 == nil || b2 == nil {
	//	if b1 != nil || b2 != nil {
	//		log.Fatalln("IsBlocked error at pc", m.gomachine.GetPC())
	//	}
	//	return nil
	//}
	//if !b1.Equals(b2) {
	//	log.Fatalln("IsBlocked error at pc", m.gomachine.GetPC())
	//}
	return b1
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

		fmt.Println(m.gomachine.GetPC())
		m.gomachine.PrintState()
		_, _ = m.gomachine.ExecuteAssertion(steps, timeBounds, inbox, timeLeft)
		a1, ranSteps1 := m.cppmachine.ExecuteAssertion(steps, timeBounds, inbox, timeLeft)

		//if ranSteps1 != ranSteps2 {
		//	_ = m.gomachine.GetPC()
		//	log.Println("cpp num steps", ranSteps1, a1.NumGas)
		//	log.Println("go num steps", ranSteps2, a2.NumGas)
		//	log.Fatalln("ExecuteAssertion error after running step", pcStart, pcEnd, a1, a2)
		//} else if !a1.Equals(a2) {
		//	pcEnd := m.gomachine.GetPC()
		//	m.cppmachine.PrintState()
		//	m.gomachine.PrintState()
		//	log.Println("cpp num steps, num gas", ranSteps1, a1.NumGas)
		//	log.Println("go num steps, num gas", ranSteps2, a2.NumGas)
		//	log.Println("pcstart", pcStart)
		//	log.Println("pcend", pcEnd)
		//	log.Fatalln("ExecuteAssertion error after running step", a1, a2)
		//}
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
	//h2, err2 := m.gomachine.MarshalForProof()
	h1, err1 := m.cppmachine.MarshalForProof()

	if err1 != nil {
		return nil, err1
	}

	//if err2 != nil {
	//	return nil, err2
	//}

	//if !bytes.Equal(h1, h2) {
	//	m.cppmachine.PrintState()
	//	m.gomachine.PrintState()
	//	log.Fatalln("MarshalForProof error at pc", m.gomachine.GetPC())
	//}

	return h1, nil
}

func (m *Machine) Checkpoint(storage machine.CheckpointStorage) bool {
	h1 := m.cppmachine.Checkpoint(storage)
	//h2 := m.gomachine.Checkpoint(storage)
	//if h1 != h2 {
	//	log.Fatalln("Checkpoint error at pc", m.gomachine.GetPC())
	//}
	return h1
}
