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
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type GlobalInbox interface {
	SendTransactionMessage(
		ctx context.Context,
		data []byte,
		vmAddress common.Address,
		contactAddress common.Address,
		amount *big.Int,
		seqNumber *big.Int,
	) error

	DeliverTransactionBatch(
		ctx context.Context,
		chain common.Address,
		transactions []message.Transaction,
		signatures [][65]byte,
	) error

	// DeliverTransactionBatchNoWait calls DeliverTransactionBatch without
	// blocking while waiting for the receipt. This behavior is different from
	// the other ArbBridge methods. At some point other methods should be
	// updated to behave this way once we can be confident that it will not
	// create any security problems
	DeliverTransactionBatchNoWait(
		ctx context.Context,
		chain common.Address,
		transactions []message.Transaction,
		signatures [][65]byte,
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
	GetTokenBalance(
		ctx context.Context,
		user common.Address,
		tokenContract common.Address,
	) (*big.Int, error)
}
