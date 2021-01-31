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
	addressTableRegisterABI      abi.Method
	addressTableLookupABI        abi.Method
	addressTableAddressExistsABI abi.Method
	addressTableSizeABI          abi.Method
	addressTableLookupIndexABI   abi.Method
	addressTableDecompressABI    abi.Method
	addressTableCompressABI      abi.Method
)

func init() {
	arbaddresstable, err := abi.JSON(strings.NewReader(arboscontracts.ArbAddressTableABI))
	if err != nil {
		panic(err)
	}

	addressTableRegisterABI = arbaddresstable.Methods["register"]
	addressTableLookupABI = arbaddresstable.Methods["lookup"]
	addressTableAddressExistsABI = arbaddresstable.Methods["addressExists"]
	addressTableSizeABI = arbaddresstable.Methods["size"]
	addressTableLookupIndexABI = arbaddresstable.Methods["lookupIndex"]
	addressTableDecompressABI = arbaddresstable.Methods["decompress"]
	addressTableCompressABI = arbaddresstable.Methods["compress"]
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
