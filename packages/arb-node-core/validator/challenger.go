package validator

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
	challenge      *ethbridge.Challenge
	lookup         core.ValidatorLookup
	challengedNode *ethbridge.NodeInfo

	inboxDelta *inboxDelta
}

func (c *Challenger) InboxDelta() (*inboxDelta, error) {
	if c.inboxDelta != nil {
		messages, err := c.lookup.GetMessages(
			c.challengedNode.Assertion.PrevState.InboxCount,
			c.challengedNode.Assertion.ExecInfo.InboxMessagesRead,
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
			return c.handleInboxDeltaChallenge(ctx)
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
		return c.handleInboxDeltaChallenge(ctx)
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

	inconsistentSegment, segmentToChallenge, err := c.findInboxInconsistency(bisection)
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
		return c.challenge.OneStepProveInboxConsistency(
			ctx,
			bisection.ChainHashes,
			segmentToChallenge,
			bisection.ChallengedSegment,
			beforeInboxAcc,
			msgs[0].CommitmentHash(),
		)
	} else {
		subSegments, err := c.generateInboxConsistencyBisection(inconsistentSegment)
		if err != nil {
			return nil, err
		}
		return c.challenge.BisectInboxConsistency(
			ctx,
			bisection.ChainHashes,
			segmentToChallenge,
			bisection.ChallengedSegment,
			subSegments,
		)
	}
}

func (c *Challenger) handleInboxDeltaChallenge(ctx context.Context) (*types.Transaction, error) {
	challengeState, err := c.challenge.ChallengeState(ctx)
	if err != nil {
		return nil, err
	}

	bisection, err := c.challenge.LookupInboxDeltaBisection(ctx, challengeState)
	if err != nil {
		return nil, err
	}

	inconsistentSegment, segmentToChallenge, err := c.findInboxDelta(bisection)
	if err != nil {
		return nil, err
	}

	if inconsistentSegment.Length.Cmp(big.NewInt(1)) == 0 {
		msgIndex := new(big.Int).Add(
			c.challengedNode.Assertion.PrevState.InboxCount,
			inconsistentSegment.Start,
		)
		msgs, err := c.lookup.GetMessages(msgIndex, big.NewInt(1))
		if err != nil {
			return nil, err
		}
		return c.challenge.OneStepProveInboxDelta(
			ctx,
			bisection.InboxAccHashes,
			bisection.InboxDeltaHashes,
			segmentToChallenge,
			bisection.ChallengedSegment,
			msgs[0],
		)
	} else {
		inboxAccHashes, inboxDeltaHashes, err := c.generateInboxDeltaBisection(inconsistentSegment)
		if err != nil {
			return nil, err
		}
		return c.challenge.BisectInboxDelta(
			ctx,
			bisection.InboxAccHashes,
			bisection.InboxDeltaHashes,
			segmentToChallenge,
			bisection.ChallengedSegment,
			inboxAccHashes,
			inboxDeltaHashes,
		)
	}
}

func (c *Challenger) handleExecutionChallenge() (*types.Transaction, error) {
	return nil, nil
}

func (c *Challenger) handleStoppedShortChallenge() (*types.Transaction, error) {
	return nil, nil
}

func (c *Challenger) findInboxInconsistency(prevBisection *ethbridge.Bisection) (*ethbridge.ChallengeSegment, int, error) {
	inboxOffset := prevBisection.ChallengedSegment.Start
	for i, chunkHash := range prevBisection.ChainHashes {
		segmentLength := calculateBisectionChunkCount(i, len(prevBisection.ChainHashes), prevBisection.ChallengedSegment.Length)

		inboxAcc, err := c.lookup.GetInboxAcc(inboxOffset)
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
		inboxOffset = new(big.Int).Add(inboxOffset, segmentLength)
	}
	return nil, 0, nil
}

func (c *Challenger) generateInboxConsistencyBisection(segment *ethbridge.ChallengeSegment) ([][32]byte, error) {
	segmentCount := calculateBisectionSegmentCount(segment.Length)
	segments := make([][32]byte, 0, segmentCount+1)

	offset := new(big.Int).Set(segment.Start)
	for i := 0; i < segmentCount+1; i++ {
		inboxAcc, err := c.lookup.GetInboxAcc(offset)
		if err != nil {
			return nil, err
		}
		segments = append(segments, inboxAcc)
		subSegmentLength := calculateBisectionChunkCount(i, segmentCount, segment.Length)
		offset = offset.Add(offset, subSegmentLength)
	}
	return segments, nil
}

func (c *Challenger) findInboxDelta(prevBisection *ethbridge.InboxDeltaBisection) (*ethbridge.ChallengeSegment, int, error) {
	correctInboxDelta, err := c.InboxDelta()
	if err != nil {
		return nil, 0, err
	}
	offset := prevBisection.ChallengedSegment.Start
	for i, inboxAccHash := range prevBisection.InboxAccHashes {
		inboxDeltaHash := prevBisection.InboxDeltaHashes[i]
		segmentLength := calculateBisectionChunkCount(i, len(prevBisection.InboxAccHashes), prevBisection.ChallengedSegment.Length)

		inboxOffset := new(big.Int).Add(c.challengedNode.Assertion.AfterInboxCount(), offset)
		inboxAcc, err := c.lookup.GetInboxAcc(inboxOffset)
		if err != nil {
			return nil, 0, err
		}
		if inboxAcc != inboxAccHash || correctInboxDelta.inboxDeltaAccs[offset.Uint64()] != inboxDeltaHash {
			if i == 0 {
				return nil, 0, errors.New("first segment was already wrong")
			}
			return &ethbridge.ChallengeSegment{
				Start:  inboxOffset,
				Length: segmentLength,
			}, i, nil
		}
		offset = new(big.Int).Add(offset, segmentLength)
	}
	return nil, 0, nil
}

func (c *Challenger) generateInboxDeltaBisection(segment *ethbridge.ChallengeSegment) ([][32]byte, [][32]byte, error) {
	correctInboxDelta, err := c.InboxDelta()
	if err != nil {
		return nil, nil, err
	}

	segmentCount := calculateBisectionSegmentCount(segment.Length)
	inboxAccHashes := make([][32]byte, 0, segmentCount+1)
	inboxDeltaHashes := make([][32]byte, 0, segmentCount+1)

	offset := new(big.Int).Set(segment.Start)
	for i := 0; i < segmentCount+1; i++ {
		inboxOffset := new(big.Int).Add(
			c.challengedNode.Assertion.PrevState.InboxCount,
			segment.Start,
		)
		inboxAcc, err := c.lookup.GetInboxAcc(inboxOffset)
		if err != nil {
			return nil, nil, err
		}
		inboxAccHashes = append(inboxAccHashes, inboxAcc)
		inboxDeltaHashes = append(inboxDeltaHashes, correctInboxDelta.inboxDeltaAccs[offset.Uint64()])
		subSegmentLength := calculateBisectionChunkCount(i, segmentCount, segment.Length)
		offset = offset.Add(offset, subSegmentLength)
	}
	return inboxAccHashes, inboxDeltaHashes, nil
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
