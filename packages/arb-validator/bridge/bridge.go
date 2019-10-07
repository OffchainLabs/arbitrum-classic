/*
 * Copyright 2019, Offchain Labs, Inc.
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

package bridge

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type Bridge interface {
	ArbVMBridge

	FinalizedUnanimousAssert(
		ctx context.Context,
		newInboxHash [32]byte,
		assertion *protocol.Assertion,
		signatures [][]byte,
	) (*types.Receipt, error)

	PendingUnanimousAssert(
		ctx context.Context,
		newInboxHash [32]byte,
		assertion *protocol.Assertion,
		sequenceNum uint64,
		signatures [][]byte,
	) (*types.Receipt, error)

	ConfirmUnanimousAsserted(
		ctx context.Context,
		newInboxHash [32]byte,
		assertion *protocol.Assertion,
	) (*types.Receipt, error)
}

type ArbVMBridge interface {
	AddedNewMessages(count uint64)

	FinalizedAssertion(
		assertion *protocol.Assertion,
		onChainTxHash []byte,
		signatures [][]byte,
		proposalResults *valmessage.UnanimousUpdateResults,
	)

	PendingDisputableAssert(
		ctx context.Context,
		precondition *protocol.Precondition,
		assertion *protocol.Assertion,
	) (*types.Receipt, error)

	ConfirmDisputableAsserted(
		ctx context.Context,
		precondition *protocol.Precondition,
		assertion *protocol.Assertion,
	) (*types.Receipt, error)

	InitiateChallenge(
		ctx context.Context,
		precondition *protocol.Precondition,
		assertionHash [32]byte,
		numSteps uint32,
	) (*types.Receipt, error)

	BisectAssertionFirst(
		ctx context.Context,
		assertion *protocol.AssertionStub,
		precondition *protocol.Precondition,
		bisections []*protocol.AssertionStub,
	) (*types.Receipt, error)

	BisectAssertionOther(
		ctx context.Context,
		firstAssertion *protocol.AssertionStub,
		secondAssertion *protocol.AssertionStub,
		precondition *protocol.Precondition,
		bisections []*protocol.AssertionStub,
	) (*types.Receipt, error)

	ContinueChallenge(
		ctx context.Context,
		assertionToChallenge uint16,
		precondition *protocol.Precondition,
		totalSteps uint32,
		assertion [32]byte,
		bisections [][32]byte,
	) (*types.Receipt, error)

	OneStepProofFirst(
		ctx context.Context,
		assertion *protocol.AssertionStub,
		precondition *protocol.Precondition,
		proof []byte,
	) (*types.Receipt, error)

	OneStepProofOther(
		ctx context.Context,
		firstAssertion *protocol.AssertionStub,
		secondAssertion *protocol.AssertionStub,
		precondition *protocol.Precondition,
		proof []byte,
	) (*types.Receipt, error)

	AsserterTimedOut(
		ctx context.Context,
	) (*types.Receipt, error)

	ChallengerTimedOut(
		ctx context.Context,
	) (*types.Receipt, error)

	IsPendingUnanimous(
		ctx context.Context,
	) (bool, error)
}
