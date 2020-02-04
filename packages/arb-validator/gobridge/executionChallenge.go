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

package gobridge

import (
	"context"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

type ExecutionChallenge struct {
	*bisectionChallenge
}

func NewExecutionChallenge(address common.Address, client *MockArbAuthClient) (*ExecutionChallenge, error) {
	fmt.Println("in NewExecutionChallenge")
	bisectionChallenge, err := newBisectionChallenge(address, client)
	if err != nil {
		return nil, err
	}
	// create new execution challenge contract
	//executionContract, err := executionchallenge.NewExecutionChallenge(address, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to ChallengeManager")
	//}
	vm := &ExecutionChallenge{bisectionChallenge: bisectionChallenge}
	//err = vm.setupContracts()
	return vm, err
}

func (c *ExecutionChallenge) setupContracts() error {
	fmt.Println("in ExecutionChallenge setupContracts")
	//	challengeManagerContract, err := executionchallenge.NewExecutionChallenge(c.address, c.Client)
	//	if err != nil {
	//		return errors2.Wrap(err, "Failed to connect to ChallengeManager")
	//	}
	//
	//	c.challenge = challengeManagerContract
	return nil
}

func (vm *ExecutionChallenge) GetEvents(ctx context.Context, blockId *structures.BlockId) ([]arbbridge.Event, error) {
	fmt.Println("in ExecutionChallenge GetEvents")
	return nil, nil
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
	totalSteps uint64,
) error {
	fmt.Println("in ExecutionChallenge BisectAssertion")
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
	fmt.Println("in ExecutionChallenge OneStepProof")
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
	//return nil
	//	bytes32 precondition = Protocol.generatePreconditionHash(
	//		_beforeHash,
	//		_timeBoundsBlocks,
	//		_beforeInbox
	//	);
	//	requireMatchesPrevState(
	//		ChallengeUtils.executionHash(
	//			1,
	//			precondition,
	//			Protocol.generateAssertionHash(
	//				_afterHash,
	//				_didInboxInsns,
	//				_gas,
	//				_firstMessage,
	//				_lastMessage,
	//				_firstLog,
	//				_lastLog
	//	)
	//)
	//);
	//
	//	uint256 correctProof = OneStepProof.validateProof(
	//		_beforeHash,
	//		_timeBoundsBlocks,
	//		_beforeInbox,
	//		_afterHash,
	//		_didInboxInsns,
	//		_firstMessage,
	//		_lastMessage,
	//		_firstLog,
	//		_lastLog,
	//		_gas,
	//		_proof
	//	);
	//
	//	require(correctProof == 0, OSP_PROOF);
	//	emit OneStepProofCompleted();

	// verify precondition
	// run one step
	// verify post condition
	// emit OneStepProofCompleted();
	return nil
}

func (c *ExecutionChallenge) ChooseSegment(
	ctx context.Context,
	assertionToChallenge uint16,
	preconditions []*valprotocol.Precondition,
	assertions []*valprotocol.ExecutionAssertionStub,
	totalSteps uint64,
) error {
	fmt.Println("in ExecutionChallenge ChooseSegment")
	bisectionHashes := make([]common.Hash, 0, len(assertions))
	for i := range assertions {
		stepCount := structures.CalculateBisectionStepCount(uint64(i), uint64(len(assertions)), totalSteps)
		bisectionHashes = append(
			bisectionHashes,
			structures.ExecutionDataHash(stepCount, preconditions[i].Hash(), assertions[i].Hash()),
		)
	}
	return c.bisectionChallenge.chooseSegment(
		ctx,
		assertionToChallenge,
		bisectionHashes,
	)
}
