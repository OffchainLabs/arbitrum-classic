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
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type PendingInbox interface {
	SendTransactionMessage(
		ctx context.Context,
		data value.Value,
		vmAddress common.Address,
		amount *big.Int,
		seqNumber *big.Int,
	) error
	DepositEthMessage(
		ctx context.Context,
		vmAddress common.Address,
		destination common.Address,
		value *big.Int,
	) error
	DepositERC20Message(
		ctx context.Context,
		vmAddress common.Address,
		tokenAddress common.Address,
		destination common.Address,
		value *big.Int,
	) error
	DepositERC721Message(
		ctx context.Context,
		vmAddress common.Address,
		tokenAddress common.Address,
		destination common.Address,
		value *big.Int,
	) error
	DepositFunds(ctx context.Context, amount *big.Int, dest common.Address) error
	GetTokenBalance(
		ctx context.Context,
		user common.Address,
		tokenContract common.Address,
	) (*big.Int, error)
}
