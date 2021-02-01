package challenge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
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
	challenge      *ethbridge.Challenge
	lookup         core.ValidatorLookup
	challengedNode *core.NodeInfo
	stakerAddress  common.Address

	inboxDelta *inboxDelta
}

func (c *Challenger) ChallengeAddress() common.Address {
	return c.challenge.Address()
}

func NewChallenger(challenge *ethbridge.Challenge, lookup core.ValidatorLookup, challengedNode *core.NodeInfo, stakerAddress common.Address) *Challenger {
	return &Challenger{
		challenge:      challenge,
		lookup:         lookup,
		challengedNode: challengedNode,
		stakerAddress:  stakerAddress,
	}
}

func (c *Challenger) getInboxDelta() (*inboxDelta, error) {
	if c.inboxDelta == nil {
		messages, err := c.lookup.GetMessages(
			c.challengedNode.Assertion.Before.InboxIndex,
			c.challengedNode.Assertion.InboxMessagesRead(),
		)
		if err != nil {
			return nil, err
		}
		inboxDeltaAccs := make([][32]byte, 0, len(messages)+1)
		inboxDeltaAcc := common.Hash{}
		inboxDeltaAccs = append(inboxDeltaAccs, inboxDeltaAcc)
		for i := range messages {
			msg := messages[len(messages)-1-i]
			inboxDeltaAcc = hashing.SoliditySHA3(hashing.Bytes32(inboxDeltaAcc), hashing.Bytes32(msg.AsValue().Hash()))
			inboxDeltaAccs = append(inboxDeltaAccs, inboxDeltaAcc)
		}
		c.inboxDelta = &inboxDelta{
			inboxDeltaAccs: inboxDeltaAccs,
			inboxMessages:  messages,
		}
	}
	return c.inboxDelta, nil
}

func (c *Challenger) HandleConflict(ctx context.Context) error {
	responder, err := c.challenge.CurrentResponder(ctx)
	if err != nil {
		return err
	}
	if responder != c.stakerAddress {
		// Not our turn
		return nil
	}
	kind, err := c.challenge.Kind(ctx)
	if err != nil {
		return err
	}

	var prevBisection *core.Bisection
	if kind == core.UNINITIALIZED {
		startCursor, err := c.lookup.GetExecutionCursor(c.challengedNode.Assertion.Before.TotalGasConsumed)
		if err != nil {
			return err
		}
		execTracker := core.NewExecutionTracker(
			c.lookup,
			startCursor,
			false,
			[]*big.Int{c.challengedNode.Assertion.GasUsed()},
		)
		kind, err = core.JudgeAssertion(c.lookup, c.challengedNode.Assertion, execTracker)
		if err != nil {
			return err
		}
		if kind == core.STOPPED_SHORT {
			panic("Not yet handled")
		}
	} else {
		challengeState, err := c.challenge.ChallengeState(ctx)
		if err != nil {
			return err
		}

		prevBisection, err = c.challenge.LookupBisection(ctx, challengeState)
		if err != nil {
			return err
		}
	}

	switch kind {
	case core.INBOX_CONSISTENCY:
		return c.handleInboxConsistencyChallenge(ctx, prevBisection)
	case core.INBOX_DELTA:
		return c.handleInboxDeltaChallenge(ctx, prevBisection)
	case core.EXECUTION:
		return c.handleExecutionChallenge(ctx, prevBisection)
	case core.STOPPED_SHORT:
		return c.handleStoppedShortChallenge()
	default:
		return errors.New("can't handle challenge")
	}
}

func (c *Challenger) handleInboxConsistencyChallenge(ctx context.Context, prevBisection *core.Bisection) error {
	challengeImpl := &InboxConsistencyImpl{
		inboxMaxCount: c.challengedNode.InboxMaxCount,
	}
	if prevBisection == nil {
		prevBisection = c.challengedNode.InitialInboxConsistencyBisection()
	}
	return handleChallenge(ctx, c.challenge, c.lookup, challengeImpl, prevBisection)
}

func (c *Challenger) handleInboxDeltaChallenge(ctx context.Context, prevBisection *core.Bisection) error {
	inboxDeltaData, err := c.getInboxDelta()
	if err != nil {
		return err
	}
	challengeImpl := &InboxDeltaImpl{
		nodeAfterInboxCount: c.challengedNode.Assertion.After.InboxIndex,
		inboxDelta:          inboxDeltaData,
	}
	if prevBisection == nil {
		prevBisection = c.challengedNode.InitialInboxDeltaBisection()
	}
	return handleChallenge(ctx, c.challenge, c.lookup, challengeImpl, prevBisection)
}

func (c *Challenger) handleExecutionChallenge(ctx context.Context, prevBisection *core.Bisection) error {
	if prevBisection == nil {
		prevBisection = c.challengedNode.InitialExecutionBisection()
	}
	inboxDeltaData, err := c.getInboxDelta()
	initialCursor, err := c.lookup.GetExecutionCursor(c.challengedNode.Assertion.Before.TotalGasConsumed)
	if err != nil {
		return err
	}
	challengeImpl := &ExecutionImpl{
		initialCursor: initialCursor,
		inboxDelta:    inboxDeltaData,
	}
	return handleChallenge(ctx, c.challenge, c.lookup, challengeImpl, prevBisection)
}

func (c *Challenger) handleStoppedShortChallenge() error {
	panic("Unimplemented")
}

type SimpleChallengerImpl interface {
	GetCut(lookup core.ValidatorLookup, offsets *big.Int) (core.Cut, error)
}

type ChallengerImpl interface {
	SegmentTarget() int

	GetCuts(lookup core.ValidatorLookup, offsets []*big.Int) ([]core.Cut, error)
	FindFirstDivergence(lookup core.ValidatorLookup, offsets []*big.Int, cuts []core.Cut) (int, error)

	Bisect(
		ctx context.Context,
		challenge *ethbridge.Challenge,
		prevBisection *core.Bisection,
		segmentToChallenge int,
		inconsistentSegment *core.ChallengeSegment,
		subCuts []core.Cut,
	) error

	OneStepProof(
		ctx context.Context,
		challenge *ethbridge.Challenge,
		lookup core.ValidatorLookup,
		prevBisection *core.Bisection,
		segmentToChallenge int,
		challengedSegment *core.ChallengeSegment,
	) error
}

func handleChallenge(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	lookup core.ValidatorLookup,
	challengeImpl ChallengerImpl,
	prevBisection *core.Bisection,
) error {
	prevCutOffsets := generateBisectionCutOffsets(prevBisection.ChallengedSegment, len(prevBisection.Cuts)-1)
	cutToChallenge, err := challengeImpl.FindFirstDivergence(lookup, prevCutOffsets, prevBisection.Cuts)
	if err != nil {
		return err
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
			cutToChallenge-1,
			inconsistentSegment,
		)
	} else {
		segmentCount := challengeImpl.SegmentTarget()
		if inconsistentSegment.Length.Cmp(big.NewInt(int64(segmentCount))) < 0 {
			// Safe since this is less than 400
			segmentCount = int(inconsistentSegment.Length.Int64())
		}
		subCutOffsets := generateBisectionCutOffsets(inconsistentSegment, segmentCount)
		subCuts, err := challengeImpl.GetCuts(lookup, subCutOffsets)
		if err != nil {
			return err
		}
		return challengeImpl.Bisect(
			ctx,
			challenge,
			prevBisection,
			cutToChallenge-1,
			inconsistentSegment,
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
				return 0, errors.New("first cut was already wrong")
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

func generateBisectionCutOffsets(segment *core.ChallengeSegment, subSegmentCount int) []*big.Int {
	cutCount := subSegmentCount + 1
	offset := new(big.Int).Set(segment.Start)
	cutOffsets := make([]*big.Int, 0, cutCount)
	for i := 0; i < cutCount; i++ {
		cutOffsets = append(cutOffsets, new(big.Int).Set(offset))
		subSegmentLength := new(big.Int).Div(segment.Length, big.NewInt(int64(subSegmentCount)))
		if i == 0 {
			subSegmentLength = subSegmentLength.Add(subSegmentLength, new(big.Int).Mod(segment.Length, big.NewInt(int64(subSegmentCount))))
		}
		offset = offset.Add(offset, subSegmentLength)
	}
	return cutOffsets
}
