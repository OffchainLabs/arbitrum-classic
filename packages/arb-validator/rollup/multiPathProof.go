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

package rollup

import (
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func MakeMultipathProof(from *Node, to []*Node) (
	[]uint64, // startingPoints
	[]common.Hash, // proofs
	[]uint64, // proofLengths
	[]uint64, // permutation
	error,
) {
	mpw, err := makeMultiProofWorkspace(from, to)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return mpw.startingPoints, mpw.proofs, mpw.proofLengths, mpw.permutation, nil
}

type multiProofWorkspace struct {
	fromNodeHash   common.Hash
	nodes          map[common.Hash]*multiProofNode
	startingPoints []uint64
	proofs         []common.Hash
	proofLengths   []uint64
	permutation    []uint64
}

type multiProofNode struct {
	refCount   uint
	prev       *multiProofNode
	node       *Node
	workspace  *multiProofWorkspace
	proofIndex int64
}

func newMultiProofWorkspace(fromNode *Node) *multiProofWorkspace {
	ret := &multiProofWorkspace{
		fromNode.hash,
		make(map[common.Hash]*multiProofNode),
		[]uint64{},
		[]common.Hash{},
		[]uint64{},
		[]uint64{},
	}
	n := &multiProofNode{
		1, // artificially inflate this, so backward search loop will stop here
		nil,
		fromNode,
		ret,
		0,
	}
	ret.nodes[fromNode.hash] = n
	ret.startingPoints = []uint64{0}
	return ret
}

func (mpw *multiProofWorkspace) addNodeWithPath(node *Node) (*multiProofNode, error) {
	if node.depth < mpw.nodes[mpw.fromNodeHash].node.depth {
		return nil, errors.New("no path proof exists")
	}
	if node.hash.Equals(mpw.fromNodeHash) {
		return mpw.nodes[mpw.fromNodeHash], nil
	}
	var prevNode *multiProofNode
	if mpw.nodes[node.hash] != nil {
		prevNode = mpw.nodes[node.hash]
	} else {
		var err error
		prevNode, err = mpw.addNodeWithPath(node.prev)
		if err != nil {
			return nil, err
		}
	}
	prevNode.refCount++
	ret := &multiProofNode{
		0,
		prevNode,
		node,
		mpw,
		-1,
	}
	mpw.nodes[node.hash] = ret
	return ret, nil
}

func makeMultiProofWorkspace(from *Node, to []*Node) (*multiProofWorkspace, error) {
	// build workspace with tree of nodes
	mpw := newMultiProofWorkspace(from)
	for _, toNode := range to {
		if _, err := mpw.addNodeWithPath(toNode); err != nil {
			return nil, err
		}
	}

	for _, toNode := range to {
		toMnode := mpw.nodes[toNode.hash]
		idx := mpw.addProofFor(toMnode)
		mpw.permutation = append(mpw.permutation, uint64(idx))
	}

	return mpw, nil
}

func (mpw *multiProofWorkspace) addProofFor(mnode *multiProofNode) int64 {
	if mnode.proofIndex >= 0 {
		return mnode.proofIndex
	}
	pred := mnode.prev
	for pred.refCount == 1 || pred.refCount <= mnode.refCount {
		pred = pred.prev
	}
	predIdx := mpw.addProofFor(pred)
	idx := int64(1 + len(mpw.proofLengths))
	mpw.startingPoints = append(mpw.startingPoints, uint64(predIdx))
	proof := GeneratePathProof(pred.node, mnode.node)
	mpw.proofLengths = append(mpw.proofLengths, uint64(len(proof)))
	mpw.proofs = append(mpw.proofs, proof...)
	return idx
}
