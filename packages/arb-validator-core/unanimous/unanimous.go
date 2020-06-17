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

package unanimous

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

func UnanimousAssertPartialPartialHash(
	newInboxHash common.Hash,
	assertion *protocol.ExecutionAssertion,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(newInboxHash),
		hashing.Bytes32(assertion.AfterHash.Unmarshal()),
	)
}

func UnanimousAssertPartialHash(
	sequenceNum uint64,
	beforeHash common.Hash,
	newInboxHash common.Hash,
	originalInboxHash common.Hash,
	assertion *protocol.ExecutionAssertion,
) (common.Hash, error) {
	stub := valprotocol.NewExecutionAssertionStubFromAssertion(assertion)
	unanRest := UnanimousAssertPartialPartialHash(
		newInboxHash,
		assertion,
	)
	return hashing.SoliditySHA3(
		hashing.Bytes32(unanRest),
		hashing.Bytes32(beforeHash),
		hashing.Bytes32(originalInboxHash),
		hashing.Uint64(sequenceNum),
		hashing.Bytes32(stub.LastMessageHash),
	), nil
}

func UnanimousAssertHash(
	vmID common.Address,
	sequenceNum uint64,
	beforeHash common.Hash,
	newInboxHash common.Hash,
	originalInboxHash common.Hash,
	assertion *protocol.ExecutionAssertion,
) (common.Hash, error) {
	partialHash, err := UnanimousAssertPartialHash(
		sequenceNum,
		beforeHash,
		newInboxHash,
		originalInboxHash,
		assertion,
	)
	if err != nil {
		return common.Hash{}, nil
	}

	return hashing.SoliditySHA3(
		hashing.Address(vmID),
		hashing.Bytes32(partialHash),
		hashing.Bytes32(valprotocol.NewExecutionAssertionStubFromAssertion(assertion).LastLogHash),
	), nil
}
