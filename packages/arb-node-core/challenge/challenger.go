package challenge

import (
	"context"
	"fmt"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/pkg/errors"
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

	challengeState, err := c.challenge.ChallengeState(ctx)
	if err != nil {
		return err
	}

	prevBisection, err := c.challenge.LookupBisection(ctx, challengeState)
	if err != nil {
		return err
	}
	fmt.Printf("Lookup bisection %v\n", prevBisection)

	challengeImpl := ExecutionImpl{}

	segment, err := c.challenge.LookupContinue(ctx, challengeState)
	if err != nil {
		return err
	}
	if segment != nil {
		fmt.Printf("Lookup subobligation start %v\n", segment)
		opcode, machine, err := challengeImpl.OneStepProofMachine(
			ctx,
			c.challenge,
			c.lookup,
			c.challengedNode.Assertion,
			segment,
		)
		if err != nil {
			return err
		}
		if opcode == 241 {
			// Get new lookup
			fmt.Printf("Found wasm test, making new lookup\n")
			storage, err := cmachine.NewArbStorage("/tmp/arbStorage2")
			fmt.Printf("Found wasm test, making new lookup ??? %v\n", err)
			storage.InitializeForWasm((machine).(cmachine.ExtendedMachine))
			arbCore := storage.GetArbCore()
			arbCore.StartThread()
			c.lookup = c.lookup.SubLookup(arbCore)
		}
	}

	if prevBisection == nil {
		prevBisection = c.challengedNode.InitialExecutionBisection()
	}
	return c.handleChallenge(ctx, c.challenge, c.challengedNode.Assertion, c.lookup, challengeImpl, prevBisection)
}

func (c *Challenger) handleChallenge(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	assertion *core.Assertion,
	lookup core.ArbCoreLookup,
	challengeImpl ExecutionImpl,
	prevBisection *core.Bisection,
) error {
	prevCutOffsets := generateBisectionCutOffsets(prevBisection.ChallengedSegment, len(prevBisection.Cuts)-1)
	fmt.Printf("hmm what %v %v prev cuts %v\n", prevBisection.ChallengedSegment, len(prevBisection.Cuts)-1, prevCutOffsets)
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
		opcode, machine, err := challengeImpl.OneStepProof(
			ctx,
			challenge,
			lookup,
			assertion,
			prevBisection,
			cutToChallenge,
			inconsistentSegment,
		)
		if opcode == 241 {
			// Get new lookup
			fmt.Printf("Found wasm test, making new lookup\n")
			storage, err := cmachine.NewArbStorage("/tmp/arbStorage")
			storage.InitializeForWasm((machine).(cmachine.ExtendedMachine))
			arbCore := storage.GetArbCore()
			arbCore.StartThread()
			// c.lookup = arbCore
			c.lookup = c.lookup.SubLookup(arbCore)
			return err
		}
		return err
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
