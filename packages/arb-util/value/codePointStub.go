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
	"encoding/binary"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"io"
)

type CodePointStub struct {
	PC   uint64
	hash common.Hash
}

func NewCodePointStubFromReader(rd io.Reader) (CodePointStub, error) {
	var insnNum uint64
	if err := binary.Read(rd, binary.BigEndian, &insnNum); err != nil {
		return CodePointStub{}, err
	}
	var hash common.Hash
	if _, err := rd.Read(hash[:]); err != nil {
		return CodePointStub{}, err
	}
	return CodePointStub{
		PC:   insnNum,
		hash: hash,
	}, nil
}

func (cp CodePointStub) String() string {
	return fmt.Sprintf("CodePointStub(%v, %v)", cp.PC, cp.hash)
}

func (cp CodePointStub) Marshal(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, &cp.PC); err != nil {
		return err
	}
	_, err := w.Write(cp.hash[:])
	return err
}

func (cp CodePointStub) TypeCode() uint8 {
	return TypeCodeCodePointStub
}

func (cp CodePointStub) Hash() common.Hash {
	return cp.hash
}

func (cp CodePointStub) Size() int64 {
	return 1
}

func (cp CodePointStub) Clone() Value {
	return cp
}

func (cp CodePointStub) Equal(val Value) bool {
	if cp2, ok := val.(CodePointStub); ok {
		return cp.hash == cp2.hash
	}
	return false
}
