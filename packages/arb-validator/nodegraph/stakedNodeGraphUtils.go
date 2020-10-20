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
	"bytes"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type SortableAddressList []common.Address

func (sa SortableAddressList) Len() int {
	return len(sa)
}

func (sa SortableAddressList) Less(i, j int) bool {
	return bytes.Compare(sa[i][:], sa[j][:]) < 0
}

func (sa SortableAddressList) Swap(i, j int) {
	sa[i], sa[j] = sa[j], sa[i]
}

func newPruneParams(
	leaf *structures.Node,
	leafAncestor *structures.Node,
	latestConfirmed *structures.Node,
) valprotocol.PruneParams {
	return valprotocol.PruneParams{
		LeafHash:     leaf.Hash(),
		AncestorHash: leafAncestor.Prev().Hash(),
		LeafProof:    structures.GeneratePathProof(leafAncestor.Prev(), leaf),
		AncProof:     structures.GeneratePathProof(leafAncestor.Prev(), latestConfirmed),
	}
}

func confirmNodeOpp(currentNode *structures.Node) valprotocol.ConfirmNodeOpportunity {
	coreOpp := &valprotocol.ConfirmNodeOpportunityCore{
		Branch:           currentNode.LinkType(),
		DeadlineTicks:    currentNode.Deadline(),
		PrevVMProtoState: currentNode.Prev().VMProtoData(),
		VMProtoState:     currentNode.VMProtoData(),
	}

	if currentNode.LinkType() == valprotocol.ValidChildType {
		// We need to know the contents of the actual assertion to confirm it
		// We've only seen the hash accumulator of the messages before whereas this requires the full values
		assertion := currentNode.Assertion()
		if assertion == nil {
			// Other thread hasn't filled assertion yet
			return nil
		}

		return valprotocol.ConfirmValidOpportunity{
			ConfirmNodeOpportunityCore: coreOpp,
			MessagesData:               assertion.OutMsgsData,
			MessageCount:               assertion.OutMsgsCount,
			LogsAcc:                    currentNode.Disputable().Assertion.LastLogHash,
		}
	}

	return valprotocol.ConfirmInvalidOpportunity{
		ConfirmNodeOpportunityCore: coreOpp,
		ChallengeNodeData:          currentNode.NodeDataHash(),
	}
}
