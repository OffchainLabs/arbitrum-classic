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

func (_ExecutionChallenge *ExecutionChallengeTransactor) BisectAssertionCall(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address, _machineHashes [][32]byte, _inboxHashes [][32]byte, _messageAccs [][32]byte, _logAccs [][32]byte, _outCounts []uint64, _gases []uint64, _totalSteps uint64) error {
	return callCheckExec(ctx, client, from, contractAddress, "bisectAssertion", _machineHashes, _inboxHashes, _messageAccs, _logAccs, _outCounts, _gases, _totalSteps)
}

func (_ExecutionChallenge *ExecutionChallengeTransactor) OneStepProofCall(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address, _inboxHash [32]byte, _firstMessage [32]byte, _firstLog [32]byte, _proof []byte) error {
	return callCheckExec(ctx, client, from, contractAddress, "oneStepProof", _inboxHash, _firstMessage, _firstLog, _proof)
}

func (_ExecutionChallenge *ExecutionChallengeTransactor) OneStepProofBufferCall(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address, _inboxHash [32]byte, _firstMessage [32]byte, _firstLog [32]byte, _proof []byte, _bproof []byte) error {
	return callCheckExec(ctx, client, from, contractAddress, "oneStepProofBuffer", _inboxHash, _firstMessage, _firstLog, _proof, _bproof)
}

func (_ExecutionChallenge *ExecutionChallengeTransactor) OneStepProofInboxCall(ctx context.Context, client ethutils.EthClient, from common.Address, contractAddress common.Address, _firstInbox [32]byte, _firstMessage [32]byte, _firstLog [32]byte, _proof []byte, _kind uint8, _blockNumber *big.Int, _timestamp *big.Int, _sender common.Address, _inboxSeqNum *big.Int, _msgData []byte) error {
	return callCheckExec(ctx, client, from, contractAddress, "oneStepProofInbox", _firstInbox, _firstMessage, _firstLog, _proof, _kind, _blockNumber, _timestamp, _sender, _inboxSeqNum, _msgData)
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
