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

package mockbridge

import (
	"context"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

type OneStepProof struct {
	//contract *onestepproof.OneStepProof
	client arbbridge.ArbClient
}

func NewOneStepProof(address common.Address, client arbbridge.ArbClient) (*OneStepProof, error) {
	//contract, err := onestepproof.NewOneStepProof(address, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to OneStepProof")
	//}

	return &OneStepProof{client}, nil
}

func (con *OneStepProof) ValidateProof(
	ctx context.Context,
	precondition *valprotocol.Precondition,
	assertion *valprotocol.ExecutionAssertionStub,
	proof []byte,
) (*big.Int, error) {
	//return con.contract.ValidateProof(
	//	auth,
	//	precondition.BeforeHash,
	//	precondition.TimeBounds.AsIntArray(),
	//	precondition.BeforeInbox.Hash(),
	//	assertion.AfterHashValue(),
	//	assertion.DidInboxInsn,
	//	assertion.FirstMessageHashValue(),
	//	assertion.LastMessageHashValue(),
	//	assertion.FirstLogHashValue(),
	//	assertion.LastLogHashValue(),
	//	assertion.NumGas,
	//	proof,
	//)
	return big.NewInt(0), nil
}
