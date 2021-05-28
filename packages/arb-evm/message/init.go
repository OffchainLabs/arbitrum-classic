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
	"bytes"
	"encoding/binary"
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

var InitOptionSetChargingParams uint64 = 2
var InitOptionSetDefaultAggregator uint64 = 3
var InitOptionSetChainID uint64 = 4

type Init struct {
	protocol.ChainParams
	Owner       common.Address
	ExtraConfig []byte
}

func NewInitMessage(params protocol.ChainParams, owner common.Address, config []ChainConfigOption) (Init, error) {
	data := make([]byte, 0)
	for _, item := range config {
		itemData := item.AsData()
		optionId := item.OptionCode()
		var w bytes.Buffer
		if err := binary.Write(&w, binary.BigEndian, &optionId); err != nil {
			return Init{}, err
		}
		length := uint64(len(itemData))
		if err := binary.Write(&w, binary.BigEndian, &length); err != nil {
			return Init{}, err
		}
		data = append(data, w.Bytes()...)
		data = append(data, itemData...)
	}
	return Init{
		ChainParams: params,
		Owner:       owner,
		ExtraConfig: data,
	}, nil
}

func NewInitFromData(data []byte) Init {
	gracePeriod, data := extractUInt256(data)
	arbGasSpeedLimit, data := extractUInt256(data)
	maxExecutionSteps, data := extractUInt256(data)
	stakeRequirement, data := extractUInt256(data)
	stakeToken, data := extractAddress(data)
	owner, data := extractAddress(data)
	return Init{
		ChainParams: protocol.ChainParams{
			StakeRequirement:          stakeRequirement,
			StakeToken:                stakeToken,
			GracePeriod:               common.NewTimeBlocks(gracePeriod),
			MaxExecutionSteps:         maxExecutionSteps.Uint64(),
			ArbGasSpeedLimitPerSecond: arbGasSpeedLimit.Uint64(),
		},
		Owner:       owner,
		ExtraConfig: data,
	}
}

func (m Init) Type() inbox.Type {
	return InitType
}

func (m Init) AsData() []byte {
	data := make([]byte, 0)
	data = append(data, math.U256Bytes(m.GracePeriod.AsInt())...)
	data = append(data, math.U256Bytes(new(big.Int).SetUint64(m.ArbGasSpeedLimitPerSecond))...)
	data = append(data, math.U256Bytes(new(big.Int).SetUint64(m.MaxExecutionSteps))...)
	data = append(data, math.U256Bytes(m.StakeRequirement)...)
	data = append(data, addressData(m.StakeToken)...)
	data = append(data, addressData(m.Owner)...)
	data = append(data, m.ExtraConfig...)
	return data
}

type ChainConfigOption interface {
	OptionCode() uint64
	AsData() []byte
}

type FeeConfig struct {
	SpeedLimitPerSecond    *big.Int
	L1GasPerL2Tx           *big.Int
	ArbGasPerL2Tx          *big.Int
	L1GasPerL2Calldata     *big.Int
	ArbGasPerL2Calldata    *big.Int
	L1GasPerStorage        *big.Int
	ArbGasPerStorage       *big.Int
	ArbGasDivisor          *big.Int
	NetFeeRecipient        common.Address
	CongestionFeeRecipient common.Address
}

func (c FeeConfig) OptionCode() uint64 {
	return InitOptionSetChargingParams
}

func (c FeeConfig) AsData() []byte {
	data := make([]byte, 0)
	data = append(data, math.U256Bytes(c.SpeedLimitPerSecond)...)
	data = append(data, math.U256Bytes(c.L1GasPerL2Tx)...)
	data = append(data, math.U256Bytes(c.ArbGasPerL2Tx)...)
	data = append(data, math.U256Bytes(c.L1GasPerL2Calldata)...)
	data = append(data, math.U256Bytes(c.ArbGasPerL2Calldata)...)
	data = append(data, math.U256Bytes(c.L1GasPerStorage)...)
	data = append(data, math.U256Bytes(c.ArbGasPerStorage)...)
	data = append(data, math.U256Bytes(c.ArbGasDivisor)...)
	data = append(data, addressData(c.NetFeeRecipient)...)
	data = append(data, addressData(c.CongestionFeeRecipient)...)
	return data
}

type DefaultAggConfig struct {
	Aggregator common.Address
}

func (c DefaultAggConfig) OptionCode() uint64 {
	return InitOptionSetDefaultAggregator
}

func (c DefaultAggConfig) AsData() []byte {
	return addressData(c.Aggregator)
}

type ChainIDConfig struct {
	ChainId *big.Int
}

func (c ChainIDConfig) OptionCode() uint64 {
	return InitOptionSetChainID
}

func (c ChainIDConfig) AsData() []byte {
	return math.U256Bytes(c.ChainId)
}
