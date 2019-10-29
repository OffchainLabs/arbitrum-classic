package machine

type CheckpointStorage interface {
	DeleteCheckpoint(checkpointName string) bool
}
