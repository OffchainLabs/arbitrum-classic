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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"math/big"
	"strings"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/executionchallenge"
)

var continuedChallengeID common.Hash

func init() {
	parsed, err := abi.JSON(strings.NewReader(executionchallenge.BisectionChallengeABI))
	if err != nil {
		panic(err)
	}
	continuedChallengeID = parsed.Events["Continued"].ID()
}

type BisectionChallenge interface {
	Challenge
	ChooseSegment(
		ctx context.Context,
		segmentToChallenge uint16,
		segments [][32]byte,
	) error
}

func NewBisectionChallenge(address common.Address, client *ethclient.Client, auth *bind.TransactOpts) (*BisectionChallenge, error) {
	challenge, err := NewChallenge(address, client, auth)
	if err != nil {
		return nil, err
	}
	vm := &BisectionChallenge{
		Challenge:          challenge,
		BisectionChallenge: nil,
	}
	err = vm.setupContracts()
	return vm, err
}
