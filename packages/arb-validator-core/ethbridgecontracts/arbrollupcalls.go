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

func (_ArbRollup *ArbRollupTransactor) ConfirmCall(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address, initalProtoStateHash [32]byte, beforeSendCount *big.Int, branches []*big.Int, deadlineTicks []*big.Int, challengeNodeData [][32]byte, logsAcc [][32]byte, vmProtoStateHashes [][32]byte, messagesLengths []*big.Int, messages []byte, stakerAddresses []common.Address, stakerProofs [][32]byte, stakerProofOffsets []*big.Int) error {
	return callCheckRollup(ctx, client, from, contractAddress, "confirm", initalProtoStateHash, beforeSendCount, branches, deadlineTicks, challengeNodeData, logsAcc, vmProtoStateHashes, messagesLengths, messages, stakerAddresses, stakerProofs, stakerProofOffsets)
}

func (_ArbRollup *ArbRollupTransactor) MakeAssertionCall(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address, fields [8][32]byte, fields2 [5]*big.Int, validBlockHashPrecondition [32]byte, validBlockHeightPrecondition *big.Int, messageCount uint64, logCount uint64, prevChildType uint32, numSteps uint64, numArbGas uint64, stakerProof [][32]byte) error {
	return callCheckRollup(ctx, client, from, contractAddress, "makeAssertion", fields, fields2, validBlockHashPrecondition, validBlockHeightPrecondition, messageCount, logCount, prevChildType, numSteps, numArbGas, stakerProof)
}

func callCheckRollup(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address, method string, params ...interface{}) error {
	contractABI, err := abi.JSON(bytes.NewReader([]byte(ArbRollupABI)))
	if err != nil {
		return err
	}
	return ethutils.CallCheck(ctx, client, from, contractAddress, contractABI, method, params...)
}
