package stack

import (
	"bytes"
	"fmt"

	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-avm/vm/warning"
)

type Tuple struct {
	stack value.Value
	count int64
}

func NewTuple(stack value.Value) *Tuple {
	ret := &Tuple{stack, 0}
	return ret
}

func (m *Tuple) Equal(yin Stack) (bool, string) {
	// TODO add tuple equal functionality
	return true, ""
}

func (m *Tuple) Clone() Stack {
	return &Tuple{m.stack.Clone(), m.count}
}

func (m *Tuple) Push(v value.Value) {
	m.stack = value.NewTuple2(v, m.stack)
	m.count++
}

func (m *Tuple) PushInt(v value.IntValue) {
	m.Push(v)
}

func (m *Tuple) PushTuple(v value.TupleValue) {
	m.Push(v)
}

func (m *Tuple) PushCodePoint(v value.CodePointValue) {
	m.Push(v)
}

func (m *Tuple) Pop() (value.Value, error) {
	topTuple, ok := m.stack.(value.TupleValue)
	if !ok {
		return nil, warning.New(fmt.Sprintf("Stack.Pop: Value in Stack was %s instead of a tuple", value.TypeCodeName(m.stack.TypeCode())))
	}
	if topTuple.Len() != 2 {
		return nil, warning.New("Stack.Pop: Tuple in Stack was not size 2")
	}
	m.stack, _ = topTuple.GetByInt64(1)
	m.count--
	return topTuple.GetByInt64(0)
}

func (m *Tuple) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	s := m.Clone()
	for !s.IsEmpty() {
		val, _ := s.Pop()
		buf.WriteString(fmt.Sprintf("%v", val))
		if !s.IsEmpty() {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

func (m *Tuple) PopInt() (value.IntValue, error) {
	val, err := m.Pop()
	if err != nil {
		return value.IntValue{}, err
	}
	v, ok := val.(value.IntValue)
	if !ok {
		return v, nil
	}
	return value.IntValue{}, TypeError{"Int", val}
}

func (m *Tuple) PopTuple() (value.TupleValue, error) {
	val, err := m.Pop()
	if err != nil {
		return value.TupleValue{}, err
	}
	v, ok := val.(value.TupleValue)
	if !ok {
		return value.TupleValue{}, TypeError{"Tuple", val}
	}
	return v, nil
}

func (m *Tuple) PopCodePoint() (value.CodePointValue, error) {
	val, err := m.Pop()
	if err != nil {
		return value.CodePointValue{}, err
	}
	v, ok := val.(value.CodePointValue)
	if !ok {
		return value.CodePointValue{}, TypeError{"CodePointValue", val}
	}
	return v, nil
}

func (m *Tuple) IsEmpty() bool {
	topTuple := m.stack.(value.TupleValue)
	return topTuple.Len() == 0
}

func (m *Tuple) Size() int64 {
	return m.stack.Size()
}

func (m *Tuple) Count() int64 {
	return m.count
}

func (m *Tuple) StateValue() value.Value {
	return value.NewHashOnlyValueFromValue(m.stack)
}

func (m *Tuple) ProofValue(stackInfo []byte) value.Value {
	c := m.Clone()
	vals := make([]value.Value, 0, len(stackInfo))
	for range stackInfo {
		val, _ := c.Pop()
		vals = append(vals, val)
	}
	stack := NewTuple(c.StateValue())
	for i := len(stackInfo) - 1; i >= 0; i-- {
		if stackInfo[i] == 1 {
			stack.Push(vals[i].CloneShallow())
		} else {
			stack.Push(value.NewHashOnlyValueFromValue(vals[i]))
		}
	}
	return stack.stack
}

func (m *Tuple) SolidityProofValue(stackInfo []byte) (value.HashOnlyValue, []value.Value) {
	c := m.Clone()
	vals := make([]value.Value, 0, len(stackInfo))
	for i := range stackInfo {
		val, _ := c.Pop()
		if stackInfo[i] == 1 {
			vals = append(vals, val.CloneShallow())
		} else {
			vals = append(vals, value.NewHashOnlyValueFromValue(val))
		}
	}
	return value.NewHashOnlyValueFromValue(c.StateValue()), vals
}

func (m *Tuple) FullyExpandedValue() value.Value {
     return m.stack
}
