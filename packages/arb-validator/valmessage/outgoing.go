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
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type FinalizedAssertion struct {
	Assertion       *protocol.Assertion
	NewLogCount     int
	Signatures      [][]byte // Unanimous Validator signatures
	ProposalResults *UnanimousUpdateResults
	OnChainTxHash   []byte // Disputable assertion on-chain Tx hash
}

func (f FinalizedAssertion) NewLogs() []value.Value {
	return f.Assertion.Logs[len(f.Assertion.Logs)-f.NewLogCount:]
}
