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

func (sl *StakerSet) Add(newStaker *Staker) {
	newStaker.location.numStakers++
	if _, ok := sl.idx[newStaker.address]; ok {
		log.Fatal("tried to insert staker twice")
	}
	sl.idx[newStaker.address] = newStaker
}

func (sl *StakerSet) Delete(staker *Staker) {
	delete(sl.idx, staker.address)
}

func (sl *StakerSet) Get(addr common.Address) *Staker {
	return sl.idx[addr]
}

func (sl *StakerSet) forall(f func(*Staker)) {
	for _, v := range sl.idx {
		f(v)
	}
}

func (staker *Staker) MarshalToBuf() *StakerBuf {
	emptyAddress := common.Address{}
	if staker.challenge == emptyAddress {
		return &StakerBuf{
			Address:       staker.address.MarshallToBuf(),
			Location:      staker.location.hash.MarshalToBuf(),
			CreationTime:  staker.creationTime.MarshalToBuf(),
			ChallengeAddr: nil,
		}
	} else {
		return &StakerBuf{
			Address:       staker.address.MarshallToBuf(),
			Location:      staker.location.hash.MarshalToBuf(),
			CreationTime:  staker.creationTime.MarshalToBuf(),
			ChallengeAddr: staker.challenge.MarshallToBuf(),
		}
	}
}

func (buf *StakerBuf) Unmarshal(chain *StakedNodeGraph) *Staker {
	// chain.nodeFromHash and chain.challenges must have already been unmarshaled
	locArr := buf.Location.Unmarshal()
	if buf.ChallengeAddr != nil {
		return &Staker{
			address:      buf.Address.Unmarshal(),
			location:     chain.nodeFromHash[locArr],
			creationTime: buf.CreationTime.Unmarshal(),
			challenge:    buf.ChallengeAddr.Unmarshal(),
		}
	} else {
		return &Staker{
			address:      buf.Address.Unmarshal(),
			location:     chain.nodeFromHash[locArr],
			creationTime: buf.CreationTime.Unmarshal(),
			challenge:    common.Address{},
		}
	}
}

func (ss *StakerSet) DebugString(prefix string) string {
	ret := prefix + "stakers:\n"
	subPrefix := prefix + "  "
	ss.forall(func(s *Staker) {
		ret = ret + s.DebugString(subPrefix)
	})
	return ret
}

func (s *Staker) DebugString(prefix string) string {
	ret := prefix + "addr:" + s.address.ShortString() + " loc:" + s.location.hash.ShortString()
	if !s.challenge.IsZero() {
		ret = ret + " chal:" + s.challenge.ShortString()
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
	if bytes.Compare(s.address[:], s2.address[:]) != 0 {
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
