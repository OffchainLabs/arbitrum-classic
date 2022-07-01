/*
 * Copyright 2021, Offchain Labs, Inc.
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

package web3

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethersphere/bee/pkg/crypto/eip712"
	"github.com/pkg/errors"
)

type Accounts struct {
	srv         *Server
	fwdSrv      *ForwarderServer
	addresses   []common.Address
	privateKeys map[common.Address]*ecdsa.PrivateKey
	signer      types.Signer
	nonMutating bool
}

func NewAccounts(ethServer *Server, privateKeys []*ecdsa.PrivateKey, nonMutating bool) *Accounts {
	keys := make(map[common.Address]*ecdsa.PrivateKey)
	addresses := make([]common.Address, 0, len(privateKeys))
	for _, privKey := range privateKeys {
		addr := crypto.PubkeyToAddress(privKey.PublicKey)
		keys[addr] = privKey
		addresses = append(addresses, addr)
	}
	return &Accounts{
		srv:         ethServer,
		addresses:   addresses,
		privateKeys: keys,
		signer:      types.NewEIP155Signer(new(big.Int).SetUint64(uint64(ethServer.ChainId()))),
		nonMutating: nonMutating,
	}
}

func (s *Accounts) Accounts() []common.Address {
	return s.addresses
}

type SendTransactionArgs struct {
	From     *common.Address `json:"from"`
	To       *common.Address `json:"to"`
	Gas      *hexutil.Uint64 `json:"gas"`
	GasPrice *hexutil.Big    `json:"gasPrice"`
	Value    *hexutil.Big    `json:"value"`
	Nonce    *hexutil.Uint64 `json:"nonce"`
	Data     *hexutil.Bytes  `json:"data"`
}

func (s *Accounts) SendTransaction(ctx context.Context, args *SendTransactionArgs) (common.Hash, error) {
	if s.nonMutating {
		return common.Hash{}, errors.New(nonMutatingModeError)
	}

	sender := s.addresses[0]
	if args.From != nil {
		sender = *args.From
	}
	privKey, ok := s.privateKeys[sender]
	if !ok {
		return common.Hash{}, errors.New("sender does not have unlocked wallet")
	}

	var nonce uint64
	if args.Nonce != nil {
		nonce = uint64(*args.Nonce)
	} else {
		pending := rpc.PendingBlockNumber
		block := rpc.BlockNumberOrHash{BlockNumber: &pending}
		rawNonce, err := s.fwdSrv.GetTransactionCount(ctx, &sender, block)
		if err != nil {
			return common.Hash{}, err
		}
		nonce = uint64(rawNonce)
	}
	gas := uint64(2000000)
	if args.Gas != nil {
		gas = uint64(*args.Gas)
	}
	val := (*big.Int)(args.Value)
	if val == nil {
		val = big.NewInt(0)
	}
	var data []byte
	if args.Data != nil {
		data = *args.Data
	}
	gasPrice := (*big.Int)(args.GasPrice)
	if gasPrice == nil {
		gasPriceRaw, err := s.srv.GasPrice(ctx)
		if err != nil {
			return [32]byte{}, err
		}
		gasPrice = (*big.Int)(gasPriceRaw)
	}
	var tx *types.Transaction
	if args.To != nil {
		tx = types.NewTransaction(
			nonce,
			*args.To,
			val,
			gas,
			gasPrice,
			data,
		)
	} else {
		tx = types.NewContractCreation(
			nonce,
			val,
			gas,
			gasPrice,
			data,
		)
	}
	signedTx, err := types.SignTx(tx, s.signer, privKey)
	if err != nil {
		return [32]byte{}, err
	}

	if err := s.srv.srv.SendTransaction(ctx, signedTx); err != nil {
		return [32]byte{}, err
	}
	return signedTx.Hash(), nil
}

func (s *Accounts) Sign(account common.Address, data hexutil.Bytes) (hexutil.Bytes, error) {
	dataHash := accounts.TextHash(data)
	sig, err := s.signHash(account, dataHash)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func (s *Accounts) SignTypedData_v4(account common.Address, typedData string) (hexutil.Bytes, error) {
	var typed eip712.TypedData
	err := json.Unmarshal([]byte(typedData), &typed)
	if err != nil {
		return nil, errors.Wrap(err, "json failed")
	}
	data, err := eip712.EncodeForSigning(&typed)
	if err != nil {
		return nil, err
	}
	dataHash := crypto.Keccak256(data)
	sig, err := s.signHash(account, dataHash)
	if err != nil {
		return nil, err
	}
	return sig, err
}

func (s *Accounts) signHash(account common.Address, dataHash []byte) (hexutil.Bytes, error) {
	privKey, ok := s.privateKeys[account]
	if !ok {
		return nil, errors.New("signer does not have unlocked wallet")
	}
	sig, err := crypto.Sign(dataHash, privKey)
	if err != nil {
		return nil, err
	}
	sig[64] += 27
	return sig, nil
}

type PersonalAccounts struct {
	privateKeys map[common.Address]*ecdsa.PrivateKey
}

func NewPersonalAccounts(privateKeys []*ecdsa.PrivateKey) *PersonalAccounts {
	keys := make(map[common.Address]*ecdsa.PrivateKey)
	for _, privKey := range privateKeys {
		addr := crypto.PubkeyToAddress(privKey.PublicKey)
		keys[addr] = privKey
	}
	return &PersonalAccounts{
		privateKeys: keys,
	}
}

func (s *PersonalAccounts) Sign(data hexutil.Bytes, account common.Address, _ *hexutil.Bytes) (hexutil.Bytes, error) {
	// Password ignored
	privKey, ok := s.privateKeys[account]
	if !ok {
		return nil, errors.New("signer does not have unlocked wallet")
	}
	sig, err := crypto.Sign(accounts.TextHash(data), privKey)
	if err != nil {
		return nil, err
	}
	sig[64] += 27
	return sig, nil
}
