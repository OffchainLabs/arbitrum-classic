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

package message

import (
	"bytes"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func TestMarshaledBytesHash(t *testing.T) {
	data, err := hexutil.Decode("0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f404142")
	if err != nil {
		t.Fatal(err)
	}
	hash := marshaledBytesHash(data)
	correct := common.HexToHash("0x4fc384a19926e9ff7ec8f2376a0d146dc273031df1db4d133236d209700e4780")
	if hash != correct {
		t.Fatal("incorrect result", hash, correct)
	}
}

func TestDecodeAddressIndex(t *testing.T) {
	rawIndex, err := rlp.EncodeToBytes(big.NewInt(1000))
	if err != nil {
		t.Fatal(err)
	}
	address, err := DecodeAddress(bytes.NewReader(rawIndex))
	if err != nil {
		t.Fatal(err)
	}
	addr, ok := address.(CompressedAddressIndex)
	if !ok {
		t.Fatal("recovered wrong address type")
	}
	if addr.Int.Cmp(big.NewInt(1000)) != 0 {
		t.Error("recovered wrong address index")
	}
}

func TestDecodeAddressFull(t *testing.T) {
	orig := common.HexToAddress("0x81183C9C61bdf79DB7330BBcda47Be30c0a85064")
	rawIndex, err := rlp.EncodeToBytes(orig)
	if err != nil {
		t.Fatal(err)
	}
	address, err := DecodeAddress(bytes.NewReader(rawIndex))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(address)
	addr, ok := address.(CompressedAddressFull)
	if !ok {
		t.Fatalf("recovered wrong address type %T", address)
	}
	if addr.Address != orig {
		t.Error("recovered wrong address")
	}
}

func TestDecodeAddressNone(t *testing.T) {
	rawIndex, err := rlp.EncodeToBytes([]byte{})
	if err != nil {
		t.Fatal(err)
	}
	address, err := DecodeAddress(bytes.NewReader(rawIndex))
	if err != nil {
		t.Fatal(err)
	}
	if address != nil {
		t.Fatalf("recovered wrong address type %T", address)
	}
}
