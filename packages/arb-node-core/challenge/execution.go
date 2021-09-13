package challenge

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

func SegmentTarget() int {
	return 400
}

var unreachableCut common.Hash

func getCutRaw(execTracker *core.ExecutionTracker, maxTotalMessagesRead *big.Int, gasTarget *big.Int) (*core.ExecutionState, bool, *big.Int, error) {
	state, steps, err := execTracker.GetExecutionState(gasTarget)
	if err != nil {
		return nil, false, nil, err
	}
	if state.TotalMessagesRead.Cmp(maxTotalMessagesRead) > 0 || state.TotalGasConsumed.Cmp(gasTarget) < 0 {
		// Execution read more messages than provided so assertion should have
		// stopped short
		return nil, false, steps, nil
	}
	return state, true, steps, nil
}

func cutHash(state *core.ExecutionState, reachable bool) common.Hash {
	if !reachable {
		return unreachableCut
	}
	return state.CutHash()
}

func getCut(execTracker *core.ExecutionTracker, maxTotalMessagesRead *big.Int, gasTarget *big.Int) (common.Hash, *big.Int, error) {
	state, reachable, steps, err := getCutRaw(execTracker, maxTotalMessagesRead, gasTarget)
	if err != nil {
		return common.Hash{}, nil, err
	}
	return cutHash(state, reachable), steps, nil
}

func GetCuts(lookup core.ArbCoreLookup, assertion *core.Assertion, offsets []*big.Int) (*core.ExecutionState, []common.Hash, error) {
	execTracker := core.NewExecutionTracker(lookup, true, offsets, true)
	cuts := make([]common.Hash, 0, len(offsets))
	var startState *core.ExecutionState
	for i, offset := range offsets {
		cut, reachable, _, err := getCutRaw(execTracker, assertion.After.TotalMessagesRead, offset)
		if err != nil {
			return nil, nil, err
		}
		if i == 0 {
			if !reachable {
				return nil, nil, errors.New("first cut is unreachable")
			}
			startState = cut
		}
		cuts = append(cuts, cutHash(cut, reachable))
	}
	return startState, cuts, nil
}

type DivergenceInfo struct {
	DifferentIndex   int
	SegmentSteps     *big.Int
	EndIsUnreachable bool
}

func FindFirstDivergence(lookup core.ArbCoreLookup, assertion *core.Assertion, offsets []*big.Int, cuts []common.Hash) (DivergenceInfo, error) {
	errRes := DivergenceInfo{
		DifferentIndex:   0,
		SegmentSteps:     big.NewInt(0),
		EndIsUnreachable: false,
	}
	execTracker := core.NewExecutionTracker(lookup, true, offsets, true)
	lastSteps := big.NewInt(0)
	for i, offset := range offsets {
		localCut, newSteps, err := getCut(execTracker, assertion.After.TotalMessagesRead, offset)
		if err != nil {
			return errRes, err
		}
		if localCut != cuts[i] {
			return DivergenceInfo{
				DifferentIndex:   i,
				SegmentSteps:     new(big.Int).Sub(newSteps, lastSteps),
				EndIsUnreachable: localCut == unreachableCut,
			}, nil
		}
		lastSteps = newSteps
	}
	return errRes, errors.New("no divergence found in cuts")
}

func getSegmentStartInfo(lookup core.ArbCoreLookup, assertion *core.Assertion, segment *core.ChallengeSegment) (*core.ExecutionState, machine.Machine, error) {
	execTracker := core.NewExecutionTracker(lookup, true, []*big.Int{segment.Start}, true)
	state, reachable, _, err := getCutRaw(execTracker, assertion.After.TotalMessagesRead, segment.Start)
	if err != nil {
		return nil, nil, err
	}
	if !reachable {
		return nil, nil, errors.New("attempted to one step prove blocked machine")
	}

	beforeMachine, err := execTracker.GetMachine(segment.Start)
	if err != nil {
		return nil, nil, err
	}

	return state, beforeMachine, nil
}

type Move interface {
	execute(context.Context, *ethbridge.Challenge) error
}

type BisectMove struct {
	prevBisection       *core.Bisection
	startState          *core.ExecutionState
	segmentToChallenge  int
	inconsistentSegment *core.ChallengeSegment
	subCuts             []common.Hash
}

func (m *BisectMove) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind                string
		PrevBisection       *core.Bisection
		StartState          *core.ExecutionState
		SegmentToChallenge  int
		InconsistentSegment *core.ChallengeSegment
		SubCuts             []common.Hash
	}{
		Kind:                "Bisect",
		PrevBisection:       m.prevBisection,
		StartState:          m.startState,
		SegmentToChallenge:  m.segmentToChallenge,
		InconsistentSegment: m.inconsistentSegment,
		SubCuts:             m.subCuts,
	})
}

func (m *BisectMove) execute(ctx context.Context, challenge *ethbridge.Challenge) error {
	logger.Info().
		Str("start", m.inconsistentSegment.Start.String()).
		Str("end", m.inconsistentSegment.GetEnd().String()).
		Msg("Bisecting challenge")
	return challenge.BisectExecution(
		ctx,
		m.prevBisection,
		m.startState,
		m.segmentToChallenge,
		m.inconsistentSegment,
		m.subCuts,
	)
}

type ProveContinuedMove struct {
	assertion          *core.Assertion
	prevBisection      *core.Bisection
	segmentToChallenge int
	challengedSegment  *core.ChallengeSegment
	previousCut        *core.ExecutionState
}

func (m *ProveContinuedMove) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind               string
		Assertion          *core.Assertion
		PrevBisection      *core.Bisection
		SegmentToChallenge int
		ChallengedSegment  *core.ChallengeSegment
		PreviousCut        *core.ExecutionState
	}{
		Kind:               "ProveContinuedExecution",
		Assertion:          m.assertion,
		PrevBisection:      m.prevBisection,
		SegmentToChallenge: m.segmentToChallenge,
		ChallengedSegment:  m.challengedSegment,
		PreviousCut:        m.previousCut,
	})
}

func (m *ProveContinuedMove) execute(ctx context.Context, challenge *ethbridge.Challenge) error {
	logger.Info().
		Str("start", m.challengedSegment.Start.String()).
		Str("end", m.challengedSegment.GetEnd().String()).
		Msg("Proving continued execution")

	return challenge.ProveContinuedExecution(
		ctx,
		m.prevBisection,
		m.segmentToChallenge,
		m.challengedSegment,
		m.previousCut,
	)
}

type OneStepProofMove struct {
	assertion          *core.Assertion
	prevBisection      *core.Bisection
	segmentToChallenge int
	challengedSegment  *core.ChallengeSegment
	previousCut        *core.ExecutionState
	proofData          []byte
	bufferProofData    []byte
}

func NewOneStepProofMove(
	ctx context.Context,
	assertion *core.Assertion,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
	sequencerInbox *ethbridge.SequencerInboxWatcher,
	lookup core.ArbCoreLookup,
) (*OneStepProofMove, error) {
	previousCut, previousMachine, err := getSegmentStartInfo(lookup, assertion, challengedSegment)
	if err != nil {
		return nil, err
	}

	proofData, bufferProofData, err := previousMachine.MarshalForProof()
	if err != nil {
		return nil, err
	}

	opcode := proofData[0]
	if opcode == 0x72 {
		// INBOX proving
		seqNum := previousCut.TotalMessagesRead
		batch, err := LookupBatchContaining(ctx, lookup, sequencerInbox, seqNum)
		if err != nil {
			return nil, err
		}
		if batch == nil {
			return nil, errors.New("Failed to lookup batch containing message")
		}
		inboxProof, err := lookup.GenInboxProof(seqNum, batch.GetBatchIndex(), batch.GetAfterCount())
		if err != nil {
			return nil, err
		}
		proofData = append(proofData, inboxProof...)
	}

	return &OneStepProofMove{
		assertion:          assertion,
		prevBisection:      prevBisection,
		segmentToChallenge: segmentToChallenge,
		challengedSegment:  challengedSegment,
		previousCut:        previousCut,
		proofData:          proofData,
		bufferProofData:    bufferProofData,
	}, nil
}

func (m *OneStepProofMove) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind               string
		Assertion          *core.Assertion
		PrevBisection      *core.Bisection
		SegmentToChallenge int
		ChallengedSegment  *core.ChallengeSegment
		PreviousCut        *core.ExecutionState
		ProofData          hexutil.Bytes
		BufferProofData    hexutil.Bytes
	}{
		Kind:               "OneStepProof",
		Assertion:          m.assertion,
		PrevBisection:      m.prevBisection,
		SegmentToChallenge: m.segmentToChallenge,
		ChallengedSegment:  m.challengedSegment,
		PreviousCut:        m.previousCut,
		ProofData:          m.proofData,
		BufferProofData:    m.bufferProofData,
	})
}

func (m *OneStepProofMove) execute(ctx context.Context, challenge *ethbridge.Challenge) error {
	opcode := m.proofData[0]
	logger.Info().Int("opcode", int(opcode)).Str("gas", m.previousCut.TotalGasConsumed.String()).Msg("Issuing one step proof")

	return challenge.OneStepProveExecution(
		ctx,
		m.prevBisection,
		m.segmentToChallenge,
		m.challengedSegment,
		m.previousCut,
		m.proofData,
		m.bufferProofData,
		opcode,
	)
}

type TimeoutMove struct {
}

func (m *TimeoutMove) execute(ctx context.Context, challenge *ethbridge.Challenge) error {
	return challenge.Timeout(ctx)
}

func (m *TimeoutMove) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind string
	}{
		Kind: "Timeout",
	})
}
