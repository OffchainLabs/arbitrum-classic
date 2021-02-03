/*
 * Copyright 2020, Offchain Labs, Inc.
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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"io"
)

type HashPreImage struct {
	hashImage common.Hash
	size      int64
}

func NewPreImage(hashImage common.Hash, size int64) HashPreImage {
	return HashPreImage{hashImage, size}
}

func NewHashPreImageFromReader(rd io.Reader) (HashPreImage, error) {
	var h common.Hash
	_, err := io.ReadFull(rd, h[:])
	if err != nil {
		return HashPreImage{}, err
	}

	intVal, err := NewIntValueFromReader(rd)
	if err != nil {
		return HashPreImage{}, err
	}

	size := intVal.BigInt().Int64()
	return NewPreImage(h, size), nil
}

func (hp HashPreImage) String() string {
	return fmt.Sprintf("HashPreImage(%v, %v)", hp.hashImage, hp.size)
}

func (hp HashPreImage) GetInnerHash() common.Hash {
	return hp.hashImage
}

func (hp HashPreImage) TypeCode() uint8 {
	return TypeCodeHashPreImage
}

func (hp HashPreImage) Clone() Value {
	return HashPreImage{hp.hashImage, hp.size}
}

func (hp HashPreImage) Equal(val Value) bool {
	o, ok := val.(HashPreImage)
	if !ok {
		return false
	}
	return hp.hashImage == o.hashImage && hp.size == o.size
}

func (hp HashPreImage) Size() int64 {
	return hp.size
}
