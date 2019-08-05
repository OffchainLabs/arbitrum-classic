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

package ethvalidator

import solsha3 "github.com/miguelmota/go-solidity-sha3"

type MerkleTree struct {
	layers [][][32]byte
}

func NewMerkleTree(elements [][32]byte) *MerkleTree {
	layers := make([][][32]byte, 0)
	layers = append(layers, elements)
	for len(layers[len(layers)-1]) > 1 {
		prevLayerSize := len(layers[len(layers)-1])
		nextLayer := make([][32]byte, 0, (prevLayerSize+1)/2)
		for i := 0; i < (prevLayerSize+1)/2; i++ {
			if 2*i+1 < prevLayerSize {
				combined := solsha3.SoliditySHA3(
					solsha3.Bytes32(layers[len(layers)-1][2*i]),
					solsha3.Bytes32(layers[len(layers)-1][2*i+1]),
				)
				var data [32]byte
				copy(data[:], combined)
				nextLayer = append(nextLayer, data)
			} else {
				nextLayer = append(nextLayer, layers[len(layers)-1][2*i])
			}
		}
		layers = append(layers, nextLayer)
	}
	return &MerkleTree{layers}
}

func (m *MerkleTree) GetRoot() [32]byte {
	return m.layers[len(m.layers)-1][0]
}

func (m *MerkleTree) GetNode(index int) [32]byte {
	return m.layers[0][index]
}

func (m *MerkleTree) GetProof(index int) [][32]byte {
	proof := make([][32]byte, 0)
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
		proofFlat = append(proofFlat, item[:]...)
	}
	return proofFlat
}
