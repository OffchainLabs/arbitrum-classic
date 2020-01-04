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
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type BridgeMessage uint

const (
	ProofAccepted BridgeMessage = iota
)

type Bridge interface {
	ArbVMBridge

	FinalizedUnanimousAssert(
		ctx context.Context,
		newInboxHash [32]byte,
		assertion *protocol.ExecutionAssertion,
		signatures [][]byte,
	) (*types.Receipt, error)

	PendingUnanimousAssert(
		ctx context.Context,
		newInboxHash [32]byte,
		assertion *protocol.ExecutionAssertion,
		sequenceNum uint64,
		signatures [][]byte,
	) (*types.Receipt, error)

	ConfirmUnanimousAsserted(
		ctx context.Context,
		newInboxHash [32]byte,
		assertion *protocol.ExecutionAssertion,
	) (*types.Receipt, error)
}

type ArbVMBridge interface {
	SendMonitorMsg(msg BridgeMessage)
	SendMonitorErr(msg Error)

	FinalizedAssertion(
		assertion *protocol.ExecutionAssertion,
		onChainTxHash []byte,
		signatures [][]byte,
		proposalResults *valmessage.UnanimousUpdateResults,
	)

	PendingDisputableAssert(
		ctx context.Context,
		precondition *protocol.Precondition,
		assertion *protocol.ExecutionAssertion,
	) (*types.Receipt, error)

	ConfirmDisputableAsserted(
		ctx context.Context,
		precondition *protocol.Precondition,
		assertion *protocol.ExecutionAssertion,
	) (*types.Receipt, error)

	InitiateChallenge(
		ctx context.Context,
		precondition *protocol.Precondition,
		assertion *protocol.ExecutionAssertionStub,
	) (*types.Receipt, error)

	IsPendingUnanimous(
		ctx context.Context,
	) (bool, error)

	Challenge(
		ctx context.Context,
		address common.Address,
		precondition *protocol.Precondition,
		machine machine.Machine,
	) error

	DefendChallenge(
		ctx context.Context,
		address common.Address,
		assDef machine.AssertionDefender,
	) error

	ObserveChallenge(
		ctx context.Context,
		address common.Address,
	) error
}

type Challenge interface {
	SendMonitorMsg(msg BridgeMessage)
	SendMonitorErr(msg Error)

	BisectAssertion(
		ctx context.Context,
		precondition *protocol.Precondition,
		assertions []*protocol.ExecutionAssertionStub,
	) (*types.Receipt, error)

	ContinueChallenge(
		ctx context.Context,
		assertionToChallenge uint16,
		preconditions *protocol.Precondition,
		assertions []*protocol.ExecutionAssertionStub,
	) (*types.Receipt, error)

	OneStepProof(
		ctx context.Context,
		precondition *protocol.Precondition,
		assertion *protocol.ExecutionAssertionStub,
		proof []byte,
	) (*types.Receipt, error)

	AsserterTimedOut(
		ctx context.Context,
	) (*types.Receipt, error)

	ChallengerTimedOut(
		ctx context.Context,
	) (*types.Receipt, error)
}

type Error struct {
	Err         error
	Message     string
	Recoverable bool
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%v: Recoverable=%v %v", e.Message, e.Recoverable, e.Err)
	}
	return e.Message
}
