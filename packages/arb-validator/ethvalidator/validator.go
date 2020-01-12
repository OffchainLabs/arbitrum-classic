/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type Validator struct {
	key *ecdsa.PrivateKey

	// Not in thread, but internal only
	serverAddress string
	arbAddresses  ethbridge.ArbAddresses
	Client        arbbridge.ArbClient

	arbbridge.ArbFactory
	arbbridge.PendingInbox
	auth *bind.TransactOpts
}

func NewValidator(
	key *ecdsa.PrivateKey,
	connectionInfo ethbridge.ArbAddresses,
	ethURL string,
) (*Validator, error) {
	auth := bind.NewKeyedTransactor(key)

	client, err := ethbridge.NewEthClient(ethURL)
	if err != nil {
		return nil, err
	}

	arbFactory, err := client.NewArbFactory(common.HexToAddress(connectionInfo.ArbFactory))
	if err != nil {
		return nil, err
	}

	pendingInbox, err := client.NewPendingInbox(common.HexToAddress(connectionInfo.GlobalPendingInbox))
	if err != nil {
		return nil, err
	}

	return &Validator{
		key:           key,
		serverAddress: ethURL,
		arbAddresses:  connectionInfo,
		Client:        client,
		ArbFactory:    arbFactory,
		PendingInbox:  pendingInbox,
		auth:          auth,
	}, nil
}

func (val *Validator) LatestHeader(ctx context.Context) (*types.Header, error) {
	return val.Client.HeaderByNumber(ctx, nil)
}

func (val *Validator) MakeAuth(ctx context.Context) *bind.TransactOpts {
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

func (val *Validator) CreateRollup(
	ctx context.Context,
	config structures.ChainParams,
	vmState [32]byte,
	owner common.Address,
) (common.Address, error) {
	return val.ArbFactory.CreateRollup(
		val.MakeAuth(ctx),
		vmState,
		config,
		owner,
	)
}
