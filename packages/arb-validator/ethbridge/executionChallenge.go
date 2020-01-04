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

package ethbridge

import (
	"context"
	"strings"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/executionchallenge"
)

var bisectedAssertionID common.Hash
var oneStepProofCompletedID common.Hash

func init() {
	parsed, err := abi.JSON(strings.NewReader(executionchallenge.ExecutionChallengeABI))
	if err != nil {
		panic(err)
	}
	bisectedAssertionID = parsed.Events["BisectedAssertion"].ID()
	oneStepProofCompletedID = parsed.Events["OneStepProofCompleted"].ID()
}

type ExecutionChallenge struct {
	*BisectionChallenge
	Challenge *executionchallenge.ExecutionChallenge
}

func NewExecutionChallenge(address common.Address, client *ethclient.Client) (*ExecutionChallenge, error) {
	bisectionChallenge, err := NewBisectionChallenge(address, client)
	if err != nil {
		return nil, err
	}
	vm := &ExecutionChallenge{BisectionChallenge: bisectionChallenge}
	err = vm.setupContracts()
	return vm, err
}

func (c *ExecutionChallenge) setupContracts() error {
	challengeManagerContract, err := executionchallenge.NewExecutionChallenge(c.address, c.Client)
	if err != nil {
		return errors2.Wrap(err, "Failed to connect to ChallengeManager")
	}

	c.Challenge = challengeManagerContract
	return nil
}

func (c *ExecutionChallenge) StartConnection(ctx context.Context, outChan chan Notification, errChan chan error) error {
	if err := c.BisectionChallenge.StartConnection(ctx, outChan, errChan); err != nil {
		return err
	}
	if err := c.setupContracts(); err != nil {
		return err
	}
	header, err := c.Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	filter := ethereum.FilterQuery{
		Addresses: []common.Address{c.address},
		Topics: [][]common.Hash{{
			bisectedAssertionID,
			oneStepProofCompletedID,
		}},
	}

	logs, err := c.Client.FilterLogs(ctx, filter)
	if err != nil {
		return err
	}
	for _, log := range logs {
		if err := c.processEvents(ctx, log, outChan); err != nil {
			return err
		}
	}

	filter.FromBlock = header.Number
	logChan := make(chan types.Log)
	logSub, err := c.Client.SubscribeFilterLogs(ctx, filter, logChan)
	if err != nil {
		return err
	}

	go func() {
		defer logSub.Unsubscribe()

		for {
			select {
			case <-ctx.Done():
				break
			case log := <-logChan:
				if err := c.processEvents(ctx, log, outChan); err != nil {
					errChan <- err
					return
				}
			case err := <-logSub.Err():
				errChan <- err
				return
			}
		}
	}()
	return nil
}

func (c *ExecutionChallenge) processEvents(ctx context.Context, log types.Log, outChan chan Notification) error {
	event, err := func() (Event, error) {
		if log.Topics[0] == bisectedAssertionID {
			bisectChal, err := c.Challenge.ParseBisectedAssertion(log)
			if err != nil {
				return nil, err
			}
			return ExecutionBisectionEvent{
				Assertions:    translateBisectionEvent(bisectChal),
				TotalSteps:    bisectChal.TotalSteps,
				DeadlineTicks: bisectChal.DeadlineTicks,
			}, nil
		} else if log.Topics[0] == oneStepProofCompletedID {
			_, err := c.Challenge.ParseOneStepProofCompleted(log)
			if err != nil {
				return nil, err
			}
			return OneStepProofEvent{}, nil
		}
		return nil, errors2.New("unknown arbitrum event type")
	}()

	if err != nil {
		return err
	}

	header, err := c.Client.HeaderByHash(ctx, log.BlockHash)
	if err != nil {
		return err
	}
	outChan <- Notification{
		Header: header,
		VMID:   c.address,
		Event:  event,
		TxHash: log.TxHash,
	}
	return nil
}

func (c *ExecutionChallenge) BisectAssertion(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertions []*protocol.AssertionStub,
	totalSteps uint32,
) (*types.Receipt, error) {
	machineHashes := make([][32]byte, 0, len(assertions)+1)
	didInboxInsns := make([]bool, 0, len(assertions))
	messageAccs := make([][32]byte, 0, len(assertions)+1)
	logAccs := make([][32]byte, 0, len(assertions)+1)
	gasses := make([]uint64, 0, len(assertions))
	machineHashes = append(machineHashes, precondition.BeforeHashValue())
	messageAccs = append(messageAccs, assertions[0].FirstMessageHashValue())
	logAccs = append(logAccs, assertions[0].FirstLogHashValue())
	for _, assertion := range assertions {
		machineHashes = append(machineHashes, assertion.AfterHashValue())
		didInboxInsns = append(didInboxInsns, assertion.DidInboxInsn)
		messageAccs = append(messageAccs, assertion.LastMessageHashValue())
		logAccs = append(logAccs, assertion.LastLogHashValue())
		gasses = append(gasses, assertion.NumGas)
	}
	tx, err := c.Challenge.BisectAssertion(
		auth,
		precondition.BeforeInboxValue(),
		precondition.TimeBounds.AsIntArray(),
		machineHashes,
		didInboxInsns,
		messageAccs,
		logAccs,
		gasses,
		totalSteps,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, c.Client, auth.From, tx, "BisectAssertion")
}

func (c *ExecutionChallenge) OneStepProof(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
	proof []byte,
) (*types.Receipt, error) {
	tx, err := c.Challenge.OneStepProof(
		auth,
		precondition.BeforeHashValue(),
		precondition.BeforeInboxValue(),
		precondition.TimeBounds.AsIntArray(),
		assertion.AfterHashValue(),
		assertion.DidInboxInsn,
		assertion.FirstMessageHashValue(),
		assertion.LastMessageHashValue(),
		assertion.FirstLogHashValue(),
		assertion.LastLogHashValue(),
		assertion.NumGas,
		proof,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, c.Client, auth.From, tx, "OneStepProof")
}

func (c *ExecutionChallenge) ChooseSegment(
	auth *bind.TransactOpts,
	assertionToChallenge uint16,
	preconditions []*protocol.Precondition,
	assertions []*protocol.AssertionStub,
) (*types.Receipt, error) {
	bisectionHashes := make([][32]byte, 0, len(assertions))
	for i := range assertions {
		bisectionHash := [32]byte{}
		copy(bisectionHash[:], solsha3.SoliditySHA3(
			solsha3.Bytes32(preconditions[i].Hash()),
			solsha3.Bytes32(assertions[i].Hash()),
		))
		bisectionHashes = append(bisectionHashes, bisectionHash)
	}
	return c.BisectionChallenge.ChooseSegment(
		auth,
		assertionToChallenge,
		bisectionHashes,
	)
}

func translateBisectionEvent(event *executionchallenge.ExecutionChallengeBisectedAssertion) []*protocol.AssertionStub {
	bisectionCount := len(event.MachineHashes) - 1
	assertions := make([]*protocol.AssertionStub, 0, bisectionCount)
	for i := 0; i < bisectionCount; i++ {
		assertion := &protocol.AssertionStub{
			AfterHash:        value.NewHashBuf(event.MachineHashes[i+1]),
			DidInboxInsn:     event.DidInboxInsns[i],
			NumGas:           event.Gases[i],
			FirstMessageHash: value.NewHashBuf(event.MessageAccs[i]),
			LastMessageHash:  value.NewHashBuf(event.MessageAccs[i+1]),
			FirstLogHash:     value.NewHashBuf(event.LogAccs[i]),
			LastLogHash:      value.NewHashBuf(event.LogAccs[i+1]),
		}
		assertions = append(assertions, assertion)
	}
	return assertions
}
