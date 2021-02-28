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
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var logger = log.With().Caller().Str("component", "message").Logger()

const (
	EthType inbox.Type = iota
	ERC20Type
	ERC721Type
	L2Type
	InitType
	L2BuddyDeploy
	EndOfBlockType
	EthDepositTxType
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
	case ERC20Type:
		return NewERC20FromData(data), nil
	case ERC721Type:
		return NewERC721FromData(data), nil
	case L2Type:
		return L2Message{Data: data}, nil
	case InitType:
		return NewInitFromData(data), nil
	case L2BuddyDeploy:
		return NewBuddyDeploymentFromData(data), nil
	case EthDepositTxType:
		return NewEthDepositTxFromData(data), nil
	default:
		return nil, errors.New("unknown inbox l2message type")
	}
}

type BuddyDeployment struct {
	MaxGas      *big.Int
	GasPriceBid *big.Int
	Payment     *big.Int
	Data        []byte
}

func (b BuddyDeployment) Type() inbox.Type {
	return L2BuddyDeploy
}

func NewBuddyDeploymentFromData(data []byte) BuddyDeployment {
	maxGas, data := extractUInt256(data)
	gasPriceBid, data := extractUInt256(data)
	payment, data := extractUInt256(data)
	return BuddyDeployment{
		MaxGas:      maxGas,
		GasPriceBid: gasPriceBid,
		Payment:     payment,
		Data:        data,
	}
}

func NewRandomBuddyDeployment() BuddyDeployment {
	return BuddyDeployment{
		MaxGas:      common.RandBigInt(),
		GasPriceBid: common.RandBigInt(),
		Payment:     common.RandBigInt(),
		Data:        common.RandBytes(200),
	}
}

func (t BuddyDeployment) AsData() []byte {
	ret := make([]byte, 0)
	ret = append(ret, math.U256Bytes(t.MaxGas)...)
	ret = append(ret, math.U256Bytes(t.GasPriceBid)...)
	ret = append(ret, math.U256Bytes(t.Payment)...)
	ret = append(ret, t.Data...)
	return ret
}

func (t BuddyDeployment) AsEthTx() *types.Transaction {
	return types.NewContractCreation(0, t.GasPriceBid, t.MaxGas.Uint64(), t.Payment, t.Data)
}
