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

package protocol

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"math/big"
)

type MerkleTree struct {
	layers [][][32]byte
}

func NewMerkleTree(elements [][32]byte) *MerkleTree {
	layers := make([][][32]byte, 0)
	layers = append(layers, elements)
	for len(layers[len(layers)-1]) > 1 {
		elements := layers[len(layers)-1]
		var nextLayer [][32]byte
		for i := 0; i < len(elements); i++ {
			if i%2 == 1 {
				continue
			}
			if i+1 >= len(elements) {
				nextLayer = append(nextLayer, elements[i])
			} else {
				data := hashing.SoliditySHA3(
					hashing.Bytes32(elements[i]),
					hashing.Bytes32(elements[i+1]),
				)
				nextLayer = append(nextLayer, data)
			}
		}
		layers = append(layers, nextLayer)
	}
	return &MerkleTree{layers}
}

func (m *MerkleTree) GetRoot() common.Hash {
	return m.layers[len(m.layers)-1][0]
}

func (m *MerkleTree) GetNode(index int) common.Hash {
	return m.layers[0][index]
}

func (m *MerkleTree) GetProof(index int) ([][32]byte, *big.Int) {
	if index == 0 && len(m.layers) == 1 {
		return nil, big.NewInt(0)
	}
	proof := make([][32]byte, 0)
	var path []bool
	for _, layer := range m.layers {
		var pairIndex int
		if index%2 == 0 {
			pairIndex = index + 1
		} else {
			pairIndex = index - 1
		}
		if pairIndex < len(layer) {
			path = append(path, index%2 == 0)
			proof = append(proof, layer[pairIndex])
		}
		index /= 2
	}

	// reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return proof, PathSliceToInt(path)
}

func PathSliceToInt(path []bool) *big.Int {
	route := big.NewInt(0)
	for _, entry := range path {
		route = route.Mul(route, big.NewInt(2))
		if entry {
			route = route.Add(route, big.NewInt(1))
		}
	}
	return route
}
