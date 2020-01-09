package machine

import "github.com/offchainlabs/arbitrum/packages/arb-util/value"

type CheckpointStorage interface {
	DeleteCheckpoint(machineHash [32]byte) bool
	CloseCheckpointStorage() bool
	GetInitialMachine() (Machine, error)
	SaveValue(val value.Value) bool
	GetValue(hashValue [32]byte) value.Value
	DeleteValue(hashValue [32]byte) bool
	SaveData(key []byte, serializedValue []byte) bool
	GetData(key []byte) []byte
	DeleteData(key []byte) bool
}
