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
	NodeNum       NodeID
	BlockProposed *common.BlockId
	Assertion     *Assertion
	InboxMaxCount *big.Int
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
			Start:  big.NewInt(0),
			Length: n.Assertion.GasUsed(),
		},
		Cuts: []Cut{
			ExecutionCut{
				GasUsed:           big.NewInt(0),
				TotalMessagesRead: n.Assertion.Before.TotalMessagesRead,
				MachineState:      n.Assertion.Before.MachineHash,
				SendAcc:           common.Hash{},
				SendCount:         big.NewInt(0),
				LogAcc:            common.Hash{},
				LogCount:          big.NewInt(0),
			},
			ExecutionCut{
				GasUsed:           n.Assertion.GasUsed(),
				TotalMessagesRead: n.Assertion.After.TotalMessagesRead,
				MachineState:      n.Assertion.After.MachineHash,
				SendAcc:           n.Assertion.SendAcc,
				SendCount:         n.Assertion.SendCount(),
				LogAcc:            n.Assertion.LogAcc,
				LogCount:          n.Assertion.LogCount(),
			},
		},
	}
}
