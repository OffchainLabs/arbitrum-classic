package challenge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
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

func (i *InboxConsistencyImpl) GetCuts(lookup core.ArbCoreLookup, offsets []*big.Int) ([]core.Cut, error) {
	return getCutsSimple(i, lookup, offsets)
}

func (i *InboxConsistencyImpl) FindFirstDivergence(lookup core.ArbCoreLookup, offsets []*big.Int, cuts []core.Cut) (int, error) {
	return findFirstDivergenceSimple(i, lookup, offsets, cuts)
}

func (i *InboxConsistencyImpl) GetCut(lookup core.ArbCoreLookup, offset *big.Int) (core.Cut, error) {
	inboxOffset := i.inboxOffset(offset)
	inboxAcc, err := lookup.GetInboxAcc(inboxOffset)
	if err != nil {
		return nil, err
	}
	return core.NewSimpleCut(inboxAcc), nil
}

func (i *InboxConsistencyImpl) Bisect(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	subCuts []core.Cut,
) error {
	return challenge.BisectInboxConsistency(
		ctx,
		prevBisection,
		segmentToChallenge,
		challengedSegment,
		subCuts,
	)
}

func (i *InboxConsistencyImpl) OneStepProof(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	lookup core.ArbCoreLookup,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
) error {
	inboxOffset := new(big.Int).Sub(i.inboxOffset(challengedSegment.Start), big.NewInt(1))
	beforeInboxAcc, err := lookup.GetInboxAcc(inboxOffset)
	if err != nil {
		return err
	}
	msgs, err := lookup.GetMessages(new(big.Int).Add(inboxOffset, big.NewInt(1)), big.NewInt(1))
	if err != nil {
		return err
	}
	return challenge.OneStepProveInboxConsistency(
		ctx,
		prevBisection,
		segmentToChallenge,
		challengedSegment,
		beforeInboxAcc,
		msgs[0].CommitmentHash(),
	)
}
