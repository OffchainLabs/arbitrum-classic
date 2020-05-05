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
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

type MessagesChallenge struct {
	*BisectionChallenge
}

func NewMessagesChallenge(address common.Address, client arbbridge.ArbClient) (*MessagesChallenge, error) {
	bisectionChallenge, err := NewBisectionChallenge(address, client)
	if err != nil {
		return nil, err
	}
	vm := &MessagesChallenge{BisectionChallenge: bisectionChallenge}
	err = vm.setupContracts()
	return vm, err
}

func (c *MessagesChallenge) setupContracts() error {
	//challengeManagerContract, err := messageschallenge.NewMessagesChallenge(c.address, c.Client)
	//if err != nil {
	//	return errors2.Wrap(err, "Failed to connect to MessagesChallenge")
	//}
	//
	//c.challenge = challengeManagerContract
	return nil
}

func (vm *MessagesChallenge) GetEvents(ctx context.Context, blockId *common.BlockId) ([]arbbridge.Event, error) {
	return nil, nil
}

//func (c *MessagesChallenge) processEvents(ctx context.Context, log types.Log, outChan chan arbbridge.Notification) error {
//	event, err := func() (arbbridge.Event, error) {
//		if log.Topics[0] == messagesBisectedID {
//			eventVal, err := c.challenge.ParseBisected(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.MessagesBisectionEvent{
//				ChainHashes:   eventVal.ChainHashes,
//				SegmentHashes: eventVal.SegmentHashes,
//				TotalLength:   eventVal.TotalLength,
//				Deadline:      structures.TimeTicks{Val: eventVal.DeadlineTicks},
//			}, nil
//		} else if log.Topics[0] == messagesOneStepProofCompletedID {
//			_, err := c.challenge.ParseOneStepProofCompleted(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.OneStepProofEvent{}, nil
//		}
//		return nil, errors2.New("unknown arbitrum event type")
//	}()
//
//	if err != nil {
//		return err
//	}
//
//	header, err := c.Client.HeaderByHash(ctx, log.BlockHash)
//	if err != nil {
//		return err
//	}
//	outChan <- arbbridge.Notification{
//		Header: header,
//		VMID:   c.address,
//		Event:  event,
//		TxHash: log.TxHash,
//	}
//	return nil
//}

func (c *MessagesChallenge) Bisect(
	ctx context.Context,
	chainHashes []common.Hash,
	segmentHashes []value.HashOnlyValue,
	chainLength *big.Int,
) error {
	//c.auth.Context = ctx
	//tx, err := c.challenge.Bisect(
	//	c.auth,
	//	chainHashes,
	//	segmentHashes,
	//	chainLength,
	//)
	//if err != nil {
	//	return err
	//}
	//return c.waitForReceipt(ctx, tx, "Bisect")
	return nil
}

func (c *MessagesChallenge) OneStepProofTransactionMessage(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB value.HashOnlyValue,
	msg message.DeliveredTransaction,
) error {
	return nil
}

func (c *MessagesChallenge) OneStepProofEthMessage(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB value.HashOnlyValue,
	msg message.DeliveredEth,
) error {
	return nil
}

func (c *MessagesChallenge) OneStepProofERC20Message(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB value.HashOnlyValue,
	msg message.DeliveredERC20,
) error {
	return nil
}

func (c *MessagesChallenge) OneStepProofERC721Message(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB value.HashOnlyValue,
	msg message.DeliveredERC721,
) error {
	return nil
}

func (c *MessagesChallenge) OneStepProofContractTransactionMessage(
	ctx context.Context,
	lowerHashA common.Hash,
	lowerHashB value.HashOnlyValue,
	msg message.DeliveredContractTransaction,
) error {
	return nil
}

func (c *MessagesChallenge) ChooseSegment(
	ctx context.Context,
	assertionToChallenge uint16,
	chainHashes []common.Hash,
	segmentHashes []common.Hash,
	chainLength *big.Int,
) error {
	return nil
}
