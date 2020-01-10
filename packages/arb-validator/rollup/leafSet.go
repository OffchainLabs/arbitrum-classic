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

package rollup

import (
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type LeafSet struct {
	idx map[[32]byte]*Node
}

func NewLeafSet() *LeafSet {
	return &LeafSet{
		make(map[[32]byte]*Node),
	}
}

func (ll *LeafSet) IsLeaf(node *Node) bool {
	_, ok := ll.idx[node.hash]
	return ok
}

func (ll *LeafSet) NumLeaves() int {
	return len(ll.idx)
}

func (ll *LeafSet) Add(node *Node) {
	log.Println("Added leaf", node.linkType, hexutil.Encode(node.hash[:]))
	if ll.IsLeaf(node) {
		log.Fatal("tried to insert leaf twice")
	}
	ll.idx[node.hash] = node
}

func (ll *LeafSet) Delete(node *Node) {
	log.Println("Removed leaf", node.linkType, hexutil.Encode(node.hash[:]))
	delete(ll.idx, node.hash)
}

func (ll *LeafSet) forall(f func(*Node)) {
	for _, v := range ll.idx {
		f(v)
	}
}

func (ll *LeafSet) Equals(ll2 *LeafSet) bool {
	if len(ll.idx) != len(ll2.idx) {
		return false
	}
	for h, n := range ll.idx {
		if ll2.idx[h] == nil || !n.Equals(ll2.idx[h]) {
			return false
		}
	}
	return true
}
