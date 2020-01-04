/*
* Copyright 2019-2020, Offchain Labs, Inc.
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

package utils

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

func MarshalHash(h [32]byte) *value.HashBuf {
	return &value.HashBuf{
		Value: append([]byte{}, h[:]...),
	}
}

func MarshalSliceOfHashes(hs [][32]byte) []*value.HashBuf {
	ret := make([]*value.HashBuf, 0, len(hs))
	for _, h := range hs {
		ret = append(ret, MarshalHash(h))
	}
	return ret
}

func UnmarshalHash(hb *value.HashBuf) [32]byte {
	var ret [32]byte
	copy(ret[:], hb.Value)
	return ret
}

func MarshalBigInt(bi *big.Int) *value.BigIntegerBuf {
	return &value.BigIntegerBuf{
		Value: bi.Bytes(),
	}
}

func UnmarshalBigInt(buf *value.BigIntegerBuf) *big.Int {
	return new(big.Int).SetBytes(buf.Value)
}

func AddressesEqual(a1, a2 common.Address) bool {
	return bytes.Compare(a1[:], a2[:]) == 0
}
