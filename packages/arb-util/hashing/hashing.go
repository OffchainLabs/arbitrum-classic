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
	"golang.org/x/crypto/sha3"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"

	solsha3 "github.com/offchainlabs/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func SoliditySHA3(data ...interface{}) common.Hash {
	var ret common.Hash
	hash := sha3.NewLegacyKeccak256()
	for _, b := range data {
		_, err := hash.Write(b.([]byte))
		if err != nil {
			// This code should never be reached
			panic("Error writing SoliditySHA3 data")
		}
	}
	hash.Sum(ret[:0])
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

func AddressArray(input []common.Address) []byte {
	addresses := make([]ethcommon.Address, 0, len(input))
	for _, address := range input {
		addresses = append(addresses, address.ToEthAddress())
	}
	return solsha3.AddressArray(addresses)
}

func Bool(input bool) []byte {
	return solsha3.Bool(input)
}

func TimeBlocks(input *common.TimeBlocks) []byte {
	return solsha3.Uint128(input.AsInt())
}

// Uint256 converts input to its packed ABI encoding
func Uint256(input *big.Int) []byte {
	return solsha3.Uint256(new(big.Int).Set(input))
}

// Uint128 converts input to its packed ABI encoding
func Uint128(input *big.Int) []byte {
	return solsha3.Uint128(new(big.Int).Set(input))
}

// Uint256Array converts input to its packed ABI encoding
func Uint256Array(input []*big.Int) []byte {
	ints := make([]*big.Int, 0, len(input))
	for _, val := range input {
		ints = append(ints, new(big.Int).Set(val))
	}
	return solsha3.Uint256Array(ints)
}

// Uint64 converts input to its packed ABI encoding
func Uint64(input uint64) []byte {
	return solsha3.Uint64(input)
}

// Uint32 converts input to its packed ABI encoding
func Uint32(input uint32) []byte {
	return solsha3.Uint32(input)
}

// Uint32Array converts input to its packed ABI encoding
func Uint32Array(input []uint32) []byte {
	return solsha3.Uint32Array(input)
}

// Uint8 converts input to its packed ABI encoding
func Uint8(input uint8) []byte {
	return solsha3.Uint8(input)
}

// Bytes32ArrayEncoded converts input to its packed ABI encoding
func Bytes32ArrayEncoded(input []common.Hash) []byte {
	var values []byte
	for _, val := range input {
		values = append(values, ethcommon.RightPadBytes(val[:], 32)...)
	}
	return values
}
