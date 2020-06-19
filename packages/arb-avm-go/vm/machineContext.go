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
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Context interface {
	Send(message value.Value)
	GetLowerBoundBlock() value.IntValue
	GetUpperBoundBlock() value.IntValue
	GetLowerBoundTimestamp() value.IntValue
	GetUpperBoundTimestamp() value.IntValue
	NotifyStep(uint64)
	LoggedValue(value.Value)
	GetInbox() value.TupleValue
	ReadInbox()

	OutMessageCount() int
}

type NoContext struct{}

func (m *NoContext) LoggedValue(data value.Value) {

}

func (m *NoContext) GetInbox() value.TupleValue {
	return value.NewEmptyTuple()
}

func (m *NoContext) ReadInbox() {

}

func (m *NoContext) Send(message value.Value) {

}

func (m *NoContext) OutMessageCount() int {
	return 0
}

func (m *NoContext) GetLowerBoundBlock() value.IntValue {
	return value.NewInt64Value(0)
}

func (m *NoContext) GetUpperBoundBlock() value.IntValue {
	return value.NewInt64Value(0)
}

func (m *NoContext) GetLowerBoundTimestamp() value.IntValue {
	return value.NewInt64Value(0)
}

func (m *NoContext) GetUpperBoundTimestamp() value.IntValue {
	return value.NewInt64Value(0)
}

func (m *NoContext) NotifyStep(uint64) {
}

type MachineAssertionContext struct {
	machine      *Machine
	timeBounds   *protocol.TimeBounds
	inbox        value.TupleValue
	didInboxInsn bool
	numSteps     uint64
	numGas       uint64
	outMsgs      []value.Value
	logs         []value.Value
}

func NewMachineAssertionContext(m *Machine, timeBounds *protocol.TimeBounds, inbox value.TupleValue) *MachineAssertionContext {
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

func (ac *MachineAssertionContext) StepCount() uint64 {
	return ac.numSteps
}

func (ac *MachineAssertionContext) GasCount() uint64 {
	return ac.numGas
}

func (ac *MachineAssertionContext) OutMessageCount() int {
	return len(ac.outMsgs)
}

func (m *MachineAssertionContext) GetLowerBoundBlock() value.IntValue {
	return value.NewIntValue(m.timeBounds.LowerBoundBlock.AsInt())
}

func (m *MachineAssertionContext) GetUpperBoundBlock() value.IntValue {
	return value.NewIntValue(m.timeBounds.UpperBoundBlock.AsInt())
}

func (m *MachineAssertionContext) GetLowerBoundTimestamp() value.IntValue {
	return value.NewIntValue(m.timeBounds.LowerBoundTimestamp)
}

func (m *MachineAssertionContext) GetUpperBoundTimestamp() value.IntValue {
	return value.NewIntValue(m.timeBounds.UpperBoundTimestamp)
}

func (ac *MachineAssertionContext) NotifyStep(numGas uint64) {
	ac.numSteps++
	ac.numGas = ac.numGas + numGas
}

func (ac *MachineAssertionContext) Finalize(m *Machine) (*protocol.ExecutionAssertion, uint64) {
	ac.machine.SetContext(&NoContext{})
	return protocol.NewExecutionAssertionFromValues(ac.machine.Hash(), ac.didInboxInsn, ac.numGas, ac.outMsgs, ac.logs), ac.numSteps
}

func (ac *MachineAssertionContext) EndContext() {
	ac.machine.SetContext(&NoContext{})
}
