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

type BisectionChallenge struct {
	*Challenge
}

func NewBisectionChallenge(address common.Address, client arbbridge.ArbClient) (*BisectionChallenge, error) {
	challenge, err := NewChallenge(address, client)
	if err != nil {
		return nil, err
	}
	vm := &BisectionChallenge{
		Challenge: challenge,
	}
	//err = vm.setupContracts()
	return vm, nil
}

//func (c *bisectionChallenge) setupContracts() error {
//	challengeManagerContract, err := executionchallenge.NewBisectionChallenge(c.address, c.Client)
//	if err != nil {
//		return errors2.Wrap(err, "Failed to connect to ChallengeManager")
//	}
//
//	c.bisectionChallenge = challengeManagerContract
//	return nil
//}

func (c *BisectionChallenge) StartConnection(ctx context.Context, startHeight *common.TimeBlocks, startLogIndex uint, outChan chan arbbridge.Event, errChan chan error) error {
	if err := c.Challenge.StartConnection(ctx, startHeight, startLogIndex, outChan, errChan); err != nil {
		return err
	}
	//if err := c.setupContracts(); err != nil {
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
	//		continuedChallengeID,
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

//func (c *bisectionChallenge) processEvents(ctx context.Context, log types.Log, outChan chan arbbridge.Notification) error {
//	header, err := c.Client.HeaderByHash(ctx, log.BlockHash)
//	if err != nil {
//		return err
//	}
//
//	if log.Topics[0] == continuedChallengeID {
//		contChal, err := c.bisectionChallenge.ParseContinued(log)
//		if err != nil {
//			return err
//		}
//		outChan <- arbbridge.Notification{
//			Header: header,
//			VMID:   c.address,
//			Event: arbbridge.ContinueChallengeEvent{
//				SegmentIndex: contChal.SegmentIndex,
//				Deadline:     structures.TimeTicks{Val: contChal.DeadlineTicks},
//			},
//			TxHash: log.TxHash,
//		}
//	}
//	return nil
//}

func (c *BisectionChallenge) ChooseSegment(
	ctx context.Context,
	segmentToChallenge uint16,
	segments []common.Hash,
) error {
	//tree := NewMerkleTree(segments)
	//c.auth.Context = ctx
	//tx, err := c.bisectionChallenge.ChooseSegment(
	//	c.auth,
	//	big.NewInt(int64(segmentToChallenge)),
	//	tree.GetProofFlat(int(segmentToChallenge)),
	//	tree.GetRoot(),
	//	tree.GetNode(int(segmentToChallenge)),
	//)
	//if err != nil {
	//	return err
	//}
	//return c.waitForReceipt(ctx, tx, "ChooseSegment")
	return nil
}
