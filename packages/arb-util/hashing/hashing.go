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

package hashing

import (
	"math/big"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func SoliditySHA3(data ...interface{}) common.Hash {
	var ret common.Hash
	copy(ret[:], solsha3.SoliditySHA3(data...))
	return ret
}

func SoliditySHA3WithPrefix(data []byte) common.Hash {
	var ret common.Hash
	copy(ret[:], solsha3.SoliditySHA3WithPrefix(data))
	return ret
}

func Bytes32(input common.Hash) []byte {
	return solsha3.Bytes32(input.Bytes())
}

func Address(input common.Address) []byte {
	return solsha3.Address(input.ToEthAddress())
}

func Bool(input bool) []byte {
	return solsha3.Bool(input)
}

func TimeTicks(input common.TimeTicks) []byte {
	return solsha3.Uint256(input.Val)
}

func TimeBlocks(input *common.TimeBlocks) []byte {
	return solsha3.Uint128(input.AsInt())
}

func Uint256(input *big.Int) []byte {
	return solsha3.Uint256(input)
}

func Uint64(input uint64) []byte {
	return solsha3.Uint64(input)
}

func Uint32(input uint32) []byte {
	return solsha3.Uint32(input)
}

func Uint8(input uint8) []byte {
	return solsha3.Uint8(input)
}
