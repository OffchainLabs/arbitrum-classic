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
	"log"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type Challenge struct {
	blockId      *common.BlockId
	logIndex     uint
	asserter     common.Address
	challenger   common.Address
	contract     common.Address
	conflictNode *Node
}

type ChallengeSet struct {
	idx map[common.Address]*Challenge
}

func NewChallengeSet() *ChallengeSet {
	return &ChallengeSet{make(map[common.Address]*Challenge)}
}

func (cs *ChallengeSet) Add(newChallenge *Challenge) {
	if _, ok := cs.idx[newChallenge.contract]; ok {
		log.Fatal("tried to insert challenge twice")
	}
	cs.idx[newChallenge.contract] = newChallenge
}

func (cs *ChallengeSet) Delete(contract common.Address) {
	delete(cs.idx, challenge.contract)
}

func (cs *ChallengeSet) Get(addr common.Address) *Challenge {
	return cs.idx[addr]
}

func (cs *ChallengeSet) forall(f func(*Challenge)) {
	for _, v := range cs.idx {
		f(v)
	}
}

func (c *Challenge) MarshalToBuf() *ChallengeBuf {
	return &ChallengeBuf{
		BlockId:              c.blockId.MarshalToBuf(),
		LogIndex:             uint64(c.logIndex),
		Asserter:             c.asserter.MarshallToBuf(),
		Challenger:           c.challenger.MarshallToBuf(),
		Contract:             c.contract.MarshallToBuf(),
		ConflictNodeHash:     c.conflictNode.hash.MarshalToBuf(),
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
}

func (m *ChallengeBuf) Unmarshal(chain *NodeGraph) *Challenge {
	// chain.nodeFromHash must have already been unmarshaled
	conflictNodeHash := m.ConflictNodeHash.Unmarshal()
	return &Challenge{
		blockId:      m.BlockId.Unmarshal(),
		logIndex:     uint(m.LogIndex),
		asserter:     m.Asserter.Unmarshal(),
		challenger:   m.Challenger.Unmarshal(),
		contract:     m.Contract.Unmarshal(),
		conflictNode: chain.nodeFromHash[conflictNodeHash],
	}
}

func (cs *ChallengeSet) Equals(cs2 *ChallengeSet) bool {
	if len(cs.idx) != len(cs2.idx) {
		return false
	}
	for addr, challenge := range cs.idx {
		challenge2 := cs2.idx[addr]
		if challenge2 == nil {
			return false
		}
		if !challenge.Equals(challenge2) {
			return false
		}
	}
	return true
}

func (c *Challenge) Equals(s2 *Challenge) bool {
	return c.blockId.Equals(s2.blockId) &&
		c.logIndex == s2.logIndex &&
		c.asserter.Equals(s2.asserter) &&
		c.challenger.Equals(s2.challenger) &&
		c.contract.Equals(s2.contract) &&
		c.conflictNode.Equals(s2.conflictNode)
}
