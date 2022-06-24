package challenge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arblog"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/pkg/errors"
)

var logger = arblog.Logger.With().Str("component", "challenge").Logger()

type Challenger struct {
	challenge           *ethbridge.Challenge
	sequencerInbox      *ethbridge.SequencerInboxWatcher
	lookup              core.ArbCoreLookup
	challengedAssertion *core.Assertion
	stakerAddress       common.Address
}

func (c *Challenger) ChallengeAddress() common.Address {
	return c.challenge.Address()
}

func NewChallenger(challenge *ethbridge.Challenge, sequencerInbox *ethbridge.SequencerInboxWatcher, lookup core.ArbCoreLookup, challengedAssertion *core.Assertion, stakerAddress common.Address) *Challenger {
	return &Challenger{
		challenge:           challenge,
		sequencerInbox:      sequencerInbox,
		lookup:              lookup,
		challengedAssertion: challengedAssertion,
		stakerAddress:       stakerAddress,
	}
}

func (c *Challenger) HandleConflict(ctx context.Context) (Move, error) {
	isTimedOut, err := c.challenge.IsTimedOut(ctx)
	if err != nil {
		return nil, err
	}
	if isTimedOut {
		move := &TimeoutMove{}
		return move, move.execute(ctx, c.challenge)
	}

	responder, err := c.challenge.CurrentResponder(ctx)
	if err != nil {
		return nil, err
	}
	if responder != c.stakerAddress {
		// Not our turn
		return nil, nil
	}

	challengeState, err := c.challenge.ChallengeState(ctx)
	if err != nil {
		return nil, err
	}

	emptyHash := common.Hash{}
	if challengeState == emptyHash {
		logger.Warn().Str("contract", c.challenge.Address().Hex()).Msg("challenge has been lost, waiting for timeout")
		return nil, nil
	}

	prevBisection, err := c.challenge.LookupBisection(ctx, challengeState)
	if err != nil {
		return nil, err
	}

	if prevBisection == nil {
		prevBisection = c.challengedAssertion.InitialExecutionBisection()
	}
	move, err := handleChallenge(ctx, c.challengedAssertion, c.lookup, c.sequencerInbox, prevBisection)
	if err != nil {
		return nil, err
	}
	return move, move.execute(ctx, c.challenge)
}

func handleChallenge(
	ctx context.Context,
	assertion *core.Assertion,
	lookup core.ArbCoreLookup,
	sequencerInbox *ethbridge.SequencerInboxWatcher,
	prevBisection *core.Bisection,
) (Move, error) {
	logger.Debug().Str("start", prevBisection.ChallengedSegment.Start.String()).Str("end", prevBisection.ChallengedSegment.GetEnd().String()).Msg("Examining opponent's bisection")
	prevCutOffsets := generateBisectionCutOffsets(prevBisection.ChallengedSegment, len(prevBisection.Cuts)-1)
	divergence, err := FindFirstDivergence(lookup, assertion, prevCutOffsets, prevBisection.Cuts)
	if err != nil {
		return nil, err
	}
	if divergence.DifferentIndex == 0 {
		return nil, errors.New("first cut was already wrong")
	}
	cutToChallenge := divergence.DifferentIndex - 1
	inconsistentSegment := &core.ChallengeSegment{
		Start:  prevCutOffsets[cutToChallenge],
		Length: new(big.Int).Sub(prevCutOffsets[cutToChallenge+1], prevCutOffsets[cutToChallenge]),
	}

	cmp := divergence.SegmentSteps.Cmp(big.NewInt(1))
	if cmp > 0 || divergence.EndIsUnreachable {
		// Steps > 1 or the endpoint is unreachable: Dissect further
		segmentCount := SegmentTarget()
		if inconsistentSegment.Length.Cmp(big.NewInt(int64(segmentCount))) < 0 {
			// Safe since this is less than 400
			segmentCount = int(inconsistentSegment.Length.Int64())
		}
		subCutOffsets := generateBisectionCutOffsets(inconsistentSegment, segmentCount)
		startState, subCuts, err := GetCuts(lookup, assertion, subCutOffsets)
		if err != nil {
			return nil, err
		}
		return &BisectMove{
			prevBisection:       prevBisection,
			startState:          startState,
			segmentToChallenge:  cutToChallenge,
			inconsistentSegment: inconsistentSegment,
			subCuts:             subCuts,
		}, nil
	} else if cmp < 0 {
		// Steps == 0: Prove that the previous instruction's execution continued through this gas window
		// Also sometimes called a zero step proof, or a constraint win
		// We specifically don't do this when we think the endpoint is unreachable,
		// as we need to dissect unreachable endpoints to force our opponent to fail to prove them
		previousCut, _, err := getSegmentStartInfo(lookup, assertion, inconsistentSegment)
		if err != nil {
			return nil, err
		}
		return &ProveContinuedMove{
			assertion:          assertion,
			prevBisection:      prevBisection,
			segmentToChallenge: cutToChallenge,
			challengedSegment:  inconsistentSegment,
			previousCut:        previousCut,
		}, nil
	} else {
		// Steps == 1: Do a one step proof, proving the execution of this step specifically
		return NewOneStepProofMove(
			ctx,
			assertion,
			prevBisection,
			cutToChallenge,
			inconsistentSegment,
			sequencerInbox,
			lookup,
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
