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
)

type AssertionDefender struct {
	assertion    *protocol.Assertion
	precondition *protocol.Precondition
	initState    Machine
}

func NewAssertionDefender(assertion *protocol.Assertion, precondition *protocol.Precondition, initState Machine) AssertionDefender {
	return AssertionDefender{assertion, precondition, initState.Clone()}
}

func (ad AssertionDefender) NumSteps() uint32 {
	return ad.assertion.NumSteps
}

func (ad AssertionDefender) GetAssertion() *protocol.Assertion {
	return ad.assertion
}

func (ad AssertionDefender) GetPrecondition() *protocol.Precondition {
	return ad.precondition
}

func (ad AssertionDefender) GetMachineState() Machine {
	return ad.initState
}

func (ad AssertionDefender) NBisect(slices uint32) []AssertionDefender {
	nsteps := ad.NumSteps()
	if nsteps < slices {
		slices = nsteps
	}
	sliceSize := nsteps / slices
	defenders := make([]AssertionDefender, 0, slices)
	m := ad.initState.Clone()

	pre := ad.precondition
	for i := uint32(0); i < slices; i++ {
		initState := m.Clone()

		stepCount := sliceSize
		if i < nsteps%slices {
			stepCount++
		}
		assertion := m.ExecuteAssertion(int32(stepCount), pre.TimeBounds)
		defenders = append(defenders, NewAssertionDefender(
			assertion,
			pre,
			initState,
		))
		pre = assertion.Stub().GeneratePostcondition(pre)
	}
	return defenders
}

func (ad AssertionDefender) SolidityOneStepProof() ([]byte, error) {
	return ad.initState.MarshalForProof()
}

func ChooseAssertionToChallenge(m Machine, assertions []*protocol.AssertionStub, preconditions []*protocol.Precondition) (uint16, Machine, error) {
	for i := range assertions {
		initState := m.Clone()
		generatedAssertion := m.ExecuteAssertion(
			int32(assertions[i].NumSteps),
			preconditions[i].TimeBounds,
		)
		if !generatedAssertion.Stub().Equals(assertions[i]) {
			return uint16(i), initState, nil
		}
	}
	return 0, nil, errors.New("all segments in false Assertion are valid")
}
