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

package common

import (
	"errors"
	"math/big"
)

//go:generate protoc -I.. -I. --go_out=paths=source_relative:. common.proto

func NewBigIntBuf(buf *big.Int) *BigIntegerBuf {
	return &BigIntegerBuf{
		Value: buf.Bytes(),
	}
}
func NewBigIntBufFromUint64(val uint64) *BigIntegerBuf {
	return NewBigIntBuf(big.NewInt(int64(val)))
}

func NewBigIntFromBuf(buf *BigIntegerBuf) *big.Int {
	return new(big.Int).SetBytes(buf.Value)
}

func Uint64FromBuf(buf *BigIntegerBuf) (uint64, error) {
	bi := NewBigIntFromBuf(buf)
	if !bi.IsUint64() {
		return 0, errors.New("block number does not fit in uint64")
	}
	return bi.Uint64(), nil
}

func NewHashBuf(h Hash) *HashBuf {
	return &HashBuf{
		Value: h[:],
	}
}

func NewHashFromBuf(buf *HashBuf) Hash {
	var ret [32]byte
	copy(ret[:], buf.Value)
	return ret
}
