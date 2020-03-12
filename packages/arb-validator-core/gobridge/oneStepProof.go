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

package gobridge

import (
	"context"
	"errors"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type oneStepProof struct {
	oneStepProofContract common.Address
	client               *goEthdata
}

func newOneStepProof(address common.Address, client *goEthdata) (*oneStepProof, error) {
	if !address.Equals(client.oneStepProof.oneStepProofContract) {
		return nil, errors.New("invalid oneStepProof address")
	}
	return client.oneStepProof, nil
}

func (con *oneStepProof) ValidateProof(
	ctx context.Context,
	precondition *valprotocol.Precondition,
	assertion *valprotocol.ExecutionAssertionStub,
	proof []byte,
) (*big.Int, error) {
	con.client.goEthMutex.Lock()
	defer con.client.goEthMutex.Unlock()
	// execution one step proof
	// for now always return true

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
