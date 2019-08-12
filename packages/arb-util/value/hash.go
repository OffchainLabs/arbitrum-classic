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
	"encoding/binary"
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type HashOnlyValue struct {
	hash [32]byte
	size int64
}

func NewHashOnlyValue(hash [32]byte, size int64) HashOnlyValue {
	return HashOnlyValue{hash, size}
}

func NewHashOnlyValueFromValue(val Value) HashOnlyValue {
	return HashOnlyValue{val.Hash(), val.Size()}
}

func NewHashOnlyValueFromReader(rd io.Reader) (HashOnlyValue, error) {
	var size int64
	err := binary.Read(rd, binary.LittleEndian, &size)
	if err != nil {
		return HashOnlyValue{}, err
	}
	var hash [32]byte
	if _, err := io.ReadFull(rd, hash[:]); err != nil {
		return HashOnlyValue{}, err
	}
	return HashOnlyValue{hash, size}, nil
}

func (nv HashOnlyValue) Marshal(wr io.Writer) error {
	// if err := binary.Write(wr, binary.LittleEndian, &nv.size); err != nil {
	//	return err
	//}
	_, err := wr.Write(nv.hash[:])
	return err
}

func (nv HashOnlyValue) TypeCode() byte {
	return TypeCodeHashOnly
}

func (nv HashOnlyValue) InternalTypeCode() byte {
	return TypeCodeHashOnly
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
	return nv.Hash() == val.Hash()
}

func (nv HashOnlyValue) String() string {
	return fmt.Sprintf("HashOnlyValue(%v)", hexutil.Encode(nv.hash[:]))
}

func (nv HashOnlyValue) Hash() [32]byte {
	return nv.hash
}
