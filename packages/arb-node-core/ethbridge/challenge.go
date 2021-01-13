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

func calculateBisectionChunkCount(segmentIndex, segmentCount int, totalLength *big.Int) *big.Int {
	size := new(big.Int).Div(totalLength, big.NewInt(int64(segmentCount)))
	if segmentIndex == 0 {
		size = size.Add(size, new(big.Int).Mod(totalLength, big.NewInt(int64(segmentCount))))
	}
	return size
}

func calculateBisectionTree(bisection *core.Bisection) ([][32]byte, *MerkleTree) {
	cutHashes := cutsToHashes(bisection.Cuts)
	segmentCount := len(cutHashes) - 1
	chunks := make([][32]byte, 0, segmentCount)
	segmentStart := new(big.Int).Set(bisection.ChallengedSegment.Start)
	for i := 0; i < segmentCount; i++ {
		segmentLength := calculateBisectionChunkCount(i, segmentCount, bisection.ChallengedSegment.Length)
		chunkHash := core.BisectionChunkHash(segmentStart, segmentLength, cutHashes[i], cutHashes[i+1])
		chunks = append(chunks, chunkHash)
		segmentStart = segmentStart.Add(segmentStart, segmentLength)
	}
	return cutHashes, NewMerkleTree(chunks)
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
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	subCuts []core.Cut,
) (*types.Transaction, error) {
	subCutHashes := cutsToHashes(subCuts)
	prevCutHashes, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	return c.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return c.con.BisectInboxConsistency(
			auth,
			nodes,
			path,
			challengedSegment.Start,
			challengedSegment.Length,
			prevCutHashes[segmentToChallenge+1],
			subCutHashes,
		)
	})
}

func (c *Challenge) OneStepProveInboxConsistency(
	ctx context.Context,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	lowerHash [32]byte,
	value [32]byte,
) (*types.Transaction, error) {
	prevCutHashes, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	return c.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return c.con.OneStepProveInboxConsistency(
			auth,
			nodes,
			path,
			challengedSegment.Start,
			prevCutHashes[segmentToChallenge+1],
			lowerHash,
			value,
		)
	})
}

func (c *Challenge) BisectInboxDelta(
	ctx context.Context,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	subCuts []core.Cut,
) (*types.Transaction, error) {
	subInboxAccHashes := make([][32]byte, 0, len(subCuts))
	subInboxDeltaHashes := make([][32]byte, 0, len(subCuts))
	for _, cut := range subCuts {
		subInboxAccHashes = append(subInboxAccHashes, cut.(core.InboxDeltaCut).InboxAccHash)
		subInboxDeltaHashes = append(subInboxDeltaHashes, cut.(core.InboxDeltaCut).InboxDeltaHash)
	}
	_, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	return c.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return c.con.BisectInboxDelta(
			auth,
			nodes,
			path,
			challengedSegment.Start,
			challengedSegment.Length,
			prevBisection.Cuts[segmentToChallenge+1].(core.InboxDeltaCut).InboxDeltaHash,
			subInboxAccHashes,
			subInboxDeltaHashes,
		)
	})
}

func (c *Challenge) OneStepProveInboxDelta(
	ctx context.Context,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	msg inbox.InboxMessage,
) (*types.Transaction, error) {
	_, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	oldBefore := prevBisection.Cuts[segmentToChallenge].(core.InboxDeltaCut)
	oldAfter := prevBisection.Cuts[segmentToChallenge+1].(core.InboxDeltaCut)
	return c.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return c.con.OneStepProveInboxDelta(
			auth,
			nodes,
			path,
			challengedSegment.Start,
			oldAfter.InboxDeltaHash,
			oldBefore.InboxDeltaHash,
			oldAfter.InboxAccHash,
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
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	subCuts []core.Cut,
) (*types.Transaction, error) {
	subCutHashes := cutsToHashes(subCuts)
	prevCutHashes, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	return c.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return c.con.BisectExecution(
			auth,
			nodes,
			path,
			challengedSegment.Start,
			challengedSegment.Length,
			prevCutHashes[segmentToChallenge+1],
			subCutHashes,
			subCuts[0].(core.ExecutionCut).GasUsed,
			subCuts[0].(core.ExecutionCut).RestHash(),
		)
	})
}

func (c *Challenge) OneStepProveExecution(
	ctx context.Context,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	beforeAssertion *core.AssertionInfo,
	executionProof []byte,
	bufferProof []byte,
) (*types.Transaction, error) {
	prevCutHashes, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	return c.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return c.con.OneStepProveExecution(
			auth,
			nodes,
			path,
			prevBisection.ChallengedSegment.Start,
			prevCutHashes[segmentToChallenge+1],
			[3][32]byte{
				beforeAssertion.InboxDelta,
				beforeAssertion.SendAcc,
				beforeAssertion.LogAcc,
			},
			beforeAssertion.GasUsed().Uint64(),
			beforeAssertion.SendCount(),
			beforeAssertion.LogCount(),
			executionProof,
			bufferProof,
		)
	})
}

func cutsToHashes(cuts []core.Cut) [][32]byte {
	cutHashes := make([][32]byte, 0, len(cuts))
	for _, cut := range cuts {
		cutHashes = append(cutHashes, cut.CutHash())
	}
	return cutHashes
}
