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

package nodegraph

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type ChallengeOpportunity struct {
	asserter              common.Address
	challenger            common.Address
	prevNodeHash          common.Hash
	deadlineTicks         common.TimeTicks
	asserterNodeType      valprotocol.ChildType
	challengerNodeType    valprotocol.ChildType
	asserterVMProtoHash   common.Hash
	challengerVMProtoHash common.Hash
	asserterProof         []common.Hash
	challengerProof       []common.Hash
	asserterNodeHash      common.Hash
	challengerDataHash    common.Hash
	challengerPeriodTicks common.TimeTicks
}

func (co *ChallengeOpportunity) Asserter() common.Address {
	return co.asserter
}

func (co *ChallengeOpportunity) Challenger() common.Address {
	return co.challenger
}

func (co *ChallengeOpportunity) PrevNodeHash() common.Hash {
	return co.prevNodeHash
}

func (co *ChallengeOpportunity) DeadlineTicks() common.TimeTicks {
	return co.deadlineTicks
}

func (co *ChallengeOpportunity) AsserterNodeType() valprotocol.ChildType {
	return co.asserterNodeType
}

func (co *ChallengeOpportunity) ChallengerNodeType() valprotocol.ChildType {
	return co.challengerNodeType
}

func (co *ChallengeOpportunity) AsserterVMProtoHash() common.Hash {
	return co.asserterVMProtoHash
}

func (co *ChallengeOpportunity) ChallengerVMProtoHash() common.Hash {
	return co.challengerVMProtoHash
}

func (co *ChallengeOpportunity) AsserterProof() []common.Hash {
	return co.asserterProof
}

func (co *ChallengeOpportunity) ChallengerProof() []common.Hash {
	return co.challengerProof
}

func (co *ChallengeOpportunity) AsserterNodeHash() common.Hash {
	return co.asserterNodeHash
}

func (co *ChallengeOpportunity) ChallengerDataHash() common.Hash {
	return co.challengerDataHash
}

func (co *ChallengeOpportunity) ChallengerPeriodTicks() common.TimeTicks {
	return co.challengerPeriodTicks
}
