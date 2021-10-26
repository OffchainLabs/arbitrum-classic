package blockcache

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"sync"
	"time"
)

type record struct {
	Header   *types.Header
	Snapshot *snapshot.Snapshot
}

type BlockCache struct {
	expiration time.Duration

	lock         sync.RWMutex
	cache        map[uint64]*record
	nextHeight   uint64
	oldestHeight uint64
}

func New(initialSize int, expiration time.Duration) (*BlockCache, error) {
	return &BlockCache{
		expiration: expiration,
		cache:      make(map[uint64]*record, initialSize),
	}, nil
}

func (bc *BlockCache) Size() int {
	bc.lock.Lock()
	defer bc.lock.Unlock()

	return len(bc.cache)
}

// emptyCacheNoLock removes all entries
func (bc *BlockCache) emptyCacheNoLock() {
	bc.cache = make(map[uint64]*record, len(bc.cache))
	bc.oldestHeight = 0
	bc.nextHeight = 0
}

// reorgNoLock removes obsolete blocks including and after nextHeight
func (bc *BlockCache) reorgNoLock(nextHeight uint64) {
	if bc.nextHeight == 0 {
		// Nothing to remove
		return
	}

	if nextHeight <= bc.oldestHeight {
		// Remove everything
		bc.emptyCacheNoLock()
		return
	}

	for currentHeight := bc.nextHeight - 1; currentHeight >= nextHeight; currentHeight-- {
		delete(bc.cache, currentHeight)
	}

	bc.nextHeight = nextHeight
}

// deleteExpiredNoLock deletes all expired log entries
func (bc *BlockCache) deleteExpiredNoLock() {
	expired := bc.expiredTimestamp()
	for bc.oldestHeight < bc.nextHeight {
		if record, exists := bc.cache[bc.oldestHeight]; exists {
			if record.Header.Time > expired {
				// Rest of blocks are not expired
				break
			}

			// Block has expired
			delete(bc.cache, bc.oldestHeight)
		}

		bc.oldestHeight++
	}
}

func (bc *BlockCache) expiredTimestamp() uint64 {
	return uint64(time.Now().Add(bc.expiration * -1).Unix())
}

// Add latest block onto cache
func (bc *BlockCache) Add(header *types.Header, value *snapshot.Snapshot) {
	bc.lock.Lock()
	defer bc.lock.Unlock()

	height := header.Number.Uint64()

	bc.reorgNoLock(height)
	bc.deleteExpiredNoLock()

	if header.Time <= bc.expiredTimestamp() {
		// Don't save expired snapshot to cache
		return
	}

	// Add new entry
	bc.cache[height] = &record{
		Header:   header,
		Snapshot: value,
	}
	bc.nextHeight = height + 1

	if len(bc.cache) == 1 {
		// Reset oldest height
		bc.oldestHeight = height
	}
}

// Get returns snapshot at requested height
func (bc *BlockCache) Get(height uint64) (*types.Header, *snapshot.Snapshot) {
	bc.lock.Lock()
	defer bc.lock.Unlock()

	record, ok := bc.cache[height]
	if !ok {
		return nil, nil
	}

	return record.Header, record.Snapshot
}

// Reorg removes all obsolete blocks up to and including nextHeight
func (bc *BlockCache) Reorg(nextHeight uint64) {
	bc.lock.Lock()
	defer bc.lock.Unlock()

	bc.reorgNoLock(nextHeight)
}
