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

package ethbridgetest

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"math/big"
	"testing"
)

func TestBytesToBytestackHash(t *testing.T) {
	datas := [][]byte{
		common.RandBytes(5),
		common.RandBytes(32),
		common.RandBytes(33),
		common.RandBytes(64),
		common.RandBytes(200),
	}
	for _, data := range datas {
		valueHash, err := valueTester.BytesToBytestackHash(nil, data, big.NewInt(0), big.NewInt(int64(len(data))))
		if err != nil {
			t.Fatal(err)
		}
		calcDataValue := message.BytesToByteStack(data)
		if calcDataValue.Hash() != valueHash {
			t.Error("hash not equal with data length", len(data))
		}
	}

}
