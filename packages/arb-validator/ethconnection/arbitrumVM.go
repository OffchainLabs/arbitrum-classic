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

package ethconnection

import (
	"bytes"
	"context"
	"errors"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethconnection/arblauncher"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethconnection/challengemanager"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/hashing"
)

type ArbitrumVM struct {
	Client             *ethclient.Client
	Tracker            *arblauncher.ArbitrumVM
	Challenge          *challengemanager.ChallengeManager
	GlobalPendingInbox *arblauncher.IGlobalPendingInbox

	address common.Address
}

func NewArbitrumVM(address common.Address, client *ethclient.Client) (*ArbitrumVM, error) {
	trackerContract, err := arblauncher.NewArbitrumVM(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to ArbChannel")
	}

	challengeManagerAddress, err := trackerContract.ChallengeManager(&bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	})
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to get ChallengeManager address")
	}
	challengeManagerContract, err := challengemanager.NewChallengeManager(challengeManagerAddress, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to ChallengeManager")
	}

	globalPendingInboxAddress, err := trackerContract.GlobalInbox(&bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	})
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to get GlobalPendingInbox address")
	}
	globalPendingContract, err := arblauncher.NewIGlobalPendingInbox(globalPendingInboxAddress, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to GlobalPendingInbox")
	}

	return &ArbitrumVM{client, trackerContract, challengeManagerContract, globalPendingContract, address}, nil
}

func (vm *ArbitrumVM) CreateListeners(ctx context.Context) (chan Notification, chan error, error) {
	outChan := make(chan Notification, 1024)
	errChan := make(chan error, 1024)

	start := uint64(0)
	watch := &bind.WatchOpts{
		Context: ctx,
		Start:   &start,
	}

	headers := make(chan *types.Header)
	headersSub, err := vm.Client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return nil, nil, err
	}

	messageDeliveredChan := make(chan *arblauncher.IGlobalPendingInboxMessageDelivered)
	messageDeliveredSub, err := vm.GlobalPendingInbox.WatchMessageDelivered(watch, messageDeliveredChan, []common.Address{vm.address})
	if err != nil {
		return nil, nil, err
	}

	dispAssChan := make(chan *arblauncher.ArbitrumVMPendingDisputableAssertion)
	dispAssSub, err := vm.Tracker.WatchPendingDisputableAssertion(watch, dispAssChan)
	if err != nil {
		return nil, nil, err
	}

	confAssChan := make(chan *arblauncher.ArbitrumVMConfirmedDisputableAssertion)
	confAssSub, err := vm.Tracker.WatchConfirmedDisputableAssertion(watch, confAssChan)
	if err != nil {
		return nil, nil, err
	}

	challengeInitiatedChan := make(chan *arblauncher.ArbitrumVMInitiatedChallenge)
	challengeInitiatedSub, err := vm.Tracker.WatchInitiatedChallenge(watch, challengeInitiatedChan)
	if err != nil {
		return nil, nil, err
	}

	challengeBisectedChan := make(chan *challengemanager.ChallengeManagerBisectedAssertion)
	challengeBisectedSub, err := vm.Challenge.WatchBisectedAssertion(watch, challengeBisectedChan, []common.Address{vm.address})
	if err != nil {
		return nil, nil, err
	}

	challengeContinuedChan := make(chan *challengemanager.ChallengeManagerContinuedChallenge)
	challengeContinuedSub, err := vm.Challenge.WatchContinuedChallenge(watch, challengeContinuedChan, []common.Address{vm.address})
	if err != nil {
		return nil, nil, err
	}

	challengeTimedOutChan := make(chan *challengemanager.ChallengeManagerTimedOutChallenge)
	challengeTimedOutSub, err := vm.Challenge.WatchTimedOutChallenge(watch, challengeTimedOutChan, []common.Address{vm.address})
	if err != nil {
		return nil, nil, err
	}

	oneStepProofChan := make(chan *challengemanager.ChallengeManagerOneStepProofCompleted)
	oneStepProofSub, err := vm.Challenge.WatchOneStepProofCompleted(watch, oneStepProofChan, []common.Address{vm.address})
	if err != nil {
		return nil, nil, err
	}

	go func() {
		defer headersSub.Unsubscribe()
		defer messageDeliveredSub.Unsubscribe()
		defer dispAssSub.Unsubscribe()
		defer confAssSub.Unsubscribe()
		defer challengeInitiatedSub.Unsubscribe()
		defer challengeBisectedSub.Unsubscribe()
		defer challengeInitiatedSub.Unsubscribe()
		defer challengeContinuedSub.Unsubscribe()
		defer oneStepProofSub.Unsubscribe()

		for {
			select {
			case <-ctx.Done():
				break
			case header := <-headers:
				outChan <- Notification{
					Header: header,
					Event:  NewTimeEvent{},
				}
			case val := <-messageDeliveredChan:
				header, err := vm.Client.HeaderByHash(context.Background(), val.Raw.BlockHash)
				if err != nil {
					errChan <- err
					return
				}
				rd := bytes.NewReader(val.Data)
				msgData, err := value.UnmarshalValue(rd)
				if err != nil {
					errChan <- err
					return
				}

				messageHash := solsha3.SoliditySHA3(
					solsha3.Bytes32(val.VmId),
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
				outChan <- Notification{
					Header: header,
					VMID:   val.VmId,
					Event: MessageDeliveredEvent{
						Msg: msg,
					},
					TxHash: val.Raw.TxHash,
				}
			case val := <-dispAssChan:
				header, err := vm.Client.HeaderByHash(ctx, val.Raw.BlockHash)
				if err != nil {
					errChan <- err
					return
				}

				precondition, assertion := translateDisputableAssertionEvent(val)
				outChan <- Notification{
					Header: header,
					VMID:   vm.address,
					Event: PendingDisputableAssertionEvent{
						Precondition: precondition,
						Assertion:    assertion,
						Asserter:     val.Asserter,
					},
					TxHash: val.Raw.TxHash,
				}
			case val := <-confAssChan:
				header, err := vm.Client.HeaderByHash(ctx, val.Raw.BlockHash)
				if err != nil {
					errChan <- err
					return
				}
				outChan <- Notification{
					Header: header,
					VMID:   vm.address,
					Event: ConfirmedDisputableAssertEvent{
						val.Raw.TxHash,
						val.LogsAccHash,
					},
					TxHash: val.Raw.TxHash,
				}
			case val := <-challengeInitiatedChan:
				header, err := vm.Client.HeaderByHash(ctx, val.Raw.BlockHash)
				if err != nil {
					errChan <- err
					return
				}
				outChan <- Notification{
					Header: header,
					VMID:   vm.address,
					Event: InitiateChallengeEvent{
						Challenger: val.Challenger,
					},
					TxHash: val.Raw.TxHash,
				}
			case val := <-challengeBisectedChan:
				header, err := vm.Client.HeaderByHash(ctx, val.Raw.BlockHash)
				if err != nil {
					errChan <- err
					return
				}
				outChan <- Notification{
					Header: header,
					VMID:   vm.address,
					Event: BisectionEvent{
						Assertions: translateBisectionEvent(val),
					},
					TxHash: val.Raw.TxHash,
				}
			case val := <-challengeTimedOutChan:
				header, err := vm.Client.HeaderByHash(ctx, val.Raw.BlockHash)
				if err != nil {
					errChan <- err
					return
				}
				if val.ChallengerWrong {
					outChan <- Notification{
						Header: header,
						VMID:   vm.address,
						Event:  AsserterTimeoutEvent{},
						TxHash: val.Raw.TxHash,
					}
				} else {
					outChan <- Notification{
						Header: header,
						VMID:   vm.address,
						Event:  ChallengerTimeoutEvent{},
						TxHash: val.Raw.TxHash,
					}
				}
			case val := <-challengeContinuedChan:
				header, err := vm.Client.HeaderByHash(ctx, val.Raw.BlockHash)
				if err != nil {
					errChan <- err
					return
				}
				outChan <- Notification{
					Header: header,
					VMID:   vm.address,
					Event: ContinueChallengeEvent{
						ChallengedAssertion: uint16(val.AssertionIndex.Uint64()),
					},
					TxHash: val.Raw.TxHash,
				}
			case val := <-oneStepProofChan:
				header, err := vm.Client.HeaderByHash(ctx, val.Raw.BlockHash)
				if err != nil {
					errChan <- err
					return
				}
				outChan <- Notification{
					Header: header,
					VMID:   vm.address,
					Event:  OneStepProofEvent{},
					TxHash: val.Raw.TxHash,
				}
			case err := <-headersSub.Err():
				errChan <- err
				return
			case err := <-messageDeliveredSub.Err():
				errChan <- err
				return
			case err := <-dispAssSub.Err():
				errChan <- err
				return
			case err := <-confAssSub.Err():
				errChan <- err
				return
			case err := <-challengeInitiatedSub.Err():
				errChan <- err
				return
			case err := <-challengeBisectedSub.Err():
				errChan <- err
				return
			case err := <-challengeContinuedSub.Err():
				errChan <- err
				return
			case err := <-challengeTimedOutSub.Err():
				errChan <- err
				return
			case err := <-oneStepProofSub.Err():
				errChan <- err
				return
			}
		}
	}()
	return outChan, errChan, nil
}

func (vm *ArbChannel) PendingDisputableAssert(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertion *protocol.Assertion,
) (*types.Receipt, error) {
	tokenNums, amounts, destinations, tokenTypes := hashing.SplitMessages(assertion.OutMsgs)

	dataHashes := make([][32]byte, 0, len(assertion.OutMsgs))
	for _, msg := range assertion.OutMsgs {
		dataHashes = append(dataHashes, msg.Data.Hash())
	}

	tx, err := vm.Tracker.PendingDisputableAssert(
		auth,
		[4][32]byte{
			precondition.BeforeHash,
			precondition.BeforeInbox.Hash(),
			assertion.AfterHash,
			assertion.LogsHash(),
		},
		assertion.NumSteps,
		precondition.TimeBounds,
		tokenTypes,
		dataHashes,
		tokenNums,
		amounts,
		destinations,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, tx.Hash())
}

func (vm *ArbitrumVM) ConfirmDisputableAsserted(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertion *protocol.Assertion,
) (*types.Receipt, error) {
	tokenNums, amounts, destinations, tokenTypes := hashing.SplitMessages(assertion.OutMsgs)

	var messageData bytes.Buffer
	for _, msg := range assertion.OutMsgs {
		err := value.MarshalValue(msg.Data, &messageData)
		if err != nil {
			return nil, err
		}
	}

	tx, err := vm.Tracker.ConfirmDisputableAsserted(
		auth,
		precondition.Hash(),
		assertion.AfterHash,
		assertion.NumSteps,
		tokenTypes,
		messageData.Bytes(),
		tokenNums,
		amounts,
		destinations,
		assertion.LogsHash(),
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, tx.Hash())
}

func (vm *ArbitrumVM) InitiateChallenge(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
) (*types.Receipt, error) {
	var preAssHash [32]byte
	copy(preAssHash[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(precondition.Hash()),
		solsha3.Bytes32(assertion.Hash()),
	))
	tx, err := vm.Tracker.InitiateChallenge(
		auth,
		preAssHash,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, tx.Hash())
}

func (vm *ArbitrumVM) BisectAssertion(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertions []*protocol.AssertionStub,
) (*types.Receipt, error) {
	afterHashAndMessageAndLogsBisections := make([][32]byte, 0, len(assertions)*3+2)
	totalMessageAmounts := make([]*big.Int, 0)
	totalSteps := uint32(0)
	for _, assertion := range assertions {
		afterHashAndMessageAndLogsBisections = append(afterHashAndMessageAndLogsBisections, assertion.AfterHash)
		totalMessageAmounts = append(totalMessageAmounts, assertion.TotalVals...)
		totalSteps += assertion.NumSteps
	}
	afterHashAndMessageAndLogsBisections = append(afterHashAndMessageAndLogsBisections, assertions[0].FirstMessageHash)
	for _, assertion := range assertions {
		afterHashAndMessageAndLogsBisections = append(afterHashAndMessageAndLogsBisections, assertion.LastMessageHash)
	}
	afterHashAndMessageAndLogsBisections = append(afterHashAndMessageAndLogsBisections, assertions[0].FirstLogHash)
	for _, assertion := range assertions {
		afterHashAndMessageAndLogsBisections = append(afterHashAndMessageAndLogsBisections, assertion.LastLogHash)
	}
	tokenTypes, amounts := precondition.BeforeBalance.GetTypesAndAmounts()
	tx, err := vm.Challenge.BisectAssertion(
		auth,
		vm.address,
		[2][32]byte{
			precondition.BeforeHash,
			precondition.BeforeInbox.Hash(),
		},
		afterHashAndMessageAndLogsBisections,
		totalMessageAmounts,
		totalSteps,
		precondition.TimeBounds,
		tokenTypes,
		amounts,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, tx.Hash())
}

func (vm *ArbitrumVM) ContinueChallenge(
	auth *bind.TransactOpts,
	assertionToChallenge uint16,
	preconditions []*protocol.Precondition,
	assertions []*protocol.AssertionStub,
) (*types.Receipt, error) {
	tree := buildBisectionTree(preconditions, assertions)
	tx, err := vm.Challenge.ContinueChallenge(
		auth,
		vm.address,
		big.NewInt(int64(assertionToChallenge)),
		tree.GetProofFlat(int(assertionToChallenge)),
		tree.GetRoot(),
		tree.GetNode(int(assertionToChallenge)),
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, tx.Hash())
}

func (vm *ArbitrumVM) OneStepProof(
	auth *bind.TransactOpts,
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
	proof []byte,
) (*types.Receipt, error) {
	tokenTypes, amounts := precondition.BeforeBalance.GetTypesAndAmounts()
	tx, err := vm.Challenge.OneStepProof(
		auth,
		vm.address,
		[2][32]byte{precondition.BeforeHash, precondition.BeforeInbox.Hash()},
		precondition.TimeBounds,
		tokenTypes,
		amounts,
		[5][32]byte{
			assertion.AfterHash,
			assertion.FirstMessageHash,
			assertion.LastMessageHash,
			assertion.FirstLogHash,
			assertion.LastLogHash,
		},
		assertion.TotalVals,
		proof,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, tx.Hash())
}

func (vm *ArbitrumVM) AsserterTimedOutChallenge(
	auth *bind.TransactOpts,
) (*types.Receipt, error) {
	tx, err := vm.Challenge.AsserterTimedOut(
		auth,
		vm.address,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, tx.Hash())
}

func (vm *ArbitrumVM) ChallengerTimedOutChallenge(
	auth *bind.TransactOpts,
) (*types.Receipt, error) {
	tx, err := vm.Challenge.ChallengerTimedOut(
		auth,
		vm.address,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, tx.Hash())
}

func (vm *ArbitrumVM) CurrentDeposit(
	auth *bind.CallOpts,
	address common.Address,
) (*big.Int, error) {
	return vm.Tracker.CurrentDeposit(auth, address)
}

func (vm *ArbitrumVM) EscrowRequired(
	auth *bind.CallOpts,
) (*big.Int, error) {
	return vm.Tracker.EscrowRequired(auth)
}

func (vm *ArbitrumVM) IsEnabled(
	auth *bind.CallOpts,
) (bool, error) {
	status, err := vm.Tracker.GetState(auth)
	return status != 0, err
}

func (vm *ArbitrumVM) VerifyVM(
	auth *bind.CallOpts,
	config *valmessage.VMConfiguration,
	machine [32]byte,
) error {
	//code, err := vm.contract.Client.CodeAt(auth.Context, vm.address, nil)
	// Verify that VM has correct code
	vmInfo, err := vm.Tracker.Vm(auth)
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

	owner, err := vm.Tracker.Owner(auth)
	if err != nil {
		return err
	}
	if protocol.NewAddressFromBuf(config.Owner) != owner {
		return errors.New("VM has different owner")
	}
	return nil
}

func buildBisectionTree(preconditions []*protocol.Precondition, assertions []*protocol.AssertionStub) *MerkleTree {
	bisectionHashes := make([][32]byte, 0, len(assertions))
	for i := range assertions {
		bisectionBytes := solsha3.SoliditySHA3(
			solsha3.Bytes32(preconditions[i].Hash()),
			solsha3.Bytes32(assertions[i].Hash()),
		)
		bisectionHash := [32]byte{}
		copy(bisectionHash[:], bisectionBytes)
		bisectionHashes = append(bisectionHashes, bisectionHash)
	}
	return NewMerkleTree(bisectionHashes)
}

func translateBisectionEvent(event *challengemanager.ChallengeManagerBisectedAssertion) []*protocol.AssertionStub {
	bisectionCount := (len(event.AfterHashAndMessageAndLogsBisections) - 2) / 3
	assertions := make([]*protocol.AssertionStub, 0, bisectionCount)
	stepCount := event.TotalSteps / uint32(bisectionCount)
	tokenTypeCount := len(event.TotalMessageAmounts) / bisectionCount
	for i := 0; i < bisectionCount; i++ {
		steps := stepCount
		if uint32(i) < event.TotalSteps%uint32(bisectionCount) {
			steps++
		}
		assertion := &protocol.AssertionStub{
			AfterHash:        event.AfterHashAndMessageAndLogsBisections[i],
			NumSteps:         steps,
			FirstMessageHash: event.AfterHashAndMessageAndLogsBisections[bisectionCount+i],
			LastMessageHash:  event.AfterHashAndMessageAndLogsBisections[bisectionCount+i+1],
			FirstLogHash:     event.AfterHashAndMessageAndLogsBisections[bisectionCount*2+1+i],
			LastLogHash:      event.AfterHashAndMessageAndLogsBisections[bisectionCount*2+2+1],
			TotalVals:        event.TotalMessageAmounts[i*tokenTypeCount : (i+1)*tokenTypeCount],
		}
		assertions = append(assertions, assertion)
	}
	return assertions
}

func translateDisputableAssertionEvent(event *arblauncher.ArbitrumVMPendingDisputableAssertion) (*protocol.Precondition, *protocol.AssertionStub) {
	balanceTracker := protocol.NewBalanceTrackerFromLists(event.TokenTypes, event.Amounts)
	precondition := protocol.NewPrecondition(
		event.Fields[0],
		event.TimeBounds,
		balanceTracker,
		value.NewHashOnlyValue(event.Fields[1], 1),
	)
	assertion := &protocol.AssertionStub{AfterHash: event.Fields[2], NumSteps: event.NumSteps, LastMessageHash: event.LastMessageHash, TotalVals: event.Amounts}
	return precondition, assertion
}
