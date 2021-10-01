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

package arbos

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
)

var (
	getPricesInWeiABI abi.Method
)

func init() {
	arbgasinfo, err := abi.JSON(strings.NewReader(arboscontracts.ArbGasInfoABI))
	if err != nil {
		panic(err)
	}

	getPricesInWeiABI = arbgasinfo.Methods["getPricesInWei"]
}

func GetPricesInWeiData() []byte {
	return makeFuncData(getPricesInWeiABI)
}

func ParseGetPricesInWeiResult(data []byte) ([6]*big.Int, error) {
	rawValues, err := getPricesInWeiABI.Outputs.UnpackValues(data)
	if err != nil {
		return [6]*big.Int{}, err
	}
	var values [6]*big.Int
	for i, rawVal := range rawValues {
		val, ok := rawVal.(*big.Int)
		if !ok {
			return [6]*big.Int{}, errors.New("unexpected tx result")
		}
		values[i] = val
	}
	return values, nil
}
