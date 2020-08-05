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

package ethbridgemachine

import (
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
)

type proofData struct {
	Assertion *valprotocol.ExecutionAssertionStub
	Proof     []byte
	Message   *inbox.InboxMessage
}

func generateProofCases(contract string) ([]*proofData, error) {
	mach, err := loader.LoadMachineFromFile(contract, true, "cpp")
	if err != nil {
		return nil, err
	}

	maxSteps := uint64(100000)
	ms := structures.NewRandomMessageStack(100)

	prevInboxHash := common.Hash{}

	proofs := make([]*proofData, 0)
	for i := uint64(0); i < maxSteps; i++ {
		proof, err := mach.MarshalForProof()
		if err != nil {
			return nil, err
		}
		beforeMach := mach.Clone()
		messages, err := ms.GetMessages(prevInboxHash, 1)
		if err != nil {
			return nil, err
		}
		a, ranSteps := mach.ExecuteAssertion(1, messages, 0)
		if ranSteps == 0 {
			break
		}
		if ranSteps != 1 {
			return nil, errors.New("executed incorrect step count")
		}
		if mach.CurrentStatus() == machine.ErrorStop {
			beforeMach.PrintState()
			mach.PrintState()
			return nil, errors.New("machine stopped in error state")
		}
		stub := structures.NewExecutionAssertionStubFromWholeAssertion(a, prevInboxHash, ms)
		var msg *inbox.InboxMessage
		if a.InboxMessagesConsumed > 0 {
			msg = &messages[0]
		}
		proofs = append(proofs, &proofData{
			Assertion: stub,
			Proof:     proof,
			Message:   msg,
		})
		prevInboxHash = stub.AfterInboxHash
	}
	return proofs, nil
}
