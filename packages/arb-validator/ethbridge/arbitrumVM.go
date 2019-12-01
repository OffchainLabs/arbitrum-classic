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
	"bytes"
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/arbchain"

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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

var pendingDisputableAssertionID common.Hash
var confirmedDisputableAssertionID common.Hash
var challengeLaunchedID common.Hash

func init() {
	parsed, err := abi.JSON(strings.NewReader(arbchain.ArbitrumVMABI))
	if err != nil {
		panic(err)
	}
	pendingDisputableAssertionID = parsed.Events["PendingDisputableAssertion"].ID()
	confirmedDisputableAssertionID = parsed.Events["ConfirmedDisputableAssertion"].ID()
	challengeLaunchedID = parsed.Events["ChallengeLaunched"].ID()
}

type ArbitrumVM struct {
	OutChan            chan Notification
	ErrChan            chan error
	Client             *ethclient.Client
	ArbitrumVM         *arbchain.ArbitrumVM
	GlobalPendingInbox *arbchain.IGlobalPendingInbox

	address common.Address
	client  *ethclient.Client
}

func NewArbitrumVM(address common.Address, client *ethclient.Client) (*ArbitrumVM, error) {
	outChan := make(chan Notification, 1024)
	errChan := make(chan error, 1024)
	vm := &ArbitrumVM{OutChan: outChan, ErrChan: errChan, Client: client, address: address}
	err := vm.setupContracts()
	return vm, err
}

func (vm *ArbitrumVM) setupContracts() error {
	arbitrumVMContract, err := arbchain.NewArbitrumVM(vm.address, vm.Client)
	if err != nil {
		return errors2.Wrap(err, "Failed to connect to ArbChannel")
	}

	globalPendingInboxAddress, err := arbitrumVMContract.GlobalInbox(&bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	})
	if err != nil {
		return errors2.Wrap(err, "Failed to get GlobalPendingInbox address")
	}
	globalPendingContract, err := arbchain.NewIGlobalPendingInbox(globalPendingInboxAddress, vm.Client)
	if err != nil {
		return errors2.Wrap(err, "Failed to connect to GlobalPendingInbox")
	}

	vm.ArbitrumVM = arbitrumVMContract
	vm.GlobalPendingInbox = globalPendingContract
	return nil
}

func (vm *ArbitrumVM) GetChans() (chan Notification, chan error) {
	return vm.OutChan, vm.ErrChan
}

func (vm *ArbitrumVM) Close() {
	close(vm.OutChan)
	close(vm.ErrChan)
}

func (vm *ArbitrumVM) StartConnection(ctx context.Context) error {
	if err := vm.setupContracts(); err != nil {
		return err
	}

	start := uint64(0)
	watch := &bind.WatchOpts{
		Context: ctx,
		Start:   &start,
	}

	headers := make(chan *types.Header)
	headersSub, err := vm.Client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return err
	}

	filter := ethereum.FilterQuery{
		Addresses: []common.Address{vm.address},
		Topics: [][]common.Hash{
			{pendingDisputableAssertionID, confirmedDisputableAssertionID, challengeLaunchedID},
		},
	}

	logChan := make(chan types.Log)
	logSub, err := vm.Client.SubscribeFilterLogs(ctx, filter, logChan)
	if err != nil {
		return err
	}

	messageDeliveredChan := make(chan *arbchain.IGlobalPendingInboxMessageDelivered)
	messageDeliveredSub, err := vm.GlobalPendingInbox.WatchMessageDelivered(watch, messageDeliveredChan, []common.Address{vm.address})
	if err != nil {
		return err
	}

	go func() {
		defer headersSub.Unsubscribe()
		defer messageDeliveredSub.Unsubscribe()
		defer logSub.Unsubscribe()

		for {
			select {
			case <-ctx.Done():
				break
			case header := <-headers:
				vm.OutChan <- Notification{
					Header: header,
					Event:  NewTimeEvent{},
				}
			case val := <-messageDeliveredChan:
				header, err := vm.Client.HeaderByHash(context.Background(), val.Raw.BlockHash)
				if err != nil {
					vm.ErrChan <- err
					return
				}
				rd := bytes.NewReader(val.Data)
				msgData, err := value.UnmarshalValue(rd)
				if err != nil {
					vm.ErrChan <- err
					return
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
					value.NewIntValue(new(big.Int).SetUint64(header.Time)),
					value.NewIntValue(header.Number),
					value.NewIntValue(msgHashInt),
				})

				msg := protocol.NewSimpleMessage(msgVal, val.TokenType, val.Value, val.Sender)
				vm.OutChan <- Notification{
					Header: header,
					VMID:   val.VmId,
					Event: MessageDeliveredEvent{
						Msg: msg,
					},
					TxHash: val.Raw.TxHash,
				}
			case log := <-logChan:
				if err := vm.processEvents(ctx, log); err != nil {
					vm.ErrChan <- err
					return
				}
			case err := <-headersSub.Err():
				vm.ErrChan <- err
				return
			case err := <-messageDeliveredSub.Err():
				vm.ErrChan <- err
				return
			case err := <-logSub.Err():
				vm.ErrChan <- err
				return
			}
		}
	}()
	return nil
}

func (vm *ArbitrumVM) processEvents(ctx context.Context, log types.Log) error {
	header, err := vm.Client.HeaderByHash(ctx, log.BlockHash)
	if err != nil {
		return err
	}
	if log.Topics[0] == pendingDisputableAssertionID {
		pendingDisVal, err := vm.ArbitrumVM.ParsePendingDisputableAssertion(log)
		if err != nil {
			return err
		}
		precondition, assertion := translateDisputableAssertionEvent(pendingDisVal)
		vm.OutChan <- Notification{
			Header: header,
			VMID:   vm.address,
			Event: PendingDisputableAssertionEvent{
				Precondition: precondition,
				Assertion:    assertion,
				Asserter:     pendingDisVal.Asserter,
				Deadline:     pendingDisVal.Deadline,
			},
			TxHash: log.TxHash,
		}
	} else if log.Topics[0] == confirmedDisputableAssertionID {
		confDisVal, err := vm.ArbitrumVM.ParseConfirmedDisputableAssertion(log)
		if err != nil {
			return err
		}
		vm.OutChan <- Notification{
			Header: header,
			VMID:   vm.address,
			Event: ConfirmedDisputableAssertEvent{
				log.TxHash,
				confDisVal.LogsAccHash,
			},
			TxHash: log.TxHash,
		}
	} else if log.Topics[0] == challengeLaunchedID {
		challengedVal, err := vm.ArbitrumVM.ParseChallengeLaunched(log)
		if err != nil {
			return err
		}
		vm.OutChan <- Notification{
			Header: header,
			VMID:   vm.address,
			Event: ChallengeLaunchedEvent{
				ChallengeAddress: challengedVal.ChallengeContract,
				Challenger:       challengedVal.Challenger,
			},
			TxHash: log.TxHash,
		}
	}
	return nil
}

func (vm *ArbitrumVM) PendingDisputableAssert(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertion *protocol.Assertion,
) (*types.Receipt, error) {
	stub := assertion.Stub()
	tx, err := vm.ArbitrumVM.PendingDisputableAssert(
		auth,
		precondition.BeforeHashValue(),
		precondition.BeforeInboxValue(),
		stub.AfterHashValue(),
		stub.LastMessageHashValue(),
		stub.LastLogHashValue(),
		assertion.NumSteps,
		[2]uint64{precondition.TimeBounds.StartTime, precondition.TimeBounds.EndTime},
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, tx.Hash(), "PendingDisputableAssert")
}

func (vm *ArbitrumVM) ConfirmDisputableAsserted(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertion *protocol.Assertion,
) (*types.Receipt, error) {
	messages := hashing.CombineMessages(assertion.OutMsgs)

	tx, err := vm.ArbitrumVM.ConfirmDisputableAsserted(
		auth,
		precondition.Hash(),
		assertion.AfterHash,
		assertion.NumSteps,
		messages,
		assertion.LogsHash(),
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, tx.Hash(), "ConfirmDisputableAsserted")
}

func (vm *ArbitrumVM) InitiateChallenge(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
) (*types.Receipt, error) {
	tx, err := vm.ArbitrumVM.InitiateChallenge(
		auth,
		precondition.BeforeHashValue(),
		precondition.BeforeInboxValue(),
		[2]uint64{precondition.TimeBounds.StartTime, precondition.TimeBounds.EndTime},
		assertion.Hash(),
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, tx.Hash(), "InitiateChallenge")
}

func (vm *ArbitrumVM) CurrentDeposit(
	auth *bind.CallOpts,
	address common.Address,
) (*big.Int, error) {
	return vm.ArbitrumVM.CurrentDeposit(auth, address)
}

func (vm *ArbitrumVM) EscrowRequired(
	auth *bind.CallOpts,
) (*big.Int, error) {
	return vm.ArbitrumVM.EscrowRequired(auth)
}

func (vm *ArbitrumVM) IsEnabled(
	auth *bind.CallOpts,
) (bool, error) {
	status, err := vm.ArbitrumVM.GetState(auth)
	return status != 0, err
}

func (vm *ArbitrumVM) IsInChallenge(
	auth *bind.CallOpts,
) (bool, error) {
	vmState, err := vm.ArbitrumVM.Vm(auth)
	if err != nil {
		return false, err
	}
	return vmState.ActiveChallengeManager != [20]byte{}, nil
}

func (vm *ArbitrumVM) IsPendingUnanimous(
	auth *bind.CallOpts,
) (bool, error) {
	status, err := vm.ArbitrumVM.GetState(auth)
	return status == 3, err
}

func (vm *ArbitrumVM) VerifyVM(
	auth *bind.CallOpts,
	config *valmessage.VMConfiguration,
	machine [32]byte,
) error {
	//code, err := vm.contract.Client.CodeAt(auth.Context, vm.address, nil)
	// Verify that VM has correct code
	vmInfo, err := vm.ArbitrumVM.Vm(auth)
	if err != nil {
		return err
	}

	if vmInfo.MachineHash != machine {
		return errors.New("VM has different machine hash")
	}

	if config.GracePeriod != uint64(vmInfo.GracePeriod) {
		return errors.New("VM has different grace period")
	}

	if value.NewBigIntFromBuf(config.EscrowRequired).Cmp(vmInfo.EscrowRequired) != 0 {
		return errors.New("VM has different escrow required")
	}

	if config.MaxExecutionStepCount != vmInfo.MaxExecutionSteps {
		return errors.New("VM has different mxa steps")
	}

	owner, err := vm.ArbitrumVM.Owner(auth)
	if err != nil {
		return err
	}
	if protocol.NewAddressFromBuf(config.Owner) != owner {
		return errors.New("VM has different owner")
	}
	return nil
}

func translateDisputableAssertionEvent(event *arbchain.ArbitrumVMPendingDisputableAssertion) (*protocol.Precondition, *protocol.AssertionStub) {
	precondition := protocol.NewPrecondition(
		event.Fields[0],
		protocol.NewTimeBounds(event.TimeBounds[0], event.TimeBounds[1]),
		value.NewHashOnlyValue(event.Fields[1], 1),
	)
	assertion := &protocol.AssertionStub{
		AfterHash:        value.NewHashBuf(event.Fields[2]),
		NumSteps:         event.NumSteps,
		FirstMessageHash: value.NewHashBuf([32]byte{}),
		LastMessageHash:  value.NewHashBuf(event.Fields[3]),
		FirstLogHash:     value.NewHashBuf([32]byte{}),
		LastLogHash:      value.NewHashBuf(event.Fields[4]),
	}
	return precondition, assertion
}
