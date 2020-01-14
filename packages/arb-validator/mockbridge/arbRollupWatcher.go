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

type EthRollupWatcher struct {
	client arbbridge.ArbClient

	address common.Address
}

func NewRollupWatcher(address common.Address, client arbbridge.ArbClient) (*EthRollupWatcher, error) {
	//vm := &EthRollupWatcher{Client: client.(*ArbClient).client, address: address}
	//err := vm.setupContracts()
	//return vm, err
	return &EthRollupWatcher{client: client}, nil
}

//func (vm *EthRollupWatcher) setupContracts() error {
//	arbitrumRollupContract, err := rollup.NewArbRollup(vm.address, vm.Client)
//	if err != nil {
//		return errors2.Wrap(err, "Failed to connect to ArbRollup")
//	}
//
//	globalPendingInboxAddress, err := arbitrumRollupContract.GlobalInbox(&bind.CallOpts{
//		Pending: false,
//		Context: context.Background(),
//	})
//	if err != nil {
//		return errors2.Wrap(err, "Failed to get GlobalPendingInbox address")
//	}
//	globalPendingContract, err := rollup.NewIGlobalPendingInbox(globalPendingInboxAddress, vm.Client)
//	if err != nil {
//		return errors2.Wrap(err, "Failed to connect to GlobalPendingInbox")
//	}
//
//	vm.ArbRollup = arbitrumRollupContract
//	vm.GlobalPendingInbox = globalPendingContract
//	return nil
//}

func (vm *EthRollupWatcher) StartConnection(ctx context.Context, startHeight *common.TimeBlocks, startLogIndex uint) (<-chan arbbridge.Event, <-chan error, error) {
	//if err := vm.setupContracts(); err != nil {
	//	return err
	//}
	//
	//start := uint64(0)
	//watch := &bind.WatchOpts{
	//	Context: ctx,
	//	Start:   &start,
	//}
	//
	//headers := make(chan *types.Header)
	//headersSub, err := vm.Client.SubscribeNewHead(ctx, headers)
	//if err != nil {
	//	return err
	//}
	//
	//filter := ethereum.FilterQuery{
	//	Addresses: []common.Address{vm.address},
	//	Topics: [][]common.Hash{
	//		{
	//			rollupStakeCreatedID,
	//			rollupChallengeStartedID,
	//			rollupChallengeCompletedID,
	//			rollupRefundedID,
	//			rollupPrunedID,
	//			rollupStakeMovedID,
	//			rollupAssertedID,
	//			rollupConfirmedID,
	//			confirmedAssertionID,
	//			debugEventID,
	//		},
	//	},
	//}
	//
	//logChan := make(chan types.Log)
	//logSub, err := vm.Client.SubscribeFilterLogs(ctx, filter, logChan)
	//if err != nil {
	//	return err
	//}
	//
	//messageDeliveredChan := make(chan *rollup.IGlobalPendingInboxMessageDelivered)
	//messageDeliveredSub, err := vm.GlobalPendingInbox.WatchMessageDelivered(watch, messageDeliveredChan, []common.Address{vm.address})
	//if err != nil {
	//	return err
	//}
	//
	//go func() {
	//	defer headersSub.Unsubscribe()
	//	defer messageDeliveredSub.Unsubscribe()
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
	//		case val := <-messageDeliveredChan:
	//			header, err := vm.Client.HeaderByHash(context.Background(), val.Raw.BlockHash)
	//			if err != nil {
	//				errChan <- err
	//				return
	//			}
	//			rd := bytes.NewReader(val.Data)
	//			msgData, err := value.UnmarshalValue(rd)
	//			if err != nil {
	//				errChan <- err
	//				return
	//			}
	//
	//			messageHash := solsha3.SoliditySHA3(
	//				solsha3.Address(val.VmId),
	//				solsha3.Bytes32(msgData.Hash()),
	//				solsha3.Uint256(val.Value),
	//				val.TokenType[:],
	//			)
	//			msgHashInt := new(big.Int).SetBytes(messageHash[:])
	//
	//			msgVal, _ := value.NewTupleFromSlice([]value.Value{
	//				msgData,
	//				value.NewIntValue(new(big.Int).SetUint64(header.Time)),
	//				value.NewIntValue(header.Number),
	//				value.NewIntValue(msgHashInt),
	//			})
	//
	//			msg := protocol.NewSimpleMessage(msgVal, val.TokenType, val.Value, val.Sender)
	//			outChan <- arbbridge.Notification{
	//				Header: header,
	//				VMID:   val.VmId,
	//				Event: arbbridge.MessageDeliveredEvent{
	//					Msg: msg,
	//				},
	//				TxHash: val.Raw.TxHash,
	//			}
	//		case log := <-logChan:
	//			if err := vm.processEvents(ctx, log, outChan); err != nil {
	//				errChan <- err
	//				return
	//			}
	//		case err := <-headersSub.Err():
	//			errChan <- err
	//			return
	//		case err := <-messageDeliveredSub.Err():
	//			errChan <- err
	//			return
	//		case err := <-logSub.Err():
	//			errChan <- err
	//			return
	//		}
	//	}
	//}()
	return nil, nil, nil
}

//func (vm *EthRollupWatcher) processEvents(ctx context.Context, log types.Log, outChan chan arbbridge.Notification) error {
//	event, err := func() (arbbridge.Event, error) {
//		if log.Topics[0] == rollupStakeCreatedID {
//			eventVal, err := vm.ArbRollup.ParseRollupStakeCreated(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.StakeCreatedEvent{
//				Staker:   eventVal.Staker,
//				NodeHash: eventVal.NodeHash,
//			}, nil
//		} else if log.Topics[0] == rollupChallengeStartedID {
//			eventVal, err := vm.ArbRollup.ParseRollupChallengeStarted(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.ChallengeStartedEvent{
//				Asserter:          eventVal.Asserter,
//				Challenger:        eventVal.Challenger,
//				ChallengeType:     structures.ChildType(eventVal.ChallengeType.Uint64()),
//				ChallengeContract: eventVal.ChallengeContract,
//			}, nil
//		} else if log.Topics[0] == rollupChallengeCompletedID {
//			eventVal, err := vm.ArbRollup.ParseRollupChallengeCompleted(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.ChallengeCompletedEvent{
//				Winner:            eventVal.Winner,
//				Loser:             eventVal.Loser,
//				ChallengeContract: eventVal.ChallengeContract,
//			}, nil
//		} else if log.Topics[0] == rollupRefundedID {
//			eventVal, err := vm.ArbRollup.ParseRollupStakeRefunded(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.StakeRefundedEvent{
//				Staker: eventVal.Staker,
//			}, nil
//		} else if log.Topics[0] == rollupPrunedID {
//			eventVal, err := vm.ArbRollup.ParseRollupPruned(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.PrunedEvent{
//				Leaf: eventVal.Leaf,
//			}, nil
//		} else if log.Topics[0] == rollupStakeMovedID {
//			eventVal, err := vm.ArbRollup.ParseRollupStakeMoved(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.StakeMovedEvent{
//				Staker:   eventVal.Staker,
//				Location: eventVal.ToNodeHash,
//			}, nil
//		} else if log.Topics[0] == rollupAssertedID {
//			eventVal, err := vm.ArbRollup.ParseRollupAsserted(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.AssertedEvent{
//				PrevLeafHash: eventVal.PrevLeaf,
//				Params: &structures.AssertionParams{
//					NumSteps: eventVal.NumSteps,
//					TimeBounds: protocol.NewTimeBoundsBlocks(
//						protocol.NewTimeBlocks(eventVal.TimeBoundsBlocks[0]),
//						protocol.NewTimeBlocks(eventVal.TimeBoundsBlocks[1]),
//					),
//					ImportedMessageCount: eventVal.ImportedMessageCount,
//				},
//				Claim: &structures.AssertionClaim{
//					AfterPendingTop:       eventVal.AfterPendingTop,
//					ImportedMessagesSlice: eventVal.ImportedMessagesSlice,
//					AssertionStub: protocol.NewExecutionAssertionStub(
//						eventVal.AfterVMHash,
//						eventVal.DidInboxInsn,
//						eventVal.NumArbGas,
//						eventVal.MessagesAccHash,
//						eventVal.LogsAccHash,
//					),
//				},
//				MaxPendingTop: eventVal.PendingValue,
//			}, nil
//		} else if log.Topics[0] == rollupConfirmedID {
//			eventVal, err := vm.ArbRollup.ParseRollupConfirmed(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.ConfirmedEvent{
//				NodeHash: eventVal.NodeHash,
//			}, nil
//		} else if log.Topics[0] == confirmedAssertionID {
//			eventVal, err := vm.ArbRollup.ParseConfirmedAssertion(log)
//			if err != nil {
//				return nil, err
//			}
//			return arbbridge.ConfirmedAssertionEvent{
//				LogsAccHash: eventVal.LogsAccHash,
//			}, nil
//		} else if log.Topics[0] == debugEventID {
//			//eventVal, err := vm.ArbRollup.ParseDebugData(log)
//			//if err != nil {
//			//	return nil, err
//			//}
//			//fmt.Println("Debug event")
//			//fmt.Println("BeforeVMHash", hexutil.Encode(eventVal.BeforeVMHash[:]))
//			//fmt.Println("TimeBounds", eventVal.TimeBoundsBlocks)
//			//fmt.Println("Inbox hash", hexutil.Encode(eventVal.Inbox[:]))
//			//fmt.Println("PreconditionHash", hexutil.Encode(eventVal.Precondition[:]))
//
//			//fmt.Println("PrevLeaf", hexutil.Encode(eventVal.PrevLeaf[:]))
//			//fmt.Println("DeadlineTicks", eventVal.DeadlineTicks)
//			//fmt.Println("BeforePendingTop", hexutil.Encode(eventVal.BeforePendingTop[:]))
//			//fmt.Println("AfterPendingTop", hexutil.Encode(eventVal.AfterPendingTop[:]))
//			//fmt.Println("ImportedMessagesSlice", hexutil.Encode(eventVal.ImportedMessagesSlice[:]))
//			//fmt.Println("ImportedMessageCount", eventVal.ImportedMessageCount)
//			//fmt.Println("ChallengePeriod", eventVal.ChallengePeriod)
//			//fmt.Println("ChildType", eventVal.ChildType)
//			//fmt.Println("VmProtoHashBefore", hexutil.Encode(eventVal.VmProtoHashBefore[:]))
//			//fmt.Println("ChallengeHash", hexutil.Encode(eventVal.ChallengeHash[:]))
//			//fmt.Println("NodeDataHash", hexutil.Encode(eventVal.NodeDataHash[:]))
//			return nil, nil
//		}
//		return nil, errors2.New("unknown arbitrum event type")
//	}()
//
//	if err != nil {
//		return err
//	}
//	if event != nil {
//		header, err := vm.Client.HeaderByHash(ctx, log.BlockHash)
//		if err != nil {
//			return err
//		}
//		outChan <- arbbridge.Notification{
//			Header: header,
//			VMID:   vm.address,
//			Event:  event,
//			TxHash: log.TxHash,
//		}
//	}
//	return nil
//}

func (vm *EthRollupWatcher) GetParams(ctx context.Context) (structures.ChainParams, error) {
	return structures.ChainParams{
		StakeRequirement:        nil,
		GracePeriod:             common.TimeTicks{},
		MaxExecutionSteps:       0,
		ArbGasSpeedLimitPerTick: 0,
	}, nil
}

func (vm *EthRollupWatcher) InboxAddress(ctx context.Context) (common.Address, error) {
	return common.Address{}, nil
}
