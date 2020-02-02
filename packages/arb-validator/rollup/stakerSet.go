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
	"bytes"
	"log"
	"strconv"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type Staker struct {
	address      common.Address
	location     *Node
	creationTime common.TimeTicks
	challenge    common.Address
}

type StakerSet struct {
	idx map[common.Address]*Staker
}

func NewStakerSet() *StakerSet {
	return &StakerSet{make(map[common.Address]*Staker)}
}

func (ss *StakerSet) Add(newStaker *Staker) {
	newStaker.location.numStakers++
	if _, ok := ss.idx[newStaker.address]; ok {
		log.Fatal("tried to insert staker twice")
	}
	ss.idx[newStaker.address] = newStaker
}

func (ss *StakerSet) Delete(staker *Staker) {
	delete(ss.idx, staker.address)
}

func (ss *StakerSet) Get(addr common.Address) *Staker {
	return ss.idx[addr]
}

func (ss *StakerSet) forall(f func(*Staker)) {
	for _, v := range ss.idx {
		f(v)
	}
}

func (s *Staker) MarshalToBuf() *StakerBuf {
	emptyAddress := common.Address{}
	if s.challenge == emptyAddress {
		return &StakerBuf{
			Address:       s.address.MarshallToBuf(),
			Location:      s.location.hash.MarshalToBuf(),
			CreationTime:  s.creationTime.MarshalToBuf(),
			ChallengeAddr: nil,
		}
	} else {
		return &StakerBuf{
			Address:       s.address.MarshallToBuf(),
			Location:      s.location.hash.MarshalToBuf(),
			CreationTime:  s.creationTime.MarshalToBuf(),
			ChallengeAddr: s.challenge.MarshallToBuf(),
		}
	}
}

func (m *StakerBuf) Unmarshal(chain *StakedNodeGraph) *Staker {
	// chain.nodeFromHash and chain.challenges must have already been unmarshaled
	locArr := m.Location.Unmarshal()
	if m.ChallengeAddr != nil {
		return &Staker{
			address:      m.Address.Unmarshal(),
			location:     chain.nodeFromHash[locArr],
			creationTime: m.CreationTime.Unmarshal(),
			challenge:    m.ChallengeAddr.Unmarshal(),
		}
	} else {
		return &Staker{
			address:      m.Address.Unmarshal(),
			location:     chain.nodeFromHash[locArr],
			creationTime: m.CreationTime.Unmarshal(),
			challenge:    common.Address{},
		}
	}
}

func (ss *StakerSet) DebugString(prefix string) string {
	ret := prefix + "stakers:\n"
	subPrefix := prefix + "  "
	ss.forall(func(s *Staker) {
		ret += s.DebugString(subPrefix)
	})
	return ret
}

func (s *Staker) DebugString(prefix string) string {
	ret := prefix + "depth:" + strconv.FormatUint(s.location.depth, 10) + " addr:" + s.address.ShortString() + " created:" + s.creationTime.String() + " loc:" + s.location.hash.ShortString()
	if !s.challenge.IsZero() {
		ret += " chal:" + s.challenge.ShortString()
	}
	return ret + "\n"
}

func (ss *StakerSet) Equals(ss2 *StakerSet) bool {
	if len(ss.idx) != len(ss2.idx) {
		return false
	}
	for addr, staker := range ss.idx {
		staker2 := ss2.idx[addr]
		if staker2 == nil {
			return false
		}
		if !staker.Equals(staker2) {
			return false
		}
	}
	return true
}

func (s *Staker) Equals(s2 *Staker) bool {
	if !bytes.Equal(s.address[:], s2.address[:]) {
		return false
	}
	if s.location.hash != s2.location.hash {
		return false
	}
	if !s.creationTime.Equals(s2.creationTime) {
		return false
	}
	return s.challenge == s2.challenge
}
