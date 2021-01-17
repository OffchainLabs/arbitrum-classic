package ethbridge

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/pkg/errors"
	"math/big"
)

// TODO: Fill this in
var opcodeProver map[uint8]uint8

func init() {
}

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
}

func NewChallenge(address ethcommon.Address, client ethutils.EthClient) (*Challenge, error) {
	watcher, err := NewChallengeWatcher(address, client)
	if err != nil {
		return nil, err
	}
	return &Challenge{
		ChallengeWatcher: watcher,
	}, nil
}

func (c *Challenge) buildTx(name string, amount *big.Int, args ...interface{}) (*RawTransaction, error) {
	data, err := challengeABI.Pack(name, args...)
	return &RawTransaction{
		Data:   data,
		Dest:   c.address,
		Amount: amount,
	}, err
}

func (c *Challenge) buildSimpleTx(name string, args ...interface{}) (*RawTransaction, error) {
	return c.buildTx(name, big.NewInt(0), args...)
}

func (c *Challenge) BisectInboxConsistency(
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	subCuts []core.Cut,
) (*RawTransaction, error) {
	subCutHashes := cutsToHashes(subCuts)
	prevCutHashes, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	return c.buildSimpleTx(
		"bisectInboxConsistency",
		nodes,
		path,
		challengedSegment.Start,
		challengedSegment.Length,
		prevCutHashes[segmentToChallenge+1],
		subCutHashes,
	)
}

func (c *Challenge) OneStepProveInboxConsistency(
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	lowerHash [32]byte,
	value [32]byte,
) (*RawTransaction, error) {
	prevCutHashes, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	return c.buildSimpleTx(
		"oneStepProveInboxConsistency",
		nodes,
		path,
		challengedSegment.Start,
		prevCutHashes[segmentToChallenge+1],
		lowerHash,
		value,
	)
}

func (c *Challenge) BisectInboxDelta(
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	subCuts []core.Cut,
) (*RawTransaction, error) {
	subInboxAccHashes := make([][32]byte, 0, len(subCuts))
	subInboxDeltaHashes := make([][32]byte, 0, len(subCuts))
	for _, cut := range subCuts {
		subInboxAccHashes = append(subInboxAccHashes, cut.(core.InboxDeltaCut).InboxAccHash)
		subInboxDeltaHashes = append(subInboxDeltaHashes, cut.(core.InboxDeltaCut).InboxDeltaHash)
	}
	_, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	return c.buildSimpleTx(
		"bisectInboxDelta",
		nodes,
		path,
		challengedSegment.Start,
		challengedSegment.Length,
		prevBisection.Cuts[segmentToChallenge+1].(core.InboxDeltaCut).InboxDeltaHash,
		subInboxAccHashes,
		subInboxDeltaHashes,
	)
}

func (c *Challenge) OneStepProveInboxDelta(
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	msg inbox.InboxMessage,
) (*RawTransaction, error) {
	_, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	oldBefore := prevBisection.Cuts[segmentToChallenge].(core.InboxDeltaCut)
	oldAfter := prevBisection.Cuts[segmentToChallenge+1].(core.InboxDeltaCut)
	return c.buildSimpleTx(
		"oneStepProveInboxDelta",
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
}

func (c *Challenge) BisectExecution(
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	subCuts []core.Cut,
) (*RawTransaction, error) {
	subCutHashes := cutsToHashes(subCuts)
	prevCutHashes, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	return c.buildSimpleTx(
		"bisectExecution",
		nodes,
		path,
		challengedSegment.Start,
		challengedSegment.Length,
		prevCutHashes[segmentToChallenge+1],
		subCuts[0].(core.ExecutionCut).GasUsed,
		subCuts[0].(core.ExecutionCut).RestHash(),
		subCutHashes,
	)
}

func (c *Challenge) OneStepProveExecution(
	prevBisection *core.Bisection,
	segmentToChallenge int,
	beforeExecInfo *core.ExecutionInfo,
	beforeInboxDelta common.Hash,
	executionProof []byte,
	bufferProof []byte,
	opcode uint8,
) (*RawTransaction, error) {
	prevCutHashes, prevTree := calculateBisectionTree(prevBisection)
	nodes, path := prevTree.GetProof(segmentToChallenge)
	prover, ok := opcodeProver[opcode]
	if !ok {
		return nil, errors.New("no prover for opcode")
	}
	return c.buildSimpleTx(
		"oneStepProveExecution",
		nodes,
		path,
		prevBisection.ChallengedSegment.Start,
		prevCutHashes[segmentToChallenge+1],
		[3][32]byte{
			beforeInboxDelta,
			beforeExecInfo.SendAcc,
			beforeExecInfo.LogAcc,
		},
		[]*big.Int{
			beforeExecInfo.GasUsed(),
			beforeExecInfo.SendCount(),
			beforeExecInfo.LogCount(),
		},
		executionProof,
		bufferProof,
		prover,
	)
}

func cutsToHashes(cuts []core.Cut) [][32]byte {
	cutHashes := make([][32]byte, 0, len(cuts))
	for _, cut := range cuts {
		cutHashes = append(cutHashes, cut.CutHash())
	}
	return cutHashes
}
