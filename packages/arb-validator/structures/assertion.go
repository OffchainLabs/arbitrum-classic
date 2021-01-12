/*
* Copyright 2019-2020, Offchain Labs, Inc.
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

package structures

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/rs/zerolog/log"
)

var logger = log.With().Caller().Str("component", "structures").Logger()

func NewExecutionAssertionStubFromWholeAssertion(
	a *protocol.ExecutionAssertion,
	beforeInboxHash common.Hash,
	inboxStack *MessageStack,
) *valprotocol.ExecutionAssertionStub {
	return NewExecutionAssertionStubFromAssertion(a, beforeInboxHash, common.Hash{}, common.Hash{}, inboxStack)
}

func NewExecutionAssertionStubFromAssertion(
	a *protocol.ExecutionAssertion,
	beforeInboxHash common.Hash,
	beforeLogsHash common.Hash,
	beforeMessagesHash common.Hash,
	inboxStack *MessageStack,
) *valprotocol.ExecutionAssertionStub {
	// The after inbox hash
	afterInboxHash, ok := inboxStack.itemSkippedAfterHash(beforeInboxHash, a.InboxMessagesConsumed)
	if !ok {
		logger.Fatal().Msg("Assertion consumed more messages then exist")
	}
	return &valprotocol.ExecutionAssertionStub{
		BeforeMachineHash: a.BeforeMachineHash.Unmarshal(),
		AfterMachineHash:  a.AfterMachineHash.Unmarshal(),
		BeforeInboxHash:   beforeInboxHash,
		AfterInboxHash:    afterInboxHash,
		NumGas:            a.NumGas,
		FirstMessageHash:  common.Hash{},
		LastMessageHash:   valprotocol.BufferAccumHash(beforeLogsHash, a.ParseOutMessages()),
		MessageCount:      a.OutMsgsCount,
		FirstLogHash:      common.Hash{},
		LastLogHash:       valprotocol.BytesArrayAccumHash(beforeMessagesHash, a.LogsData, a.LogsCount),
		LogCount:          a.LogsCount,
	}
}
