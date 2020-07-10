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
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func addressSliceToRaw(slice []common.Address) []ethcommon.Address {
	ret := make([]ethcommon.Address, 0, len(slice))
	for _, a := range slice {
		ret = append(ret, a.ToEthAddress())
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
