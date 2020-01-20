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
	"errors"
	"log"
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
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/rollup"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

var rollupStakeCreatedID ethcommon.Hash
var rollupChallengeStartedID ethcommon.Hash
var rollupChallengeCompletedID ethcommon.Hash
var rollupRefundedID ethcommon.Hash
var rollupPrunedID ethcommon.Hash
var rollupStakeMovedID ethcommon.Hash
var rollupAssertedID ethcommon.Hash
var rollupConfirmedID ethcommon.Hash
var confirmedAssertionID ethcommon.Hash
var transactionMessageDeliveredID ethcommon.Hash
var ethDepositMessageDeliveredID ethcommon.Hash
var depositERC20MessageDeliveredID ethcommon.Hash
var depositERC721MessageDeliveredID ethcommon.Hash

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

	transactionMessageDeliveredID = parsedInbox.Events["TransactionMessageDelivered"].ID()
	ethDepositMessageDeliveredID = parsedInbox.Events["EthDepositMessageDelivered"].ID()
	depositERC20MessageDeliveredID = parsedInbox.Events["ERC20DepositMessageDelivered"].ID()
	depositERC721MessageDeliveredID = parsedInbox.Events["ERC721DepositMessageDelivered"].ID()
}

type ethRollupWatcher struct {
	ArbRollup          *rollup.ArbRollup
	GlobalPendingInbox *rollup.IGlobalPendingInbox

	address             ethcommon.Address
	pendingInboxAddress ethcommon.Address
	client              *ethclient.Client
}

func newRollupWatcher(address ethcommon.Address, client *ethclient.Client) (*ethRollupWatcher, error) {
	vm := &ethRollupWatcher{client: client, address: address}
	err := vm.setupContracts()
	return vm, err
}

func (vm *ethRollupWatcher) messageFilter() ethereum.FilterQuery {
	addressIndex := ethcommon.Hash{}
	copy(addressIndex[:], ethcommon.LeftPadBytes(vm.address.Bytes(), 32))
	return ethereum.FilterQuery{
		Addresses: []ethcommon.Address{vm.pendingInboxAddress},
		Topics: [][]ethcommon.Hash{
			{transactionMessageDeliveredID},
			{ethDepositMessageDeliveredID},
			{depositERC20MessageDeliveredID},
			{depositERC721MessageDeliveredID},
			{addressIndex},
		},
	}
}

func (vm *ethRollupWatcher) rollupFilter() ethereum.FilterQuery {
	return ethereum.FilterQuery{
		Addresses: []ethcommon.Address{vm.address},
		Topics: [][]ethcommon.Hash{
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
}

func (vm *ethRollupWatcher) setupContracts() error {
	arbitrumRollupContract, err := rollup.NewArbRollup(vm.address, vm.client)
	if err != nil {
		return errors2.Wrap(err, "Failed to connect to arbRollup")
	}

	globalPendingInboxAddress, err := arbitrumRollupContract.GlobalInbox(&bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	})
	if err != nil {
		return errors2.Wrap(err, "Failed to get GlobalPendingInbox address")
	}
	vm.pendingInboxAddress = globalPendingInboxAddress
	globalPendingContract, err := rollup.NewIGlobalPendingInbox(globalPendingInboxAddress, vm.client)
	if err != nil {
		return errors2.Wrap(err, "Failed to connect to GlobalPendingInbox")
	}

	vm.ArbRollup = arbitrumRollupContract
	vm.GlobalPendingInbox = globalPendingContract
	return nil
}

func (vm *ethRollupWatcher) StartConnection(ctx context.Context, startHeight *common.TimeBlocks, startLogIndex uint) (<-chan arbbridge.MaybeEvent, error) {
	if err := vm.setupContracts(); err != nil {
		return nil, err
	}

	headers := make(chan *types.Header)
	headersSub, err := vm.client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return nil, err
	}

	logCtx, cancelFunc := context.WithCancel(ctx)

	rollupMaybeLogChan, err := getLogs(logCtx, vm.client, vm.rollupFilter(), startHeight, startLogIndex)
	if err != nil {
		return nil, err
	}
	inboxMaybeLogChan, err := getLogs(logCtx, vm.client, vm.messageFilter(), startHeight, startLogIndex)
	if err != nil {
		return nil, err
	}

	eventChan := make(chan arbbridge.MaybeEvent, 1024)
	go func() {
		defer close(eventChan)
		defer cancelFunc()
		defer headersSub.Unsubscribe()

		for {
			select {
			case <-ctx.Done():
				return
			case header := <-headers:
				eventChan <- arbbridge.MaybeEvent{Event: arbbridge.NewTimeEvent{arbbridge.ChainInfo{
					BlockId: getBlockID(header),
				}}}

			case maybeLog, ok := <-rollupMaybeLogChan:
				log.Println("event check 1111")

				if !ok {
					eventChan <- arbbridge.MaybeEvent{Err: errors.New("rollupMaybeLogChan terminated early")}
					return
				}
				if maybeLog.err != nil {
					eventChan <- arbbridge.MaybeEvent{Err: err}
					return
				}
				event, err := vm.processEvents(ctx, maybeLog.log)
				if err != nil {
					eventChan <- arbbridge.MaybeEvent{Err: err}
					return
				}
				eventChan <- arbbridge.MaybeEvent{Event: event}
			case maybeLog, ok := <-inboxMaybeLogChan:
				log.Println("event check 2222")
				if !ok {
					eventChan <- arbbridge.MaybeEvent{Err: errors.New("inboxMaybeLogChan terminated early")}
					return
				}
				if maybeLog.err != nil {
					eventChan <- arbbridge.MaybeEvent{Err: err}
					return
				}
				event, err := vm.processEvents(ctx, maybeLog.log)
				if err != nil {
					eventChan <- arbbridge.MaybeEvent{Err: err}
					return
				}
				eventChan <- arbbridge.MaybeEvent{Event: event}
			case err := <-headersSub.Err():
				log.Println("event check 33333")
				eventChan <- arbbridge.MaybeEvent{Err: err}
				return
			}
		}
	}()
	return eventChan, nil
}

func AddressToIntValue(address common.Address) value.IntValue {
	addressBytes := [32]byte{}
	copy(addressBytes[12:], address[:])
	addressVal := big.NewInt(0).SetBytes(addressBytes[:])

	return value.NewIntValue(addressVal)
}

func (vm *ethRollupWatcher) ProcessMessageDeliveredEvents(chainInfo arbbridge.ChainInfo, ethLog types.Log) (arbbridge.Event, error) {
	log.Println("event check")
	if ethLog.Topics[0] == transactionMessageDeliveredID {
		val, err := vm.GlobalPendingInbox.ParseTransactionMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}

		rd := bytes.NewReader(val.Data)
		msgData, err := value.UnmarshalValue(rd)
		if err != nil {
			return nil, err
		}

		messageHash := hashing.SoliditySHA3(
			hashing.Address(common.NewAddressFromEth(val.VmReceiverId)),
			hashing.Address(common.NewAddressFromEth(val.VmSenderId)),
			hashing.Uint256(val.SeqNumber),
			hashing.Uint256(val.Value),
			hashing.Bytes32(msgData.Hash()),
		)

		msgHashInt := new(big.Int).SetBytes(messageHash.Bytes())

		msgVal, _ := value.NewTupleFromSlice([]value.Value{
			value.NewIntValue(val.SeqNumber),
			value.NewIntValue(val.Value),
			msgData,
		})

		msgType, _ := value.NewTupleFromSlice([]value.Value{
			value.NewIntValue(big.NewInt(0)),
			AddressToIntValue(common.NewAddressFromEth(val.VmSenderId)),
			msgVal,
		})

		dataMsg, _ := value.NewTupleFromSlice([]value.Value{
			value.NewIntValue(new(big.Int).SetUint64(ethLog.BlockNumber)),
			value.NewIntValue(msgHashInt),
			msgType,
		})

		return arbbridge.MessageDeliveredEvent{
			ChainInfo: chainInfo,
			MsgValue:  dataMsg,
		}, nil

	} else if ethLog.Topics[0] == ethDepositMessageDeliveredID {
		val, err := vm.GlobalPendingInbox.ParseEthDepositMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}

		messageHash := hashing.SoliditySHA3(
			hashing.Address(common.NewAddressFromEth(val.VmReceiverId)),
			hashing.Address(common.NewAddressFromEth(val.Sender)),
			hashing.Address(common.NewAddressFromEth(val.Destination)),
			hashing.Uint256(val.Value),
		)

		msgHashInt := new(big.Int).SetBytes(messageHash.Bytes())

		msgVal, _ := value.NewTupleFromSlice([]value.Value{
			AddressToIntValue(common.NewAddressFromEth(val.Destination)),
			value.NewIntValue(val.Value),
		})

		msgType, _ := value.NewTupleFromSlice([]value.Value{
			value.NewIntValue(big.NewInt(1)),
			AddressToIntValue(common.NewAddressFromEth(val.Sender)),
			msgVal,
		})

		dataMsg, _ := value.NewTupleFromSlice([]value.Value{
			value.NewIntValue(new(big.Int).SetUint64(ethLog.BlockNumber)),
			value.NewIntValue(msgHashInt),
			msgType,
		})

		return arbbridge.MessageDeliveredEvent{
			ChainInfo: chainInfo,
			MsgValue:  dataMsg,
		}, nil

	} else if ethLog.Topics[0] == depositERC20MessageDeliveredID {
		val, err := vm.GlobalPendingInbox.ParseERC20DepositMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}

		log.Printf("event")

		log.Printf("%+v\n", val)

		messageHash := hashing.SoliditySHA3(
			hashing.Address(common.NewAddressFromEth(val.VmReceiverId)),
			hashing.Address(common.NewAddressFromEth(val.Sender)),
			hashing.Address(common.NewAddressFromEth(val.Destination)),
			hashing.Address(common.NewAddressFromEth(val.TokenAddress)),
			hashing.Uint256(val.Value),
		)

		msgHashInt := new(big.Int).SetBytes(messageHash.Bytes())

		msgVal, _ := value.NewTupleFromSlice([]value.Value{
			AddressToIntValue(common.NewAddressFromEth(val.TokenAddress)),
			AddressToIntValue(common.NewAddressFromEth(val.Destination)),
			value.NewIntValue(val.Value),
		})

		msgType, _ := value.NewTupleFromSlice([]value.Value{
			value.NewIntValue(big.NewInt(2)),
			AddressToIntValue(common.NewAddressFromEth(val.Sender)),
			msgVal,
		})

		dataMsg, _ := value.NewTupleFromSlice([]value.Value{
			value.NewIntValue(new(big.Int).SetUint64(ethLog.BlockNumber)),
			value.NewIntValue(msgHashInt),
			msgType,
		})

		return arbbridge.MessageDeliveredEvent{
			ChainInfo: chainInfo,
			MsgValue:  dataMsg,
		}, nil

	} else if ethLog.Topics[0] == depositERC721MessageDeliveredID {
		val, err := vm.GlobalPendingInbox.ParseERC721DepositMessageDelivered(ethLog)
		if err != nil {
			return nil, err
		}

		messageHash := hashing.SoliditySHA3(
			hashing.Address(common.NewAddressFromEth(val.VmReceiverId)),
			hashing.Address(common.NewAddressFromEth(val.Sender)),
			hashing.Address(common.NewAddressFromEth(val.Destination)),
			hashing.Address(common.NewAddressFromEth(val.TokenAddress)),
			hashing.Uint256(val.Value),
		)

		msgHashInt := new(big.Int).SetBytes(messageHash.Bytes())

		msgVal, _ := value.NewTupleFromSlice([]value.Value{
			AddressToIntValue(common.NewAddressFromEth(val.TokenAddress)),
			AddressToIntValue(common.NewAddressFromEth(val.Destination)),
			value.NewIntValue(val.Value),
		})

		msgType, _ := value.NewTupleFromSlice([]value.Value{
			value.NewIntValue(big.NewInt(3)),
			AddressToIntValue(common.NewAddressFromEth(val.Sender)),
			msgVal,
		})

		dataMsg, _ := value.NewTupleFromSlice([]value.Value{
			value.NewIntValue(new(big.Int).SetUint64(ethLog.BlockNumber)),
			value.NewIntValue(msgHashInt),
			msgType,
		})

		return arbbridge.MessageDeliveredEvent{
			ChainInfo: chainInfo,
			MsgValue:  dataMsg,
		}, nil

	} else {
		return nil, errors2.New("unknown arbitrum event type")
	}
}

func (vm *ethRollupWatcher) processEvents(ctx context.Context, ethLog types.Log) (arbbridge.Event, error) {

	log.Println("processEvents: ", ethLog.Topics[0], " message")

	header, err := vm.client.HeaderByHash(ctx, ethLog.BlockHash)
	if err != nil {
		return nil, err
	}
	chainInfo := getChainInfo(ethLog, header)
	if ethLog.Topics[0] == rollupStakeCreatedID {
		eventVal, err := vm.ArbRollup.ParseRollupStakeCreated(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.StakeCreatedEvent{
			ChainInfo: chainInfo,
			Staker:    common.NewAddressFromEth(eventVal.Staker),
			NodeHash:  eventVal.NodeHash,
		}, nil
	} else if ethLog.Topics[0] == rollupChallengeStartedID {
		eventVal, err := vm.ArbRollup.ParseRollupChallengeStarted(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.ChallengeStartedEvent{
			ChainInfo:         chainInfo,
			Asserter:          common.NewAddressFromEth(eventVal.Asserter),
			Challenger:        common.NewAddressFromEth(eventVal.Challenger),
			ChallengeType:     structures.ChildType(eventVal.ChallengeType.Uint64()),
			ChallengeContract: common.NewAddressFromEth(eventVal.ChallengeContract),
		}, nil
	} else if ethLog.Topics[0] == rollupChallengeCompletedID {
		eventVal, err := vm.ArbRollup.ParseRollupChallengeCompleted(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.ChallengeCompletedEvent{
			ChainInfo:         chainInfo,
			Winner:            common.NewAddressFromEth(eventVal.Winner),
			Loser:             common.NewAddressFromEth(eventVal.Loser),
			ChallengeContract: common.NewAddressFromEth(eventVal.ChallengeContract),
		}, nil
	} else if ethLog.Topics[0] == rollupRefundedID {
		eventVal, err := vm.ArbRollup.ParseRollupStakeRefunded(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.StakeRefundedEvent{
			ChainInfo: chainInfo,
			Staker:    common.NewAddressFromEth(eventVal.Staker),
		}, nil
	} else if ethLog.Topics[0] == rollupPrunedID {
		eventVal, err := vm.ArbRollup.ParseRollupPruned(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.PrunedEvent{
			ChainInfo: chainInfo,
			Leaf:      eventVal.Leaf,
		}, nil
	} else if ethLog.Topics[0] == rollupStakeMovedID {
		eventVal, err := vm.ArbRollup.ParseRollupStakeMoved(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.StakeMovedEvent{
			ChainInfo: chainInfo,
			Staker:    common.NewAddressFromEth(eventVal.Staker),
			Location:  eventVal.ToNodeHash,
		}, nil
	} else if ethLog.Topics[0] == rollupAssertedID {
		eventVal, err := vm.ArbRollup.ParseRollupAsserted(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.AssertedEvent{
			ChainInfo:    chainInfo,
			PrevLeafHash: eventVal.Fields[0],
			Params: &structures.AssertionParams{
				NumSteps: eventVal.NumSteps,
				TimeBounds: &protocol.TimeBoundsBlocks{
					common.NewTimeBlocks(eventVal.TimeBoundsBlocks[0]),
					common.NewTimeBlocks(eventVal.TimeBoundsBlocks[1]),
				},
				ImportedMessageCount: eventVal.ImportedMessageCount,
			},
			Claim: &structures.AssertionClaim{
				AfterPendingTop:       eventVal.Fields[2],
				ImportedMessagesSlice: eventVal.Fields[3],
				AssertionStub: &valprotocol.ExecutionAssertionStub{
					AfterHash:        eventVal.Fields[4],
					DidInboxInsn:     eventVal.DidInboxInsn,
					NumGas:           eventVal.NumArbGas,
					FirstMessageHash: [32]byte{},
					LastMessageHash:  eventVal.Fields[5],
					FirstLogHash:     [32]byte{},
					LastLogHash:      eventVal.Fields[6],
				},
			},
			MaxPendingTop:   eventVal.Fields[1],
			MaxPendingCount: eventVal.PendingCount,
		}, nil
	} else if ethLog.Topics[0] == rollupConfirmedID {
		eventVal, err := vm.ArbRollup.ParseRollupConfirmed(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.ConfirmedEvent{
			ChainInfo: chainInfo,
			NodeHash:  eventVal.NodeHash,
		}, nil
	} else if ethLog.Topics[0] == confirmedAssertionID {
		eventVal, err := vm.ArbRollup.ParseConfirmedAssertion(ethLog)
		if err != nil {
			return nil, err
		}
		return arbbridge.ConfirmedAssertionEvent{
			LogsAccHash: eventVal.LogsAccHash,
		}, nil
	}
	return vm.ProcessMessageDeliveredEvents(chainInfo, ethLog)
}

func (vm *ethRollupWatcher) GetParams(ctx context.Context) (structures.ChainParams, error) {
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
		GracePeriod:             common.TimeTicks{rawParams.GracePeriodTicks},
		MaxExecutionSteps:       rawParams.MaxExecutionSteps,
		ArbGasSpeedLimitPerTick: rawParams.ArbGasSpeedLimitPerTick.Uint64(),
	}, nil
}

func (vm *ethRollupWatcher) InboxAddress(ctx context.Context) (common.Address, error) {
	addr, err := vm.ArbRollup.GlobalInbox(nil)
	return common.NewAddressFromEth(addr), err
}
