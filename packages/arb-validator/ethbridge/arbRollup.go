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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/rollup"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/hashing"
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

func init() {
	parsed, err := abi.JSON(strings.NewReader(rollup.ArbRollupABI))
	if err != nil {
		panic(err)
	}
	rollupStakeCreatedID = parsed.Events["RollupStakeCreated"].ID()
	rollupChallengeStartedID = parsed.Events["RollupChallengeStarted"].ID()
	rollupChallengeCompletedID = parsed.Events["RollupChallengeCompleted"].ID()
	rollupRefundedID = parsed.Events["RollupStakeRefunded"].ID()
	rollupPrunedID = parsed.Events["RollupPruned"].ID()
	rollupStakeMovedID = parsed.Events["RollupStakeMoved"].ID()
	rollupAssertedID = parsed.Events["RollupAsserted"].ID()
	rollupConfirmedID = parsed.Events["RollupConfirmed"].ID()
	confirmedAssertionID = parsed.Events["ConfirmedAssertion"].ID()
}

type ArbRollup struct {
	OutChan            chan Notification
	ErrChan            chan error
	Client             *ethclient.Client
	ArbRollup          *rollup.ArbRollup
	GlobalPendingInbox *rollup.IGlobalPendingInbox

	address common.Address
	client  *ethclient.Client
}

func NewRollup(address common.Address, client *ethclient.Client) (*ArbRollup, error) {
	outChan := make(chan Notification, 1024)
	errChan := make(chan error, 1024)
	vm := &ArbRollup{OutChan: outChan, ErrChan: errChan, Client: client, address: address}
	err := vm.setupContracts()
	return vm, err
}

func (vm *ArbRollup) setupContracts() error {
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
	globalPendingContract, err := rollup.NewIGlobalPendingInbox(globalPendingInboxAddress, vm.Client)
	if err != nil {
		return errors2.Wrap(err, "Failed to connect to GlobalPendingInbox")
	}

	vm.ArbRollup = arbitrumRollupContract
	vm.GlobalPendingInbox = globalPendingContract
	return nil
}

func (vm *ArbRollup) GetChans() (chan Notification, chan error) {
	return vm.OutChan, vm.ErrChan
}

func (vm *ArbRollup) Close() {
	close(vm.OutChan)
	close(vm.ErrChan)
}

func (vm *ArbRollup) StartConnection(ctx context.Context) error {
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

	logChan := make(chan types.Log)
	logSub, err := vm.Client.SubscribeFilterLogs(ctx, filter, logChan)
	if err != nil {
		return err
	}

	messageDeliveredChan := make(chan *rollup.IGlobalPendingInboxMessageDelivered)
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

func (vm *ArbRollup) processEvents(ctx context.Context, log types.Log) error {
	event, err := func() (Event, error) {
		if log.Topics[0] == rollupStakeCreatedID {
			eventVal, err := vm.ArbRollup.ParseRollupStakeCreated(log)
			if err != nil {
				return nil, err
			}
			return StakeCreatedEvent{
				Staker:   eventVal.Staker,
				NodeHash: eventVal.NodeHash,
			}, nil
		} else if log.Topics[0] == rollupChallengeStartedID {
			eventVal, err := vm.ArbRollup.ParseRollupChallengeStarted(log)
			if err != nil {
				return nil, err
			}
			return ChallengeStartedEvent{
				Asserter:          eventVal.Asserter,
				Challenger:        eventVal.Challenger,
				ChallengeType:     eventVal.ChallengeType.Uint64(),
				ChallengeContract: eventVal.ChallengeContract,
			}, nil
		} else if log.Topics[0] == rollupChallengeCompletedID {
			eventVal, err := vm.ArbRollup.ParseRollupChallengeCompleted(log)
			if err != nil {
				return nil, err
			}
			return ChallengeCompletedEvent{
				Winner:            eventVal.Winner,
				Loser:             eventVal.Loser,
				ChallengeContract: eventVal.ChallengeContract,
			}, nil
		} else if log.Topics[0] == rollupRefundedID {
			eventVal, err := vm.ArbRollup.ParseRollupStakeRefunded(log)
			if err != nil {
				return nil, err
			}
			return StakeRefundedEvent{
				Staker: eventVal.Staker,
			}, nil
		} else if log.Topics[0] == rollupPrunedID {
			eventVal, err := vm.ArbRollup.ParseRollupPruned(log)
			if err != nil {
				return nil, err
			}
			return PrunedEvent{
				Leaf: eventVal.Leaf,
			}, nil
		} else if log.Topics[0] == rollupStakeMovedID {
			eventVal, err := vm.ArbRollup.ParseRollupStakeMoved(log)
			if err != nil {
				return nil, err
			}
			return StakeMovedEvent{
				Staker:   eventVal.Staker,
				Location: eventVal.ToNodeHash,
			}, nil
		} else if log.Topics[0] == rollupAssertedID {
			eventVal, err := vm.ArbRollup.ParseRollupAsserted(log)
			if err != nil {
				return nil, err
			}
			return AssertedEvent{
				PrevLeafHash:          eventVal.Fields[0],
				TimeBoundsBlocks:      eventVal.TimeBoundsBlocks,
				AfterPendingTop:       eventVal.Fields[1],
				ImportedMessagesSlice: eventVal.Fields[2],
				ImportedMessageCount:  eventVal.ImportedMessageCount,
				Assertion: protocol.NewAssertionStub(
					eventVal.Fields[3],
					eventVal.DidInboxInsn,
					eventVal.NumSteps,
					eventVal.NumArbGas,
					eventVal.Fields[4],
					eventVal.Fields[5],
				),
			}, nil
		} else if log.Topics[0] == rollupConfirmedID {
			eventVal, err := vm.ArbRollup.ParseRollupConfirmed(log)
			if err != nil {
				return nil, err
			}
			return ConfirmedEvent{
				NodeHash: eventVal.NodeHash,
			}, nil
		} else if log.Topics[0] == confirmedAssertionID {
			eventVal, err := vm.ArbRollup.ParseConfirmedAssertion(log)
			if err != nil {
				return nil, err
			}
			return ConfirmedAssertionEvent{
				LogsAccHash: eventVal.LogsAccHash,
			}, nil
		}
		return nil, errors2.New("unknown arbitrum event type")
	}()

	if err != nil {
		return err
	}
	header, err := vm.Client.HeaderByHash(ctx, log.BlockHash)
	if err != nil {
		return err
	}
	vm.OutChan <- Notification{
		Header: header,
		VMID:   vm.address,
		Event:  event,
		TxHash: log.TxHash,
	}

	return nil
}

func protoStateHash(
	machineHash [32]byte,
	inboxHash [32]byte,
	pendingTop [32]byte,
	pendingCount *big.Int,
) [32]byte {
	var ret [32]byte
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(machineHash),
		solsha3.Bytes32(inboxHash),
		solsha3.Bytes32(pendingTop),
		solsha3.Uint256(pendingCount),
	))
	return ret
}

func (vm *ArbRollup) PlaceStake(
	auth *bind.TransactOpts,
	stakeAmount *big.Int,
	location [32]byte,
	leaf [32]byte,
	proof1 [][32]byte,
	proof2 [][32]byte,
) (*types.Receipt, error) {
	call := &bind.TransactOpts{
		From:     auth.From,
		Nonce:    auth.Nonce,
		Signer:   auth.Signer,
		Value:    stakeAmount,
		GasPrice: auth.GasPrice,
		GasLimit: 100000,
		Context:  auth.Context,
	}
	tx, err := vm.ArbRollup.PlaceStake(
		call,
		location,
		leaf,
		proof1,
		proof2,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, auth, tx, "PlaceStake")
}

func (vm *ArbRollup) RecoverStakeConfirmed(
	auth *bind.TransactOpts,
	proof [][32]byte,
) (*types.Receipt, error) {
	tx, err := vm.ArbRollup.RecoverStakeConfirmed(
		auth,
		proof,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, auth, tx, "RecoverStakeConfirmed")
}

func (vm *ArbRollup) RecoverStakeOld(
	auth *bind.TransactOpts,
	staker common.Address,
	proof [][32]byte,
) (*types.Receipt, error) {
	tx, err := vm.ArbRollup.RecoverStakeOld(
		auth,
		staker,
		proof,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, auth, tx, "RecoverStakeOld")
}

func (vm *ArbRollup) RecoverStakeMooted(
	auth *bind.TransactOpts,
	disputableHash [32]byte,
	staker common.Address,
	latestConfirmedProof [][32]byte,
	nodeProof [][32]byte,
) (*types.Receipt, error) {
	tx, err := vm.ArbRollup.RecoverStakeMooted(
		auth,
		staker,
		disputableHash,
		latestConfirmedProof,
		nodeProof,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, auth, tx, "RecoverStakeMooted")
}

func (vm *ArbRollup) RecoverStakePassedDeadline(
	auth *bind.TransactOpts,
	stakerAddress common.Address,
	deadlineTicks *big.Int,
	disputableNodeHashVal [32]byte,
	childType uint64,
	vmProtoStateHash [32]byte,
	leaf [32]byte,
	proof [][32]byte,
) (*types.Receipt, error) {
	tx, err := vm.ArbRollup.RecoverStakePassedDeadline(
		auth,
		stakerAddress,
		deadlineTicks,
		disputableNodeHashVal,
		new(big.Int).SetUint64(childType),
		vmProtoStateHash,
		leaf,
		proof,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, auth, tx, "RecoverStakePassedDeadline")
}

func (vm *ArbRollup) MoveStake(
	auth *bind.TransactOpts,
	newLocation [32]byte,
	leaf [32]byte,
	proof1 [][32]byte,
	proof2 [][32]byte,
) (*types.Receipt, error) {
	tx, err := vm.ArbRollup.MoveStake(
		auth,
		newLocation,
		leaf,
		proof1,
		proof2,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, auth, tx, "MoveStake")
}

func (vm *ArbRollup) PruneLeaf(
	auth *bind.TransactOpts,
	leaf [32]byte,
	from [32]byte,
	proof1 [][32]byte,
	proof2 [][32]byte,
) (*types.Receipt, error) {
	tx, err := vm.ArbRollup.PruneLeaf(
		auth,
		leaf,
		from,
		proof1,
		proof2,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, auth, tx, "PruneLeaf")
}

type VMProtoState struct {
	VMHash       [32]byte
	InboxHash    [32]byte
	PendingTop   [32]byte
	PendingCount *big.Int
}

type AssertionParams struct {
	timeBoundsBlocks  [2]*big.Int
	afterPendingCount *big.Int
	numSteps          uint32
}

type PendingTopOutput struct {
	AfterPendingTop [32]byte
}

type MessagesOutput struct {
	ImportedMessagesSlice [32]byte
}

type ExecutionOutput struct {
	vmHash          [32]byte
	didInboxInsn    bool
	numArbGas       uint64
	messagesAccHash [32]byte
	logsAccHash     [32]byte
}

func (vm *ArbRollup) MakeAssertion(
	auth *bind.TransactOpts,

	prevPrevLeafHash [32]byte,
	prevDisputableNodeHash [32]byte,
	prevDeadlineTicks *big.Int,
	prevChildType uint32,

	beforeState VMProtoState,
	assertionParams AssertionParams,
	pendingTopOutput PendingTopOutput,
	messagesOutput MessagesOutput,
	executionOutput ExecutionOutput,
	stakerProof [][32]byte,

) (*types.Receipt, error) {
	tx, err := vm.ArbRollup.MakeAssertion(
		auth,
		[10][32]byte{
			beforeState.VMHash,
			beforeState.InboxHash,
			beforeState.PendingTop,
			prevPrevLeafHash,
			prevDisputableNodeHash,
			pendingTopOutput.AfterPendingTop,
			messagesOutput.ImportedMessagesSlice,
			executionOutput.vmHash,
			executionOutput.messagesAccHash,
			executionOutput.logsAccHash,
		},

		beforeState.PendingCount,
		prevDeadlineTicks,
		prevChildType,
		assertionParams.numSteps,
		assertionParams.timeBoundsBlocks,
		assertionParams.afterPendingCount,
		executionOutput.didInboxInsn,
		executionOutput.numArbGas,
		stakerProof,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, auth, tx, "MakeAssertion")
}

func (vm *ArbRollup) ConfirmValid(
	auth *bind.TransactOpts,
	deadlineTics *big.Int,
	outMsgs []value.Value,
	logsAccHash [32]byte,
	protoHash [32]byte,
	stakerAddresses []common.Address,
	stakerProofs [][32]byte,
	stakerProofOffsets []*big.Int,
) (*types.Receipt, error) {
	messages := hashing.CombineMessages(outMsgs)
	tx, err := vm.ArbRollup.ConfirmValid(
		auth,
		deadlineTics,
		messages,
		logsAccHash,
		protoHash,
		stakerAddresses,
		stakerProofs,
		stakerProofOffsets,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, auth, tx, "ConfirmValid")
}

func (vm *ArbRollup) ConfirmInvalid(
	auth *bind.TransactOpts,
	deadlineTics *big.Int,
	challengeNodeData [32]byte,
	branch uint64,
	protoHash [32]byte,
	stakerAddresses []common.Address,
	stakerProofs [][32]byte,
	stakerProofOffsets []*big.Int,
) (*types.Receipt, error) {
	tx, err := vm.ArbRollup.ConfirmInvalid(
		auth,
		deadlineTics,
		challengeNodeData,
		new(big.Int).SetUint64(branch),
		protoHash,
		stakerAddresses,
		stakerProofs,
		stakerProofOffsets,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, auth, tx, "ConfirmInvalid")
}

func (vm *ArbRollup) StartChallenge(
	auth *bind.TransactOpts,
	asserterAddress common.Address,
	challengerAddress common.Address,
	node [32]byte,
	disputableDeadline *big.Int,
	staker1Position uint64,
	staker2Position uint64,
	vmProtoHash1 [32]byte,
	vmProtoHash2 [32]byte,
	proof1 [][32]byte,
	proof2 [][32]byte,
	challenge1DataHash [32]byte,
	challenge1PeriodTicks *big.Int,
	challenge2NodeHash [32]byte,
) (*types.Receipt, error) {
	tx, err := vm.ArbRollup.StartChallenge(
		auth,
		asserterAddress,
		challengerAddress,
		node,
		disputableDeadline,
		[2]*big.Int{
			new(big.Int).SetUint64(staker1Position),
			new(big.Int).SetUint64(staker2Position),
		},
		[2][32]byte{
			vmProtoHash1,
			vmProtoHash2,
		},
		proof1,
		proof2,
		challenge1DataHash,
		challenge1PeriodTicks,
		challenge2NodeHash,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, auth, tx, "StartExecutionChallenge")
}

//func (vm *ArbRollup) VerifyVM(
//	auth *bind.CallOpts,
//	config *valmessage.VMConfiguration,
//	machine [32]byte,
//) error {
//	//code, err := vm.contract.Client.CodeAt(auth.Context, vm.address, nil)
//	// Verify that VM has correct code
//	vmInfo, err := vm.ArbRollup.Vm(auth)
//	if err != nil {
//		return err
//	}
//
//	if vmInfo.MachineHash != machine {
//		return errors.New("VM has different machine hash")
//	}
//
//	if config.GracePeriod != uint64(vmInfo.GracePeriod) {
//		return errors.New("VM has different grace period")
//	}
//
//	if value.NewBigIntFromBuf(config.EscrowRequired).Cmp(vmInfo.EscrowRequired) != 0 {
//		return errors.New("VM has different escrow required")
//	}
//
//	if config.MaxExecutionStepCount != vmInfo.MaxExecutionSteps {
//		return errors.New("VM has different mxa steps")
//	}
//
//	owner, err := vm.ArbRollup.Owner(auth)
//	if err != nil {
//		return err
//	}
//	if protocol.NewAddressFromBuf(config.Owner) != owner {
//		return errors.New("VM has different owner")
//	}
//	return nil
//}
