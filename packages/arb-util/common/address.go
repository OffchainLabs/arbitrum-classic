/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

package common

import (
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Address [20]byte

var zeroAddress Address

func init() {
	zeroAddress = Address{}
}

func NewAddressFromEth(a ethcommon.Address) Address {
	return Address(a)
}

func NewAddressFromBig(int *big.Int) Address {
	return NewAddressFromEth(ethcommon.BigToAddress(int))
}

func AddressArrayFromEth(addresses []ethcommon.Address) []Address {
	ret := make([]Address, 0, len(addresses))
	for _, a := range addresses {
		ret = append(ret, NewAddressFromEth(a))
	}
	return ret
}

func AddressArrayToEth(addresses []Address) []ethcommon.Address {
	ret := make([]ethcommon.Address, 0, len(addresses))
	for _, a := range addresses {
		ret = append(ret, a.ToEthAddress())
	}
	return ret
}

func (a Address) Bytes() []byte {
	return a[:]
}

func (a Address) String() string {
	return a.Hex()
}

func (a Address) ShortString() string {
	return a.Hex()[2:8]
}

func (a Address) IsZero() bool {
	return a == zeroAddress
}

func (a Address) Equals(a2 Address) bool {
	return a == a2
}

func (a Address) ToEthAddress() ethcommon.Address {
	return ethcommon.Address(a)
}

func (a Address) Hex() string {
	return hexutil.Encode(a[:])
}

func HexToAddress(hex string) Address {
	return NewAddressFromEth(ethcommon.HexToAddress(hex))
}
