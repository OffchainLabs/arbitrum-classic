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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

type ExecutionChallenge struct {
	*BisectionChallenge
}

func NewExecutionChallenge(address common.Address, client arbbridge.ArbClient) (*ExecutionChallenge, error) {
	bisectionChallenge, err := NewBisectionChallenge(address, client)
	if err != nil {
		return nil, err
	}
	vm := &ExecutionChallenge{BisectionChallenge: bisectionChallenge}
	//err = vm.setupContracts()
	return vm, err
}

func (c *ExecutionChallenge) setupContracts() error {
	//	challengeManagerContract, err := executionchallenge.NewExecutionChallenge(c.address, c.Client)
	//	if err != nil {
	//		return errors2.Wrap(err, "Failed to connect to ChallengeManager")
	//	}
	//
	//	c.challenge = challengeManagerContract
	return nil
}

func (c *ExecutionChallenge) StartConnection(ctx context.Context, outChan chan arbbridge.Notification, errChan chan error) error {
	if err := c.BisectionChallenge.StartConnection(ctx, outChan, errChan); err != nil {
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
	//		bisectedAssertionID,
	//		oneStepProofCompletedID,
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

//func (c *ExecutionChallenge) processEvents(ctx context.Context, log types.Log, outChan chan arbbridge.Notification) error {
//	event, err := func() (arbbridge.Event, error) {
//		if log.Topics[0] == bisectedAssertionID {
//			bisectChal, err := c.challenge.ParseBisectedAssertion(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.ExecutionBisectionEvent{
//				Assertions: translateBisectionEvent(bisectChal),
//				TotalSteps: bisectChal.TotalSteps,
//				Deadline:   structures.TimeTicks{Val: bisectChal.DeadlineTicks},
//			}, nil
//		} else if log.Topics[0] == oneStepProofCompletedID {
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

func (c *ExecutionChallenge) BisectAssertion(
	ctx context.Context,
	precondition *valprotocol.Precondition,
	assertions []*valprotocol.ExecutionAssertionStub,
	totalSteps uint32,
) error {
	//machineHashes := make([][32]byte, 0, len(assertions)+1)
	//didInboxInsns := make([]bool, 0, len(assertions))
	//messageAccs := make([][32]byte, 0, len(assertions)+1)
	//logAccs := make([][32]byte, 0, len(assertions)+1)
	//gasses := make([]uint64, 0, len(assertions))
	//machineHashes = append(machineHashes, precondition.BeforeHash)
	//messageAccs = append(messageAccs, assertions[0].FirstMessageHashValue())
	//logAccs = append(logAccs, assertions[0].FirstLogHashValue())
	//for _, assertion := range assertions {
	//	machineHashes = append(machineHashes, assertion.AfterHashValue())
	//	didInboxInsns = append(didInboxInsns, assertion.DidInboxInsn)
	//	messageAccs = append(messageAccs, assertion.LastMessageHashValue())
	//	logAccs = append(logAccs, assertion.LastLogHashValue())
	//	gasses = append(gasses, assertion.NumGas)
	//}
	//c.auth.Context = ctx
	//tx, err := c.challenge.BisectAssertion(
	//	c.auth,
	//	precondition.BeforeHash,
	//	precondition.TimeBounds.AsIntArray(),
	//	machineHashes,
	//	didInboxInsns,
	//	messageAccs,
	//	logAccs,
	//	gasses,
	//	totalSteps,
	//)
	//if err != nil {
	//	return err
	//}
	//return c.waitForReceipt(ctx, tx, "BisectAssertion")
	return nil
}

func (c *ExecutionChallenge) OneStepProof(
	ctx context.Context,
	precondition *valprotocol.Precondition,
	assertion *valprotocol.ExecutionAssertionStub,
	proof []byte,
) error {
	//c.auth.Context = ctx
	//tx, err := c.challenge.OneStepProof(
	//	c.auth,
	//	precondition.BeforeHash,
	//	precondition.BeforeInbox.Hash(),
	//	precondition.TimeBounds.AsIntArray(),
	//	assertion.AfterHashValue(),
	//	assertion.DidInboxInsn,
	//	assertion.FirstMessageHashValue(),
	//	assertion.LastMessageHashValue(),
	//	assertion.FirstLogHashValue(),
	//	assertion.LastLogHashValue(),
	//	assertion.NumGas,
	//	proof,
	//)
	//if err != nil {
	//	return err
	//}
	//return c.waitForReceipt(ctx, tx, "OneStepProof")
	return nil
}

func (c *ExecutionChallenge) ChooseSegment(
	ctx context.Context,
	assertionToChallenge uint16,
	preconditions []*valprotocol.Precondition,
	assertions []*valprotocol.ExecutionAssertionStub,
	totalSteps uint32,
) error {
	//bisectionHashes := make([][32]byte, 0, len(assertions))
	//for i := range assertions {
	//	bisectionHash := [32]byte{}
	//	copy(bisectionHash[:], solsha3.SoliditySHA3(
	//		solsha3.Bytes32(preconditions[i].Hash()),
	//		solsha3.Bytes32(assertions[i].Hash()),
	//	))
	//	bisectionHashes = append(bisectionHashes, bisectionHash)
	//}
	//return c.bisectionChallenge.ChooseSegment(
	//	ctx,
	//	assertionToChallenge,
	//	bisectionHashes,
	//)
	return nil
}
