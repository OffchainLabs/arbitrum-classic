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
	"math/big"
)

type ConfirmNodeOpp struct {
	Branch           ChildType
	DeadlineTicks    common.TimeTicks
	PrevVMProtoState *VMProtoData
	VMProtoState     *VMProtoData
}

func (c *ConfirmNodeOpp) Clone() *ConfirmNodeOpp {
	return &ConfirmNodeOpp{
		Branch:           c.Branch,
		DeadlineTicks:    c.DeadlineTicks.Clone(),
		PrevVMProtoState: c.PrevVMProtoState.Clone(),
		VMProtoState:     c.VMProtoState.Clone(),
	}
}

type ConfirmValidOpportunity struct {
	*ConfirmNodeOpp
	MessagesData []byte
	MessageCount uint64
	LogsAcc      common.Hash
}

func (opp ConfirmValidOpportunity) Clone() ConfirmNodeOpportunity {
	messagesData := make([]byte, 0, len(opp.MessagesData))
	messagesData = append(messagesData, opp.MessagesData...)
	return ConfirmValidOpportunity{
		ConfirmNodeOpp: opp.ConfirmNodeOpp.Clone(),
		MessagesData:   messagesData,
		MessageCount:   opp.MessageCount,
		LogsAcc:        opp.LogsAcc,
	}
}

func (opp ConfirmValidOpportunity) CoreOpp() *ConfirmNodeOpp {
	return opp.ConfirmNodeOpp
}

func (opp ConfirmValidOpportunity) ProofSize() int {
	return 3 + len(opp.MessagesData)
}

type ConfirmInvalidOpportunity struct {
	*ConfirmNodeOpp
	ChallengeNodeData common.Hash
}

func (opp ConfirmInvalidOpportunity) Clone() ConfirmNodeOpportunity {
	return ConfirmInvalidOpportunity{
		ConfirmNodeOpp:    opp.ConfirmNodeOpp.Clone(),
		ChallengeNodeData: opp.ChallengeNodeData,
	}
}

func (opp ConfirmInvalidOpportunity) CoreOpp() *ConfirmNodeOpp {
	return opp.ConfirmNodeOpp
}

func (opp ConfirmInvalidOpportunity) ProofSize() int {
	return 3
}

type ConfirmNodeOpportunity interface {
	Clone() ConfirmNodeOpportunity
	CoreOpp() *ConfirmNodeOpp
	ProofSize() int
}

type ConfirmOpportunity struct {
	Nodes                  []ConfirmNodeOpportunity
	CurrentLatestConfirmed common.Hash
	StakerAddresses        []common.Address
	StakerProofs           [][]common.Hash
}

func (opp *ConfirmOpportunity) Clone() *ConfirmOpportunity {
	nodes := make([]ConfirmNodeOpportunity, 0, len(opp.Nodes))
	for _, node := range opp.Nodes {
		nodes = append(nodes, node.Clone())
	}
	stakerAddresses := append([]common.Address{}, opp.StakerAddresses...)
	stakerProofs := make([][]common.Hash, 0, len(opp.StakerProofs))
	for _, proof := range opp.StakerProofs {
		stakerProofs = append(stakerProofs, append([]common.Hash{}, proof...))
	}

	return &ConfirmOpportunity{
		Nodes:                  nodes,
		CurrentLatestConfirmed: opp.CurrentLatestConfirmed,
		StakerAddresses:        stakerAddresses,
		StakerProofs:           stakerProofs,
	}
}

type ConfirmProof struct {
	InitalProtoStateHash common.Hash
	BeforeSendCount      *big.Int
	BranchesNums         []*big.Int
	DeadlineTicks        []*big.Int
	ChallengeNodeData    [][32]byte
	LogsAcc              [][32]byte
	VMProtoStateHashes   [][32]byte
	MessageCounts        []*big.Int
	Messages             []byte
	CombinedProofs       [][32]byte
	StakerProofOffsets   []*big.Int
}

func (opp *ConfirmOpportunity) PrepareProof() ConfirmProof {
	nodeOpps := opp.Nodes
	branchesNums := make([]*big.Int, 0, len(nodeOpps))
	deadlineTicks := make([]*big.Int, 0, len(nodeOpps))
	challengeNodeData := make([]common.Hash, 0)
	logsAcc := make([]common.Hash, 0)
	vmProtoStateHashes := make([]common.Hash, 0)
	messageCounts := make([]*big.Int, 0)
	messageData := make([]byte, 0)

	for _, opp := range nodeOpps {
		core := opp.CoreOpp()
		branchesNums = append(branchesNums, new(big.Int).SetUint64(uint64(core.Branch)))
		deadlineTicks = append(deadlineTicks, core.DeadlineTicks.Val)

		switch opp := opp.(type) {
		case ConfirmValidOpportunity:
			logsAcc = append(logsAcc, opp.LogsAcc)
			vmProtoStateHashes = append(vmProtoStateHashes, core.VMProtoState.Hash())
			messageData = append(messageData, opp.MessagesData...)
			messageCounts = append(messageCounts, new(big.Int).SetUint64(opp.MessageCount))
		case ConfirmInvalidOpportunity:
			challengeNodeData = append(challengeNodeData, opp.ChallengeNodeData)
		}
	}

	combinedProofs := make([]common.Hash, 0)
	stakerProofOffsets := make([]*big.Int, 0, len(opp.StakerAddresses))
	stakerProofOffsets = append(stakerProofOffsets, big.NewInt(0))
	for _, proof := range opp.StakerProofs {
		combinedProofs = append(combinedProofs, proof...)
		stakerProofOffsets = append(stakerProofOffsets, big.NewInt(int64(len(combinedProofs))))
	}

	return ConfirmProof{
		InitalProtoStateHash: nodeOpps[0].CoreOpp().VMProtoState.Hash(),
		BeforeSendCount:      nodeOpps[0].CoreOpp().PrevVMProtoState.MessageCount,
		BranchesNums:         branchesNums,
		DeadlineTicks:        deadlineTicks,
		ChallengeNodeData:    common.HashSliceToRaw(challengeNodeData),
		LogsAcc:              common.HashSliceToRaw(logsAcc),
		VMProtoStateHashes:   common.HashSliceToRaw(vmProtoStateHashes),
		MessageCounts:        messageCounts,
		Messages:             messageData,
		CombinedProofs:       common.HashSliceToRaw(combinedProofs),
		StakerProofOffsets:   stakerProofOffsets,
	}
}
