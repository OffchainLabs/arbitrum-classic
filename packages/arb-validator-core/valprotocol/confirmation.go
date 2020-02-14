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

package valprotocol

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type ConfirmValidOpportunity struct {
	DeadlineTicks    common.TimeTicks
	Messages         []value.Value
	LogsAcc          common.Hash
	VMProtoStateHash common.Hash
}

func (opp ConfirmValidOpportunity) Clone() ConfirmNodeOpportunity {
	messages := append([]value.Value{}, opp.Messages...)
	return ConfirmValidOpportunity{
		DeadlineTicks:    opp.DeadlineTicks.Clone(),
		Messages:         messages,
		LogsAcc:          opp.LogsAcc,
		VMProtoStateHash: opp.VMProtoStateHash,
	}
}

func (opp ConfirmValidOpportunity) BranchType() ChildType {
	return ValidChildType
}

func (opp ConfirmValidOpportunity) Deadline() common.TimeTicks {
	return opp.DeadlineTicks
}

func (opp ConfirmValidOpportunity) StateHash() common.Hash {
	return opp.VMProtoStateHash
}

func (opp ConfirmValidOpportunity) ProofSize() int {
	totalSize := 3
	for _, val := range opp.Messages {
		buf := value.MarshalValueToBytes(val)
		totalSize += len(buf)
	}
	return totalSize
}

type ConfirmInvalidOpportunity struct {
	DeadlineTicks     common.TimeTicks
	ChallengeNodeData common.Hash
	Branch            ChildType
	VMProtoStateHash  common.Hash
}

func (opp ConfirmInvalidOpportunity) Clone() ConfirmNodeOpportunity {
	return ConfirmInvalidOpportunity{
		opp.DeadlineTicks.Clone(),
		opp.ChallengeNodeData,
		opp.Branch,
		opp.VMProtoStateHash,
	}
}

func (opp ConfirmInvalidOpportunity) BranchType() ChildType {
	return opp.Branch
}

func (opp ConfirmInvalidOpportunity) Deadline() common.TimeTicks {
	return opp.DeadlineTicks
}

func (opp ConfirmInvalidOpportunity) StateHash() common.Hash {
	return opp.VMProtoStateHash
}

func (opp ConfirmInvalidOpportunity) ProofSize() int {
	return 3
}

type ConfirmNodeOpportunity interface {
	Clone() ConfirmNodeOpportunity
	BranchType() ChildType
	Deadline() common.TimeTicks
	StateHash() common.Hash
	ProofSize() int
}

type ConfirmOpportunity struct {
	Nodes                  []ConfirmNodeOpportunity
	CurrentLatestConfirmed common.Hash
	StakerAddresses        []common.Address
	StakerProofs           [][]common.Hash
}

func (c *ConfirmOpportunity) Clone() *ConfirmOpportunity {
	nodes := make([]ConfirmNodeOpportunity, 0, len(c.Nodes))
	for _, node := range c.Nodes {
		nodes = append(nodes, node.Clone())
	}
	stakerAddresses := append([]common.Address{}, c.StakerAddresses...)
	stakerProofs := make([][]common.Hash, 0, len(c.StakerProofs))
	for _, proof := range c.StakerProofs {
		stakerProofs = append(stakerProofs, append([]common.Hash{}, proof...))
	}

	return &ConfirmOpportunity{
		Nodes:                  nodes,
		CurrentLatestConfirmed: c.CurrentLatestConfirmed,
		StakerAddresses:        stakerAddresses,
		StakerProofs:           stakerProofs,
	}
}
