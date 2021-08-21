package challenge

import (
	"context"
	"math/big"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
)

var logger = log.With().Caller().Stack().Str("component", "challenge").Logger()

type Challenger struct {
	challenge      *ethbridge.Challenge
	sequencerInbox *ethbridge.SequencerInboxWatcher
	lookup         core.ArbCoreLookup
	challengedNode *core.NodeInfo
	stakerAddress  common.Address
}

func (c *Challenger) ChallengeAddress() common.Address {
	return c.challenge.Address()
}

func NewChallenger(challenge *ethbridge.Challenge, sequencerInbox *ethbridge.SequencerInboxWatcher, lookup core.ArbCoreLookup, challengedNode *core.NodeInfo, stakerAddress common.Address) *Challenger {
	return &Challenger{
		challenge:      challenge,
		sequencerInbox: sequencerInbox,
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

	challengeState, err := c.challenge.ChallengeState(ctx)
	if err != nil {
		return err
	}

	prevBisection, err := c.challenge.LookupBisection(ctx, challengeState)
	if err != nil {
		return err
	}

	if prevBisection == nil {
		prevBisection = c.challengedNode.InitialExecutionBisection()
	}
	challengeImpl := ExecutionImpl{}
	return handleChallenge(ctx, c.challenge, c.sequencerInbox, c.challengedNode.Assertion, c.lookup, challengeImpl, prevBisection)
}

func handleChallenge(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	sequencerInbox *ethbridge.SequencerInboxWatcher,
	assertion *core.Assertion,
	lookup core.ArbCoreLookup,
	challengeImpl ExecutionImpl,
	prevBisection *core.Bisection,
) error {
	logger.Debug().Str("start", prevBisection.ChallengedSegment.Start.String()).Str("end", prevBisection.ChallengedSegment.GetEnd().String()).Msg("Examining opponent's bisection")
	prevCutOffsets := generateBisectionCutOffsets(prevBisection.ChallengedSegment, len(prevBisection.Cuts)-1)
	divergence, err := challengeImpl.FindFirstDivergence(lookup, assertion, prevCutOffsets, prevBisection.Cuts)
	if err != nil {
		return err
	}
	if divergence.DifferentIndex == 0 {
		return errors.New("first cut was already wrong")
	}
	cutToChallenge := divergence.DifferentIndex - 1
	inconsistentSegment := &core.ChallengeSegment{
		Start:  prevCutOffsets[cutToChallenge],
		Length: new(big.Int).Sub(prevCutOffsets[cutToChallenge+1], prevCutOffsets[cutToChallenge]),
	}

	cmp := divergence.SegmentSteps.Cmp(big.NewInt(1))
	if cmp > 0 || divergence.EndIsUnreachable {
		// Steps > 1 or the endpoint is unreachable: Dissect further
		segmentCount := challengeImpl.SegmentTarget()
		if inconsistentSegment.Length.Cmp(big.NewInt(int64(segmentCount))) < 0 {
			// Safe since this is less than 400
			segmentCount = int(inconsistentSegment.Length.Int64())
		}
		subCutOffsets := generateBisectionCutOffsets(inconsistentSegment, segmentCount)
		subCuts, err := challengeImpl.GetCuts(lookup, assertion, subCutOffsets)
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
	} else if cmp < 0 {
		// Steps == 0: Prove that the previous instruction's execution continued through this gas window
		// Also sometimes called a zero step proof, or a constraint win
		// We specifically don't do this when we think the endpoint is unreachable,
		// as we need to dissect unreachable endpoints to force our opponent to fail to prove them
		return challengeImpl.ProveContinuedExecution(
			ctx,
			challenge,
			lookup,
			assertion,
			prevBisection,
			cutToChallenge,
			inconsistentSegment,
		)
	} else {
		// Steps == 1: Do a one step proof, proving the execution of this step specifically
		return challengeImpl.OneStepProof(
			ctx,
			challenge,
			sequencerInbox,
			lookup,
			assertion,
			prevBisection,
			cutToChallenge,
			inconsistentSegment,
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
