/*
 * Copyright 2019, Offchain Labs, Inc.
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

package valmessage

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type UnanimousRequestData struct {
	BeforeHash  common.Hash
	BeforeInbox common.Hash
	SequenceNum uint64
	TimeBounds  *protocol.TimeBounds
}

func (r UnanimousRequestData) Hash() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(r.BeforeHash),
		hashing.Bytes32(r.BeforeInbox),
		hashing.Uint64(r.SequenceNum),
		hashing.TimeBlocks(r.TimeBounds.LowerBoundBlock),
		hashing.TimeBlocks(r.TimeBounds.UpperBoundBlock),
	)
}

type UnanimousRequest struct {
	UnanimousRequestData
	NewMessages []valprotocol.Message
}

type UnanimousUpdateResults struct {
	UnanimousRequestData
	NewInboxHash common.Hash
	Assertion    *protocol.ExecutionAssertion
	NewLogCount  int
}
