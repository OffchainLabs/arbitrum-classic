package challenge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/pkg/errors"
	"math/big"
)

type Challenger struct {
	challenge      *ethbridge.Challenge
	lookup         core.ArbCoreLookup
	challengedNode *core.NodeInfo
	stakerAddress  common.Address
}

func (c *Challenger) ChallengeAddress() common.Address {
	return c.challenge.Address()
}

func NewChallenger(challenge *ethbridge.Challenge, lookup core.ArbCoreLookup, challengedNode *core.NodeInfo, stakerAddress common.Address) *Challenger {
	return &Challenger{
		challenge:      challenge,
		lookup:         lookup,
		challengedNode: challengedNode,
		stakerAddress:  stakerAddress,
	}
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
	if kind == core.Uninitialized {
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
		kind, err = core.JudgeAssertion(c.challengedNode.Assertion, execTracker)
		if err != nil {
			return err
		}
		if kind == core.StoppedShort {
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
	case core.Execution:
		return c.handleExecutionChallenge(ctx, prevBisection)
	case core.StoppedShort:
		return c.handleStoppedShortChallenge()
	default:
		return errors.New("can't handle challenge")
	}
}

func (c *Challenger) handleExecutionChallenge(ctx context.Context, prevBisection *core.Bisection) error {
	if prevBisection == nil {
		prevBisection = c.challengedNode.InitialExecutionBisection()
	}
	initialCursor, err := c.lookup.GetExecutionCursor(c.challengedNode.Assertion.Before.TotalGasConsumed)
	if err != nil {
		return err
	}
	challengeImpl := &ExecutionImpl{
		initialCursor: initialCursor,
	}
	return handleChallenge(ctx, c.challenge, c.lookup, challengeImpl, prevBisection)
}

func (c *Challenger) handleStoppedShortChallenge() error {
	panic("Unimplemented")
}

type SimpleChallengerImpl interface {
	GetCut(lookup core.ArbCoreLookup, offsets *big.Int) (core.Cut, error)
}

type ChallengerImpl interface {
	SegmentTarget() int

	GetCuts(lookup core.ArbCoreLookup, offsets []*big.Int) ([]core.Cut, error)
	FindFirstDivergence(lookup core.ArbCoreLookup, offsets []*big.Int, cuts []core.Cut) (int, error)

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
		lookup core.ArbCoreLookup,
		prevBisection *core.Bisection,
		segmentToChallenge int,
		challengedSegment *core.ChallengeSegment,
	) error
}

func handleChallenge(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	lookup core.ArbCoreLookup,
	challengeImpl ChallengerImpl,
	prevBisection *core.Bisection,
) error {
	prevCutOffsets := generateBisectionCutOffsets(prevBisection.ChallengedSegment, len(prevBisection.Cuts)-1)
	cutToChallenge, err := challengeImpl.FindFirstDivergence(lookup, prevCutOffsets, prevBisection.Cuts)
	if err != nil {
		return err
	}
	if cutToChallenge >= len(prevCutOffsets) {
		return errors.New("cannot challenge last cut")
	}
	inconsistentSegment := &core.ChallengeSegment{
		Start:  prevCutOffsets[cutToChallenge],
		Length: new(big.Int).Sub(prevCutOffsets[cutToChallenge+1], prevCutOffsets[cutToChallenge]),
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
			cutToChallenge,
			inconsistentSegment,
			subCuts,
		)
	}
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
