/*
* Copyright 020, Offchain Labs, Inc.
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

package nodegraph

import (
	"bytes"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"github.com/rs/zerolog"
	"strconv"
)

type Staker struct {
	address      common.Address
	location     *structures.Node
	creationTime common.TimeTicks
	challenge    common.Address
}

func (staker *Staker) MarshalZerologObject(e *zerolog.Event) {
	e.Hex("address", staker.address.Bytes()).
		Hex("location", staker.location.Hash().Bytes()).
		Uint64("depth", staker.location.Depth()).
		Str("created", staker.creationTime.Val.String()).
		Hex("challenge", staker.challenge.Bytes())
}

func (staker *Staker) Challenge() common.Address {
	return staker.challenge
}

func (staker *Staker) Location() *structures.Node {
	return staker.location
}

func (staker *Staker) CreationTime() common.TimeTicks {
	return staker.creationTime
}

func (staker *Staker) Address() common.Address {
	return staker.address
}

func (staker *Staker) MarshalToBuf() *StakerBuf {
	emptyAddress := common.Address{}
	if staker.challenge == emptyAddress {
		return &StakerBuf{
			Address:       staker.address.MarshallToBuf(),
			Location:      staker.location.Hash().MarshalToBuf(),
			CreationTime:  staker.creationTime.MarshalToBuf(),
			ChallengeAddr: nil,
		}
	} else {
		return &StakerBuf{
			Address:       staker.address.MarshallToBuf(),
			Location:      staker.location.Hash().MarshalToBuf(),
			CreationTime:  staker.creationTime.MarshalToBuf(),
			ChallengeAddr: staker.challenge.MarshallToBuf(),
		}
	}
}

func (buf *StakerBuf) Unmarshal(chain *NodeGraph) *Staker {
	// chain.nodeFromHash must have already been unmarshaled
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

func (s *Staker) DebugString(prefix string) string {
	ret := prefix + "depth:" + strconv.FormatUint(s.location.Depth(), 10) + " addr:" + s.address.ShortString() + " created:" + s.creationTime.String() + " loc:" + s.location.Hash().ShortString()
	if !s.challenge.IsZero() {
		ret = ret + " chal:" + s.challenge.ShortString()
	}
	return ret + "\n"
}

func (s *Staker) Equals(s2 *Staker) bool {
	if !bytes.Equal(s.address[:], s2.address[:]) {
		return false
	}
	if s.location.Hash() != s2.location.Hash() {
		return false
	}
	if !s.creationTime.Equals(s2.creationTime) {
		return false
	}
	return s.challenge == s2.challenge
}
