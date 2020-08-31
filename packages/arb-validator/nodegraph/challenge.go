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

package nodegraph

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type Challenge struct {
	blockId      *common.BlockId
	logIndex     uint
	asserter     common.Address
	challenger   common.Address
	contract     common.Address
	conflictNode *structures.Node
}

func NewChallenge(
	blockId *common.BlockId,
	logIndex uint,
	asserter common.Address,
	challenger common.Address,
	contract common.Address,
	conflictNode *structures.Node,
) *Challenge {
	return &Challenge{
		blockId:      blockId,
		logIndex:     logIndex,
		asserter:     asserter,
		challenger:   challenger,
		contract:     contract,
		conflictNode: conflictNode,
	}
}

func NewChallengeFromEvent(event arbbridge.ChallengeStartedEvent, challengerAncestor *structures.Node) *Challenge {
	return &Challenge{
		blockId:      event.BlockId,
		logIndex:     event.LogIndex,
		asserter:     event.Asserter,
		challenger:   event.Challenger,
		contract:     event.ChallengeContract,
		conflictNode: challengerAncestor,
	}
}

func (c *Challenge) ConflictNode() *structures.Node {
	return c.conflictNode
}

func (c *Challenge) Contract() common.Address {
	return c.contract
}

func (c *Challenge) Challenger() common.Address {
	return c.challenger
}

func (c *Challenge) Asserter() common.Address {
	return c.asserter
}

func (c *Challenge) LogIndex() uint {
	return c.logIndex
}

func (c *Challenge) BlockId() *common.BlockId {
	return c.blockId
}
