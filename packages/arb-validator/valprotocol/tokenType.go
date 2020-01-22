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

package valprotocol

import (
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type TokenType [21]byte

func TokenTypeFromIntValue(val value.IntValue) TokenType {
	var tokType TokenType
	tokBytes := val.ToBytes()
	copy(tokType[:], tokBytes[11:])
	return tokType
}

func (t TokenType) ToIntValue() value.IntValue {
	var bigtok [32]byte
	copy(bigtok[11:], t[:])
	return value.NewIntValue(new(big.Int).SetBytes(bigtok[:]))
}

func (t TokenType) IsToken() bool {
	return t[20] == 0
}

func NewTokenTypeBuf(tok [21]byte) *TokenTypeBuf {
	return &TokenTypeBuf{
		Value: tok[:],
	}
}

func (buf *TokenTypeBuf) Unmarshal() TokenType {
	var ret [21]byte
	copy(ret[:], buf.Value)
	return ret
}
