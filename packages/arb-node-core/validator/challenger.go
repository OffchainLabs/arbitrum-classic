package validator

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/pkg/errors"
	"math/big"
)

func (s *Staker) handleConflict(ctx context.Context, info *ethbridge.StakerInfo) (*types.Transaction, error) {
	if info.CurrentChallenge == nil {
		return nil, nil
	}
	challenge, err := ethbridge.NewChallengeWatcher(info.CurrentChallenge.ToEthAddress(), s.client)
	if err != nil {
		return nil, err
	}
	responder, err := challenge.CurrentResponder(ctx)
	if err != nil {
		return nil, err
	}
	if responder != s.address {
		// Not our turn
		return nil, nil
	}
	kind, err := challenge.Kind(ctx)
	if err != nil {
		return nil, err
	}

	challengedNodeNum, err := challenge.ChallengedNodeNum(ctx)
	if err != nil {
		return nil, err
	}

	nodeInfo, err := s.lookupNode(ctx, challengedNodeNum)
	if err != nil {
		return nil, err
	}

	switch kind {
	case ethbridge.UNINITIALIZED:
		judgment, err := s.judgeNode(nodeInfo, nil)
		if err != nil {
			return nil, err
		}
		switch judgment {
		case ethbridge.INBOX_CONSISTENCY:
			return s.handleInboxConsistencyChallenge()
		case ethbridge.INBOX_DELTA:
			return s.handleInboxDeltaChallenge()
		case ethbridge.EXECUTION:
			return s.handleExecutionChallenge()
		case ethbridge.STOPPED_SHORT:
			return s.handleStoppedShortChallenge()
		default:
			return nil, errors.New("can't handle challenge")
		}
	case ethbridge.INBOX_CONSISTENCY:
		return s.handleInboxConsistencyChallenge()
	case ethbridge.INBOX_DELTA:
		return s.handleInboxDeltaChallenge()
	case ethbridge.EXECUTION:
		return s.handleExecutionChallenge()
	case ethbridge.STOPPED_SHORT:
		return s.handleStoppedShortChallenge()
	default:
		return nil, errors.New("can't handle challenge")
	}
}

func (s *Staker) handleInboxConsistencyChallenge(ctx context.Context, challenge *ethbridge.Challenge, info *ethbridgecontracts.ChallengeBisected) (*types.Transaction, error) {
	prevSegment := &ethbridge.ChallengeSegment{
		Start:  info.SegmentStart,
		Length: info.SegmentLength,
	}
	inconsistentSegment, segmentToChallenge, err := s.findInboxInconsistency(prevSegment, info.ChainHashes)
	if err != nil {
		return nil, err
	}

	if inconsistentSegment.Length.Cmp(big.NewInt(1)) == 0 {
		beforeInboxAcc, err := s.lookup.GetInboxAcc(inconsistentSegment.Start)
		if err != nil {
			return nil, err
		}
		msgs, err := s.lookup.GetMessages(inconsistentSegment.Start, big.NewInt(1))
		if err != nil {
			return nil, err
		}
		return challenge.OneStepProveInboxConsistency(ctx, prevSegment, info.ChainHashes, segmentToChallenge, beforeInboxAcc, msgs[0].CommitmentHash())
	} else {
		subSegments, err := s.generateInboxBisection(inconsistentSegment)
		if err != nil {
			return nil, err
		}
		return challenge.BisectInboxConsistency(ctx, prevSegment, info.ChainHashes, segmentToChallenge, subSegments)
	}
}

func (s *Staker) findInboxInconsistency(segment *ethbridge.ChallengeSegment, chunkHashes [][32]byte) (*ethbridge.ChallengeSegment, int, error) {
	inboxOffset := segment.Start
	for i, chunkHash := range chunkHashes {
		segmentLength := calculateBisectionChunkCount(i, len(chunkHashes), segment.Length)
		newInboxOffset := new(big.Int).Add(inboxOffset, segmentLength)
		inboxAcc, err := s.lookup.GetInboxAcc(newInboxOffset)
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

func (s *Staker) generateInboxBisection(segment *ethbridge.ChallengeSegment) ([][32]byte, error) {
	segmentCount := calculateBisectionSegmentCount(segment.Length)
	segments := make([][32]byte, 0, segmentCount+1)

	offset := new(big.Int).Set(segment.Start)
	inboxAcc, err := s.lookup.GetInboxAcc(offset)
	if err != nil {
		return nil, err
	}
	segments = append(segments, inboxAcc)
	for i := 0; i < segmentCount; i++ {
		subSegmentLength := calculateBisectionChunkCount(i, segmentCount, segment.Length)
		offset = offset.Add(offset, subSegmentLength)
		inboxAcc, err := s.lookup.GetInboxAcc(offset)
		if err != nil {
			return nil, err
		}
		segments = append(segments, inboxAcc)
	}
	return segments, nil
}

func (s *Staker) handleInboxDeltaChallenge() (*types.Transaction, error) {
	return nil, nil
}

func (s *Staker) handleExecutionChallenge() (*types.Transaction, error) {
	return nil, nil
}

func (s *Staker) handleStoppedShortChallenge() (*types.Transaction, error) {
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
