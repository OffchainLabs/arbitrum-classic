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
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var (
	txCountABI abi.Method
	txCountSig []byte

	withdrawEthABI abi.Method
	withdrawEthSig []byte

	getStorageAtABI abi.Method
	getStorageAtSig []byte
)

func init() {
	arbsys, err := abi.JSON(strings.NewReader(arboscontracts.ArbSysABI))
	if err != nil {
		panic(err)
	}

	txCountABI = arbsys.Methods["getTransactionCount"]
	txCountSig = hexutil.MustDecode("0x23ca0cd2")

	withdrawEthABI = arbsys.Methods["withdrawEth"]
	withdrawEthSig = hexutil.MustDecode("0x25e16063")

	getStorageAtABI = arbsys.Methods["getStorageAt"]
	getStorageAtSig = hexutil.MustDecode("0xa169625f")
}

func getTransactionCountData(address common.Address) []byte {
	txData, err := txCountABI.Inputs.Pack(address)
	if err != nil {
		panic(err)
	}
	return append(txCountSig, txData...)
}

func parseTransactionCountResult(res *evm.TxResult) (*big.Int, error) {
	vals, err := txCountABI.Outputs.UnpackValues(res.ReturnData)
	if err != nil {
		return nil, err
	}
	val, ok := vals[0].(*big.Int)
	if !ok {
		return nil, errors.New("unexpected tx result")
	}
	return val, nil
}

func GetWithdrawEthData(address common.Address) []byte {
	txData, err := withdrawEthABI.Inputs.Pack(address)
	if err != nil {
		panic(err)
	}
	return append(withdrawEthSig, txData...)
}

func getStorageAtData(address common.Address, index *big.Int) []byte {
	txData, err := getStorageAtABI.Inputs.Pack(address, index)
	if err != nil {
		panic(err)
	}
	return append(getStorageAtSig, txData...)
}

func parseGetStorageAtResult(res *evm.TxResult) (*big.Int, error) {
	vals, err := getStorageAtABI.Outputs.UnpackValues(res.ReturnData)
	if err != nil {
		return nil, err
	}
	val, ok := vals[0].(*big.Int)
	if !ok {
		return nil, errors.New("unexpected tx result")
	}
	return val, nil
}
