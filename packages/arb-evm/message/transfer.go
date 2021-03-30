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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
)

type Eth struct {
	Dest  common.Address
	Value *big.Int
}

func NewEthFromData(data []byte) Eth {
	destAddress, data := extractAddress(data)
	// Last value returned is not an error type
	payment, _ := extractUInt256(data)
	return Eth{
		Dest:  destAddress,
		Value: payment,
	}
}

func NewRandomEth() Eth {
	return Eth{
		Dest:  common.RandAddress(),
		Value: common.RandBigInt(),
	}
}

func (e Eth) AsData() []byte {
	data := make([]byte, 0)
	data = append(data, addressData(e.Dest)...)
	data = append(data, math.U256Bytes(e.Value)...)
	return data
}

func (e Eth) Type() inbox.Type {
	return EthType
}

type EthDepositTx struct {
	L2Message
}

func NewEthDepositTxFromData(data []byte) EthDepositTx {
	return EthDepositTx{L2Message: L2Message{Data: data}}
}

func (e EthDepositTx) AsData() []byte {
	return e.L2Message.Data
}

func (e EthDepositTx) Type() inbox.Type {
	return EthDepositTxType
}
