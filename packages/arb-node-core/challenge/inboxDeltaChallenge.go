package challenge

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
)

type InboxDeltaImpl struct {
	nodeAfterInboxCount *big.Int
	inboxDelta          *inboxDelta
}

func (i *InboxDeltaImpl) GetCuts(lookup core.ValidatorLookup, offsets []*big.Int) ([]core.Cut, error) {
	return getCutsSimple(i, lookup, offsets)
}

func (i *InboxDeltaImpl) FindFirstDivergence(lookup core.ValidatorLookup, offsets []*big.Int, cuts []core.Cut) (int, error) {
	return findFirstDivergenceSimple(i, lookup, offsets, cuts)
}

func (i *InboxDeltaImpl) GetCut(lookup core.ValidatorLookup, offset *big.Int) (core.Cut, error) {
	inboxOffset := new(big.Int).Sub(i.nodeAfterInboxCount, offset)
	//inboxOffset = inboxOffset.Add(inboxOffset, big.NewInt(1))
	inboxAcc, err := lookup.GetInboxAcc(inboxOffset)
	if err != nil {
		fmt.Println("GetCut failed", offset, inboxOffset)
		return nil, err
	}
	fmt.Println("GetCut", offset, inboxOffset, inboxAcc, common.Hash(i.inboxDelta.inboxDeltaAccs[offset.Uint64()]))
	return core.InboxDeltaCut{
		InboxAccHash:   inboxAcc,
		InboxDeltaHash: i.inboxDelta.inboxDeltaAccs[offset.Uint64()],
	}, nil
}

func (i *InboxDeltaImpl) Bisect(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	inconsistentSegment *core.ChallengeSegment,
	subCuts []core.Cut,
) (*types.Transaction, error) {
	return challenge.BisectInboxDelta(
		ctx,
		prevBisection,
		segmentToChallenge,
		inconsistentSegment,
		subCuts,
	)
}

func (i *InboxDeltaImpl) OneStepProof(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	lookup core.ValidatorLookup,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
) (*types.Transaction, error) {
	inboxOffset := new(big.Int).Sub(i.nodeAfterInboxCount, challengedSegment.Start)
	inboxOffset = inboxOffset.Sub(inboxOffset, big.NewInt(1))
	msgs, err := lookup.GetMessages(inboxOffset, big.NewInt(1))
	if err != nil {
		return nil, err
	}
	return challenge.OneStepProveInboxDelta(
		ctx,
		prevBisection,
		segmentToChallenge,
		challengedSegment,
		msgs[0],
	)
}
