package checkpointing

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/pkg/errors"
	"math/big"
)

type memoryCheckpointData struct {
	blockHash common.Hash
	contents  []byte
}

type InMemoryCheckpointer struct {
	initialMachine machine.Machine
	minHeight      uint64
	maxHeight      uint64
	checkpoints    map[uint64][]memoryCheckpointData
	store          *InMemoryStore
}

func NewInMemoryCheckpointer() *InMemoryCheckpointer {
	return &InMemoryCheckpointer{
		checkpoints: make(map[uint64][]memoryCheckpointData),
		store:       NewInMemoryStore(),
	}
}

func (cp *InMemoryCheckpointer) Initialize(arbitrumCodefilePath string) error {
	mach, err := cmachine.New(arbitrumCodefilePath)
	if err != nil {
		return err
	}
	cp.initialMachine = mach
	return nil
}

func (cp *InMemoryCheckpointer) Initialized() bool {
	return cp.initialMachine != nil
}

func (cp *InMemoryCheckpointer) HasCheckpointedState() bool {
	return false
}

func (cp *InMemoryCheckpointer) RestoreLatestState(ctx context.Context, clnt arbbridge.ChainTimeGetter, restoreFunc func([]byte, ckptcontext.RestoreContext, *common.BlockId) error) error {
	for height := cp.maxHeight; height >= cp.minHeight; height-- {
		checkpoints := cp.checkpoints[height]
		if len(checkpoints) > 0 {
			blockId, err := clnt.BlockIdForHeight(ctx, common.NewTimeBlocksInt(int64(height)))
			if err != nil {
				return err
			}
			for _, c := range checkpoints {
				if c.blockHash == blockId.HeaderHash {
					if err := restoreFunc(c.contents, cp.store, blockId); err != nil {
						logger.Error().Stack().Err(err).Msg("Failed load checkpoint")
						continue
					}
					return nil
				}
			}
		}
	}
	return errNoMatchingCheckpoint
}

func (cp *InMemoryCheckpointer) GetInitialMachine(machine.ValueCache) (machine.Machine, error) {
	return cp.initialMachine.Clone(), nil
}

func (cp *InMemoryCheckpointer) AsyncSaveCheckpoint(
	blockId *common.BlockId,
	contents []byte,
	checkpointContext *ckptcontext.CheckpointContext,
) <-chan error {
	errChan := make(chan error, 1)
	height := blockId.Height.AsInt().Uint64()
	cp.checkpoints[height] = append(cp.checkpoints[height], memoryCheckpointData{
		contents: contents,
	})
	for machHash, mach := range checkpointContext.Machines() {
		cp.store.machines[machHash] = mach
	}
	for valHash, val := range checkpointContext.Values() {
		cp.store.values[valHash] = val
	}
	return errChan
}

func (cp *InMemoryCheckpointer) MaxReorgHeight() *big.Int {
	return big.NewInt(100)
}

type InMemoryStore struct {
	values   map[common.Hash]value.Value
	machines map[common.Hash]machine.Machine
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		values:   make(map[common.Hash]value.Value),
		machines: make(map[common.Hash]machine.Machine),
	}
}

func (ms *InMemoryStore) GetValue(hash common.Hash) (value.Value, error) {
	val, ok := ms.values[hash]
	if !ok {
		return nil, errors.New("value not found in db")
	}
	return val, nil
}

func (ms *InMemoryStore) GetMachine(hash common.Hash) (machine.Machine, error) {
	mach, ok := ms.machines[hash]
	if !ok {
		return nil, errors.New("machine not found in db")
	}
	return mach, nil
}
