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
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

func CreateVMHash(data *valmessage.CreateVMValidatorRequest) [32]byte {
	var ret [32]byte
	keys := make([]common.Address, 0, len(data.Config.AssertKeys))
	for _, key := range data.Config.AssertKeys {
		var address common.Address
		copy(address[:], key.Value)
		keys = append(keys, address)
	}
	var owner common.Address
	copy(owner[:], data.Config.Owner.Value)
	var escrowCurrency common.Address
	copy(escrowCurrency[:], data.Config.EscrowCurrency.Value)
	createHash := solsha3.SoliditySHA3(
		solsha3.Uint32(uint32(data.Config.GracePeriod)),
		solsha3.Uint128(value.NewBigIntFromBuf(data.Config.EscrowRequired)),
		solsha3.Address(escrowCurrency),
		solsha3.Uint32(data.Config.MaxExecutionStepCount),
		solsha3.Bytes32(value.NewHashFromBuf(data.VmState)),
		solsha3.Uint16(data.ChallengeManagerNum),
		solsha3.Address(owner),
		solsha3.AddressArray(keys),
	)
	copy(ret[:], createHash)
	createHash = nil
	return ret
}

func SplitMessages(
	outMsgs []protocol.Message,
	balance *protocol.BalanceTracker,
) ([]uint16, []*big.Int, [][32]byte) {
	tokenNums := make([]uint16, 0, len(outMsgs))
	amounts := make([]*big.Int, 0, len(outMsgs))
	destinations := make([][32]byte, 0, len(outMsgs))
	for _, msg := range outMsgs {
		tokenNums = append(tokenNums,
			uint16(balance.TokenIndex(msg.TokenType, msg.Currency)))
		amounts = append(amounts, msg.Currency)
		destinations = append(destinations, msg.Destination)
	}
	return tokenNums, amounts, destinations
}

func UnanimousAssertPartialPartialHash(
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
	messageData bytes.Buffer,
	destinations [][32]byte,
) []byte {
	return solsha3.SoliditySHA3(
		solsha3.Bytes32(newInboxHash),
		solsha3.Bytes32(assertion.AfterHash),
		messageData.Bytes(),
		value.Bytes32ArrayEncoded(destinations),
	)
}

func UnanimousAssertPartialHash(
	sequenceNum uint64,
	beforeHash [32]byte,
	newInboxHash [32]byte,
	originalInboxHash [32]byte,
	assertion *protocol.Assertion,
) ([32]byte, error) {
	balance := protocol.NewBalanceTrackerFromMessages(assertion.OutMsgs)
	tokenNums, amounts, destinations := SplitMessages(
		assertion.OutMsgs,
		balance,
	)

	var messageData bytes.Buffer
	for _, msg := range assertion.OutMsgs {
		err := value.MarshalValue(msg.Data, &messageData)
		if err != nil {
			return [32]byte{}, err
		}
	}

	var ret [32]byte
	if sequenceNum == math.MaxUint64 {
		copy(ret[:], solsha3.SoliditySHA3(
			UnanimousAssertPartialPartialHash(
				newInboxHash,
				assertion,
				messageData,
				destinations,
			),
			solsha3.Bytes32(beforeHash),
			solsha3.Bytes32(originalInboxHash),
			protocol.TokenTypeArrayEncoded(balance.TokenTypes),
			solsha3.Uint16Array(tokenNums),
			solsha3.Uint256Array(amounts),
		))
	} else {
		copy(ret[:], solsha3.SoliditySHA3(
			UnanimousAssertPartialPartialHash(
				newInboxHash,
				assertion,
				messageData,
				destinations,
			),
			solsha3.Bytes32(beforeHash),
			solsha3.Bytes32(originalInboxHash),
			protocol.TokenTypeArrayEncoded(balance.TokenTypes),
			solsha3.Uint16Array(tokenNums),
			solsha3.Uint256Array(amounts),
			solsha3.Uint64(sequenceNum),
		))
	}
	return ret, nil
}

func UnanimousAssertHash(
	vmID [32]byte,
	sequenceNum uint64,
	beforeHash [32]byte,
	newInboxHash [32]byte,
	originalInboxHash [32]byte,
	assertion *protocol.Assertion,
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
		solsha3.Bytes32(vmID),
		solsha3.Bytes32(partialHash),
		solsha3.Bytes32(assertion.LogsHash()),
	))
	return hash, nil
}
