package core

import (
	"github.com/pkg/errors"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
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
		TotalGasUsed:   n.Assertion.AfterTotalGasUsed(),
		MachineHash:    n.Assertion.ExecInfo.AfterMachineHash,
		InboxHash:      n.Assertion.AfterInboxHash,
		InboxCount:     n.Assertion.AfterInboxCount(),
		TotalSendCount: n.Assertion.AfterTotalSendCount(),
		TotalLogCount:  n.Assertion.AfterTotalLogCount(),
		InboxMaxCount:  n.InboxMaxCount,
	}
}

func (n *NodeInfo) InitialInboxConsistencyBisection() *Bisection {
	return &Bisection{
		ChallengedSegment: &ChallengeSegment{
			Start:  n.Assertion.PrevState.InboxCount,
			Length: new(big.Int).Sub(n.InboxMaxCount, n.Assertion.AfterInboxCount()),
		},
		Cuts: []Cut{
			NewSimpleCut(n.InboxMaxHash),
			NewSimpleCut(n.Assertion.AfterInboxHash),
		},
	}
}

func (n *NodeInfo) InitialInboxDeltaBisection() *Bisection {
	return &Bisection{
		ChallengedSegment: &ChallengeSegment{
			Start:  big.NewInt(0),
			Length: n.Assertion.ExecInfo.InboxMessagesRead,
		},
		Cuts: []Cut{
			InboxDeltaCut{
				InboxAccHash:   n.Assertion.AfterInboxHash,
				InboxDeltaHash: [32]byte{},
			},
			InboxDeltaCut{
				InboxAccHash:   n.Assertion.PrevState.InboxHash,
				InboxDeltaHash: n.Assertion.InboxDelta,
			},
		},
	}
}

func (n *NodeInfo) InitialExecutionBisection() *Bisection {
	return &Bisection{
		ChallengedSegment: &ChallengeSegment{
			Start:  big.NewInt(0),
			Length: n.Assertion.ExecInfo.GasUsed,
		},
		Cuts: []Cut{
			NewSimpleCut(n.Assertion.BeforeExecutionHash()),
			NewSimpleCut(n.Assertion.AfterExecutionHash()),
		},
	}
}

func JudgeAssertion(lookup ValidatorLookup, assertion *Assertion, mach machine.Machine) (ChallengeKind, error) {
	afterInboxHash, err := lookup.GetInboxAcc(assertion.AfterInboxCount())
	if err != nil {
		return 0, err
	}
	if assertion.AfterInboxHash != afterInboxHash {
		// Failed inbox consistency
		return INBOX_CONSISTENCY, nil
	}
	messages, err := lookup.GetMessages(assertion.PrevState.InboxCount, assertion.ExecInfo.InboxMessagesRead)
	if err != nil {
		return 0, err
	}
	if assertion.InboxDelta != CalculateInboxDeltaAcc(messages) {
		// Failed inbox delta
		return INBOX_DELTA, nil
	}
	if mach == nil {
		mach, err = lookup.GetMachine(assertion.PrevState.TotalGasUsed)
		if err != nil {
			return 0, err
		}
	}
	if mach.Hash() != assertion.PrevState.MachineHash {
		return 0, errors.New("before machine state inconsistent with local db")
	}
	localExecutionInfo, err := lookup.GetExecutionInfoWithMaxMessages(mach, assertion.ExecInfo.GasUsed, assertion.ExecInfo.InboxMessagesRead)
	if err != nil {
		return 0, err
	}

	if localExecutionInfo.GasUsed.Cmp(assertion.ExecInfo.GasUsed) < 0 {
		return STOPPED_SHORT, nil
	}

	if !assertion.ExecInfo.Equals(localExecutionInfo) {
		return EXECUTION, nil
	}
	return NO_CHALLENGE, nil
}

func CalculateInboxDeltaAcc(messages []inbox.InboxMessage) common.Hash {
	acc := common.Hash{}
	for i := range messages {
		valHash := messages[len(messages)-1-i].AsValue().Hash()
		acc = hashing.SoliditySHA3(hashing.Bytes32(acc), hashing.Bytes32(valHash))
	}
	return acc
}
