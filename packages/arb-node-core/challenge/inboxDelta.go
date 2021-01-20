package challenge

import (
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"math/big"
)

type InboxDeltaImpl struct {
	nodeAfterInboxCount *big.Int
	inboxDelta          *inboxDelta
}

func (i *InboxDeltaImpl) SegmentTarget() int {
	return 250
}

func (i *InboxDeltaImpl) GetCuts(lookup core.ValidatorLookup, offsets []*big.Int) ([]core.Cut, error) {
	return getCutsSimple(i, lookup, offsets)
}

func (i *InboxDeltaImpl) FindFirstDivergence(lookup core.ValidatorLookup, offsets []*big.Int, cuts []core.Cut) (int, error) {
	return findFirstDivergenceSimple(i, lookup, offsets, cuts)
}

func (i *InboxDeltaImpl) GetCut(lookup core.ValidatorLookup, offset *big.Int) (core.Cut, error) {
	inboxOffset := new(big.Int).Sub(i.nodeAfterInboxCount, offset)
	inboxAcc, err := lookup.GetInboxAcc(inboxOffset)
	if err != nil {
		return nil, err
	}
	return core.InboxDeltaCut{
		InboxAccHash:   inboxAcc,
		InboxDeltaHash: i.inboxDelta.inboxDeltaAccs[offset.Uint64()],
	}, nil
}

func (i *InboxDeltaImpl) Bisect(
	challenge *ethbridge.Challenge,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	inconsistentSegment *core.ChallengeSegment,
	subCuts []core.Cut,
) (*ethbridge.RawTransaction, error) {
	return challenge.BisectInboxDelta(
		prevBisection,
		segmentToChallenge,
		inconsistentSegment,
		subCuts,
	)
}

func (i *InboxDeltaImpl) OneStepProof(
	challenge *ethbridge.Challenge,
	lookup core.ValidatorLookup,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
) (*ethbridge.RawTransaction, error) {
	inboxOffset := new(big.Int).Sub(i.nodeAfterInboxCount, challengedSegment.Start)
	inboxOffset = inboxOffset.Sub(inboxOffset, big.NewInt(1))
	msgs, err := lookup.GetMessages(inboxOffset, big.NewInt(1))
	if err != nil {
		return nil, err
	}
	return challenge.OneStepProveInboxDelta(
		prevBisection,
		segmentToChallenge,
		challengedSegment,
		msgs[0],
	)
}
