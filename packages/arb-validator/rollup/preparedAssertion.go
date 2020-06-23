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
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"math/big"
	"time"
)

type PreparedAssertion struct {
	prev        *structures.Node
	beforeState *valprotocol.VMProtoData
	params      *valprotocol.AssertionParams
	claim       *valprotocol.AssertionClaim
	assertion   *protocol.ExecutionAssertion
	machine     machine.Machine
}

func (pa *PreparedAssertion) String() string {
	return fmt.Sprintf(
		"PreparedAssertion(%v, %v, %v, %v, %v)",
		pa.prev.Hash(),
		pa.beforeState,
		pa.params,
		pa.claim,
		pa.assertion,
	)
}

func (pa *PreparedAssertion) Clone() *PreparedAssertion {
	return &PreparedAssertion{
		prev:        pa.prev,
		beforeState: pa.beforeState.Clone(),
		params:      pa.params.Clone(),
		claim:       pa.claim.Clone(),
		assertion:   pa.assertion,
		machine:     pa.machine,
	}
}

func (pa *PreparedAssertion) PossibleFutureNode(chainParams valprotocol.ChainParams) *structures.Node {
	node := structures.NewValidNodeFromPrev(
		pa.prev,
		valprotocol.NewDisputableNode(
			pa.params,
			pa.claim,
			common.Hash{},
			big.NewInt(0),
		),
		chainParams,
		common.BlocksFromSeconds(time.Now().Unix()),
		common.Hash{},
	)
	_ = node.UpdateValidOpinion(pa.machine, pa.assertion)
	return node
}

func (prep *PreparedAssertion) getAssertionParams() [9][32]byte {
	return [9][32]byte{
		prep.beforeState.MachineHash,
		prep.beforeState.InboxTop,
		prep.prev.PrevHash(),
		prep.prev.NodeDataHash(),
		prep.claim.AfterInboxTop,
		prep.claim.ImportedMessagesSlice,
		prep.claim.AssertionStub.AfterHash,
		prep.claim.AssertionStub.LastMessageHash,
		prep.claim.AssertionStub.LastLogHash,
	}
}
