/*
 * Copyright 2019, Offchain Labs, Inc.
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
	"math/big"
	"strings"

	errors2 "github.com/pkg/errors"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/challenge"
)

var initiatedChallengeID common.Hash
var bisectedAssertionID common.Hash
var timedOutChallengeID common.Hash
var continuedChallengeID common.Hash
var oneStepProofCompletedID common.Hash

func init() {
	parsed, err := abi.JSON(strings.NewReader(challenge.ChallengeABI))
	if err != nil {
		panic(err)
	}
	initiatedChallengeID = parsed.Events["InitiatedChallenge"].ID()
	bisectedAssertionID = parsed.Events["BisectedAssertion"].ID()
	timedOutChallengeID = parsed.Events["TimedOutChallenge"].ID()
	continuedChallengeID = parsed.Events["ContinuedChallenge"].ID()
	oneStepProofCompletedID = parsed.Events["OneStepProofCompleted"].ID()
}

type Challenge struct {
	OutChan   chan Notification
	ErrChan   chan error
	Client    *ethclient.Client
	Challenge *challenge.Challenge

	address common.Address
	client  *ethclient.Client
}

func NewChallenge(address common.Address, client *ethclient.Client) (*Challenge, error) {
	outChan := make(chan Notification, 1024)
	errChan := make(chan error, 1024)
	vm := &Challenge{OutChan: outChan, ErrChan: errChan, Client: client, address: address}
	err := vm.setupContracts()
	return vm, err
}

func (c *Challenge) setupContracts() error {
	challengeManagerContract, err := challenge.NewChallenge(c.address, c.Client)
	if err != nil {
		return errors2.Wrap(err, "Failed to connect to ChallengeManager")
	}

	c.Challenge = challengeManagerContract
	return nil
}

func (c *Challenge) GetChans() (chan Notification, chan error) {
	return c.OutChan, c.ErrChan
}

func (c *Challenge) Close() {
	close(c.OutChan)
	close(c.ErrChan)
}

func (c *Challenge) StartConnection(ctx context.Context) error {
	if err := c.setupContracts(); err != nil {
		return err
	}
	headers := make(chan *types.Header)
	headersSub, err := c.Client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return err
	}

	header, err := c.Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	filter := ethereum.FilterQuery{
		Addresses: []common.Address{c.address},
		Topics: [][]common.Hash{{
			initiatedChallengeID,
			bisectedAssertionID,
			timedOutChallengeID,
			continuedChallengeID,
			oneStepProofCompletedID,
		}},
	}

	logs, err := c.Client.FilterLogs(ctx, filter)
	if err != nil {
		return err
	}
	for _, log := range logs {
		if err := c.processEvents(ctx, log); err != nil {
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
		defer headersSub.Unsubscribe()
		defer logSub.Unsubscribe()

		for {
			select {
			case <-ctx.Done():
				break
			case header := <-headers:
				c.OutChan <- Notification{
					Header: header,
					Event:  NewTimeEvent{},
				}
			case log := <-logChan:
				if err := c.processEvents(ctx, log); err != nil {
					c.ErrChan <- err
					return
				}
			case err := <-headersSub.Err():
				c.ErrChan <- err
				return
			case err := <-logSub.Err():
				c.ErrChan <- err
				return
			}
		}
	}()
	return nil
}

func (c *Challenge) processEvents(ctx context.Context, log types.Log) error {
	header, err := c.Client.HeaderByHash(ctx, log.BlockHash)
	if err != nil {
		return err
	}

	if log.Topics[0] == initiatedChallengeID {
		initChal, err := c.Challenge.ParseInitiatedChallenge(log)
		if err != nil {
			return err
		}
		c.OutChan <- Notification{
			Header: header,
			VMID:   c.address,
			Event: InitiateChallengeEvent{
				Deadline: initChal.Deadline,
			},
			TxHash: log.TxHash,
		}
	} else if log.Topics[0] == bisectedAssertionID {
		bisectChal, err := c.Challenge.ParseBisectedAssertion(log)
		if err != nil {
			return err
		}
		c.OutChan <- Notification{
			Header: header,
			VMID:   c.address,
			Event: BisectionEvent{
				Assertions: translateBisectionEvent(bisectChal),
				Deadline:   bisectChal.Deadline,
			},
			TxHash: log.TxHash,
		}
	} else if log.Topics[0] == timedOutChallengeID {
		timeoutChal, err := c.Challenge.ParseTimedOutChallenge(log)
		if err != nil {
			return err
		}
		if timeoutChal.ChallengerWrong {
			c.OutChan <- Notification{
				Header: header,
				VMID:   c.address,
				Event:  AsserterTimeoutEvent{},
				TxHash: log.TxHash,
			}
		} else {
			c.OutChan <- Notification{
				Header: header,
				VMID:   c.address,
				Event:  ChallengerTimeoutEvent{},
				TxHash: log.TxHash,
			}
		}
	} else if log.Topics[0] == continuedChallengeID {
		contChal, err := c.Challenge.ParseContinuedChallenge(log)
		if err != nil {
			return err
		}
		c.OutChan <- Notification{
			Header: header,
			VMID:   c.address,
			Event: ContinueChallengeEvent{
				ChallengedAssertion: uint16(contChal.AssertionIndex.Uint64()),
				Deadline:            contChal.Deadline,
			},
			TxHash: log.TxHash,
		}
	} else if log.Topics[0] == oneStepProofCompletedID {
		_, err = c.Challenge.ParseOneStepProofCompleted(log)
		if err != nil {
			return err
		}
		c.OutChan <- Notification{
			Header: header,
			VMID:   c.address,
			Event:  OneStepProofEvent{},
			TxHash: log.TxHash,
		}
	}
	return nil
}

func (c *Challenge) BisectAssertion(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertions []*protocol.AssertionStub,
) (*types.Receipt, error) {
	machineHashes := make([][32]byte, 0, len(assertions)+1)
	messageAccs := make([][32]byte, 0, len(assertions)+1)
	logAccs := make([][32]byte, 0, len(assertions)+1)
	numGases := make([]uint64, 0, len(assertions)+1)
	totalSteps := uint32(0)
	machineHashes = append(machineHashes, precondition.BeforeHashValue())
	messageAccs = append(messageAccs, assertions[0].FirstMessageHashValue())
	logAccs = append(logAccs, assertions[0].FirstLogHashValue())
	for _, assertion := range assertions {
		machineHashes = append(machineHashes, assertion.AfterHashValue())
		messageAccs = append(messageAccs, assertion.LastMessageHashValue())
		logAccs = append(logAccs, assertion.LastLogHashValue())
		numGases = append(numGases, assertion.NumGas)
		totalSteps += assertion.NumSteps
	}
	var preData [32]byte
	copy(preData[:], solsha3.SoliditySHA3(
		solsha3.Uint64(precondition.TimeBounds.StartTime),
		solsha3.Uint64(precondition.TimeBounds.EndTime),
		solsha3.Bytes32(precondition.BeforeInbox.Value),
	))
	tx, err := c.Challenge.BisectAssertion(
		auth,
		preData,
		machineHashes,
		messageAccs,
		logAccs,
		numGases,
		totalSteps,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, c.Client, auth, tx, "BisectAssertion")
}

func (c *Challenge) ContinueChallenge(
	auth *bind.TransactOpts,
	assertionToChallenge uint16,
	precondition *protocol.Precondition,
	assertions []*protocol.AssertionStub,
) (*types.Receipt, error) {
	var preData [32]byte
	copy(preData[:], solsha3.SoliditySHA3(
		solsha3.Uint64(precondition.TimeBounds.StartTime),
		solsha3.Uint64(precondition.TimeBounds.EndTime),
		solsha3.Bytes32(precondition.BeforeInbox.Value),
	))

	bisectionHashes := make([][32]byte, 0, len(assertions))
	preconditions := protocol.GeneratePreconditions(precondition, assertions)
	for i := range assertions {
		bisectionHash := [32]byte{}
		copy(bisectionHash[:], solsha3.SoliditySHA3(
			solsha3.Bytes32(preData),
			solsha3.Bytes32(preconditions[i].BeforeHashValue()),
			solsha3.Bytes32(assertions[i].Hash()),
		))
		bisectionHashes = append(bisectionHashes, bisectionHash)
	}
	tree := NewMerkleTree(bisectionHashes)
	tx, err := c.Challenge.ContinueChallenge(
		auth,
		big.NewInt(int64(assertionToChallenge)),
		tree.GetProofFlat(int(assertionToChallenge)),
		tree.GetRoot(),
		tree.GetNode(int(assertionToChallenge)),
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, c.Client, auth, tx, "ContinueChallenge")
}

func (c *Challenge) OneStepProof(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
	proof []byte,
) (*types.Receipt, error) {
	tx, err := c.Challenge.OneStepProof(
		auth,
		precondition.BeforeHashValue(),
		precondition.BeforeInboxValue(),
		[2]uint64{precondition.TimeBounds.StartTime, precondition.TimeBounds.EndTime},
		assertion.AfterHashValue(),
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
	return waitForReceipt(auth.Context, c.Client, auth, tx, "OneStepProof")
}

func (c *Challenge) AsserterTimedOutChallenge(
	auth *bind.TransactOpts,
) (*types.Receipt, error) {
	tx, err := c.Challenge.AsserterTimedOut(
		auth,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, c.Client, auth, tx, "AsserterTimedOut")
}

func (c *Challenge) ChallengerTimedOutChallenge(
	auth *bind.TransactOpts,
) (*types.Receipt, error) {
	tx, err := c.Challenge.ChallengerTimedOut(
		auth,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, c.Client, auth, tx, "ChallengerTimedOut")
}

func translateBisectionEvent(event *challenge.ChallengeBisectedAssertion) []*protocol.AssertionStub {
	bisectionCount := len(event.MachineHashes) - 1
	assertions := make([]*protocol.AssertionStub, 0, bisectionCount)
	for i := 0; i < bisectionCount; i++ {
		steps := uint32(0)
		if i == 0 {
			steps = event.TotalSteps/uint32(bisectionCount) + event.TotalSteps%uint32(bisectionCount)
		} else {
			steps = event.TotalSteps / uint32(bisectionCount)
		}
		assertion := &protocol.AssertionStub{
			AfterHash:        value.NewHashBuf(event.MachineHashes[i+1]),
			NumSteps:         steps,
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
