package challenge

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/pkg/errors"
	"math/big"
)

type inboxDelta struct {
	inboxDeltaAccs [][32]byte
	inboxMessages  []inbox.InboxMessage
}

type Challenger struct {
	challenge           *ethbridge.Challenge
	lookup              core.ValidatorLookup
	challengedAssertion *core.Assertion

	inboxDelta *inboxDelta
}

func (c *Challenger) InboxDelta() (*inboxDelta, error) {
	if c.inboxDelta != nil {
		messages, err := c.lookup.GetMessages(
			c.challengedAssertion.PrevState.InboxCount,
			c.challengedAssertion.ExecInfo.InboxMessagesRead,
		)
		if err != nil {
			return nil, err
		}
		inboxDeltaAccs := make([][32]byte, 0, len(messages)+1)
		inboxDeltaAcc := common.Hash{}
		inboxDeltaAccs = append(inboxDeltaAccs, inboxDeltaAcc)
		for i := range messages {
			msg := messages[len(messages)-1-i]
			inboxDeltaAcc = hashing.SoliditySHA3(hashing.Bytes32(inboxDeltaAcc), msg.AsValue().Hash())
			inboxDeltaAccs = append(inboxDeltaAccs, inboxDeltaAcc)
		}
		c.inboxDelta = &inboxDelta{
			inboxDeltaAccs: inboxDeltaAccs,
			inboxMessages:  messages,
		}
	}
	return c.inboxDelta, nil
}

func NewChallenger(challenge *ethbridge.Challenge, lookup core.ValidatorLookup, challengedAssertion *core.Assertion) *Challenger {
	return &Challenger{
		challenge:           challenge,
		lookup:              lookup,
		challengedAssertion: challengedAssertion,
	}
}

func (c *Challenger) handleConflict(ctx context.Context) (*types.Transaction, error) {
	responder, err := c.challenge.CurrentResponder(ctx)
	if err != nil {
		return nil, err
	}
	if responder != c.challenge.Transactor() {
		// Not our turn
		return nil, nil
	}
	kind, err := c.challenge.Kind(ctx)
	if err != nil {
		return nil, err
	}

	switch kind {
	case core.UNINITIALIZED:
		judgment, err := core.JudgeAssertion(c.lookup, c.challengedAssertion, nil)
		if err != nil {
			return nil, err
		}
		switch judgment {
		case core.INBOX_CONSISTENCY:
			return c.handleInboxConsistencyChallenge(ctx)
		case core.INBOX_DELTA:
			return c.handleInboxDeltaChallenge(ctx)
		case core.EXECUTION:
			return c.handleExecutionChallenge()
		case core.STOPPED_SHORT:
			return c.handleStoppedShortChallenge()
		default:
			return nil, errors.New("can't handle challenge")
		}
	case core.INBOX_CONSISTENCY:
		return c.handleInboxConsistencyChallenge(ctx)
	case core.INBOX_DELTA:
		return c.handleInboxDeltaChallenge(ctx)
	case core.EXECUTION:
		return c.handleExecutionChallenge()
	case core.STOPPED_SHORT:
		return c.handleStoppedShortChallenge()
	default:
		return nil, errors.New("can't handle challenge")
	}

}

func (c *Challenger) handleInboxConsistencyChallenge(ctx context.Context) (*types.Transaction, error) {
	challengeImpl := &InboxConsistencyImpl{}
	return handleChallenge(ctx, c.challenge, c.lookup, challengeImpl)
}

func (c *Challenger) handleInboxDeltaChallenge(ctx context.Context) (*types.Transaction, error) {
	inboxDeltaData, err := c.InboxDelta()
	if err != nil {
		return nil, err
	}
	challengeImpl := &InboxDeltaImpl{
		nodeAfterInboxCount: c.challengedAssertion.AfterInboxCount(),
		inboxDelta:          inboxDeltaData,
	}
	return handleChallenge(ctx, c.challenge, c.lookup, challengeImpl)
}

func (c *Challenger) handleExecutionChallenge() (*types.Transaction, error) {
	return nil, nil
}

func (c *Challenger) handleStoppedShortChallenge() (*types.Transaction, error) {
	return nil, nil
}

type SimpleChallengerImpl interface {
	GetCut(lookup core.ValidatorLookup, offsets *big.Int) (core.Cut, error)
}

type ChallengerImpl interface {
	GetCuts(lookup core.ValidatorLookup, offsets []*big.Int) ([]core.Cut, error)
	FindFirstDivergence(lookup core.ValidatorLookup, offsets []*big.Int, cuts []core.Cut) (int, error)

	Bisect(
		ctx context.Context,
		challenge *ethbridge.Challenge,
		prevBisection *core.Bisection,
		segmentToChallenge int,
		subCuts []core.Cut,
	) (*types.Transaction, error)

	OneStepProof(
		ctx context.Context,
		challenge *ethbridge.Challenge,
		lookup core.ValidatorLookup,
		prevBisection *core.Bisection,
		segmentToChallenge int,
		challengedSegment *core.ChallengeSegment,
	) (*types.Transaction, error)
}

func handleChallenge(ctx context.Context, challenge *ethbridge.Challenge, lookup core.ValidatorLookup, challengeImpl ChallengerImpl) (*types.Transaction, error) {
	challengeState, err := challenge.ChallengeState(ctx)
	if err != nil {
		return nil, err
	}

	prevBisection, err := challenge.LookupBisection(ctx, challengeState)
	if err != nil {
		return nil, err
	}

	prevCutOffsets := generateBisectionCutOffsets(prevBisection.ChallengedSegment)
	cutToChallenge, err := challengeImpl.FindFirstDivergence(lookup, prevCutOffsets, prevBisection.Cuts)
	if err != nil {
		return nil, err
	}
	inconsistentSegment := &core.ChallengeSegment{
		Start:  prevCutOffsets[cutToChallenge-1],
		Length: new(big.Int).Sub(prevCutOffsets[cutToChallenge], prevCutOffsets[cutToChallenge-1]),
	}

	if inconsistentSegment.Length.Cmp(big.NewInt(1)) == 0 {
		return challengeImpl.OneStepProof(
			ctx,
			challenge,
			lookup,
			prevBisection,
			cutToChallenge+1,
			inconsistentSegment,
		)
	} else {
		subCutOffsets := generateBisectionCutOffsets(inconsistentSegment)
		subCuts, err := challengeImpl.GetCuts(lookup, subCutOffsets)
		if err != nil {
			return nil, err
		}
		return challengeImpl.Bisect(
			ctx,
			challenge,
			prevBisection,
			cutToChallenge+1,
			subCuts,
		)
	}
}

func findFirstDivergenceSimple(impl SimpleChallengerImpl, lookup core.ValidatorLookup, cutOffsets []*big.Int, cuts []core.Cut) (int, error) {
	for i, cutOffset := range cutOffsets {
		correctCut, err := impl.GetCut(lookup, cutOffset)
		if err != nil {
			return 0, err
		}
		if !correctCut.Equals(cuts[i]) {
			if i == 0 {
				return 0, errors.New("first segment was already wrong")
			}
			return i, nil
		}
	}
	return 0, errors.New("all cuts correct")
}

func getCutsSimple(impl SimpleChallengerImpl, lookup core.ValidatorLookup, offsets []*big.Int) ([]core.Cut, error) {
	cuts := make([]core.Cut, 0, len(offsets))
	for _, cutOffset := range offsets {
		cut, err := impl.GetCut(lookup, cutOffset)
		if err != nil {
			return nil, err
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func generateBisectionCutOffsets(segment *core.ChallengeSegment) []*big.Int {
	segmentCount := 20
	if segment.Length.Cmp(big.NewInt(int64(segmentCount))) < 0 {
		// Safe since this is less than 20
		segmentCount = int(segment.Length.Int64())
	}
	offset := new(big.Int).Set(segment.Start)
	cutOffsets := make([]*big.Int, 0, segmentCount+1)
	cutOffsets = append(cutOffsets, offset)
	for i := 0; i < segmentCount; i++ {
		subSegmentLength := new(big.Int).Div(segment.Length, big.NewInt(int64(segmentCount)))
		if i == 0 {
			subSegmentLength = subSegmentLength.Add(subSegmentLength, new(big.Int).Mod(segment.Length, big.NewInt(int64(segmentCount))))
		}
		offset = offset.Add(offset, subSegmentLength)
		cutOffsets = append(cutOffsets, offset)
	}
	return cutOffsets
}
