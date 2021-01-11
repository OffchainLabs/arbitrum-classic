package ethbridge

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
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
		chunkHash := core.BisectionChunkHash(segmentStart, segmentLength, segmentHashes[i], segmentHashes[i+1])
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
	prevBisection *Bisection,
	segmentToChallenge int,
	subCuts []Cut,
) (*types.Transaction, error) {
	prevCutHashes := cutsToHashes(prevBisection.Cuts)
	subCutHashes := cutsToHashes(subCuts)
	prevTree := NewMerkleTree(calculateBisectionLeaves(prevBisection.ChallengedSegment, prevCutHashes))
	return c.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return c.con.BisectInboxConsistency(
			auth,
			big.NewInt(int64(segmentToChallenge)),
			prevTree.GetProofFlat(segmentToChallenge),
			prevBisection.ChallengedSegment.Start,
			prevBisection.ChallengedSegment.Length,
			prevCutHashes[segmentToChallenge+1],
			subCutHashes,
		)
	})
}

func (c *Challenge) OneStepProveInboxConsistency(
	ctx context.Context,
	prevBisection *Bisection,
	segmentToChallenge int,
	lowerHash [32]byte,
	value [32]byte,
) (*types.Transaction, error) {
	prevCutHashes := cutsToHashes(prevBisection.Cuts)
	prevTree := NewMerkleTree(calculateBisectionLeaves(prevBisection.ChallengedSegment, prevCutHashes))
	return c.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return c.con.OneStepProveInboxConsistency(
			auth,
			big.NewInt(int64(segmentToChallenge)),
			prevTree.GetProofFlat(segmentToChallenge),
			prevBisection.ChallengedSegment.Start,
			prevCutHashes[segmentToChallenge+1],
			lowerHash,
			value,
		)
	})
}

func (c *Challenge) BisectInboxDelta(
	ctx context.Context,
	prevBisection *Bisection,
	segmentToChallenge int,
	subCuts []Cut,
) (*types.Transaction, error) {
	subInboxAccHashes := make([][32]byte, 0, len(subCuts))
	subInboxDeltaHashes := make([][32]byte, 0, len(subCuts))
	for _, cut := range subCuts {
		subInboxAccHashes = append(subInboxAccHashes, cut.(InboxDeltaCut).InboxAccHash)
		subInboxDeltaHashes = append(subInboxDeltaHashes, cut.(InboxDeltaCut).InboxDeltaHash)
	}
	prevCutHashes := cutsToHashes(prevBisection.Cuts)
	prevTree := NewMerkleTree(calculateBisectionLeaves(prevBisection.ChallengedSegment, prevCutHashes))
	return c.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return c.con.BisectInboxDelta(
			auth,
			big.NewInt(int64(segmentToChallenge)),
			prevTree.GetProofFlat(segmentToChallenge),
			prevBisection.ChallengedSegment.Start,
			prevBisection.ChallengedSegment.Length,
			prevBisection.Cuts[segmentToChallenge+1].(InboxDeltaCut).InboxDeltaHash,
			subInboxAccHashes,
			subInboxDeltaHashes,
		)
	})
}

func (c *Challenge) OneStepProveInboxDelta(
	ctx context.Context,
	prevBisection *Bisection,
	segmentToChallenge int,
	msg inbox.InboxMessage,
) (*types.Transaction, error) {
	prevCutHashes := cutsToHashes(prevBisection.Cuts)
	prevTree := NewMerkleTree(calculateBisectionLeaves(prevBisection.ChallengedSegment, prevCutHashes))
	return c.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return c.con.OneStepProveInboxDelta(
			auth,
			big.NewInt(int64(segmentToChallenge)),
			prevTree.GetProofFlat(segmentToChallenge),
			prevBisection.ChallengedSegment.Start,
			prevCutHashes[segmentToChallenge+1],
			prevBisection.Cuts[segmentToChallenge].(InboxDeltaCut).InboxDeltaHash,
			prevBisection.Cuts[segmentToChallenge+1].(InboxDeltaCut).InboxAccHash,
			uint8(msg.Kind),
			msg.ChainTime.BlockNum.AsInt(),
			msg.ChainTime.Timestamp,
			msg.Sender.ToEthAddress(),
			msg.InboxSeqNum,
			msg.Data,
		)
	})
}

func (c *Challenge) BisectExecution(
	ctx context.Context,
	prevBisection *Bisection,
	segmentToChallenge int,
	subCuts []Cut,
) (*types.Transaction, error) {
	prevCutHashes := cutsToHashes(prevBisection.Cuts)
	subCutHashes := cutsToHashes(subCuts)
	prevTree := NewMerkleTree(calculateBisectionLeaves(prevBisection.ChallengedSegment, prevCutHashes))
	return c.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return c.con.BisectExecution(
			auth,
			big.NewInt(int64(segmentToChallenge)),
			prevTree.GetProofFlat(segmentToChallenge),
			prevBisection.ChallengedSegment.Start,
			prevBisection.ChallengedSegment.Length,
			prevCutHashes[segmentToChallenge+1],
			subCutHashes,
			subCuts[0].(ExpandedExecutionCut).GasUsed,
			subCuts[0].(ExpandedExecutionCut).Rest,
		)
	})
}

func (c *Challenge) OneStepProveExecution(
	ctx context.Context,
	prevBisection *Bisection,
	segmentToChallenge int,
	beforeAssertion *core.AssertionInfo,
	executionProof []byte,
	bufferProof []byte,
) (*types.Transaction, error) {
	prevCutHashes := cutsToHashes(prevBisection.Cuts)
	prevTree := NewMerkleTree(calculateBisectionLeaves(prevBisection.ChallengedSegment, prevCutHashes))
	return c.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return c.con.OneStepProveExecution(
			auth,
			big.NewInt(int64(segmentToChallenge)),
			prevTree.GetProofFlat(segmentToChallenge),
			prevBisection.ChallengedSegment.Start,
			prevCutHashes[segmentToChallenge+1],
			[3][32]byte{
				beforeAssertion.InboxDelta,
				beforeAssertion.ExecInfo.SendAcc,
				beforeAssertion.ExecInfo.LogAcc,
			},
			beforeAssertion.ExecInfo.GasUsed.Uint64(),
			beforeAssertion.ExecInfo.SendCount,
			beforeAssertion.ExecInfo.LogCount,
			executionProof,
			bufferProof,
		)
	})
}

func cutsToHashes(cuts []Cut) [][32]byte {
	cutHashes := make([][32]byte, 0, len(cuts))
	for _, cut := range cuts {
		cutHashes = append(cutHashes, cut.Hash())
	}
	return cutHashes
}
