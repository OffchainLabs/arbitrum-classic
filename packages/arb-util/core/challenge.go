package core

import (
	"fmt"
	"math/big"
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
	return c.CutHash() == other.CutHash()
}

func (c SimpleCut) CutHash() [32]byte {
	return c.hash
}

type Bisection struct {
	ChallengedSegment *ChallengeSegment
	Cuts              []Cut
}

type ChallengeSegment struct {
	Start  *big.Int
	Length *big.Int
}

func (s ChallengeSegment) GetEnd() *big.Int {
	return new(big.Int).Add(s.Start, s.Length)
}
