package machine

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
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

	ExecuteAssertion(maxSteps uint32, timeBounds *protocol.TimeBoundsBlocks, inbox value.TupleValue) (*protocol.ExecutionAssertion, uint32)
	MarshalForProof() ([]byte, error)

	Checkpoint(storage CheckpointStorage) bool
	RestoreCheckpoint(storage CheckpointStorage, machineHash common.Hash) bool
}
