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

package snapshot

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var (
	registerBLSKeyABI  abi.Method
	getBLSPublicKeyABI abi.Method
)

func init() {
	arbbls, err := abi.JSON(strings.NewReader(arboscontracts.ArbBLSABI))
	if err != nil {
		panic(err)
	}

	registerBLSKeyABI = arbbls.Methods["register"]
	getBLSPublicKeyABI = arbbls.Methods["getPublicKey"]
}

func RegisterBLSKeyData(x0, x1, y0, y1 *big.Int) []byte {
	return makeFuncData(registerBLSKeyABI, x0, x1, y0, y1)
}

func GetBLSPublicKeyData(address common.Address) []byte {
	return makeFuncData(getBLSPublicKeyABI, address)
}
