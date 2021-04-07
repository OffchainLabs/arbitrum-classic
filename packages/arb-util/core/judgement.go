package core

import (
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type ChallengeKind uint8

const (
	Uninitialized ChallengeKind = iota
	Execution
	StoppedShort
	NoChallenge
)

type NodeID *big.Int

type NodeInfo struct {
	NodeNum            NodeID
	BlockProposed      *common.BlockId
	Assertion          *Assertion
	InboxMaxCount      *big.Int
	NodeHash           common.Hash
	AfterInboxBatchAcc [32]byte
}

func (n *NodeInfo) AfterState() *NodeState {
	return &NodeState{
		ProposedBlock:  n.BlockProposed.Height.AsInt(),
		InboxMaxCount:  n.InboxMaxCount,
		ExecutionState: n.Assertion.After,
	}
}

func (n *NodeInfo) InitialExecutionBisection() *Bisection {
	return &Bisection{
		ChallengedSegment: &ChallengeSegment{
			Start:  n.Assertion.Before.TotalGasConsumed,
			Length: new(big.Int).Sub(n.Assertion.After.TotalGasConsumed, n.Assertion.Before.TotalGasConsumed),
		},
		Cuts: []Cut{
			n.Assertion.Before,
			n.Assertion.After,
		},
	}
}
