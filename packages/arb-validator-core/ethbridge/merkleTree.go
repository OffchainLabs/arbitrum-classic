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

package ethbridge

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
)

type MerkleTree struct {
	layers [][]common.Hash
}

func NewMerkleTree(elements []common.Hash) *MerkleTree {
	layers := make([][]common.Hash, 0)
	layers = append(layers, elements)
	for len(layers[len(layers)-1]) > 1 {
		prevLayerSize := len(layers[len(layers)-1])
		nextLayer := make([]common.Hash, 0, (prevLayerSize+1)/2)
		for i := 0; i < (prevLayerSize+1)/2; i++ {
			if 2*i+1 < prevLayerSize {
				data := hashing.SoliditySHA3(
					hashing.Bytes32(layers[len(layers)-1][2*i]),
					hashing.Bytes32(layers[len(layers)-1][2*i+1]),
				)
				nextLayer = append(nextLayer, data)
			} else {
				nextLayer = append(nextLayer, layers[len(layers)-1][2*i])
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

func (m *MerkleTree) GetProof(index int) []common.Hash {
	proof := make([]common.Hash, 0)
	for _, layer := range m.layers {
		var pairIndex int
		if index%2 == 0 {
			pairIndex = index + 1
		} else {
			pairIndex = index - 1
		}
		if pairIndex < len(layer) {
			proof = append(proof, layer[pairIndex])
		}
		index /= 2
	}
	return proof
}

func (m *MerkleTree) GetProofFlat(index int) []byte {
	proofList := m.GetProof(index)
	proofFlat := make([]byte, 0, 32*len(proofList))
	for _, item := range proofList {
		proofFlat = append(proofFlat, item.Bytes()...)
	}
	return proofFlat
}
