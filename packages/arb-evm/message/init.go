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
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

type Init struct {
	protocol.ChainParams
	Owner       common.Address
	ExtraConfig []byte
}

func NewInitMessage(params protocol.ChainParams, owner common.Address, config []ChainConfigOption) (Init, error) {
	data := make([]byte, 0)
	for _, item := range config {
		data = append(data, item.AsData()...)
	}
	return Init{
		ChainParams: params,
		Owner:       owner,
		ExtraConfig: data,
	}, nil
}

var challengePeriodParamId = hashing.SoliditySHA3([]byte("ChallengePeriodEthBlocks"))
var speedLimitParamId = hashing.SoliditySHA3([]byte("SpeedLimitPerSecond"))
var chainOwnerParamId = hashing.SoliditySHA3([]byte("ChainOwner"))

func NewInitFromData(data []byte) (Init, error) {
	paramId, data := extractUInt256(data)
	if paramId != new(big.Int).SetBytes(challengePeriodParamId[:]) {
		return Init{}, errors.New("Unexpected challenge period parameter id in init message")
	}
	gracePeriod, data := extractUInt256(data)
	paramId, data = extractUInt256(data)
	if paramId != new(big.Int).SetBytes(speedLimitParamId[:]) {
		return Init{}, errors.New("Unexpected speed limit parameter id in init message")
	}
	arbGasSpeedLimit, data := extractUInt256(data)
	paramId, data = extractUInt256(data)
		if paramId != new(big.Int).SetBytes(chainOwnerParamId[:]) {
		return Init{}, errors.New("Unexpected owner parameter id in init message")
	}
	owner, data := extractAddress(data)
	return Init{
		ChainParams: protocol.ChainParams{
			GracePeriod:               common.NewTimeBlocks(gracePeriod),
			ArbGasSpeedLimitPerSecond: arbGasSpeedLimit.Uint64(),
		},
		Owner:       owner,
		ExtraConfig: data,
	}, nil
}

func (m Init) Type() inbox.Type {
	return InitType
}

func (m Init) AsData() []byte {
	data := make([]byte, 0)
 	data = append(data, challengePeriodParamId[:]...)
	data = append(data, math.U256Bytes(m.GracePeriod.AsInt())...)
	data = append(data, speedLimitParamId[:]...)
	data = append(data, math.U256Bytes(new(big.Int).SetUint64(m.ArbGasSpeedLimitPerSecond))...)
	data = append(data, chainOwnerParamId[:]...)
	data = append(data, addressData(m.Owner)...)
	data = append(data, m.ExtraConfig...)
	return data
}

type ChainConfigOption interface {
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

func (c FeeConfig) AsData() []byte {
	data := make([]byte, 0)
	data = append(data, speedLimitParamId[:]...)
	data = append(data, math.U256Bytes(c.SpeedLimitPerSecond)...)
	data = append(data, hashing.SoliditySHA3([]byte("L1GasPerL1CalldataUnit")).Bytes()...)
	data = append(data, math.U256Bytes(c.L1GasPerL2Calldata)...)
	data = append(data, hashing.SoliditySHA3([]byte("L1GasPerStorage")).Bytes()...)
	data = append(data, math.U256Bytes(c.L1GasPerStorage)...)
	data = append(data, hashing.SoliditySHA3([]byte("ArbGasDivisor")).Bytes()...)
	data = append(data, math.U256Bytes(c.ArbGasDivisor)...)
	data = append(data, hashing.SoliditySHA3([]byte("NetworkFeeRecipient")).Bytes()...)
	data = append(data, addressData(c.NetFeeRecipient)...)
	data = append(data, hashing.SoliditySHA3([]byte("CongestionFeeRecipient")).Bytes()...)
	data = append(data, addressData(c.CongestionFeeRecipient)...)
	return data
}

type DefaultAggConfig struct {
	Aggregator common.Address
}

func (c DefaultAggConfig) AsData() []byte {
	var data []byte
	data = append(data, hashing.SoliditySHA3([]byte("DefaultAggregator")).Bytes()...)
	data = append(data, addressData(c.Aggregator)...)
	return data
}

type ChainIDConfig struct {
	ChainId *big.Int
}

func (c ChainIDConfig) AsData() []byte {
	var data []byte
	data = append(data, hashing.SoliditySHA3([]byte("ChainID")).Bytes()...)
	data = append(data, math.U256Bytes(c.ChainId)...)
	return data
}
