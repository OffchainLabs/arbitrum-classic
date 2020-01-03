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

package rollup

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"math/big"
)

func marshalHash(h [32]byte) *value.HashBuf {
	return &value.HashBuf{
		Value: append([]byte{}, h[:]...),
	}
}

func marshalSliceOfHashes(hs [][32]byte) []*value.HashBuf {
	ret := make([]*value.HashBuf, 0, len(hs))
	for _, h := range hs {
		ret = append(ret, marshalHash(h))
	}
	return ret
}

func unmarshalHash(hb *value.HashBuf) [32]byte {
	var ret [32]byte
	copy(ret[:], hb.Value)
	return ret
}

func marshalBigInt(bi *big.Int) *value.BigIntegerBuf {
	return &value.BigIntegerBuf{
		Value: bi.Bytes(),
	}
}

func unmarshalBigInt(buf *value.BigIntegerBuf) *big.Int {
	return new(big.Int).SetBytes(buf.Value)
}
