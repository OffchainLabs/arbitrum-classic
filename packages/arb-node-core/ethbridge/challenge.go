package ethbridge

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
)

type ChallengeSegment struct {
	Start  *big.Int
	Length *big.Int
}

func calculateBisectionChunkCount(segmentIndex, segmentCount int, totalLength *big.Int) *big.Int {
	size := new(big.Int).Div(totalLength, big.NewInt(int64(segmentCount)))
	if segmentIndex == 0 {
		size = size.Add(size, new(big.Int).Mod(totalLength, big.NewInt(int64(segmentCount))))
	}
	return size
}

func calculateBisectionLeaves(segment *ChallengeSegment, segmentHashes [][32]byte) [][32]byte {
	chunks := make([][32]byte, 0, len(segmentHashes))
	segmentStart := new(big.Int).Set(segment.Start)
	for i := 0; i < len(segmentHashes)-1; i++ {
		segmentLength := calculateBisectionChunkCount(i, len(segmentHashes), segment.Length)
		chunkHash := BisectionChunkHash(segmentStart, segmentLength, segmentHashes[i], segmentHashes[i+1])
		chunks = append(chunks, chunkHash)
		segmentStart = segmentStart.Add(segmentStart, segmentLength)
	}
	return chunks
}

type Challenge struct {
	*ChallengeWatcher
	auth *TransactAuth
}

func NewChallenge(address ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*Challenge, error) {
	watcher, err := NewChallengeWatcher(address, client)
	if err != nil {
		return nil, err
	}
	return &Challenge{
		ChallengeWatcher: watcher,
		auth:             auth,
	}, nil
}

func (c *Challenge) Transactor() common.Address {
	return common.NewAddressFromEth(c.auth.auth.From)
}

func (c *Challenge) BisectInboxConsistency(
	ctx context.Context,
	prevChainHashes [][32]byte,
	segmentToChallenge int,
	prevSegment *ChallengeSegment,
	subSegments [][32]byte,
) (*types.Transaction, error) {
	prevTree := NewMerkleTree(calculateBisectionLeaves(prevSegment, prevChainHashes))
	return c.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return c.con.BisectInboxConsistency(
			auth,
			big.NewInt(int64(segmentToChallenge)),
			prevTree.GetProofFlat(segmentToChallenge),
			prevSegment.Start,
			prevSegment.Length,
			prevChainHashes[segmentToChallenge+1],
			subSegments,
		)
	})
}

func (c *Challenge) OneStepProveInboxConsistency(
	ctx context.Context,
	prevChainHashes [][32]byte,
	segmentToChallenge int,
	prevSegment *ChallengeSegment,
	lowerHash [32]byte,
	value [32]byte,
) (*types.Transaction, error) {
	prevTree := NewMerkleTree(calculateBisectionLeaves(prevSegment, prevChainHashes))
	return c.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return c.con.OneStepProveInboxConsistency(
			auth,
			big.NewInt(int64(segmentToChallenge)),
			prevTree.GetProofFlat(segmentToChallenge),
			prevSegment.Start,
			prevChainHashes[segmentToChallenge+1],
			lowerHash,
			value,
		)
	})
}

//func (c *Challenge) BisectInboxConsistency(
//	ctx context.Context,
//	prevSegment *ChallengeSegment,
//	prevChainHashes [][32]byte,
//	segmentToChallenge int,
//	subSegments [][32]byte,
//) (*types.Transaction, error) {
//	prevTree := NewMerkleTree(calculateBisectionLeaves(prevSegment, prevChainHashes))
//	return c.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
//		return c.con.BisectInboxDelta(
//			auth,
//			big.NewInt(int64(segmentToChallenge)),
//			prevTree.GetProofFlat(segmentToChallenge),
//			subSegments,
//			prevSegment.Start,
//			prevSegment.Length,
//			prevChainHashes[segmentToChallenge+1],
//		)
//	})
//}
