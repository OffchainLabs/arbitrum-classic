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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/pkg/errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var (
	txCountABI                   abi.Method
	withdrawEthABI               abi.Method
	withdrawERC20ABI             abi.Method
	withdrawERC721ABI            abi.Method
	getStorageAtABI              abi.Method
	addressTableRegisterABI      abi.Method
	addressTableLookupABI        abi.Method
	addressTableAddressExistsABI abi.Method
	addressTableSizeABI          abi.Method
	addressTableLookupIndexABI   abi.Method
	addressTableDecompressABI    abi.Method
	addressTableCompressABI      abi.Method
	registerBLSKeyABI            abi.Method
	getBLSPublicKeyABI           abi.Method
	uploadFunctionTableABI       abi.Method
	functionTableSizeABI         abi.Method
	functionTableGetABI          abi.Method

	ethWithdrawal    ethcommon.Hash
	erc20Withdrawal  ethcommon.Hash
	erc721Withdrawal ethcommon.Hash

	arbsysConn *bind.BoundContract
)

func init() {
	arbsys, err := abi.JSON(strings.NewReader(arboscontracts.ArbSysABI))
	if err != nil {
		panic(err)
	}

	txCountABI = arbsys.Methods["getTransactionCount"]
	withdrawEthABI = arbsys.Methods["withdrawEth"]
	withdrawERC20ABI = arbsys.Methods["withdrawERC20"]
	withdrawERC721ABI = arbsys.Methods["withdrawERC721"]
	getStorageAtABI = arbsys.Methods["getStorageAt"]
	addressTableRegisterABI = arbsys.Methods["addressTable_register"]
	addressTableLookupABI = arbsys.Methods["addressTable_lookup"]
	addressTableAddressExistsABI = arbsys.Methods["addressTable_addressExists"]
	addressTableSizeABI = arbsys.Methods["addressTable_size"]
	addressTableLookupIndexABI = arbsys.Methods["addressTable_lookupIndex"]
	addressTableDecompressABI = arbsys.Methods["addressTable_decompress"]
	addressTableCompressABI = arbsys.Methods["addressTable_compress"]
	registerBLSKeyABI = arbsys.Methods["registerBlsKey"]
	getBLSPublicKeyABI = arbsys.Methods["getBlsPublicKey"]
	uploadFunctionTableABI = arbsys.Methods["uploadFunctionTable"]
	functionTableSizeABI = arbsys.Methods["functionTableSize"]
	functionTableGetABI = arbsys.Methods["functionTableGet"]

	ethWithdrawal = arbsys.Events["EthWithdrawal"].ID
	erc20Withdrawal = arbsys.Events["ERC20Withdrawal"].ID
	erc721Withdrawal = arbsys.Events["ERC721Withdrawal"].ID

	arbsysConn = bind.NewBoundContract(arbos.ARB_SYS_ADDRESS, arbsys, nil, nil, nil)
}

func TransactionCountData(address common.Address) []byte {
	return makeFuncData(txCountABI, address)
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

func WithdrawEthData(address common.Address) []byte {
	return makeFuncData(withdrawEthABI, address)
}

func WithdrawERC20Data(address common.Address, amount *big.Int) []byte {
	return makeFuncData(withdrawERC20ABI, address, amount)
}

func WithdrawERC721Data(address common.Address, id *big.Int) []byte {
	return makeFuncData(withdrawERC721ABI, address, id)
}

func encodeLog(log evm.Log) types.Log {
	return types.Log{
		Address: log.Address.ToEthAddress(),
		Topics:  common.NewEthHashesFromHashes(log.Topics),
		Data:    log.Data,
	}
}

func ParseEthWithdrawalEvent(log evm.Log) (*arboscontracts.ArbSysEthWithdrawal, error) {
	if log.Topics[0].ToEthHash() != ethWithdrawal {
		return nil, errors.New("wrong event type")
	}
	event := new(arboscontracts.ArbSysEthWithdrawal)
	if err := arbsysConn.UnpackLog(event, "EthWithdrawal", encodeLog(log)); err != nil {
		return nil, err
	}
	return event, nil
}

func ParseERC20WithdrawalEvent(log evm.Log) (*arboscontracts.ArbSysERC20Withdrawal, error) {
	if log.Topics[0].ToEthHash() != erc20Withdrawal {
		return nil, errors.New("wrong event type")
	}
	event := new(arboscontracts.ArbSysERC20Withdrawal)
	if err := arbsysConn.UnpackLog(event, "ERC20Withdrawal", encodeLog(log)); err != nil {
		return nil, err
	}
	return event, nil
}

func ParseERC721WithdrawalEvent(log evm.Log) (*arboscontracts.ArbSysERC721Withdrawal, error) {
	if log.Topics[0].ToEthHash() != erc721Withdrawal {
		return nil, errors.New("wrong event type")
	}
	event := new(arboscontracts.ArbSysERC721Withdrawal)
	if err := arbsysConn.UnpackLog(event, "ERC721Withdrawal", encodeLog(log)); err != nil {
		return nil, err
	}
	return event, nil
}

func StorageAtData(address common.Address, index *big.Int) []byte {
	return makeFuncData(getStorageAtABI, address, index)
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

func AddressTableRegisterData(address common.Address) []byte {
	return makeFuncData(addressTableRegisterABI, address)
}

func AddressTableLookupData(address common.Address) []byte {
	return makeFuncData(addressTableLookupABI, address)
}

func AddressTableAddressExistsData(address common.Address) []byte {
	return makeFuncData(addressTableAddressExistsABI, address)
}

func AddressTableSizeData() []byte {
	return makeFuncData(addressTableSizeABI)
}

func AddressTableLookupIndexData(index *big.Int) []byte {
	return makeFuncData(addressTableLookupIndexABI, index)
}

func AddressTableDecompressData(buf []byte, offset *big.Int) []byte {
	return makeFuncData(addressTableDecompressABI, buf, offset)
}

func AddressTableCompressData(address common.Address) []byte {
	return makeFuncData(addressTableCompressABI, address)
}

func RegisterBLSKeyData(x0, x1, y0, y1 *big.Int) []byte {
	return makeFuncData(registerBLSKeyABI, x0, x1, y0, y1)
}

func GetBLSPublicKeyData(address common.Address) []byte {
	return makeFuncData(getBLSPublicKeyABI, address)
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

func makeFuncData(funcABI abi.Method, params ...interface{}) []byte {
	txData, err := funcABI.Inputs.Pack(params...)
	if err != nil {
		panic(err)
	}
	return append(funcABI.ID, txData...)
}
