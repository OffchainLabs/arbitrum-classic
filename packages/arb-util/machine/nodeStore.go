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

package machine

import "C"
import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

// ConfirmedNodeStore provides a mechanism for recording data base rollup nodes indexed
// by their height and hash. The intention of this interface is to be used exclusively
// for recording nodes that will not be removed in the future. As such the only
// current usage is for recording confirmed nodes. This limitation is because
// the GetNodeHeight function described below relies on the assumption that the
// most recently recorded node at a given height is the correct/relevant one
type ConfirmedNodeStore interface {
	// PutNode records a record into the database with the given height, hash, and data
	PutNode(height uint64, hash common.Hash, data []byte) error

	// GetNode looks up a record in the database with the given height and hash
	GetNode(height uint64, hash common.Hash) ([]byte, error)

	// GetNodeHeight returns the height of the record with the given hash which
	// should be unique across all possible nodes
	GetNodeHeight(hash common.Hash) (uint64, error)

	// GetNodeHash returns the hash of the node most recently saved at that height
	GetNodeHash(height uint64) (common.Hash, error)

	// Empty returns true if the store is empty and false otherwise
	Empty() bool

	// MaxHeight returns maximum height node stored in the database
	// If the database is empty, MaxHeight returns 0
	MaxHeight() uint64
}
