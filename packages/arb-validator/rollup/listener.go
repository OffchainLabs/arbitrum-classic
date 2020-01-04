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
	"github.com/ethereum/go-ethereum/common"
	"log"
)

type ChainEventListener interface {
	Notify(interface{})
}

type ChanCEListener struct {
	chain *ChainObserver
	ch    chan interface{}
}

func NewChanCEListener(chain *ChainObserver, runLoop func(*ChanCEListener)) {
	ret := &ChanCEListener{chain, make(chan interface{}, 1024)}
	go runLoop(ret)
}

func (lis *ChanCEListener) Notify(ev interface{}) {
	select {
	case lis.ch <- ev:
		// do nothing
	default:
		log.Fatal("ChanCEListener: ran out of buffer space")
	}
}

type StakeCreatedChainEvent struct {
	staker       common.Address
	nodeHash     [32]byte
	creationTime RollupTime
}

type StakeMovedChainEvent struct {
	staker     common.Address
	toNodeHash [32]byte
}

type StakeRefundedChainEvent struct {
	staker common.Address
}

type ChallengeStartedChainEvent struct {
	challengeContract common.Address
	asserter          common.Address
	challenger        common.Address
	kind              ChallengeType
}

type ChallengeCompletedChainEvent struct {
	challengeContract common.Address
	winner            common.Address
	loser             common.Address
}

func sinkRunLoop(lis *ChanCEListener) {
	for {
		_, ok := <-lis.ch
		if !ok {
			return
		}
	}
}

func templateRunLoop(lis *ChanCEListener) {
	for {
		inEv, ok := <-lis.ch
		if !ok {
			return
		}
		switch ev := inEv.(type) {
		case StakeCreatedChainEvent:
		case StakeMovedChainEvent:
		case StakeRefundedChainEvent:
		case ChallengeStartedChainEvent:
		case ChallengeCompletedChainEvent:
		default:
			_ = ev //suppress compiler warning
			log.Fatal("unrecognized event type in rollup chain listener")
		}
	}
}
