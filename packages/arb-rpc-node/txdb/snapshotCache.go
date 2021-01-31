/*
* Copyright 2020, Offchain Labs, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package txdb

import (
	"github.com/emirpasic/gods/trees/redblacktree"

	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

type snapshotCache struct {
	tree *redblacktree.Tree
	max  int
}

func newSnapshotCache(max int) *snapshotCache {
	if max == 0 {
		panic("must use max greater than 0")
	}
	return &snapshotCache{
		tree: redblacktree.NewWith(func(a, b interface{}) int {
			return a.(*common.TimeBlocks).Cmp(b.(*common.TimeBlocks))
		}),
		max: max,
	}
}

func (sc *snapshotCache) latest() *snapshot.Snapshot {
	node := sc.tree.Right()
	if node == nil {
		return nil
	}
	return node.Value.(*snapshot.Snapshot)
}

func (sc *snapshotCache) getSnapshot(time inbox.ChainTime) *snapshot.Snapshot {
	// If the time is past the most recent value, we have no snapshot
	if sc.tree.Right().Key.(*common.TimeBlocks).Cmp(time.BlockNum) < 0 {
		return nil
	}

	nearest, found := sc.tree.Floor(time.BlockNum)
	if !found {
		return nil
	}

	snap := nearest.Value.(*snapshot.Snapshot)
	if snap.Height().Cmp(time.BlockNum) == 0 {
		return snap
	}

	// advance time on found snapshot
	snap = snap.Clone()
	snap.AdvanceTime(time)
	return snap
}

func (sc *snapshotCache) addSnapshot(snap *snapshot.Snapshot) {
	// Clear out any snapshots that occur at or after this snapshot's height
	for sc.tree.Size() > 0 {
		if sc.tree.Right().Key.(*common.TimeBlocks).Cmp(snap.Height()) < 0 {
			break
		}
		sc.tree.Remove(sc.tree.Right().Key)
	}

	sc.tree.Put(snap.Height(), snap)
	for sc.tree.Size() > sc.max {
		sc.tree.Remove(sc.tree.Left().Key)
	}
}
