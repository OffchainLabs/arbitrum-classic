package machine

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"time"
)

type Status int

const (
	Extensive Status = iota
	ErrorStop
	Halt
)

type Machine interface {
	Hash() common.Hash
	Clone() Machine
	PrintState()

	CurrentStatus() Status
	IsBlocked(currentTime *common.TimeBlocks, newMessages bool) BlockReason

	ExecuteAssertion(
		maxSteps uint64,
		timeBounds *protocol.TimeBoundsBlocks,
		inbox value.TupleValue,
		maxWallTime time.Duration,
	) (*protocol.ExecutionAssertion, uint64)

	MarshalForProof() ([]byte, error)

	Checkpoint(storage CheckpointStorage) bool
}
