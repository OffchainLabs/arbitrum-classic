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

package vm

import "github.com/offchainlabs/arbitrum/packages/arb-util/value"

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
