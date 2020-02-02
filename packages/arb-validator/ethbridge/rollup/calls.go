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

package rollup

import (
	"bytes"
	"context"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethutils"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum/common"
)

// Solidity: function confirmValid(uint256 deadlineTicks, bytes _messages, bytes32 logsAcc, bytes32 vmProtoStateHash, address[] stakerAddresses, bytes32[] stakerProofs, uint256[] stakerProofOffsets) returns()
func (_ArbRollup *ArbRollupTransactor) ConfirmValidCall(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, deadlineTicks *big.Int, _messages []byte, logsAcc [32]byte, vmProtoStateHash [32]byte, stakerAddresses []common.Address, stakerProofs [][32]byte, stakerProofOffsets []*big.Int) error {
	return callCheck(ctx, client, from, contractAddress, "confirmValid", deadlineTicks, _messages, logsAcc, vmProtoStateHash, stakerAddresses, stakerProofs, stakerProofOffsets)
}

// Solidity: function confirmInvalid(uint256 deadlineTicks, bytes32 challengeNodeData, uint256 branch, bytes32 vmProtoStateHash, address[] stakerAddresses, bytes32[] stakerProofs, uint256[] stakerProofOffsets) returns()
func (_ArbRollup *ArbRollupTransactor) ConfirmInvalidCall(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, deadlineTicks *big.Int, challengeNodeData [32]byte, branch *big.Int, vmProtoStateHash [32]byte, stakerAddresses []common.Address, stakerProofs [][32]byte, stakerProofOffsets []*big.Int) error {
	return callCheck(ctx, client, from, contractAddress, "confirmInvalid", deadlineTicks, challengeNodeData, branch, vmProtoStateHash, stakerAddresses, stakerProofs, stakerProofOffsets)
}

func callCheck(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, method string, params ...interface{}) error {
	contractABI, err := abi.JSON(bytes.NewReader([]byte(ArbRollupABI)))
	if err != nil {
		return err
	}
	return ethutils.CallCheck(ctx, client, from, contractAddress, contractABI, method, params...)
}
