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
	"bytes"
	"context"
	"log"
	"math/big"
	"strings"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"

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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/rollup"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

var rollupStakeCreatedID common.Hash
var rollupChallengeStartedID common.Hash
var rollupChallengeCompletedID common.Hash
var rollupRefundedID common.Hash
var rollupPrunedID common.Hash
var rollupStakeMovedID common.Hash
var rollupAssertedID common.Hash
var rollupConfirmedID common.Hash
var confirmedAssertionID common.Hash
var messageDeliveredID common.Hash

func init() {
	parsedRollup, err := abi.JSON(strings.NewReader(rollup.ArbRollupABI))
	if err != nil {
		panic(err)
	}
	parsedInbox, err := abi.JSON(strings.NewReader(rollup.IGlobalPendingInboxABI))
	if err != nil {
		panic(err)
	}
	rollupStakeCreatedID = parsedRollup.Events["RollupStakeCreated"].ID()
	rollupChallengeStartedID = parsedRollup.Events["RollupChallengeStarted"].ID()
	rollupChallengeCompletedID = parsedRollup.Events["RollupChallengeCompleted"].ID()
	rollupRefundedID = parsedRollup.Events["RollupStakeRefunded"].ID()
	rollupPrunedID = parsedRollup.Events["RollupPruned"].ID()
	rollupStakeMovedID = parsedRollup.Events["RollupStakeMoved"].ID()
	rollupAssertedID = parsedRollup.Events["RollupAsserted"].ID()
	rollupConfirmedID = parsedRollup.Events["RollupConfirmed"].ID()
	confirmedAssertionID = parsedRollup.Events["ConfirmedAssertion"].ID()

	messageDeliveredID = parsedInbox.Events["MessageDelivered"].ID()
}

type EthRollupWatcher struct {
	Client             *ethclient.Client
	ArbRollup          *rollup.ArbRollup
	GlobalPendingInbox *rollup.IGlobalPendingInbox

	address             common.Address
	pendingInboxAddress common.Address
	client              *ethclient.Client
}

func NewRollupWatcher(address common.Address, client *ethclient.Client) (*EthRollupWatcher, error) {
	vm := &EthRollupWatcher{Client: client, address: address}
	err := vm.setupContracts()
	return vm, err
}

func (vm *EthRollupWatcher) setupContracts() error {
	arbitrumRollupContract, err := rollup.NewArbRollup(vm.address, vm.Client)
	if err != nil {
		return errors2.Wrap(err, "Failed to connect to ArbRollup")
	}

	globalPendingInboxAddress, err := arbitrumRollupContract.GlobalInbox(&bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	})
	if err != nil {
		return errors2.Wrap(err, "Failed to get GlobalPendingInbox address")
	}
	vm.pendingInboxAddress = globalPendingInboxAddress
	globalPendingContract, err := rollup.NewIGlobalPendingInbox(globalPendingInboxAddress, vm.Client)
	if err != nil {
		return errors2.Wrap(err, "Failed to connect to GlobalPendingInbox")
	}

	vm.ArbRollup = arbitrumRollupContract
	vm.GlobalPendingInbox = globalPendingContract
	return nil
}

func (vm *EthRollupWatcher) StartConnection(ctx context.Context, outChan chan arbbridge.Notification, errChan chan error) error {
	if err := vm.setupContracts(); err != nil {
		return err
	}

	headers := make(chan *types.Header)
	headersSub, err := vm.Client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return err
	}

	currentHeader, err := vm.Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	filter := ethereum.FilterQuery{
		Addresses: []common.Address{vm.address},
		Topics: [][]common.Hash{
			{
				rollupStakeCreatedID,
				rollupChallengeStartedID,
				rollupChallengeCompletedID,
				rollupRefundedID,
				rollupPrunedID,
				rollupStakeMovedID,
				rollupAssertedID,
				rollupConfirmedID,
				confirmedAssertionID,
			},
		},
	}

	addressIndex := common.Hash{}
	copy(addressIndex[:], common.LeftPadBytes(vm.address.Bytes(), 32))
	messagesFilter := ethereum.FilterQuery{
		Addresses: []common.Address{vm.pendingInboxAddress},
		Topics: [][]common.Hash{
			{messageDeliveredID},
			{addressIndex},
		},
	}
	messagesFilter.ToBlock = currentHeader.Number
	messageLogs, err := vm.Client.FilterLogs(ctx, messagesFilter)
	if err != nil {
		return err
	}
	for _, log := range messageLogs {
		if err := vm.processEvents(ctx, log, outChan); err != nil {
			return err
		}
	}

	filter.ToBlock = currentHeader.Number
	logs, err := vm.Client.FilterLogs(ctx, filter)
	if err != nil {
		return err
	}
	for _, log := range logs {
		if err := vm.processEvents(ctx, log, outChan); err != nil {
			return err
		}
	}

	filter.FromBlock = currentHeader.Number
	filter.ToBlock = nil
	logChan := make(chan types.Log)
	logSub, err := vm.Client.SubscribeFilterLogs(ctx, filter, logChan)
	if err != nil {
		return err
	}

	messagesFilter.FromBlock = currentHeader.Number
	messagesFilter.ToBlock = nil
	messagesLogChan := make(chan types.Log)
	messagesLogSub, err := vm.Client.SubscribeFilterLogs(ctx, messagesFilter, messagesLogChan)
	if err != nil {
		return err
	}

	go func() {
		defer headersSub.Unsubscribe()
		defer messagesLogSub.Unsubscribe()
		defer logSub.Unsubscribe()

		for {
			select {
			case <-ctx.Done():
				break
			case header := <-headers:
				outChan <- arbbridge.Notification{
					Header: header,
					Event:  arbbridge.NewTimeEvent{},
				}
			case log := <-messagesLogChan:
				if err := vm.processEvents(ctx, log, outChan); err != nil {
					errChan <- err
					return
				}
			case log := <-logChan:
				if err := vm.processEvents(ctx, log, outChan); err != nil {
					errChan <- err
					return
				}
			case err := <-headersSub.Err():
				errChan <- err
				return
			case err := <-messagesLogSub.Err():
				errChan <- err
				return
			case err := <-logSub.Err():
				errChan <- err
				return
			}
		}
	}()
	return nil
}

func (vm *EthRollupWatcher) processEvents(ctx context.Context, log types.Log, outChan chan arbbridge.Notification) error {
	event, err := func() (arbbridge.Event, error) {
		if log.Topics[0] == rollupStakeCreatedID {
			eventVal, err := vm.ArbRollup.ParseRollupStakeCreated(log)
			if err != nil {
				return nil, err
			}
			return arbbridge.StakeCreatedEvent{
				Staker:   eventVal.Staker,
				NodeHash: eventVal.NodeHash,
			}, nil
		} else if log.Topics[0] == rollupChallengeStartedID {
			eventVal, err := vm.ArbRollup.ParseRollupChallengeStarted(log)
			if err != nil {
				return nil, err
			}
			return arbbridge.ChallengeStartedEvent{
				Asserter:          eventVal.Asserter,
				Challenger:        eventVal.Challenger,
				ChallengeType:     structures.ChildType(eventVal.ChallengeType.Uint64()),
				ChallengeContract: eventVal.ChallengeContract,
			}, nil
		} else if log.Topics[0] == rollupChallengeCompletedID {
			eventVal, err := vm.ArbRollup.ParseRollupChallengeCompleted(log)
			if err != nil {
				return nil, err
			}
			return arbbridge.ChallengeCompletedEvent{
				Winner:            eventVal.Winner,
				Loser:             eventVal.Loser,
				ChallengeContract: eventVal.ChallengeContract,
			}, nil
		} else if log.Topics[0] == rollupRefundedID {
			eventVal, err := vm.ArbRollup.ParseRollupStakeRefunded(log)
			if err != nil {
				return nil, err
			}
			return arbbridge.StakeRefundedEvent{
				Staker: eventVal.Staker,
			}, nil
		} else if log.Topics[0] == rollupPrunedID {
			eventVal, err := vm.ArbRollup.ParseRollupPruned(log)
			if err != nil {
				return nil, err
			}
			return arbbridge.PrunedEvent{
				Leaf: eventVal.Leaf,
			}, nil
		} else if log.Topics[0] == rollupStakeMovedID {
			eventVal, err := vm.ArbRollup.ParseRollupStakeMoved(log)
			if err != nil {
				return nil, err
			}
			return arbbridge.StakeMovedEvent{
				Staker:   eventVal.Staker,
				Location: eventVal.ToNodeHash,
			}, nil
		} else if log.Topics[0] == rollupAssertedID {
			eventVal, err := vm.ArbRollup.ParseRollupAsserted(log)
			if err != nil {
				return nil, err
			}
			return arbbridge.AssertedEvent{
				PrevLeafHash: eventVal.PrevLeaf,
				Params: &structures.AssertionParams{
					NumSteps: eventVal.NumSteps,
					TimeBounds: protocol.NewTimeBoundsBlocks(
						protocol.NewTimeBlocks(eventVal.TimeBoundsBlocks[0]),
						protocol.NewTimeBlocks(eventVal.TimeBoundsBlocks[1]),
					),
					ImportedMessageCount: eventVal.ImportedMessageCount,
				},
				Claim: &structures.AssertionClaim{
					AfterPendingTop:       eventVal.AfterPendingTop,
					ImportedMessagesSlice: eventVal.ImportedMessagesSlice,
					AssertionStub: &valprotocol.ExecutionAssertionStub{
						AfterHash:        eventVal.AfterVMHash,
						DidInboxInsn:     eventVal.DidInboxInsn,
						NumGas:           eventVal.NumArbGas,
						FirstMessageHash: [32]byte{},
						LastMessageHash:  eventVal.MessagesAccHash,
						FirstLogHash:     [32]byte{},
						LastLogHash:      eventVal.LogsAccHash,
					},
				},
				MaxPendingTop: eventVal.PendingValue,
			}, nil
		} else if log.Topics[0] == rollupConfirmedID {
			eventVal, err := vm.ArbRollup.ParseRollupConfirmed(log)
			if err != nil {
				return nil, err
			}
			return arbbridge.ConfirmedEvent{
				NodeHash: eventVal.NodeHash,
			}, nil
		} else if log.Topics[0] == confirmedAssertionID {
			eventVal, err := vm.ArbRollup.ParseConfirmedAssertion(log)
			if err != nil {
				return nil, err
			}
			return arbbridge.ConfirmedAssertionEvent{
				LogsAccHash: eventVal.LogsAccHash,
			}, nil
		} else if log.Topics[0] == messageDeliveredID {
			val, err := vm.GlobalPendingInbox.ParseMessageDelivered(log)
			if err != nil {
				return nil, err
			}

			rd := bytes.NewReader(val.Data)
			msgData, err := value.UnmarshalValue(rd)
			if err != nil {
				return nil, err
			}

			messageHash := solsha3.SoliditySHA3(
				solsha3.Address(val.VmId),
				solsha3.Bytes32(msgData.Hash()),
				solsha3.Uint256(val.Value),
				val.TokenType[:],
			)
			msgHashInt := new(big.Int).SetBytes(messageHash[:])

			msgVal, _ := value.NewTupleFromSlice([]value.Value{
				msgData,
				value.NewIntValue(new(big.Int).SetUint64(log.BlockNumber)),
				value.NewIntValue(msgHashInt),
			})

			msg := protocol.NewSimpleMessage(msgVal, val.TokenType, val.Value, val.Sender)
			return arbbridge.MessageDeliveredEvent{
				Msg: msg,
			}, nil
		}
		return nil, errors2.New("unknown arbitrum event type")
	}()

	if err != nil {
		return err
	}
	if event != nil {
		header, err := vm.Client.HeaderByHash(ctx, log.BlockHash)
		if err != nil {
			return err
		}
		outChan <- arbbridge.Notification{
			Header: header,
			VMID:   vm.address,
			Event:  event,
			TxHash: log.TxHash,
		}
	}
	return nil
}

func (vm *EthRollupWatcher) GetParams(ctx context.Context) (structures.ChainParams, error) {
	log.Println("Calling GetParams")
	rawParams, err := vm.ArbRollup.VmParams(nil)
	if err != nil {
		return structures.ChainParams{}, err
	}
	stakeRequired, err := vm.ArbRollup.GetStakeRequired(nil)
	if err != nil {
		return structures.ChainParams{}, err
	}
	return structures.ChainParams{
		StakeRequirement:        stakeRequired,
		GracePeriod:             structures.TimeTicks{rawParams.GracePeriodTicks},
		MaxExecutionSteps:       rawParams.MaxExecutionSteps,
		ArbGasSpeedLimitPerTick: rawParams.ArbGasSpeedLimitPerTick.Uint64(),
	}, nil
}

func (vm *EthRollupWatcher) InboxAddress(ctx context.Context) (common.Address, error) {
	return vm.ArbRollup.GlobalInbox(nil)
}
