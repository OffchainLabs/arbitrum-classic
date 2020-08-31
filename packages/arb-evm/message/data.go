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

package message

import (
	ethmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func extractUInt256(data []byte) (*big.Int, []byte) {
	val := new(big.Int).SetBytes(data[:32])
	data = data[32:]
	return val, data
}

func extractAddress(data []byte) (common.Address, []byte) {
	data = data[12:] // Skip first 12 bytes of 32 byte address data
	var addr common.Address
	copy(addr[:], data[:])
	data = data[20:]
	return addr, data
}

func addressData(addr common.Address) []byte {
	ret := make([]byte, 0, 32)
	ret = append(ret, make([]byte, 12)...)
	ret = append(ret, addr[:]...)
	return ret
}

func marshaledBytesHash(data []byte) common.Hash {
	var ret common.Hash
	copy(ret[:], ethmath.U256Bytes(big.NewInt(int64(len(data)))))
	chunks := make([]common.Hash, 0)
	for len(data) > 0 {
		var nextVal common.Hash
		copy(nextVal[:], data[:])
		chunks = append(chunks, nextVal)
		if len(data) <= 32 {
			break
		}
		data = data[32:]
	}

	for i := range chunks {
		ret = hashing.SoliditySHA3(
			hashing.Bytes32(ret),
			hashing.Bytes32(chunks[len(chunks)-1-i]),
		)
	}
	return ret
}
