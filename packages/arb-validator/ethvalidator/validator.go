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

package ethvalidator

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethconnection"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type Validator struct {
	key *ecdsa.PrivateKey

	// Not in thread, but internal only
	serverAddress string
	arbAddresses  ethconnection.ArbAddresses
	client        *ethclient.Client

	*ethconnection.ArbLauncher
	*ethconnection.PendingInbox
	auth *bind.TransactOpts
}

func NewValidator(
	key *ecdsa.PrivateKey,
	connectionInfo ethconnection.ArbAddresses,
	ethURL string,
) (*Validator, error) {
	auth := bind.NewKeyedTransactor(key)

	client, err := ethclient.Dial(ethURL)
	if err != nil {
		return nil, err
	}

	vmCreator, err := ethconnection.NewArbLauncher(common.HexToAddress(connectionInfo.VMCreatorAddress), client)
	if err != nil {
		return nil, err
	}

	pendingInbox, err := ethconnection.NewPendingInbox(common.HexToAddress(connectionInfo.GlobalPendingInbox), client)
	if err != nil {
		return nil, err
	}

	return &Validator{
		key:           key,
		serverAddress: ethURL,
		arbAddresses:  connectionInfo,
		client:        client,
		ArbLauncher:   vmCreator,
		PendingInbox:  pendingInbox,
		auth:          auth,
	}, nil
}

func (val *Validator) NewFollower(
	name string,
	machine machine.Machine,
	config *valmessage.VMConfiguration,
	challengeEverything bool,
	maxCallSteps int32,
	maxStepsUnanSteps int32,
	coordinatorURL string,
) (*ValidatorFollower, error) {
	return NewValidatorFollower(
		name,
		val,
		machine,
		config,
		challengeEverything,
		maxCallSteps,
		maxStepsUnanSteps,
		coordinatorURL,
	)
}

func (val *Validator) NewCoordinator(
	name string,
	vmID common.Address,
	machine machine.Machine,
	config *valmessage.VMConfiguration,
	challengeEverything bool,
	maxCallSteps int32,
	maxStepsUnanSteps int32,
) (*ValidatorCoordinator, error) {
	return NewCoordinator(
		name,
		val,
		vmID,
		machine,
		config,
		challengeEverything,
		maxCallSteps,
		maxStepsUnanSteps,
	)
}

func (val *Validator) LatestHeader(ctx context.Context) (*types.Header, error) {
	return val.client.HeaderByNumber(ctx, nil)
}

func (val *Validator) makeAuth(ctx context.Context) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:     val.auth.From,
		Nonce:    val.auth.Nonce,
		Signer:   val.auth.Signer,
		Value:    val.auth.Value,
		GasPrice: val.auth.GasPrice,
		GasLimit: 0,
		Context:  ctx,
	}
}

func (val *Validator) Address() common.Address {
	return crypto.PubkeyToAddress(val.key.PublicKey)
}

func (val *Validator) Sign(msgHash []byte) ([]byte, error) {
	data := solsha3.SoliditySHA3WithPrefix(solsha3.Bytes32(msgHash))
	return crypto.Sign(data, val.key)
}

func EthSigToPub(hash, sig []byte) (*ecdsa.PublicKey, error) {
	data := solsha3.SoliditySHA3WithPrefix(solsha3.Bytes32(hash))
	return crypto.SigToPub(data, sig)
}

func (val *Validator) GetTokenBalance(
	user common.Address,
	tokenContract common.Address,
) (*big.Int, error) {
	auth := &bind.CallOpts{
		Pending: false,
		From:    val.auth.From,
		Context: val.auth.Context,
	}
	amt, err := val.PendingInbox.GetTokenBalance(auth, user, tokenContract)
	return amt, err
}

func (val *Validator) LaunchChannel(
	ctx context.Context,
	config *valmessage.VMConfiguration,
	vmState [32]byte,
) (common.Address, error) {
	return val.ArbLauncher.LaunchChannel(
		val.makeAuth(ctx),
		config,
		vmState,
	)
}

func (val *Validator) LaunchChain(
	ctx context.Context,
	config *valmessage.VMConfiguration,
	vmState [32]byte,
) (common.Address, error) {
	return val.ArbLauncher.LaunchChain(
		val.makeAuth(ctx),
		config,
		vmState,
	)
}
