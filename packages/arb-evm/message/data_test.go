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
