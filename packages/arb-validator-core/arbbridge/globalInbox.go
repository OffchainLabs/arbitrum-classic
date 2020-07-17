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
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type GlobalInboxWatcher interface {
	ContractWatcher

	GetDeliveredEvents(
		ctx context.Context,
		fromBlock *big.Int,
		toBlock *big.Int,
	) ([]MessageDeliveredEvent, error)

	GetERC20Balance(
		ctx context.Context,
		user common.Address,
		tokenContract common.Address,
	) (*big.Int, error)

	GetEthBalance(
		ctx context.Context,
		user common.Address,
	) (*big.Int, error)
}

type GlobalInbox interface {
	GlobalInboxWatcher

	SendL2Message(
		ctx context.Context,
		chain common.Address,
		data message.L2Message,
	) (MessageDeliveredEvent, error)

	// SendL2MessageNoWait calls SendL2Message without
	// blocking while waiting for the receipt. This behavior is different from
	// the other ArbBridge methods. At some point other methods should be
	// updated to behave this way once we can be confident that it will not
	// create any security problems
	SendL2MessageNoWait(
		ctx context.Context,
		chain common.Address,
		data message.L2Message,
	) error

	DepositEthMessage(
		ctx context.Context,
		chain common.Address,
		destination common.Address,
		value *big.Int,
	) error
	DepositERC20Message(
		ctx context.Context,
		chain common.Address,
		tokenAddress common.Address,
		destination common.Address,
		value *big.Int,
	) error
	DepositERC721Message(
		ctx context.Context,
		chain common.Address,
		tokenAddress common.Address,
		destination common.Address,
		value *big.Int,
	) error
}
