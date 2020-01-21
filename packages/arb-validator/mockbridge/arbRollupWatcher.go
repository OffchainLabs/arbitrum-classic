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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type ethRollupWatcher struct {
	client *MockArbClient

	address common.Address
}

func newRollupWatcher(address common.Address, client *MockArbClient) (*ethRollupWatcher, error) {
	vm := &ethRollupWatcher{client: client, address: address}
	err := vm.setupContracts()
	return vm, err
	//arbitrumRollupContract, err := rollup.NewArbRollup(rollupAddress, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to arbRollup")
	//}
	//
	//globalPendingInboxAddress, err := arbitrumRollupContract.GlobalInbox(&bind.CallOpts{
	//	Pending: false,
	//	Context: context.Background(),
	//})
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to get GlobalPendingInbox address")
	//}
	//globalPendingContract, err := rollup.NewIGlobalPendingInbox(globalPendingInboxAddress, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to GlobalPendingInbox")
	//}
	//
	//return &ethRollupWatcher{
	//	ArbRollup:          arbitrumRollupContract,
	//	GlobalPendingInbox: globalPendingContract,
	//	rollupAddress:      rollupAddress,
	//	inboxAddress:       globalPendingInboxAddress,
	//	client:             client,
	//}, nil
}

func (vm *ethRollupWatcher) setupContracts() error {
	return nil
}

//	//	arbitrumRollupContract, err := rollup.NewArbRollup(vm.address, vm.Client)
//	//	if err != nil {
//	//		return errors2.Wrap(err, "Failed to connect to arbRollup")
//	//	}
//	//
//	//	globalPendingInboxAddress, err := arbitrumRollupContract.GlobalInbox(&bind.CallOpts{
//	//		Pending: false,
//	//		Context: context.Background(),
//	//	})
//	//	if err != nil {
//	//		return errors2.Wrap(err, "Failed to get GlobalPendingInbox address")
//	//	}
//	//	globalPendingContract, err := rollup.NewIGlobalPendingInbox(globalPendingInboxAddress, vm.Client)
//	//	if err != nil {
//	//		return errors2.Wrap(err, "Failed to connect to GlobalPendingInbox")
//	//	}
//	//
//	//	vm.arbRollup = arbitrumRollupContract
//	//	vm.GlobalPendingInbox = globalPendingContract
//	return nil
//}
//
////func (vm *ethRollupWatcher) StartConnection(ctx context.Context, startHeight *common.TimeBlocks, startLogIndex uint) (<-chan MaybeEvent, error) {
////	if err := vm.setupContracts(); err != nil {
////		return nil, err
////	}
////
////	logChan := make(chan types.Log, 1024)
////	logErrChan := make(chan error, 10)
////	headers := make(chan *types.Header)
////	headersSub, err := vm.client.SubscribeNewHead(ctx, headers)
////	if err != nil {
////		return nil, err
////	}
////
////	//if err := getLogs(ctx, vm.client, vm.rollupFilter(), big.NewInt(0), logChan, logErrChan); err != nil {
////	//	return err
////	//}
////	//
////	//if err := getLogs(ctx, vm.client, vm.messageFilter(), big.NewInt(0), logChan, logErrChan); err != nil {
////	//	return err
////	//}
////
////	eventChan := make(chan arbbridge.MaybeEvent, 1024)
////	go func() {
////
////		for {
////			select {
////			case <-ctx.Done():
////				break
////			case ethLog := <-logChan:
////				if err := vm.processEvents(ctx, ethLog, outChan); err != nil {
////					errChan <- err
////					return
////				}
////			case err := <-logErrChan:
////				errChan <- err
////				return
////			}
////		}
////	}()
////	return eventChan, nil
////}
//
//func (vm *ethRollupWatcher) StartConnection(ctx context.Context, startHeight *common.TimeBlocks, startLogIndex uint) (<-chan arbbridge.MaybeEvent, error) {
//	if err := vm.setupContracts(); err != nil {
//		return nil, err
//	}
//
//	headers := make(chan arbbridge.MaybeEvent)
//	vm.client.MockEthClient.registerOutChan(headers)
//	//headersSub, err := vm.client.MockEthClient.registerOutChan(headers)
//	//if err != nil {
//	//	return nil, err
//	//}
//
//	//logCtx, cancelFunc := context.WithCancel(ctx)
//
//	//rollupMaybeLogChan, err := getLogs(logCtx, vm.client, vm.rollupFilter(), startHeight, startLogIndex)
//	//if err != nil {
//	//	return nil, err
//	//}
//	//inboxMaybeLogChan, err := getLogs(logCtx, vm.client, vm.messageFilter(), startHeight, startLogIndex)
//	//if err != nil {
//	//	return nil, err
//	//}
//
//	eventChan := make(chan arbbridge.MaybeEvent, 1024)
//	go func() {
//		defer close(eventChan)
//		//defer cancelFunc()
//
//		for {
//			select {
//			case <-ctx.Done():
//				return
//			case header := <-headers:
//				eventChan <- header
//				//case maybeLog, ok := <-rollupMaybeLogChan:
//				//	if !ok {
//				//		eventChan <- arbbridge.MaybeEvent{Err: errors.New("rollupMaybeLogChan terminated early")}
//				//		return
//				//	}
//				//	if maybeLog.err != nil {
//				//		eventChan <- arbbridge.MaybeEvent{Err: err}
//				//		return
//				//	}
//				//	event, err := vm.processEvents(ctx, maybeLog.log)
//				//	if err != nil {
//				//		eventChan <- arbbridge.MaybeEvent{Err: err}
//				//		return
//				//	}
//				//	eventChan <- arbbridge.MaybeEvent{Event: event}
//				//case maybeLog, ok := <-inboxMaybeLogChan:
//				//	if !ok {
//				//		eventChan <- arbbridge.MaybeEvent{Err: errors.New("inboxMaybeLogChan terminated early")}
//				//		return
//				//	}
//				//	if maybeLog.err != nil {
//				//		eventChan <- arbbridge.MaybeEvent{Err: err}
//				//		return
//				//	}
//				//	event, err := vm.processEvents(ctx, maybeLog.log)
//				//	if err != nil {
//				//		eventChan <- arbbridge.MaybeEvent{Err: err}
//				//		return
//				//	}
//				//	eventChan <- arbbridge.MaybeEvent{Event: event}
//				//case err := <-headersSub.Err():
//				//	eventChan <- arbbridge.MaybeEvent{Err: err}
//				//	return
//			}
//		}
//	}()
//	return eventChan, nil
//}
//
//func (vm *ethRollupWatcher) processEvents(ctx context.Context, log types.Log, outChan chan arbbridge.MaybeEvent) error {
//	//	event, err := func() (arbbridge.Event, error) {
//	//		if log.Topics[0] == rollupStakeCreatedID {
//	//			eventVal, err := vm.arbRollup.ParseRollupStakeCreated(log)
//	//			if err != nil {
//	//				return nil, err
//	//			}
//	//			return arbbridge.StakeCreatedEvent{
//	//				staker:   eventVal.staker,
//	//				NodeHash: eventVal.NodeHash,
//	//			}, nil
//	//		} else if log.Topics[0] == rollupChallengeStartedID {
//	//			eventVal, err := vm.arbRollup.ParseRollupChallengeStarted(log)
//	//			if err != nil {
//	//				return nil, err
//	//			}
//	//			return arbbridge.ChallengeStartedEvent{
//	//				Asserter:          eventVal.Asserter,
//	//				Challenger:        eventVal.Challenger,
//	//				ChallengeType:     structures.ChildType(eventVal.ChallengeType.Uint64()),
//	//				ChallengeContract: eventVal.ChallengeContract,
//	//			}, nil
//	//		} else if log.Topics[0] == rollupChallengeCompletedID {
//	//			eventVal, err := vm.arbRollup.ParseRollupChallengeCompleted(log)
//	//			if err != nil {
//	//				return nil, err
//	//			}
//	//			return arbbridge.ChallengeCompletedEvent{
//	//				Winner:            eventVal.Winner,
//	//				Loser:             eventVal.Loser,
//	//				ChallengeContract: eventVal.ChallengeContract,
//	//			}, nil
//	//		} else if log.Topics[0] == rollupRefundedID {
//	//			eventVal, err := vm.arbRollup.ParseRollupStakeRefunded(log)
//	//			if err != nil {
//	//				return nil, err
//	//			}
//	//			return arbbridge.StakeRefundedEvent{
//	//				staker: eventVal.staker,
//	//			}, nil
//	//		} else if log.Topics[0] == rollupPrunedID {
//	//			eventVal, err := vm.arbRollup.ParseRollupPruned(log)
//	//			if err != nil {
//	//				return nil, err
//	//			}
//	//			return arbbridge.PrunedEvent{
//	//				Leaf: eventVal.Leaf,
//	//			}, nil
//	//		} else if log.Topics[0] == rollupStakeMovedID {
//	//			eventVal, err := vm.arbRollup.ParseRollupStakeMoved(log)
//	//			if err != nil {
//	//				return nil, err
//	//			}
//	//			return arbbridge.StakeMovedEvent{
//	//				staker:   eventVal.staker,
//	//				Location: eventVal.ToNodeHash,
//	//			}, nil
//	//		} else if log.Topics[0] == rollupAssertedID {
//	//			eventVal, err := vm.arbRollup.ParseRollupAsserted(log)
//	//			if err != nil {
//	//				return nil, err
//	//			}
//	//			return arbbridge.AssertedEvent{
//	//				PrevLeafHash: eventVal.PrevLeaf,
//	//				Params: &structures.AssertionParams{
//	//					NumSteps: eventVal.NumSteps,
//	//					TimeBounds: protocol.NewTimeBoundsBlocks(
//	//						protocol.NewTimeBlocks(eventVal.TimeBoundsBlocks[0]),
//	//						protocol.NewTimeBlocks(eventVal.TimeBoundsBlocks[1]),
//	//					),
//	//					ImportedMessageCount: eventVal.ImportedMessageCount,
//	//				},
//	//				Claim: &structures.AssertionClaim{
//	//					AfterPendingTop:       eventVal.AfterPendingTop,
//	//					ImportedMessagesSlice: eventVal.ImportedMessagesSlice,
//	//					AssertionStub: protocol.NewExecutionAssertionStub(
//	//						eventVal.AfterVMHash,
//	//						eventVal.DidInboxInsn,
//	//						eventVal.NumArbGas,
//	//						eventVal.MessagesAccHash,
//	//						eventVal.LogsAccHash,
//	//					),
//	//				},
//	//				MaxPendingTop: eventVal.PendingValue,
//	//			}, nil
//	//		} else if log.Topics[0] == rollupConfirmedID {
//	//			eventVal, err := vm.arbRollup.ParseRollupConfirmed(log)
//	//			if err != nil {
//	//				return nil, err
//	//			}
//	//			return arbbridge.ConfirmedEvent{
//	//				NodeHash: eventVal.NodeHash,
//	//			}, nil
//	//		} else if log.Topics[0] == confirmedAssertionID {
//	//			eventVal, err := vm.arbRollup.ParseConfirmedAssertion(log)
//	//			if err != nil {
//	//				return nil, err
//	//			}
//	//			return arbbridge.ConfirmedAssertionEvent{
//	//				LogsAccHash: eventVal.LogsAccHash,
//	//			}, nil
//	//		} else if log.Topics[0] == debugEventID {
//	//			//eventVal, err := vm.arbRollup.ParseDebugData(log)
//	//			//if err != nil {
//	//			//	return nil, err
//	//			//}
//	//			//fmt.Println("Debug event")
//	//			//fmt.Println("BeforeVMHash", hexutil.Encode(eventVal.BeforeVMHash[:]))
//	//			//fmt.Println("TimeBounds", eventVal.TimeBoundsBlocks)
//	//			//fmt.Println("Inbox hash", hexutil.Encode(eventVal.Inbox[:]))
//	//			//fmt.Println("PreconditionHash", hexutil.Encode(eventVal.Precondition[:]))
//	//
//	//			//fmt.Println("PrevLeaf", hexutil.Encode(eventVal.PrevLeaf[:]))
//	//			//fmt.Println("DeadlineTicks", eventVal.DeadlineTicks)
//	//			//fmt.Println("BeforePendingTop", hexutil.Encode(eventVal.BeforePendingTop[:]))
//	//			//fmt.Println("AfterPendingTop", hexutil.Encode(eventVal.AfterPendingTop[:]))
//	//			//fmt.Println("ImportedMessagesSlice", hexutil.Encode(eventVal.ImportedMessagesSlice[:]))
//	//			//fmt.Println("ImportedMessageCount", eventVal.ImportedMessageCount)
//	//			//fmt.Println("ChallengePeriod", eventVal.ChallengePeriod)
//	//			//fmt.Println("ChildType", eventVal.ChildType)
//	//			//fmt.Println("VmProtoHashBefore", hexutil.Encode(eventVal.VmProtoHashBefore[:]))
//	//			//fmt.Println("ChallengeHash", hexutil.Encode(eventVal.ChallengeHash[:]))
//	//			//fmt.Println("NodeDataHash", hexutil.Encode(eventVal.NodeDataHash[:]))
//	//			return nil, nil
//	//		}
//	//		return nil, errors2.New("unknown arbitrum event type")
//	//	}()
//	//
//	//	if err != nil {
//	//		return err
//	//	}
//	//	if event != nil {
//	//		header, err := vm.Client.HeaderByHash(ctx, log.BlockHash)
//	//		if err != nil {
//	//			return err
//	//		}
//	//		outChan <- arbbridge.Notification{
//	//			Header: header,
//	//			VMID:   vm.address,
//	//			Event:  event,
//	//			TxHash: log.TxHash,
//	//		}
//	//	}
//	return nil
//}

func (vm *ethRollupWatcher) GetEvents(ctx context.Context, blockId *structures.BlockId) ([]arbbridge.Event, error) {
	return vm.client.MockEthClient.rollups[vm.address].events, nil
}

func (vm *ethRollupWatcher) GetParams(ctx context.Context) (structures.ChainParams, error) {
	return structures.ChainParams{
		StakeRequirement:        nil,
		GracePeriod:             common.TimeTicks{},
		MaxExecutionSteps:       0,
		ArbGasSpeedLimitPerTick: 0,
	}, nil
}

func (vm *ethRollupWatcher) InboxAddress(ctx context.Context) (common.Address, error) {
	return common.Address{}, nil
}

func (vm *ethRollupWatcher) GetCreationHeight(ctx context.Context) (*structures.BlockId, error) {
	return vm.client.MockEthClient.rollups[vm.address].creation, nil
}
