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

package chainlistener

import (
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"math/big"
)

type PreparedAssertion struct {
	Prev          *structures.Node
	BeforeState   *valprotocol.VMProtoData
	Params        *valprotocol.AssertionParams
	AssertionStub *valprotocol.ExecutionAssertionStub
	Assertion     *protocol.ExecutionAssertion
	Machine       machine.Machine
	ValidBlock    *common.BlockId
}

func (pa *PreparedAssertion) String() string {
	return fmt.Sprintf(
		"PreparedAssertion(%v, %v, %v, %v, %v, %v)",
		pa.Prev.Hash(),
		pa.BeforeState,
		pa.Params,
		pa.AssertionStub,
		pa.Assertion,
		pa.ValidBlock,
	)
}

func (pa *PreparedAssertion) Clone() *PreparedAssertion {
	return &PreparedAssertion{
		Prev:          pa.Prev,
		BeforeState:   pa.BeforeState.Clone(),
		Params:        pa.Params.Clone(),
		AssertionStub: pa.AssertionStub.Clone(),
		Assertion:     pa.Assertion,
		Machine:       pa.Machine,
		ValidBlock:    pa.ValidBlock.Clone(),
	}
}

func (prep *PreparedAssertion) GetAssertionParams() [8][32]byte {
	return [8][32]byte{
		prep.AssertionStub.BeforeMachineHash,
		prep.AssertionStub.AfterMachineHash,
		prep.AssertionStub.BeforeInboxHash,
		prep.AssertionStub.AfterInboxHash,
		prep.AssertionStub.LastMessageHash,
		prep.AssertionStub.LastLogHash,
		prep.Prev.PrevHash(),
		prep.Prev.NodeDataHash(),
	}
}

func (prep *PreparedAssertion) GetAssertionParams2() [5]*big.Int {
	return [5]*big.Int{
		prep.BeforeState.InboxCount,
		prep.Prev.Deadline().Val,
		prep.Params.ImportedMessageCount,
		prep.BeforeState.MessageCount,
		prep.BeforeState.LogCount,
	}
}
