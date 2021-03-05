package ethbridge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
)

func calculateBisectionChunkCount(segmentIndex, segmentCount int, totalLength *big.Int) *big.Int {
	size := new(big.Int).Div(totalLength, big.NewInt(int64(segmentCount)))
	if segmentIndex == 0 {
		size = size.Add(size, new(big.Int).Mod(totalLength, big.NewInt(int64(segmentCount))))
	}
	return size
}

func calculateBisectionTree(bisection *core.Bisection) ([][32]byte, *protocol.MerkleTree) {
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
	return cutHashes, protocol.NewMerkleTree(chunks)
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
	challengedSegment *core.ChallengeSegment,
	beforeCut core.ExecutionCut,
	executionProof []byte,
	bufferProof []byte,
	opcode uint8,
) error {
	prevCutHashes, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	var prover uint8
	if (opcode >= 0xa1 && opcode <= 0xa6) || opcode == 0x70 {
		// OSP2 (covers buffer related stuff)
		prover = 1
	} else if opcode >= 0x20 && opcode <= 0x24 {
		// OSPHash
		prover = 2
	} else {
		// OSP
		prover = 0
	}
	_, err := c.builderCon.OneStepProveExecution(
		authWithContext(ctx, c.builderAuth),
		nodes,
		path,
		challengedSegment.Start,
		challengedSegment.Length,
		prevCutHashes[segmentToChallenge+1],
		beforeCut.TotalMessagesRead,
		beforeCut.SendAcc,
		beforeCut.LogAcc,
		[3]*big.Int{
			beforeCut.GasUsed,
			beforeCut.SendCount,
			beforeCut.LogCount,
		},
		executionProof,
		bufferProof,
		prover,
	)
	return err
}

func (c *Challenge) ProveContinuedExecution(
	ctx context.Context,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	beforeCut core.ExecutionCut,
) error {
	prevCutHashes, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	_, err := c.builderCon.ProveContinuedExecution(
		authWithContext(ctx, c.builderAuth),
		nodes,
		path,
		challengedSegment.Start,
		challengedSegment.Length,
		prevCutHashes[segmentToChallenge+1],
		beforeCut.GasUsed,
		beforeCut.RestHash(),
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
