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

package hashing

import (
	"bytes"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

func CombineMessages(
	messages []value.Value,
) []byte {
	var messageData bytes.Buffer
	for _, msg := range messages {
		_ = value.MarshalValue(msg, &messageData)
	}
	return messageData.Bytes()
}

func UnanimousAssertPartialPartialHash(
	newInboxHash common.Hash,
	assertion *protocol.ExecutionAssertion,
) []byte {
	return solsha3.SoliditySHA3(
		solsha3.Bytes32(newInboxHash.Bytes()),
		solsha3.Bytes32(assertion.AfterHash.Bytes()),
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
	var ret common.Hash
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(unanRest),
		solsha3.Bytes32(beforeHash.Bytes()),
		solsha3.Bytes32(originalInboxHash.Bytes()),
		solsha3.Uint64(sequenceNum),
		solsha3.Bytes32(stub.LastMessageHash.Bytes()),
	))
	return ret, nil
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

	var hash common.Hash
	copy(hash[:], solsha3.SoliditySHA3(
		solsha3.Address(vmID.ToEthAddress()),
		solsha3.Bytes32(partialHash.Bytes()),
		solsha3.Bytes32(assertion.LogsHash().Bytes()),
	))
	return hash, nil
}
