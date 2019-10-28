package machine

import "unsafe"

type CheckpointStorage interface {
	DeleteCheckpoint(checkpointName string) bool
	GetCStorage() unsafe.Pointer
}
