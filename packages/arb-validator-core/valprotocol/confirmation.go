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
	"bytes"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"math/big"
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
	initalProtoStateHash := nodeOpps[0].StateHash()
	branchesNums := make([]*big.Int, 0, len(nodeOpps))
	deadlineTicks := make([]*big.Int, 0, len(nodeOpps))
	challengeNodeData := make([]common.Hash, 0)
	logsAcc := make([]common.Hash, 0)
	vmProtoStateHashes := make([]common.Hash, 0)

	messageCounts := make([]*big.Int, 0)
	var messageData bytes.Buffer

	for _, opp := range nodeOpps {
		branchesNums = append(branchesNums, new(big.Int).SetUint64(uint64(opp.BranchType())))
		deadlineTicks = append(deadlineTicks, opp.Deadline().Val)

		switch opp := opp.(type) {
		case ConfirmValidOpportunity:
			logsAcc = append(logsAcc, opp.LogsAcc)
			vmProtoStateHashes = append(vmProtoStateHashes, opp.VMProtoStateHash)

			for _, msg := range opp.Messages {
				_ = value.MarshalValue(msg, &messageData)
			}
			messageCounts = append(messageCounts, big.NewInt(int64(len(opp.Messages))))
		case ConfirmInvalidOpportunity:
			challengeNodeData = append(challengeNodeData, opp.ChallengeNodeData)
		}
	}

	messages := messageData.Bytes()
	combinedProofs := make([]common.Hash, 0)
	stakerProofOffsets := make([]*big.Int, 0, len(opp.StakerAddresses))
	stakerProofOffsets = append(stakerProofOffsets, big.NewInt(0))
	for _, proof := range opp.StakerProofs {
		combinedProofs = append(combinedProofs, proof...)
		stakerProofOffsets = append(stakerProofOffsets, big.NewInt(int64(len(combinedProofs))))
	}

	return ConfirmProof{
		InitalProtoStateHash: initalProtoStateHash,
		BranchesNums:         branchesNums,
		DeadlineTicks:        deadlineTicks,
		ChallengeNodeData:    common.HashSliceToRaw(challengeNodeData),
		LogsAcc:              common.HashSliceToRaw(logsAcc),
		VMProtoStateHashes:   common.HashSliceToRaw(vmProtoStateHashes),
		MessageCounts:        messageCounts,
		Messages:             messages,
		CombinedProofs:       common.HashSliceToRaw(combinedProofs),
		StakerProofOffsets:   stakerProofOffsets,
	}
}
