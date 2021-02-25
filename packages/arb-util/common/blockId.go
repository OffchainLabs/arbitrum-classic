/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

package common

import (
	"fmt"
	"github.com/rs/zerolog"
)

type BlockId struct {
	Height     *TimeBlocks
	HeaderHash Hash
}

func (id *BlockId) MarshalZerologObject(e *zerolog.Event) {
	e.Str("height", id.Height.String()).
		Hex("header", id.HeaderHash.Bytes())
}

func (id *BlockId) Clone() *BlockId {
	return &BlockId{
		Height:     id.Height.Clone(),
		HeaderHash: id.HeaderHash,
	}
}

func (id *BlockId) Equals(id2 *BlockId) bool {
	return id.Height.Cmp(id2.Height) == 0 &&
		id.HeaderHash.Equals(id2.HeaderHash)
}

func (id *BlockId) String() string {
	return fmt.Sprintf("Block(%v, %v)", id.Height.AsInt(), id.HeaderHash)
}
