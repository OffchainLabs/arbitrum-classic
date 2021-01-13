package core

import (
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

type ChallengeKind uint8

const (
	UNINITIALIZED ChallengeKind = iota
	INBOX_CONSISTENCY
	INBOX_DELTA
	EXECUTION
	STOPPED_SHORT
	NO_CHALLENGE
)

type NodeID *big.Int

type NodeInfo struct {
	NodeNum       NodeID
	BlockProposed *common.BlockId
	Assertion     *Assertion
	InboxMaxCount *big.Int
	InboxMaxHash  common.Hash
}

func (n *NodeInfo) AfterState() *NodeState {
	return &NodeState{
		ProposedBlock:  n.BlockProposed.Height.AsInt(),
		InboxMaxCount:  n.InboxMaxCount,
		ExecutionState: n.Assertion.After,
	}
}

func (n *NodeInfo) InitialInboxConsistencyBisection() *Bisection {
	return &Bisection{
		ChallengedSegment: &ChallengeSegment{
			Start:  big.NewInt(0),
			Length: new(big.Int).Sub(n.InboxMaxCount, n.Assertion.After.InboxIndex),
		},
		Cuts: []Cut{
			NewSimpleCut(n.InboxMaxHash),
			NewSimpleCut(n.Assertion.After.InboxHash),
		},
	}
}

func (n *NodeInfo) InitialInboxDeltaBisection() *Bisection {
	beforeCut := InboxDeltaCut{
		InboxAccHash:   n.Assertion.After.InboxHash,
		InboxDeltaHash: [32]byte{},
	}
	afterCut := InboxDeltaCut{
		InboxAccHash:   n.Assertion.Before.InboxHash,
		InboxDeltaHash: n.Assertion.InboxDelta,
	}
	return &Bisection{
		ChallengedSegment: &ChallengeSegment{
			Start:  big.NewInt(0),
			Length: n.Assertion.InboxMessagesRead(),
		},
		Cuts: []Cut{beforeCut, afterCut},
	}
}

func (n *NodeInfo) InitialExecutionBisection() *Bisection {
	return &Bisection{
		ChallengedSegment: &ChallengeSegment{
			Start:  big.NewInt(0),
			Length: n.Assertion.GasUsed(),
		},
		Cuts: []Cut{
			ExecutionCut{
				GasUsed:      big.NewInt(0),
				InboxDelta:   n.Assertion.InboxDelta,
				MachineState: n.Assertion.Before.MachineHash,
				SendAcc:      common.Hash{},
				SendCount:    big.NewInt(0),
				LogAcc:       common.Hash{},
				LogCount:     big.NewInt(0),
			},
			ExecutionCut{
				GasUsed:      n.Assertion.GasUsed(),
				InboxDelta:   common.Hash{},
				MachineState: n.Assertion.After.MachineHash,
				SendAcc:      n.Assertion.SendAcc,
				SendCount:    n.Assertion.SendCount(),
				LogAcc:       n.Assertion.LogAcc,
				LogCount:     n.Assertion.LogCount(),
			},
		},
	}
}

func CalculateInboxDeltaAcc(messages []inbox.InboxMessage) common.Hash {
	acc := common.Hash{}
	for i := range messages {
		valHash := messages[len(messages)-1-i].AsValue().Hash()
		acc = hashing.SoliditySHA3(hashing.Bytes32(acc), hashing.Bytes32(valHash))
	}
	return acc
}
