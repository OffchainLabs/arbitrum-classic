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

package ethvalidator

import (
	"bytes"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"

	"github.com/offchainlabs/arb-validator/valmessage"
)

func tokenTypeEncoded(input [21]byte) []byte {
	return common.RightPadBytes(input[:], 21)
}

func tokenTypeArrayEncoded(input [][21]byte) []byte {
	var values []byte
	for _, val := range input {
		values = append(values, tokenTypeEncoded(val)...)
	}
	return values
}

func bytes32ArrayEncoded(input [][32]byte) []byte {
	var values []byte
	for _, val := range input {
		values = append(values, common.RightPadBytes(val[:], 32)...)
	}
	return values
}

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

func UnanimousAssertHash(
	vmId [32]byte,
	sequenceNum uint64,
	beforeHash [32]byte,
	timeBounds protocol.TimeBounds,
	newInboxHash [32]byte,
	originalInboxHash [32]byte,
	assertion *protocol.Assertion,
) ([32]byte, error) {
	tokenNums := make([]uint16, 0, len(assertion.OutMsgs))
	amounts := make([]*big.Int, 0, len(assertion.OutMsgs))
	destinations := make([][32]byte, 0, len(assertion.OutMsgs))
	balance := protocol.NewBalanceTrackerFromMessages(assertion.OutMsgs)
	var messageData bytes.Buffer

	for _, msg := range assertion.OutMsgs {
		tokenNums = append(tokenNums, uint16(balance.TokenIndex(msg.TokenType, msg.Currency)))
		amounts = append(amounts, msg.Currency)
		destinations = append(destinations, msg.Destination)
		err := msg.Data.Marshal(&messageData)
		if err != nil {
			return [32]byte{}, err
		}
	}

	var ret [32]byte
	if sequenceNum == math.MaxUint64 {
		copy(ret[:], solsha3.SoliditySHA3(
			solsha3.Bytes32(vmId),
			solsha3.Bytes32(beforeHash),
			solsha3.Bytes32(originalInboxHash),
			solsha3.Bytes32(assertion.AfterHash),
			solsha3.Bytes32(assertion.LogsHash()),
			solsha3.Bytes32(newInboxHash),
			solsha3.Int64Array(timeBounds),
			tokenTypeArrayEncoded(balance.TokenTypes),
			messageData.Bytes(),
			solsha3.Uint16Array(tokenNums),
			solsha3.Uint256Array(amounts),
			bytes32ArrayEncoded(destinations),
		))
	} else {
		copy(ret[:], solsha3.SoliditySHA3(
			solsha3.Bytes32(vmId),
			solsha3.Bytes32(solsha3.SoliditySHA3(
				solsha3.Bytes32(newInboxHash),
				solsha3.Bytes32(assertion.AfterHash),
				solsha3.Bytes32(assertion.LogsHash()),
				messageData.Bytes(),
				bytes32ArrayEncoded(destinations),
			)),
			solsha3.Uint64(sequenceNum),
			solsha3.Int64Array(timeBounds),
			solsha3.Bytes32(beforeHash),
			solsha3.Bytes32(originalInboxHash),
			tokenTypeArrayEncoded(balance.TokenTypes),
			solsha3.Uint16Array(tokenNums),
			solsha3.Uint256Array(amounts),
		))
	}
	return ret, nil
}
