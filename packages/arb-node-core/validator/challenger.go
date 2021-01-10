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
	challengeImpl := &InboxConsistencyImpl{
		lookup: c.lookup,
	}

	return handleChallenge(ctx, c.challenge, challengeImpl)
}

func (c *Challenger) handleInboxDeltaChallenge(ctx context.Context) (*types.Transaction, error) {
	inboxDeltaData, err := c.InboxDelta()
	if err != nil {
		return nil, err
	}
	challengeImpl := &InboxDeltaImpl{
		lookup:              c.lookup,
		nodeAfterInboxCount: c.challengedNode.Assertion.AfterInboxCount(),
		inboxDelta:          inboxDeltaData,
	}
	return handleChallenge(ctx, c.challenge, challengeImpl)
}

func (c *Challenger) handleExecutionChallenge() (*types.Transaction, error) {
	return nil, nil
}

func (c *Challenger) handleStoppedShortChallenge() (*types.Transaction, error) {
	return nil, nil
}

type ChallengerImpl interface {
	GetCut(offset *big.Int) (ethbridge.Cut, error)

	Bisect(
		ctx context.Context,
		challenge *ethbridge.Challenge,
		prevBisection *ethbridge.Bisection,
		segmentToChallenge int,
		subCuts []ethbridge.Cut,
	) (*types.Transaction, error)

	OneStepProof(
		ctx context.Context,
		challenge *ethbridge.Challenge,
		prevBisection *ethbridge.Bisection,
		segmentToChallenge int,
		challengedSegment *ethbridge.ChallengeSegment,
	) (*types.Transaction, error)
}

type InboxConsistencyImpl struct {
	lookup core.ValidatorLookup
}

func (i *InboxConsistencyImpl) GetCut(offset *big.Int) (ethbridge.Cut, error) {
	inboxAcc, err := i.lookup.GetInboxAcc(offset)
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
	prevBisection *ethbridge.Bisection,
	segmentToChallenge int,
	challengedSegment *ethbridge.ChallengeSegment,
) (*types.Transaction, error) {
	beforeInboxAcc, err := i.lookup.GetInboxAcc(challengedSegment.Start)
	if err != nil {
		return nil, err
	}
	msgs, err := i.lookup.GetMessages(challengedSegment.Start, big.NewInt(1))
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

type InboxDeltaImpl struct {
	lookup              core.ValidatorLookup
	nodeAfterInboxCount *big.Int
	inboxDelta          *inboxDelta
}

func (i *InboxDeltaImpl) GetCut(offset *big.Int) (ethbridge.Cut, error) {
	inboxOffset := new(big.Int).Add(i.nodeAfterInboxCount, offset)
	inboxAcc, err := i.lookup.GetInboxAcc(inboxOffset)
	if err != nil {
		return nil, err
	}
	return ethbridge.InboxDeltaCut{
		InboxAccHash:   inboxAcc,
		InboxDeltaHash: i.inboxDelta.inboxDeltaAccs[offset.Uint64()],
	}, nil
}

func (i *InboxDeltaImpl) Bisect(
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

func (i *InboxDeltaImpl) OneStepProof(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	prevBisection *ethbridge.Bisection,
	segmentToChallenge int,
	challengedSegment *ethbridge.ChallengeSegment,
) (*types.Transaction, error) {
	msgIndex := new(big.Int).Add(i.nodeAfterInboxCount, challengedSegment.Start)
	msgs, err := i.lookup.GetMessages(msgIndex, big.NewInt(1))
	if err != nil {
		return nil, err
	}
	return challenge.OneStepProveInboxDelta(
		ctx,
		prevBisection.Cuts,
		segmentToChallenge,
		prevBisection.ChallengedSegment,
		msgs[0],
	)
}

func handleChallenge(ctx context.Context, challenge *ethbridge.Challenge, challengeImpl ChallengerImpl) (*types.Transaction, error) {
	challengeState, err := challenge.ChallengeState(ctx)
	if err != nil {
		return nil, err
	}

	prevBisection, err := challenge.LookupBisection(ctx, challengeState)
	if err != nil {
		return nil, err
	}

	inconsistentSegment, segmentToChallenge, err := findIncorrectSegment(prevBisection, challengeImpl)
	if err != nil {
		return nil, err
	}

	if inconsistentSegment.Length.Cmp(big.NewInt(1)) == 0 {
		return challengeImpl.OneStepProof(
			ctx,
			challenge,
			prevBisection,
			segmentToChallenge,
			inconsistentSegment,
		)
	} else {
		subCuts, err := generateBisection(inconsistentSegment, challengeImpl)
		if err != nil {
			return nil, err
		}
		return challengeImpl.Bisect(
			ctx,
			challenge,
			prevBisection,
			segmentToChallenge,
			subCuts,
		)
	}
}

func generateBisection(segment *ethbridge.ChallengeSegment, impl ChallengerImpl) ([]ethbridge.Cut, error) {
	segmentCount := 20
	if segment.Length.Cmp(big.NewInt(int64(segmentCount))) < 0 {
		// Safe since this is less than 20
		segmentCount = int(segment.Length.Int64())
	}
	cuts := make([]ethbridge.Cut, 0, segmentCount+1)
	offset := new(big.Int).Set(segment.Start)
	for i := 0; i < segmentCount+1; i++ {
		cut, err := impl.GetCut(offset)
		if err != nil {
			return nil, err
		}
		cuts = append(cuts, cut)
		subSegmentLength := calculateBisectionChunkCount(i, segmentCount, segment.Length)
		offset = offset.Add(offset, subSegmentLength)
	}
	return cuts, nil
}

func findIncorrectSegment(prevBisection *ethbridge.Bisection, impl ChallengerImpl) (*ethbridge.ChallengeSegment, int, error) {
	offset := prevBisection.ChallengedSegment.Start
	for i, cut := range prevBisection.Cuts {
		correctCut, err := impl.GetCut(offset)
		if err != nil {
			return nil, 0, err
		}
		segmentLength := calculateBisectionChunkCount(i, len(prevBisection.Cuts), prevBisection.ChallengedSegment.Length)
		if !correctCut.Equals(cut) {
			if i == 0 {
				return nil, 0, errors.New("first segment was already wrong")
			}
			return &ethbridge.ChallengeSegment{
				Start:  offset,
				Length: segmentLength,
			}, i, nil
		}
		offset = new(big.Int).Add(offset, segmentLength)
	}
	return nil, 0, nil
}

func calculateBisectionChunkCount(segmentIndex, segmentCount int, totalLength *big.Int) *big.Int {
	size := new(big.Int).Div(totalLength, big.NewInt(int64(segmentCount)))
	if segmentIndex == 0 {
		size = size.Add(size, new(big.Int).Mod(totalLength, big.NewInt(int64(segmentCount))))
	}
	return size
}
