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
	"errors"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge/channellauncher"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/hashing"
)

type ArbChannel struct {
	*ArbitrumVM
	contract *channellauncher.ArbChannel
}

func NewArbChannel(address common.Address, client *ethclient.Client) (*ArbChannel, error) {
	arbVM, err := NewArbitrumVM(address, client)
	if err != nil {
		return nil, err
	}
	channel := &ArbChannel{ArbitrumVM: arbVM}
	err = channel.setupContracts()
	return channel, err
}

func (vm *ArbChannel) setupContracts() error {
	trackerContract, err := channellauncher.NewArbChannel(vm.address, vm.Client)
	if err != nil {
		return errors2.Wrap(err, "Failed to connect to ArbChannel")
	}
	vm.contract = trackerContract
	return nil
}

func (vm *ArbChannel) StartConnection(ctx context.Context) error {
	if err := vm.ArbitrumVM.StartConnection(ctx); err != nil {
		return err
	}
	if err := vm.setupContracts(); err != nil {
		return err
	}
	start := uint64(0)
	watch := &bind.WatchOpts{
		Context: ctx,
		Start:   &start,
	}

	unanAssChan := make(chan *channellauncher.ArbChannelFinalizedUnanimousAssertion)
	unanAssSub, err := vm.contract.WatchFinalizedUnanimousAssertion(watch, unanAssChan)
	if err != nil {
		return err
	}

	unanPropChan := make(chan *channellauncher.ArbChannelPendingUnanimousAssertion)
	unanPropSub, err := vm.contract.WatchPendingUnanimousAssertion(watch, unanPropChan)
	if err != nil {
		return err
	}

	unanConfChan := make(chan *channellauncher.ArbChannelConfirmedUnanimousAssertion)
	unanConfSub, err := vm.contract.WatchConfirmedUnanimousAssertion(watch, unanConfChan)
	if err != nil {
		return err
	}

	go func() {
		defer unanAssSub.Unsubscribe()
		defer unanConfSub.Unsubscribe()
		defer unanPropSub.Unsubscribe()

		for {
			select {
			case <-ctx.Done():
				break
			case val := <-unanAssChan:
				header, err := vm.Client.HeaderByHash(ctx, val.Raw.BlockHash)
				if err != nil {
					vm.ErrChan <- err
					return
				}
				vm.OutChan <- Notification{
					Header: header,
					VMID:   vm.address,
					Event: FinalizedUnanimousAssertEvent{
						UnanHash: val.UnanHash,
					},
					TxHash: val.Raw.TxHash,
				}
			case val := <-unanPropChan:
				header, err := vm.Client.HeaderByHash(ctx, val.Raw.BlockHash)
				if err != nil {
					vm.ErrChan <- err
					return
				}
				vm.OutChan <- Notification{
					Header: header,
					VMID:   vm.address,
					Event: PendingUnanimousAssertEvent{
						UnanHash:    val.UnanHash,
						SequenceNum: val.SequenceNum,
					},
					TxHash: val.Raw.TxHash,
				}
			case val := <-unanConfChan:
				header, err := vm.Client.HeaderByHash(ctx, val.Raw.BlockHash)
				if err != nil {
					vm.ErrChan <- err
					return
				}
				vm.OutChan <- Notification{
					Header: header,
					VMID:   vm.address,
					Event: ConfirmedUnanimousAssertEvent{
						SequenceNum: val.SequenceNum,
					},
					TxHash: val.Raw.TxHash,
				}
			case err := <-unanAssSub.Err():
				vm.ErrChan <- err
				return
			case err := <-unanPropSub.Err():
				vm.ErrChan <- err
				return
			case err := <-unanConfSub.Err():
				vm.ErrChan <- err
				return
			}
		}
	}()
	return nil
}

func (vm *ArbChannel) IncreaseDeposit(
	auth *bind.TransactOpts,
	amount *big.Int,
) (*types.Receipt, error) {
	call := &bind.TransactOpts{
		From:     auth.From,
		Nonce:    auth.Nonce,
		Signer:   auth.Signer,
		Value:    amount,
		GasPrice: auth.GasPrice,
		GasLimit: 100000,
		Context:  auth.Context,
	}
	tx, err := vm.contract.IncreaseDeposit(call)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, tx.Hash(), "IncreaseDeposit")
}

func (vm *ArbChannel) FinalizedUnanimousAssert(
	auth *bind.TransactOpts,
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
	signatures [][]byte,
) (*types.Receipt, error) {
	messageData, tokenNums, amounts, destinations, tokenTypes := hashing.SplitMessages(assertion.OutMsgs)

	tx, err := vm.contract.FinalizedUnanimousAssert(
		auth,
		assertion.AfterHash,
		newInboxHash,
		tokenTypes,
		messageData,
		tokenNums,
		amounts,
		destinations,
		assertion.LogsHash(),
		sigsToBlock(signatures),
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, tx.Hash(), "FinalizedUnanimousAssert")
}

func (vm *ArbChannel) PendingUnanimousAssert(
	auth *bind.TransactOpts,
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
	sequenceNum uint64,
	signatures [][]byte,
) (*types.Receipt, error) {
	var unanRest [32]byte
	copy(unanRest[:], hashing.UnanimousAssertPartialPartialHash(
		newInboxHash,
		assertion,
	))

	stub := assertion.Stub()

	tx, err := vm.contract.PendingUnanimousAssert(
		auth,
		unanRest,
		sequenceNum,
		stub.LastMessageHashValue(),
		stub.LastLogHashValue(),
		sigsToBlock(signatures),
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, tx.Hash(), "PendingUnanimousAssert")
}

func (vm *ArbChannel) ConfirmUnanimousAsserted(
	auth *bind.TransactOpts,
	newInboxHash [32]byte,
	assertion *protocol.Assertion,
) (*types.Receipt, error) {
	messageData, tokenNums, amounts, destinations, tokenTypes := hashing.SplitMessages(assertion.OutMsgs)

	tx, err := vm.contract.ConfirmUnanimousAsserted(
		auth,
		assertion.AfterHash,
		newInboxHash,
		tokenTypes,
		messageData,
		tokenNums,
		amounts,
		destinations,
	)
	if err != nil {
		return nil, err
	}
	return waitForReceipt(auth.Context, vm.Client, tx.Hash(), "ConfirmUnanimousAsserted")
}

func (vm *ArbChannel) VerifyVM(
	auth *bind.CallOpts,
	config *valmessage.VMConfiguration,
	machine [32]byte,
) error {
	err := vm.ArbitrumVM.VerifyVM(auth, config, machine)
	validators := make([]common.Address, 0, len(config.AssertKeys))
	for _, assertKey := range config.AssertKeys {
		validators = append(validators, protocol.NewAddressFromBuf(assertKey))
	}
	correctValidators, err := vm.contract.IsValidatorList(auth, validators)
	if err != nil {
		return err
	}
	if !correctValidators {
		return errors.New("VM has different validator list")
	}
	return nil
}

func sigsToBlock(signatures [][]byte) []byte {
	sigData := make([]byte, 0, len(signatures)*65)
	for _, sig := range signatures {
		sigData = append(sigData, sig[:64]...)
		v := uint8(int(sig[64]))
		if v < 27 {
			v += 27
		}
		sigData = append(sigData, v)
	}
	return sigData
}
