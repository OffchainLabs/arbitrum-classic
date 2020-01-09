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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type ArbClient interface {
	NewArbFactory(address common.Address) (ArbFactory, error)
	NewRollup(address common.Address, auth *bind.TransactOpts) (ArbRollup, error)
	NewRollupWatcher(address common.Address) (ArbRollupWatcher, error)
	NewExecutionChallenge(address common.Address, auth *bind.TransactOpts) (ExecutionChallenge, error)
	NewMessagesChallenge(address common.Address, auth *bind.TransactOpts) (MessagesChallenge, error)
	NewOneStepProof(address common.Address) (OneStepProof, error)
	NewPendingInbox(address common.Address) (PendingInbox, error)
	NewPendingTopChallenge(address common.Address, auth *bind.TransactOpts) (PendingTopChallenge, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}
