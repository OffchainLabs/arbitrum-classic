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

package ethbridge

import (
	"bytes"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

func hashSliceToRaw(slice []common.Hash) [][32]byte {
	ret := make([][32]byte, 0, len(slice))
	for _, h := range slice {
		ret = append(ret, h)
	}
	return ret
}

func addressSliceToRaw(slice []common.Address) []ethcommon.Address {
	ret := make([]ethcommon.Address, 0, len(slice))
	for _, a := range slice {
		ret = append(ret, a.ToEthAddress())
	}
	return ret
}

func addressSliceToAddresses(slice []ethcommon.Address) []common.Address {
	ret := make([]common.Address, 0, len(slice))
	for _, a := range slice {
		ret = append(ret, common.NewAddressFromEth(a))
	}
	return ret
}

func hashSliceToHashes(slice [][32]byte) []common.Hash {
	ret := make([]common.Hash, 0, len(slice))
	for _, a := range slice {
		ret = append(ret, a)
	}
	return ret
}

func combineMessages(
	messages []value.Value,
) []byte {
	var messageData bytes.Buffer
	for _, msg := range messages {
		_ = value.MarshalValue(msg, &messageData)
	}
	return messageData.Bytes()
}
