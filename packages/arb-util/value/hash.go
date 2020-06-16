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

package value

import (
	"fmt"
	"io"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type HashOnlyValue struct {
	hash common.Hash
	size int64
}

func NewHashOnlyValueFromValue(val Value) HashOnlyValue {
	return HashOnlyValue{val.Hash(), val.Size()}
}

func NewHashOnlyValueFromReader(rd io.Reader) (HashOnlyValue, error) {
	var hash common.Hash
	if _, err := io.ReadFull(rd, hash[:]); err != nil {
		return HashOnlyValue{}, err
	}

	intVal, err2 := NewIntValueFromReader(rd)

	if err2 != nil {
		return HashOnlyValue{}, err2
	}
	size := intVal.val.Int64()

	return HashOnlyValue{hash, size}, nil
}

func (nv HashOnlyValue) Marshal(wr io.Writer) error {
	// if err := binary.Write(wr, binary.LittleEndian, &nv.size); err != nil {
	//	return err
	//}
	_, err := wr.Write(nv.hash[:])
	sizeVal := NewInt64Value(nv.Size())
	sizeVal.Marshal(wr)

	return err
}

func (tv HashOnlyValue) MarshalForProof(wr io.Writer) error {
	return tv.Marshal(wr)
}

func (nv HashOnlyValue) TypeCode() byte {
	return TypeCodeHashPreImage
}

func (nv HashOnlyValue) InternalTypeCode() byte {
	return TypeCodeHashPreImage
}

func (nv HashOnlyValue) Clone() Value {
	return HashOnlyValue{nv.hash, nv.size}
}

func (nv HashOnlyValue) CloneShallow() Value {
	return HashOnlyValue{nv.hash, nv.size}
}

func (nv HashOnlyValue) Size() int64 {
	return nv.size
}

func (nv HashOnlyValue) Equal(val Value) bool {
	return nv.Hash() == val.Hash() && nv.Size() == val.Size()
}

func (nv HashOnlyValue) String() string {
	return fmt.Sprintf("HashOnlyValue(%v)", nv.hash)
}

func (nv HashOnlyValue) Hash() common.Hash {
	return nv.hash
}
