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
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/pkg/errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var (
	uploadFunctionTableABI abi.Method
	functionTableSizeABI   abi.Method
	functionTableGetABI    abi.Method
)

func init() {
	arbfunctable, err := abi.JSON(strings.NewReader(arboscontracts.ArbFunctionTableABI))
	if err != nil {
		panic(err)
	}
	uploadFunctionTableABI = arbfunctable.Methods["upload"]
	functionTableSizeABI = arbfunctable.Methods["size"]
	functionTableGetABI = arbfunctable.Methods["get"]
}

func UploadFunctionTableData(buf []byte) []byte {
	return makeFuncData(uploadFunctionTableABI, buf)
}

func FunctionTableSizeData(address common.Address) []byte {
	return makeFuncData(functionTableSizeABI, address)
}

func FunctionTableGetData(address common.Address, index *big.Int) []byte {
	return makeFuncData(functionTableGetABI, address, index)
}

func ParseFunctionTableGetDataResult(data []byte) (message.FunctionTableEntry, error) {
	failRet := message.FunctionTableEntry{}
	vals, err := functionTableGetABI.Outputs.UnpackValues(data)
	if err != nil {
		return failRet, err
	}
	if len(vals) != 3 {
		return failRet, errors.New("unexpected return param count")
	}
	funcIdRaw, ok := vals[0].(*big.Int)
	if !ok {
		return failRet, errors.New("unexpected type for func id")
	}
	payableRaw, ok := vals[1].(bool)
	if !ok {
		return failRet, errors.New("unexpected type for payable")
	}
	maxGas, ok := vals[2].(*big.Int)
	if !ok {
		return failRet, errors.New("unexpected type for max gas")
	}
	funcBytes := funcIdRaw.Bytes()
	if len(funcBytes) < 4 {
		return failRet, errors.New("unexpected func id length")
	}
	var funcId [4]byte
	copy(funcId[:], funcBytes)
	var payable byte
	if payableRaw {
		payable = 1
	}
	return message.FunctionTableEntry{
		FuncID:  funcId,
		Payable: payable,
		MaxGas:  maxGas,
	}, nil
}
