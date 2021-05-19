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

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var (
	txCountABI      abi.Method
	withdrawEthABI  abi.Method
	getStorageAtABI abi.Method
	arbOSVersionABI abi.Method

	ethWithdrawal ethcommon.Hash

	arbsysConn *bind.BoundContract

	L2ToL1TransactionID ethcommon.Hash
)

func init() {
	arbsys, err := abi.JSON(strings.NewReader(arboscontracts.ArbSysABI))
	if err != nil {
		panic(err)
	}

	txCountABI = arbsys.Methods["getTransactionCount"]
	withdrawEthABI = arbsys.Methods["withdrawEth"]
	getStorageAtABI = arbsys.Methods["getStorageAt"]
	arbOSVersionABI = arbsys.Methods["arbOSVersion"]

	ethWithdrawal = arbsys.Events["EthWithdrawal"].ID
	L2ToL1TransactionID = arbsys.Events["L2ToL1Transaction"].ID

	arbsysConn = bind.NewBoundContract(ARB_SYS_ADDRESS, arbsys, nil, nil, nil)
}

func TransactionCountData(address common.Address) []byte {
	return makeFuncData(txCountABI, address)
}

func ParseTransactionCountResult(data []byte) (*big.Int, error) {
	vals, err := txCountABI.Outputs.UnpackValues(data)
	if err != nil {
		return nil, err
	}
	val, ok := vals[0].(*big.Int)
	if !ok {
		return nil, errors.New("unexpected tx result")
	}
	return val, nil
}

func WithdrawEthData(address common.Address) []byte {
	return makeFuncData(withdrawEthABI, address)
}

func StorageAtData(address common.Address, index *big.Int) []byte {
	return makeFuncData(getStorageAtABI, address, index)
}

func ParseGetStorageAtResult(data []byte) (*big.Int, error) {
	vals, err := getStorageAtABI.Outputs.UnpackValues(data)
	if err != nil {
		return nil, err
	}
	val, ok := vals[0].(*big.Int)
	if !ok {
		return nil, errors.New("unexpected tx result")
	}
	return val, nil
}

func ArbOSVersionData() []byte {
	return makeFuncData(arbOSVersionABI)
}

func ParseArbOSVersionResult(data []byte) (*big.Int, error) {
	vals, err := arbOSVersionABI.Outputs.UnpackValues(data)
	if err != nil {
		return nil, err
	}
	val, ok := vals[0].(*big.Int)
	if !ok {
		return nil, errors.New("unexpected tx result")
	}
	return val, nil
}

func makeFuncData(funcABI abi.Method, params ...interface{}) []byte {
	txData, err := funcABI.Inputs.Pack(params...)
	if err != nil {
		panic(err)
	}
	return append(funcABI.ID, txData...)
}
