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
	"context"
	"math/big"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge/executionchallenge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type oneStepProof struct {
	contract *executionchallenge.OneStepProof
	client   *ethclient.Client
}

func newOneStepProof(address ethcommon.Address, client *ethclient.Client) (*oneStepProof, error) {
	contract, err := executionchallenge.NewOneStepProof(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to oneStepProof")
	}

	return &oneStepProof{contract, client}, nil
}

func (con *oneStepProof) ValidateProof(
	ctx context.Context,
	precondition *valprotocol.Precondition,
	assertion *valprotocol.ExecutionAssertionStub,
	proof []byte,
) (*big.Int, error) {
	hashPreImage := precondition.BeforeInbox.GetPreImage()
	return con.contract.ValidateProof(
		&bind.CallOpts{Context: ctx},
		precondition.BeforeHash,
		precondition.TimeBounds.AsIntArray(),
		hashPreImage.GetPreImageHash(),
		big.NewInt(hashPreImage.Size()),
		assertion.AfterHash,
		assertion.DidInboxInsn,
		assertion.FirstMessageHash,
		assertion.LastMessageHash,
		assertion.FirstLogHash,
		assertion.LastLogHash,
		assertion.NumGas,
		proof,
	)
}
