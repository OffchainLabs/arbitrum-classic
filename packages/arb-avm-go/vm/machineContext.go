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

package vm

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type MachineAssertionContext struct {
	machine      *Machine
	timeBounds   *protocol.TimeBoundsBlocks
	inbox        value.TupleValue
	didInboxInsn bool
	numSteps     uint32
	numGas       uint64
	outMsgs      []value.Value
	logs         []value.Value
}

func NewMachineAssertionContext(m *Machine, timeBounds *protocol.TimeBoundsBlocks, inbox value.TupleValue) *MachineAssertionContext {
	ret := &MachineAssertionContext{
		m,
		timeBounds,
		inbox,
		false,
		0,
		0,
		make([]value.Value, 0),
		make([]value.Value, 0),
	}
	ret.machine.SetContext(ret)
	return ret
}

func (ac *MachineAssertionContext) LoggedValue(data value.Value) {
	ac.logs = append(ac.logs, data)
}

func (m *MachineAssertionContext) GetInbox() value.TupleValue {
	return m.inbox
}

func (m *MachineAssertionContext) ReadInbox() {
	m.didInboxInsn = true
	m.inbox = value.NewEmptyTuple()
}

func (ac *MachineAssertionContext) Send(message value.Value) {
	ac.outMsgs = append(ac.outMsgs, message)
}

func (ac *MachineAssertionContext) StepCount() uint32 {
	return ac.numSteps
}

func (ac *MachineAssertionContext) GasCount() uint64 {
	return ac.numGas
}

func (ac *MachineAssertionContext) OutMessageCount() int {
	return len(ac.outMsgs)
}

func (ac *MachineAssertionContext) GetTimeBounds() value.TupleValue {
	return ac.timeBounds.AsValue()
}

func (ac *MachineAssertionContext) NotifyStep(numGas uint64) {
	ac.numSteps++
	ac.numGas = ac.numGas + numGas
}

func (ac *MachineAssertionContext) Finalize(m *Machine) (*protocol.ExecutionAssertion, uint32) {
	ac.machine.SetContext(&machine.NoContext{})
	return protocol.NewExecutionAssertion(ac.machine.Hash(), ac.didInboxInsn, ac.numGas, ac.outMsgs, ac.logs), ac.numSteps
}

func (ac *MachineAssertionContext) EndContext() {
	ac.machine.SetContext(&machine.NoContext{})
}
