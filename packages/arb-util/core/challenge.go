package core

import (
	"fmt"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type Cut interface {
	Equals(other Cut) bool
	CutHash() [32]byte
}

type SimpleCut struct {
	hash [32]byte
}

func NewSimpleCut(hash [32]byte) SimpleCut {
	return SimpleCut{hash: hash}
}

func (c SimpleCut) String() string {
	return fmt.Sprintf("SimpleCut(0x%x)", c.hash)
}

func (c SimpleCut) Equals(other Cut) bool {
	o, ok := other.(SimpleCut)
	if !ok {
		return false
	}
	return c.hash == o.hash
}

func (c SimpleCut) CutHash() [32]byte {
	return c.hash
}

type ExecutionCut struct {
	GasUsed           *big.Int
	TotalMessagesRead *big.Int
	MachineState      common.Hash
	SendAcc           common.Hash
	SendCount         *big.Int
	LogAcc            common.Hash
	LogCount          *big.Int
}

func (c ExecutionCut) Equals(other Cut) bool {
	o, ok := other.(ExecutionCut)
	if !ok {
		return false
	}
	return c.GasUsed.Cmp(o.GasUsed) == 0 &&
		c.TotalMessagesRead.Cmp(o.TotalMessagesRead) == 0 &&
		c.MachineState == o.MachineState &&
		c.SendAcc == o.SendAcc &&
		c.SendCount.Cmp(o.SendCount) == 0 &&
		c.LogAcc == o.LogAcc &&
		c.LogCount.Cmp(o.LogCount) == 0
}

func (c ExecutionCut) RestHash() [32]byte {
	return assertionRestHash(
		c.TotalMessagesRead,
		c.MachineState,
		c.SendAcc,
		c.SendCount,
		c.LogAcc,
		c.LogCount,
	)
}

func (c ExecutionCut) CutHash() [32]byte {
	return assertionHash(c.GasUsed, c.RestHash())
}

type Bisection struct {
	ChallengedSegment *ChallengeSegment
	Cuts              []Cut
}

type ChallengeSegment struct {
	Start  *big.Int
	Length *big.Int
}
