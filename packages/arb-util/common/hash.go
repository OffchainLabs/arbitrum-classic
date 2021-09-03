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

func (h Hash) ShortString() string {
	return h.String()[2:8]
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

func (h Hash) MarshalText() ([]byte, error) {
	return []byte(h.String()), nil
}

func NewEthHashesFromHashes(a []Hash) []ethcommon.Hash {
	ret := make([]ethcommon.Hash, 0, len(a))
	for _, t := range a {
		ret = append(ret, ethcommon.BytesToHash(t[:]))
	}
	return ret
}

func HexToHash(hex string) Hash {
	return NewHashFromEth(ethcommon.HexToHash(hex))
}

func HashSliceToRaw(slice []Hash) [][32]byte {
	ret := make([][32]byte, 0, len(slice))
	for _, h := range slice {
		ret = append(ret, h)
	}
	return ret
}

func HashArrayFromEth(hashes []ethcommon.Hash) []Hash {
	ret := make([]Hash, 0, len(hashes))
	for _, a := range hashes {
		ret = append(ret, NewHashFromEth(a))
	}
	return ret
}
