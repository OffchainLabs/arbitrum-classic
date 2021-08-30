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
	"fmt"
	"math/big"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum/common/math"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

var logger = log.With().Caller().Stack().Str("component", "message").Logger()

const (
	L2Type            inbox.Type = 3
	OldInitType       inbox.Type = 4 // remove after upgrade 5
	EndOfBlockType    inbox.Type = 6
	EthDepositTxType  inbox.Type = 7
	RetryableType     inbox.Type = 9
	GasEstimationType inbox.Type = 10
	InitType          inbox.Type = 11
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
	case L2Type:
		return L2Message{Data: data}, nil
	case InitType:
		return NewInitFromData(data)
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

type GasEstimationMessage struct {
	Aggregator       common.Address
	ComputationLimit *big.Int
	TxData           []byte
}

func (t GasEstimationMessage) String() string {
	batch := newTransactionBatchFromData(t.TxData)
	return fmt.Sprintf("GasEstimationMessage{aggregator=%v, computeLimit=%v, tx=%v}", t.Aggregator, t.ComputationLimit, batch)
}

func NewGasEstimationMessage(aggregator common.Address, computationLimit *big.Int, tx CompressedECDSATransaction) (GasEstimationMessage, error) {
	// Make sure upper bound of estimate is accurate
	tx.R = math.MaxBig256
	tx.S = math.MaxBig256
	tx.V = 1
	tx.SequenceNum = big.NewInt(1)

	batch, err := NewTransactionBatchFromMessages([]AbstractL2Message{tx})
	if err != nil {
		return GasEstimationMessage{}, err
	}
	batchData, err := batch.AsData()

	if err != nil {
		return GasEstimationMessage{}, err
	}
	return GasEstimationMessage{
		Aggregator:       aggregator,
		ComputationLimit: computationLimit,
		TxData:           batchData,
	}, nil
}

func (t GasEstimationMessage) AsData() []byte {
	return t.AsDataSafe()
}

func (t GasEstimationMessage) AsDataSafe() []byte {
	ret := make([]byte, 0)
	ret = append(ret, 3)
	ret = append(ret, AddressData(t.Aggregator)...)
	ret = append(ret, math.U256Bytes(t.ComputationLimit)...)
	ret = append(ret, t.TxData...)
	return ret
}

func (t GasEstimationMessage) Type() inbox.Type {
	return GasEstimationType
}

type EndBlockMessage struct {
}

func (t EndBlockMessage) Type() inbox.Type {
	return EndOfBlockType
}

func (t EndBlockMessage) AsData() []byte {
	return nil
}

func L2RemapAccount(account common.Address) common.Address {

	if account == (common.Address{}) {
		return account
	}

	magic, _ := new(big.Int).SetString("1111000000000000000000000000000000001111", 16)
	overflow := new(big.Int).Exp(big.NewInt(2), big.NewInt(20*8), nil)

	translated := new(big.Int).SetBytes(account.Bytes())
	translated.Add(translated, magic)
	if translated.Cmp(overflow) == 1 {
		translated.Sub(translated, overflow)
	}

	return common.NewAddressFromBig(translated)
}
