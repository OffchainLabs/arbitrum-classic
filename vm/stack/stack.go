package stack

import (
	"fmt"

	"github.com/offchainlabs/arb-avm/value"
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

	Pop() (value.Value, error)
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
