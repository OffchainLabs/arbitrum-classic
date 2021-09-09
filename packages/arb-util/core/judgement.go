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
	NodeNum                 NodeID
	BlockProposed           *common.BlockId
	Assertion               *Assertion
	InboxMaxCount           *big.Int
	NodeHash                common.Hash
	AfterInboxBatchEndCount *big.Int
	AfterInboxBatchAcc      common.Hash
}

func (n *NodeInfo) AfterState() *NodeState {
	return &NodeState{
		ProposedBlock:  n.BlockProposed.Height.AsInt(),
		InboxMaxCount:  n.InboxMaxCount,
		ExecutionState: n.Assertion.After,
	}
}
