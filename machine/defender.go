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
	"io"

	"github.com/offchainlabs/arb-util/protocol"
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
	machine := ad.initState.Clone()

	timeBounds := ad.precondition.TimeBounds
	for i := uint32(0); i < slices; i++ {
		runState := machine.Clone()

		stepCount := sliceSize
		if i < nsteps%slices {
			stepCount++
		}
		defender, _ := ExecuteMachineAssertion(runState, int32(stepCount), timeBounds)
		defenders = append(defenders, defender)
		machine = runState
	}
	return defenders
}

func (ad AssertionDefender) SolidityOneStepProof(proofWr io.Writer) error {
	return ad.initState.MarshalForProof(proofWr)
}

func ChooseAssertionToChallenge(m Machine, assertions []*protocol.AssertionStub, preconditions []*protocol.Precondition) (uint16, Machine, error) {
	for i := range assertions {
		ad, _ := ExecuteMachineAssertion(
			m,
			int32(assertions[i].NumSteps),
			preconditions[i].TimeBounds,
		)
		generatedAssertion := ad.GetAssertion()
		if !generatedAssertion.Stub().Equals(assertions[i]) {
			return uint16(i), ad.initState, nil
		}
	}
	return 0, nil, errors.New("all segments in false Assertion are valid")
}
