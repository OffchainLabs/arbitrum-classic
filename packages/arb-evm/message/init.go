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
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"math/big"
)

type Init struct {
	protocol.ChainParams
	Owner       common.Address
	ExtraConfig []byte
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
