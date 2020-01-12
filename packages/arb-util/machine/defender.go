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

package machine

import (
	"errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type AssertionDefender struct {
	precondition *protocol.Precondition
	numSteps     uint32
	initState    Machine
}

func NewAssertionDefender(precondition *protocol.Precondition, numSteps uint32, initState Machine) AssertionDefender {
	return AssertionDefender{precondition, numSteps, initState.Clone()}
}

func (ad AssertionDefender) NumSteps() uint32 {
	return ad.numSteps
}

func (ad AssertionDefender) GetPrecondition() *protocol.Precondition {
	return ad.precondition
}

func (ad AssertionDefender) GetMachineState() Machine {
	return ad.initState
}

func (ad AssertionDefender) NBisect(slices uint32) ([]AssertionDefender, []*protocol.ExecutionAssertionStub) {
	nsteps := ad.NumSteps()
	if nsteps < slices {
		slices = nsteps
	}
	defenders := make([]AssertionDefender, 0, slices)
	assertions := make([]*protocol.ExecutionAssertionStub, 0, slices)
	m := ad.initState.Clone()

	pre := ad.precondition
	for i := uint32(0); i < slices; i++ {
		steps := CalculateBisectionStepCount(uint32(i), slices, nsteps)
		initState := m.Clone()

		assertion, numSteps := m.ExecuteAssertion(steps, pre.TimeBounds, pre.BeforeInbox.(value.TupleValue))
		defenders = append(defenders, NewAssertionDefender(
			pre,
			numSteps,
			initState,
		))
		stub := assertion.Stub()
		assertions = append(assertions, stub)
		pre = stub.GeneratePostcondition(pre)
	}
	return defenders, assertions
}

func (ad AssertionDefender) SolidityOneStepProof() ([]byte, error) {
	return ad.initState.MarshalForProof()
}

func CalculateBisectionStepCount(chunkIndex, segmentCount, totalSteps uint32) uint32 {
	if chunkIndex == 0 {
		return totalSteps/segmentCount + totalSteps%segmentCount
	} else {
		return totalSteps / segmentCount
	}
}

func ChooseAssertionToChallenge(
	m Machine,
	pre *protocol.Precondition,
	assertions []*protocol.ExecutionAssertionStub,
	totalSteps uint32,
) (uint16, Machine, error) {
	assertionCount := uint32(len(assertions))
	for i := range assertions {
		steps := CalculateBisectionStepCount(uint32(i), assertionCount, totalSteps)
		initState := m.Clone()
		generatedAssertion, numSteps := m.ExecuteAssertion(
			steps,
			pre.TimeBounds,
			pre.BeforeInbox.(value.TupleValue),
		)
		stub := generatedAssertion.Stub()
		if numSteps != steps || !stub.Equals(assertions[i]) {
			return uint16(i), initState, nil
		}
		pre = stub.GeneratePostcondition(pre)
	}
	return 0, nil, errors.New("all segments in false ExecutionAssertion are valid")
}
