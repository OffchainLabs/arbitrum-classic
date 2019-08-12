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
	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

type UnanimousRequestData struct {
	BeforeHash  [32]byte
	BeforeInbox [32]byte
	SequenceNum uint64
	TimeBounds  protocol.TimeBounds
}

func (r UnanimousRequestData) Hash() [32]byte {
	var ret [32]byte
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(r.BeforeHash),
		solsha3.Bytes32(r.BeforeInbox),
		solsha3.Uint64(r.SequenceNum),
		solsha3.Uint64(r.TimeBounds[0]),
		solsha3.Uint64(r.TimeBounds[1]),
	))
	return ret
}

type UnanimousRequest struct {
	UnanimousRequestData
	NewMessages []protocol.Message
}

type UnanimousUpdateResults struct {
	UnanimousRequestData
	NewInboxHash [32]byte
	Assertion    *protocol.Assertion
	NewLogCount  int
}
