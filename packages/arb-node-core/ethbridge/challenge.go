package ethbridge

import (
	"context"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/pkg/errors"
	"math/big"
)

// TODO: Fill this in
var opcodeProver map[uint8]uint8

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
	*BuilderBackend
	builderCon *ethbridgecontracts.Challenge
}

func NewChallenge(address ethcommon.Address, client ethutils.EthClient, builder *BuilderBackend) (*Challenge, error) {
	builderCon, err := ethbridgecontracts.NewChallenge(address, builder)
	if err != nil {
		return nil, err
	}
	watcher, err := NewChallengeWatcher(address, client)
	if err != nil {
		return nil, err
	}
	return &Challenge{
		ChallengeWatcher: watcher,
		BuilderBackend:   builder,
		builderCon:       builderCon,
	}, nil
}

func (c *Challenge) BisectInboxConsistency(
	ctx context.Context,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	subCuts []core.Cut,
) error {
	subCutHashes := cutsToHashes(subCuts)
	prevCutHashes, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	_, err := c.builderCon.BisectInboxConsistency(
		authWithContext(ctx, c.builderAuth),
		nodes,
		path,
		challengedSegment.Start,
		challengedSegment.Length,
		prevCutHashes[segmentToChallenge+1],
		subCutHashes,
	)
	return err
}

func (c *Challenge) OneStepProveInboxConsistency(
	ctx context.Context,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	lowerHash [32]byte,
	value [32]byte,
) error {
	prevCutHashes, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	_, err := c.builderCon.OneStepProveInboxConsistency(
		authWithContext(ctx, c.builderAuth),
		nodes,
		path,
		challengedSegment.Start,
		prevCutHashes[segmentToChallenge+1],
		lowerHash,
		value,
	)
	return err
}

func (c *Challenge) BisectInboxDelta(
	ctx context.Context,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	subCuts []core.Cut,
) error {
	subInboxAccHashes := make([][32]byte, 0, len(subCuts))
	subInboxDeltaHashes := make([][32]byte, 0, len(subCuts))
	for _, cut := range subCuts {
		subInboxAccHashes = append(subInboxAccHashes, cut.(core.InboxDeltaCut).InboxAccHash)
		subInboxDeltaHashes = append(subInboxDeltaHashes, cut.(core.InboxDeltaCut).InboxDeltaHash)
	}
	_, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	_, err := c.builderCon.BisectInboxDelta(
		authWithContext(ctx, c.builderAuth),
		nodes,
		path,
		challengedSegment.Start,
		challengedSegment.Length,
		prevBisection.Cuts[segmentToChallenge+1].(core.InboxDeltaCut).InboxDeltaHash,
		subInboxAccHashes,
		subInboxDeltaHashes,
	)
	return err
}

func (c *Challenge) OneStepProveInboxDelta(
	ctx context.Context,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	msg inbox.InboxMessage,
) error {
	_, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	oldBefore := prevBisection.Cuts[segmentToChallenge].(core.InboxDeltaCut)
	oldAfter := prevBisection.Cuts[segmentToChallenge+1].(core.InboxDeltaCut)
	_, err := c.builderCon.OneStepProveInboxDelta(
		authWithContext(ctx, c.builderAuth),
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
	return err
}

func (c *Challenge) BisectExecution(
	ctx context.Context,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	subCuts []core.Cut,
) error {
	subCutHashes := cutsToHashes(subCuts)
	prevCutHashes, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	_, err := c.builderCon.BisectExecution(
		authWithContext(ctx, c.builderAuth),
		nodes,
		path,
		challengedSegment.Start,
		challengedSegment.Length,
		prevCutHashes[segmentToChallenge+1],
		subCuts[0].(core.ExecutionCut).GasUsed,
		subCuts[0].(core.ExecutionCut).RestHash(),
		subCutHashes,
	)
	return err
}

func (c *Challenge) OneStepProveExecution(
	ctx context.Context,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	beforeExecInfo *core.ExecutionInfo,
	beforeInboxDelta common.Hash,
	executionProof []byte,
	bufferProof []byte,
	opcode uint8,
) error {
	prevCutHashes, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	prover, ok := opcodeProver[opcode]
	if !ok {
		return errors.New("no prover for opcode")
	}
	_, err := c.builderCon.OneStepProveExecution(
		authWithContext(ctx, c.builderAuth),
		nodes,
		path,
		prevBisection.ChallengedSegment.Start,
		prevCutHashes[segmentToChallenge+1],
		[3][32]byte{
			beforeInboxDelta,
			beforeExecInfo.SendAcc,
			beforeExecInfo.LogAcc,
		},
		[3]*big.Int{
			beforeExecInfo.GasUsed(),
			beforeExecInfo.SendCount(),
			beforeExecInfo.LogCount(),
		},
		executionProof,
		bufferProof,
		prover,
	)
	return err
}

func cutsToHashes(cuts []core.Cut) [][32]byte {
	cutHashes := make([][32]byte, 0, len(cuts))
	for _, cut := range cuts {
		cutHashes = append(cutHashes, cut.CutHash())
	}
	return cutHashes
}
