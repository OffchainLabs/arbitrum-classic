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

package message

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var logger = log.With().Caller().Stack().Str("component", "message").Logger()

const (
	EthType          inbox.Type = 0
	L2Type           inbox.Type = 3
	InitType         inbox.Type = 4
	EthDepositTxType inbox.Type = 7
	RetryableType    inbox.Type = 9
)

type Message interface {
	Type() inbox.Type
	AsData() []byte
}

func NewInboxMessage(msg Message, sender common.Address, inboxSeqNum *big.Int, gasPrice *big.Int, time inbox.ChainTime) inbox.InboxMessage {
	return inbox.InboxMessage{
		Kind:        msg.Type(),
		Sender:      sender,
		InboxSeqNum: inboxSeqNum,
		GasPrice:    gasPrice,
		Data:        msg.AsData(),
		ChainTime:   time,
	}
}

func NewRandomInboxMessage(msg Message) inbox.InboxMessage {
	return NewInboxMessage(
		msg,
		common.RandAddress(),
		common.RandBigInt(),
		common.RandBigInt(),
		inbox.NewRandomChainTime(),
	)
}

func NestedMessage(data []byte, kind inbox.Type) (Message, error) {
	switch kind {
	case EthType:
		return NewEthFromData(data), nil
	case L2Type:
		return L2Message{Data: data}, nil
	case InitType:
		return NewInitFromData(data), nil
	case EthDepositTxType:
		return NewEthDepositTxFromData(data), nil
	case RetryableType:
		return NewRetryableTxFromData(data), nil
	default:
		return nil, errors.New("unknown inbox l2message type")
	}
}

func CalculateRequestId(chainId *big.Int, msgCount *big.Int) common.Hash {
	return hashing.SoliditySHA3(hashing.Uint256(chainId), hashing.Uint256(msgCount))
}

func RetryableId(requestId common.Hash) common.Hash {
	return hashing.SoliditySHA3(hashing.Bytes32(requestId), hashing.Uint256(big.NewInt(0)))
}
