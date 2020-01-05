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
	"errors"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-util) -I. -I .. --go_out=paths=source_relative:. *.proto"

type ChainObserver struct {
	*StakedNodeGraph
	rollupAddr       common.Address
	pendingInbox     *structures.PendingInbox
	listenForAddress common.Address
	listener         ChainEventListener
}

func NewChain(
	_rollupAddr common.Address,
	_machine machine.Machine,
	_vmParams structures.ChainParams,
	_listenForAddress common.Address,
	_listener ChainEventListener,
) *ChainObserver {
	ret := &ChainObserver{
		StakedNodeGraph:  NewStakedNodeGraph(_machine, _vmParams),
		rollupAddr:       _rollupAddr,
		pendingInbox:     structures.NewPendingInbox(),
		listenForAddress: _listenForAddress,
		listener:         _listener,
	}
	return ret
}

func (chain *ChainObserver) MarshalToBuf() *ChainObserverBuf {
	return &ChainObserverBuf{
		StakedNodeGraph: chain.StakedNodeGraph.MarshalToBuf(),
		ContractAddress: chain.rollupAddr.Bytes(),
		PendingInbox:    chain.pendingInbox.MarshalToBuf(),
	}
}

func (m *ChainObserverBuf) Unmarshal(_listenForAddress common.Address, _listener ChainEventListener) *ChainObserver {
	chain := &ChainObserver{
		StakedNodeGraph:  m.StakedNodeGraph.Unmarshal(),
		rollupAddr:       common.BytesToAddress(m.ContractAddress),
		pendingInbox:     &structures.PendingInbox{m.PendingInbox.Unmarshal()},
		listenForAddress: _listenForAddress,
		listener:         _listener,
	}
	return chain
}

func (chain *ChainObserver) notifyAssert(
	prevLeafHash [32]byte,
	params *structures.AssertionParams,
	claim *structures.AssertionClaim,
	maxPendingTop [32]byte,
	currentTime *protocol.TimeBlocks,
) error {
	topPendingCount, ok := chain.pendingInbox.GetHeight(maxPendingTop)
	if !ok {
		return errors.New("Couldn't find top message in inbox")
	}
	disputableNode := structures.NewDisputableNode(
		params,
		claim,
		maxPendingTop,
		topPendingCount,
	)
	chain.CreateNodesOnAssert(chain.nodeFromHash[prevLeafHash], disputableNode, nil, currentTime)
	return nil
}

func (chain *ChainObserver) notifyNewBlockNumber(blockNum *big.Int) {
	chain.Lock()
	defer chain.Unlock()
	//TODO: checkpoint, and take other appropriate actions for new block
}

func (co *ChainObserver) Equals(co2 *ChainObserver) bool {
	return co.StakedNodeGraph.Equals(co2.StakedNodeGraph) &&
		bytes.Compare(co.rollupAddr[:], co2.rollupAddr[:]) == 0 &&
		co.pendingInbox.Equals(co2.pendingInbox) &&
		bytes.Compare(co.listenForAddress[:], co2.listenForAddress[:]) == 0
}
