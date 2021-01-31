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
	"github.com/pkg/errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var (
	getBalanceABI abi.Method
	getBalanceSig []byte
	getCodeABI    abi.Method
	getCodeSig    []byte
)

func init() {
	arbinfo, err := abi.JSON(strings.NewReader(arboscontracts.ArbInfoABI))
	if err != nil {
		panic(err)
	}

	getBalanceABI = arbinfo.Methods["getBalance"]
	getBalanceSig, err = hexutil.Decode("0xf8b2cb4f")
	if err != nil {
		panic(err)
	}

	getCodeABI = arbinfo.Methods["getCode"]
	getCodeSig, err = hexutil.Decode("0x7e105ce2")
	if err != nil {
		panic(err)
	}
}

func GetBalanceData(address common.Address) []byte {
	txData, err := getBalanceABI.Inputs.Pack(address)
	if err != nil {
		panic(err)
	}
	return append(getBalanceSig, txData...)
}

func ParseBalanceResult(res *evm.TxResult) (*big.Int, error) {
	vals, err := getBalanceABI.Outputs.UnpackValues(res.ReturnData)
	if err != nil {
		return nil, err
	}
	val, ok := vals[0].(*big.Int)
	if !ok {
		return nil, errors.New("unexpected tx result")
	}
	return val, nil
}

func getCodeData(address common.Address) []byte {
	txData, err := getCodeABI.Inputs.Pack(address)
	if err != nil {
		panic(err)
	}
	return append(getCodeSig, txData...)
}

func parseCodeResult(res *evm.TxResult) ([]byte, error) {
	vals, err := getCodeABI.Outputs.UnpackValues(res.ReturnData)
	if err != nil {
		return nil, err
	}
	val, ok := vals[0].([]byte)
	if !ok {
		return nil, errors.New("unexpected tx result")
	}
	return val, nil
}
