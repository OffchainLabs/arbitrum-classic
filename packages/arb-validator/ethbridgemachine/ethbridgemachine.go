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
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
)

type proofData struct {
	BeforeHash common.Hash
	Assertion  *valprotocol.ExecutionAssertionStub
	InboxInner common.Hash
	InboxSize  int64
	Proof      []byte
}

func generateProofCases(contract string) ([]*proofData, error) {
	mach, err := loader.LoadMachineFromFile(contract, true, "cpp")
	if err != nil {
		return nil, err
	}

	maxSteps := uint64(100000)
	inboxMessages := make([]inbox.InboxMessage, 0)

	proofs := make([]*proofData, 0)
	for i := uint64(0); i < maxSteps; i++ {
		proof, err := mach.MarshalForProof()
		if err != nil {
			return nil, err
		}
		beforeHash := mach.Hash()
		beforeMach := mach.Clone()
		a, ranSteps := mach.ExecuteAssertion(1, inboxMessages, 0)
		if ranSteps == 0 {
			break
		}
		if ranSteps != 1 {
			return nil, errors.New("Executed incorrect step count")
		}
		if mach.CurrentStatus() == machine.ErrorStop {
			beforeMach.PrintState()
			mach.PrintState()
			return nil, errors.New("machine stopped in error state")
		}
		hashPreImage := inbox.InboxValue(inboxMessages).GetPreImage()
		proofs = append(proofs, &proofData{
			BeforeHash: beforeHash,
			Assertion:  valprotocol.NewExecutionAssertionStubFromAssertion(a),
			InboxInner: hashPreImage.GetInnerHash(),
			InboxSize:  hashPreImage.Size(),
			Proof:      proof,
		})

		if a.DidInboxInsn {
			inboxMessages = make([]inbox.InboxMessage, 0)
		}
	}
	return proofs, nil
}
