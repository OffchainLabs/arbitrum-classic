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

package common

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Hash [32]byte

func NewHashFromEth(a ethcommon.Hash) Hash {
	return Hash(a)
}

func (h Hash) String() string {
	return hexutil.Encode(h[:])
}

func (h Hash) Bytes() []byte {
	return h[:]
}

func (h Hash) Equals(h2 Hash) bool {
	return h == h2
}

func (h Hash) ToEthHash() ethcommon.Hash {
	return ethcommon.Hash(h)
}

func (h Hash) MarshalToBuf() *HashBuf {
	return &HashBuf{
		Value: append([]byte{}, h[:]...),
	}
}

func MarshalSliceOfHashes(hs []Hash) []*HashBuf {
	ret := make([]*HashBuf, 0, len(hs))
	for _, h := range hs {
		ret = append(ret, h.MarshalToBuf())
	}
	return ret
}

func (hb *HashBuf) Unmarshal() Hash {
	var ret Hash
	copy(ret[:], hb.Value)
	return ret
}
