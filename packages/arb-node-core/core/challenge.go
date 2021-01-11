package core

import (
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"math/big"
)

type Cut interface {
	Equals(other Cut) bool
	Hash() [32]byte
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

func (c SimpleCut) Hash() [32]byte {
	return c.hash
}

type InboxDeltaCut struct {
	InboxAccHash   [32]byte
	InboxDeltaHash [32]byte
}

func (c InboxDeltaCut) String() string {
	return fmt.Sprintf("InboxDeltaCut(0x%x, 0x%x)", c.InboxAccHash, c.InboxDeltaHash)
}

func (c InboxDeltaCut) Equals(other Cut) bool {
	o, ok := other.(InboxDeltaCut)
	if !ok {
		return false
	}
	return c.InboxAccHash == o.InboxAccHash && c.InboxDeltaHash == o.InboxDeltaHash
}

func (c InboxDeltaCut) Hash() [32]byte {
	return InboxDeltaHash(c.InboxAccHash, c.InboxDeltaHash)
}

type ExpandedExecutionCut struct {
	GasUsed *big.Int
	Rest    common.Hash
}

func (c ExpandedExecutionCut) Equals(other Cut) bool {
	o, ok := other.(ExpandedExecutionCut)
	if !ok {
		return false
	}
	return c.GasUsed.Cmp(o.GasUsed) == 0 && c.Rest == o.Rest
}

func (c ExpandedExecutionCut) Hash() [32]byte {
	return hashing.SoliditySHA3(hashing.Uint256(c.GasUsed), hashing.Bytes32(c.Rest))
}

type Bisection struct {
	ChallengedSegment *ChallengeSegment
	Cuts              []Cut
}

type ChallengeSegment struct {
	Start  *big.Int
	Length *big.Int
}
