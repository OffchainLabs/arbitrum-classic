package challenge

import (
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"math/big"
)

type InboxConsistencyImpl struct {
	inboxMaxCount *big.Int
}

func (i *InboxConsistencyImpl) SegmentTarget() int {
	return 400
}

func (i *InboxConsistencyImpl) inboxOffset(offset *big.Int) *big.Int {
	return new(big.Int).Sub(i.inboxMaxCount, offset)
}

func (i *InboxConsistencyImpl) GetCuts(lookup core.ValidatorLookup, offsets []*big.Int) ([]core.Cut, error) {
	return getCutsSimple(i, lookup, offsets)
}

func (i *InboxConsistencyImpl) FindFirstDivergence(lookup core.ValidatorLookup, offsets []*big.Int, cuts []core.Cut) (int, error) {
	return findFirstDivergenceSimple(i, lookup, offsets, cuts)
}

func (i *InboxConsistencyImpl) GetCut(lookup core.ValidatorLookup, offset *big.Int) (core.Cut, error) {
	inboxOffset := i.inboxOffset(offset)
	inboxAcc, err := lookup.GetInboxAcc(inboxOffset)
	if err != nil {
		return nil, err
	}
	return core.NewSimpleCut(inboxAcc), nil
}

func (i *InboxConsistencyImpl) Bisect(
	challenge *ethbridge.Challenge,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	subCuts []core.Cut,
) (*ethbridge.RawTransaction, error) {
	return challenge.BisectInboxConsistency(
		prevBisection,
		segmentToChallenge,
		challengedSegment,
		subCuts,
	)
}

func (i *InboxConsistencyImpl) OneStepProof(
	challenge *ethbridge.Challenge,
	lookup core.ValidatorLookup,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
) (*ethbridge.RawTransaction, error) {
	inboxOffset := new(big.Int).Sub(i.inboxOffset(challengedSegment.Start), big.NewInt(1))
	beforeInboxAcc, err := lookup.GetInboxAcc(inboxOffset)
	if err != nil {
		return nil, err
	}
	msgs, err := lookup.GetMessages(inboxOffset, big.NewInt(1))
	if err != nil {
		return nil, err
	}
	return challenge.OneStepProveInboxConsistency(
		prevBisection,
		segmentToChallenge,
		challengedSegment,
		beforeInboxAcc,
		msgs[0].CommitmentHash(),
	)
}
