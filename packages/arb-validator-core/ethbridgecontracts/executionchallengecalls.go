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

func (_BisectionChallenge *BisectionChallengeTransactor) ChooseSegmentCall(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address, _segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) error {
	return callCheckExec(ctx, client, from, contractAddress, "chooseSegment", _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

func (_ExecutionChallenge *ExecutionChallengeTransactor) BisectAssertionCall(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address, _beforeInbox [32]byte, _machineHashes [][32]byte, inboxInsnIndex uint32, _messageAccs [][32]byte, _logAccs [][32]byte, _outCounts []uint64, _gases []uint64, _totalSteps uint64) error {
	return callCheckExec(ctx, client, from, contractAddress, "bisectAssertion", _beforeInbox, _machineHashes, inboxInsnIndex, _messageAccs, _logAccs, _outCounts, _gases, _totalSteps)
}

func (_ExecutionChallenge *ExecutionChallengeTransactor) OneStepProofCall(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address, _beforeInbox [32]byte, _beforeInboxValueSize *big.Int, _firstMessage [32]byte, _firstLog [32]byte, _proof []byte) error {
	return callCheckExec(ctx, client, from, contractAddress, "oneStepProof", _beforeInbox, _beforeInboxValueSize, _firstMessage, _firstLog, _proof)
}

func (_Challenge *ChallengeTransactor) TimeoutChallengeCall(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address) error {
	return callCheckExec(ctx, client, from, contractAddress, "timeoutChallenge")
}

func callCheckExec(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address, method string, params ...interface{}) error {
	contractABI, err := abi.JSON(bytes.NewReader([]byte(ExecutionChallengeABI)))
	if err != nil {
		return err
	}
	return ethutils.CallCheck(ctx, client, from, contractAddress, contractABI, method, params...)
}
