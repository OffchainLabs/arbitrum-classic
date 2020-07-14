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

package ethbridgecontracts

import (
	"bytes"
	"context"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func (_InboxTopChallenge *InboxTopChallengeTransactor) BisectCall(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address, _chainHashes [][32]byte, _chainLength *big.Int) error {
	return callCheckInbox(ctx, client, from, contractAddress, "bisect", _chainHashes, _chainLength)
}

func callCheckInbox(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address, method string, params ...interface{}) error {
	contractABI, err := abi.JSON(bytes.NewReader([]byte(InboxTopChallengeABI)))
	if err != nil {
		return err
	}
	return ethutils.CallCheck(ctx, client, from, contractAddress, contractABI, method, params...)
}
