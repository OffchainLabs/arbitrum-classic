package challenge

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"math/big"
)

type ExecutionImpl struct {
	initialGasUsed *big.Int
}

func (e ExecutionImpl) GetCuts(lookup core.ValidatorLookup, offsets []*big.Int) ([]ethbridge.Cut, error) {
	panic("implement me")
}

func (e ExecutionImpl) FindFirstDivergence(lookup core.ValidatorLookup, offsets []*big.Int, cuts []ethbridge.Cut) (int, error) {
	panic("implement me")
}

func (e ExecutionImpl) Bisect(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	prevBisection *ethbridge.Bisection,
	segmentToChallenge int,
	subCuts []ethbridge.Cut,
) (*types.Transaction, error) {
	return challenge.BisectExecution(
		ctx,
		prevBisection,
		segmentToChallenge,
		subCuts,
	)
}

func (e ExecutionImpl) OneStepProof(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	lookup core.ValidatorLookup,
	prevBisection *ethbridge.Bisection,
	segmentToChallenge int,
	challengedSegment *ethbridge.ChallengeSegment,
) (*types.Transaction, error) {
	startMachine, err := lookup.GetMachine(e.initialGasUsed)
	if err != nil {
		return nil, err
	}
	beforeAssertion, err := lookup.GetExecutionInfo(startMachine, challengedSegment.Start)
	if err != nil {
		return nil, err
	}

	beforeMachine, err := lookup.GetMachine(new(big.Int).Add(e.initialGasUsed, challengedSegment.Start))
	if err != nil {
		return nil, err
	}

	proofData, err := beforeMachine.MarshalForProof()
	if err != nil {
		return nil, err
	}

	bufferProofData, err := beforeMachine.MarshalBufferProof()
	if err != nil {
		return nil, err
	}

	return challenge.OneStepProveExecution(
		ctx,
		prevBisection,
		segmentToChallenge,
		beforeAssertion,
		proofData,
		bufferProofData,
	)
}
