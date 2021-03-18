/*
* Copyright 2021, Offchain Labs, Inc.
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
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"strings"
)

var (
	RetryCanceledEvent abi.Event
	RetryRedeemedEvent abi.Event
)

func init() {
	parsedABI, err := abi.JSON(strings.NewReader(arboscontracts.ArbRetryableTxABI))
	if err != nil {
		panic(err)
	}

	RetryCanceledEvent = parsedABI.Events["Canceled"]
	RetryRedeemedEvent = parsedABI.Events["Redeemed"]
}
