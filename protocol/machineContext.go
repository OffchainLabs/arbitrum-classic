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

package protocol

import (
	"github.com/offchainlabs/arb-avm/vm"
	"github.com/offchainlabs/arb-util/value"
)

type MachineAssertionContext struct {
	machine      *vm.Machine
	precondition *Precondition
	beforeInbox  value.Value
	numSteps     uint32
	afterBalance *BalanceTracker
	outMsgs      []Message
	logs         []value.Value
}

func NewMachineAssertionContext(m *vm.Machine, beforeBalance *BalanceTracker, timeBounds TimeBounds, beforeInbox value.Value) *MachineAssertionContext {
	outMsgs := make([]Message, 0)
	ret := &MachineAssertionContext{
		m,
		&Precondition{
			BeforeHash:    m.Hash(),
			TimeBounds:    timeBounds,
			BeforeBalance: beforeBalance,
			BeforeInbox:   value.NewHashOnlyValueFromValue(beforeInbox),
		},
		beforeInbox,
		0,
		beforeBalance.Clone(),
		outMsgs,
		make([]value.Value, 0),
	}
	ret.machine.SetContext(ret)
	return ret
}

func (ac *MachineAssertionContext) CanSpend(tokenType value.IntValue, currency value.IntValue) bool {
	tokenTypeBytes := tokenType.ToBytes()
	var tok TokenType
	// Cut off at 21 bytes
	copy(tok[:], tokenTypeBytes[:])
	return ac.afterBalance.CanSpend(tok, currency.BigInt())
}

func (ac *MachineAssertionContext) LoggedValue(data value.Value) error {
	ac.logs = append(ac.logs, data)
	return nil
}

func (ac *MachineAssertionContext) Send(data value.Value, tokenType value.IntValue, currency value.IntValue, dest value.IntValue) error {
	tokType := [21]byte{}
	tokBytes := tokenType.ToBytes()
	copy(tokType[:], tokBytes[:])
	newMsg := NewMessage(data, tokType, currency.BigInt(), dest.ToBytes())
	err := ac.afterBalance.Spend(tokType, newMsg.Currency)
	if err != nil {
		return err
	}
	ac.outMsgs = append(ac.outMsgs, newMsg)
	return nil

}

func (ac *MachineAssertionContext) OutMessageCount() int {
	return len(ac.outMsgs)
}

func (ac *MachineAssertionContext) ReadInbox() value.Value {
	return ac.beforeInbox
}

func (ac *MachineAssertionContext) GetTimeBounds() value.Value {
	return ac.precondition.TimeBounds.AsValue()
}

func (ac *MachineAssertionContext) NotifyStep() {
	ac.numSteps++
}

func (ac *MachineAssertionContext) GetAssertion() *Assertion {
	return NewAssertion(ac.machine.Hash(), ac.numSteps, ac.outMsgs, ac.logs)
}

func (ac *MachineAssertionContext) GetPostcondition() *Precondition {
	return NewPrecondition(ac.machine.Hash(), ac.precondition.TimeBounds, ac.afterBalance, ac.beforeInbox)
}

func (ac *MachineAssertionContext) Finalize(m *vm.Machine) AssertionDefender {
	ac.machine.SetContext(&vm.MachineNoContext{})
	return NewAssertionDefender(
		ac.GetAssertion(),
		NewPrecondition(
			ac.precondition.BeforeHash,
			ac.precondition.TimeBounds,
			NewBalanceTrackerFromMessages(ac.outMsgs),
			ac.precondition.BeforeInbox,
		),
		ac.beforeInbox,
		m,
	)
}

func (ac *MachineAssertionContext) EndContext() {
	ac.machine.SetContext(&vm.MachineNoContext{})
}
