package machine

import (
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
	Hash() [32]byte
	Clone() Machine
	PrintState()

	CurrentStatus() Status
	LastBlockReason() BlockReason

	ExecuteAssertion(maxSteps uint32, timeBounds *protocol.TimeBoundsBlocks, inbox value.TupleValue) (*protocol.ExecutionAssertion, uint32)
	MarshalForProof() ([]byte, error)

	Checkpoint(storage CheckpointStorage) bool
	RestoreCheckpoint(storage CheckpointStorage, checkpointName string) bool
}

func IsMachineBlocked(machine Machine, currentTime *protocol.TimeBlocks) bool {
	lastReason := machine.LastBlockReason()
	if lastReason == nil {
		return false
	}
	return lastReason.IsBlocked(machine, currentTime)
}
