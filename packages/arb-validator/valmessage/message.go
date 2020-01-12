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

package valmessage

import (
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

func NewVMConfiguration(
	gracePeriod uint64,
	escrowRequired *big.Int,
	escrowCurrency ethcommon.Address,
	assertKeys []ethcommon.Address,
	maxSteps uint32,
	owner ethcommon.Address,
) *VMConfiguration {
	keys := make([]*valprotocol.AddressBuf, 0, len(assertKeys))
	for _, key := range assertKeys {
		keys = append(keys, valprotocol.NewAddressBuf(key))
	}

	return &VMConfiguration{
		GracePeriod:           gracePeriod,
		EscrowRequired:        common.NewBigIntBuf(escrowRequired),
		EscrowCurrency:        &valprotocol.AddressBuf{Value: escrowCurrency.Bytes()},
		AssertKeys:            keys,
		MaxExecutionStepCount: maxSteps,
		Owner:                 &valprotocol.AddressBuf{Value: owner.Bytes()},
	}
}
