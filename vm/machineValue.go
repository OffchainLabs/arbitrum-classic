package vm

import (
	"github.com/offchainlabs/arb-avm/value"
)

type MachineValue struct {
	value value.Value
	size  int64
}

func (r *MachineValue) Equal(comp *MachineValue) (bool, string) {
	if !value.Eq(r.value, comp.value) {
		return false, "MachineValues different"
	}
	if r.size != comp.size {
		return false, "MachineValues different size"
	}
	return true, ""
}

func NewMachineValue(val value.Value) *MachineValue {
	return &MachineValue{val, val.Size()}
}

func (r *MachineValue) Clone() *MachineValue {
	return &MachineValue{r.value.Clone(), r.size}
}

func (r *MachineValue) ProofValue() value.Value {
	return value.NewHashOnlyValueFromValue(r.value)
}

func (r *MachineValue) StateValue() value.Value {
	return value.NewHashOnlyValueFromValue(r.value)
}

func (r *MachineValue) Size() int64 {
	return r.size
}

func (r *MachineValue) Get() value.Value {
	return r.value
}

func (r *MachineValue) Set(value value.Value) {
	r.value = value
	r.size = value.Size()
}
