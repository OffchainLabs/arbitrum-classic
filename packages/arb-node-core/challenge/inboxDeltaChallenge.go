package challenge

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"math/big"
)

type InboxDeltaImpl struct {
	nodeAfterInboxCount *big.Int
	inboxDelta          *inboxDelta
}

func (i *InboxDeltaImpl) GetCuts(lookup core.ValidatorLookup, offsets []*big.Int) ([]Cut, error) {
	return getCutsSimple(i, lookup, offsets)
}

func (i *InboxDeltaImpl) FindFirstDivergence(lookup core.ValidatorLookup, offsets []*big.Int, cuts []Cut) (int, error) {
	return findFirstDivergenceSimple(i, lookup, offsets, cuts)
}

func (i *InboxDeltaImpl) GetCut(lookup core.ValidatorLookup, offset *big.Int) (Cut, error) {
	inboxOffset := new(big.Int).Add(i.nodeAfterInboxCount, offset)
	inboxAcc, err := lookup.GetInboxAcc(inboxOffset)
	if err != nil {
		return nil, err
	}
	return InboxDeltaCut{
		InboxAccHash:   inboxAcc,
		InboxDeltaHash: i.inboxDelta.inboxDeltaAccs[offset.Uint64()],
	}, nil
}

func (i *InboxDeltaImpl) Bisect(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	prevBisection *Bisection,
	segmentToChallenge int,
	subCuts []Cut,
) (*types.Transaction, error) {
	return challenge.BisectInboxDelta(
		ctx,
		prevBisection,
		segmentToChallenge,
		subCuts,
	)
}

func (i *InboxDeltaImpl) OneStepProof(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	lookup core.ValidatorLookup,
	prevBisection *Bisection,
	segmentToChallenge int,
	challengedSegment *ethbridge.ChallengeSegment,
) (*types.Transaction, error) {
	msgIndex := new(big.Int).Add(i.nodeAfterInboxCount, challengedSegment.Start)
	msgs, err := lookup.GetMessages(msgIndex, big.NewInt(1))
	if err != nil {
		return nil, err
	}
	return challenge.OneStepProveInboxDelta(
		ctx,
		prevBisection,
		segmentToChallenge,
		msgs[0],
	)
}
