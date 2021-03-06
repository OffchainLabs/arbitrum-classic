/*
 * Copyright 2021, Offchain Labs, Inc.
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

package evm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/pkg/errors"
	"math/big"
)

type MerkleNode interface {
	Hash() common.Hash
	Lowest() uint64
	Highest() uint64
	ContainsIndex(index uint64) bool
	Entries() [][]byte
}

type MerkleLeaf struct {
	Data []byte

	index uint64
}

func (m *MerkleLeaf) Hash() common.Hash {
	return hashing.SoliditySHA3(hashing.Bytes32(hashing.SoliditySHA3(m.Data)))
}

func (m *MerkleLeaf) String() string {
	return hexutil.Encode(m.Data)
}

func (m *MerkleLeaf) Lowest() uint64 {
	return m.index
}

func (m *MerkleLeaf) Highest() uint64 {
	return m.index
}

func (m *MerkleLeaf) ContainsIndex(index uint64) bool {
	return index == m.index
}

func (m *MerkleLeaf) Entries() [][]byte {
	return [][]byte{m.Data}
}

type MerkleInteriorNode struct {
	Left  MerkleNode
	Right MerkleNode

	lowest  uint64
	highest uint64
}

func NewMerkleInteriorNode(left, right MerkleNode) *MerkleInteriorNode {
	return &MerkleInteriorNode{
		Left:    left,
		Right:   right,
		lowest:  left.Lowest(),
		highest: right.Highest(),
	}
}

func (m *MerkleInteriorNode) String() string {
	return fmt.Sprintf("(%v, %v)", m.Left, m.Right)
}

func (m *MerkleInteriorNode) Hash() common.Hash {
	return hashing.SoliditySHA3(hashing.Bytes32(m.Left.Hash()), hashing.Bytes32(m.Right.Hash()))
}

func (m *MerkleInteriorNode) Lowest() uint64 {
	return m.lowest
}

func (m *MerkleInteriorNode) Highest() uint64 {
	return m.highest
}

func (m *MerkleInteriorNode) ContainsIndex(index uint64) bool {
	return index >= m.lowest && index <= m.highest
}

func (m *MerkleInteriorNode) Entries() [][]byte {
	return append(m.Left.Entries(), m.Right.Entries()...)
}

type MerkleRootResult struct {
	BatchNumber *big.Int
	NumInBatch  *big.Int
	Tree        MerkleNode
}

func NewMerkleRootLogResultFromValue(tup *value.TupleValue) (*MerkleRootResult, error) {
	if tup.Len() != 4 {
		return nil, errors.New("expected merkle root info tuple of length 4")
	}

	resultKindVal, _ := tup.GetByInt64(0)
	batchNumberVal, _ := tup.GetByInt64(1)
	numInBatchVal, _ := tup.GetByInt64(2)
	treeVal, _ := tup.GetByInt64(3)

	resultKindInt, ok := resultKindVal.(value.IntValue)
	if !ok {
		return nil, errors.New("resultKind must be an int")
	}
	if resultKindInt.BigInt().Uint64() != 3 {
		return nil, errors.New("incorrect result kind for merkle root log result")
	}

	batchNumber, ok := batchNumberVal.(value.IntValue)
	if !ok {
		return nil, errors.New("batchNumber must be an int")
	}
	numInBatch, ok := numInBatchVal.(value.IntValue)
	if !ok {
		return nil, errors.New("numInBatch must be an int")
	}
	tree, err := newMerkleTreeFromValue(treeVal, 0)
	if err != nil {
		return nil, err
	}
	return &MerkleRootResult{
		BatchNumber: batchNumber.BigInt(),
		NumInBatch:  numInBatch.BigInt(),
		Tree:        tree,
	}, nil
}

func newMerkleTreeFromValue(val value.Value, minIndex uint64) (MerkleNode, error) {
	treeTup, ok := val.(*value.TupleValue)
	if !ok {
		return nil, errors.New("tree must be a 2-tuple")
	}
	if treeTup.Len() == 2 {
		node1Val, _ := treeTup.GetByInt64(0)
		node2Val, _ := treeTup.GetByInt64(1)
		node1, err := newMerkleTreeFromValue(node1Val, minIndex)
		if err != nil {
			return nil, err
		}
		node2, err := newMerkleTreeFromValue(node2Val, node1.Highest()+1)
		if err != nil {
			return nil, err
		}
		return NewMerkleInteriorNode(node1, node2), nil
	} else if treeTup.Len() == 3 {
		dataSizeVal, _ := treeTup.GetByInt64(0)
		dataContentsVal, _ := treeTup.GetByInt64(1)
		dataSizeInt, ok := dataSizeVal.(value.IntValue)
		if !ok {
			return nil, errors.New("dataSize must be an int")
		}
		dataContentsBuf, ok := dataContentsVal.(*value.Buffer)
		if !ok {
			return nil, errors.New("dataContents must be a buffer")
		}
		data, err := inbox.BufAndLengthToBytes(dataSizeInt.BigInt(), dataContentsBuf)
		if err != nil {
			return nil, err
		}
		return &MerkleLeaf{
			Data:  data,
			index: minIndex,
		}, nil
	} else {
		return nil, errors.New("tree node must be a 2 or 3 tuple")
	}
}

type MerkleRootProof struct {
	Nodes []common.Hash
	Path  []bool
	Data  []byte
}

func (m *MerkleRootResult) GenerateProof(index uint64) (*MerkleRootProof, error) {
	nodes := make([]common.Hash, 0)
	path := make([]bool, 0)
	nd := m.Tree
	for {
		switch node := nd.(type) {
		case *MerkleInteriorNode:
			if node.Left.ContainsIndex(index) {
				nodes = append(nodes, node.Right.Hash())
				path = append(path, true)
				nd = node.Left
			} else if node.Right.ContainsIndex(index) {
				nodes = append(nodes, node.Left.Hash())
				nd = node.Right
				path = append(path, false)
			} else {
				return nil, errors.New("invalid merkle tree")
			}
		case *MerkleLeaf:
			if index != node.index {
				return nil, errors.New("invalid merkle tree")
			}
			for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
				nodes[i], nodes[j] = nodes[j], nodes[i]
			}
			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}
			return &MerkleRootProof{
				Nodes: nodes,
				Path:  path,
				Data:  node.Data,
			}, nil
		}
	}
}
