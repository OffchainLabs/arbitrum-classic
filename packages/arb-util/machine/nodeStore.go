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

type NodeStore interface {
	PutNode(height uint64, hash common.Hash, data []byte) error

	GetNode(height uint64, hash common.Hash) ([]byte, error)
	GetNodeHeight(hash common.Hash) (uint64, error)
	GetNodeHash(height uint64) (common.Hash, error)
	Empty() bool
	MaxHeight() uint64
}
