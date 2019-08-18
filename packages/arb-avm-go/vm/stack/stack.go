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

package stack

import (
	"fmt"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type EmptyError struct{}

func (e EmptyError) Error() string {
	return "tried to pop empty stack"
}

type TypeError struct {
	expectedType string
	value        value.Value
}

func (e TypeError) Error() string {
	return fmt.Sprintf("popped stack expecting %s but received %v", e.expectedType, e.value)
}

type Stack interface {
	Clone() Stack

	Push(value.Value)
	PushInt(value.IntValue)
	PushTuple(value.TupleValue)
	PushCodePoint(value.CodePointValue)

	Pop() (value.Value, error) // Only error than can return is EmptyError
	PopInt() (value.IntValue, error)
	PopTuple() (value.TupleValue, error)
	PopCodePoint() (value.CodePointValue, error)

	Equal(Stack) (bool, string) // current usage is for testing only. Revisit return value if other usage identified
	IsEmpty() bool
	Size() int64
	Count() int64

	StateValue() value.Value
	ProofValue([]byte) value.Value
	SolidityProofValue([]byte) (value.HashOnlyValue, []value.Value)
	FullyExpandedValue() value.Value
}
