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

package challenges

import (
	"errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type AssertionDefender struct {
	precondition *valprotocol.Precondition
	numSteps     uint64
	initState    machine.Machine
}

func NewAssertionDefender(precondition *valprotocol.Precondition, numSteps uint64, initState machine.Machine) AssertionDefender {
	return AssertionDefender{precondition, numSteps, initState.Clone()}
}

func (ad AssertionDefender) NumSteps() uint64 {
	return ad.numSteps
}

func (ad AssertionDefender) GetPrecondition() *valprotocol.Precondition {
	return ad.precondition
}

func (ad AssertionDefender) GetMachineState() machine.Machine {
	return ad.initState
}

func (ad AssertionDefender) NBisect(slices uint64) ([]AssertionDefender, []*valprotocol.ExecutionAssertionStub) {
	nsteps := ad.NumSteps()
	if nsteps < slices {
		slices = nsteps
	}
	defenders := make([]AssertionDefender, 0, slices)
	assertions := make([]*valprotocol.ExecutionAssertionStub, 0, slices)
	m := ad.initState.Clone()

	pre := ad.precondition
	for i := uint64(0); i < slices; i++ {
		steps := valprotocol.CalculateBisectionStepCount(i, slices, nsteps)
		initState := m.Clone()

		assertion, numSteps := m.ExecuteAssertion(
			steps,
			pre.BeforeInbox,
			0,
		)
		defenders = append(defenders, NewAssertionDefender(
			pre,
			numSteps,
			initState,
		))
		stub := valprotocol.NewExecutionAssertionStubFromAssertion(assertion)
		assertions = append(assertions, stub)
		pre = pre.GeneratePostcondition(stub)
	}
	return defenders, assertions
}

func (ad AssertionDefender) SolidityOneStepProof() ([]byte, error) {
	return ad.initState.MarshalForProof()
}

func ChooseAssertionToChallenge(
	m machine.Machine,
	pre *valprotocol.Precondition,
	assertions []*valprotocol.ExecutionAssertionStub,
	totalSteps uint64,
) (uint16, machine.Machine, error) {
	assertionCount := uint64(len(assertions))
	for i := range assertions {
		steps := valprotocol.CalculateBisectionStepCount(uint64(i), assertionCount, totalSteps)
		initState := m.Clone()
		generatedAssertion, numSteps := m.ExecuteAssertion(
			steps,
			pre.BeforeInbox,
			0,
		)
		stub := valprotocol.NewExecutionAssertionStubFromAssertion(generatedAssertion)
		if uint64(numSteps) != steps || !stub.Equals(assertions[i]) {
			return uint16(i), initState, nil
		}
		pre = pre.GeneratePostcondition(stub)
	}
	return 0, nil, errors.New("all segments in false ExecutionAssertion are valid")
}
