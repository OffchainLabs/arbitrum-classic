package machine

import "github.com/offchainlabs/arbitrum/packages/arb-util/value"

type CheckpointStorage interface {
	DeleteCheckpoint(checkpointName string) bool
	SaveValue(val value.Value) bool
}
