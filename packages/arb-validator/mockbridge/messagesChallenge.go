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
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
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

func (c *MessagesChallenge) StartConnection(ctx context.Context, startHeight *common.TimeBlocks, startLogIndex uint, eventChan chan<- arbbridge.Event, errChan chan<- error) error {
	if err := c.BisectionChallenge.StartConnection(ctx, startHeight, startLogIndex, eventChan, errChan); err != nil {
		return err
	}
	if err := c.setupContracts(); err != nil {
		return err
	}
	//header, err := c.Client.HeaderByNumber(ctx, nil)
	//if err != nil {
	//	return err
	//}
	//
	//filter := ethereum.FilterQuery{
	//	Addresses: []common.Address{c.address},
	//	Topics: [][]common.Hash{{
	//		messagesBisectedID,
	//		messagesOneStepProofCompletedID,
	//	}},
	//}
	//
	//logs, err := c.Client.FilterLogs(ctx, filter)
	//if err != nil {
	//	return err
	//}
	//for _, log := range logs {
	//	if err := c.processEvents(ctx, log, outChan); err != nil {
	//		return err
	//	}
	//}
	//
	//filter.FromBlock = header.Number
	//logChan := make(chan types.Log)
	//logSub, err := c.Client.SubscribeFilterLogs(ctx, filter, logChan)
	//if err != nil {
	//	return err
	//}
	//
	//go func() {
	//	defer logSub.Unsubscribe()
	//
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			break
	//		case log := <-logChan:
	//			if err := c.processEvents(ctx, log, outChan); err != nil {
	//				errChan <- err
	//				return
	//			}
	//		case err := <-logSub.Err():
	//			errChan <- err
	//			return
	//		}
	//	}
	//}()
	return nil
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
	segmentHashes []common.Hash,
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

func (c *MessagesChallenge) OneStepProof(
	ctx context.Context,
	lowerHashA common.Hash,
	topHashA common.Hash,
	lowerHashB common.Hash,
	topHashB common.Hash,
	value common.Hash,
) error {
	//c.auth.Context = ctx
	//tx, err := c.challenge.OneStepProof(
	//	c.auth,
	//	lowerHashA,
	//	topHashA,
	//	lowerHashB,
	//	topHashB,
	//	value,
	//)
	//if err != nil {
	//	return err
	//}
	//return c.waitForReceipt(ctx, tx, "OneStepProof")
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
