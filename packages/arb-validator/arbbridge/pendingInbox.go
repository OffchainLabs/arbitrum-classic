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
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

type PendingInbox interface {
	SendMessage(
		auth *bind.TransactOpts,
		msg valprotocol.Message,
	) error
	ForwardMessage(
		auth *bind.TransactOpts,
		msg valprotocol.Message,
		sig []byte,
	) error
	SendEthMessage(
		auth *bind.TransactOpts,
		data value.Value,
		destination common.Address,
		amount *big.Int,
	) (uint64, error)
	DepositFunds(auth *bind.TransactOpts, amount *big.Int, dest common.Address) error
	GetTokenBalance(
		auth *bind.CallOpts,
		user common.Address,
		tokenContract common.Address,
	) (*big.Int, error)
}
