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

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

type Challenge struct {
	*ClientConnection

	address common.Address
	client  arbbridge.ArbClient
}

func NewChallenge(address common.Address, client arbbridge.ArbClient) (*Challenge, error) {
	vm := &Challenge{ClientConnection: &ClientConnection{client}, address: address}
	//err := vm.setupContracts()
	return vm, nil
}

//func (c *challenge) setupContracts() error {
//	challengeManagerContract, err := executionchallenge.NewChallenge(c.address, c.Client)
//	if err != nil {
//		return errors2.Wrap(err, "Failed to connect to ChallengeManager")
//	}
//
//	c.challenge = challengeManagerContract
//	return nil
//}

func (c *Challenge) StartConnection(ctx context.Context, startHeight *common.TimeBlocks, startLogIndex uint, eventChan chan<- arbbridge.Event, errChan chan<- error) error {
	//if err := c.setupContracts(); err != nil {
	//	return err
	//}
	//headers := make(chan *types.Header)
	//headersSub, err := c.Client.SubscribeNewHead(ctx, headers)
	//if err != nil {
	//	return err
	//}
	//
	//header, err := c.Client.HeaderByNumber(ctx, nil)
	//if err != nil {
	//	return err
	//}
	//
	//filter := ethereum.FilterQuery{
	//	Addresses: []common.Address{c.address},
	//	Topics: [][]common.Hash{{
	//		initiatedChallengeID,
	//		timedOutAsserterID,
	//		timedOutChallengerID,
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
	//	defer headersSub.Unsubscribe()
	//	defer logSub.Unsubscribe()
	//
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			break
	//		case header := <-headers:
	//			outChan <- arbbridge.Notification{
	//				Header: header,
	//				Event:  arbbridge.NewTimeEvent{},
	//			}
	//		case log := <-logChan:
	//			if err := c.processEvents(ctx, log, outChan); err != nil {
	//				errChan <- err
	//				return
	//			}
	//		case err := <-headersSub.Err():
	//			errChan <- err
	//			return
	//		case err := <-logSub.Err():
	//			errChan <- err
	//			return
	//		}
	//	}
	//}()
	return nil
}

//func (c *challenge) processEvents(ctx context.Context, log types.Log, outChan chan arbbridge.Notification) error {
//	event, err := func() (arbbridge.Event, error) {
//		if log.Topics[0] == initiatedChallengeID {
//			eventVal, err := c.challenge.ParseInitiatedChallenge(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.InitiateChallengeEvent{
//				Deadline: structures.TimeTicks{Val: eventVal.DeadlineTicks},
//			}, nil
//		} else if log.Topics[0] == timedOutAsserterID {
//			_, err := c.challenge.ParseAsserterTimedOut(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.AsserterTimeoutEvent{}, nil
//		} else if log.Topics[0] == timedOutChallengerID {
//			_, err := c.challenge.ParseChallengerTimedOut(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.ChallengerTimeoutEvent{}, nil
//		}
//		return nil, errors2.New("unknown arbitrum event type")
//	}()
//	if err != nil {
//		return err
//	}
//
//	header, err := c.Client.HeaderByHash(ctx, log.BlockHash)
//	if err != nil {
//		return err
//	}
//
//	outChan <- arbbridge.Notification{
//		Header: header,
//		VMID:   c.address,
//		Event:  event,
//		TxHash: log.TxHash,
//	}
//
//	return nil
//}

func (c *Challenge) TimeoutChallenge(
	ctx context.Context,
) error {
	//c.auth.Context = ctx
	//tx, err := c.challenge.TimeoutChallenge(c.auth)
	//if err != nil {
	//	return err
	//}
	//return c.waitForReceipt(ctx, tx, "TimeoutChallenge")
	return nil
}

//func (c *challenge) waitForReceipt(ctx context.Context, tx *types.Transaction, methodName string) error {
//	return c.ClientConnection.waitForReceipt(ctx, c.auth.From, tx, methodName)
//}
