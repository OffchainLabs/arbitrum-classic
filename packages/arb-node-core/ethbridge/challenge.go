/*
 * Copyright 2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ethbridge

import (
	"context"
	"math/big"

	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
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

func NewChallenge(address ethcommon.Address, fromBlock int64, client ethutils.EthClient, builder *BuilderBackend, callOpts bind.CallOpts) (*Challenge, error) {
	builderCon, err := ethbridgecontracts.NewChallenge(address, builder)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	watcher, err := NewChallengeWatcher(address, fromBlock, client, callOpts)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Challenge{
		ChallengeWatcher: watcher,
		BuilderBackend:   builder,
		builderCon:       builderCon,
	}, nil
}

func addStackTrace(err error) error {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}
	_, ok := err.(stackTracer)
	if ok {
		return err
	}
	return errors.WithStack(err)
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
		subCuts[0].(*core.ExecutionState).TotalGasConsumed,
		subCuts[0].(*core.ExecutionState).RestHash(),
		subCutHashes,
	)

	return addStackTrace(err)
}

func (c *Challenge) OneStepProveExecution(
	ctx context.Context,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	beforeCut *core.ExecutionState,
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
		[2][32]byte{beforeCut.SendAcc, beforeCut.LogAcc},
		[3]*big.Int{
			beforeCut.TotalGasConsumed,
			beforeCut.TotalSendCount,
			beforeCut.TotalLogCount,
		},
		executionProof,
		bufferProof,
		prover,
	)
	return errors.WithStack(err)
}

func (c *Challenge) ProveContinuedExecution(
	ctx context.Context,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	beforeCut *core.ExecutionState,
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
		beforeCut.TotalGasConsumed,
		beforeCut.RestHash(),
	)
	return errors.WithStack(err)
}

func (c *Challenge) Timeout(
	ctx context.Context,
) error {
	_, err := c.builderCon.Timeout(authWithContext(ctx, c.builderAuth))
	return errors.WithStack(err)
}

func cutsToHashes(cuts []core.Cut) [][32]byte {
	cutHashes := make([][32]byte, 0, len(cuts))
	for _, cut := range cuts {
		cutHashes = append(cutHashes, cut.CutHash())
	}
	return cutHashes
}
