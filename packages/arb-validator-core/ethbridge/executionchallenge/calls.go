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

package executionchallenge

import (
	"bytes"
	"context"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum/common"
)

func (_BisectionChallenge *BisectionChallengeTransactor) ChooseSegmentCall(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, _segmentToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) error {
	return callCheck(ctx, client, from, contractAddress, "chooseSegment", _segmentToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

func (_ExecutionChallenge *ExecutionChallengeTransactor) BisectAssertionCall(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, _beforeInbox [32]byte, _timeBoundsBlocks [2]*big.Int, _machineHashes [][32]byte, _didInboxInsns []bool, _messageAccs [][32]byte, _logAccs [][32]byte, _gases []uint64, _totalSteps uint64) error {
	return callCheck(ctx, client, from, contractAddress, "bisectAssertion", _beforeInbox, _timeBoundsBlocks, _machineHashes, _didInboxInsns, _messageAccs, _logAccs, _gases, _totalSteps)
}

func (_ExecutionChallenge *ExecutionChallengeTransactor) OneStepProofCall(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, _beforeHash [32]byte, _beforeInbox [32]byte, _beforeInboxValueSize *big.Int, _timeBoundsBlocks [2]*big.Int, _afterHash [32]byte, _didInboxInsns bool, _firstMessage [32]byte, _lastMessage [32]byte, _firstLog [32]byte, _lastLog [32]byte, _gas uint64, _proof []byte) error {
	return callCheck(ctx, client, from, contractAddress, "oneStepProof", _beforeHash, _beforeInbox, _beforeInboxValueSize, _timeBoundsBlocks, _afterHash, _didInboxInsns, _firstMessage, _lastMessage, _firstLog, _lastLog, _gas, _proof)
}

func (_Challenge *ChallengeTransactor) TimeoutChallengeCall(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address) error {
	return callCheck(ctx, client, from, contractAddress, "timeoutChallenge")
}

func callCheck(ctx context.Context, client *ethclient.Client, from common.Address, contractAddress common.Address, method string, params ...interface{}) error {
	contractABI, err := abi.JSON(bytes.NewReader([]byte(ExecutionChallengeABI)))
	if err != nil {
		return err
	}
	return ethutils.CallCheck(ctx, client, from, contractAddress, contractABI, method, params...)
}
