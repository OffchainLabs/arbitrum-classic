package blockcache

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"math/big"
	"testing"
	"time"
)

func TestBlockCacheAdd(t *testing.T) {
	expiration := 3 * time.Second
	cache, err := New(10, expiration)
	if err != nil {
		t.Fatalf("error creating new block cache: %s\n", err.Error())
	}

	// Test that expired block is not added
	cache.Add(&types.Header{
		Number: big.NewInt(0),
		Time:   uint64(time.Now().Add(expiration * -1).Unix())},
		&snapshot.Snapshot{})

	if len(cache.cache) != 0 {
		t.Error("expired block was incorrectly added to cache")
	}

	// Test that non-expired block is added
	cache.Add(&types.Header{Number: big.NewInt(0), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})

	if len(cache.cache) != 1 {
		t.Error("non-expired block was not added to cache")
	}
}

func TestBlockCacheGet(t *testing.T) {
	cache, err := New(10, 3*time.Second)
	if err != nil {
		t.Fatalf("error creating new block cache: %s\n", err.Error())
	}

	cache.Add(&types.Header{Number: big.NewInt(42), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})
	cache.Add(&types.Header{Number: big.NewInt(43), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})

	if len(cache.cache) != 2 {
		t.Fatalf("cache size %v does not equal 2\n", len(cache.cache))
	}

	header, value := cache.Get(42)
	if header == nil || value == nil {
		t.Fatalf("nil value returned by Get\n")
	}

	if header.Number.Uint64() != 42 {
		t.Fatalf("header number %v does not equal 42\n", header.Number.Uint64())
	}

	header2, value2 := cache.Get(44)
	if header2 != nil || value2 != nil {
		t.Fatalf("non-nil value returned by Get for invalid  height\n")
	}

}

func TestBlockCacheReorg(t *testing.T) {
	cache, err := New(10, 30*time.Second)
	if err != nil {
		t.Fatalf("error creating new block cache: %s\n", err.Error())
	}

	cache.Add(&types.Header{Number: big.NewInt(0), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})
	cache.Add(&types.Header{Number: big.NewInt(1), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})
	cache.Add(&types.Header{Number: big.NewInt(2), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})
	cache.Add(&types.Header{Number: big.NewInt(3), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})
	cache.Add(&types.Header{Number: big.NewInt(4), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})

	if len(cache.cache) != 5 {
		t.Fatalf("cache size %v does not equal 5\n", len(cache.cache))
	}

	if cache.oldestHeight != 0 {
		t.Fatalf("oldestHeight %v  does not equal 0\n", cache.oldestHeight)
	}

	if cache.nextHeight != 5 {
		t.Fatalf("nextHeight %v does not equal 5\n", cache.nextHeight)
	}

	// Test reorg above current height
	cache.Reorg(5)

	if len(cache.cache) != 5 {
		t.Fatalf("cache size %v does not equal 5\n", len(cache.cache))
	}

	if cache.oldestHeight != 0 {
		t.Fatalf("oldestHeight %v  does not equal 0\n", cache.oldestHeight)
	}

	if cache.nextHeight != 5 {
		t.Fatalf("nextHeight %v does not equal 5\n", cache.nextHeight)
	}

	// Test reorg single value
	cache.Reorg(4)

	if len(cache.cache) != 4 {
		t.Fatalf("cache size %v does not equal 4\n", len(cache.cache))
	}

	if cache.oldestHeight != 0 {
		t.Fatalf("oldestHeight %v  does not equal 0\n", cache.oldestHeight)
	}

	if cache.nextHeight != 4 {
		t.Fatalf("nextHeight %v does not equal 4\n", cache.nextHeight)
	}

	// Test reorg entire cache
	cache.Reorg(0)

	if len(cache.cache) != 0 {
		t.Fatalf("cache size %v does not equal 0\n", len(cache.cache))
	}

	if cache.oldestHeight != 0 {
		t.Fatalf("oldestHeight %v  does not equal 0\n", cache.oldestHeight)
	}

	if cache.nextHeight != 0 {
		t.Fatalf("nextHeight %v does not equal 0\n", cache.nextHeight)
	}

	cache.Add(&types.Header{Number: big.NewInt(40), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})
	cache.Add(&types.Header{Number: big.NewInt(42), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})
	cache.Add(&types.Header{Number: big.NewInt(43), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})
	cache.Add(&types.Header{Number: big.NewInt(44), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})
	cache.Add(&types.Header{Number: big.NewInt(45), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})

	if len(cache.cache) != 5 {
		t.Fatalf("cache size %v does not equal 5\n", len(cache.cache))
	}

	if cache.oldestHeight != 40 {
		t.Fatalf("oldestHeight %v  does not equal 40\n", cache.oldestHeight)
	}

	if cache.nextHeight != 46 {
		t.Fatalf("nextHeight %v does not equal 46\n", cache.nextHeight)
	}

	// Test reorg to single value
	cache.Reorg(41)

	if len(cache.cache) != 1 {
		t.Fatalf("cache size %v does not equal 1\n", len(cache.cache))
	}

	if cache.oldestHeight != 40 {
		t.Fatalf("oldestHeight %v  does not equal 40\n", cache.oldestHeight)
	}

	if cache.nextHeight != 41 {
		t.Fatalf("nextHeight %v does not equal 41\n", cache.nextHeight)
	}

	// Test reorg below current stack
	cache.Reorg(0)

	if len(cache.cache) != 0 {
		t.Fatalf("cache size %v does not equal 0\n", len(cache.cache))
	}

	if cache.oldestHeight != 0 {
		t.Fatalf("oldestHeight %v  does not equal 0\n", cache.oldestHeight)
	}

	if cache.nextHeight != 0 {
		t.Fatalf("nextHeight %v does not equal 0\n", cache.nextHeight)
	}

	cache.Add(&types.Header{Number: big.NewInt(40), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})
	cache.Add(&types.Header{Number: big.NewInt(42), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})

	// Test implicit reorg to value below current oldest
	cache.Add(&types.Header{Number: big.NewInt(39), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})

	if len(cache.cache) != 1 {
		t.Fatalf("cache size %v does not equal 1\n", len(cache.cache))
	}

	if cache.oldestHeight != 39 {
		t.Fatalf("oldestHeight %v  does not equal 39\n", cache.oldestHeight)
	}

	if cache.nextHeight != 40 {
		t.Fatalf("nextHeight %v does not equal 40\n", cache.nextHeight)
	}

}

func TestBlockCacheExpire(t *testing.T) {
	expiration := 2 * time.Second
	cache, err := New(10, expiration)
	if err != nil {
		t.Fatalf("error creating new block cache: %s\n", err.Error())
	}

	cache.Add(&types.Header{Number: big.NewInt(0), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})
	cache.Add(&types.Header{Number: big.NewInt(1), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})
	cache.Add(&types.Header{Number: big.NewInt(2), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})
	cache.Add(&types.Header{Number: big.NewInt(3), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})
	cache.Add(&types.Header{Number: big.NewInt(4), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})

	if len(cache.cache) != 5 {
		t.Fatalf("cache size %v does not equal 5\n", len(cache.cache))
	}

	if cache.oldestHeight != 0 {
		t.Fatalf("oldestHeight %v  does not equal 0\n", cache.oldestHeight)
	}

	if cache.nextHeight != 5 {
		t.Fatalf("nextHeight %v does not equal 5\n", cache.nextHeight)
	}

	// Let cache expire
	time.Sleep(expiration)

	// Add one more record
	cache.Add(&types.Header{Number: big.NewInt(5), Time: uint64(time.Now().Unix())}, &snapshot.Snapshot{})

	if len(cache.cache) != 1 {
		t.Fatalf("cache size %v does not equal 1\n", len(cache.cache))
	}

	if cache.oldestHeight != 5 {
		t.Fatalf("oldestHeight %v  does not equal 5\n", cache.oldestHeight)
	}

	if cache.nextHeight != 6 {
		t.Fatalf("nextHeight %v does not equal 6\n", cache.nextHeight)
	}
}
