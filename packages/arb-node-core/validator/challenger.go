package validator

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/pkg/errors"
	"math/big"
)

type Challenger struct {
	challenge      *ethbridge.Challenge
	lookup         core.ValidatorLookup
	challengedNode *ethbridge.NodeInfo
}

func NewChallenger(challenge *ethbridge.Challenge, lookup core.ValidatorLookup, challengedNode *ethbridge.NodeInfo) *Challenger {
	return &Challenger{
		challenge:      challenge,
		lookup:         lookup,
		challengedNode: challengedNode,
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
	case ethbridge.UNINITIALIZED:
		judgment, err := judgeNode(c.lookup, c.challengedNode, nil)
		if err != nil {
			return nil, err
		}
		switch judgment {
		case ethbridge.INBOX_CONSISTENCY:
			return c.handleInboxConsistencyChallenge(ctx)
		case ethbridge.INBOX_DELTA:
			return c.handleInboxDeltaChallenge()
		case ethbridge.EXECUTION:
			return c.handleExecutionChallenge()
		case ethbridge.STOPPED_SHORT:
			return c.handleStoppedShortChallenge()
		default:
			return nil, errors.New("can't handle challenge")
		}
	case ethbridge.INBOX_CONSISTENCY:
		return c.handleInboxConsistencyChallenge(ctx)
	case ethbridge.INBOX_DELTA:
		return c.handleInboxDeltaChallenge()
	case ethbridge.EXECUTION:
		return c.handleExecutionChallenge()
	case ethbridge.STOPPED_SHORT:
		return c.handleStoppedShortChallenge()
	default:
		return nil, errors.New("can't handle challenge")
	}

}

func (c *Challenger) handleInboxConsistencyChallenge(ctx context.Context) (*types.Transaction, error) {
	challengeState, err := c.challenge.ChallengeState(ctx)
	if err != nil {
		return nil, err
	}

	bisection, err := c.challenge.LookupBisection(ctx, challengeState)
	if err != nil {
		return nil, err
	}

	inconsistentSegment, segmentToChallenge, err := c.findInboxInconsistency(bisection.PrevSegment, bisection.ChainHashes)
	if err != nil {
		return nil, err
	}

	if inconsistentSegment.Length.Cmp(big.NewInt(1)) == 0 {
		beforeInboxAcc, err := c.lookup.GetInboxAcc(inconsistentSegment.Start)
		if err != nil {
			return nil, err
		}
		msgs, err := c.lookup.GetMessages(inconsistentSegment.Start, big.NewInt(1))
		if err != nil {
			return nil, err
		}
		return c.challenge.OneStepProveInboxConsistency(ctx, bisection.ChainHashes, segmentToChallenge, bisection.PrevSegment, beforeInboxAcc, msgs[0].CommitmentHash())
	} else {
		subSegments, err := c.generateInboxBisection(inconsistentSegment)
		if err != nil {
			return nil, err
		}
		return c.challenge.BisectInboxConsistency(ctx, bisection.ChainHashes, segmentToChallenge, bisection.PrevSegment, subSegments)
	}
}

func (c *Challenger) findInboxInconsistency(segment *ethbridge.ChallengeSegment, chunkHashes [][32]byte) (*ethbridge.ChallengeSegment, int, error) {
	inboxOffset := segment.Start
	for i, chunkHash := range chunkHashes {
		segmentLength := calculateBisectionChunkCount(i, len(chunkHashes), segment.Length)
		newInboxOffset := new(big.Int).Add(inboxOffset, segmentLength)
		inboxAcc, err := c.lookup.GetInboxAcc(newInboxOffset)
		if err != nil {
			return nil, 0, err
		}
		if inboxAcc != chunkHash {
			if i == 0 {
				return nil, 0, errors.New("first segment was already wrong")
			}
			return &ethbridge.ChallengeSegment{
				Start:  inboxOffset,
				Length: segmentLength,
			}, i, nil
		}
		inboxOffset = newInboxOffset
	}
	return nil, 0, nil
}

func (c *Challenger) generateInboxBisection(segment *ethbridge.ChallengeSegment) ([][32]byte, error) {
	segmentCount := calculateBisectionSegmentCount(segment.Length)
	segments := make([][32]byte, 0, segmentCount+1)

	offset := new(big.Int).Set(segment.Start)
	inboxAcc, err := c.lookup.GetInboxAcc(offset)
	if err != nil {
		return nil, err
	}
	segments = append(segments, inboxAcc)
	for i := 0; i < segmentCount; i++ {
		subSegmentLength := calculateBisectionChunkCount(i, segmentCount, segment.Length)
		offset = offset.Add(offset, subSegmentLength)
		inboxAcc, err := c.lookup.GetInboxAcc(offset)
		if err != nil {
			return nil, err
		}
		segments = append(segments, inboxAcc)
	}
	return segments, nil
}

func (c *Challenger) handleInboxDeltaChallenge() (*types.Transaction, error) {
	return nil, nil
}

func (c *Challenger) handleExecutionChallenge() (*types.Transaction, error) {
	return nil, nil
}

func (c *Challenger) handleStoppedShortChallenge() (*types.Transaction, error) {
	return nil, nil
}

func calculateBisectionChunkCount(segmentIndex, segmentCount int, totalLength *big.Int) *big.Int {
	size := new(big.Int).Div(totalLength, big.NewInt(int64(segmentCount)))
	if segmentIndex == 0 {
		size = size.Add(size, new(big.Int).Mod(totalLength, big.NewInt(int64(segmentCount))))
	}
	return size
}

func calculateBisectionSegmentCount(totalLength *big.Int) int {
	maxSegmentCount := 20
	if totalLength.Cmp(big.NewInt(int64(maxSegmentCount))) < 0 {
		// Safe since this is less than 20
		return int(totalLength.Int64())
	} else {
		return maxSegmentCount
	}
}
