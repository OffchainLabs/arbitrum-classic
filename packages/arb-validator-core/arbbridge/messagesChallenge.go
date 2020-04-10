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

package arbbridge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

type MessagesChallenge interface {
	Challenge

	Bisect(
		ctx context.Context,
		chainHashes []common.Hash,
		segmentHashes []value.HashOnlyValue,
		chainLength *big.Int,
	) error

	OneStepProofTransactionMessage(
		ctx context.Context,
		lowerHashA common.Hash,
		lowerHashB value.HashOnlyValue,
		msg message.DeliveredTransaction,
	) error

	OneStepProofEthMessage(
		ctx context.Context,
		lowerHashA common.Hash,
		lowerHashB value.HashOnlyValue,
		msg message.DeliveredEth,
	) error

	OneStepProofERC20Message(
		ctx context.Context,
		lowerHashA common.Hash,
		lowerHashB value.HashOnlyValue,
		msg message.DeliveredERC20,
	) error

	OneStepProofERC721Message(
		ctx context.Context,
		lowerHashA common.Hash,
		lowerHashB value.HashOnlyValue,
		msg message.DeliveredERC721,
	) error

	OneStepProofContractTransactionMessage(
		ctx context.Context,
		lowerHashA common.Hash,
		lowerHashB value.HashOnlyValue,
		msg message.DeliveredContractTransaction,
	) error

	ChooseSegment(
		ctx context.Context,
		assertionToChallenge uint16,
		chainHashes []common.Hash,
		segmentHashes []common.Hash,
		chainLength *big.Int,
	) error
}

type MessagesChallengeWatcher interface {
	ContractWatcher
}
