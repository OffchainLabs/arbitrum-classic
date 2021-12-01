package core

import (
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type Bisection struct {
	ChallengedSegment *ChallengeSegment
	Cuts              []common.Hash
}

type ChallengeSegment struct {
	Start  *big.Int
	Length *big.Int
}

func (s ChallengeSegment) GetEnd() *big.Int {
	return new(big.Int).Add(s.Start, s.Length)
}
