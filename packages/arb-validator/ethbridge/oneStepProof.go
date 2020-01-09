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

package ethbridge

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/onestepproof"
	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type OneStepProof struct {
	contract *onestepproof.OneStepProof
	client   *ethclient.Client
}

func NewOneStepProof(address common.Address, client *ethclient.Client) (*OneStepProof, error) {
	contract, err := onestepproof.NewOneStepProof(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to OneStepProof")
	}

	return &OneStepProof{contract, client}, nil
}

func (con *OneStepProof) ValidateProof(
	auth *bind.CallOpts,
	precondition *protocol.Precondition,
	assertion *protocol.ExecutionAssertionStub,
	proof []byte,
) (*big.Int, error) {
	return con.contract.ValidateProof(
		auth,
		precondition.BeforeHash,
		precondition.TimeBounds.AsIntArray(),
		precondition.BeforeInbox.Hash(),
		assertion.AfterHashValue(),
		assertion.DidInboxInsn,
		assertion.FirstMessageHashValue(),
		assertion.LastMessageHashValue(),
		assertion.FirstLogHashValue(),
		assertion.LastLogHashValue(),
		assertion.NumGas,
		proof,
	)
}
