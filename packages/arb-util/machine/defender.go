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
	"math/rand"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

type AssertionDefender struct {
	assertion    *protocol.AssertionStub
	precondition *protocol.Precondition
	initState    Machine
}

func NewAssertionDefender(assertion *protocol.AssertionStub, precondition *protocol.Precondition, initState Machine) AssertionDefender {
	return AssertionDefender{assertion, precondition, initState.Clone()}
}

func (ad AssertionDefender) NumSteps() uint32 {
	return ad.assertion.NumSteps
}

func (ad AssertionDefender) GetAssertion() *protocol.AssertionStub {
	return ad.assertion
}

func (ad AssertionDefender) GetPrecondition() *protocol.Precondition {
	return ad.precondition
}

func (ad AssertionDefender) GetMachineState() Machine {
	return ad.initState
}

func (ad AssertionDefender) NBisect(slices uint32) []AssertionDefender {
	stepCounts := BisectionStepCounts(slices, ad.NumSteps())
	defenders := make([]AssertionDefender, 0, slices)
	m := ad.initState.Clone()

	prevAssertion := protocol.NewAssertion([32]byte{}, 0, nil, nil)
	pre := ad.precondition
	for i := uint32(0); i < slices; i++ {
		initState := m.Clone()
		chunkAssertion := m.ExecuteAssertion(int32(stepCounts[i]), pre.TimeBounds)
		if i > 0 {
			prevAssertion = protocol.NewAssertion(
				chunkAssertion.AfterHash,
				prevAssertion.NumSteps+chunkAssertion.NumSteps,
				append(prevAssertion.OutMsgs, chunkAssertion.OutMsgs...),
				append(prevAssertion.Logs, chunkAssertion.Logs...),
			)
		} else {
			prevAssertion = chunkAssertion
		}

		stub := prevAssertion.Stub()
		defenders = append(defenders, NewAssertionDefender(
			stub,
			pre,
			initState,
		))
		pre = stub.GeneratePostcondition(pre)
	}
	return defenders
}

func (ad AssertionDefender) SolidityOneStepProof() ([]byte, error) {
	return ad.initState.MarshalForProof()
}

func BisectionStepCounts(bisectionCount uint32, totalSteps uint32) []uint32 {
	stepCounts := make([]uint32, bisectionCount)
	sliceSize := totalSteps / bisectionCount
	for i := uint32(0); i < bisectionCount; i++ {
		stepCount := sliceSize
		if i < totalSteps%bisectionCount {
			stepCount++
		}
		stepCounts[i] = stepCount
	}
	return stepCounts
}

func updateAssertion(
	m Machine,
	timeBounds protocol.TimeBounds,
	prevAssertion *protocol.Assertion,
	stepCount uint32,
) *protocol.Assertion {
	chunkAssertion := m.ExecuteAssertion(
		int32(stepCount),
		timeBounds,
	)
	if prevAssertion != nil {
		return protocol.NewAssertion(
			chunkAssertion.AfterHash,
			prevAssertion.NumSteps+chunkAssertion.NumSteps,
			append(prevAssertion.OutMsgs, chunkAssertion.OutMsgs...),
			append(prevAssertion.Logs, chunkAssertion.Logs...),
		)
	} else {
		return chunkAssertion
	}
}

func ChooseAssertionToChallenge(
	m Machine,
	bisectionHashes [][32]byte,
	precondition *protocol.Precondition,
	totalSteps uint32,
) (uint16, Machine, error) {
	var prevAssertion *protocol.Assertion
	stepCounts := BisectionStepCounts(uint32(len(bisectionHashes)), totalSteps)
	for i := range bisectionHashes {
		initState := m.Clone()
		prevAssertion = updateAssertion(m, precondition.TimeBounds, prevAssertion, stepCounts[i])
		if prevAssertion.Stub().Hash() != bisectionHashes[i] {
			return uint16(i), initState, nil
		}
	}
	return 0, nil, errors.New("all segments in false Assertion are valid")
}

func ChooseRandomAssertionToChallenge(
	m Machine,
	bisectionHashes [][32]byte,
	precondition *protocol.Precondition,
	totalSteps uint32,
) (uint16, Machine) {
	assertionNum := uint16(rand.Int31n(int32(len(bisectionHashes))))
	var prevAssertion *protocol.Assertion
	stepCounts := BisectionStepCounts(uint32(len(bisectionHashes)), totalSteps)
	for i := range bisectionHashes {
		initState := m.Clone()
		prevAssertion = updateAssertion(m, precondition.TimeBounds, prevAssertion, stepCounts[i])
		if uint16(i) == assertionNum {
			return uint16(i), initState
		}
	}
	panic("Can't reach")
}
