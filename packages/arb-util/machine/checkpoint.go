package machine

import "github.com/offchainlabs/arbitrum/packages/arb-util/value"

type CheckpointStorage interface {
	DeleteCheckpoint(checkpointName string) bool
	GetInitialMachine() (Machine, error)
	SaveValue(val value.Value) bool
	GetValue(hashValue value.Value) value.Value
	DeleteValue(hashValue value.Value) bool
	SaveData(key string, serializedValue string) bool
	GetData(key string) string
	DeleteData(key string) bool
}
