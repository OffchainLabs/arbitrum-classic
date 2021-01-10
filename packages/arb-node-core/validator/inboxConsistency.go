package validator

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"math/big"
)

type InboxConsistencyImpl struct{}

func (i *InboxConsistencyImpl) GetCuts(lookup core.ValidatorLookup, offsets []*big.Int) ([]ethbridge.Cut, error) {
	return getCutsSimple(i, lookup, offsets)
}

func (i *InboxConsistencyImpl) FindFirstDivergence(lookup core.ValidatorLookup, offsets []*big.Int, cuts []ethbridge.Cut) (int, error) {
	return findFirstDivergenceSimple(i, lookup, offsets, cuts)
}

func (i *InboxConsistencyImpl) GetCut(lookup core.ValidatorLookup, offset *big.Int) (ethbridge.Cut, error) {
	inboxAcc, err := lookup.GetInboxAcc(offset)
	if err != nil {
		return nil, err
	}
	return ethbridge.InboxConsistencyCut{InboxAccHash: inboxAcc}, nil
}

func (i *InboxConsistencyImpl) Bisect(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	prevBisection *ethbridge.Bisection,
	segmentToChallenge int,
	subCuts []ethbridge.Cut,
) (*types.Transaction, error) {
	return challenge.BisectInboxDelta(
		ctx,
		prevBisection.Cuts,
		segmentToChallenge,
		prevBisection.ChallengedSegment,
		subCuts,
	)
}

func (i *InboxConsistencyImpl) OneStepProof(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	lookup core.ValidatorLookup,
	prevBisection *ethbridge.Bisection,
	segmentToChallenge int,
	challengedSegment *ethbridge.ChallengeSegment,
) (*types.Transaction, error) {
	beforeInboxAcc, err := lookup.GetInboxAcc(challengedSegment.Start)
	if err != nil {
		return nil, err
	}
	msgs, err := lookup.GetMessages(challengedSegment.Start, big.NewInt(1))
	if err != nil {
		return nil, err
	}
	return challenge.OneStepProveInboxConsistency(
		ctx,
		prevBisection.Cuts,
		segmentToChallenge,
		prevBisection.ChallengedSegment,
		beforeInboxAcc,
		msgs[0].CommitmentHash(),
	)
}
