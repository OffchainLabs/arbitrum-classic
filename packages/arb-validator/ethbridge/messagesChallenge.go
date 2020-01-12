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
	"math/big"
	"strings"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/messageschallenge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

var messagesBisectedID ethcommon.Hash
var messagesOneStepProofCompletedID ethcommon.Hash

func init() {
	parsed, err := abi.JSON(strings.NewReader(messageschallenge.MessagesChallengeABI))
	if err != nil {
		panic(err)
	}
	messagesBisectedID = parsed.Events["Bisected"].ID()
	messagesOneStepProofCompletedID = parsed.Events["OneStepProofCompleted"].ID()
}

type messagesChallenge struct {
	*bisectionChallenge
	contract *messageschallenge.MessagesChallenge
}

func newMessagesChallenge(address ethcommon.Address, client *ethclient.Client, auth *bind.TransactOpts) (*messagesChallenge, error) {
	bisectionChallenge, err := newBisectionChallenge(address, client, auth)
	if err != nil {
		return nil, err
	}
	vm := &messagesChallenge{bisectionChallenge: bisectionChallenge}
	err = vm.setupContracts()
	return vm, err
}

func (c *messagesChallenge) setupContracts() error {
	challengeManagerContract, err := messageschallenge.NewMessagesChallenge(c.address, c.client)
	if err != nil {
		return errors2.Wrap(err, "Failed to connect to messagesChallenge")
	}

	c.contract = challengeManagerContract
	return nil
}

func (c *messagesChallenge) StartConnection(ctx context.Context, outChan chan arbbridge.Notification, errChan chan error) error {
	if err := c.bisectionChallenge.StartConnection(ctx, outChan, errChan); err != nil {
		return err
	}
	if err := c.setupContracts(); err != nil {
		return err
	}
	header, err := c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	filter := ethereum.FilterQuery{
		Addresses: []ethcommon.Address{c.address},
		Topics: [][]ethcommon.Hash{{
			messagesBisectedID,
			messagesOneStepProofCompletedID,
		}},
	}

	logs, err := c.client.FilterLogs(ctx, filter)
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
	logSub, err := c.client.SubscribeFilterLogs(ctx, filter, logChan)
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

func (c *messagesChallenge) processEvents(ctx context.Context, log types.Log, outChan chan arbbridge.Notification) error {
	event, err := func() (arbbridge.Event, error) {
		if log.Topics[0] == messagesBisectedID {
			eventVal, err := c.contract.ParseBisected(log)
			if err != nil {
				return nil, err
			}
			return arbbridge.MessagesBisectionEvent{
				ChainHashes:   hashSliceToHashes(eventVal.ChainHashes),
				SegmentHashes: hashSliceToHashes(eventVal.SegmentHashes),
				TotalLength:   eventVal.TotalLength,
				Deadline:      common.TimeTicks{Val: eventVal.DeadlineTicks},
			}, nil
		} else if log.Topics[0] == messagesOneStepProofCompletedID {
			_, err := c.contract.ParseOneStepProofCompleted(log)
			if err != nil {
				return nil, err
			}
			return arbbridge.OneStepProofEvent{}, nil
		}
		return nil, errors2.New("unknown arbitrum event type")
	}()

	if err != nil {
		return err
	}

	header, err := c.client.HeaderByHash(ctx, log.BlockHash)
	if err != nil {
		return err
	}
	outChan <- arbbridge.Notification{
		BlockHeader: common.NewHashFromEth(header.Hash()),
		BlockHeight: header.Number,
		VMID:        common.NewAddressFromEth(c.address),
		Event:       event,
		TxHash:      log.TxHash,
	}
	return nil
}

func (c *messagesChallenge) Bisect(
	ctx context.Context,
	chainHashes []common.Hash,
	segmentHashes []common.Hash,
	chainLength *big.Int,
) error {
	c.auth.Context = ctx
	tx, err := c.contract.Bisect(
		c.auth,
		hashSliceToRaw(chainHashes),
		hashSliceToRaw(segmentHashes),
		chainLength,
	)
	if err != nil {
		return err
	}
	return c.waitForReceipt(ctx, tx, "Bisect")
}

func (c *messagesChallenge) OneStepProof(
	ctx context.Context,
	lowerHashA common.Hash,
	topHashA common.Hash,
	lowerHashB common.Hash,
	topHashB common.Hash,
	value common.Hash,
) error {
	c.auth.Context = ctx
	tx, err := c.contract.OneStepProof(
		c.auth,
		lowerHashA,
		topHashA,
		lowerHashB,
		topHashB,
		value,
	)
	if err != nil {
		return err
	}
	return c.waitForReceipt(ctx, tx, "OneStepProof")
}

func (c *messagesChallenge) ChooseSegment(
	ctx context.Context,
	assertionToChallenge uint16,
	chainHashes []common.Hash,
	segmentHashes []common.Hash,
	chainLength *big.Int,
) error {
	bisectionCount := uint32(len(chainHashes) - 1)
	bisectionHashes := make([]common.Hash, 0, bisectionCount)
	for i := uint32(0); i < bisectionCount; i++ {
		stepCount := structures.CalculateBisectionStepCount(i, bisectionCount, uint32(chainLength.Uint64()))
		bisectionHashes = append(
			bisectionHashes,
			structures.MessageChallengeDataHash(
				chainHashes[i],
				chainHashes[i+1],
				segmentHashes[i],
				segmentHashes[i+1],
				new(big.Int).SetUint64(uint64(stepCount)),
			),
		)
	}
	return c.bisectionChallenge.chooseSegment(
		ctx,
		assertionToChallenge,
		bisectionHashes,
	)
}
