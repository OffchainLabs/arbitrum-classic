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

	"github.com/ethereum/go-ethereum/common"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

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
	newInboxHash [32]byte,
	assertion *protocol.ExecutionAssertion,
) []byte {
	return solsha3.SoliditySHA3(
		solsha3.Bytes32(newInboxHash),
		solsha3.Bytes32(assertion.AfterHash),
	)
}

func UnanimousAssertPartialHash(
	sequenceNum uint64,
	beforeHash [32]byte,
	newInboxHash [32]byte,
	originalInboxHash [32]byte,
	assertion *protocol.ExecutionAssertion,
) ([32]byte, error) {
	stub := valprotocol.NewExecutionAssertionStubFromAssertion(assertion)
	unanRest := UnanimousAssertPartialPartialHash(
		newInboxHash,
		assertion,
	)
	var ret [32]byte
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(unanRest),
		solsha3.Bytes32(beforeHash),
		solsha3.Bytes32(originalInboxHash),
		solsha3.Uint64(sequenceNum),
		solsha3.Bytes32(stub.LastMessageHash),
	))
	return ret, nil
}

func UnanimousAssertHash(
	vmID common.Address,
	sequenceNum uint64,
	beforeHash [32]byte,
	newInboxHash [32]byte,
	originalInboxHash [32]byte,
	assertion *protocol.ExecutionAssertion,
) ([32]byte, error) {
	partialHash, err := UnanimousAssertPartialHash(
		sequenceNum,
		beforeHash,
		newInboxHash,
		originalInboxHash,
		assertion,
	)
	if err != nil {
		return [32]byte{}, nil
	}

	var hash [32]byte
	copy(hash[:], solsha3.SoliditySHA3(
		solsha3.Address(vmID),
		solsha3.Bytes32(partialHash),
		solsha3.Bytes32(assertion.LogsHash()),
	))
	return hash, nil
}
