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

package inbox

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
)

type DelayedMessage struct {
	DelayedSequenceNumber *big.Int
	DelayedAccumulator    common.Hash
	Message               []byte
}

func NewDelayedMessage(beforeAcc common.Hash, message InboxMessage) DelayedMessage {
	return DelayedMessage{
		DelayedSequenceNumber: message.InboxSeqNum,
		DelayedAccumulator: hashing.SoliditySHA3(
			hashing.Bytes32(beforeAcc),
			hashing.Bytes32(message.CommitmentHash()),
		),
		Message: message.ToBytes(),
	}
}

func (m DelayedMessage) ToBytesWithSeqNum() []byte {
	var data []byte
	data = append(data, math.U256Bytes(m.DelayedSequenceNumber)...)
	data = append(data, m.DelayedAccumulator.Bytes()...)
	data = append(data, m.Message...)
	return data
}
